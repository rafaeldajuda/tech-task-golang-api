package entity

type User struct {
	ID    int64  `json:"ID,omitempty"`
	Nome  string `json:"Nome,omitempty"`
	Email string `json:"Email,omitempty"`
	Senha string `json:"Senha,omitempty"`
}
