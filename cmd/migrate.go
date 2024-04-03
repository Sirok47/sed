package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"sed/dbHandler"
	"sed/model"
)

func init() {
	rootCmd.AddCommand(migrateSchemas)
}

var migrateSchemas = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate all schemas to DB",
	Long:  "Creates all absent tables in database",
	Run: func(cmd *cobra.Command, args []string) {
		db := dbHandler.ConnectToDB()
		if err := db.AutoMigrate(&model.User{}, &model.CreditCard{}); err != nil {
			log.Println("Error while migrating schemas: ", err)
		}
	},
}
