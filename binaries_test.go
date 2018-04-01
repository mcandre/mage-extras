package mageextras_test

import (
	"testing"

	"github.com/mcandre/mage-extras"
)

func TestGoBinaryPath(t *testing.T) {
	if err := mageextras.LoadGoBinariesPath(); err != nil {
		t.Error(err)
	}

	if mageextras.LoadedGoBinariesPath == "" {
		t.Errorf("Expected %v to be non-empty", mageextras.LoadedGoBinariesPath)
	}
}
