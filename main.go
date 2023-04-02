package main

import (
	"fmt"
	"github.com/rootiens/twitchannouncer/discord"
	"github.com/rootiens/twitchannouncer/twitch"
	"sync"
	"time"
)

func main() {
	streamers, err := twitch.GetStreamers()
	if err != nil {
		panic(err)
	}

	discord.InitiateDiscord()

	var wg sync.WaitGroup

	for {
		twitch.CheckToken()

		for _, streamer := range streamers.Streamers {
			wg.Add(1)
			go func(streamer twitch.StreamerData) {
				defer wg.Done()
				ok, err := twitch.IsStreamerOnline(streamer.TwitchName)
				if err != nil {
					fmt.Println(err)
				}
				if ok {
					fmt.Println(streamer.TwitchName, "is live")
				} else {
					message := discord.Message{
						ChannelID: streamer.DiscordChannel,
						Content:   streamer.TwitchName + "is offline",
					}
					discord.SendMessage(message)

					fmt.Println(streamer.TwitchName, "is offline")
				}
			}(streamer)
		}
		wg.Wait()
		fmt.Println("=======================")
		time.Sleep(60 * time.Second)
	}

}
