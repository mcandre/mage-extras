package mageextras

import (
	"os"
	"os/exec"
)

// GoVet runs go tool vet against all Go packages in a project.
func GoVet(args ...string) error {
	cmdName := "go"

	cmdParameters := []string{cmdName}
	cmdParameters = append(cmdParameters, "vet")
	cmdParameters = append(cmdParameters, args...)
	cmdParameters = append(cmdParameters, AllPackagesPath)

	cmd := exec.Command(cmdName)
	cmd.Args = cmdParameters
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
