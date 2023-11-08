package sync

import (
	"fmt"
	"os/exec"

	c "github.com/cgund98/gosync/internal/config"
)

/*
 * runCmd is a wrapper for exec.Command() that reads the output of a shell command as it executes.
 */
func runCmd(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		return err
	}
	if err = cmd.Start(); err != nil {
		return err
	}
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		fmt.Print(string(tmp))
		if err != nil {
			break
		}
	}

	return cmd.Wait()
}

type SyncFunc func(source, destination string, excludePatterns []string)

func MapSyncTypeToFunc(t string) (SyncFunc, error) {

	if t == c.DestinationLocal {
		return SyncToLocal, nil
	}

	if t == c.DestinationS3 {
		return SyncToS3, nil
	}

	return nil, fmt.Errorf("no sync function found for type: '%s'", t)
}
