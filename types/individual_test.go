package types_test

import (
	"testing"

	"github.com/adamisrael/gedcom/types"
)

func TestType_individual(t *testing.T) {
	individual := types.Individual{
		"",  // Xref
		"M", // Sex
		[]*types.Name{
			{
				"Adam /Doe/",
				"Adam",
				"Doe",
				"Sr",
				[]*types.Citation{},
				[]*types.Note{},
			},
		}, // Name
		[]*types.Event{},      // Events
		[]*types.Event{},      // Attributes
		[]*types.FamilyLink{}, // Parents
		[]*types.FamilyLink{}, // Family
	}

	if !individual.IsValid() {
		t.Fatalf("Individual is invalid")
	}

	_ = individual.String()

	_ = individual.JSON()
}
