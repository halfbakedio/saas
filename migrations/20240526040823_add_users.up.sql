-- create "users" table
CREATE TABLE "users" (
  "id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY,
  "email" character varying NOT NULL DEFAULT 'unknown',
  PRIMARY KEY ("id")
);