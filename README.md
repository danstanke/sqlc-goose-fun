# sqlc-goose-fun
Playing with goose/sqlc/postgres/go

# Install

## `goose`

```shell
go install github.com/pressly/goose/v3/cmd/goose@latest
```

### usage 

```shell
goose -dir ./migrations create create_database sql
```

for project:
```shell
go get github.com/pressly/goose/v3
```

### why not `go get`?

`go get github.com/pressly/goose/v3`
go: go.mod file not found in current directory or any parent directory.
	'go get' is no longer supported outside a module.
	To build and install a command, use 'go install' with a version,
	like 'go install example.com/cmd@latest'
	For more information, see https://golang.org/doc/go-get-install-deprecation
	or run 'go help get' or 'go help install'.

`go install github.com/pressly/goose/v3`
go: 'go install' requires a version when current directory is not in a module
	Try 'go install github.com/pressly/goose/v3@latest' to install the latest version

## `sqlc`

> sqlc `github.com/sqlc-dev/sqlc v1.27.0` uses `github.com/google/uuid v1.6.0`

```shell
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```
