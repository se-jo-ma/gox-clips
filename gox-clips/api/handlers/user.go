package handlers

import (
    "encoding/json"
    "github.com/se-jo-ma/gox-clips/gox-clips/internal/services"
    "net/http"
)

type UserHandler struct {
    Service services.UserService
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
    user := h.Service.GetUser(1)
    json.NewEncoder(w).Encode(user)
}
