// 2. Defines the actual functions that handle API requests

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

// Serves a list of "Todo" items in JSON format
func TodolistIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(list); err != nil { // "list" is from repo.go
		panic(err)
	}
}

// This function is responsible to display information about a specific "Todo" item based on its ID
func TodolistShow(w http.ResponseWriter, r *http.Request) { // function signature
	// this variable call mux.Vars function to extratct variable from the URL
	vars := mux.Vars(r)

	// Parsing idTodo
	var idTodo int
	var err error
	if idTodo, err = strconv.Atoi(vars["idTodo"]); err != nil { // strconv.Atoi converts from string (e.g. "1") to int (1)
		panic(err)
	}

	// Finding the todo item with call a repository
	todo := RepoFindTodo(idTodo)

	// Checking if the todo item exists
	if todo.Id > 0 {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(todo); err != nil {
			panic(err)
		}
		return
	}

	// Handling not found
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}

// This function in charge of the process of a new todo
func TodolistCreate(w http.ResponseWriter, r *http.Request) { // function signature
	// buat variable dengan memanggil struct ToDoList dari todo.go
	var todo ToDoList
	
	// Read all requests
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	// Close the request body
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	// process converting (unmarshalling)
	if err := json.Unmarshal(body, &todo); err != nil { // json.Unmarshal aims to decode the json in body into the todo struct
		// if unmarshalling fails
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(422)
		// encode the error as json and write it to the response
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		// Exit the function to prevent further execution
		return
	}

	// Repository call
	t := ReposCreateTodo(todo)

	// Sucessfull Response 
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)

	if err = json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}