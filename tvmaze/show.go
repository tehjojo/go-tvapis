package tvmaze

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"time"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// ShowResponse wraps a TV Maze search response
type ShowResponse struct {
	Score float64
	Show  Show
}

// Show wraps a TV Maze show object
type Show struct {
	ID         int
	Url        string
	Name       string
	Type       string
	Language   string `json:"language"`
	Genres     []string
	Status     string
	Runtime    int
	Premiered  Date
	Ended      Date `json:"ended"`
	Summary    string
	Network    network
	Updated    int `json:"updated"`
	WebChannel struct {
		ID           int     `json:"id"`
		Name         string  `json:"name"`
		Country      country `json:"country"` //Had to make it interface to avoid interface{}
		OfficialSite string  `json:"officialSite"`
	} `json:"webChannel"`
	Embeds struct {
		Episodes []Episode
	} `json:"_embedded"`
	Remotes map[string]*json.RawMessage `json:"externals"`
	Image   struct {
		Medium   string
		Original string
	}
	Seasons []Season
}

func (s Show) get(url url.URL, ret interface{}) (status int, err error) {
	r, err := http.Get(url.String())
	if err != nil {
		return 0, errors.Wrapf(err, "failed to get url: %s", url.String())
	}

	if r.StatusCode >= http.StatusBadRequest {
		return r.StatusCode, errors.Errorf("received error status code (%d): %s", r.StatusCode, r.Status)
	}

	defer r.Body.Close()
	return r.StatusCode, json.NewDecoder(r.Body).Decode(&ret)
}

// GetTitle return the show title
func (s Show) GetTitle() string {
	return s.Name
}

// GetDescription returns a summary of the show
func (s Show) GetDescription() string {
	return s.Summary
}

// IsWebOnly returns whether the show is online only (ie. Netflix/Amazon Prime or other networks)
func (s Show) IsWebOnly() bool {
	return s.WebChannel.ID != 0
}

// GetNetwork returns the network that currently broadcasts the show
func (s Show) GetNetwork() string {
	if s.WebChannel.Name != "" {
		return s.WebChannel.Name
	} else {
		return s.Network.Name
	}
}

// GetCountry returns the country that currently broadcasts the show
func (s Show) GetCountry() string {
	if s.WebChannel.Country.Name != "" {
		return s.WebChannel.Country.Name
	} else {
		return s.Network.Country.Name
	}
}

// GetFirstAired return the time the first episode was aired
func (s Show) GetFirstAired() time.Time {
	if s.Premiered.Valid {
		return s.Premiered.Time
	}
	return time.Time{}
}

// GetMediumPoster returns the URL to a medium sized poster
func (s Show) GetMediumPoster() string {
	return s.Image.Medium
}

// GetOriginalPoster returns the URL to an original sized poster
func (s Show) GetOriginalPoster() string {
	return s.Image.Original
}

// GetTVRageID returns the show's ID on tvrage.com
func (s Show) GetTVRageID() int {
	if s.Remotes["tvrage"] == nil {
		return 0
	}
	var val int
	if err := json.Unmarshal(*s.Remotes["tvrage"], &val); err != nil {
		log.WithError(err).WithField("tvrage_id", s.Remotes["tvrage"]).Error("failed to parse tvrage id")
	}
	return val
}

// GetTVDBID returns the show's ID on thetvdb.com
func (s Show) GetTVDBID() int {
	if s.Remotes["thetvdb"] == nil {
		return 0
	}
	var val int
	if err := json.Unmarshal(*s.Remotes["thetvdb"], &val); err != nil {
		log.WithError(err).WithField("thetvdb_id", s.Remotes["thetvdb"]).Error("failed to parse thetvdb id")
	}
	return val
}

// GetIMDBID returns the show's ID on imdb.com
func (s Show) GetIMDBID() string {
	if s.Remotes["imdb"] == nil {
		return ""
	}
	var val string
	if err := json.Unmarshal(*s.Remotes["imdb"], &val); err != nil {
		log.WithError(err).WithField("imdb_id", s.Remotes["imdb"]).Error("failed to parse imdb id")
	}
	return val
}

