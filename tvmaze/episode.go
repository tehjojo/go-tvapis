package tvmaze

import (
	"log"
	"strconv"
	"time"
)

// Episode wraps a TV Maze episode resource
type Episode struct {
	Id      int
	Name    string
	Season  int
	Number  int
	AirDate time.Time `json:"airstamp"`
	Runtime int
	Summary string
}

// GetEpisodes finds all episodes for the given show
func (c *Client) GetEpisodes(s *Show) (episodes []Episode) {
	path := "/shows/" + strconv.FormatInt(int64(s.Id), 10) + "/episodes"

	if err := c.get(path, &episodes); err != nil {
		log.Fatal(err)
	}

	return episodes
}
