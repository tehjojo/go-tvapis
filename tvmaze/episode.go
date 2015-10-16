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
func (c Client) GetEpisodes(s Show) (episodes []Episode, err error) {
	url := baseURLWithPath(fmt.Sprintf("shows/%d/episodes", s.ID))

	if err = c.get(url, &episodes); err != nil {
		return nil, err
	}

	return episodes, nil
}

// GetNextEpisode returns the next un-air episode for the show
func (c Client) GetNextEpisode(s Show) (*Episode, error) {
	url := baseURLWithPathQuery(fmt.Sprintf("shows/%d", s.ID), "embed", "nextepisode")

	var embed embeddedNextEpisode
	if err := c.get(url, &embed); err != nil {
		return nil, err
	}

	if embed.Embedded.NextEpisode.ID == 0 {
		return nil, nil
	}
	return &embed.Embedded.NextEpisode, nil
}

type embeddedNextEpisode struct {
	Embedded struct {
		NextEpisode Episode `json:"nextepisode"`
	} `json:"_embedded"`
}
