CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    firstname VARCHAR(100) NOT NULL,
    lastname VARCHAR(100) NOT NULL,
    fullname VARCHAR(200) NOT NULL,
    age INT NOT NULL,
    is_married BOOLEAN NOT NULL,
    passwordhash VARCHAR(255) NOT NULL,
	created_at timestamptz DEFAULT NOW(),
	updated_at timestamptz DEFAULT NOW()
);

ALTER TABLE users ADD CONSTRAINT age_check CHECK (age >= 0);

CREATE TABLE products (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    description TEXT NOT NULL,
    quantity BIGINT NOT NULL,
	created_at timestamptz DEFAULT NOW(),
	updated_at timestamptz DEFAULT NOW()
);

ALTER TABLE products ADD CONSTRAINT quantity_check CHECK (quantity >= 0);

CREATE TABLE tags (
    id SERIAL PRIMARY KEY,
    tag VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE product_tags (
    product_id UUID REFERENCES products(id),
    tag_id INT REFERENCES tags(id),
    PRIMARY KEY (product_id, tag_id)
);

CREATE TABLE orders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "status" VARCHAR(100) NOT NULL,
    total_cost INT NOT NULL,
	user_id UUID REFERENCES users(id),
	created_at timestamptz DEFAULT NOW(),
    updated_at timestamptz DEFAULT NOW()
);

CREATE TABLE order_items (
    order_id UUID REFERENCES orders(id),
    product_id UUID NOT NULL,
    price INT NOT NULL,
    description TEXT NOT NULL,
    quantity BIGINT NOT NULL,
	created_at timestamptz DEFAULT NOW(),
	updated_at timestamptz DEFAULT NOW(),
    PRIMARY KEY (order_id, product_id)
);