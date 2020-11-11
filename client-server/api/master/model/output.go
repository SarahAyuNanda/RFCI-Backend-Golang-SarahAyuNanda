package model

type Author struct {
	Author string `json:"author"`
}
type Title struct {
	Title string `json:"title"`
}
type Comment struct {
	Comment []Message `json:"comment"`
}