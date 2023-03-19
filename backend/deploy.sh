#!/bin/sh

# build backend image
docker build -t haisin-official/backend:latest -f docker/production/app/Dockerfile .
docker push backend

# apply kubernetes
kubectl apply -f database-secret.yaml
kubectl apply -f redis-secret.yaml
kubectl apply -f app-secret.yaml

kubectl apply -f app-