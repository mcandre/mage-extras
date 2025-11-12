package mageextras

import (
	"github.com/walle/targz"

	"fmt"
)

// Archive compresses build artifacts.
func Archive(portBasename string, artifactsPath string) error {
	archiveFilename := fmt.Sprintf("%s.tgz", portBasename)
	return targz.Compress(artifactsPath, archiveFilename)
}
