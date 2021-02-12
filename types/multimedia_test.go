package types_test

import (
	"testing"

	"github.com/adamisrael/gedcom/types"
)

func TestType_multimedia(t *testing.T) {
	multimedia := types.MultiMedia{}

	if !multimedia.IsValid() {
		t.Fatalf("MultiMedia is invalid")
	}
}
