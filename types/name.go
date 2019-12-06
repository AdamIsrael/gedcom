package types

type Name struct {
	// The raw, as-is string from the GEDCOM.
	Name string

	// The given and surname, and suffix
	Given   string
	Surname string
	Suffix  string

	Citation []*Citation
	Note     []*Note
}

func (n Name) IsValid() bool {
	valid := false
	// if len(n.Name) <= 8 {
	// 	valid = true
	// }
	return valid
}

// func (n Name) String() string {
// 	return fmt.Sprintf("%v (%v)", i.Name[0], i.Sex)
// }
