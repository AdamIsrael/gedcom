package types

type CharacterSet struct {
	Name    string
	Version string
}

func (cs CharacterSet) IsValid() bool {
	valid := false
	if len(cs.Name) <= 8 {
		valid = true
	}
	return valid
}
