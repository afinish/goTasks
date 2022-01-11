# Go package managers (mod)

[go get](https://go.dev/ref/mod#go-get)

[go install](https://go.dev/ref/mod#go-install)

[go list -m](https://go.dev/ref/mod#go-list-m)

[go mod download](https://go.dev/ref/mod#go-mod-download)

[go mod edit](https://go.dev/ref/mod#go-mod-edit)

[go mod graph](https://go.dev/ref/mod#go-mod-graph)

[go mod init](https://go.dev/ref/mod#go-mod-init)

[go mod tidy](https://go.dev/ref/mod#go-mod-tidy)

[go mod vendor](https://go.dev/ref/mod#go-mod-vendor)

[go mod verify](https://go.dev/ref/mod#go-mod-verify)

[go mod why](https://go.dev/ref/mod#go-mod-why)

#### go get
```go
examples:
$ go get -d golang.org/x/net
# 

$ go get -d ./...
#

$ go get -d golang.org/x/net@v0.3.2
#  

$ go get -d golang.org/x/net@master
# 

$ go get -d golang.org/x/net@none
# 

```

**Description**: `go get` updates module dependencies in the go.mod file, builds and installs specified packages.

Firstly, it determines which modules have to be updated. We can specify list of packages, package patterns, and module paths. 
- Package as an argument
`go get` will update the module that provides the package. 
- Patters as an argument 
`go get` will expand the pattern to a set of packages, then updates the modules that provide the packages
- Module as an argument
`go get` will update the module without building a package
- No argument
`go get` behaves same when `.` is specified. This can be used with `-u` flag to update modules that provide imported packages

Arguments can contain version query suffix.
- @*version* (e.g. `@v3.4.2`) 
- @*version prefix* (e.g. `@v3.0`)
- @*branch/tag-name* (e.g. `@master`)
- @*revision* (e.g. `1234abcd`)
- @*patch*
- @*upgrade*
- @*latest*
- when no version is given @*upgrade* is used

Then, `go get` adds, removes, or changes the require directive in the go.mod file.

Other modules might be upgraded or downgraded when the specified module is added/upgraded/downgraded or downgraded/removed if the module requires the other modules at higher/lower versions.

Module requirement might be removed using the @*none* suffix. 

When a module is needed at two versions (specified in the command line, satisfy downgrade, upgrade), `go get` will report an error.

After `go get` has selected a new set versions, it checks whether the versions are deprecated or retracted. If so, it reports a warning. To list all the warnings `go list -m -u all` can be used.

After `go get` updates the `go.mod` file, it builds the specified packages. Executables will be installed in directory specified by `GOBIN` environment variable, which is set to `$GOPATH/bin` or `$HOME/go/bin` when the `GOPATH` environment variable is not set.

**`go get` flags**
- `-d` flag is deprecated, it is always enabled. When it is used `go get` will manage only dependencies in `go.mod` file without building or installing packages.
- `-u` upgrades modules providing packages imported directly or indirectly imported by packages specified in the command line.
- `-u=patch` acts similar to version query suffix *@patch*.  Upgrades dependencies to the latest patch version.
- `-t` consider modules need to build tests of the specified packages. In combination with `-u`, `go get` will update dependencies as well.
NOTE: `go install` is the recommended command for building and installing programs. When used with version prefix, `go install` builds packages in module-aware mode, ignoring go.mod file. 
`go get` is more focused on managing requirements in the go.mod file.
#### go install
Usage:
```go
go install [build flags] [packages]
```
Examples:
```go
$ go install golang.org/x/tools/gopls@latest
# 

$ go install golang.org/x/tools/gopls@v0.6.4
# 

$ go install golang.org/x/tools/gopls
# 

$ go install ./cmd/...
# 

```

**Description**: `go install` command builds and installs packages named by the paths on the command line. Executables (`main` packages) are installed to the directory named by the `GOBIN` environment variable, which is `$GOPATH/bin` by default or `$HOME/go/bin` when the `GOPATH` environment variables is not set.

If the arguments have version suffixes, `go install` acts in module-aware mode and ignores the go.mod file. It is useful when installing executables without affecting the dependencies of the main module.

#### go list -m all (brief)
**Description:** use this command to check all versions in the build list

#### go mod init
Usage:
```go
go mod init [module-path]
```
Example:
```go
go mod init
go mod init example.com/m
```
**Description:** `go mod init` initializes and creates a new go.mod file in the current directory. 

`init` accepts an optional argument, the *module path* for the new module.


*Module path* is a canonical name for a module, declared with the *module directive* in the module's go.mod file. A module's path is a prefix for package paths within the module 

*Module directive* defines main module's path. A go.mod file must contain only one module directive

#### go mod tidy
Usage:
```go
	go mod tidy [-e] [-v] [-go=version] [-compat=version]
```
**Description:** ensures that the go.mod file matches the source code in the module. It adds any missing module requirements necessary to build the current module's packages and dependencies, and it removes requirements on modules that do not provide any relevant packages.
Also it adds missing entries or removes unnecessary entries in the go.sum.

Flags:
- `-e` makes `go mod tidy` to attempt to proceed despite errors encountered while loading packages.
- `-v` makes `go mod tidy` print info about removed modules to standard error.

#not-complete
- [ ] go mod vendor
- [ ] go mod verify
- [ ] go mod why
