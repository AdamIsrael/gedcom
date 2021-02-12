package types_test

import (
	"testing"

	"github.com/adamisrael/gedcom/types"
)

func TestType_submission(t *testing.T) {
	submission := types.Submission{}

	if !submission.IsValid() {
		t.Fatalf("Submission is invalid")
	}
}
