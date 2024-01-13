package config

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
)

// BedrockConfig represents a defualt bedrock client

var BRc *bedrockruntime.Client

// Initialize the configuration and create aclient from the cfg
func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	BRc = bedrockruntime.NewFromConfig(cfg)
}

var PORT string = os.Getenv("PORT")

func Setup() {

	if PORT == "" {
		log.Fatal("no listen port defined, set PORT :XX")
	}
}
