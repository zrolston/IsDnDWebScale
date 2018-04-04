package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

func putGerald(w http.ResponseWriter, r *http.Request) {
	char := makeGerald()
	setupResponse(&w, r)

	ctx := appengine.NewContext(r)
	gKey := datastore.NewKey(ctx, "character", "Gerald", 0, nil)

	_, err := datastore.Put(ctx, gKey, &char)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, (err))
		return
	}

	w.WriteHeader(200)
	fmt.Fprintln(w, "Correctly Stored Gerald")
}

func getGeraldDB(w http.ResponseWriter, r *http.Request) {

	setupResponse(&w, r)

	var char Character

	ctx := appengine.NewContext(r)
	gKey := datastore.NewKey(ctx, "character", "Gerald", 0, nil)

	if err := datastore.Get(ctx, gKey, &char); err != nil {
		w.WriteHeader(418)
		fmt.Fprintln(w, "Not found")
		return
	}

	response, _ := json.Marshal(char)
	w.WriteHeader(200)
	w.Write(response)
}

func getGeraldWep(w http.ResponseWriter, r *http.Request) {
	char := makeGerald()
	wep := char.EquipedWeapon[0]
	response, _ := json.Marshal(wep)
	setupResponse(&w, r)
	w.WriteHeader(200)
	w.Write(response)
}

func addSkill(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var skill Skill
	err := decoder.Decode(&skill)
	setupResponse(&w, r)
	if err == nil {
		w.WriteHeader(500)
	}
	log.Print(skill.Name)
	w.WriteHeader(200)
}

func welcomeTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the fucking show")
}

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
	router.HandleFunc("/", welcomeTest).Methods("GET")
	router.HandleFunc("/api/getGerald", getGeraldDB).Methods("GET")
	router.HandleFunc("/api/putGerald", putGerald)
	router.HandleFunc("/api/getGerald/weapon", getGeraldWep).Methods("GET")
	router.HandleFunc("/api/putChar/skill", addSkill).Methods("POST")
	router.HandleFunc("/api/login", login)
	router.HandleFunc("/api/signup", signUp)
	router.HandleFunc("/api/testDB/put/{num}", dbTestPut)
	router.HandleFunc("/api/testDB/get", dbTestGet)
	router.HandleFunc("/api/testAuth", testTokens)
	router.HandleFunc("/api/you/should/not/call/this", removeUser)
	router.HandleFunc("/api/putChar", putChar).Methods("POST")
	router.HandleFunc("/api/testBody", testBody).Methods("POST")
	router.HandleFunc("/api/getChars", getCharList)

	http.Handle("/", router)
}
