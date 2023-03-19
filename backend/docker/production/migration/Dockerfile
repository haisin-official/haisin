# docker build -t haisin-official/migration:latest -t haisin-official/migration:v1 -f docker/production/migration/Dockerfile .
FROM arigaio/atlas:latest

COPY ./ent/migrate/migrations /migrations

CMD ["migrate", "apply", "--url", "${DB_URL}"]