package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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

	http.HandleFunc("/", welcomeTest)
	http.HandleFunc("/api/getGerald", getGerald)
	http.HandleFunc("/api/getGerald/weapon", getGeraldWep)
	http.HandleFunc("/api/putChar/skill", addSkill)

	appengine.Main()
}
