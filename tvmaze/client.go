package tvmaze

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// DefaultClient is the default TV Maze client
var DefaultClient = NewClient()

// Client represents a TV Maze client
type Client struct {
	baseURL string
	client  *http.Client
}

// NewClient returns a new TV Maze client
func NewClient() *Client {
	return &Client{
		baseURL: "http://api.tvmaze.com",
		client:  &http.Client{},
	}
}

func (c *Client) get(path string, ret interface{}) (err error) {
	var r *http.Response
	var URL string
	var b []byte

	URL = c.baseURL + path
	r, err = http.Get(URL)
	if err != nil {
		return err
	}

	defer r.Body.Close()
	b, err = ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, &ret)
}

type date struct {
	time.Time
}

func (d *date) UnmarshalJSON(b []byte) (err error) {
	const format = "2006-01-02"
	var v string

	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}

	t, err := time.Parse(format, v)
	if err != nil {
		return err
	}

	*d = date{t}

	return nil
}
