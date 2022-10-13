# BUILDTIME REQUIREMENTS

* [Go](https://golang.org/) 1.19+
* a POSIX compatible shell (e.g., `bash`, `ksh`, `sh`, `zsh`)
* Go development tools (`sh acquire`)

## Recommended

* [ASDF](https://asdf-vm.com/) 0.10
* [snyk](https://www.npmjs.com/package/snyk) 1.996.0 (`npm install -g snyk@1.996.0`)

# SECURITY AUDIT

```console
$ snyk test
```

# UNIT TEST

```console
$ go test
```

# COVERAGE

```console
$ mage coverageHTML
$ karp cover.html
```

# LINT

```console
$ mage lint
```

# CLEAN ALL ARTIFACTS

```console
$ mage clean; mage -clean
```
