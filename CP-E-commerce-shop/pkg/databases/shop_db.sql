CREATE TABLE "users" (
  "id" varchar PRIMARY KEY,
  "username" varchar UNIQUE,
  "password" varchar,
  "email" varchar,
  "role_id" int,
  "create_at" timestamp,
  "update_at" timestamp
);

CREATE TABLE "oauth" (
  "id" varchar PRIMARY KEY,
  "user_id" varchar,
  "access_token" varchar,
  "refresh_token" varchar,
  "create_at" timestamp,
  "update_at" timestamp
);

CREATE TABLE "roles" (
  "id" int PRIMARY KEY,
  "title" varchar
);

CREATE TABLE "products" (
  "id" varchar PRIMARY KEY,
  "title" varchar,
  "description" varchar,
  "price" float,
  "create_at" timestamp,
  "update_at" timestamp
);

CREATE TABLE "products_catagories" (
  "id" varchar PRIMARY KEY,
  "products_id" varchar,
  "catagories_id" int
);

CREATE TABLE "catagories" (
  "id" int PRIMARY KEY,
  "title" varchar UNIQUE
);

CREATE TABLE "images" (
  "id" varchar PRIMARY KEY,
  "file_name" varchar,
  "url" varchar,
  "products_id" varchar,
  "create_at" timestamp,
  "update_at" timestamp
);

CREATE TABLE "orders" (
  "id" varchar PRIMARY KEY,
  "user_id" varchar,
  "contract" varchar,
  "address" varchar,
  "tranfer_slip" jsonb,
  "status" varchar,
  "create_at" timestamp,
  "update_at" timestamp
);

CREATE TABLE "products_orders" (
  "id" varchar PRIMARY KEY,
  "orders_id" varchar,
  "qty" int,
  "product" jsonb
);

ALTER TABLE "users" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");

ALTER TABLE "oauth" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "products_catagories" ADD FOREIGN KEY ("products_id") REFERENCES "products" ("id");

ALTER TABLE "products_catagories" ADD FOREIGN KEY ("catagories_id") REFERENCES "catagories" ("id");

ALTER TABLE "images" ADD FOREIGN KEY ("products_id") REFERENCES "products" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "products_orders" ADD FOREIGN KEY ("orders_id") REFERENCES "orders" ("id");
