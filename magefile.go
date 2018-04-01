// +build mage

package main

import (
	"fmt"

	"github.com/magefile/mage/mg"
	"github.com/mcandre/mage-extras"
)

// Default references the default build task.
var Default = Test

// Test executes the unit test suite.
func Test() error { return mageextras.UnitTest() }

// GoVet runs go tool vet.
func GoVet() error { return mageextras.GoVet("-shadow") }

// GoLint runs golint.
func GoLint() error { return mageextras.GoLint() }

// Gofmt runs gofmt.
func GoFmt() error { return mageextras.GoFmt("-s", "-w") }

// GoImports runs goimports.
func GoImports() error { return mageextras.GoImports("-w") }

// Errcheck runs errcheck.
func Errcheck() error { return mageextras.Errcheck("-blank") }

// Nakedret runs nakedret.
func Nakedret() error { return mageextras.Nakedret("-l", "0") }

// Lint runs the lint suite.
func Lint() error {
	mg.Deps(GoVet)
	mg.Deps(GoLint)
	mg.Deps(GoFmt)
	mg.Deps(GoImports)
	mg.Deps(Errcheck)
	mg.Deps(Nakedret)
	return nil
}

// NoVendor lists non-vendored Go source files.
func NoVendor() error {
	mg.Deps(mageextras.CollectGoFiles)

	for pth, _ := range mageextras.CollectedGoFiles {
		fmt.Println(pth)
	}

	return nil
}
