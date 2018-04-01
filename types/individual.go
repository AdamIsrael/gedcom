package types

type Individual struct {
	Xref      string
	Sex       string
	Name      []*Name
	Event     []*Event
	Attribute []*Event
	Parents   []*FamilyLink
	Family    []*FamilyLink
}

func (i Individual) IsValid() bool {
	valid := false

	return valid
}
