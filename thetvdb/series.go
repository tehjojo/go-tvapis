package thetvdb

import (
	"log"
	"net/url"
	"strconv"
)

// seriesResponse contains the response of a query for matching series
type seriesResponse struct {
	Series []Series `xml:"Series"`
}

// episodeResponse contains the response with information for a single episode
type episodeResponse struct {
	Episode Episode `xml:"Episode"`
}

// fullResponse contains the response for a full-series lookup
type fullResponse struct {
	Series   Series    `xml:"Series"`
	Episodes []Episode `xml:"Episode"`
}

type Series struct {
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

// GetSeries finds a series by name
func (c *Client) GetSeries(name string) []Series {
	var URL string
	var series seriesResponse

	URL = "http://thetvdb.com/api/GetSeries.php?seriesname=" + url.QueryEscape(name)
	if err := c.get(URL, &series); err != nil {
		log.Fatal(err)
	}

	return series.Series
}

func (c *Client) GetEpisodes(series Series) []Episode {
	var URL string
	var episodes fullResponse

	URL = c.BaseURL + "/api/" + c.ApiKey + "/series/" + strconv.FormatInt(int64(series.Id), 10) + "/all/en.xml"
	if err := c.get(URL, &episodes); err != nil {
		log.Fatal(err)
	}

	return episodes.Episodes
}
