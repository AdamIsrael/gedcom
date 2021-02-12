package types

type FamilyLink struct {
	Family *Family
	Type   string
	Note   []*Note
}

func (f FamilyLink) IsValid() bool {
	valid := true

	return valid
}
