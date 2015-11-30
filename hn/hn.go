package hn

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Client wraps the HackerNews API https://github.com/HackerNews/API
type Client struct {
	httpClient *http.Client
}

// New returns a new HN client.
func New() *Client {
	return &Client{
		httpClient: http.DefaultClient,
	}
}

// TopStories returns the ids of the top HN posts.
// https://hacker-news.firebaseio.com/v0/topstories.
func (c *Client) TopStories() ([]int, error) {
	request, err := http.NewRequest("GET", "https://hacker-news.firebaseio.com/v0/topstories.json", nil)
	if err != nil {
		return make([]int, 0), err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return make([]int, 0), err
	}

	var results []int
	err = json.NewDecoder(response.Body).Decode(&results)
	if err != nil {
		return make([]int, 0), err
	}

	return results, nil
}

// Item represents a HN post.
type Item struct {
	By          *string `json:"by"`
	Descendants *int    `json:"descendants"`
	ID          int     `json:"id"`
	Kids        *[]int  `json:"kids"`
	Score       *int    `json:"score"`
	Time        *int    `json:"time"`
	Title       *string `json:"title"`
	Type        *string `json:"type"`
	URL         *string `json:"url"`
}

// GetPost returns the post with the given ID.
func (c *Client) GetPost(id int) (Item, error) {
	url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%s.json", strconv.Itoa(id))
	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return Item{}, err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return Item{}, err
	}

	var result Item
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return Item{}, err
	}

	return result, nil
}
