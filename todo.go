// 1. defined the data structure 

package main

import("time")

type ToDoList struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Completed bool `json:"completed"`
	Due time.Time `json:"due"`
}

type List []ToDoList