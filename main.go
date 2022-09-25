package main

import (
	"akaibu/scrape"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Home(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello World")
	if err != nil {
		fmt.Println(err)
	}
}

func Channels(w http.ResponseWriter, r *http.Request) {
	var response = "{channelData: {channel: true, id: 1}}"
	fmt.Fprintf(w, response)
}

func main() {
	fmt.Println("Main started!")
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for {
			select {
			case <-ticker.C:
				scrape.RunScrape()
			}
		}
	}()

	<-quit
	fmt.Println("Ticker Stopped")
	ticker.Stop()
}
