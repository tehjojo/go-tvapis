package tvmaze

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTVMaze(t *testing.T) {
	c := DefaultClient

	t.Run("get shows", func(t *testing.T) {
		t.Parallel()
		results, err := c.GetShows(0)
		require.NoError(t, err)
		require.NotEmpty(t, len(results))
		results, err = c.GetShows(99999) // This should trigger a 404, as there aren't these many pages in the index
		require.NoError(t, err)
		require.Nil(t, results)
	})

	t.Run("find show", func(t *testing.T) {
		t.Parallel()
		results, err := c.FindShows("archer")
		require.NoError(t, err)
		require.NotEmpty(t, results, "expected results")
		require.NotEmpty(t, results[0].Show.GetTitle())
		require.NotEmpty(t, results[0].Show.GetDescription())
	})

	t.Run("get show by id", func(t *testing.T) {
		t.Parallel()
		result, err := c.GetShowWithID("315") // Archer
		require.NoError(t, err)
		require.NotNil(t, result, "expected a result")
		require.Equal(t, "Archer", result.GetTitle())
		require.NotEmpty(t, result.GetDescription())
		require.Equal(t, 110381, result.GetTVDBID())
		require.Equal(t, 23354, result.GetTVRageID())
		require.Equal(t, "tt1486217", result.GetIMDBID())
		require.NotEmpty(t, result.GetMediumPoster())
		require.NotEmpty(t, result.GetOriginalPoster())
	})

	t.Run("get episode", func(t *testing.T) {
		t.Parallel()
		show := Show{ID: 315}
		result, err := show.GetEpisode(4, 5)
		require.NoError(t, err)
		require.NotNil(t, result, "expected a result")
		require.NotEmpty(t, result.Name)
		require.NotEmpty(t, result.Summary)
	})

	t.Run("get show with tvrage id", func(t *testing.T) {
		t.Parallel()
		result, err := c.GetShowWithTVRageID("23354") // Archer
		require.NoError(t, err)
		require.NotNil(t, result, "expected a result")
		require.NotEmpty(t, result.GetTitle())
		require.NotEmpty(t, result.GetDescription())
	})

	t.Run("get show by name", func(t *testing.T) {
		t.Parallel()
		result, err := c.GetShow("archer")
		require.NoError(t, err)
		require.NotNil(t, result, "expected a result")
		require.NotEmpty(t, result.GetTitle())
		require.NotEmpty(t, result.GetDescription())
	})

	t.Run("refresh show", func(t *testing.T) {
		t.Parallel()
		show := Show{ID: 315} // Archer
		err := c.RefreshShow(&show)

		require.NoError(t, err)
		require.NotEmpty(t, show.GetTitle())
		require.NotEmpty(t, show.GetDescription())
	})

	t.Run("get episodes", func(t *testing.T) {
		t.Parallel()
		show := Show{ID: 315} // Archer
		episodes, err := show.GetEpisodes()
		require.NoError(t, err)
		require.NotEmpty(t, episodes, "expected to get episodes")
	})

	t.Run("get next episode", func(t *testing.T) {
		t.Parallel()
		show := Show{ID: 35144} // jersey-shore-family-vacation
		episode, err := show.GetNextEpisode()
		require.NoError(t, err)
		require.NotNil(t, episode)
	})

	t.Run("null times", func(t *testing.T) {
		t.Parallel()
		show := Show{ID: 180} // Firefly
		episodes, err := show.GetEpisodes()
		require.NoError(t, err)
		require.NotEmpty(t, episodes, "expected to get episodes")
	})

	t.Run("seasons", func(t *testing.T) {
		t.Parallel()
		show := Show{ID: 35144} // jersey-shore-family-vacation
		seasons, err := show.GetSeasons()
		require.NoError(t, err)
		require.NotEmpty(t, seasons, "expected to get seasons")
	})

	t.Run("season 2 present", func(t *testing.T) {
		t.Parallel()
		show := Show{ID: 35144} // jersey-shore-family-vacation
		season, err := show.GetSeason(2)
		require.NoError(t, err)
		require.NotEmpty(t, season, "expected to get season 2")
	})

	t.Run("season 12 missing", func(t *testing.T) {
		t.Parallel()
		show := Show{ID: 35144} // jersey-shore-family-vacation
		season, err := show.GetSeason(12)
		require.Error(t, err)
		require.Empty(t, season, "expected to not get season 12")
	})
}
