package tvmaze

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

// DefaultClient is the default TV Maze client
var DefaultClient = NewClient()
var baseURL = url.URL{
	Scheme: "http",
	Host:   "api.tvmaze.com",
}

// Client represents a TV Maze client
type Client struct{}

// NewClient returns a new TV Maze client
func NewClient() Client {
	return Client{}
}

func (c Client) get(url url.URL, ret interface{}) (status int, err error) {
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

func baseURLWithPath(path string) url.URL {
	ret := baseURL
	ret.Path = path
	return ret
}

func baseURLWithPathQuery(path, key, val string) url.URL {
	ret := baseURL
	ret.Path = path
	ret.RawQuery = fmt.Sprintf("%s=%s", key, url.QueryEscape(val))
	return ret
}

func baseURLWithPathQueries(path string, vals map[string]string) url.URL {
	ret := baseURL
	ret.Path = path
	var queryStrings []string
	for key, val := range vals {
		queryStrings = append(queryStrings, fmt.Sprintf("%s=%s", key, url.QueryEscape(val)))
	}
	ret.RawQuery = strings.Join(queryStrings, "&")
	return ret
}
