package types

import "fmt"

// Individual contains the Individual record
type Individual struct {
	Xref      string        `json:"xref"`
	Sex       string        `json:"sex"`
	Name      []*Name       `json:"names"`
	Event     []*Event      `json:"events"`
	Attribute []*Event      `json:"attributes"`
	Parents   []*FamilyLink `json:"parents"`
	Family    []*FamilyLink `json:"family"`
}

// IsValid performs validation against the record to
// determine if it represents a valid Individual
func (i Individual) IsValid() bool {
	valid := true

	return valid
}

func (i Individual) String() string {
	return fmt.Sprintf("%v (%v)", i.Name[0], i.Sex)
}

// JSON returns a JSON-encoded version of the Individual record
func (i Individual) JSON() string {

	return fmt.Sprintf(
		`{name: "%s", sex: "%s"}`,
		i.Name[0].Name,
		i.Sex,
	)
}
