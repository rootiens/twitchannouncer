package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
    streamers, err := GetStreamers()
    if err != nil {
        panic(err)
    }

	var wg sync.WaitGroup

	for {
		CheckToken()

		for _, streamer := range streamers.Streamers {
			wg.Add(1)
			go func(name string) {
				defer wg.Done()
				ok, err := IsStreamerOnline(name)
				if err != nil {
					fmt.Println(err)
				}
				if ok {
					fmt.Println(name, "is live")
				} else {
					fmt.Println(name, "is offline")
				}
			}(streamer.TwitchName)
		}
		wg.Wait()
		fmt.Println("=======================")
		time.Sleep(60 * time.Second)
	}

}
