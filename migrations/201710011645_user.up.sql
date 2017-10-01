CREATE TABLE "user" (
    ID serial PRIMARY KEY,
    username text NOT NULL UNIQUE,
    first_name text NOT NULL,
    last_name text NOT NULL,
    email text NOT NULL UNIQUE,
    password_hash text NOT NULL UNIQUE
);