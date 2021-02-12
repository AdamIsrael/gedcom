package types_test

import (
	"testing"

	"github.com/adamisrael/gedcom/types"
)

func TestType_place(t *testing.T) {
	place := types.Place{
		"",                  // Name
		[]*types.Citation{}, // Citations
		[]*types.Note{},     // Notes
	}

	if !place.IsValid() {
		t.Fatalf("Place is invalid")
	}
}
