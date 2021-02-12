package types_test

import (
	"testing"

	"github.com/adamisrael/gedcom/types"
)

func TestType_trailer(t *testing.T) {
	trailer := types.Trailer{}

	if !trailer.IsValid() {
		t.Fatalf("Trailer is invalid")
	}
}
