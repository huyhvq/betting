package cmd

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/huyhvq/betting/pkg/database"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate",
	Long:  `migrate.`,
	Run:   migrateExecute,
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

func migrateExecute(cmd *cobra.Command, args []string) {
	db := database.NewDB()
	conn, err := db.Open()
	if err != nil {
		panic(err)
	}
	driver, _ := mysql.WithInstance(conn, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"mysql",
		driver,
	)
	if err != nil {
		panic(err)
	}

	if len(args) > 0 && args[0] == "down" {
		fmt.Println("down")
		if err := m.Down(); err != nil {
			fmt.Println(err)
		}
	}

	if err := m.Up(); err != nil {
		fmt.Println(err)
	}
}