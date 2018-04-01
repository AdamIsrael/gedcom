package types

type Address struct {
	Full       string
	Line1      string
	Line2      string
	City       string
	State      string
	PostalCode string
	Country    string
	Phone      string
	// Email      string
	// Fax        string
	// WWW        string
}

func (a Address) IsValid() bool {
	valid := false

	return valid
}
