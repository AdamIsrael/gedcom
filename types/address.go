package types

type Address struct {
	Lines      [3]string
	City       string
	State      string
	PostalCode string
	Country    string
	Email      string
	Fax        string
	WWW        string
}

func (a Address) IsValid() bool {
	valid := false

	return valid
}
