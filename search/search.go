package search

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/adamisrael/gedcom/date"
	"github.com/adamisrael/gedcom/types"
)

func FindHomeIndividual(g types.Gedcom) *types.Individual {
	// This might not work in all cases
	// TODO: what if you run this against an empty gedcom? Will probably throw an exception
	return g.Individual[0]
}

// FindIndividualsByNameDate finds Individuals by name, matching their year
// of birth and death to limit results
func FindIndividualsByNameDate(g types.Gedcom, name string, yob int, yod int) []types.Individual {
	var individuals []types.Individual
	// var err error
	// var t time.Time

	re := regexp.MustCompile(name)

	for _, i := range g.Individual {
		var birth int
		var death int

		if re.Find([]byte(i.Name[0].Name)) != nil {
			// TODO: Implement better date handling through the Individual object
			for _, event := range i.Event {
				switch event.Tag {
				case "BIRT":
					t, err := date.Parse(event.Date)
					if err == nil {
						birth = t.Year()
					}

				case "DEAT":
					t, err := date.Parse(event.Date)
					if err == nil {
						death = t.Year()
					} else {
						fmt.Printf("Failed to parse %s\n", event.Date)
					}
				}
			}
			if birth == yob && death == yod {
				individuals = append(individuals, *i)
			}
		}
	}
	return individuals

}

func FindIndividualsByName(g types.Gedcom, name string) []types.Individual {
	var individuals []types.Individual

	// Remove extra spaces
	name = strings.ReplaceAll(name, "  ", " ")
	name = strings.TrimSpace(name)

	re := regexp.MustCompile(name)

	for _, i := range g.Individual {
		if re.Find([]byte(i.Name[0].Name)) != nil {
			individuals = append(individuals, *i)
		}
	}
	if len(individuals) == 0 {
		fmt.Printf("Couldn't find %s\n", name)
	}
	return individuals
}

func FindIndividualByXref(g types.Gedcom, Xref string) *types.Individual {
	for _, i := range g.Individual {
		if i.Xref == Xref {
			return i
		}
	}
	return nil
}
