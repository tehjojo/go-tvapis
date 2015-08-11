package thetvdb

import (
	"log"
	"net/url"
	"strconv"
	"time"
)

// seriesResponse contains the response of a query for matching series
type seriesResponse struct {
	Series []*Show `xml:"Series"`
}

// episodeResponse contains the response with information for a single episode
type episodeResponse struct {
	Episode *Episode `xml:"Episode"`
}

// fullResponse contains the response for a full-series lookup
type fullResponse struct {
	Series   *Show      `xml:"Series"`
	Episodes []*Episode `xml:"Episode"`
}

// Show contains all data related to the TV show
type Show struct {
	Id          int    `xml:"seriesid"`
	Locale      string `xml:"language"`
	Name        string `xml:"SeriesName"`
	Description string `xml:"Overview"`
	Network     string
	FirstAired  Date
}

type Episode struct {
	Id          int    `xml:"id"`
	Name        string `xml:"EpisodeName"`
	Number      int    `xml:"EpisodeNumber"`
	Season      int    `xml:"SeasonNumber"`
	Description string `xml:"Overview"`
	FirstAired  Date
}

// GetTitle return the show title
func (s *Show) GetTitle() string {
	return s.Name
}

// GetDescription returns a summary of the show
func (s *Show) GetDescription() string {
	return s.Description
}

// GetNetwork returns the network that currently broadcasts the show
func (s *Show) GetNetwork() string {
	return s.Network
}

// GetFirstAired return the time the first episode was aired
func (s *Show) GetFirstAired() time.Time {
	return s.FirstAired
}

func (c *Client) FindShow(name string) ([]Show, error) {
	var URL string
	var series seriesResponse

	URL = "http://thetvdb.com/api/GetSeries.php?seriesname=" + url.QueryEscape(name)
	if err := c.get(URL, &series); err != nil {
		return nil, err
	}

	return series.Series
}

func (c *Client) GetShow(name string) (*Show, error) {
	shows, err := c.FindShow(name)
	if err != nil {
		return nil, err
	}

	return shows[0]
}

func (c *Client) RefreshShow(show *Show) error {
	var URL string
	var response fullResponse

	URL = c.BaseURL + "/api/" + c.ApiKey + "/series/" + strconv.FormatInt(int64(show.Id), 10) + "/en.xml"
	if err := c.get(URL, &response); err != nil {
		return err
	}

	return nil
}

func (c *Client) GetEpisodes(series Show) ([]Episode, error) {
	var URL string
	var episodes fullResponse

	URL = c.BaseURL + "/api/" + c.ApiKey + "/series/" + strconv.FormatInt(int64(series.Id), 10) + "/all/en.xml"
	if err := c.get(URL, &episodes); err != nil {
		return nil, err
	}

	return episodes.Episodes
}
