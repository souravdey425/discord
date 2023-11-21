package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/souravdey425/discord/config"
)

var BotId string
var gobot *discordgo.Session

func Start() {
	gobot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println(err)
	}
	user, err := gobot.User("@me")
	fmt.Println(user)
	if err != nil {
		fmt.Println(err.Error())
	}
	BotId = user.ID

	gobot.AddHandler(messageHandler)
	err = gobot.Open()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Bot is running")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotId {
		return
	}
	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
	}
}
