package utils

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2/log"
	"github.com/rafaeldajuda/tech-task-golang-api/entity"
)

var mapStatus = map[string]int{
	"pendente":     1,
	"em andamento": 2,
	"concluida":    4,
}

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

func GetUser(email string, senha string, db *sql.DB) (id int64, exist bool, err error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	query := "SELECT ID FROM usuario WHERE Email=? AND Senha=?"
	result, err := db.QueryContext(ctx, query, email, senha)
	if err != nil {
		return
	}
	defer result.Close()
	for result.Next() {
		exist = true
		err = result.Scan(&id)
		if err != nil {
			return
		}
	}
	return
}

func InsertUser(nome string, email string, senha string, db *sql.DB) (err error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	query := fmt.Sprintf(`INSERT INTO usuario (Nome, Email, Senha) VALUES ("%s", "%s", "%s");`, nome, email, senha)
	_, err = db.ExecContext(ctx, query)
	if err != nil {
		return
	}

	return
}

func InsertTask(id int64, email string, task entity.Task, db *sql.DB) (idTask int64, err error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	query := "SELECT ID FROM usuario WHERE ID=? AND Email=?"
	result, err := db.QueryContext(ctx, query, id, email)
	if err != nil {
		return
	}
	defer result.Close()
	if result.Next() {
		status := mapStatus["pendente"]
		query := fmt.Sprintf("INSERT INTO tarefa (UserID, Titulo, Descricao, `Status`)"+` VALUES (%d,"%s", "%s", %d)`, id, task.Titulo, task.Descricao, status)
		fmt.Println(query)
		result, errorDb := db.ExecContext(ctx, query)
		if errorDb != nil {
			err = errorDb
			return
		}
		idTask, err = result.LastInsertId()
	}

	return
}

func UpdateTask(idTask int64, idUser int64, email string, task entity.Task, db *sql.DB) (err error) {
	dataConclusao := ""

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	query := "SELECT ID FROM usuario WHERE ID=? AND Email=?"
	result, err := db.QueryContext(ctx, query, idUser, email)
	if err != nil {
		return
	}
	defer result.Close()

	if result.Next() {
		status := mapStatus[task.Status]
		if status == 0 {
			err = errors.New("invalid status")
			return
		} else if status == 4 {
			dataConclusao = `, DataDeConclusao="` + time.Now().Format("2006-01-02 15:04:05") + `"`
		}

		query := fmt.Sprintf(`UPDATE tarefa SET Titulo="%s", Descricao="%s", `+"`Status`=%d"+dataConclusao+` WHERE ID=? AND UserID=?`, task.Titulo, task.Descricao, status)
		log.Debug(query)
		result, errorDb := db.ExecContext(ctx, query, idTask, idUser)
		if errorDb != nil {
			err = errorDb
			return
		}
		_, err = result.LastInsertId()
	}

	return
}

func SelectTasks(idTask int64, id int64, email string, db *sql.DB) (tasks []entity.Task, err error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	query := "SELECT tarefa.ID, tarefa.Titulo, tarefa.Descricao, DATE_FORMAT(tarefa.DataDeCriacao," + `"%d/%m/%Y %k:%i:%s")` + ", tarefa.DataDeConclusao, `tarefa_status`.Descricao FROM tarefa, tarefa_status WHERE tarefa.ID=? AND UserID=? OR UserID=? AND tarefa.`Status`=tarefa_status.ID"
	result, err := db.QueryContext(ctx, query, idTask, id, id)
	if err != nil {
		return
	}
	defer result.Close()
	for result.Next() {
		task := entity.Task{}
		var dataCriacao interface{}
		var dataConclusao interface{}
		err = result.Scan(&task.ID, &task.Titulo, &task.Descricao, &dataCriacao, &dataConclusao, &task.Status)
		if err != nil {
			return
		}
		if dataCriacao != nil {
			task.DataDeCriacao = string(dataCriacao.([]uint8))
		}
		if dataConclusao != nil {
			task.DataDeConclusao = string(dataConclusao.([]uint8))
		}

		tasks = append(tasks, task)
	}
	return
}

func SelectTask(idTask int64, id int64, email string, db *sql.DB) (task entity.Task, err error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	query := "SELECT tarefa.ID, tarefa.Titulo, tarefa.Descricao, DATE_FORMAT(tarefa.DataDeCriacao," + `"%d/%m/%Y %k:%i:%s")` + ", DATE_FORMAT(tarefa.DataDeConclusao," + `"%d/%m/%Y %k:%i:%s")` + ", `tarefa_status`.Descricao FROM tarefa, tarefa_status WHERE tarefa.ID=? AND UserID=? AND tarefa.`Status`=tarefa_status.ID"
	result, err := db.QueryContext(ctx, query, idTask, id)
	if err != nil {
		return
	}
	defer result.Close()
	for result.Next() {

		var dataCriacao interface{}
		var dataConclusao interface{}
		err = result.Scan(&task.ID, &task.Titulo, &task.Descricao, &dataCriacao, &dataConclusao, &task.Status)
		if err != nil {
			return
		}
		if dataCriacao != nil {
			task.DataDeCriacao = string(dataCriacao.([]uint8))
		}
		if dataConclusao != nil {
			task.DataDeConclusao = string(dataConclusao.([]uint8))
		}
	}
	return
}
