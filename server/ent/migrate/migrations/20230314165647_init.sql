-- create "users" table
CREATE TABLE "users" ("id" uuid NOT NULL, "create_time" timestamptz NOT NULL, "update_time" timestamptz NOT NULL, "email" character varying NOT NULL, "slug" character varying NOT NULL, "ga" character varying NULL, PRIMARY KEY ("id"));
-- create index "users_email_key" to table: "users"
CREATE UNIQUE INDEX "users_email_key" ON "users" ("email");
-- create index "users_slug_key" to table: "users"
CREATE UNIQUE INDEX "users_slug_key" ON "users" ("slug");
-- create "services" table
CREATE TABLE "services" ("id" uuid NOT NULL, "create_time" timestamptz NOT NULL, "update_time" timestamptz NOT NULL, "service" character varying NOT NULL, "url" character varying NOT NULL, "user_uuid" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "services_users_uuid" FOREIGN KEY ("user_uuid") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
