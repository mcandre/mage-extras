package mx

import (
	"bytes"
	"io"
	"log"
	"os"

	"github.com/magefile/mage/sh"
)

// GoReleaser executes goreleaser.
func GoReleaser(env map[string]string, args ...string) error {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	_, err := sh.Exec(env, &stdout, &stderr, "goreleaser", args...)

	if err != nil {
		if _, err2 := io.Copy(os.Stdout, &stdout); err2 != nil {
			log.Println(err2)
		}

		if _, err2 := io.Copy(os.Stderr, &stderr); err2 != nil {
			log.Println(err2)
		}
	}

	return err
}
