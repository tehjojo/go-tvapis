package tvmaze

import (
	"fmt"
	"strconv"
	"time"
)

// Episode wraps a TV Maze episode resource
type Episode struct {
	ID      int
	Name    string
	Season  int
	Number  int
	AirDate *time.Time `json:"airstamp"`
	Runtime int
	Summary string
}

// GetEpisodes finds all episodes for the given show
func (s Show) GetEpisodes() (episodes []Episode, err error) {
	url := baseURLWithPath(fmt.Sprintf("shows/%d/episodes", s.ID))

	if _, err = s.get(url, &episodes); err != nil {
		return nil, err
	}

	return episodes, nil
}

// GetNextEpisode returns the next un-air episode for the show
func (s Show) GetNextEpisode() (*Episode, error) {
	url := baseURLWithPathQuery(fmt.Sprintf("shows/%d", s.ID), "embed", "nextepisode")

	var embed embeddedNextEpisode
	if _, err := s.get(url, &embed); err != nil {
		return nil, err
	}

	if embed.Embedded.NextEpisode.ID == 0 {
		return nil, nil
	}
	return &embed.Embedded.NextEpisode, nil
}

// GetEpisode returns a specific episode for a show
func (s Show) GetEpisode(season int, episode int) (*Episode, error) {
	url := baseURLWithPathQueries(fmt.Sprintf("shows/%d/episodebynumber", s.ID), map[string]string{
		"season": strconv.Itoa(season),
		"number": strconv.Itoa(episode),
	})

	var epOut Episode
	if _, err := s.get(url, &epOut); err != nil {
		return nil, err
	}
	return &epOut, nil
}

/*
	Backwards compatibility
*/
// GetEpisodes finds all episodes for the given show
func (c Client) GetEpisodes(s Show) (episodes []Episode, err error) {
	fmt.Println("[WARN] (Client).GetEpisodes(Show) is deprecated.")
	return s.GetEpisodes()
}

// GetNextEpisode returns the next un-air episode for the show
func (c Client) GetNextEpisode(s Show) (*Episode, error) {
	fmt.Println("[WARN] (Client).GetNextEpisode(Show) is deprecated.")
	return s.GetNextEpisode()
}

// GetEpisode returns a specific episode for a show
func (c Client) GetEpisode(s Show, season int, episode int) (*Episode, error) {
	fmt.Println("[WARN] (Client).GetEpisode(Show) is deprecated.")
	return s.GetEpisode(season, episode)
}

type embeddedNextEpisode struct {
	Embedded struct {
		NextEpisode Episode `json:"nextepisode"`
	} `json:"_embedded"`
}
