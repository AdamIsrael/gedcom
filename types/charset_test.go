package types_test

import (
	"testing"

	"github.com/adamisrael/gedcom/types"
)

func TestType_Charset(t *testing.T) {
	charset := types.CharacterSet{
		"UTF-8",
		"",
	}

	if !charset.IsValid() {
		t.Fatalf("CharacterSet is invalid")
	}
}
