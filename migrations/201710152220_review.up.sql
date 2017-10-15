CREATE TABLE "review" (
    ID serial PRIMARY KEY,
    user_id integer NOT NULL REFERENCES "user" (ID),
    dish_id integer NOT NULL REFERENCES "dish" (ID),
    title text NOT NULL,
    rating integer NOT NULL,
    body text NOT NULL,
    date_created timestamp NOT NULL,
    last_edited timestamp NOT NULL
)