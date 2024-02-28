package domain

type GamMoney struct {
	// Three letter code
	CurrencyCode string
	// Money values are always specified in terms of micros which are a
	// millionth of the fundamental currency unit.
	// For US dollars, $1 is 1,000,000 micros.
	MicroAmount uint64
}

// Money representation
type Money struct {
	CurrencyCode string
	Amount       float64
}
