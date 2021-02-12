package types_test

import (
	"testing"

	"github.com/adamisrael/gedcom/types"
)

func TestType_source(t *testing.T) {
	source := types.Source{
		"",                    // Xref
		"",                    // Title
		[]*types.MultiMedia{}, // MultiMedia
		[]*types.Note{},       // Notes
	}

	if !source.IsValid() {
		t.Fatalf("Source is invalid")
	}
}

func TestType_sourcedata(t *testing.T) {
	sourcedata := types.SourceData{
		"", // Date
		"", // Copyright
	}

	if !sourcedata.IsValid() {
		t.Fatalf("SourceData is invalid")
	}
}
