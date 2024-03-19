package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	json "sed/jsonHandler"
)

const defJsonPath = "S:/GoProj/json.txt"

func init() {
	jsonMarshal.Flags().StringVarP(&wrMode, "mode", "m", "append", "Defines write mode [append(default)|rewrite]")
	rootCmd.AddCommand(jsonMarshal)
}

var wrMode string

var jsonMarshal = &cobra.Command{
	Use:   "marshal [--mode] [path]",
	Short: "Prints input as json",
	Long:  "Encodes inputs as json and prints the result",
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
