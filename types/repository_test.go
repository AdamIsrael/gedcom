package types_test

import (
	"testing"

	"github.com/adamisrael/gedcom/types"
)

func TestType_repository(t *testing.T) {
	repository := types.Repository{}

	if !repository.IsValid() {
		t.Fatalf("Repository is invalid")
	}
}
