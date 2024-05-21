INSERT INTO users (first_name, last_name, email, avatar_url, address, is_admin)
VALUES ('John', 'Doe', 'john.doe@example.com', 'http://example.com/avatar.jpg', '123 Main St', false);

INSERT INTO products (name, description)
VALUES ('tomatoes', NULL),
    ('cucumbers', NULL),
    ('oil', NULL);

INSERT INTO diets (name, description)
VALUES ('Keto Diet', 'A low-carb, high-fat diet'),
    ('Mediterranean Diet', 'Emphasizes fruits, vegetables, whole grains, and healthy fats'),
    ('Paleo Diet', 'Based on the types of foods presumed to have been eaten by early humans'),
    ('Vegetarian Diet', 'Excludes meat and sometimes other animal products'),
    ('Vegan Diet', 'Excludes all animal products');

INSERT INTO tags (name)
VALUES ('vegan'),
    ('vegetarian');