package mx_test

import (
	"os"
	"reflect"
	"testing"

	mx "github.com/mcandre/mx"
)

func TestSources(t *testing.T) {
	if err := mx.CollectGoFiles(); err != nil {
		t.Error(err)
	}

	if len(mx.CollectedGoFiles) < 1 {
		t.Errorf("Expected %v to be non-empty", mx.CollectedGoFiles)
	}

	for pth := range mx.CollectedGoFiles {
		_, err := os.Stat(pth)

		if os.IsNotExist(err) {
			t.Errorf("Invalid path: %v", pth)
		}
	}

	if len(mx.CollectedGoSourceFiles) < 1 {
		t.Errorf("Expected %v to be non-empty", mx.CollectedGoSourceFiles)
	}

	for pth := range mx.CollectedGoTestFiles {
		_, err := os.Stat(pth)

		if os.IsNotExist(err) {
			t.Errorf("Invalid path: %v", pth)
		}
	}

	if len(mx.CollectedGoSourceFiles) < 1 {
		t.Errorf("Expected %v to be non-empty", mx.CollectedGoTestFiles)
	}

	for pth := range mx.CollectedGoTestFiles {
		_, err := os.Stat(pth)

		if os.IsNotExist(err) {
			t.Errorf("Invalid path: %v", pth)
		}
	}

	var sourceWithTestFiles = make(map[string]bool)

	for k, v := range mx.CollectedGoSourceFiles {
		sourceWithTestFiles[k] = v
	}

	for k, v := range mx.CollectedGoTestFiles {
		sourceWithTestFiles[k] = v
	}

	if !reflect.DeepEqual(mx.CollectedGoFiles, sourceWithTestFiles) {
		t.Errorf("Expected collected go files to include both source and test files, got: %v", mx.CollectedGoFiles)
	}
}
