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

## Running migrations

```shell
goose postgres "user=myuser password=mypassword dbname=mydb sslmode=disable" up
```

## UUID

```sql
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
```

## Generating OpenAPI code

### install
```shell
go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
```

### addd
```yaml
package: api
generate:
  std-http-server: true
  models: true
output: gen.go
```

```shell
oapi-codegen -generate types,std-http -package api -o api/server.gen.go api/openapi.yaml
```

```
/endpoint/{id}:
    get:
      operationId: getById
      parameters:
        - name: id
          in: path
          required: true
          schema:
            type: integer
            format: int32
      responses:
        '200':
          description: Successful retrieval by ID
          content:
            application/json:
              schema:
                type: string
                example: "Successfully retrieved the entity by ID."
  /endpoint/{uuid}:
    get:
      operationId: getByUuid
      parameters:
        - name: uuid
          in: path
          required: true
          schema:
            type: string
            format: uuid
      responses:
        '200':
          description: Successful retrieval by UUID
          content:
            application/json:
              schema:
                type: string
                example: "Successfully retrieved the entity by UUID."
```

apply migrations
```shell
goose -dir migrations postgres "user=myuser dbname=mydb sslmode=disable" up
```