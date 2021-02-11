package types_test

import (
	"testing"

	"github.com/adamisrael/gedcom/types"
)

func TestType_copyright(t *testing.T) {
	copyright := types.Copyright{
		"",
	}

	if !copyright.IsValid() {
		t.Fatalf("Copyright is invalid")
	}
}
