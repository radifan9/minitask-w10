--- --- --- Create Database --- --- ---
create database simple_store owner radif;


CREATE TABLE "users" (
  "id" uuid default gen_random_uuid() PRIMARY KEY,
  "email" text unique not null ,
  "password" text not null,
  "created_at" timestamptz default current_timestamp,
  "updated_at" timestamptz default current_timestamp
);

CREATE TABLE "products" (
  "id" serial PRIMARY KEY,
  "name" text,
  "price" int
);

CREATE TABLE "transactions" (
  "id" UUID default gen_random_uuid() PRIMARY KEY,
  "user_id" uuid,
  "product_id" int,
  "created_at" timestamptz default current_timestamp,
  "updated_at" timestamptz default current_timestamp
);

ALTER TABLE "transactions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");
