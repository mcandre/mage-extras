package mageextras

import (
	"fmt"
	"os/exec"
)

// Factorio cross-compiles Go binaries for a multitude of platforms.
func Factorio(banner string, args ...string) error {
	cmd := exec.Command("factorio")
	cmd.Args = append(cmd.Args, args...)
	cmd.Env = append(cmd.Env, fmt.Sprintf("FACTORIO_BANNER=%s", banner))
	return cmd.Run()
}
