package controllers

import (
	"database/sql"
	"ff/database/models"
	"log"
)

func GetNewsByUuid(uuid string) (models.News, error) {
	SELECT_NEWS_BY_UUID := "SELECT UUID, title, content, author FROM news WHERE UUID = ?"

	var news models.News
	err := db.QueryRow(SELECT_NEWS_BY_UUID, uuid).Scan(&news.UUID, &news.Title, &news.Content, &news.Author)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No news found with UUID: %s", uuid)
			return models.News{}, nil
		}
		log.Printf("Error querying database: %v", err)
		return models.News{}, err
	}

	return news, nil
}
