package types

type Data struct {
	Date string
	Text []string
}

func (d Data) IsValid() bool {
	valid := true
	// if len(d.Data) <= 8 {
	// 	valid = true
	// }
	return valid
}
