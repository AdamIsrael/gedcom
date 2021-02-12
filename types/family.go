package types

type Family struct {
	Xref    string
	Husband *Individual
	Wife    *Individual
	Child   []*Individual
	Event   []*Event
}

func (f Family) IsValid() bool {
	valid := true

	return valid
}
