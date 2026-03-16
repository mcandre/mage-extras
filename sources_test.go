package mx_test

import (
	"os"
	"reflect"
	"testing"

	mx "github.com/mcandre/mx"
)

func TestSources(t *testing.T) {
	gopaths, err := mx.NoVendor()

	if err != nil {
		t.Error(err)
	}

	if len(gopaths.All) < 1 {
		t.Errorf("Expected %v to be non-empty", gopaths.All)
	}

	for pth := range gopaths.All {
		_, err := os.Stat(pth)

		if os.IsNotExist(err) {
			t.Errorf("Invalid path: %v", pth)
		}
	}

	if len(gopaths.Regular) < 1 {
		t.Errorf("Expected %v to be non-empty", gopaths.Regular)
	}

	for pth := range gopaths.Regular {
		_, err := os.Stat(pth)

		if os.IsNotExist(err) {
			t.Errorf("Invalid path: %v", pth)
		}
	}

	if len(gopaths.Test) < 1 {
		t.Errorf("Expected %v to be non-empty", gopaths.Test)
	}

	for pth := range gopaths.Test {
		_, err := os.Stat(pth)

		if os.IsNotExist(err) {
			t.Errorf("Invalid path: %v", pth)
		}
	}

	var sourceWithTestFiles = make(map[string]bool)

	for k, v := range gopaths.All {
		sourceWithTestFiles[k] = v
	}

	for k, v := range gopaths.Test {
		sourceWithTestFiles[k] = v
	}

	if !reflect.DeepEqual(gopaths.All, sourceWithTestFiles) {
		t.Errorf("Expected collected go files to include both source and test files, got: %v", gopaths.All)
	}
}
