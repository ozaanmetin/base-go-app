#!/bin/bash

# Ensure Go binary is available in PATH
export PATH=$PATH:/usr/local/go/bin

# Set the Go module path
export GO111MODULE=on

# Load environment variables from .env file
if [ -f .env ]; then
  echo "Loading environment variables from .env file"
  source .env
fi

# PostgreSQL connection string
connection_string="postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@$POSTGRES_HOST:$POSTGRES_PORT/$POSTGRES_DB?sslmode=disable"

# Define the location of the migrations for each app
declare -a MIGRATION_DIRS=(
    "src/database/migrations/users"
    # Add more apps here
)



# Function to run migrations for each app
run_migrations() {
    for dir in "${MIGRATION_DIRS[@]}"; do
        echo "Running migrations for $dir..."

        migrate --path $dir --database $connection_string up

        if [ $? -ne 0 ]; then
            echo "Error running migrations for $dir. Exiting."
            exit 1
        fi
    done
}

# Run migrations
run_migrations

echo "All migrations have been successfully applied!"