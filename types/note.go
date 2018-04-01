package types

type Note struct {
	Note     string
	Citation []*Citation
}

func (n Note) IsValid() bool {
	valid := false

	return valid
}
