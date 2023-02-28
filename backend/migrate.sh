#!/bin/sh

# Generate migration file
go generate ./ent

atlas migrate diff init \
  --dir "file://ent/migrate/migrations" \
  --to "ent://ent/schema" \
  --dev-url "docker://postgres/14/haisin?search_path=public"

source ./.env
atlas migrate apply \
  --dir "file://ent/migrate/migrations" \
  --url ${DB_URL}