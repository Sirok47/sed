package cmd

import (
	"github.com/spf13/cobra"
	json "sed/jsonHandler"
)

func init() {
	rootCmd.AddCommand(jsonUnmarshal)
}

var jsonUnmarshal = &cobra.Command{
	Use:   "unmarshal",
	Short: "Prints input as json",
	Long:  "Encodes inputs as json and prints the result",
	Run: func(cmd *cobra.Command, args []string) {
		json.Unmarshall()
	},
}
