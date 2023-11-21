package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var (
	Token     string
	BotPrefix string
	config    *configStruct
)

type configStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
}

func ReadConfig() error {
	fmt.Println("Read config ...")
	file, err := os.ReadFile("./config.json")
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println(string(file))
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	Token = config.Token
	BotPrefix = config.BotPrefix
	return nil
}
