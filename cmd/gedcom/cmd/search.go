package cmd

import (
	"fmt"
	"strings"

	"github.com/adamisrael/gedcom"
	"github.com/adamisrael/gedcom/types"
	"github.com/spf13/cobra"
)

var (
	location      bool
	caseSensitive bool = false
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search the contents of a GEDCOM file",
	Long: `A simple, lightweight interface to searching the contents of a GEDCOM file.
For more advanced searching, use gquery.`,
	Run: func(cmd *cobra.Command, args []string) {
		g, err := gedcom.OpenGedcom(gedcomFile)
		if g == nil || err != nil {
			fmt.Printf("Invalid GEDCOM file: %s\n", err)
			return
		}

		if location {
			fmt.Printf("Search location\n")
		}
		// Do a simple search of the GEDCOM, searching Name and optionally location(s)

		if len(args) > 0 {
			var found = make(map[*types.Individual]int)

			if location {
				found = findByLocation(args, g.Individual, caseSensitive)
			} else {
				found = findByName(args, g.Individual, caseSensitive)
			}

			for i := range found {
				if found[i] >= len(args) {
					fmt.Printf("[%d] %s\n", found[i], i.Name[0].Name)
				}
			}
		}
	},
}

// findByName resturns individuals who match the search criteria
func findByName(needles []string, individuals []*types.Individual, caseSensitive bool) map[*types.Individual]int {
	var found = make(map[*types.Individual]int)

	for _, i := range individuals {
		var name = i.Name[0].Name

		if caseSensitive == false {
			name = strings.ToLower(name)
		}

		for _, needle := range needles {
			if caseSensitive == false {
				needle = strings.ToLower(needle)
			}
			if strings.Contains(name, needle) {
				found[i] += 1
			}
		}
	}

	return found
}

// findByLocation finds individuals who match the location
func findByLocation(needles []string, individuals []*types.Individual, caseSensitive bool) map[*types.Individual]int {
	var found = make(map[*types.Individual]int)

	// There's lots of normalization issues with the data, so we're just doing a token match
	for _, i := range individuals {
		for _, e := range i.Event {
			for _, needle := range needles {
				if caseSensitive == false {
					needle = strings.ToLower(needle)
				}
				// We only want to mark found once per needle
				if strings.Contains(matchCase(e.Address.Full, caseSensitive), needle) {
					fmt.Printf(".")
					found[i] += 1
				} else if strings.Contains(matchCase(e.Place.Name, caseSensitive), needle) {
					fmt.Printf(".")
					found[i] += 1
				} else if strings.Contains(matchCase(e.Address.City, caseSensitive), needle) {
					found[i] += 1
				}
			}
		}
	}
	return found
}

func matchCase(s string, caseSensitive bool) string {
	if caseSensitive {
		return s
	} else {
		return strings.ToLower(s)
	}
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	searchCmd.Flags().BoolVar(&location, "location", false, "Include locations in search")
	searchCmd.Flags().BoolVar(&caseSensitive, "case-sensitive", false, "Perform a case-sensitive search")
}
