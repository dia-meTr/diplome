BEGIN;

CREATE TABLE IF NOT EXISTS dishes (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name text NOT NULL,
    photo_url text,
    description text
);

CREATE TABLE IF NOT EXISTS dish_tags (
    dish_id uuid REFERENCES dishes(id) on delete cascade,
    tag_id uuid REFERENCES tags(id) on delete cascade,
    PRIMARY KEY (dish_id, tag_id)
);

CREATE TABLE IF NOT EXISTS dish_diets (
    dish_id uuid REFERENCES dishes(id) on delete cascade,
    diet_id uuid REFERENCES diets(id) on delete cascade,
    PRIMARY KEY (dish_id, diet_id)
);

CREATE TABLE IF NOT EXISTS dish_products (
    dish_id uuid REFERENCES dishes(id) on delete cascade,
    product_id uuid REFERENCES products(id) on delete cascade,
    amount integer NOT NULL,
    PRIMARY KEY (dish_id, product_id)
);

COMMIT;