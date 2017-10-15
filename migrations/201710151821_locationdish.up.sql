CREATE TABLE "location" (
    ID serial PRIMARY KEY,
    name text NOT NULL
);

CREATE TABLE "dish" (
    ID serial PRIMARY KEY,
    name text NOT NULL,
    location_id integer REFERENCES "location" (ID)
)