package types_test

import (
	"testing"

	"github.com/adamisrael/gedcom/types"
)

func TestType_individual(t *testing.T) {
	individual := types.Individual{
		"",                    // Xref
		"",                    // Sex
		[]*types.Name{},       // Name
		[]*types.Event{},      // Events
		[]*types.Event{},      // Attributes
		[]*types.FamilyLink{}, // Parents
		[]*types.FamilyLink{}, // Family
	}

	if !individual.IsValid() {
		t.Fatalf("Individual is invalid")
	}
}
