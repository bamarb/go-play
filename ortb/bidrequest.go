package ortb

import "github.com/bsm/openrtb/v3"

// BidRequest is a wrapper around openrtb.BidRequest to add some usefull methods
// Also prepare for enriching the request. More fields might be added here depending
// on the type of enrichment
type BidRequest struct {
	*openrtb.BidRequest
}

// GetAppDomain returns the (publisher)domain of app or site
// "" if not found
func (br *BidRequest) GetAppDomain() string {
	if br == nil || br.BidRequest == nil {
		return ""
	}
	if br.App == nil {
		return ""
	}
	if br.App.Domain != "" {
		return br.App.Domain
	}
	if br.App.Publisher != nil && br.App.Publisher.Domain != "" {
		return br.App.Publisher.Domain
	}
	if br.App.Bundle != "" {
		// TODO: Reverse this string, also check if it is domain name
		return br.App.Bundle
	}
	return ""
}

func (br *BidRequest) GetSiteDomain() string {
	if br == nil || br.BidRequest == nil {
		return ""
	}
	if br.Site == nil {
		return ""
	}
	if br.Site.Domain != "" {
		return br.Site.Domain
	}
	if br.Site.Publisher != nil && br.Site.Publisher.Domain != "" {
		return br.Site.Publisher.Domain
	}
	return ""
}

// GetGeo used for location targetting
func (br *BidRequest) GetGeo() *openrtb.Geo {
	// Geo Info can be found in the User struct or Device struct
	// Device wins if both are specified
	if br == nil || br.BidRequest == nil {
		return nil
	}
	if br.Device != nil && br.Device.Geo != nil {
		return br.Device.Geo
	}
	if br.User != nil && br.User.Geo != nil {
		return br.User.Geo
	}

	return nil
}

// GetGeoPoint Lon/Lat 2d Point index 0==Lon 1==Lat
func (br *BidRequest) GetGeoPoint() [2]float64 {
	var ret [2]float64
	if br == nil || br.BidRequest == nil {
		return ret
	}

	if br.User != nil && br.User.Geo != nil {
		ret[0] = br.User.Geo.Longitude
		ret[1] = br.User.Geo.Latitude
		return ret
	}

	if br.Device != nil && br.Device.Geo != nil {
		ret[0] = br.User.Geo.Longitude
		ret[1] = br.User.Geo.Latitude
		return ret
	}
	return ret
}

// GetUserId the id of the user suitable to use for freq cap
func (br *BidRequest) GetUserId() string {
	if br == nil || br.BidRequest == nil {
		return ""
	}

	if br.User != nil && br.User.BuyerID != "" {
		return br.User.BuyerID
	}

	if br.User != nil && br.User.ID != "" {
		return br.User.ID
	}

	if br.Device != nil && br.Device.IFA != "" {
		return br.Device.IFA
	}
	return ""
}
