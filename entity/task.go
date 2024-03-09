package entity

type Task struct {
	ID              int    `json:"ID"`
	UserID          int    `json:"UserID"`
	Titulo          string `json:"Titulo"`
	Descricao       string `json:"Descricao"`
	DataDeCriacao   string `json:"DataDeCriacao"`
	DataDeConclusao string `json:"DataDeConclusao"`
	Status          string `json:"Status"`
}
