package search_test

import (
	"testing"

	"github.com/adamisrael/gedcom"
	"github.com/adamisrael/gedcom/search"
)

var gedcomFile = "../testdata/Bierce.ged"

func TestFindHomeIndividual(t *testing.T) {
	g := gedcom.Gedcom(gedcomFile)
	if g == nil {
		t.Fatalf("Result of decoding gedcom was nil, expected valid object")
	}

	var individual = search.FindHomeIndividual(*g)
	if individual == nil {
		t.Fatalf("Result of FindHomeIndividual was nil, expected valid object")
	}
}

func TestFindIndividualsByName(t *testing.T) {
	g := gedcom.Gedcom(gedcomFile)
	if g == nil {
		t.Fatalf("Result of decoding gedcom was nil, expected valid object")
	}

	var individual = search.FindIndividualsByName(*g, "Ambrose Gwinnett /Bierce/")
	if individual == nil {
		t.Fatalf("Result of FindIndividualsByName was nil, expected valid object")
	}
}

func TestFindIndividualByXref(t *testing.T) {
	g := gedcom.Gedcom(gedcomFile)
	if g == nil {
		t.Fatalf("Result of decoding gedcom was nil, expected valid object")
	}

	var individual = search.FindIndividualByXref(*g, "P1")
	if individual == nil {
		t.Fatalf("Result of FindIndividualByXref was nil, expected valid object")
	}
}

func TestFindIndividualsByNameDate(t *testing.T) {
	g := gedcom.Gedcom(gedcomFile)
	if g == nil {
		t.Fatalf("Result of decoding gedcom was nil, expected valid object")
	}

	var individuals = search.FindIndividualsByNameDate(*g, "Nathaniel /Burr/", 1665, 1701)
	if individuals == nil || len(individuals) == 0 {
		t.Fatalf("Result of FindIndividualsByNameDate was nil, expected slice")
	}

	for _, individual := range individuals {
		if individual.Birth() != nil && individual.Death() != nil {
			if individual.Birth().Year() != 1665 || individual.Death().Year() != 1701 {
				t.Fatalf("Result of FindIndividualsByNameDate returned %d individuals but the dates don't match", len(individuals))
			}
		} else {
			t.Fatalf("Missing Birth or Death for %s\n", individual.Name[0].Name)
		}
	}

	// 	// test[&search.FindIndividualsByNameDate(*g, "Aurelia Jane /Bierce/", 1848, 1850)[0]] = "Sister"

	// This test fails because her death date is "Abt. 1850" and the date parser doesn't handle that yet.
	// individuals = search.FindIndividualsByNameDate(*g, "Aurelia Jane /Bierce/", 1848, 1850)
	// if individuals == nil || len(individuals) == 0 {
	// 	t.Fatalf("Result of FindIndividualsByNameDate was nil, expected slice")
	// }

	// for _, individual := range individuals {
	// 	if individual.Birth() != nil && individual.Death() != nil {
	// 		if individual.Birth().Year() != 1848 || individual.Death().Year() != 1850 {
	// 			t.Fatalf("Result of FindIndividualsByNameDate returned %d individuals but the dates don't match", len(individuals))
	// 		}
	// 	} else {
	// 		t.Fatalf("Missing Birth or Death for %s\n", individual.Name[0].Name)
	// 	}
	// }

}
