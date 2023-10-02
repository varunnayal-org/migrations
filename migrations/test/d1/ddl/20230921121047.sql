-- Create enum type "user_status_enum"
CREATE TYPE "public"."user_status_enum" AS ENUM ('active', 'inactive');
-- Modify "users" table
ALTER TABLE "public"."users" ADD COLUMN "status" "public"."user_status_enum" NOT NULL;
