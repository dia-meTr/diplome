INSERT INTO users (first_name, last_name, email, avatar_url, address, is_admin)
VALUES ('John', 'Doe', 'john.doe@example.com', 'http://example.com/avatar.jpg', '123 Main St', false);

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


INSERT INTO dish_tags (dish_id, tag_id)
VALUES 
('75f41eb4-a31e-4a96-85ef-c7ef7c5cd02d', 3),
('a4f89ea0-f605-4b24-8e87-8ef2eb43023f', 4),
('82e4350e-7c9f-40dd-81f4-a139ca2266bd', 4),
('7e17cd81-c198-4cb4-bc23-c9eb3f5648fe', 4),
('7e17cd81-c198-4cb4-bc23-c9eb3f5648fe', 1),
('7e17cd81-c198-4cb4-bc23-c9eb3f5648fe', 2),
('82e4350e-7c9f-40dd-81f4-a139ca2266bd', 2),
('25df769c-b09f-4b2e-892c-194f5cf717a5', 3),
('25df769c-b09f-4b2e-892c-194f5cf717a5', 2),
('5b27e998-a282-4870-b3f3-728378c88c89', 3),
('5b27e998-a282-4870-b3f3-728378c88c89', 1),
('5b27e998-a282-4870-b3f3-728378c88c89', 2),
('d7d32a5b-ec2b-4927-8ce4-4b956d190f35', 1),
('d7d32a5b-ec2b-4927-8ce4-4b956d190f35', 2),
('25df769c-b09f-4b2e-892c-194f5cf717a5', 4);
