#!/bin/sh
set -a && source .env.local && set +a

# load the migration from the mysql env in dbconfig.yml file
sqlboiler mysql --no-hooks
 