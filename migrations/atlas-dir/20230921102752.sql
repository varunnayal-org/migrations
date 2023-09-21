-- Create enum type "category_name_enum"
CREATE TYPE "public"."category_name_enum" AS ENUM ('electricity', 'water', 'postpaid', 'loan');
-- Create "categories" table
CREATE TABLE "public"."categories" (
  "id" uuid NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "is_active" boolean NULL DEFAULT true,
  "name" character varying(100) NOT NULL,
  "type" "public"."category_name_enum" NULL,
  PRIMARY KEY ("id")
);
