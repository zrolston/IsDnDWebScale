package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"google.golang.org/appengine"
)

func getGerald(w http.ResponseWriter, r *http.Request) {
	char := makeGerald()
	response, _ := json.Marshal(char)
	w.Write(response)
}

func getGeraldWep(w http.ResponseWriter, r *http.Request) {
	char := makeGerald()
	wep := char.EquipedWeapon[0]
	response, _ := json.Marshal(wep)
	w.Write(response)
}

func addSkill(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var skill Skill
	err := decoder.Decode(&skill)
	if err == nil {
		panic(err)
	}
	log.Print(skill.Name)
}

func welcomeTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the fucking show")
}

func main() {
	RegisterHandlers()

	appengine.Main()
}

func RegisterHandlers() {
	router := mux.NewRouter()
	router.HandleFunc("/", welcomeTest).Methods("GET")
	router.HandleFunc("/api/getGerald", getGerald).Methods("GET")
	router.HandleFunc("/api/getGerald/weapon", getGeraldWep).Methods("GET")
	router.HandleFunc("/api/putChar/skill", addSkill).Methods("POST")
	//router.HandleFunc("/api/login", login).Methods("POST")
	//router.HandleFunc("/api/signup", signUp).Methods("POST")
	router.HandleFunc("/api/testDB/put/{num}", dbTestPut)
	router.HandleFunc("/api/testDB/get", dbTestGet)

	http.Handle("/", router)
}
