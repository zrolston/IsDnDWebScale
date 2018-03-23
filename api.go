package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"crypto/md5"
  "encoding/hex"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/appengine"
)

type apiResponse struct {
	errorCode int           `json:"err"`
	message   []interface{} `json:"msg"`
}

func newAPIR(err int, data []interface{}) *apiResponse{
	newOne := apiResponse{
		errorCode: err,
		message: data
	}

	return &newOne
}

func (api apiResponse) setMessage(data []interface{}){
	api.message = data
}

//Should error out when username already exists in the database
func signUp(w http.ResponseWriter, r *http.Request) {
	username := "Davin"
	password := "pa$$word"
	//TODO GET USERNAME/PASSWORD FROM httpRequest

	err := nil
	//TODO TEST IF USERNAME IS TAKEN

	if err != nil {
		w.WriteHeader(418)
		w.Write("Username already taken")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		w.WriteHeader(500)
		w.Write("Could not encrypt for some reason")
		log.Fatal(err)
	}

	//TODO Write username and hashed password to database

	w.WriteHeader(200)
	w.Write("User registered successfully")
}

func login(w http.ResponseWriter, r *http.Request) {
	user, pass, err := r.BasicAuth()

	realPass := "maymays"

	//TODO get the users pwd from the database

	if err != nil {
		w.WriteHeader(500)
		w.Write("Could not encrypt for some reason")
		log.Fatal(err)
	}

	if err := bcrypt.CompareHashAndPassword(realPass, []byte(pass)); err != nil {
        w.WriteHeader(400)
				w.Write("Password did not match")
				return
  }

	hash := md5.New()
	hash.write(realPass)
	authToken := hex.EncodeToString(hash.Sum(nil))

	w.WriteHeader(200)
	w.Write(authToken)
}




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

	http.Handle("/", router)
}
