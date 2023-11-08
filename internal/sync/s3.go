package sync

import (
	"fmt"
	"log"

	"github.com/cgund98/gosync/internal/shell"
)

func SyncToS3(source, destination string, excludePatterns []string) {
	args := []string{"s3", "sync", source, destination, "--storage-class", "INTELLIGENT_TIERING"}

	for _, pattern := range excludePatterns {
		args = append(args, "--exclude")
		args = append(args, pattern)
	}
	fmt.Printf("aws s3 sync '%s' => '%s'\n", source, destination)

	if err := shell.RunCmd("aws", args...); err != nil {
		log.Fatalf("command finished with error: %v", err)
	}
}
