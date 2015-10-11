package tvmaze

import (
	"fmt"
	"time"
)

// Episode wraps a TV Maze episode resource
type Episode struct {
	ID      int
	Name    string
	Season  int
	Number  int
	AirDate time.Time `json:"airstamp"`
	Runtime int
	Summary string
}

// GetEpisodes finds all episodes for the given show
func (c *Client) GetEpisodes(s *Show) (episodes []Episode, err error) {
	path := fmt.Sprintf("/shows/%d/episodes", s.ID)

	if err := c.get(path, &episodes); err != nil {
		return nil, err
	}

	return episodes, nil
}
