package main

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	flagConfig = "config"
)

func init() {
	cobra.OnInitialize(initViper)

	// Commands
	rootCmd.AddCommand(parseCmd)
	rootCmd.AddCommand(syncCmd)
	rootCmd.AddCommand(archiveCmd)

	// Flags
	rootCmd.PersistentFlags().String(flagConfig, "", "config file (default is $HOME/.gosync.yaml)")
	viper.BindPFlag("config.path", rootCmd.PersistentFlags().Lookup(flagConfig))
}

func initViper() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	viper.SetDefault("config.path", homedir+"/.gosync.yaml")
}

func main() {
	rootCmd.Execute()
}
