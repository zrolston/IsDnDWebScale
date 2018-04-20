package main

import (
	"encoding/json"
	"net/http"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

func refreshTime(username string, ctx context.Context) error {

	key := datastore.NewKey(ctx, "userAuth", username, 0, nil)

	var updatedAuth UserAuth

	err := datastore.Get(ctx, key, &updatedAuth)
	if err != nil {
		return err
	}

	updatedAuth.LastUsed = time.Now()

	_, err = datastore.Put(ctx, key, &updatedAuth)
	if err != nil {
		return err
	}

	return nil
}

func writeMessage(w http.ResponseWriter, s string) {
	mapD := map[string]string{"str": s}
	mapB, _ := json.Marshal(mapD)
	w.Write(mapB)
}

func checkTime(t time.Time) bool {
	currTime := time.Now()
	duration := currTime.Sub(t)
	if duration.Hours() >= 24 {
		return false
	}

	return true
}

func removeUser(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	username, password, ok := r.BasicAuth()

	if !ok {
		w.WriteHeader(418)
		writeMessage(w, "Basic Auth incorrectly passed")
		return
	}

	if password != "root" {
		w.WriteHeader(418)
		writeMessage(w, "Fuck off Ricky")
		return
	}

	ctx := appengine.NewContext(r)

	key := datastore.NewKey(ctx, "userAuth", username, 0, nil)

	var dst []UserChar
	keys, err := datastore.NewQuery("userChar").Filter("username =", username).GetAll(ctx, &dst)
	if err != nil && err != datastore.ErrNoSuchEntity {
		w.WriteHeader(510)
		writeMessage(w, err.Error())
		return
	}

	topts := &datastore.TransactionOptions{
		XG: true,
	}
	err = datastore.RunInTransaction(ctx, func(ctx context.Context) error {
		return deletTransaction(ctx, key, keys)
	}, topts)
	if err != nil {
		w.WriteHeader(505)
		writeMessage(w, err.Error())
		return
	}

	w.WriteHeader(200)
	writeMessage(w, "Successful deletion")
}

func deletTransaction(ctx context.Context, userKey *datastore.Key, charKeys []*datastore.Key) error {
	err := datastore.Delete(ctx, userKey)
	if err != nil && err != datastore.ErrNoSuchEntity {
		return err
	}

	err = datastore.DeleteMulti(ctx, charKeys)
	if err != nil && err != datastore.ErrNoSuchEntity {
		return err
	}

	return nil
}

func deleteChars(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	authToken, _, ok := r.BasicAuth()
	if !ok {
		w.WriteHeader(420)
		writeMessage(w, "Unable to get basic auth")
		return
	}

	ctx := appengine.NewContext(r)

	username, ok := checkValidToken(authToken, ctx)
	if !ok {
		w.WriteHeader(418)
		writeMessage(w, "Invalid Token!")
		return
	}

	var chars []UserChar
	keys, err := datastore.NewQuery("userChar").Filter("username =", username).GetAll(ctx, &chars)
	if err != nil {
		w.WriteHeader(500)
		writeMessage(w, "Unable to perform query "+err.Error())
		return
	}

	err = datastore.DeleteMulti(ctx, keys)
	if err != nil {
		w.WriteHeader(500)
		writeMessage(w, "Unable to delete: "+err.Error())
		return
	}

	w.WriteHeader(200)
	writeMessage(w, "Successful Deletions")
}

func deleteChar(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	authToken, name, ok := r.BasicAuth()
	if !ok {
		w.WriteHeader(420)
		writeMessage(w, "Incorrect Basic Auth")
	}

	ctx := appengine.NewContext(r)

	username, ok := checkValidToken(authToken, ctx)
	if !ok {
		w.WriteHeader(418)
		writeMessage(w, "Invalid Authentication")
		return
	}

	key := datastore.NewKey(ctx, "userChar", username+" "+name, 0, nil)
	err := datastore.Delete(ctx, key)
	if err != nil {
		w.WriteHeader(440)
		writeMessage(w, "Deletion impossible: "+err.Error())
		return
	}

	w.WriteHeader(200)
	writeMessage(w, "Successful Deletion of "+name)
}

func validToken(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	authToken, _, ok := r.BasicAuth()
	if !ok {
		w.WriteHeader(418)
		writeMessage(w, "Could not get basic auth")
		return
	}
	ctx := appengine.NewContext(r)

	_, ok = checkValidToken(authToken, ctx)
	if !ok {
		w.WriteHeader(418)
		writeMessage(w, "Invalid token")
		return
	}

	w.WriteHeader(200)
	writeMessage(w, "You're guccini")
}
