SHELL=/user/bin/bash

include .env

dbseed:
	psql -U $(POSTGRES_USER) -d $(POSTGRES_DB) -a -f ./scripts/00-db-init.sql ./scripts/01-db-seed.sql

protogen:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/arex.proto
