package thetvdb

import (
	"encoding/xml"
	"time"
)

const DATE_FORMAT = "2006-01-02"

type Date struct {
	time.Time
}

func (date *Date) UnmarshalXML(d *xml.Decoder, el xml.StartElement) (err error) {
	var v string
	var t time.Time

	if err = d.DecodeElement(&v, &el); err != nil {
		return err
	}

	t, err = time.Parse(DATE_FORMAT, v)
	if err != nil {
		return err
	}

	*date = Date{t}

	return nil
}
