package types

type Header struct {
	ID              string
	Source          Source
	CharacterSet    CharacterSet
	Version         string
	ProductName     string
	BusinessName    string
	BusinessAddress Address
	Language        string
}

func (h Header) IsValid() bool {
	valid := true
	if len(h.ID) > 20 {
		valid = false
	}
	return valid
}
