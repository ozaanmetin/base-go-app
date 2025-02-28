#!/bin/bash

# Load environment variables from .env file
if [ -f .env ]; then
  source .env
fi

# PostgreSQL connection string
connection_string="postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@$POSTGRES_HOST:$POSTGRES_PORT/$POSTGRES_DB?sslmode=disable"
migrations_path="src/database/migrations"

# Ensure we have the required arguments
if [ "$#" -ne 4 ]; then
    echo "Usage: $0 --path $migrations_path --number <migrations_to_roll_back>"
    exit 1
fi

# Parse command line arguments
while [ "$1" != "" ]; do
    case $1 in
        --number)
            shift
            MIGRATION_NUMBER=$1
            ;;
        *)
            echo "Invalid option: $1"
            exit 1
            ;;
    esac
    shift
done

# Run migrations down for the specified number
echo "Rolling back $MIGRATION_NUMBER migrations for $MIGRATION_PATH..."
migrate --path $MIGRATION_PATH --database $connection_string down $MIGRATION_NUMBER

if [ $? -ne 0 ]; then
    echo "Error rolling back migrations. Exiting."
    exit 1
fi

echo "Migrations rolled back successfully!"