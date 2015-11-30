package main

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/f2prateek/hn2instapaper/hn"
	"github.com/f2prateek/hn2instapaper/instapaper"
	"github.com/tj/docopt"
)

const (
	usage = `hn2instapaper.

Save top HN articles to Instapaper.

Usage:
  hn2instapaper <username> <password> [--limit l]
  hn2instapaper -h | --help
  hn2instapaper --version

Options:
  --limit l     Number of articles to save [default: 500].
  -h --help     Show this screen.
  --version     Show version.`

	version = "1.0.0"
)

func main() {
	arguments, err := docopt.Parse(usage, nil, true, version, false)
	check(err)

	hnClient := hn.New()
	instapaperClient := instapaper.New(arguments["<username>"].(string), arguments["<password>"].(string))
	limit, err := strconv.Atoi(arguments["--limit"].(string))
	check(err)

	stories, err := hnClient.TopStories()
	check(err)

	var wg sync.WaitGroup
	for i, id := range stories {
		if i >= limit {
			break
		}

		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			story, err := hnClient.GetPost(id)
			check(err)
			if story.URL == nil {
				fmt.Println("Skipping", *story.Title)
			}

			_, err = instapaperClient.Add(instapaper.AddParams{
				URL:   *story.URL,
				Title: story.Title,
			})
			check(err)
			fmt.Println("Saved", *story.Title)
		}(id)
	}
	wg.Wait()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
