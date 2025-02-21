package services

import (
	"htmx-go-app/internal/auth"
	"net/http"
)

type AuthService struct {
	auth.Auth
}

func (s *AuthService) Login(w http.ResponseWriter, r *http.Request) {
	s.LoginHandler(w, r)
}

func (s *AuthService) Protected(w http.ResponseWriter, r *http.Request) {
	s.ProtectedHandler(w, r)
}

func (s *AuthService) Logout(w http.ResponseWriter, r *http.Request) {
	s.LogoutHandler(w, r)
}

