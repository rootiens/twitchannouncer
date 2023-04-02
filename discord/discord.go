package discord

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

var discord *discordgo.Session

func InitiateDiscord() {
	session, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
	if err != nil {
		panic(err)
	}

	discord = session
}

func SendMessage(message Message) {
	mes, err := discord.ChannelMessageSend(message.ChannelID, message.Content)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(mes)
}
