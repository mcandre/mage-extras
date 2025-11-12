module github.com/mcandre/mage-extras

go 1.25.3

require (
	github.com/magefile/mage v1.15.0
	github.com/walle/targz v0.0.0-20140417120357-57fe4206da5a
)

require (
	github.com/BurntSushi/toml v1.5.0 // indirect
	github.com/alexkohler/nakedret/v2 v2.0.6 // indirect
	github.com/kisielk/errcheck v1.9.0 // indirect
	golang.org/x/exp/typeparams v0.0.0-20250408133849-7e4ce0ab07d0 // indirect
	golang.org/x/mod v0.24.0 // indirect
	golang.org/x/sync v0.13.0 // indirect
	golang.org/x/tools v0.32.0 // indirect
	honnef.co/go/tools v0.6.1 // indirect
)

tool (
	github.com/alexkohler/nakedret/v2/cmd/nakedret
	github.com/kisielk/errcheck
	github.com/magefile/mage
	honnef.co/go/tools/cmd/staticcheck
)
