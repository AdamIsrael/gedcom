package types

type Name struct {
	Name     string
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
