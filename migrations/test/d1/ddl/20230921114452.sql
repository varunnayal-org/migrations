-- Modify "users" table
ALTER TABLE "public"."users" ADD COLUMN "phone" character varying(15) NOT NULL, ADD COLUMN "dob" date NOT NULL;
