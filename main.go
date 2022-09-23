package main

import (
	"fmt"
	"go-sandbox/scrape"
	"net/http"
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

func Initialize() {

}

func main() {
	fmt.Println("Main started!")
	scrape.RunScrape()
}
