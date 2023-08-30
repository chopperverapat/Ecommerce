BEGIN ;

--set timezone
SET TIME ZONE 'Asia/Bangkok';

--install extension uuid 
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

--create user_id_seqr_id > u00001
--create products_id_seq > u00001
--create orders_id_seq > u00001
CREATE SEQUENCE user_id_seq START WITH  1 INCREMENT BY 1 ;
CREATE SEQUENCE products_id_seq START WITH  1 INCREMENT BY 1 ;
CREATE SEQUENCE orders_id_seq START WITH  1 INCREMENT BY 1 ;

--Auto update because progret
CREATE OR REPLACE FUNCTION set_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';

-- create enum status 
CREATE TYPE "order_status" AS ENUM (
    'waiting',
    'shipping',
    'completed',
    'cancled'
);

CREATE TABLE "users" (
  "id" VARCHAR(7) PRIMARY KEY DEFAULT CONCAT('U',LPAD(NEXTVAL('user_id_seq')::TEXT,6,'0')),
  "username" VARCHAR UNIQUE NOT NULL,
  "password" VARCHAR  NOT NULL,
  "email" VARCHAR UNIQUE NOT NULL,
  "role_id" INT NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE "oauth" (
  "id" uuid NOT NULL UNIQUE PRIMARY KEY DEFAULT uuid_generate_v4(),
  "user_id" VARCHAR NOT NULL,
  "access_token" VARCHAR NOT NULL,
  "refresh_token" VARCHAR NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP NOT NULL DEFAULT now()
);

-- serail int will increate continous 
CREATE TABLE "roles" (
  "id" SERIAL PRIMARY KEY,
  "title" VARCHAR NOT NULL UNIQUE
);

CREATE TABLE "products" (
  "id" VARCHAR(7) PRIMARY KEY DEFAULT CONCAT('P',LPAD(NEXTVAL('products_id_seq')::TEXT,6,'0')),
  "title" VARCHAR NOT NULL,
  "description" VARCHAR NOT NULL DEFAULT '',
  "price" FLOAT NOT NULL DEFAULT 0,
  "created_at" TIMESTAMP NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE "products_catagories" (
  "id" uuid NOT NULL UNIQUE PRIMARY KEY DEFAULT uuid_generate_v4(),
  "products_id" VARCHAR NOT NULL,
  "catagories_id" INT NOT NULL
);

CREATE TABLE "catagories" (
  "id" SERIAL PRIMARY KEY,
  "title" VARCHAR UNIQUE NOT NULL
);

CREATE TABLE "images" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "filename" VARCHAR NOT NULL,
  "url" VARCHAR NOT NULL,
  "products_id" VARCHAR NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE "orders" (
  "id" VARCHAR(7) PRIMARY KEY DEFAULT CONCAT('O',LPAD(NEXTVAL('orders_id_seq')::TEXT,6,'0')),
  "user_id" VARCHAR NOT NULL,
  "contact" VARCHAR NOT NULL,
  "address" VARCHAR NOT NULL,
  "transfer_slip" jsonb,
  "status" order_status NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE "products_orders" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "orders_id" VARCHAR,
  "qty" INT,
  "product" jsonb
);

ALTER TABLE "users" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");
ALTER TABLE "oauth" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "products_catagories" ADD FOREIGN KEY ("products_id") REFERENCES "products" ("id");
ALTER TABLE "products_catagories" ADD FOREIGN KEY ("catagories_id") REFERENCES "catagories" ("id");
ALTER TABLE "images" ADD FOREIGN KEY ("products_id") REFERENCES "products" ("id");
ALTER TABLE "orders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "products_orders" ADD FOREIGN KEY ("orders_id") REFERENCES "orders" ("id");

CREATE TRIGGER set_updated_at_timestamp_users_table BEFORE UPDATE ON "users" FOR EACH ROW EXECUTE PROCEDURE set_updated_at_column();
CREATE TRIGGER set_updated_at_timestamp_oauth_table BEFORE UPDATE ON "oauth" FOR EACH ROW EXECUTE PROCEDURE set_updated_at_column();
CREATE TRIGGER set_updated_at_timestamp_products_table BEFORE UPDATE ON "products" FOR EACH ROW EXECUTE PROCEDURE set_updated_at_column();
CREATE TRIGGER set_updated_at_timestamp_images_table BEFORE UPDATE ON "images" FOR EACH ROW EXECUTE PROCEDURE set_updated_at_column();
CREATE TRIGGER set_updated_at_timestamp_orders_table BEFORE UPDATE ON "orders" FOR EACH ROW EXECUTE PROCEDURE set_updated_at_column();

COMMIT; 