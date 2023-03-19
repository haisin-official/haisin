# docker build -t haisin-official/migration:latest -f docker/production/migration/Dockerfile .
FROM arigaio/atlas:latest

COPY ./ent/migrate/migrations /migrations

ENTRYPOINT ["migrate", "apply", "--url", "${DB_URL}"]