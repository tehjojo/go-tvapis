package tvmaze

import (
	"fmt"
	"net/url"
	"time"
)

// Show wraps a TV Maze show object
type Show struct {
	ID        int
	Name      string
	Type      string
	Genres    []string
	Status    string
	Runtime   int
	Premiered date
	Summary   string
	Network   network
	Embeds    struct {
		Episodes []Episode
	} `json:"_embedded"`
	Remotes map[string]int `json:"externals"`
}

// GetTitle return the show title
func (s *Show) GetTitle() string {
	return s.Name
}

// GetDescription returns a summary of the show
func (s *Show) GetDescription() string {
	return s.Summary
}

// GetNetwork returns the network that currently broadcasts the show
func (s *Show) GetNetwork() string {
	return s.Network.Name
}

// GetFirstAired return the time the first episode was aired
func (s *Show) GetFirstAired() time.Time {
	return s.Premiered.Time
}

// GetTVRageID returns the show's ID on tvrage.com
func (s *Show) GetTVRageID() int {
	return s.Remotes["tvrage"]
}

// FindShow finds all matches for a given search string
func (c *Client) FindShow(name string) (s []*Show, err error) {
	path := fmt.Sprintf("/search/shows?q=%s", url.QueryEscape(name))

	if err := c.get(path, &s); err != nil {
		return nil, err
	}

	return s, nil
}

// GetShow finds all matches for a given search string
func (c *Client) GetShow(name string) (s []*Show, err error) {
	path := fmt.Sprintf("/singlesearch/shows?q=%s", url.QueryEscape(name))

	if err := c.get(path, &s); err != nil {
		return nil, err
	}

	return s, nil
}

// RefreshShow refreshes a show from the server
func (c *Client) RefreshShow(show *Show) (err error) {
	path := fmt.Sprintf("/shows/%d", show.ID)

	if err := c.get(path, &show); err != nil {
		return err
	}

	return nil
}
