package types_test

import (
	"testing"

	"github.com/adamisrael/gedcom/types"
)

func TestType_name(t *testing.T) {
	name := types.Name{
		"Adam /Doe/",        // Name
		"Adam",              // Given
		"Doe",               // Surname
		"Sr.",               // Suffix
		[]*types.Citation{}, // Citations
		[]*types.Note{},     // Notes
	}

	if !name.IsValid() {
		t.Fatalf("Name is invalid")
	}
}
