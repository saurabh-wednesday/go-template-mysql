#!/bin/sh
set -a && source .env.local && set +a

# load the migration from the mysql env in dbconfig.yml file
sql-migrate up -env mysql 