-- DROP DATABASE desafio_tecnico_backend;
CREATE DATABASE IF NOT EXISTS desafio_tecnico_backend;
USE desafio_tecnico_backend;

CREATE TABLE empresa (
    empresa_id INT PRIMARY KEY AUTO_INCREMENT,
    cnpj VARCHAR(14) UNIQUE,
    senha VARCHAR(255)
);

CREATE TABLE destinatario (
    destinatario_id INT PRIMARY KEY AUTO_INCREMENT,
    cnpj VARCHAR(14) UNIQUE,
    xNome VARCHAR(255),
    email VARCHAR(255),
    enderDest JSON
);

CREATE TABLE produto (
    produto_id INT PRIMARY KEY AUTO_INCREMENT,
    cProd VARCHAR(20) UNIQUE,
    cEAN VARCHAR(14) UNIQUE,
    xProd VARCHAR(255),
    uCom VARCHAR(10),
    qCom DECIMAL(15, 2),
    vUnCom DECIMAL(15, 2),
    vCusto DECIMAL(15, 2),
    vPreco DECIMAL(15, 2)
);

CREATE TABLE endereco (
    id INT AUTO_INCREMENT PRIMARY KEY,
    destinatario_id INT, 
    xLgr VARCHAR(255) NOT NULL,
    nro VARCHAR(20),
    xCpl VARCHAR(255),
    xBairro VARCHAR(255) NOT NULL,
    cMun INT NOT NULL,
    CEP VARCHAR(8) NOT NULL,
    fone VARCHAR(20),
    FOREIGN KEY (destinatario_id) REFERENCES destinatario(destinatario_id) 
);

INSERT INTO empresa VALUES (null, 10541434000152, "$2a$10$1TXJzkmloaQQ9npKzSPAp.AL1hP0SP85pp6tnMPdnm9W8w2aFr8xa");

