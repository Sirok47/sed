package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	json "sed/jsonHandler"
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
		if len(args) == 0 {
			path = defJsonPath
		} else {
			path = args[0]
		}
		result, err := json.Marshal(path, wrMode)
		if err != nil {
			fmt.Println("Error occurred during writing to a file", err)
			return
		}
		fmt.Println(string(result))
	},
}
