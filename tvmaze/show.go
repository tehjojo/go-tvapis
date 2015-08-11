package tvmaze

import (
	"log"
	"net/url"
	"strconv"
)

// Show wraps a TV Maze show object
type Show struct {
	Id        int
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

// GetTVRageID returns the show's ID on tvrage.com
func (s *Show) GetTVRageID() int {
	return s.Remotes["tvrage"]
}

// LookupShow tries to find a match for the show name and return a Show object if successful
func (c *Client) LookupShow(name string) (s *Show) {
	path := "/singlesearch/shows?q=" + url.QueryEscape(name)

	if err := c.get(path, &s); err != nil {
		log.Fatal(err)
	}

	return s
}

// GetShow retrieves a show by id
func (c *Client) GetShow(id int, withEpisodes bool) (s *Show) {
	path := "/shows/" + strconv.FormatInt(int64(id), 10)

	if withEpisodes {
		path += "?embed=episodes"
	}

	if err := c.get(path, &s); err != nil {
		log.Fatal(err)
	}

	return s
}
