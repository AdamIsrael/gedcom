package types_test

import (
	"testing"

	"github.com/adamisrael/gedcom/types"
)

func TestType_Association(t *testing.T) {
	assoc := types.Association{}

	if !assoc.IsValid() {
		t.Fatalf("Association is invalid")
	}
}
