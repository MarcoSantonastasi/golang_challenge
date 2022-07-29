SHELL=/bin/bash

include .env

dbseed:
	psql -U $(POSTGRES_USER) -d $(POSTGRES_DB) -h $(POSTGRES_HOSTNAME) -a \
	  -f ./scripts/db/init/00-db-init.sql \
	  -f ./scripts/db/init/01-db-seed.sql
.PHONY: dbseed

protogen:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative api/arex/v1/arexApiV1.proto
.PHONY: protogen
