DROP DATABASE IF EXISTS golang_cqrs;
CREATE DATABASE golang_cqrs;
\c golang_cqrs;

DROP TABLE IF EXISTS roles;
CREATE TABLE roles (
    role_id SERIAL PRIMARY KEY,
    code VARCHAR(50) UNIQUE NOT NULL,
    role_name VARCHAR(100) NOT NULL,
    unique_id UUID DEFAULT gen_random_uuid(),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

DROP TABLE IF EXISTS users;
CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    role_id INT NOT NULL,
    user_name VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(150) NOT NULL,
    user_image VARCHAR(100), 
    is_active BOOLEAN default true,
    token VARCHAR(150) UNIQUE,
    unique_id UUID DEFAULT gen_random_uuid(),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT fk_user_role FOREIGN KEY(role_id) REFERENCES roles(role_id)
);

DROP TABLE IF EXISTS categories;
CREATE TABLE categories (
    category_id SERIAL PRIMARY KEY,
    category_name VARCHAR(100) UNIQUE,
    description VARCHAR(100) UNIQUE,
    unique_id UUID DEFAULT gen_random_uuid(),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS products;
CREATE TABLE products (
    product_id SERIAL PRIMARY KEY,
    category_id INT NOT NULL,
    code VARCHAR(30) UNIQUE NOT NULL,
    product_name VARCHAR(100) NOT NULL,
    unit_price FLOAT NOT NULL DEFAULT 0,
    image VARCHAR(150),
    description VARCHAR(200),
    unique_id UUID DEFAULT gen_random_uuid(),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_category_product FOREIGN KEY(category_id) REFERENCES categories(category_id)
);

-- VIEWS 
-- view_product_data
DROP VIEW IF EXISTS view_product_data;
CREATE VIEW view_product_data AS 
SELECT pr.product_id, pr.unique_id,
    pr.product_name, pr.code,
    pr.unit_price, pr.image,
    pr.description, 
    pr.created_at, pr.updated_at,
    ca.category_id, ca.category_name
FROM products pr
JOIN categories ca ON(ca.category_id = pr.category_id)
ORDER BY pr.created_at DESC;

-- View: view_user_data
DROP VIEW IF EXISTS view_user_data;
CREATE VIEW view_user_data AS 
SELECT us.user_id, us.unique_id,
	us.user_name, us.password, 
	us.is_active, us.user_image,
    us.token, 
	TO_CHAR(us.created_at, 'YYYY-MM-DD HH24:MI:SS') AS created_at,
    TO_CHAR(us.updated_at, 'YYYY-MM-DD HH24:MI:SS') AS updated_at,
	ro.role_id, ro.role_name,
	ro.code AS role_code
FROM users us
JOIN roles ro ON(ro.role_id = us.role_id)
ORDER BY created_at DESC;
