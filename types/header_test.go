package types_test

import (
	"testing"

	"github.com/adamisrael/gedcom/types"
)

func TestType_header(t *testing.T) {
	header := types.Header{
		"",                   // ID
		types.Source{},       // Source
		types.CharacterSet{}, // CharacterSet
		"5.5",                // Version
		"",                   // ProductName
		"",                   // BusinessName
		types.Address{},      // BusinessAddress
		"",                   // Language
	}

	if !header.IsValid() {
		t.Fatalf("Header is invalid")
	}
}
