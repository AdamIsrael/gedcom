package types_test

import (
	"testing"

	"github.com/adamisrael/gedcom/types"
)

func TestType_event(t *testing.T) {
	event := types.Event{
		"", // tag
		"", // value
		"", // type
		"", // Date
		types.Place{},
		types.Address{},
		"", // Age
		"", // Agency
		"", // Cause
		[]*types.Citation{},
		[]*types.MultiMedia{},
		[]*types.Note{},
	}

	if !event.IsValid() {
		t.Fatalf("Event is invalid")
	}
}
