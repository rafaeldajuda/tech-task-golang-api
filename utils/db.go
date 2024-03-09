package utils

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2/log"
)

func Connection() *sql.DB {
	db, err := sql.Open("mysql", "root:admin@tcp(172.17.0.2:3306)/db_techtask")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(-1)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	log.Debug("database connection ok")
	return db
}

func ConfigTables(db *sql.DB) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	createTableUser(db, ctx)
	createTableTaskStatus(db, ctx)
	createTableTask(db, ctx)
}

func createTableUser(db *sql.DB, ctx context.Context) {
	query := `CREATE TABLE IF NOT EXISTS usuario (
		ID INTEGER PRIMARY KEY AUTO_INCREMENT,
		Nome VARCHAR(100),
		Email VARCHAR(100),
		Senha VARCHAR(100)
	);`
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}
	log.Debug("table usuario ok")
}

func createTableTaskStatus(db *sql.DB, ctx context.Context) {
	query := `CREATE TABLE IF NOT EXISTS tarefa_status (
		ID INTEGER PRIMARY KEY AUTO_INCREMENT,
		Descricao  VARCHAR(255)
   	);`
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}
	log.Debug("table tarefa_status ok")
}

func createTableTask(db *sql.DB, ctx context.Context) {
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS tarefa (ID INTEGER PRIMARY KEY AUTO_INCREMENT,UserID INTEGER,Titulo VARCHAR(100),Descricao  VARCHAR(255),DataDeCriacao DATETIME DEFAULT CURRENT_TIMESTAMP,DataDeConclusao DATETIME,`%s` INTEGER,FOREIGN KEY (UserID) REFERENCES usuario(ID),FOREIGN KEY (`Status`) REFERENCES tarefa_status(ID));", "Status")
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}
	log.Debug("table tarefa ok")
}

func GetUser(email string, senha string, db *sql.DB) (exist bool, err error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	query := "SELECT ID FROM usuario WHERE Email=? AND Senha=?"
	result, err := db.QueryContext(ctx, query, email, senha)
	if err != nil {
		return
	}
	defer result.Close()
	exist = result.Next()

	return
}
