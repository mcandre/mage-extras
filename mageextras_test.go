package mageextras_test

import (
	"testing"

	mageextras "github.com/mcandre/mage-extras"
)

func TestVersion(t *testing.T) {
	if mageextras.Version == "" {
		t.Errorf("Expected %v to be non-empty", mageextras.Version)
	}
}
