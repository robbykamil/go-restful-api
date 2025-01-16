// ToDo Repository that provides CRUD operations

package main

import("fmt")


var currentID int
var list List // "List" is a struct from todo.go

func init() {
    ReposCreateTodo(ToDoList{Name: "Plan A"})
    ReposCreateTodo(ToDoList{Name: "Plan B"})
}

// return as TodoList at todo.go
// this function is used for TodolistShow() at handler.go
func RepoFindTodo(id int) ToDoList{
    for _, t := range list {
        if t.Id == id {
            return t
        }
    }
    return ToDoList{}
}


// return as TodoList at todo.go
// type of parameter as TodoList struct
// this function is used for TodolistCreate() at handler.go
func ReposCreateTodo(t ToDoList) ToDoList{
    currentID+=1
    t.Id = currentID
    list = append(list, t)
    return t
}

// 
func RepoDestroyTodo(id int) error {
    for i, t := range list { // i represent the index, and t for actual elements
        if t.Id == id {
            // list[:i] is select for all elements from the beginning up to (but not including) the current index
            // 
            list = append(list[:i],list[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("Couln't find %d", id)
}