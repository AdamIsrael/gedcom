package types

import "strconv"

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

	version, _ := strconv.ParseFloat(h.Version, 32)

	if version < 5.5 || version >= 5.6 {
		valid = false
	}

	return valid
}
