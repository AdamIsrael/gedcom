package cmd

import (
	"fmt"

	"github.com/adamisrael/gedcom"
	"github.com/spf13/cobra"
)

// statsCmd represents the stats command
var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		g, err := gedcom.OpenGedcom(gedcomFile)
		if g == nil || err != nil {
			fmt.Printf("Invalid GEDCOM file: %s\n", err)
			return
		}

		// Display statistics about this GEDCOM
		fmt.Println("GEDCOM Statistics:")
		fmt.Printf("%d individuals\n", len(g.Individual))
		fmt.Printf("%d families\n", len(g.Family))
		fmt.Printf("%d sources\n", len(g.Source))
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
