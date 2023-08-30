BEGIN;

TRUNCATE TABLE "users" CASCADE;
TRUNCATE TABLE "oauth" CASCADE;
TRUNCATE TABLE "roles" CASCADE;
TRUNCATE TABLE "products" CASCADE;
TRUNCATE TABLE "catagories" CASCADE;
TRUNCATE TABLE "products_catagories" CASCADE;
TRUNCATE TABLE "images" CASCADE;
TRUNCATE TABLE "orders" CASCADE;
TRUNCATE TABLE "products_orders" CASCADE;

SELECT SETVAL ((SELECT PG_GET_SERIAL_SEQUENCE('"roles"', 'id')), 1, FALSE);
SELECT SETVAL ((SELECT PG_GET_SERIAL_SEQUENCE('"catagories"', 'id')), 1, FALSE);

COMMIT;