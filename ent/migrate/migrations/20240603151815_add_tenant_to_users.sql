-- Modify "organizations" table
ALTER TABLE "organizations" ADD COLUMN "user_tenant" bigint NULL, ADD
 CONSTRAINT "organizations_users_tenant" FOREIGN KEY ("user_tenant") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
