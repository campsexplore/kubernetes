package main

import (
	"fmt"
	"net/http" 
	"os"
	"io/ioutil"
	"log"
) 

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/config", ConfigMap)
	http.HandleFunc("/secret", Secret)
	http.ListenAndServe(":80", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")

	fmt.Fprintf(w, "Hello, my name is %s and I am %s years old", name, age)
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("config/config.txt")

	if err != nil {
		log.Fatalf("File reading error", err)
		return
	}

	fmt.Fprintf(w, "Config: %s", string(data))
}

func Secret(w http.ResponseWriter, r *http.Request) {
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")

	fmt.Fprintf(w, "User: %s. Password: %s", user, password)
}

