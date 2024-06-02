
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table users
(
    id         uuid          default uuid_generate_v4(),
    first_name text not null,
    last_name  text not null default '',
    email      text not null unique,
    user_role  text not null,
    primary key (id)
);

CREATE TABLE products (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name text NOT NULL,
    description text
);

CREATE TABLE tags (
    id SERIAL PRIMARY KEY,
    name text NOT NULL
);

CREATE TABLE IF NOT EXISTS dishes (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name text NOT NULL,
    photo_url text,
    description text,
    price float NOT NULL
);

CREATE TABLE IF NOT EXISTS dish_tags (
    dish_id uuid REFERENCES dishes(id) on delete cascade,
    tag_id integer REFERENCES tags(id) on delete cascade,
    PRIMARY KEY (dish_id, tag_id)
);

CREATE TABLE IF NOT EXISTS dish_products (
    dish_id uuid REFERENCES dishes(id) on delete cascade,
    product_id uuid REFERENCES products(id) on delete cascade,
    amount integer NOT NULL,
    PRIMARY KEY (dish_id, product_id)
);

CREATE TABLE shopping_cards (
    user_id uuid REFERENCES users(id) on delete cascade,
    dish_id uuid REFERENCES dishes(id) on delete cascade,
    amount integer NOT NULL,
    PRIMARY KEY (user_id, dish_id)
);

CREATE TABLE orders (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id uuid REFERENCES users(id) on delete cascade,
    order_date timestamp NOT NULL DEFAULT NOW(),
    delivery_date date NOT NULL DEFAULT NOW(),
    delivery_address text NOT NULL,
    order_status text NOT NULL
);

CREATE TABLE orders_dishes (
    order_id uuid REFERENCES orders(id) on delete cascade,
    dish_id uuid REFERENCES dishes(id),
    amount integer NOT NULL,
    PRIMARY KEY (order_id, dish_id)
);


INSERT INTO products (name, description)
VALUES ('tomatoes', NULL),
    ('cucumbers', NULL),
    ('oil', NULL);

INSERT INTO tags (name)
VALUES ('vegan'),
    ('vegetarian'),
    ('breakfest'),
    ('lunch'),
    ('diner'),
    ('glutenfree');

INSERT INTO dishes (name, photo_url, description, price)
VALUES 
 ('Steak'   , 'https://bestofood-storage.s3.us-east-2.amazonaws.com/food-basic.jpg', 'Steak Description',   100),
 ('Salad'   , 'https://bestofood-storage.s3.us-east-2.amazonaws.com/food-basic.jpg', 'Salad Description',   150),
 ('Soup'    , 'https://bestofood-storage.s3.us-east-2.amazonaws.com/food-basic.jpg', 'Soup Description',  200),
 ('Soup2'   , 'https://bestofood-storage.s3.us-east-2.amazonaws.com/food-basic.jpg', 'Soup2 Description',   200),
 ('Soup3'   , 'https://bestofood-storage.s3.us-east-2.amazonaws.com/food-basic.jpg', 'Soup3 Description',   200),
 ('Scrumble', 'https://bestofood-storage.s3.us-east-2.amazonaws.com/food-basic.jpg', 'Scrumble Description',200),
 ('Pancakes', 'https://bestofood-storage.s3.us-east-2.amazonaws.com/food-basic.jpg', 'Pancakes Description', 80);
