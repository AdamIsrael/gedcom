package types_test

import (
	"testing"

	"github.com/adamisrael/gedcom/types"
)

func TestType_familylink(t *testing.T) {
	link := types.FamilyLink{
		&types.Family{},
		"", // Type
		[]*types.Note{},
	}

	if !link.IsValid() {
		t.Fatalf("Copyright is invalid")
	}
}
