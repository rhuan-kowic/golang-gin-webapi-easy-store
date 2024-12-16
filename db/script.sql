CREATE TABLE
    products (
        id_product SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        description TEXT,
        price NUMERIC(10, 2) NOT NULL,
        stock_quantity INT NOT NULL DEFAULT 0,
        minimum_stock INT NOT NULL DEFAULT 0
    );

CREATE TABLE
    sales (
        id_sale SERIAL PRIMARY KEY,
        product_id INT REFERENCES products (id_product) ON DELETE CASCADE,
        date_sale TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        total NUMERIC(10, 2) NOT NULL,
        quantity INT NOT NULL,
        CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES products (id_product) ON DELETE CASCADE
    );

INSERT INTO
    products (
        name,
        description,
        price,
        stock_quantity,
        minimum_stock
    )
VALUES
    (
        'Camisa Polo',
        'Camisa polo de algodão',
        59.90,
        100,
        10
    ),
    (
        'Calça Jeans',
        'Calça jeans masculina',
        129.90,
        50,
        5
    ),
    (
        'Tênis Esportivo',
        'Tênis esportivo confortável',
        199.90,
        80,
        10
    ),
    (
        'Jaqueta Couro',
        'Jaqueta de couro estilo motociclista',
        249.90,
        30,
        5
    ),
    (
        'Camiseta Básica',
        'Camiseta básica de algodão',
        39.90,
        200,
        20
    );

INSERT INTO
    sales (product_id, total, quantity)
VALUES
    (1, 119.70, 2),
    (2, 259.80, 2),
    (3, 399.80, 2),
    (4, 499.80, 2),
    (5, 79.80, 2);