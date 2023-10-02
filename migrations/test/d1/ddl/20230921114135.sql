-- Create "users" table
CREATE TABLE "public"."users" (
  "id" uuid NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "name" character varying(100) NOT NULL,
  "email" character varying(100) NOT NULL,
  PRIMARY KEY ("id")
);
