package main

import (
	"fmt"

	"github.com/souravdey425/discord/bot"
	"github.com/souravdey425/discord/config"
)

func main() {

	// http://discordapp.com/oauth2/authorize?&client_id=1164004561489514556&scope=bot
	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err)
	}
	bot.Start()
	<-make(chan struct{})
	return
}
