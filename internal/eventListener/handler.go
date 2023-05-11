package eventlistener

import (
	"fmt"
	"log"
	"time"

	"github.com/deepfake-db/contracts"
	"github.com/deepfake-db/internal/config"
	"github.com/deepfake-db/internal/models"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
)

var (
	imageAsserted          string = crypto.Keccak256Hash([]byte("ImageAsserted(bytes32,string,address,bytes32)")).Hex()
	imageAssertionResolved string = crypto.Keccak256Hash([]byte("ImageAssertionResolved(bytes32,string,address,bytes32)")).Hex()
)

func HandleEvents(client *ethclient.Client, l types.Log) {
	eventSig := l.Topics[0].Hex()

	fmt.Println("event logged: ", eventSig)
	switch eventSig {
	case imageAsserted:
		go handleImageAssertion(config.DB, l)
	case imageAssertionResolved:
		go handleImageAssertionResolved(config.DB, l)
	}
}

func handleImageAssertion(db *gorm.DB, l types.Log) {
	var id [32]byte
	copy(id[:], l.Topics[1].Bytes()[:32])
	idHex := hexutil.Encode(id[:])
	fmt.Println("ID: ", idHex)

	asserter := l.Topics[2].Bytes()[12:]
	asserterAddr := hexutil.Encode(asserter)
	fmt.Println("asserter addr: ", asserterAddr)

	var assertionId [32]byte
	copy(assertionId[:], l.Topics[3].Bytes()[:32])
	assertionIdHex := hexutil.Encode(assertionId[:])
	fmt.Println("assertion id: ", assertionIdHex)

	t := time.Now().Unix()

	resolvableAt := t + config.DbConfig.AssertionLiveness

	// getting unindexed event execution timeframe
	f, err := contracts.NewDeepFakeDetectorFilterer(common.HexToAddress(config.DbConfig.DeepfakeContractAddress), config.Client)
	if err != nil {
		log.Fatal(err)
	}

	imgAssertedEvent, err := f.ParseImageAsserted(l)
	if err != nil {
		log.Fatal(err)
	}

	imgAssertion := models.ImageAssertion{
		ImageId:      idHex,
		ImageUrl:     imgAssertedEvent.ImageUrl,
		Asserter:     asserterAddr,
		AssertionId:  assertionIdHex,
		Resolved:     false,
		Seen:         false,
		LoggedAt:     t,
		ResolvableAt: resolvableAt,
	}
	db.Create(&imgAssertion)
}

func handleImageAssertionResolved(db *gorm.DB, l types.Log) {
	var id [32]byte
	copy(id[:], l.Topics[3].Bytes()[:32])
	idHex := hexutil.Encode(id[:])

	// get event and opdate in DB
	var localAssertion models.ImageAssertion
	db.First(&localAssertion, "assertion_id = ?", idHex)

	db.Model(&localAssertion).Updates(models.ImageAssertion{Resolved: true})
}
