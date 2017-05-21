package main

import (
	"log"
	"os/exec"

	"github.com/zmb3/spotify"
)

func searchSpotify(url string) {
	//Query Spotify API
	results, err := spotify.Search(url, spotify.SearchTypePlaylist|spotify.SearchTypeAlbum)
	if err != nil {
		log.Fatal(err)
	}
	//Loop through album results
	//Opens Albums into default external browser
	if results.Albums != nil {
		for _, item := range results.Albums.Albums {
			for _, val := range item.ExternalURLs {
				exec.Command("open", val).Start()
			}
		}

	}
}
