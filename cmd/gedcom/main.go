package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/adamisrael/gedcom"
	"github.com/adamisrael/gedcom/date"
	"github.com/adamisrael/gedcom/relationship"
	"github.com/adamisrael/gedcom/search"
	"github.com/adamisrael/gedcom/types"
	// "gedcom"
)

// var dateFormats = []string{
// 	"12 Feb 2006",
// 	"03 February 2013",
// }

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// verbose := flag.Bool("verbose", false, "verbosity")
	gedcomFile := flag.String("gedcom", "", "The path to the GEDCOM file to analyze.")
	// format := flag.String("format", "", "The format to output in (text, json).")
	anniversary := flag.Bool("anniversary", false, "Print all individuals who have an anniverary on this date.")

	flag.Parse()

	// TODO: Check if file exists
	if *gedcomFile == "" {
		fmt.Println("Invalid gedcom")
		return
	}

	g := gedcom.Gedcom(*gedcomFile)

	var homeIndividual = search.FindHomeIndividual(*g)

	// var test = make(map[*types.Individual]string)

	// // TODO: Move this into relationship_test.go
	// // TODO: Replace this with sample data
	// test[&findIndividualsByName(*g, "Matthew Bryan /Israel/")[0]] = "Brother"

	// test[&findIndividualsByName(*g, "Deborah Sue /Appleyard/")[0]] = "Aunt"
	// test[&findIndividualsByName(*g, "Alice Francis \"Lissie\" /Allen/")[0]] = "Great-Aunt"

	// test[&findIndividualsByName(*g, "Donald Martin /Israel/")[0]] = "Father"
	// test[&findIndividualsByName(*g, "Roy Edgar /Appleyard/")[0]] = "Grandfather"
	// test[&findIndividualsByName(*g, "Edgar Andas /Appleyard/")[0]] = "Great-Grandfather"
	// test[&findIndividualsByName(*g, "Joseph Henry /Appleyard/")[0]] = "2nd Great-Grandfather"
	// test[&findIndividualsByName(*g, "Alois /Auer/")[0]] = "4th Great-Grandfather"
	// test[&findIndividualsByName(*g, "Georg Salas /Auer/")[0]] = "5th Great-Grandfather"

	// test[&findIndividualsByName(*g, "Lin /Meyer/")[0]] = "1st Cousin 1x Removed"
	// test[&findIndividualsByName(*g, "Betty Ann /Appleyard/")[0]] = "1st Cousin 2x Removed"

	// test[&findIndividualsByName(*g, "Edith May /Appleyard/")[0]] = "Great-Grandaunt"
	// test[&findIndividualsByName(*g, "Louisa /Appleyard/")[0]] = "3rd Great-Grandaunt"
	// test[&findIndividualsByName(*g, "Audrey Ordery /Appleyard/")[0]] = "4th Great-Grandaunt"
	// test[&findIndividualsByName(*g, "Frances \"Fanny\" /Appleyard/")[0]] = "6th Great-Grandaunt"

	// test[&findIndividualsByName(*g, "William Kenneth /Rusk/")[0]] = "1st Cousin 2x Removed"
	// test[&findIndividualsByName(*g, "Diane Lee /Rusk/")[0]] = "2nd Cousin 1x Removed"
	// test[&findIndividualsByName(*g, "Arthur Edward /Appleyard/")[0]] = "2nd Cousin 4x Removed"

	// test[&findIndividualsByName(*g, "Melissa /Moore/")[0]] = "3rd Cousin"
	// test[&findIndividualsByName(*g, "Maisie /Appleyard/")[0]] = "4th Cousin 2x Removed"
	// test[&findIndividualsByName(*g, "Dorothy /Jones/")[0]] = "5th Cousin"

	// test[&findIndividualsByName(*g, "/Brocklesby-Atkinson/")[0]] = "5th Cousin 1x Removed"

	// for match := range test {
	// 	var relation = calculateRelationship(homeIndividual, *match)

	// 	fmt.Printf("\n** Relationship to %s: %s **\n\n", match.Name[0].Name, relation)
	// 	if relation != test[match] {
	// 		fmt.Printf("Relation to %s doesn't match: %s != %s\n", match.Name[0].Name, relation, test[match])
	// 		return
	// 	}
	// }

	// var matches = findIndividualsByName(*g, "Solomon /Israel/") // No relation

	fmt.Printf("Home Individual: %q\n", homeIndividual.Name[0].Name)

	for _, i := range g.Individual {

		// if i.Xref == "P1" {
		// 	if *format == "json" {
		// 		fmt.Println("JSON selected")
		// 		fmt.Println(i.JSON())
		// 	} else {
		// 		for _, n := range i.Name {
		// 			fmt.Printf("Name: %q\n", n.Name)
		// 		}

		// 		// for _, n := range i.Event {
		// 		// 	fmt.Printf("Event: %q\n", n.Date)
		// 		// }
		// 	}

		// 	// fmt.Printf("%+v\n", i)
		// 	// fmt.Printf("%#v\n", i)
		// }

		if *anniversary {
			// currentTime := time.Now()

			for _, event := range i.Event {
				if date.IsSameDay(*event) {
					switch event.Tag {
					case "BIRT":
						fmt.Printf("%s was born on %s\n", i.Name[0].Name, event.Date)
						// fmt.Printf("%s was born on this date %d years ago\n", i.Name[0].Name, year)
					case "DEAT":
						fmt.Printf("%s died on %s\n", i.Name[0].Name, event.Date)
					}
				}
			}
		}
	}

	fmt.Printf("%d - %s\n", 1, homeIndividual.Name[0].Name)
	printAncestors(*homeIndividual, 1)
}

