package types_test

import (
	"testing"

	"github.com/adamisrael/gedcom/types"
)

func TestType_Address(t *testing.T) {
	address := types.Address{
		"1234 Main Street, P.O. Box 1975, Anytown, IL 60506 USA", // Full address
		"1234 Main Street", // Line 1
		"P.O. Box 1975",    // Line 2
		"",                 // Line 3
		"Anytown",          // City
		"IL",               // State
		"60506",            // Postal Code
		"USA",              // Country
		[]string{
			"555-555-5555", // Phone
		},
		[]string{
			"john@doe.com", // Email
		},
		[]string{
			"555-555-1234", // Fax
		},
		[]string{
			"http://john.doe.com", // WWW
		},
	}

	if !address.IsValid() {
		t.Fatalf("Address returned nil")
	}
}
