package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	c "github.com/cgund98/gosync/internal/config"
)

var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Parse the configuration file and list out backups.",
	Run:   runParseCmd,
}

func runParseCmd(cmd *cobra.Command, args []string) {
	configPath := viper.GetString("config.path")

	fmt.Printf("Parsing config file at '%v'...", configPath)
	config := c.ParseYaml(configPath)

	// Construct output message
	msg := ""
	for name, section := range config.Backups {
		msg += "\n"
		msg += fmt.Sprintf("%s:\n", name)
		msg += fmt.Sprintf("  source: %v\n", section.Source)
		msg += "  destinations:\n"
		for _, dest := range section.Destinations {
			msg += fmt.Sprintf("    - %s (%s)\n", dest.Path, dest.Type)
		}
		msg += "  exclude_patterns:\n"
		for _, pattern := range section.ExcludePatterns {
			msg += fmt.Sprintf("    - %s\n", pattern)
		}
	}
	msg += "global:\n"
	for _, pattern := range config.Global.ExcludePatterns {
		msg += fmt.Sprintf(" - %s\n", pattern)
	}

	fmt.Println(msg)
}