func printAnniversary(i types.Individual) {
	for _, event := range i.Event {
		if date.IsSameDay(*event) {
			switch event.Tag {
			case "BIRT":
				fmt.Printf("%s was born on %s\n", i.Name[0].Name, event.Date)
				// fmt.Printf("%s was born on this date %d years ago\n", i.Name[0].Name, year)
			case "DEAT":
				fmt.Printf("%s died on %s\n", i.Name[0].Name, event.Date)
			}
		}
	}
}

func printAncestors(i types.Individual, generation int) {

	for _, p := range i.Parents {
		if p != nil {
			// Parents
			if p.Family.Husband != nil {
				fmt.Printf("%s%d └── %s (%s)\n", strings.Repeat("   ", generation), generation, p.Family.Husband.Name[0].Name, relationship.GetAncestorRelationship(generation, "M"))
			}
			if p.Family.Wife != nil {
				fmt.Printf("%s%d ┌── %s (%s)\n", strings.Repeat("   ", generation), generation, p.Family.Wife.Name[0].Name, relationship.GetAncestorRelationship(generation, "F"))
				// printAnniversary(*p.Family.Wife)
				// printAncestors(*p.Family.Wife, generation+1)
			}

			// // Children
			// for _, child := range p.Family.Child {
			// 	if child != nil {
			// 		fmt.Printf("%s      └── %s\n", strings.Repeat("   ", generation), child.Name[0].Name)
			// 		// printAnniversary(*child)
			// 	}
			// }

			if p.Family.Husband != nil {
				printAnniversary(*p.Family.Husband)
			}
			if p.Family.Husband != nil {
				printAncestors(*p.Family.Husband, generation+1)
			}
			if p.Family.Wife != nil {
				printAnniversary(*p.Family.Wife)
			}
			if p.Family.Wife != nil {
				printAncestors(*p.Family.Wife, generation+1)
			}

		}
	}
}

// // calculateGeneration should take the home person
// // and ancestor and calculate the relation, i.e.,
// // 5th great-grandfather
// func calculateGeneration(home types.Individual, ancestor types.Individual) string {
// 	fmt.Printf("Calculating generations between self and %s\n", ancestor.Name[0].Name)
// 	var generation = 0
// 	var removed = 0
// 	var err error

// 	// Figure out how many generations there are between the two Individuals
// 	var scanGeneration func(types.Individual, int, int) (int, error)

// 	scanGeneration = func(person types.Individual, generation int, removed int) (int, error) {
// 		if ancestor.Xref == person.Xref {
// 			return generation, nil
// 		}

// 		generation++

// 		for _, p := range person.Parents {

// 			if p.Family.Husband != nil {

// 				generation, err = scanGeneration(*p.Family.Husband, generation, removed)
// 				if err == nil {
// 					return generation, err
// 				}
// 				generation--
// 			}
// 			if p.Family.Wife != nil {

// 				generation, err = scanGeneration(*p.Family.Wife, generation, removed)
// 				if err == nil {
// 					return generation, err
// 				}
// 				generation--
// 			}
// 		}

// 		return generation, errors.New("individual not found")
// 	}

// 	generation, err = scanGeneration(home, generation, removed)
// 	if err != nil {
// 		return ""
// 	} else {
// 		return getGeneration(generation, ancestor.Sex)
// 	}
// }
