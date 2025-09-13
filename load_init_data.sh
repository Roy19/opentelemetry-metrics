#!/bin/bash
# Script to load initial cart data into Postgres

# Usage: ./load_init_data.sh <PGHOST> <PGPORT> <PGUSER> <PGDATABASE>

set -e

if [ "$#" -ne 4 ]; then
  echo "Usage: $0 <PGHOST> <PGPORT> <PGUSER> <PGDATABASE>"
  exit 1
fi

PGHOST=$1
PGPORT=$2
PGUSER=$3
PGDATABASE=$4

psql \
  --host="$PGHOST" \
  --port="$PGPORT" \
  --username="$PGUSER" \
  --dbname="$PGDATABASE" \
  --file="$(dirname "$0")/init_data.sql"
