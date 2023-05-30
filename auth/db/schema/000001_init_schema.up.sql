CREATE TYPE "type" AS ENUM (
  'admin',
  'blogger',
  'regular'
);

CREATE TABLE "users" (
  "id" int PRIMARY KEY,
  "full_name" varchar NOT NULL,
  "username" varchar UNIQUE NOT NULL,
  "email" varchar NOT NULL,
  "password" varchar NOT NULL,
  "active" integer DEFAULT 0,
  "type" type,
  "last_login" timestamp DEFAULT (now())
);

COMMENT ON COLUMN "users"."full_name" IS 'User first + last names';