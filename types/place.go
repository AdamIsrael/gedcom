package types

type Place struct {
	Name     string
	Citation []*Citation
	Note     []*Note
}

func (p Place) IsValid() bool {
	valid := false
	// if len(d.Data) <= 8 {
	// 	valid = true
	// }
	return valid
}
