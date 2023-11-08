package config

const (
	DestinationLocal     = "local"
	DestinationS3        = "s3"
	DestinationS3Glacier = "s3glacier"
)

type BackupSection struct {
	Source       string `validate:"required,dirpath"`
	Destinations []struct {
		Path string `validate:"required,dirpath"`
		Type string `validate:"oneof=local s3"`
	}
	ExcludePatterns []string `yaml:"exclude_patterns"`
}

type ArchiveSection struct {
	Source      string `validate:"required,dirpath"`
	Destination struct {
		Path string `validate:"required,dirpath"`
		Type string `validate:"oneof=s3glacier"`
	}
}

type BackupConfig struct {
	Backups  map[string]BackupSection
	Archives map[string]ArchiveSection
	Global   struct {
		ExcludePatterns []string `yaml:"exclude_patterns"`
		StageDir        string   `yaml:"stage_dir" validate:"required,dirpath"`
	}
}
