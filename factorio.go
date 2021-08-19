package mageextras

import (
	"fmt"
	"os"
	"os/exec"
)

// Factorio cross-compiles Go binaries for a multitude of platforms.
func Factorio(banner string, args ...string) error {
	cmd := exec.Command("factorio")
	cmd.Args = []string{"factorio"}
	cmd.Args = append(cmd.Args, args...)
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, fmt.Sprintf("FACTORIO_BANNER=%s", banner))
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
