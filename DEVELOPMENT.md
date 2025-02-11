# BUILDTIME REQUIREMENTS

* [Go](https://go.dev/) 1.23.3+
* [Snyk](https://snyk.io/)
* Provision additional dev tools with `go install tool`

## Recommended

* [ASDF](https://asdf-vm.com/) 0.10 (run `asdf reshim` after provisioning)
* [direnv](https://direnv.net/) 2
* macOS [open](https://ss64.com/mac/open.html) or equivalent alias

Non-UNIX environments may produce subtle adverse effects when linting or generating application ports.

## Windows

Apply a user environment variable `GODEBUG=modcacheunzipinplace=1` per [access denied resolution](https://github.com/golang/go/wiki/Modules/e93463d3e853031af84204dc5d3e2a9a710a7607#go-115), for native Windows development environments (Command Prompt / PowerShell, not WLS, not Cygwin, not MSYS2, not MinGW, not msysGit, not Git Bash, not etc).

# AUDIT

```console
$ mage audit
```

# TEST

```console
$ mage test
```

# COVERAGE

```console
$ mage coverageHTML
$ open cover.html
```

# LINT

```console
$ mage lint
```

# CLEAN

```console
$ mage clean
```
