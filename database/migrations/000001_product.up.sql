  CREATE TABLE products (
    product_id BIGSERIAL PRIMARY KEY,
    aggregate_identifier UUID NOT NULL,
    code INTEGER NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    stock INTEGER NOT NULL,
    total_stock INTEGER NOT NULL,
    cut_stock INTEGER NOT NULL,
    available_stock INTEGER NOT NULL,
    price_from FLOAT NOT NULL,
    price_to FLOAT NOT NULL,
    created_at TIMESTAMP(3) NOT NULL,
    updated_at TIMESTAMP(3) NOT NULL,
    created_by SMALLINT NOT NULL, 
    updated_by SMALLINT NOT NULL
  );