package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/go-playground/validator/v10"
)

func ParseYaml(path string) *BackupConfig {
	config := &BackupConfig{}

	yfile, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("unable to open config file: %v", err)
	}

	err = yaml.Unmarshal(yfile, config)
	if err != nil {
		log.Fatalf("error parsing config file: %v", err)
	}

	var validate = validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(config); err != nil {
		log.Fatal(err)
	}

	return config
}
