package tvapi

import "time"

// Show represents a single show
type Show interface {
	GetTitle() string
	GetDescription() string
	GetFirstAired() time.Time
	GetNetwork() string
}
