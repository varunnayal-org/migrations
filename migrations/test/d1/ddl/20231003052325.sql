-- Modify "friends" table
ALTER TABLE "public"."friends" ADD COLUMN "friend_id" uuid NOT NULL, ADD
 CONSTRAINT "fk_friends_friend" FOREIGN KEY ("friend_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD
 CONSTRAINT "fk_friends_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
