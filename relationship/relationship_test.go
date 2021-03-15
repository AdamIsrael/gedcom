package relationship_test

import (
	"testing"

	"github.com/adamisrael/gedcom"
	"github.com/adamisrael/gedcom/relationship"
	"github.com/adamisrael/gedcom/search"
	"github.com/adamisrael/gedcom/types"
)

func TestGedcom_asdf(t *testing.T) {

	g := gedcom.Gedcom("../testdata/Bierce.ged")
	if g == nil {
		t.Fatalf("Unable to open Bierce.ged\n")
	}
	var homeIndividual = search.FindHomeIndividual(*g)

	var test = make(map[*types.Individual]string)

	// TODO: This will panic if no individuals are found
	// var individuals []types.Individual
	// if individuals = search.FindIndividualsByNameDate(*g, "Aurelius /Bierce/", 1830, 1862); len(individuals) > 0 {
	// 	test[&individuals[0]] = "Brother"
	// }

	test[&search.FindIndividualsByNameDate(*g, "Aurelius /Bierce/", 1830, 1862)[0]] = "Brother"
	test[&search.FindIndividualsByName(*g, "Aurelia Jane /Bierce/")[0]] = "Sister"

	// Aunt/Uncles
	test[&search.FindIndividualsByName(*g, "Mary Pierce /Sherwood/")[0]] = "Aunt"

	test[&search.FindIndividualsByName(*g, "Amelia Anna /McCall/")[0]] = "1st Cousin"
	test[&search.FindIndividualsByName(*g, "Anna McCall\\s+/Brush/")[0]] = "1st Cousin 1x Removed"
	test[&search.FindIndividualsByName(*g, "Sherwood /Ake/")[0]] = "1st Cousin 2x Removed"
	test[&search.FindIndividualsByName(*g, "Emily Alden /Ake/")[0]] = "1st Cousin 3x Removed"
	test[&search.FindIndividualsByName(*g, "Mary Helen /Ipson/")[0]] = "1st Cousin 4x Removed"

	test[&search.FindIndividualsByNameDate(*g, "Nathaniel /Burr/", 1665, 1701)[0]] = "3rd Great-Granduncle"
	test[&search.FindIndividualsByNameDate(*g, "Ephraim /Burr/", 1700, 1776)[0]] = "1st Cousin 4x Removed"
	test[&search.FindIndividualsByNameDate(*g, "Ephraim /Burr/", 1736, 1779)[0]] = "2nd Cousin 3x Removed"
	test[&search.FindIndividualsByNameDate(*g, "Catherine /Burr/", 1766, 1835)[0]] = "3rd Cousin 2x Removed"
	test[&search.FindIndividualsByNameDate(*g, "Silas Burr /Sherwood/", 1805, 1861)[0]] = "4th Cousin 1x Removed"
	test[&search.FindIndividualsByNameDate(*g, "Silas Burr /Sherwood/", 1830, 1896)[0]] = "5th Cousin"
	test[&search.FindIndividualsByName(*g, "Alexander West /Sherwood/")[0]] = "5th Cousin 1x Removed"
	test[&search.FindIndividualsByName(*g, "Maybelle Adeline /Sherwood/")[0]] = "5th Cousin 2x Removed"

	test[&search.FindIndividualsByName(*g, "Vesta Iola /Chappell/")[0]] = "Grandniece"
	test[&search.FindIndividualsByName(*g, "Joann T /Johnson/")[0]] = "Great-Grandniece"

	// Direct ancestors
	test[&search.FindIndividualsByName(*g, "Marcus Aurelius /Bierce/")[0]] = "Father"
	test[&search.FindIndividualsByNameDate(*g, "Nathaniel /Burr/", 1640, 1712)[0]] = "4th Great-Grandfather"

	test[&search.FindIndividualsByName(*g, "Laura /Sherwood/")[0]] = "Mother"

	for match, _ := range test {
		var relation = relationship.CalculateRelationship(*homeIndividual, *match)

		// 	fmt.Printf("\n** Relationship to %s: %s **\n\n", match.Name[0].Name, relation)
		if relation != test[match] {
			t.Fatalf("Relation to %s doesn't match: %s != %s\n", match.Name[0].Name, relation, test[match])
			return
		}
	}

}
