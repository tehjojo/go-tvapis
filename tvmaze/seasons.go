package tvmaze

import (
	"fmt"
	"time"
)

type Season struct {
	ID           int    `json:"id"`
	URL          string `json:"url"`
	Number       int    `json:"number"`
	Name         string `json:"name"`
	EpisodeOrder int    `json:"episodeOrder"`
	PremiereDate Date   `json:"premiereDate"`
	EndDate      Date   `json:"endDate"`
	Network      struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Country struct {
			Name     string `json:"name"`
			Code     string `json:"code"`
			Timezone string `json:"timezone"`
		} `json:"country"`
		OfficialSite interface{} `json:"officialSite"`
	} `json:"network"`
	WebChannel interface{} `json:"webChannel"`
	Image      struct {
		Medium   string `json:"medium"`
		Original string `json:"original"`
	} `json:"image"`
	Summary interface{} `json:"summary"`
	Links   struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
}

//IsFuture returns whether the season is in the future.
// TVMaze often tracks seasons that have been ordered, but not started
func (s Season) IsFuture() bool {
	//Safety check. Often future seasons dont have a date set.
	if s.PremiereDate.IsZero() {
		return true
	}
	return s.PremiereDate.After(time.Now())
}

//getSeasons retrieves the seasons for a TV Show
func (s Show) getSeasons() (seasons []Season, err error) {
	url := baseURLWithPath(fmt.Sprintf("shows/%d/seasons", s.ID))

	if _, err = s.get(url, &seasons); err != nil {
		return nil, err
	}

	return seasons, nil
}

// GetSeason returns a specific season for a show
func (s Show) GetSeason(seasonNumber int) (*Season, error) {
	for _, season := range s.Seasons {
		if season.Number == seasonNumber {
			return &season, nil
		}
	}
	return nil, fmt.Errorf("Season not found")
}

// GetActiveSeason returns the active season for the date provided
// Note: Web Shows often have a Premiere and End Date set to same date and it wont be retrievable with this function
func (s Show) GetActiveSeason(dt time.Time) (*Season, error) {
	for _, season := range s.Seasons {
		if dt.After(season.PremiereDate.Time) && dt.Before(season.EndDate.Time) {
			return &season, nil
		}
	}

	return nil, fmt.Errorf("no active seasons found for provided date %s", dt)
}

//GetCurrentSeason returns the current season that is not in the future.
func (s Show) GetCurrentSeason() (*Season, error) {
	//Return last season. Assume they are pre-sorted in the API response.
	for i := len(s.Seasons) - 1; i >= 0; i-- {
		if !s.Seasons[i].IsFuture() {
			return &s.Seasons[i], nil
		}
	}

	return nil, fmt.Errorf("no current season found")
}
