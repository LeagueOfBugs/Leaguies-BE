#!/bin/bash

echo "Building Go binary for Lambda..."
GOOS=linux GOARCH=amd64 go build -o bootstrap cmd/main.go
if [ $? -ne 0 ]; then
  echo "Build failed"
  exit 1
fi

echo "Deploying with Serverless..."
serverless deploy --aws-profile bagel-site
if [ $? -ne 0 ]; then
  echo "Deploy failed"
  exit 1
fi

echo "Deployment complete!"
