package main

import (
	"encoding/json"
	"net/http"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

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
