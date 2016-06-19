package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ngurajeka/ng"
)

func StdLog(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	resp := ng.ErrorResponse(
		ng.MapString{"message": http.StatusText(http.StatusNotFound)},
	)
	json.NewEncoder(w).Encode(resp)
}
