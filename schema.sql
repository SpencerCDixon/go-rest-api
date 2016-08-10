DROP TABLE IF EXISTS "todos";

CREATE TABLE "todos" (
  id serial PRIMARY KEY,
  name varchar(255) NOT NULL,
  complete boolean
)
