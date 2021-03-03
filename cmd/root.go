package cmd

import (
	"fmt"
	"github.com/huyhvq/betting/pkg/database"
	"github.com/huyhvq/betting/pkg/repository"
	"github.com/huyhvq/betting/pkg/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "betting",
	Short: "Betting API",
	Long:  `Betting API`,
	Run:   serve,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.betting.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.Getwd()
		cobra.CheckErr(err)
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName("betting")
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	viper.SetEnvPrefix("bet")
	for _, cfgName := range viper.AllKeys() {
		viper.BindEnv(cfgName)
	}
}

func serve(cmd *cobra.Command, args []string) {
	db := database.NewDB()
	conn, err := db.Open()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	wr := repository.NewWager(conn)
	srv := server.NewServer(wr)
	srv.Start()
}
