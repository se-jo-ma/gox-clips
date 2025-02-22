package auth

// import (
// 	"net/http"

// 	"github.com/gorilla/sessions"
// )

// var store = sessions.NewCookieStore([]byte("your-secret-key"))

// func loginHandler(w http.ResponseWriter, r *http.Request) {
// 	session, _ := store.Get(r, "session-name")
// 	session.Values["authenticated"] = true
// 	session.Save(r, w)
// 	http.Redirect(w, r, "/", http.StatusSeeOther)
// }

// func protectedHandler(w http.ResponseWriter, r *http.Request) {
// 	session, _ := store.Get(r, "session-name")
// 	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
// 		http.Error(w, "Forbidden", http.StatusForbidden)
// 		return
// 	}
// 	w.Write([]byte("Protected content"))
// }

// func main() {
// 	http.HandleFunc("/login", loginHandler)
// 	http.HandleFunc("/protected", protectedHandler)
// 	http.ListenAndServe(":8080", nil)
// }
