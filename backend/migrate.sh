#!/bin/sh
# Generate Ent files before atlas
go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/upsert ./ent/schema

# description by ent
# go run -mod=mod entgo.io/ent/cmd/ent describe ./ent/schema
atlas migrate diff init \
  --dir "file://ent/migrate/migrations" \
  --to "ent://ent/schema" \
  --dev-url "docker://postgres/14/haisin?search_path=public"

. ./.env
atlas migrate apply \
  --dir "file://ent/migrate/migrations" \
  --url ${DB_URL}