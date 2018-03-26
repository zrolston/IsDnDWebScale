package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

type UserAuth struct {
	authToken string    `datastore:"authToken" json:"auth"`
	username  string    `datastore:"user" json:"user"`
	password  string    `datastore:"pass" json:"pass"`
	lastUsed  time.Time `datastore:"time" json:"time"`
}

//Should error out when username already exists in the database
func signUp(w http.ResponseWriter, r *http.Request) {
	username := "Davin"
	password := "pa$$word"

	//TODO GET USERNAME/PASSWORD FROM httpRequest

	ctx := appengine.NewContext(r)

	key := datastore.newKey(ctx, "userAuth", username, 0, nil)

	var user userAuth

//TEST IF USERNAME IS TAKEN
	if err := datastore.Get(ctx, key, &user); err == nil{
		w.WriteHeader(418)
		w.Write("Username already taken")
		return
	}

	err := nil

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		w.WriteHeader(500)
		w.Write("Could not encrypt for some reason")
		log.Fatal(err)
	}

	user := userAuth{
		authToken: uuid.NewV4(),
		username : username,
		password : hash,
		lastUsed : time.Now()
	}

	//Write username and hashed password to database
	newKey, err := datastore.Put(ctx, key, &user)
	if err != nil {
							http.Error(w, err.Error(), http.StatusInternalServerError)
							return
	}

	w.WriteHeader(200)
	w.Write("User registered successfully")
}

func login(w http.ResponseWriter, r *http.Request) {
	user, pass, err := r.BasicAuth()

	realPass := "maymays"

	ctx := appengine.NewContext(r)
	key := datastore.newKey(ctx, "userAuth", username, 0, nil)

	var uAuth userAuth

//TODO get the users pwd from the database
	if err := datastore.Get(ctx, key, &uAuth); err != nil{
		w.WriteHeader(418)
		w.Write("Username not registered")
		return
	}

	realPass := uAuth.password

	if err := bcrypt.CompareHashAndPassword(realPass, []byte(pass)); err != nil {
		w.WriteHeader(400)
		w.Write("Password did not match")
		return
	}

	authToken, err := uuid.NewV4()

	if err != nil {
		log.Printf("Something went wrong: %s", err)
		w.WriteHeader(500)
		w.Write("Could not generate token, sorry")
		return
	}

	uAuth.authToken = authToken
	uAuth.lastUsed = time.Now()

	newKey, err := datastore.Put(ctx, key, &uAuth)
	if err != nil {
							http.Error(w, err.Error(), http.StatusInternalServerError)
							return
	}

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
