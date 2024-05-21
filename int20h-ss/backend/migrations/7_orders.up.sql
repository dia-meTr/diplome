BEGIN;

CREATE TABLE orders (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id uuid REFERENCES users(id) on delete cascade,
    order_date timestamp NOT NULL DEFAULT NOW()
);

CREATE TABLE orders_dishes (
    order_id uuid REFERENCES orders(id) on delete cascade,
    dish_id uuid REFERENCES dishes(id),
    quantity integer NOT NULL,
    PRIMARY KEY (order_id, dish_id)
);

COMMIT;