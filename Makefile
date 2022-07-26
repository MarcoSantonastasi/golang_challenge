SHELL=/bin/bash

include .env

dbseed:
	psql -U $(POSTGRES_USER) -d $(POSTGRES_DB) -h $(POSTGRES_HOSTNAME) -a \
	  -f ./scripts/00-db-init.sql \
	  -f ./scripts/01-db-seed.sql

protogen:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/arex.proto
