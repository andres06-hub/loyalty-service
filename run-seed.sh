#!/bin/bash

# Bash script information
TITLE="Loyalty Database Seeding"
VERSION=v1.0.0

echo "$TITLE $VERSION"

# Usage
if [ $# -eq 0 ] || [ $1 == "--help" ] || [ $1 == "-h" ] || [ $1 == "help" ] || [ $1 == "-help" ]; then
    echo "Usage: ./run-seed.sh [dev|prod]"
    echo 'Env variables: 
        APP_DATABASE_PASSWORD: Postgres password
        APP_DATABASE_HOST: Postgres host
        APP_DATABASE_NAME: Postgres database name
        APP_DATABASE_USER: Postgres user
    '
    exit
fi

MODE=$1

# Check if variables are set, if not, exit with an error message
: "${APP_DATABASE_PASSWORD:?APP_DATABASE_PASSWORD is not set}"
: "${APP_DATABASE_HOST:?APP_DATABASE_HOST is not set}"
: "${APP_DATABASE_NAME:?APP_DATABASE_NAME is not set}"
: "${APP_DATABASE_USER:?APP_DATABASE_USER is not set}"

if [ $MODE == "dev" ]; then
    CONN_STRING=postgresql://$APP_DATABASE_USER:$APP_DATABASE_PASSWORD@$APP_DATABASE_HOST/$APP_DATABASE_NAME
else
    echo "Error: unrecognized option \"$MODE\"."
    exit
fi

echo "Current execution mode is: $MODE"
echo "Current connection string is: $CONN_STRING"
echo -e "\n"

function _seed() {
    echo "Seeding loyalty POSTGRES database started."

    SEEDERS=$(ls -l ./seeders | grep ^- | grep -Po "[\w\-\_]*\.sql" | sed "s;^;./seeders/;")

    for f in $SEEDERS; do
        # Get only filename without ext
        # So example.sql -> example
        FILENAME=$(basename $f .sql)
        echo -e "\nSeeding $FILENAME..."
        psql $CONN_STRING -f $f
    done
    echo -e "\nSeeding complete."
}

# Run the seed script
_seed