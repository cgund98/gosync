package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	c "github.com/cgund98/gosync/internal/config"
	"github.com/cgund98/gosync/internal/sync"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync all directories",
	Run:   runSyncCmd,
}

func runSyncCmd(cmd *cobra.Command, args []string) {
	configPath := viper.GetString("config.path")
	config := c.ParseYaml(configPath)

	// If args are provided, only sync those backups
	backups := map[string]c.BackupSection{}
	if len(args) > 0 {
		for _, sel := range args {
			if _, ok := config.Backups[sel]; !ok {
				log.Fatalf("no backup with name '%s' in config file.", sel)
			}

			backups[sel] = config.Backups[sel]
		}
	} else {
		backups = config.Backups
	}

	// Sync all backups
	for name, section := range backups {
		fmt.Printf("Syncing %s...\n", name)

		for _, destination := range section.Destinations {
			f, err := sync.MapSyncTypeToFunc(destination.Type)
			if err != nil {
				log.Fatal(err)
			}

			excludes := append(section.ExcludePatterns, config.Global.ExcludePatterns...)
			f(section.Source, destination.Path, excludes)
			fmt.Println()
		}
	}
}
