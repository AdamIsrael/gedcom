package main

import (
	"regexp"

	"github.com/adamisrael/gedcom/types"
)

func findHomeIndividual(g types.Gedcom) types.Individual {
	// This might not work in all cases
	// TODO: what if you run this against an empty gedcom? Will probably throw an exception
	return *g.Individual[0]
}

func findIndividualsByName(g types.Gedcom, name string) []types.Individual {
	var individuals []types.Individual

	re := regexp.MustCompile(name)

	for _, i := range g.Individual {
		if re.Find([]byte(i.Name[0].Name)) != nil {
			individuals = append(individuals, *i)
		}
	}
	return individuals
}

func findIndividualByXref(g types.Gedcom, Xref string) *types.Individual {
	for _, i := range g.Individual {
		if i.Xref == Xref {
			return i
		}
	}
	return nil
}
