package controllers

import (
	"api/db/models"
	"errors"
)

func AddBookmark(userId uint, title string, link string, keywords []string) (*models.Bookmark, error) {

	if userId == 0 {
		return nil, errors.New("No user id")
	}
	if link == "" {
		return nil, errors.New("No link")
	}
	bookmark := models.Bookmark{
		UserID: userId,
		Title:  title,
		Link:   link,
	}
	return &bookmark, nil
}
