CREATE TABLE "users" (
  "id" uuid UNIQUE NOT NULL,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "email" varchar NOT NULL
);

CREATE TABLE "book" (
  "id" uuid UNIQUE NOT NULL,
  "name" varchar NOT NULL,
  "image" text 
);

CREATE TABLE "fav_book" (
  "book_id" uuid NOT NULL,
  "user_id" uuid NOT NULL,
  PRIMARY KEY ("book_id", "user_id")
);

ALTER TABLE "fav_book" ADD FOREIGN KEY ("book_id") REFERENCES "book" ("id");

ALTER TABLE "fav_book" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
