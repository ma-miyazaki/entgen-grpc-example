# ent-grpc-example

Example app that using ent and gRPC.

## Create database.
```
docker-compose exec go 'sh' '-c' 'sqlite3 users.sqlite < users.sql'
```

## Generate ent's schema
```
docker-compose exec go entgen -driver sqlite3 -dsn ./users.sqlite -rplural
```
__Notice__
entgen cannot output collect types from sqlite. So, overwrite field types in _./ent/schema/user.go_ .
