package sync

import (
	"fmt"
	"log"
)

func SyncToLocal(source, destination string, excludePatterns []string) {
	args := []string{"-a"}

	for _, pattern := range excludePatterns {
		args = append(args, "--exclude")
		args = append(args, pattern)
	}

	args = append(args, source)
	args = append(args, destination)

	fmt.Printf("rsync '%s' => '%s'\n", source, destination)

	if err := runCmd("rsync", args...); err != nil {
		log.Fatalf("command finished with error: %v", err)
	}
}
