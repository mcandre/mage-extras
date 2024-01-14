package mageextras

import (
	"fmt"
	"os"
	"os/exec"
)

// Archive compresses build artifacts.
//
// Assumes a conventional bin build directory.
func Archive(portBasename string, artifactsPath string) error {
	archiveFilename := fmt.Sprintf("%s.tgz", portBasename)

	cmdName := "tar"

	cmdParameters := []string{cmdName}
	cmdParameters = append(cmdParameters, "czf")
	cmdParameters = append(cmdParameters, archiveFilename)
	cmdParameters = append(cmdParameters, "-C")
	cmdParameters = append(cmdParameters, "bin")
	cmdParameters = append(cmdParameters, artifactsPath)

	cmd := exec.Command(cmdName)
	cmd.Args = cmdParameters
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
