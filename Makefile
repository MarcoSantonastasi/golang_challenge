SHELL=/bin/sh

include .env

dbseed:
	psql -U $POSTGRES_USER -d $POSTGRES_DB -a -f ./scripts/00-db-init.sql ./scripts/01-db-seed.sql