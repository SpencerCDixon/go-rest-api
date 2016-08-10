package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"upper.io/db.v2/lib/sqlbuilder"
	"upper.io/db.v2/postgresql"
)

// Database Connection
var dbSettings = postgresql.ConnectionURL{
	Database: `goapi`,
	Host:     `localhost`,
	User:     os.Getenv("DB_USER"),
	Password: os.Getenv("DB_PASS"),
}

func dbConnection() sqlbuilder.Database {
	sess, err := postgresql.Open(dbSettings)
	if err != nil {
		log.Fatal(err)
	}
	return sess
}

// Model
type Todo struct {
	ID       uint   `json:"id,omitempty" db:"id,omitempty"`
	Name     string `json:"name" db:"name"`
	Complete bool   `json:"complete" db:"complete"`
}

type Todos []Todo

// Actions
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	var todos Todos
	db := dbConnection()
	defer db.Close()

	if err := db.Collection("todos").Find().All(&todos); err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

// TODO
func TodoCreate(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var todo Todo

	if err := decoder.Decode(&todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)

		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	// TODO: FIgure out how to not pass in thd ID when creating object

	log.Printf("todo is name: ", todo.Name)
	log.Printf("todo is complete: ", todo.Complete)

	// create todo
	db := dbConnection()
	defer db.Close()
	_, err := db.Collection("todos").Insert(todo)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(todo); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]

	var todo Todo
	db := dbConnection()
	defer db.Close()

	if err := db.Collection("todos").Find("id", todoId).One(&todo); err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(todo); err != nil {
		panic(err)
	}
}
