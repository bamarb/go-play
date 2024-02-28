package domain

// TimeUnit used to specify time units in frequency caps
type TimeUnit byte

const (
	TIMEUNIT_UNSPECIFIED TimeUnit = iota
	TIMEUNIT_MINUTE
	TIMEUNIT_HOUR
	TIMEUNIT_DAY
	TIMEUNIT_WEEK
	TIMEUNIT_MONTH
	TIMEUNIT_LIFETIME
)

type TimeZoneType byte

const (
	// TZT_SYSTEM use the Time zone of the system (adserver)
	TZT_SYSTEM TimeZoneType = iota
	// TZT_PUBLISHER use tthe timezone of the publisher
	TZT_PUBLISHER
	// TZT_BROWSER use the browser timezone
	TZT_BROWSER
)

type CreativeRotationType byte

const (
	// CRT_EVEN Creatives are displayed roughly the same number of
	// times over the duration of the line item.
	CRT_EVEN CreativeRotationType = iota
	// Creatives are served roughly proportionally to their performance.
	CRT_OPTIMIZED
	// Creatives are served roughly proportionally to their weights,
	// set on the LineItemCreativeAssociation.
	CRT_MANUAL
	// Creatives are served exactly in sequential order, aka Storyboarding.
	// Set on the LineItemCreativeAssociation.
	SEQUENTIAL
)

// Priority are ordered ascending 0 being lowest 100 being Highest
type Priority byte

const (
	// PRIORITY_HOUSE are no bid in-house ads for free, this is the default
	PRIORITY_HOUSE Priority = iota
	// PRIORITY_STANDARD items are involved in auction
	PRIORITY_STANDARD
	// PRIORITY_SPONSORSHIP has highest priority
	PRIORITY_SPONSORSHIP
)

// State the state of an entity (LineItem , Advertiser , Creative ...)
type State byte

const (
	STATE_UNKNOWN State = iota
	STATE_INACTIVE
	STATE_PAUSED
	STATE_ACTIVE
)

// Status the status of the entity , mostly for display in UI purposes
type Status byte

const (
	STATUS_DRAFT Status = iota
	STATUS_PENDING_APPROVAL
	STATUS_DISAPPROVED
	STATUS_PAUSED
	STATUS_CANCELLED
	STATUS_READY
	STATUS_DELIVERING
)

type DeliveryPacingType byte

const (
	DPT_NONE DeliveryPacingType = iota
	DPT_EVEN
	DPT_FRONT_LOADED
)

// GoalType The type of the goal for the LineItem.
// It defines the period over which the goal for LineItem should be reached.
type GoalType byte

const (
	GOALTYPE_NONE GoalType = iota
	GOALTYPE_LIFETIME
	GOALTYPE_DAILY
)
