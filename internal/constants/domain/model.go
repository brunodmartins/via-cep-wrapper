package domain

type Address struct {
	ZipCode  string `json:"zipCode"`
	State    string `json:"state"`
	City     string `json:"city"`
	Address  string `json:"address"`
	District string `json:"district"`
}
