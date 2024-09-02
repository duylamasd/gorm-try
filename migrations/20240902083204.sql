-- Drop index "idx_users_email" from table: "users"
DROP INDEX "public"."idx_users_email";
-- Create index "idx_users_email" to table: "users"
CREATE UNIQUE INDEX "idx_users_email" ON "public"."users" ((lower(email)));
