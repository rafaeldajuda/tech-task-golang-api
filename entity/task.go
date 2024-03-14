package entity

type Task struct {
	ID              int    `json:"ID"`
	UserID          int    `json:"UserID,omitempty"`
	Titulo          string `json:"Titulo"`
	Descricao       string `json:"Descricao"`
	DataDeCriacao   string `json:"DataDeCriacao"`
	DataDeConclusao string `json:"DataDeConclusao"`
	Status          string `json:"Status"`
}

type PostTaskSuccess struct {
	IDTask int64 `json:"id_task"`
}
