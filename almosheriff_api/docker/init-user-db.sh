#!/bin/bash
set -e

ps ql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
	CREATE DATABASE almosheriff;
	GRANT ALL PRIVILEGES ON DATABASE almosheriff TO "$POSTGRES_USER";
EOSQL
