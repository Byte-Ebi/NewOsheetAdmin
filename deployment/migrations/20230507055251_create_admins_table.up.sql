BEGIN;

CREATE TABLE "admins" (
  "id" BIGSERIAL PRIMARY KEY,
  "account" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "status" varchar NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp
);

COMMIT;
