package services

import (
	"github.com/se-jo-ma/gox-clips/gox-clips/pkg/models"
)

type UserService struct{}

func (s *UserService) GetUser(id int) models.User {
    return models.User{ID: uint(id), Name: "Alice", Email: "alice@example.com"}
}