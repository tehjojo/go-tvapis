package tvmaze

import (
	"encoding/json"
	"testing"

	log "github.com/Sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
)

func TestTVMaze(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	Convey("I have setup a TVMaze client", t, func() {
		c := DefaultClient

		Convey("I can find a show", func() {
			results, err := c.FindShows("archer")
			So(err, ShouldBeNil)
			So(len(results), ShouldBeGreaterThan, 0)
			//fmt.Println(JSONString(results))
			So(results[0].Show.GetTitle(), ShouldNotResemble, "")
			So(results[0].Show.GetDescription(), ShouldNotResemble, "")
		})

		Convey("I can get a show by its tvmaze id", func() {
			result, err := c.GetShowWithID("315") // Archer
			So(err, ShouldBeNil)
			So(result, ShouldNotBeNil)
			//fmt.Println(JSONString(result))
			So(result.GetTitle(), ShouldNotResemble, "")
			So(result.GetDescription(), ShouldNotResemble, "")
		})

		Convey("I can get a specific episode of a show", func() {
			result, err := c.GetEpisode(Show{ID: 315}, 4, 5)
			So(err, ShouldBeNil)
			So(result, ShouldNotBeNil)
			So(result.Name, ShouldNotResemble, "")
			So(result.Summary, ShouldNotResemble, "")
		})

		Convey("I can get a show by its tvrage id", func() {
			result, err := c.GetShowWithTVRageID("23354") // Archer
			So(err, ShouldBeNil)
			So(result, ShouldNotBeNil)
			//fmt.Println(JSONString(result))
			So(result.GetTitle(), ShouldNotResemble, "")
			So(result.GetDescription(), ShouldNotResemble, "")
		})

		Convey("I can get a show", func() {
			result, err := c.GetShow("archer")
			So(err, ShouldBeNil)
			So(result, ShouldNotBeNil)
			//fmt.Println(JSONString(result))
			So(result.GetTitle(), ShouldNotResemble, "")
			So(result.GetDescription(), ShouldNotResemble, "")
		})

		Convey("I can refresh a show", func() {
			show := Show{ID: 315} // Archer
			err := c.RefreshShow(&show)
			So(err, ShouldBeNil)
			So(show.GetTitle(), ShouldNotResemble, "")
			So(show.GetDescription(), ShouldNotResemble, "")
		})

		Convey("I can get episodes for a show", func() {
			show := Show{ID: 315} // Archer
			episodes, err := c.GetEpisodes(show)
			So(err, ShouldBeNil)
			//fmt.Println(JSONString(episodes))
			So(len(episodes), ShouldBeGreaterThan, 0)
		})

		Convey("I can get the next episode for a show", func() {
			show := Show{ID: 75} // Mindy Project
			episode, err := c.GetNextEpisode(show)
			So(err, ShouldBeNil)
			So(episode, ShouldNotBeNil)
			//fmt.Println(JSONString(episode))
		})

		Convey("null times are parsed correctly", func() {
			show := Show{ID: 180} // Firefly
			episodes, err := c.GetEpisodes(show)
			So(err, ShouldBeNil)
			//fmt.Println(JSONString(episodes))
			So(len(episodes), ShouldBeGreaterThan, 0)
		})
	})
}

func JSONString(val interface{}) string {
	bytes, err := json.MarshalIndent(val, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
