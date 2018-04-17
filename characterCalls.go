package main

import (
	"encoding/json"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

type UserChar struct {
	Username string    `datastore:"username" json:"username"`
	Char     Character `datastore:"char" json:"char"`
}

func getCharList(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	authToken, _, _ := r.BasicAuth()
	ctx := appengine.NewContext(r)

	username, ok := checkValidToken(authToken, ctx)
	if !ok {
		w.WriteHeader(418)
		writeMessage(w, "Invalid Token!")
	}

	var userCharList []UserChar
	_, err := datastore.NewQuery("userChar").Filter("username =", username).GetAll(ctx, &userCharList)
	if err != nil {
		w.WriteHeader(500)
		writeMessage(w, err.Error())
		return
	}

	leng := len(userCharList)

	charList := make([]Character, leng)
	for i := 0; i < leng; i++ {
		charList[i] = userCharList[i].Char
	}

	response, _ := json.Marshal(charList)
	w.WriteHeader(200)
	w.Write(response)
}

func putChar(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	authToken, _, ok := r.BasicAuth()
	if !ok {
		w.WriteHeader(418)
		writeMessage(w, "Could not get basic auth")
		return
	}
	ctx := appengine.NewContext(r)

	username, ok := checkValidToken(authToken, ctx)
	if !ok {
		w.WriteHeader(418)
		writeMessage(w, "Invalid token")
		return
	}

	decoder := json.NewDecoder(r.Body)
	var char Character
	err := decoder.Decode(&char)
	if err != nil {
		w.WriteHeader(420)
		writeMessage(w, err.Error())
		return
	}

	userChar := UserChar{
		Username: username,
		Char:     char,
	}

	key := datastore.NewKey(ctx, "userChar", username+" "+char.Name, 0, nil)

	var user UserChar
	if err := datastore.Get(ctx, key, &user); err == nil {
		w.WriteHeader(440)
		writeMessage(w, "Name already taken")
		return
	}

	_, err = datastore.Put(ctx, key, &userChar)
	if err != nil {
		w.WriteHeader(505)
		writeMessage(w, "WHAT THE FUCK")
		return
	}

	w.WriteHeader(200)
	writeMessage(w, "Successful Write")
}

func emptyChar(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	char := Character{}

	response, err := json.Marshal(char)
	if err != nil {
		w.WriteHeader(505)
		writeMessage(w, "Unable to marshal")
		return
	}
	w.WriteHeader(200)
	w.Write(response)
}
