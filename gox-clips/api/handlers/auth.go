package handlers

import (
	"net/http"
	"github.com/se-jo-ma/gox-clips/gox-clips/internal/services"
)
type AuthHandler struct {
    Service services.AuthService
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
    h.Service.Login(w, r)
}

func (h *AuthHandler) Protected(w http.ResponseWriter, r *http.Request) {
    h.Service.Protected(w, r)
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
    h.Service.Logout(w, r)
}


