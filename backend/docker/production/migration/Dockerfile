# docker build -t haisin-official/migration:latest -t haisin-official/migration:k8s -f docker/production/migration/Dockerfile .
FROM arigaio/atlas:latest

COPY ./ent/migrate/migrations /migrations

ENTRYPOINT ["migrate", "apply", "--url", "${DB_URL}"]