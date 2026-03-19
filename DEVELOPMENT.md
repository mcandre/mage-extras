# DEVELOPMENT GUIDE

We follows standard, `go` based operations for compiling and unit testing Go code.

For advanced operations, such as linting, we further supplement with some software industry tools.

# DEV ENVIRONMENT

## Prerequisites

* [Go](https://go.dev/)
* [make](https://pubs.opengroup.org/onlinepubs/9799919799/utilities/make.html)
* Provision additional dev tools with `make -f install.mk`

## Recommended

* [asdf](https://asdf-vm.com/) 0.18

## Security Audit

```sh
mage audit
```

## Lint

```sh
mage lint
```

## Test

```sh
mage test
```

## Clean Workspace

```sh
mage -clean
```
