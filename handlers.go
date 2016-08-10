package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"upper.io/db.v2/postgresql"
)

var dbSettings = postgresql.ConnectionURL{
	Database: `goapi`,
	Host:     `localhost`,
	User:     os.Getenv("DB_USER"),
	Password: os.Getenv("DB_PASS"),
}

type Todo struct {
	ID        uint   `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Completed bool   `json:"completed" db:"completed"`
}

type Todos []Todo

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Testing the env file", os.Getenv("DB_HOST"))
	var todos Todos
	sess, err := postgresql.Open(dbSettings)
	if err != nil {
		log.Fatal(err)
	}
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

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}
