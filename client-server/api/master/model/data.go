package model

type Data struct {
	Author  string    `json:"author"`
	Title   string    `json:"title"`
	Comment []Message `json:"comments"`
}

type Message struct {
	Message string `json:"message"`
}
