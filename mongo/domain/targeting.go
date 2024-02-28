package domain

import (
	"net"
	"time"

	"github.com/bsm/openrtb/v3"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
)

type TargetingSpec interface {
	Kind() string
}

type BaseSpec struct {
	Type string
}

func (tb BaseSpec) Kind() string {
	return tb.Type
}

// ALL Time Based Targeting
// FrequencyCap cap specified by the user
type FreqencyCap struct {
	MaxImpressions uint32
	NumTimeUnits   uint8
	TimeUnit       TimeUnit
}

// DayPart day parts specified by the user
type DayPart struct {
	StartTime time.Time // Time in UTC with Date set to Epoch
	EndTime   time.Time
	WeekDay   time.Weekday
	TimeZone  TimeZoneType
}

/* TODO: Pre Define these convinient times  job for UI?
Work hours,After work, Sleeping time, Weekends, Primetime */

// DayPartSpec the day pary targeting specification
type DayPartSpec struct {
	BaseSpec
	Excluded []DayPart
	Targeted []DayPart
}

// GEO Targeting

// GeoFence simple geo fence
type GeoFence struct {
	Point orb.Point
	// Radius Fence radius in Meters Min resolution is 1 meter
	// Max ~ 500 km -- Max in GAM
	RadiusMeters int
}

// GeoFenceSpec Targeting based on geofence (point and radius)
type GeoFenceSpec struct {
	BaseSpec
	Excluded []GeoFence
	Targeted []GeoFence
}

// GeoLocationSpec Targeting based on City, Country, Town ...
type GeoLocationSpec struct {
	BaseSpec
	Excluded []openrtb.Geo
	Targeted []openrtb.Geo
}

// GeoJsonSpec Targeting based on Geo Polygons chosen by the user
type GeoJsonSpec struct {
	BaseSpec
	Excluded []geojson.FeatureCollection
	Targeted []geojson.FeatureCollection
}

// Ip Address Targeting
type IpAddrSpec struct {
	BaseSpec `bson:",inline"`
	Excluded []net.IPNet
	Targeted []net.IPNet
}

func NewIpAddrSpec(inc []string, exc []string) *IpAddrSpec {
	ret := &IpAddrSpec{BaseSpec: BaseSpec{Type: "ipaddrspec"}}
	for _, iip := range inc {
		_, ipn, err := net.ParseCIDR(iip)
		if err != nil {
			continue
		}
		ret.Targeted = append(ret.Targeted, *ipn)
	}
	for _, eip := range exc {
		_, epn, err := net.ParseCIDR(eip)
		if err != nil {
			continue
		}
		ret.Excluded = append(ret.Excluded, *epn)
	}
	return ret
}
