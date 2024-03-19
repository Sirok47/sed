package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	json "sed/jsonHandler"
)

func init() {
	rootCmd.AddCommand(jsonUnmarshal)
}

var jsonUnmarshal = &cobra.Command{
	Use:   "unmarshal [path]",
	Short: "Prints input as json",
	Long:  "Encodes inputs as json and prints the result",
	Run: func(cmd *cobra.Command, args []string) {
		var path string
		if len(args) == 0 {
			path = defJsonPath
		} else {
			path = args[0]
		}
		result, err := json.Unmarshall(path)
		if err != nil {
			fmt.Println("Error occurred during reading from json", err)
			return
		}
		fmt.Println(result)
	},
}
