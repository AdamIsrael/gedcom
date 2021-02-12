package types_test

import (
	"testing"

	"github.com/adamisrael/gedcom/types"
)

func TestType_note(t *testing.T) {
	note := types.Note{
		"",                  // Note
		[]*types.Citation{}, // Citations
	}

	if !note.IsValid() {
		t.Fatalf("Note is invalid")
	}
}
