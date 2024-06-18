package templates

import (
	"ff/api/controllers"
	"ff/database/models"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type PostsTemplate struct {
	CurrentUser models.User
	LoggedIn    bool
	Posts       []models.Post
}

func Posts(w http.ResponseWriter, r *http.Request) {
	currentUser, err := controllers.GetCurrentLoggedInUser(r)

	if err != nil {
		if err != http.ErrNoCookie {
			http.Error(w, "Failed to fetch user", http.StatusInternalServerError)
			log.Printf("Error fetching user: %v", err)
			return
		}
	}

	posts, err := controllers.GetAllPosts()
	if err != nil {
		http.Error(w, "Failed to fetch posts", http.StatusInternalServerError)
		log.Printf("Error fetching posts: %v", err)
		return
	}

	// data := PostsTemplate{
	// 	CurrentUser: currentUser,
	// 	LoggedIn:    currentUser.UUID != "",
	// 	Posts:       posts,
	// }

	tmpl, err := template.ParseFiles("web/pages/posts.html")
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		log.Printf("Template parsing error: %v", err)
		return
	}

	var formattedPosts []models.Post
	for _, p := range posts {
		formattedContent := template.HTML(strings.ReplaceAll(string(p.Content), "\n", "<br>"))
		p.Content = formattedContent
		formattedPosts = append(formattedPosts, p)
	}

	data := PostsTemplate{
		CurrentUser: currentUser,
		LoggedIn:    currentUser.UUID != "",
		Posts:       formattedPosts,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
		log.Printf("Template execution error: %v", err)
	}
}
