package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

type UserAuth struct {
	authToken string    `datastore:"authToken" json:"auth"`
	username  string    `datastore:"user" json:"user"`
	password  string    `datastore:"pass" json:"pass"`
	lastUsed  time.Time `datastore:"time" json:"time"`
}

/*
//Should error out when username already exists in the database
func signUp(w http.ResponseWriter, r *http.Request) {
	username := "Davin"
	password := "pa$$word"

	//TODO GET USERNAME/PASSWORD FROM httpRequest

	ctx := appengine.NewContext(r)

	key := datastore.newKey(ctx, "userAuth", username, 0, nil)

	var user UserAuth

	//TEST IF USERNAME IS TAKEN
	if err := datastore.Get(ctx, key, &user); err == nil {
		w.WriteHeader(418)
		fmt.Fprintln(w, "Username already taken")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintln(w, "Could not encrypt for some reason")
		log.Fatal(err)
	}

	user = UserAuth{
		authToken: uuid.New().String(),
		username:  username,
		password:  string(hash),
		lastUsed:  time.Now(),
	}

	//Write username and hashed password to database
	newKey, err := datastore.Put(ctx, key, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	fmt.Fprintln(w, "User registered successfully")
}

func login(w http.ResponseWriter, r *http.Request) {
	username, pass, err := r.BasicAuth()

	realPass := "maymays"

	ctx := appengine.NewContext(r)
	key := datastore.newKey(ctx, "userAuth", username, 0, nil)

	var uAuth UserAuth

	//TODO get the users pwd from the database
	if err := datastore.Get(ctx, key, &uAuth); err != nil {
		w.WriteHeader(418)
		fmt.Fprintln(w, "Username not registered")
		return
	}

	realPass = uAuth.password

	if err := bcrypt.CompareHashAndPassword([]byte(realPass), []byte(pass)); err != nil {
		w.WriteHeader(400)
		fmt.Fprintln(w, "Password did not match")
		return
	}

	authToken := uuid.New().String()

	uAuth.authToken = authToken
	uAuth.lastUsed = time.Now()

	if _, err := datastore.Put(ctx, key, &uAuth); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	fmt.Fprintln(w, authToken)
}
*/

type TestDB struct {
	Num int `datastore:"fug" json:"fug"`
}

func dbTestPut(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	myNum, err := strconv.Atoi(vars["num"])

	if err != nil {
		fmt.Fprintln(w, "Thats not a number boyo")
	}

	dbt := TestDB{
		Num: myNum,
	}

	ctx := appengine.NewContext(r)

	key := datastore.NewKey(ctx, "dbTest", "maymay", 0, nil)

	_, err = datastore.Put(ctx, key, &dbt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	key = datastore.NewKey(ctx, "dbTest", "maymay", 0, nil)

	var newt TestDB

	if err := datastore.Get(ctx, key, &newt); err != nil {
		w.WriteHeader(418)
		fmt.Fprintln(w, "Not found")
		return
	}

	fmt.Fprintln(w, "Successfully added "+strconv.Itoa(newt.Num)+", should be "+strconv.Itoa(dbt.Num))
}

func dbTestGet(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	key := datastore.NewKey(ctx, "dbTest", "maymay", 0, nil)

	dbt := TestDB{
		Num: 69,
	}

	if err := datastore.Get(ctx, key, &dbt); err != nil {
		w.WriteHeader(418)
		fmt.Fprintln(w, "Not found")
		return
	}

	writeString := "The number was " + strconv.Itoa(dbt.Num)

	fmt.Fprintln(w, writeString)
}
