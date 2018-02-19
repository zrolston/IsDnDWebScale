package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getChar(w http.ResponseWriter, r *http.Request) {
	char := makeGerald()
	response, _ := json.Marshal(char)
	w.Write(response)
}

func getWep(w http.ResponseWriter, r *http.Request) {
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

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/getChar/", getChar).Methods("GET")
	router.HandleFunc("/getChar/weapon", getWep).Methods("GET")
	router.HandleFunc("/putChar/skil", addSkill).Methods("POST")

}
