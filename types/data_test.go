package types_test

import (
	"testing"

	"github.com/adamisrael/gedcom/types"
)

func TestType_data(t *testing.T) {
	data := types.Data{
		"",
		[]string{
			"line 1",
		},
	}

	if !data.IsValid() {
		t.Fatalf("Data is invalid")
	}
}
