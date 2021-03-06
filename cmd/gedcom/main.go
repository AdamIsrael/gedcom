package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/adamisrael/gedcom"
	"github.com/adamisrael/gedcom/types"
	// "gedcom"
)

var dateFormats = []string{
	"12 Feb 2006",
	"03 February 2013",
}

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

	// // Map families
	// for _, f := range g.Family {
	// 	fmt.Println(f.Xref)
	// }

	var homeIndividual = findHomeIndividual(*g)

	// var matches = findIndividualsByName(*g, "Wilhelmine Louise Auguste /Berlin/")
	// var matches = findIndividualsByName(*g, "Matthew Bryan /Israel/") // Brother
	// var matches = findIndividualsByName(*g, "Deborah Sue /Appleyard/") // Aunt
	// var matches = findIndividualsByName(*g, "Joseph Henry /Appleyard/") // 2nd great-grandfather
	// var matches = findIndividualsByName(*g, "Edith May /Appleyard/") // Great-grandaunt
	// var matches = findIndividualsByName(*g, "Alice Francis /Allen/") // Great-aunt
	// var matches = findIndividualsByName(*g, "Betty Ann /Appleyard/") // 1st Cousin 2x Removed
	var matches = findIndividualsByName(*g, "Diane Lee /Rusk/") // 2nd cousin 1x removed

	fmt.Printf("%+v\n", matches)

	// var gen = calculateGeneration(homeIndividual, matches[0])
	var gen = calculateRelationshipTest(homeIndividual, matches[0])
	// var gen = calculateRelationship(homeIndividual, homeIndividual)

	fmt.Printf("Relationship: %s\n", gen)
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
				if isSameDay(*event) {
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
	printAncestors(homeIndividual, 1)
}

func printAnniversary(i types.Individual) {
	for _, event := range i.Event {
		if isSameDay(*event) {
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
				fmt.Printf("%s%d └── %s (%s)\n", strings.Repeat("   ", generation), generation, p.Family.Husband.Name[0].Name, getAncestorRelationship(generation, "M"))
			}
			if p.Family.Wife != nil {
				fmt.Printf("%s%d ┌── %s (%s)\n", strings.Repeat("   ", generation), generation, p.Family.Wife.Name[0].Name, getAncestorRelationship(generation, "F"))
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
