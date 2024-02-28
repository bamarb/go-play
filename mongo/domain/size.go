package domain

// Size represents the dimensions of an AdUnit,LineItem or Creative
type Size struct {
	Width         uint
	Height        uint
	IsAspectRatio bool
}
