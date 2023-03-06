-- create "users" table
CREATE TABLE "users" ("user_id" uuid NOT NULL, "create_time" timestamptz NOT NULL, "update_time" timestamptz NOT NULL, "email" character varying NOT NULL, "slug" character varying NOT NULL, "ga" character varying NULL, PRIMARY KEY ("user_id"));
-- create index "users_email_key" to table: "users"
CREATE UNIQUE INDEX "users_email_key" ON "users" ("email");
-- create index "users_slug_key" to table: "users"
CREATE UNIQUE INDEX "users_slug_key" ON "users" ("slug");
-- create "urls" table
CREATE TABLE "urls" ("uuid" uuid NOT NULL, "create_time" timestamptz NOT NULL, "update_time" timestamptz NOT NULL, "service" character varying NOT NULL, "url" character varying NOT NULL, "user_id" uuid NOT NULL, PRIMARY KEY ("uuid"), CONSTRAINT "urls_users_id" FOREIGN KEY ("user_id") REFERENCES "users" ("user_id") ON UPDATE NO ACTION ON DELETE CASCADE);
