package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// lintCmd represents the lint command
var lintCmd = &cobra.Command{
	Use:   "lint",
	Short: "Analyze a GEDCOM file",
	Long:  `Analyze a GEDCOM file for errors, duplication, and common errors.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("lint called")

		// TODO: Perform sort of linting against the GEDCOM.
	},
}

func init() {
	rootCmd.AddCommand(lintCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lintCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lintCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
