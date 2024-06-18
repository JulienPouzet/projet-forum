// models/post.go
package models

type Post struct {
	UUID     string    `json:"uuid"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Author   string    `json:"author"`
	Category string    `json:"category"`
	Date     string    `json:"date"`
	Comments []Comment `json:"comments"`
}

type Posts []Post

type Comment struct {
	UUID     string `json:"uuid"`
	PostUUID string `json:"post_uuid"`
	Author   string `json:"author"`
	Content  string `json:"content"`
	Date     string `json:"date"`
}
