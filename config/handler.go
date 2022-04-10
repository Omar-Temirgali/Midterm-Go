package config

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/Omar-Temirgali/go-service/models"
	"github.com/Omar-Temirgali/go-service/repository"
	"github.com/gorilla/mux"
)

type jsonErr struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "RESTful API with the Go language\n")
}

func ShowAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(repository.Kvmap); err != nil {
		panic(err)
	}
}

func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	pair := repository.RepoFind(key)

	if pair.Value != "" {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(pair); err != nil {
			panic(err)
		}
	} else if pair.Key == "" && pair.Value == "" {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
			panic(err)
		}
	}
}

func UpdateAndInsert(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	p := models.KVmap{Key: vars["key"], Value: vars["value"]}
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &p); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	pair := repository.RepoUpdateAndInsert(p)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(pair); err != nil {
		panic(err)
	}
}
