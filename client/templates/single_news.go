package templates

import (
	"ff/api/controllers"
	"ff/database/models"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type SingleNewsTemplate struct {
	News models.News
}

func SingleNews(w http.ResponseWriter, r *http.Request) {
	uuid := strings.TrimPrefix(r.URL.Path, "/news/")
	if uuid == "" {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	news, err := controllers.GetNewsByUuid(uuid)

	if err != nil {
		http.Error(w, "Failed to fetch news", http.StatusInternalServerError)
		return
	}

	data := SingleNewsTemplate{
		News: news,
	}

	tmpl, err := template.ParseFiles("web/pages/news/index.html")
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		log.Printf("Template parsing error: %v", err)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
		log.Printf("Template execution error: %v", err)
	}

}
