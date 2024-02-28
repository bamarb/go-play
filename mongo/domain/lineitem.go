package domain

import (
	"math"
	"time"
)

//go:generate stringer -type=Priority,MediaType,Status,TimeUnit

/*
https://support.google.com/admanager/answer/7085745?sjid=12135498202416396811-AP#add-labels
These are the maximum values you can set for frequency caps.
Impressions: 511 ,Minutes: 120, Hours: 48, Days: 14 Weeks: 26 Months: 6
*/

func ToDuration(f FreqencyCap) time.Duration {
	switch f.TimeUnit {
	case TIMEUNIT_MINUTE:
		return time.Duration(f.NumTimeUnits) * time.Minute
	case TIMEUNIT_HOUR:
		return time.Duration(f.NumTimeUnits) * time.Hour
	case TIMEUNIT_DAY:
		return time.Duration(f.NumTimeUnits) * time.Hour * 24
	case TIMEUNIT_WEEK:
		return time.Duration(f.NumTimeUnits) * time.Hour * 24 * 7
	case TIMEUNIT_MONTH:
		return time.Duration(f.NumTimeUnits) * time.Hour * 24 * 30
	default: // Covers UNSPECIFIED and LIFETIME
		return time.Duration(math.MaxInt64) // ~292 YEARS
	}
}

type LineItem struct {
	Base
	AdvertiserId string
	CampaignId   string
	Name         string
	FCap         FreqencyCap
	Priority     Priority
	Targeting    []TargetingSpec
}
