# mage-extras: some predefined tasks for common mage workflows

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/mcandre/mage-extras) [![Test](https://github.com/mcandre/mage-extras/actions/workflows/test.yml/badge.svg)](https://github.com/mcandre/mage-extras/actions/workflows/test.yml) [![license](https://img.shields.io/badge/license-BSD-3)](LICENSE.md) [![Donate](https://img.shields.io/badge/-any?logo=gumroad&label=Donate&color=grey)](https://mcandre.gumroad.com/)

# SUMMARY

mage-extras provides a collection of prebaked tasks for common Go project software development needs.

# ABOUT

mage-extras defines some reusable task predicates for common workflows, in a platform-agnostic way:

* security audits
* checking that Go source code actually compiles
* running unit tests
* generating code coverage reports
* linting with assorted Go linting tools
* formatting Go code
* installing and uninstall Go applications
* collecting Go source file paths
* obtaining the GOPATH/bin directory
* referencing all local Go packages
* referencing all local Go commands
* cross-compiling applications with factorio, gox, goxcart, and xgo
* archiving artifacts
* manipulating the path separator as a string

Mage is highly agnostic about workflows. mage-extras is a little more opinionated, introducing some useful conventions on top, such as reliably obtaining a list of non-vendored Go files paths, while allowing developers to customize builds to suit their project needs.

# EXAMPLES

```console
% mage noVendor
/Users/andrew/go/src/github.com/mcandre/mage-extras/sources_test.go
/Users/andrew/go/src/github.com/mcandre/mage-extras/archive.go
/Users/andrew/go/src/github.com/mcandre/mage-extras/golint.go
/Users/andrew/go/src/github.com/mcandre/mage-extras/xgo.go
/Users/andrew/go/src/github.com/mcandre/mage-extras/version.go
/Users/andrew/go/src/github.com/mcandre/mage-extras/factorio.go
/Users/andrew/go/src/github.com/mcandre/mage-extras/sources.go
/Users/andrew/go/src/github.com/mcandre/mage-extras/staticcheck.go
/Users/andrew/go/src/github.com/mcandre/mage-extras/unmake.go
/Users/andrew/go/src/github.com/mcandre/mage-extras/mageextras_test.go
/Users/andrew/go/src/github.com/mcandre/mage-extras/chandler.go
/Users/andrew/go/src/github.com/mcandre/mage-extras/compile.go
/Users/andrew/go/src/github.com/mcandre/mage-extras/pathseparator.go
/Users/andrew/go/src/github.com/mcandre/mage-extras/revive.go
/Users/andrew/go/src/github.com/mcandre/mage-extras/goimports.go
/Users/andrew/go/src/github.com/mcandre/mage-extras/install.go
/Users/andrew/go/src/github.com/mcandre/mage-extras/test.go
/Users/andrew/go/src/github.com/mcandre/mage-extras/binaries_test.go
/Users/andrew/go/src/github.com/mcandre/mage-extras/dockerscout.go
/Users/andrew/go/src/github.com/mcandre/mage-extras/yamllint.go
/Users/andrew/go/src/github.com/mcandre/mage-extras/binaries.go
/Users/andrew/go/src/github.com/mcandre/mage-extras/deadcode.go
/Users/andrew/go/src/github.com/mcandre/mage-extras/errcheck.go
/Users/andrew/go/src/github.com/mcandre/mage-extras/mageextras.go
/Users/andrew/go/src/github.com/mcandre/mage-extras/nakedret.go
/Users/andrew/go/src/github.com/mcandre/mage-extras/tuggy.go
/Users/andrew/go/src/github.com/mcandre/mage-extras/coverage.go
/Users/andrew/go/src/github.com/mcandre/mage-extras/govulncheck.go
/Users/andrew/go/src/github.com/mcandre/mage-extras/vet.go
/Users/andrew/go/src/github.com/mcandre/mage-extras/packages.go
```

# RUNTIME REQUIREMENTS

* [Go](https://go.dev/) 1.25.4+
* [Mage](https://magefile.org/) (e.g., `go get -tool github.com/magefile/mage`)

## Recommended

* a UNIX environment, such as macOS, Linux, BSD, [WSL](https://learn.microsoft.com/en-us/windows/wsl/), etc.
* [GNU](https://www.gnu.org/software/tar/)/[BSD](https://man.freebsd.org/cgi/man.cgi?tar(1))/[Windows](https://ss64.com/nt/tar.html) tar with gzip support
* [tree](https://linux.die.net/man/1/tree)

tar is a portable archiver suitable for creating `*.tgz` tarball archives. Users can then download the tarball and extract the executable relevant to their platform. Tarballs are especially well suited for use in Docker containers, as the tar command is more likely to be installed than unzip.

Note that non-UNIX file systems may not preserve crucial chmod acl bits during port generation. This can corrupt downstream artifacts, such as compressed archives and installation procedures.

# CONTRIBUTING

For more details on developing mage-extras itself, see [DEVELOPMENT.md](DEVELOPMENT.md).
