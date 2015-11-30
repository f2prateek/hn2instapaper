package main

import (
	"fmt"
	"log"

	"github.com/f2prateek/hn-to-instapaper/hn"
)

func main() {
	hnClient := hn.New()

	stories, err := hnClient.TopStories()
	if err != nil {
		log.Fatal(err)
	}

	for i, id := range stories {
		story, err := hnClient.GetPost(id)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(i, *story.Title)
	}
}
