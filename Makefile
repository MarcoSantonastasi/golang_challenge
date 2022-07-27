SHELL=/bin/bash

include .env

dbseed:
	psql -U $(POSTGRES_USER) -d $(POSTGRES_DB) -h $(POSTGRES_HOSTNAME) -a \
	  -f ./db/init/00-db-init.sql \
	  -f ./db/init/01-db-seed.sql
.PHONY: dbseed

protogen:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/arex.proto
.PHONY: protogen
