package config

import (
	"log"

	"github.com/deepfake-db/internal/models"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	NodeURL                 string `mapstructure:"NODE_URL"`
	Port                    int    `mapstructure:"PORT"`
	DeepfakeContractAddress string `mapstructure:"DEEPFAKE_CONTRACT_ADDRESS"`
	AssertionLiveness       int64  `mapstructure:"ASSERTION_LIVENESS"`
}

var DbConfig Config
var Client *ethclient.Client

var DB *gorm.DB

func LoadConfig() {
	viper.AddConfigPath("../../")
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Failed to read env file: ", err)
	}

	err = viper.Unmarshal(&DbConfig)
	if err != nil {
		log.Fatal("Failed to decode env file: ", err)
	}
}

func LoadClient() {
	client, err := ethclient.Dial(DbConfig.NodeURL)
	if err != nil {
		log.Fatal(err)
	}
	Client = client
}

func InitialMigrationDB() {
	db, err := gorm.Open(sqlite.Open("deepfake.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.ImageAssertion{})
	DB = db
}
