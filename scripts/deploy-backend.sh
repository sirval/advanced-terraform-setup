#!/bin/bash
set -e

echo "Deploying backend..."

if [ -z "$ECR_REPOSITORY" ]; then
  echo "ECR_REPOSITORY is required"
  exit 1
fi

cd backend/MuchToDo

docker build -t "$ECR_REPOSITORY:latest" .
docker push "$ECR_REPOSITORY:latest"

echo "Backend image built and pushed successfully"