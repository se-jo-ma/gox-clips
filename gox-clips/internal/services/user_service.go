package services

import (
	"htmx-go-app/pkg/models"
)

type UserService struct{}

func (s *UserService) GetUser(id int) models.User {
    return models.User{ID: uint(id), Name: "Alice", Email: "alice@example.com"}
}