package cmd

import (
	"fmt"
	json "sed/jsonHandler"
	"sed/model"

	"github.com/spf13/cobra"
)

var stamp *bool

func init() {
	stamp = jsonUnmarshal.Flags().BoolP("db", "d", false, "use it to record the result to DB")
	rootCmd.AddCommand(jsonUnmarshal)
}

var jsonUnmarshal = &cobra.Command{
	Use:   "unmarshal [model] [path]",
	Short: "Prints json as object, and optionally inserts result into DB",
	Long:  "Decodes data from text file and prints the result, using tag -d you can stamp object to database",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		var scheme json.JSONable
		switch args[0] {
		case "user":
			scheme = &model.Users{}
		case "creditCard":
			scheme = &model.CreditCards{}
		}

		var path string
		if len(args) >= 2 {
			path = args[1]
		} else {
			path = defJsonPath + args[0] + ".txt"
		}

		err := json.Unmarshal(&scheme, path, *stamp)
		if err != nil {
			fmt.Println("Error occurred during reading from json", err)
			return
		}
		fmt.Println(scheme.GetValue()...)
	},
}
