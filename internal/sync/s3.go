package sync

import (
	"fmt"
	"log"
)

func SyncToS3(source, destination string, excludePatterns []string) {
	args := []string{"s3", "sync", source, destination}

	for _, pattern := range excludePatterns {
		args = append(args, "--exclude")
		args = append(args, pattern)
	}
	fmt.Printf("aws s3 sync '%s' => '%s'\n", source, destination)

	if err := runCmd("aws", args...); err != nil {
		log.Fatalf("command finished with error: %v", err)
	}
}
