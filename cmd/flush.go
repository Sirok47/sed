package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"sed/dbHandler"
)

func init() {
	rootCmd.AddCommand(deleteAll)
}

var deleteAll = &cobra.Command{
	Use:   "flush",
	Short: "Flushes the table",
	Long:  "Deletes all(!) entries of current table",
	Run: func(cmd *cobra.Command, args []string) {
		db := dbHandler.ConnectToDB()
		if err := db.Exec("DELETE FROM users"); err != nil {
			log.Println("Error while flushing: ", err)
		}
	},
}
