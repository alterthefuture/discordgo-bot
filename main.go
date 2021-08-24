package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

const token string = "token-here"

var BotID string

func main() {
	dg, err := discordgo.New("Bot " + token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := dg.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID

	dg.AddHandler(messageHandler)

	err = dg.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running!")

	<-make(chan struct{})
	return

}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == BotID {
		return
	}

	if m.Content == "!ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong!")
	}

	if m.Content == "!hello" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Hello, "+m.Author.Username)
	}

	if m.Content == "!userinfo" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "User: "+m.Author.Username+"\nID: "+m.Author.ID+"\nTag: "+m.Author.Discriminator)
	}
}
