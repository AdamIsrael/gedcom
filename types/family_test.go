package types_test

import (
	"testing"

	"github.com/adamisrael/gedcom/types"
)

func TestType_family(t *testing.T) {
	family := types.Family{
		"", // Xref
		&types.Individual{},
		&types.Individual{},
		[]*types.Individual{},
		[]*types.Event{},
	}

	if !family.IsValid() {
		t.Fatalf("Family is invalid")
	}
}
