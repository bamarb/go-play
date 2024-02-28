package domain

type Advertiser struct {
	Base
	Name   string
	Url    string
	IabCat []string
	TaxId  string
}
