CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    budget DECIMAL(10, 2) NOT NULL CHECK (budget >= 0),
    expense DECIMAL(10, 2) NOT NULL DEFAULT 0 CHECK (expense >= 0)
);

INSERT INTO categories (name, budget) VALUES
    ('Groceries', 500),
    ('Rent', 1000),
    ('Transportation', 100),
    ('Entertainment', 50);
