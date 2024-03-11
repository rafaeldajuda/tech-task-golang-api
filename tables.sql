CREATE TABLE IF NOT EXISTS usuario (
    ID INTEGER PRIMARY KEY AUTO_INCREMENT,
    Nome VARCHAR(100),
    Email VARCHAR(100),
    Senha VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS tarefa_status (
     ID INTEGER PRIMARY KEY AUTO_INCREMENT,
     Descricao  VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS tarefa (ID INTEGER PRIMARY KEY AUTO_INCREMENT,UserID INTEGER,Titulo VARCHAR(100),Descricao  VARCHAR(255),DataDeCriacao DATETIME DEFAULT CURRENT_TIMESTAMP,DataDeConclusao DATETIME,`Status` INTEGER,FOREIGN KEY (UserID) REFERENCES usuario(ID),FOREIGN KEY (`Status`) REFERENCES tarefa_status(ID));

INSERT INTO usuario (Nome, Email, Senha) VALUES ("Rafael", "rafael@gmail.com", "aaa123");
INSERT INTO usuario (Nome, Email, Senha) VALUES ("?", "?", "?");

INSERT INTO tarefa_status (Descricao) SELECT 'pendente' WHERE NOT EXISTS (SELECT 1 FROM tarefa_status WHERE Descricao = 'pendente'); 
INSERT INTO tarefa_status (Descricao) SELECT 'em andamento' WHERE NOT EXISTS (SELECT 1 FROM tarefa_status WHERE Descricao = 'em andamento'); 
INSERT INTO tarefa_status (Descricao) SELECT 'concluida' WHERE NOT EXISTS (SELECT 1 FROM tarefa_status WHERE Descricao = 'concluida'); 