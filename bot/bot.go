package bot

import (
	"discord-bot/config"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var BotId string
var goBot *discordgo.Session

func Start() {

	goBot, err := discordgo.New("Bot " + config.Conf.DiscordToken)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	goBot.AddHandler(messageHandler)
	goBot.Identify.Intents = discordgo.IntentsGuildMessages

	err = goBot.Open()
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("4")
		return
	}
	fmt.Println("Bot is running !")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	switch {
	case m.Content == config.Conf.Prefix+"ping":
		s.ChannelMessageSend(m.ChannelID, "pong")
	case strings.HasPrefix(m.Content, config.Conf.Prefix+"player"):
		s.ChannelMessageSend(m.ChannelID, (PlayerInfo(strings.TrimPrefix(m.Content, "!player#"))).Name)
	case strings.HasPrefix(m.Content, config.Conf.Prefix+"clan"):
		s.ChannelMessageSend(m.ChannelID, (ClanInfo(strings.TrimPrefix(m.Content, "!clan#"))).Description)
	}
}
