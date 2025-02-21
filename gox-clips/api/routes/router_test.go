package routes

import (
	"net/http"
)

func Router_test() *http.ServeMux {
	mux := http.NewServeMux()

	return mux
}
