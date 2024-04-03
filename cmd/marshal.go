package cmd

import (
	"fmt"
	json "sed/jsonHandler"
	"sed/model"

	"github.com/spf13/cobra"
)

func init() {
	jsonMarshal.Flags().StringVarP(&wrMode, "mode", "m", "", "Defines write mode [append|rewrite]")
	rootCmd.AddCommand(jsonMarshal)
}

var wrMode string

var jsonMarshal = &cobra.Command{
	Use:   "marshal [--mode] [path]",
	Short: "Prints input as json, and optionally inserts result into txt",
	Long:  "Encodes data from DB as json and prints the result, using tag --mode you can print json to text file",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var path string
		if len(args) >= 2 {
			path = args[1]
		} else {
			path = defJsonPath + args[0] + ".txt"
		}

		var object json.JSONable
		switch args[0] {
		case "user":
			object = &model.Users{}
		case "creditCard":
			object = &model.CreditCards{}
		}

		result, err := json.Marshal(&object, path, wrMode)
		if err != nil {
			fmt.Println("Error occurred during writing to a file", err)
			return
		}
		fmt.Println(string(result))
	},
}
