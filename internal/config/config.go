package config

const (
	DestinationLocal     = "local"
	DestinationS3        = "s3"
	DestinationS3Glacial = "s3glacial"
)

type BackupSection struct {
	Source       string `validate:"required,dirpath"`
	Destinations []struct {
		Path string `validate:"required,dirpath"`
		Type string `validate:"oneof=local s3 s3glacial"`
	}
	ExcludePatterns []string `yaml:"exclude_patterns"`
}

type BackupConfig struct {
	Backups map[string]BackupSection
	Global  struct {
		ExcludePatterns []string `yaml:"exclude_patterns"`
	}
}
