package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	json "sed/jsonHandler"
)

var stamp *bool

func init() {
	stamp = jsonUnmarshal.Flags().BoolP("db", "d", false, "use it to record the result to DB")
	rootCmd.AddCommand(jsonUnmarshal)
}

var jsonUnmarshal = &cobra.Command{
	Use:   "unmarshal [path]",
	Short: "Prints json as object, and optionally inserts result into DB",
	Long:  "Decodes data from text file and prints the result, using tag -d you can stamp object to database",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var path string
		if len(args) == 0 {
			path = defJsonPath
		} else {
			path = args[0]
		}
		result, err := json.Unmarshall(path, *stamp)
		if err != nil {
			fmt.Println("Error occurred during reading from json", err)
			return
		}
		fmt.Println(result)
	},
}
