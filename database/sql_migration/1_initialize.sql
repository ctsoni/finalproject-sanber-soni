-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE users (
	id SERIAL PRIMARY KEY,
	full_name VARCHAR NOT NULL,
	email VARCHAR NOT NULL,
	password_hash VARCHAR NOT NULL,
    is_admin BOOLEAN DEFAULT FALSE,
	created_at TIMESTAMP NOT NULL DEFAULT (NOW()),
	updated_at TIMESTAMP NOT NULL DEFAULT (NOW())
);

CREATE TABLE inventory_categories (
	id SERIAL PRIMARY KEY,
	name VARCHAR NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT (NOW()),
	updated_at TIMESTAMP NOT NULL DEFAULT (NOW())
);

CREATE TABLE inventories (
	id SERIAL PRIMARY KEY,
	cat_id int,
	name VARCHAR NOT NULL,
	description VARCHAR NOT NULL,
	is_available BOOLEAN DEFAULT TRUE,
	created_at TIMESTAMP NOT NULL DEFAULT (NOW()),
	updated_at TIMESTAMP NOT NULL DEFAULT (NOW()),
	CONSTRAINT fk_cat_id 
	FOREIGN KEY (cat_id) REFERENCES inventory_categories(id) ON DELETE CASCADE
);

CREATE TABLE inventory_stocks (
	id SERIAL PRIMARY KEY,
	inven_id int,
	stock_unit INT NOT NULL,
	price_per_unit INT NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT (NOW()),
	updated_at TIMESTAMP NOT NULL DEFAULT (NOW()),
	CONSTRAINT fk_inven_id 
	FOREIGN KEY (inven_id) REFERENCES inventories(id) ON DELETE CASCADE
);

DROP TYPE IF EXISTS status;
CREATE TYPE status AS ENUM ('Failed', 'Unpaid', 'Paid', 'Canceled');

CREATE TABLE transactions (
	id SERIAL PRIMARY KEY,
	user_id int,
	inven_id int,
	unit int NOT NULL,
	total_price int NOT NULL,
	status status NOT NULL,
	start_at TIMESTAMP NOT NULL DEFAULT (NOW()),
	finish_at TIMESTAMP NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT (NOW()),
	updated_at TIMESTAMP NOT NULL DEFAULT (NOW()),
	CONSTRAINT fk_user_id 
	FOREIGN KEY (user_id) REFERENCES users(id),
	CONSTRAINT fk_inven_id 
	FOREIGN KEY (inven_id) REFERENCES inventories(id)
);

CREATE TABLE reviews (
	id SERIAL PRIMARY KEY,
	user_id INT,
	trans_id INT,
	review TEXT,
	rating INT DEFAULT 5,
	created_at TIMESTAMP NOT NULL DEFAULT (NOW()),
	updated_at TIMESTAMP NOT NULL DEFAULT (NOW()),
	CONSTRAINT fk_user_id 
	FOREIGN KEY (user_id) REFERENCES users(id),
	CONSTRAINT fk_trans_id 
	FOREIGN KEY (trans_id) REFERENCES transactions(id)
)

-- +migrate StatementEnd
