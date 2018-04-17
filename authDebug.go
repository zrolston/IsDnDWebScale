package main

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

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
	token, _, ok := r.BasicAuth()
	if !ok {
		w.WriteHeader(440)
		writeMessage(w, "Unable to get basic auth")
		return
	}
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
