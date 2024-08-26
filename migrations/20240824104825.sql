-- Create "users" table
CREATE TABLE "public"."users" (
 "id" uuid NOT NULL DEFAULT gen_random_uuid(),
 "created_at" timestamptz NULL,
 "updated_at" timestamptz NULL,
 "deleted_at" timestamptz NULL,
 "name" text NOT NULL,
 "email" text NOT NULL,
 "password" text NOT NULL,
 "date_of_birth" timestamptz NOT NULL,
 "first_name" text NOT NULL,
 "last_name" text NOT NULL,
 PRIMARY KEY ("id")
);
-- Create index "idx_users_deleted_at" to table: "users"
CREATE INDEX "idx_users_deleted_at" ON "public"."users" ("deleted_at");
-- Create index "idx_users_email" to table: "users"
CREATE UNIQUE INDEX "idx_users_email" ON "public"."users" ((lower(name)));
-- Create index "idx_users_name" to table: "users"
CREATE UNIQUE INDEX "idx_users_name" ON "public"."users" ((lower(name)));
-- Create "todos" table
CREATE TABLE "public"."todos" (
 "id" uuid NOT NULL DEFAULT gen_random_uuid(),
 "created_at" timestamptz NULL,
 "updated_at" timestamptz NULL,
 "deleted_at" timestamptz NULL,
 "uid" uuid NOT NULL,
 "name" text NOT NULL,
 "description" text NOT NULL,
 "ttl" timestamptz NULL,
 PRIMARY KEY ("id"),
 CONSTRAINT "fk_users_todos" FOREIGN KEY ("uid") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_todos_deleted_at" to table: "todos"
CREATE INDEX "idx_todos_deleted_at" ON "public"."todos" ("deleted_at");
