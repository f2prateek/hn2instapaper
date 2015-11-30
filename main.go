package main

import (
	"fmt"
	"log"

	"github.com/f2prateek/hn-to-instapaper/hn"
	"github.com/f2prateek/hn-to-instapaper/instapaper"
	"github.com/tj/docopt"
)

const (
	usage = `hn2instapaper.

Usage:
  hn2instapaper <username> <password>
  hn2instapaper -h | --help
  hn2instapaper --version

Options:
  -h --help     Show this screen.
  --version     Show version.`
)

func main() {
	arguments, err := docopt.Parse(usage, nil, true, "Naval Fate 2.0", false)
	if err != nil {
		log.Fatal(err)
	}

	hnClient := hn.New()
	instapaperClient := instapaper.New(arguments["<username>"].(string), arguments["<password>"].(string))

	stories, err := hnClient.TopStories()
	if err != nil {
		log.Fatal(err)
	}

	for i, id := range stories {
		story, err := hnClient.GetPost(id)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Adding", *story.Title)

		result, err := instapaperClient.Add(instapaper.AddParams{
			URL:   *story.URL,
			Title: story.Title,
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(i, result.BookmarkID)
	}
}
