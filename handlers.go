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
	ID        uint   `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Completed bool   `json:"completed" db:"completed"`
}

type Todos []Todo

// Actions
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	var todos Todos
	sess := dbConnection()
	defer sess.Close()

	if err := sess.Collection("todos").Find().All(&todos); err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// log.Println("Todos: ")
	// for _, todo := range todos {
	// log.Printf("%q (ID: %d)\n", todo.Name, todo.ID)
	// }

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}
