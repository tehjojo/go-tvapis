package tvmaze

import "fmt"

type Season struct {
	ID           int    `json:"id"`
	URL          string `json:"url"`
	Number       int    `json:"number"`
	Name         string `json:"name"`
	EpisodeOrder int    `json:"episodeOrder"`
	PremiereDate string `json:"premiereDate"`
	EndDate      string `json:"endDate"`
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

func (s Show) GetSeasons() (seasons []Season, err error) {
	url := baseURLWithPath(fmt.Sprintf("shows/%d/seasons", s.ID))

	if _, err = s.get(url, &seasons); err != nil {
		return nil, err
	}

	return seasons, nil
}

// GetSeason returns a specific season for a show
func (s Show) GetSeason(seasonNumber int) (*Season, error) {
	seasons, err := s.GetSeasons()
	if err != nil {
		return nil, err
	}
	for _, season := range seasons {
		if season.Number == seasonNumber {
			return &season, nil
		}
	}
	return nil, fmt.Errorf("Season not found")
}
