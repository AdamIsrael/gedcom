package cmd

import (
	"fmt"
	"strings"

	"github.com/adamisrael/gedcom"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search the contents of a GEDCOM file",
	Long: `A simple, lightweight interface to searching the contents of a GEDCOM file.
For more advanced searching, use gquery.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("search called")

		g, err := gedcom.OpenGedcom(gedcomFile)
		if g == nil || err != nil {
			fmt.Printf("Invalid GEDCOM file: %s\n", err)
		}

		// Do a simple search of the GEDCOM, searching Name and optionally location(s)
		fmt.Printf("%v\n", args)
		if len(args) > 0 {
			for _, i := range g.Individual {
				if strings.Contains(i.Name[0].Name, args[0]) {
					fmt.Printf("%s\n", i.Name[0].Name)
				}
			}
		}
	},
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

	searchCmd.Flags().BoolP("location", "l", false, "Include locations in search")

}
