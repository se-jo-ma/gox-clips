package auth

import (
	"net/http"
	"github.com/gorilla/sessions"
)

type Auth struct {
	loginHandler func(w http.ResponseWriter, r *http.Request)
	protectedHandler func(w http.ResponseWriter, r *http.Request)
}
var store = sessions.NewCookieStore([]byte("your-secret-key"))

func (a *Auth) LoginHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	session.Values["authenticated"] = true
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (a *Auth) ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	w.Write([]byte("Protected content"))
}

func (a *Auth) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	session.Values["authenticated"] = false
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

