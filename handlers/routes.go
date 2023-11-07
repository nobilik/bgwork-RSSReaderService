package handlers

import (
	"net/http"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/parse", http.HandlerFunc(ParseHandler))
	return mux
}
