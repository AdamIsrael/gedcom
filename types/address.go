package types

// The Address structure
type Address struct {
	Full       string
	Line1      string
	Line2      string
	Line3      string
	City       string
	State      string
	PostalCode string
	Country    string
	// These can have up to three records each
	Phone []string
	Email []string
	Fax   []string
	WWW   []string
}

// IsValid checks if the address structure is valid
func (a Address) IsValid() bool {
	valid := true

	// Phone, Email, Fax, and WWW can have a max of 3
	return valid
}
