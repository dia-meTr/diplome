BEGIN;

CREATE TABLE shopping_cart (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    dish_id uuid REFERENCES dishes(id) on delete cascade,
);

COMMIT;
