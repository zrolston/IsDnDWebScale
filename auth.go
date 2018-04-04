package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

type UserAuth struct {
	AuthToken string    `datastore:"authToken" json:"auth"`
	Username  string    `datastore:"user" json:"user"`
	Password  []byte    `datastore:"pass" json:"pass"`
	LastUsed  time.Time `datastore:"time" json:"time"`
}

//Should error out when username already exists in the database
func signUp(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	username, password, ok := r.BasicAuth()

	if !ok {
		w.WriteHeader(418)
		writeMessage(w, "Basic Auth incorrectly passed")
		return
	}

	ctx := appengine.NewContext(r)

	key := datastore.NewKey(ctx, "userAuth", username, 0, nil)

	var user UserAuth

	//TEST IF USERNAME IS TAKEN
	if err := datastore.Get(ctx, key, &user); err == nil {
		w.WriteHeader(418)
		writeMessage(w, "Username already taken")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		w.WriteHeader(500)
		writeMessage(w, "Could not encrypt for some reason")
		return
	}

	user = UserAuth{
		AuthToken: uuid.New().String(),
		Username:  username,
		Password:  hash,
		LastUsed:  time.Now(),
	}

	//Write username and hashed password to database
	_, err = datastore.Put(ctx, key, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		writeMessage(w, "Could not put in DB")
		return
	}

	w.WriteHeader(200)
	writeMessage(w, "User registered successfully with username "+string(user.Username))
}

func login(w http.ResponseWriter, r *http.Request) {

	setupResponse(&w, r)

	username, pass, _ := r.BasicAuth()

	ctx := appengine.NewContext(r)
	key := datastore.NewKey(ctx, "userAuth", username, 0, nil)

	var uAuth UserAuth

	//TODO get the users pwd from the database
	if err := datastore.Get(ctx, key, &uAuth); err != nil {
		w.WriteHeader(418)
		writeMessage(w, "Username not registered")
		return
	}

	realPass := uAuth.Password

	if err := bcrypt.CompareHashAndPassword(realPass, []byte(pass)); err != nil {
		w.WriteHeader(400)
		writeMessage(w, err.Error())
		return
	}

	authToken := uuid.New().String()

	uAuth.AuthToken = authToken
	uAuth.LastUsed = time.Now()

	if _, err := datastore.Put(ctx, key, &uAuth); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeMessage(w, "Could not put in DB")
		return
	}

	w.WriteHeader(200)
	writeMessage(w, authToken)
}

type TestDB struct {
	Num int `datastore:"fug" json:"fug"`
}

func dbTestPut(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	myNum, err := strconv.Atoi(vars["num"])

	setupResponse(&w, r)

	if err != nil {
		writeMessage(w, "Not a number")
		return
	}

	dbt := TestDB{
		Num: myNum,
	}

	ctx := appengine.NewContext(r)

	key := datastore.NewKey(ctx, "dbTest", "maymay", 0, nil)

	_, err = datastore.Put(ctx, key, &dbt)
	if err != nil {
		w.WriteHeader(500)
		writeMessage(w, "Could not write to DB")
		return
	}

	key = datastore.NewKey(ctx, "dbTest", "maymay", 0, nil)

	var newt TestDB

	if err := datastore.Get(ctx, key, &newt); err != nil {
		w.WriteHeader(418)
		writeMessage(w, "Not Found in DB")
		return
	}

	w.WriteHeader(200)
	writeMessage(w, "Successfully added "+strconv.Itoa(newt.Num)+", should be "+strconv.Itoa(dbt.Num))
}

func dbTestGet(w http.ResponseWriter, r *http.Request) {

	setupResponse(&w, r)
	ctx := appengine.NewContext(r)

	key := datastore.NewKey(ctx, "dbTest", "maymay", 0, nil)

	dbt := TestDB{
		Num: 69,
	}

	if err := datastore.Get(ctx, key, &dbt); err != nil {
		w.WriteHeader(418)
		writeMessage(w, "Not found in DB")
		return
	}

	write := "The number was " + strconv.Itoa(dbt.Num)

	w.WriteHeader(200)
	writeMessage(w, write)
}

func basicAuthTest(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	username, password, _ := r.BasicAuth()

	s := "Successful auth with Username " + username + " and password " + password
	w.WriteHeader(200)
	writeMessage(w, s)
}

func testTokens(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	token, _, _ := r.BasicAuth()
	ctx := appengine.NewContext(r)

	resp, ok := checkValidToken(token, ctx)

	if ok {
		w.WriteHeader(200)
	}
	if !ok {
		w.WriteHeader(404)
	}
	writeMessage(w, resp)
}
