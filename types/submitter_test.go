package types_test

import (
	"testing"

	"github.com/adamisrael/gedcom/types"
)

func TestType_submitter(t *testing.T) {
	submitter := types.Submitter{}

	if !submitter.IsValid() {
		t.Fatalf("Submitter is invalid")
	}
}
