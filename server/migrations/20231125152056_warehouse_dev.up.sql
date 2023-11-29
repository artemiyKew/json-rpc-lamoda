CREATE TABLE Warehouse (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    availability BOOLEAN NOT NULL 
);

CREATE TABLE Product (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    size VARCHAR(50),
    unique_code VARCHAR(100) UNIQUE NOT NULL
);

CREATE TABLE Shipping (
    id SERIAL PRIMARY KEY,
    unique_code VARCHAR(100) REFERENCES Product(unique_code) ON DELETE CASCADE,
    warehouse_id INTEGER REFERENCES Warehouse(id) ON DELETE CASCADE,
    quantity INTEGER DEFAULT 0
);

-- Сделать таблицу Reseravtion 
CREATE TABLE Reservation (
    id SERIAL PRIMARY KEY,
    unique_code VARCHAR(100) REFERENCES Product(unique_code) ON DELETE CASCADE,
    warehouse_id INTEGER REFERENCES Warehouse(id) ON DELETE CASCADE,
    quantity INTEGER,
    status VARCHAR(255) NOT NULL
);

-- Вставка данных в таблицу Warehouse
INSERT INTO Warehouse (name, availability) VALUES
    ('Warehouse1', true),
    ('Warehouse2', true);

-- Вставка данных в таблицу Product
INSERT INTO Product (name, size, unique_code) VALUES
    ('Product1', 'Size1', 'Code1'),
    ('Product2', 'Size2', 'Code2'),
    ('Product3', 'Size3', 'Code3'),
    ('Product4', 'Size4', 'Code4'),
    ('Product5', 'Size5', 'Code5'),
    ('Product6', 'Size6', 'Code6'),
    ('Product7', 'Size7', 'Code7'),
    ('Product8', 'Size8', 'Code8'),
    ('Product9', 'Size9', 'Code9'),
    ('Product10', 'Size10', 'Code10');

-- Вставка данных в таблицу Shipping
INSERT INTO Shipping (unique_code, warehouse_id, quantity) VALUES
    ('Code1', 1, 100),
    ('Code2', 1, 150),
    ('Code1', 2, 200),
    ('Code3', 2, 50),
    ('Code3', 1, 80),
    ('Code5', 2, 120),
    ('Code4', 1, 90),
    ('Code10', 2, 60),
    ('Code10', 1, 110),
    ('Code9', 2, 70);