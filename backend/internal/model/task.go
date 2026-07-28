package model

import (
	"time"
)

type Task struct {
    ID          string 	  `json:"id"`
    Nome        string    `json:"nome"`
    Descricao   string    `json:"descricao"`
    DataCriacao time.Time `json:"data_criacao"`
    Importancia string    `json:"importancia"`
    Coluna      string    `json:"coluna"`
}