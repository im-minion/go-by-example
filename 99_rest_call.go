package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"bytes"
)

// Todo struct
type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	get()
	getList()
	post()
}

func get() {
	fmt.Println("1. Performing Http Get...")
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println("API Response as String:\n" + bodyString)

	// Convert response body to Todo struct
	var todoStruct Todo
	json.Unmarshal(bodyBytes, &todoStruct)
	fmt.Printf("API Response as struct %+v\n", todoStruct)
}

func getList() {
	fmt.Println("2. Performing Http GetList...")
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println("API Response as String:\n" + bodyString)

	// Convert response body to Todo struct
	var data []Todo
	json.Unmarshal(bodyBytes, &data)
	fmt.Printf("API Response as struct %+v\n", data)
}

func post() {
    fmt.Println("3. Performing Http Post...")
    todo := Todo{1, 2, "lorem ipsum dolor sit amet", true}
    jsonReq, err := json.Marshal(todo)
    resp, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
    if err != nil {
        log.Fatalln(err)
    }

    defer resp.Body.Close()
    bodyBytes, _ := ioutil.ReadAll(resp.Body)

    // Convert response body to string
    bodyString := string(bodyBytes)
    fmt.Println(bodyString)

    // Convert response body to Todo struct
    var todoStruct Todo
    json.Unmarshal(bodyBytes, &todoStruct)
    fmt.Printf("%+v\n", todoStruct)
}
