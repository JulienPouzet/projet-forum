package models

type News struct {
	UUID     string `json:"uuid"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Author   string `json:"author"`
	Category string `json:"category"`
	Date     string `json:"date"`
}
