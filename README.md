# ent-grpc-example

Example app that using ent and gRPC.

## Create database

```sh
docker-compose exec go 'sh' '-c' 'sqlite3 users.sqlite < users.sql'
```

## Generate ent schema

```sh
docker-compose exec go entgen -driver sqlite3 -dsn ./users.sqlite -rplural
```
__Notice__
entgen cannot output collect types from sqlite. So, overwrite field types in _./ent/schema/user.go_ .
Add ID field if the type is not _int_ .

## Generate ent assets

Write _ent/generate.go_ .
```go:ent/generate.go
package ent

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate ./schema
```

Generate asstes
```sh
docker-compose exec go go generate ./ent
```

## Generate proto file

Add a last line to _ent/generate.go_ .
```go:ent/generate.go
//go:generate go run -mod=mod entgo.io/contrib/entproto/cmd/entproto -path ./schema
```
