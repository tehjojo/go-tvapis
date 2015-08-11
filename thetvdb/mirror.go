package thetvdb

import (
	"log"
)

type mirrors struct {
	Mirrors []mirror `xml:"Mirror"`
}

type mirror struct {
	Id     int    `xml:"id"`
	Prefix string `xml:"mirrorpath"`
	Type   int    `xml:"typemask"`
}

func (c *Client) getMirrors() []mirror {
	var URL string
	var mirrors mirrors

	URL = "http://thetvdb.com/api/" + c.ApiKey + "/mirrors.xml"
	if err := c.get(URL, &mirrors); err != nil {
		log.Fatal(err)
	}

	return mirrors.Mirrors
}
