#!/bin/bash

if [ -f .env ]; then
  export $(grep -v '^#' .env | xargs)
else
  echo ".env file not found!"
  exit 1
fi

LAMBDA_NAME="leaguies-backend-dev-api"

echo "Updating Lambda environment variables for function: $LAMBDA_NAME"

aws lambda update-function-configuration \
  --function-name "$LAMBDA_NAME" \
  --environment Variables="{DB_HOST=$DB_HOST,DB_PORT=$DB_PORT,DB_USER=$DB_USER,DB_PASSWORD=$DB_PASSWORD,DB_NAME=$DB_NAME,DB_SSLMODE=$DB_SSLMODE}"

if [ $? -eq 0 ]; then
  echo "Lambda environment variables updated successfully!"
else
  echo "Failed to update Lambda environment variables."
fi
