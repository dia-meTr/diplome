BEGIN;


CREATE TABLE tags (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name text NOT NULL
);

COMMIT;
