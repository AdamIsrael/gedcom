package types

type Event struct {
	Tag      string
	Value    string
	Type     string
	Date     string
	Place    Place
	Address  Address
	Age      string
	Agency   string
	Cause    string
	Citation []*Citation
	Media    []*MultiMedia
	Note     []*Note
}

func (e Event) IsValid() bool {
	valid := true

	return valid
}
