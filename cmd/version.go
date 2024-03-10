package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

const (
	version string = "v1.0"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints version of SED you have",
	Long:  "Prints version of SED you have",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Current version of SED: ", version)
	},
}
