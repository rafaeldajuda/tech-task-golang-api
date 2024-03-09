package entity

type User struct {
	ID    int    `json:"ID,omitempty"`
	Nome  string `json:"Nome"`
	Email string `json:"Email"`
	Senha string `json:"Senha,omitempty"`
}
