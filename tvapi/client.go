package tvapi

// Client is the main interface for a TV API provider.
// It contains the methods to allow querying the API.
type Client interface {
	// FindShow returns all matches for a given title.
	FindShow(name string) ([]Show, error)
	// GetShow returns the best match for a given title.
	GetShow(name string) (Show, error)
	// RefreshShow loads fresh information for a show.
	RefreshShow(show Show) error
	// GetEpisodes loads the episodes for a show.
	GetEpisodes(show Show) (interface{}, error)
}
