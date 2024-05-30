BEGIN;

CREATE TABLE shopping_cards (
    user_id uuid REFERENCES users(id) on delete cascade,
    dish_id uuid REFERENCES dishes(id) on delete cascade,
    amount integer NOT NULL,
    PRIMARY KEY (user_id, dish_id)
);

COMMIT;
