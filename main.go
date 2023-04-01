package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	streamers := []string{"rootiens", "theprimeagen"}

	var wg sync.WaitGroup

	for {
		CheckToken()

		for _, streamer := range streamers {
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
			}(streamer)
		}
		wg.Wait()
		fmt.Println("=======================")
		time.Sleep(60 * time.Second)
	}

}
