#!/bin/sh

# build backend image
docker build -t haisin_official/backend:latest -f docker/production/app/Dockerfile .
docker push backend
