package shell

import (
	"fmt"
	"os/exec"
)

/*
 * RunCmd is a wrapper for exec.Command() that reads the output of a shell command as it executes.
 */
func RunCmd(name string, arg ...string) error {
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
