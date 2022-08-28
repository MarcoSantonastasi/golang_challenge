SHELL=/bin/bash

ENV := $(PWD)/.env
-include $(ENV)
export

.PHONY: demo
demo:
	make proddbseed
	go run ./cmd/server
	go run ./cmd/client


.PHONY: e2etests
e2etests:
	make testingdbseed
	go test -v ./tests/e2e/...

.PHONY: unittests
unittests:
	make stubdbseed
	go test -v ./internal/...


.PHONY: protogen
protogen:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative api/arex/v1/arexApiV1.proto


.PHONY: proddbseed
proddbseed:
	psql -U $$POSTGRES_USER -d $$POSTGRES_PROD_DB -h $$POSTGRES_HOSTNAME -a \
	  -f ./scripts/db/init/00-db-init.sql \
	  -f ./scripts/db/init/01-db-seed.sql


.PHONY: testingdbseed
testingdbseed:
	psql -U $$POSTGRES_USER -d $$POSTGRES_PROD_DB -h $$POSTGRES_HOSTNAME -a \
	  -c "DROP DATABASE IF EXISTS $$POSTGRES_TESTING_DB;" \
	  -c "CREATE DATABASE $$POSTGRES_TESTING_DB;"

	psql -U $$POSTGRES_USER -d $$POSTGRES_TESTING_DB -h $$POSTGRES_HOSTNAME -a \
	  -f ./scripts/db/init/00-db-init.sql \
	  -f ./scripts/db/init/01-db-seed.sql


.PHONY: stubdbseed
stubdbseed:
	psql -U $$POSTGRES_USER -d $$POSTGRES_PROD_DB -h $$POSTGRES_HOSTNAME -a \
	  -c "DROP DATABASE IF EXISTS $$POSTGRES_STUB_DB;" \
	  -c "CREATE DATABASE $$POSTGRES_STUB_DB;"

	psql -U $$POSTGRES_USER -d $$POSTGRES_STUB_DB -h $$POSTGRES_HOSTNAME -a \
	  -f ./scripts/db/init/00-db-init.sql \
	  -f ./scripts/db/init/01-db-seed.sql
