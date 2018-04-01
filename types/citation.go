package types

type Citation struct {
	Source *Source
	Page   string
	Data   Data
	Quay   string
	Media  []*MultiMedia
	Note   []*Note
}

func (c Citation) IsValid() bool {
	valid := false
	// if len(c.Name) <= 8 {
	// 	valid = true
	// }
	return valid
}
