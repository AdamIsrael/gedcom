package types_test

import (
	"testing"

	"github.com/adamisrael/gedcom/types"
)

func TestType_Citation(t *testing.T) {
	source := *&types.Source{
		"", // Xref
		"", // Title
		[]*types.MultiMedia{},
		[]*types.Note{},
	}

	citation := types.Citation{
		&source,
		"",           // Page
		types.Data{}, // Data
		"",           // Quay
		[]*types.MultiMedia{},
		[]*types.Note{},
	}

	if !citation.IsValid() {
		t.Fatalf("Citation is invalid")
	}
}
