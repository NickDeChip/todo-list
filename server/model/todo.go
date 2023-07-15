package model

type Todo struct {
	ID   int64  `json:"id"`
	Info string `json:"info"`
}

type Info struct {
	Info string `json:"info"`
}
