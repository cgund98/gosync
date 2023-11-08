package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cgund98/gosync/internal/archive"
	c "github.com/cgund98/gosync/internal/config"
	"github.com/cgund98/gosync/internal/shell"
)

var archiveCmd = &cobra.Command{
	Use:   "archive",
	Short: "Archive directories",
	Run:   runArchiveCmd,
}

func checkDirExists(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return fileInfo.IsDir(), err
}

func runArchiveCmd(cmd *cobra.Command, args []string) {
	configPath := viper.GetString("config.path")
	config := c.ParseYaml(configPath)

	// If args are provided, only upload those achives
	achives := map[string]c.ArchiveSection{}
	if len(args) > 0 {
		for _, sel := range args {
			if _, ok := config.Archives[sel]; !ok {
				log.Fatalf("no achive with name '%s' in config file.", sel)
			}

			achives[sel] = config.Archives[sel]
		}
	} else {
		achives = config.Archives
	}

	// Archive all achives
	for name, section := range achives {
		fmt.Printf("Archiving %s...\n", name)

		stageFile := config.Global.StageDir + fmt.Sprintf("/%s.zip", name)

		f, err := archive.MapArchiveTypeToFunc(section.Destination.Type)
		if err != nil {
			log.Fatal(err)
		}

		// Copy config file to zipfile
		dirExists, err := checkDirExists(section.Source)
		if err != nil {
			log.Fatalf("unable to validate source path: %v", err)
		} else if !dirExists {
			log.Fatalf("invalid source directory: %s", section.Source)
		}
		// Copy archive directory to stage
		stageDir := config.Global.StageDir + fmt.Sprintf("/%s", name)
		if err := os.RemoveAll(stageDir); err != nil {
			log.Fatalf("unable to remove stage directory: %v", err)
		}
		if err := shell.RunCmd("cp", "-r", section.Source, stageDir); err != nil {
			log.Fatalf("unable to copy to archive folder to '%s'. %v", stageDir, err)
		}
		if err := shell.RunCmd("cp", configPath, section.Source); err != nil {
			log.Fatalf("unable to copy config file to archive folder: %v", err)
		}

		// Create zipfile
		err = archive.ZipSource(stageDir, stageFile)
		if err != nil {
			log.Fatalf("unable to zip directory: %v", err)
		}

		// Upload archive
		f(stageFile, section.Destination.Path)
		fmt.Println()

		// Cleanup zipfile
		// if err := os.Remove(stageFile); err != nil {
		// 	log.Fatalf("unable to delete zipfile: %v", err)
		// }

		// Cleanup stage directory
		if err := os.RemoveAll(stageDir); err != nil {
			log.Fatalf("unable to remove stage directory: %v", err)
		}
	}
}
