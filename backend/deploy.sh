#!/bin/sh

# eval $(minikube docker-env)

# build backend image
docker build -t haisin-official/backend:latest -t haisin-official/backend:v1 -f docker/production/app/Dockerfile .
docker build -t haisin-official/migration:latest -t haisin-official/migration:v1 -f docker/production/migration/Dockerfile .

# load images to minikube
# minikube image load haisin-official/backend:v1
# minikube image load haisin-official/migration:v1

# apply kubernetes
kubectl delete -f database-secret.yaml
kubectl delete -f redis-secret.yaml
kubectl delete -f migration-secret.yaml
kubectl delete -f app-secret.yaml

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
kubectl delete -f migration-job.yaml
kubectl apply -f migration-job.yaml

# Execute golang apps
kubectl apply -f app-deployment.yaml
kubectl apply -f app-service.yaml

# Apply ingress
kubectl apply -f https://raw.githubusercontent.com/traefik/traefik/v2.9/docs/content/reference/dynamic-configuration/kubernetes-crd-definition-v1.yml
kubectl apply -f https://raw.githubusercontent.com/traefik/traefik/v2.9/docs/content/reference/dynamic-configuration/kubernetes-crd-rbac.yml
kubectl apply -f app-ingress.yaml