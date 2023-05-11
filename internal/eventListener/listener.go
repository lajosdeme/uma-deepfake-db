package eventlistener

import (
	"context"
	"fmt"
	"log"

	"github.com/deepfake-db/internal/config"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func Listen() {
	detectorAddr := common.HexToAddress(config.DbConfig.DeepfakeContractAddress)
	ooAddr := common.HexToAddress(config.DbConfig.OOV3Address)
	fmt.Println("OO ADDR: ", ooAddr)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{detectorAddr, ooAddr},
	}

	logs := make(chan types.Log)

	sub, err := config.Client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			fmt.Println(err)
		case vLog := <-logs:
			fmt.Println(vLog)
			HandleEvents(config.Client, vLog)
		}
	}
}
