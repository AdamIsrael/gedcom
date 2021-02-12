package types_test

import (
	"testing"

	"github.com/adamisrael/gedcom/types"
)

func TestType_gedcom(t *testing.T) {
	gedcom := types.Gedcom{
		types.Header{},        // Header
		&types.Submission{},   // Submission
		[]*types.Family{},     // Family
		[]*types.Individual{}, // Individual
		[]*types.MultiMedia{}, // MultiMedia
		[]*types.Repository{}, // Repository
		[]*types.Source{},     // Source
		[]*types.Submitter{},  // Submittor
		&types.Trailer{},      // Trailer
	}

	if !gedcom.IsValid() {
		t.Fatalf("Gedcom is invalid")
	}
}
