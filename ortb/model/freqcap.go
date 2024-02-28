package model

import (
	"math"
	"sync"
	"time"
)

//go:generate stringer -type=TimeUnit

type TimeUnit int

const (
	UNSPECIFIED TimeUnit = iota
	MINUTE
	HOUR
	DAY
	WEEK
	MONTH
	LIFETIME
)

/*
https://support.google.com/admanager/answer/7085745?sjid=12135498202416396811-AP#add-labels
These are the maximum values you can set for frequency caps.
Impressions: 511 ,Minutes: 120, Hours: 48, Days: 14 Weeks: 26 Months: 6
*/
// FrequencyCap cap specified by the user
//
type FreqencyCap struct {
	MaxImpressions uint32
	NumTimeUnits   uint8
	TimeUnit       TimeUnit
	Span           time.Duration // computed duration
}

func ToDuration(f FreqencyCap) time.Duration {
	switch f.TimeUnit {
	case MINUTE:
		return time.Duration(f.NumTimeUnits) * time.Minute
	case HOUR:
		return time.Duration(f.NumTimeUnits) * time.Hour
	case DAY:
		return time.Duration(f.NumTimeUnits) * time.Hour * 24
	case WEEK:
		return time.Duration(f.NumTimeUnits) * time.Hour * 24 * 7
	case MONTH:
		return time.Duration(f.NumTimeUnits) * time.Hour * 24 * 30
	default: // Covers UNSPECIFIED and LIFETIME
		return time.Duration(math.MaxInt64) // ~292 YEARS
	}
}

// Facp counter the zero value limits to count of 1
type Fcap struct {
	mu    sync.Mutex
	last  time.Time
	span  time.Duration
	limit uint32
	count uint32
}

func NewFcapOf(f FreqencyCap) *Fcap {
	dur := ToDuration(f)
	return &Fcap{span: dur, limit: f.MaxImpressions}
}

// Check checks whether frequency cap is met, true if met
func (fc *Fcap) CheckCapMet() bool {
	fc.mu.Lock()
	if fc.count > fc.limit {
		fc.mu.Unlock()
		return true
	}
	fc.mu.Unlock()
	return false
}

// Mark checks whether the Cap Is NOT Met
// returns true if the cap is not met and increments count
func (fc *Fcap) Mark(now time.Time) bool {
	fc.mu.Lock()
	defer fc.mu.Unlock()
	// A zero val or any fcap should allow atleast 1
	if fc.count == 0 {
		fc.last = now
		fc.count++
		return true
	}
	if fc.span == 0 {
		return false
	}

	diff := now.Sub(fc.last)
	// If we are within the span and the count is less than limit
	if diff < fc.span {
		if fc.count < fc.limit {
			fc.count++
			return true
		}
	}
	// If the current Check happenned after the span
	if diff >= fc.span {
		fc.count = 1
		fc.last = now
		return true
	}

	return false
}