// GetShows returns a list of all shows in the TVMaze database. When the end of the index
// is reached, a nil slice is returned.
func (c Client) GetShows(page int) ([]Show, error) {
	url := baseURLWithPathQuery("shows", "page", strconv.Itoa(page))
	shows := []Show{}
	if status, err := c.get(url, &shows); err != nil {
		if status == http.StatusNotFound {
			return nil, nil
		}
		log.Printf("err on")
		return nil, err
	}
	return shows, nil
}

// FindShows finds all matches for a given search string
func (c Client) FindShows(name string) (s []ShowResponse, err error) {
	url := baseURLWithPathQuery("search/shows", "q", name)

	if _, err := c.get(url, &s); err != nil {
		return nil, err
	}

	return s, nil
}

// FindShowsUrl Returns the search URL for a given search string
func (c Client) FindShowsUrl(name string) url.URL {
	return baseURLWithPathQuery("search/shows", "q", name)
}

// GetShow finds all matches for a given search string
func (c Client) GetShow(name string) (*Show, error) {
	url := baseURLWithPathQuery("singlesearch/shows", "q", name)

	show := &Show{}
	if _, err := c.get(url, show); err != nil {
		return nil, err
	}
	if show.ID != 0 {
		seasons, err := show.getSeasons()
		if err != nil {
			fmt.Println("Unable to fetch seasons for show")
		} else {
			show.Seasons = seasons
		}
	}

	return show, nil
}

// GetShowWithID finds a show by its TVMaze ID
func (c Client) GetShowWithID(tvMazeID string) (*Show, error) {
	url := baseURLWithPath(fmt.Sprintf("shows/%s", tvMazeID))

	show := &Show{}
	if _, err := c.get(url, show); err != nil {
		return nil, err
	}

	if show.ID != 0 {
		seasons, err := show.getSeasons()
		if err != nil {
			fmt.Println("Unable to fetch seasons for show")
		} else {
			show.Seasons = seasons
		}
	}

	return show, nil
}

// GetShowWithTVRageID finds a show by its TVRage ID
func (c Client) GetShowWithTVRageID(tvRageID string) (*Show, error) {
	url := baseURLWithPathQuery("lookup/shows", "tvrage", tvRageID)

	show := &Show{}
	if _, err := c.get(url, show); err != nil {
		return nil, err
	}

	if show.ID != 0 {
		seasons, err := show.getSeasons()
		if err != nil {
			fmt.Println("Unable to fetch seasons for show")
		} else {
			show.Seasons = seasons
		}
	}
	return show, nil
}

// GetShowWithTVDBID finds a show by its TVDB ID
func (c Client) GetShowWithTVDBID(TVDBID string) (*Show, error) {
	url := baseURLWithPathQuery("lookup/shows", "thetvdb", TVDBID)

	show := &Show{}
	if _, err := c.get(url, show); err != nil {
		return nil, err
	}

	if show.ID != 0 {
		seasons, err := show.getSeasons()
		if err != nil {
			fmt.Println("Unable to fetch seasons for show")
		} else {
			show.Seasons = seasons
		}
	}

	return show, nil
}

// RefreshShow refreshes a show from the server
func (c Client) RefreshShow(show *Show) error {
	url := baseURLWithPath(fmt.Sprintf("shows/%d", show.ID))
	_, err := c.get(url, &show)
	//Blindly refresh
	show.Seasons, _ = show.getSeasons()
	return err
}

func (c Client) GetSeasons(showID int) (seasons []Season, err error) {
	url := baseURLWithPath(fmt.Sprintf("shows/%d/seasons", showID))

	if _, err = c.get(url, &seasons); err != nil {
		return nil, err
	}

	return seasons, nil
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
		return errors.Wrap(err, "failed to unmarshal JSON response")
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
		err = errors.Errorf("json: cannot unmarshal %v into Go value of type tvmaze.Date", reflect.TypeOf(v).Name())
	}
	d.Valid = err == nil
	return err
}
