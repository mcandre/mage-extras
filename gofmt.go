package mageextras

import (
	"os"
	"os/exec"

	"github.com/magefile/mage/mg"
)

// GoFmt runs gofmt.
func GoFmt(args ...string) error {
	mg.Deps(CollectGoFiles)

	cmdName := "gofmt"

	for pth := range CollectedGoFiles {
		cmdParameters := []string{cmdName}
		cmdParameters = append(cmdParameters, args...)
		cmdParameters = append(cmdParameters, pth)

		cmd := exec.Command(cmdName)
		cmd.Args = cmdParameters
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return err
		}
	}

	return nil
}
