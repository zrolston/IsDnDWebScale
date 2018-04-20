package main

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"google.golang.org/appengine"
)

func main() {
	RegisterHandlers()

	appengine.Main()
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Testing")
}

func RegisterHandlers() {
	router := mux.NewRouter()
	router.HandleFunc("/", welcomeTest)
	router.HandleFunc("/api/getGerald", getGeraldDB)
	router.HandleFunc("/api/putGerald", putGerald)
	router.HandleFunc("/api/login", login)
	router.HandleFunc("/api/signup", signUp)
	router.HandleFunc("/api/you/should/not/call/this", removeUser)
	router.HandleFunc("/api/putChar", putChar)
	router.HandleFunc("/api/getChars", getCharList)
	router.HandleFunc("/api/getRaces", getRaces)
	router.HandleFunc("/api/getClasses", getClasses)
	router.HandleFunc("/api/deleteChars", deleteChars)
	router.HandleFunc("/api/deleteChar", deleteChar)
	router.HandleFunc("/api/emptyChar", emptyChar)
	router.HandleFunc("/api/checkToken", validToken)
	router.HandleFunc("/api/getSkills", getSkills)

	originsOK := handlers.AllowedOrigins([]string{"*"})
	headersOk := handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "X-Testing"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	exposed := handlers.ExposedHeaders([]string{"Authorization"})
	http.Handle("/", corsHandler(handlers.CORS(originsOK, headersOk, methodsOk, exposed)(router)))
}

func corsHandler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			setupResponse(&w, r)
			w.WriteHeader(200)
		} else {
			h.ServeHTTP(w, r)
		}
	}
}
