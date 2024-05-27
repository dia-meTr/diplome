BEGIN;


CREATE TABLE tags (
    id SERIAL PRIMARY KEY,
    name text NOT NULL
);

COMMIT;
