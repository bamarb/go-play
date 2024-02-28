package domain

import "time"

// Campaign a collection of line items grouped around a theme
type Campaign struct {
	Base
	Id            string
	Name          string
	AdvertiserId  string
	StartDateTime time.Time
	EndDateTime   time.Time
}
