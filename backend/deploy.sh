#!/bin/sh

# build backend image
docker build -t haisin-official/backend:latest -f docker/production/app/Dockerfile .
docker push backend

# apply kubernetes
kubectl apply -f database-secret.yaml
kubectl apply -f redis-secret.yaml
kubectl apply -f migration-secret.yaml
kubectl apply -f app-secret.yaml

# Execute database
kubectl apply -f database-pv.yaml
kubectl apply -f database-pvc.yaml
kubectl apply -f database-cm.yaml
kubectl apply -f database-deployment.yaml
kubectl apply -f database-service.yaml

# Execute redis
kubectl apply -f redis-pv.yaml
kubectl apply -f redis-pvc.yaml
kubectl apply -f redis-deployment.yaml
kubectl apply -f redis-service.yaml

# Execute migrations
kubectl apply -f migration-deployment.yaml
kubectl apply -f migration-service.yaml

# Execute golang apps
kubectl apply -f app-deployment.yaml
kubectl apply -f app-service.yaml