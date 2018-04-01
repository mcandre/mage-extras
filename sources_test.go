package mageextras_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/mcandre/mage-extras"
)

func TestSources(t *testing.T) {
	if err := mageextras.CollectGoFiles(); err != nil {
		t.Error(err)
	}

	if len(mageextras.CollectedGoFiles) < 1 {
		t.Errorf("Expected %v to be non-empty", mageextras.CollectedGoFiles)
	}

	for pth, _ := range mageextras.CollectedGoFiles {
		_, err := os.Stat(pth)

		if os.IsNotExist(err) {
			t.Errorf("Invalid path: %v", pth)
		}
	}

	if len(mageextras.CollectedGoSourceFiles) < 1 {
		t.Errorf("Expected %v to be non-empty", mageextras.CollectedGoSourceFiles)
	}

	for pth, _ := range mageextras.CollectedGoTestFiles {
		_, err := os.Stat(pth)

		if os.IsNotExist(err) {
			t.Errorf("Invalid path: %v", pth)
		}
	}

	if len(mageextras.CollectedGoSourceFiles) < 1 {
		t.Errorf("Expected %v to be non-empty", mageextras.CollectedGoTestFiles)
	}

	for pth, _ := range mageextras.CollectedGoTestFiles {
		_, err := os.Stat(pth)

		if os.IsNotExist(err) {
			t.Errorf("Invalid path: %v", pth)
		}
	}

	var sourceWithTestFiles = make(map[string]bool)

	for k, v := range mageextras.CollectedGoSourceFiles {
		sourceWithTestFiles[k] = v
	}

	for k, v := range mageextras.CollectedGoTestFiles {
		sourceWithTestFiles[k] = v
	}

	if !reflect.DeepEqual(mageextras.CollectedGoFiles, sourceWithTestFiles) {
		t.Errorf("Expected collected go files to include both source and test files, got: %v", mageextras.CollectedGoFiles)
	}
}
