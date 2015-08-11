package thetvdb

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
)

const (
	MIRROR_XML     = 1
	MIRROR_BANNERS = 2
	MIRROR_ZIPS    = 4
)

// Client is the main entry point to the TheTVDB API
type Client struct {
	ApiKey     string
	BaseURL    string
	httpClient *http.Client
}

// NewClient returns a new instance of a TheTVDB client
func NewClient(apiKey string) (c *Client, err error) {
	c = new(Client)

	c.ApiKey = apiKey
	c.httpClient = http.DefaultClient
	mirrors := c.getMirrors()

	for _, mirror := range mirrors {
		if mirror.Type&MIRROR_XML != 0 {
			c.BaseURL = mirror.Prefix
			break
		}
	}

	if c.BaseURL == "" {
		err = errors.New("Coudld not determine base URL for The TV DB.")
		return nil, err
	}

	return c, nil
}

func (c *Client) get(url string, data interface{}) (err error) {
	var r *http.Response
	var b []byte

	r, err = c.httpClient.Get(url)
	if err != nil {
		return err
	}

	defer r.Body.Close()
	b, err = ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err = xml.Unmarshal(b, data); err != nil {
		return err
	}

	return nil
}
