package archive

import (
	"fmt"
	"log"

	"github.com/cgund98/gosync/internal/shell"
)

func ArchiveS3Glacier(source, destination string) {
	args := []string{"s3", "cp", source, destination, "--storage-class", "GLACIER"}

	fmt.Printf("aws s3 cp '%s' => '%s'\n", source, destination)

	if err := shell.RunCmd("aws", args...); err != nil {
		log.Fatalf("command finished with error: %v", err)
	}
}
