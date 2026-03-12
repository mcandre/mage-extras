package mx_test

import (
	"testing"

	"github.com/mcandre/mx"
)

func TestVersion(t *testing.T) {
	if mx.Version == "" {
		t.Errorf("Expected %v to be non-empty", mx.Version)
	}
}
