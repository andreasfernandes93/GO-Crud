-- database: /home/andreas/Documentos/Workspace/PROJETOS/Go/CRUD/github.com/andreasfernandes93/GO-Crud/src/infra/db/db.db

-- Pressione o botão ▷ no canto superior direito da janela para executar todo o arquivo.

DROP TABLE IF EXISTS Products;

SELECT * FROM Products;

CREATE TABLE Products (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    Name VARCHAR,
    Description VARCHAR,
    Price REAL,
    Quantity INTEGER
);

INSERT INTO Products (Name, Description, Price, Quantity) VALUES
("Mouse Gamer Logitec","180dpi sem fio", 150.0, 500),
("Placa de Video TOP","RTX 80GB", 1500.0, 500),
('Notebook', 'CoreI3 | 8GB RAM | SSD 1TB', 2000.00, 200),
('Smartphone', '6.5'' Display | 128GB Storage | 5000mAh Battery', 1200.00, 150),
('TV', '55'' 4K Ultra HD | Smart TV | HDR', 2500.00, 100),
('Headphones', 'Wireless | Noise Cancelling | Bluetooth', 150.00, 300),
('PS5', 'Videogame da Sony', 5500.99, 300)



