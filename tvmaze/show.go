package tvmaze

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

// ShowResponse wraps a TV Maze search response
type ShowResponse struct {
	Score float64
	Show  Show
}

// Show wraps a TV Maze show object
type Show struct {
	ID        int
	Name      string
	Type      string
	Genres    []string
	Status    string
	Runtime   int
	Premiered Date
	Summary   string
	Network   network
	Embeds    struct {
		Episodes []Episode
	} `json:"_embedded"`
	Remotes map[string]interface{} `json:"externals"`
}

// GetTitle return the show title
func (s Show) GetTitle() string {
	return s.Name
}

// GetDescription returns a summary of the show
func (s Show) GetDescription() string {
	return s.Summary
}

// GetNetwork returns the network that currently broadcasts the show
func (s Show) GetNetwork() string {
	return s.Network.Name
}

// GetFirstAired return the time the first episode was aired
func (s Show) GetFirstAired() time.Time {
	if s.Premiered.Valid {
		return s.Premiered.Time
	}
	return time.Time{}
}

// GetTVRageID returns the show's ID on tvrage.com
func (s Show) GetTVRageID() int {
	return int(s.Remotes["tvrage"].(float64))
}

// GetTVDBID returns the show's ID on thetvdb.com
func (s Show) GetTVDBID() int {
	return int(s.Remotes["thetvdb"].(float64))
}

// GetIMDBID returns the show's ID on imdb.com
func (s Show) GetIMDBID() string {
	return s.Remotes["imdb"].(string)
}

// FindShows finds all matches for a given search string
func (c Client) FindShows(name string) (s []ShowResponse, err error) {
	url := baseURLWithPathQuery("search/shows", "q", name)

	if err := c.get(url, &s); err != nil {
		return nil, err
	}

	return s, nil
}

// GetShow finds all matches for a given search string
func (c Client) GetShow(name string) (*Show, error) {
	url := baseURLWithPathQuery("singlesearch/shows", "q", name)

	show := &Show{}
	if err := c.get(url, show); err != nil {
		return nil, err
	}

	return show, nil
}

// GetShowWithID finds a show by its TVMaze ID
func (c Client) GetShowWithID(tvMazeID string) (*Show, error) {
	url := baseURLWithPath(fmt.Sprintf("shows/%s", tvMazeID))

	show := &Show{}
	if err := c.get(url, show); err != nil {
		return nil, err
	}

	return show, nil
}

// GetShowWithTVRageID finds a show by its TVRage ID
func (c Client) GetShowWithTVRageID(tvRageID string) (*Show, error) {
	url := baseURLWithPathQuery("lookup/shows", "tvrage", tvRageID)

	show := &Show{}
	if err := c.get(url, show); err != nil {
		return nil, err
	}

	return show, nil
}

// GetShowWithTVDBID finds a show by its TVDB ID
func (c Client) GetShowWithTVDBID(TVDBID string) (*Show, error) {
	url := baseURLWithPathQuery("lookup/shows", "thetvdb", TVDBID)

	show := &Show{}
	if err := c.get(url, show); err != nil {
		return nil, err
	}

	return show, nil
}

// RefreshShow refreshes a show from the server
func (c Client) RefreshShow(show *Show) (err error) {
	url := baseURLWithPath(fmt.Sprintf("shows/%d", show.ID))
	return c.get(url, &show)
}

// Date represents a date from tvmaze, supporting nullability
type Date struct {
	time.Time
	Valid bool
}

// MarshalJSON implements json.Marshaler.
// It will encode null if this Date is null.
func (d *Date) MarshalJSON() ([]byte, error) {
	if !d.Valid {
		return []byte("null"), nil
	}
	return d.Time.MarshalJSON()
}

// UnmarshalJSON implements json.Unmarshaler.
// It supports string and null input.
func (d *Date) UnmarshalJSON(data []byte) error {
	var err error
	var v interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch x := v.(type) {
	case string:
		var parsedTime time.Time
		parsedTime, err = time.Parse(time.RFC3339[:10], x)
		*d = Date{parsedTime, true}
	case nil:
		d.Valid = false
		return nil
	default:
		err = fmt.Errorf("json: cannot unmarshal %v into Go value of type tvmaze.Date", reflect.TypeOf(v).Name())
	}
	d.Valid = err == nil
	return err
}
