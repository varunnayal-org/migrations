-- Create enum type "category_name_enum"
CREATE TYPE "public"."category_name_enum" AS ENUM ('electricity', 'water', 'postpaid', 'loan');
-- Create enum type "user_status_enum"
CREATE TYPE "public"."user_status_enum" AS ENUM ('active', 'inactive');
-- Create "categories" table
CREATE TABLE "public"."categories" (
  "id" uuid NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "is_active" boolean NULL DEFAULT true,
  "name" character varying(100) NOT NULL,
  "type" "public"."category_name_enum" NULL,
  "desc" character varying(100) NOT NULL,
  PRIMARY KEY ("id")
);
-- Create "users" table
CREATE TABLE "public"."users" (
  "id" uuid NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "name" character varying(100) NOT NULL,
  "email" character varying(100) NOT NULL,
  "phone" character varying(15) NOT NULL,
  "status" "public"."user_status_enum" NOT NULL,
  "dob" date NOT NULL,
  "sec_phone" character varying(15) NOT NULL,
  PRIMARY KEY ("id")
);
