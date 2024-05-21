BEGIN;

CREATE TABLE diets (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name text NOT NULL,
    description text
);

COMMIT;
