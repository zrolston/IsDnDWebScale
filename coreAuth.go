package main

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
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

	username, pass, ok := r.BasicAuth()
	if !ok {
		w.WriteHeader(440)
		writeMessage(w, "Unable tot get basic auth")
		return
	}

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

func checkValidToken(authToken string, ctx context.Context) (string, bool) {

	var tokens []UserAuth

	_, err := datastore.NewQuery("userAuth").Filter("authToken =", authToken).GetAll(ctx, &tokens)

	if len(tokens) == 0 || err != nil {
		return "Frigg off Ricky", false
	}

	if !checkTime(tokens[0].LastUsed) {
		return "Bad Time", false
	}

	return tokens[0].Username, true
}
