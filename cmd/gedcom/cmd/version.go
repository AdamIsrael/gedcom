package cmd

import (
	"fmt"

	"github.com/adamisrael/gedcom"
	"github.com/spf13/cobra"
)

// lintCmd represents the lint command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display version",
	Long:  `Display the version of gedcom`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Gedcom CLI Version %s\n", version)

		g, err := gedcom.OpenGedcom(gedcomFile)
		if g != nil || err == nil {
			fmt.Printf("GEDCOM version: %s\n", g.Header.Version)
			fmt.Printf("GEDCOM CharSet: %s\n", g.Header.CharacterSet.Name)
			// fmt.Printf("GEDCOM ID: %s\n", g.Header.ID)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lintCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lintCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
