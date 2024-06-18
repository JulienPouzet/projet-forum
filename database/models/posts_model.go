// models/post.go

package models

import "html/template"

type Post struct {
	UUID     string        `json:"uuid"`
	Title    string        `json:"title"`
	Content  template.HTML `json:"content"`
	Author   string        `json:"author"`
	Category string        `json:"category"`
	Date     string        `json:"date"`
	Comments []Comment     `json:"comments"`
}

type Posts []Post

type Comment struct {
	UUID     string        `json:"uuid"`
	PostUUID string        `json:"post_uuid"`
	Author   string        `json:"author"`
	Content  template.HTML `json:"content"`
	Date     string        `json:"date"`
}
