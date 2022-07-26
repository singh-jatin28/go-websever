package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type user struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		/*
			//first part of the question
			person := user{Name: "HTTP", Email: "test@example.com"}
			res, err := json.Marshal(person)
			if err != nil {
				fmt.Fprintf(w, "error while marshaling: ")
			} else {
				w.Header().Set("Content-Type", "application/octet-stream")
				w.WriteHeader(http.StatusOK)
				w.Write(res)
			}
		*/
		w.Header().Set("Content-Type", "application/octet-stream")
		w.WriteHeader(http.StatusOK)
		file, err := ioutil.ReadFile("test.txt")
		if err != nil {
			fmt.Fprintf(w, "Error: %v", err)
		}

		w.Write(file)
	case "POST":
		person := user{}

		if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		f, err := os.OpenFile("test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Fprintf(w, "Error while opening file: %v", err)
		}
		f.WriteString("Name: " + person.Name + "\n")
		f.WriteString("Email: " + person.Email + "\n")
		f.WriteString("\n")

		fmt.Fprintf(w, "Written: \n")
		fmt.Fprintf(w, "Name = %s\n", person.Name)
		fmt.Fprintf(w, "Email = %s\n", person.Email)
		f.Close()
	default:
		fmt.Fprintf(w, "This webserver only supports GET and POST methods.")
	}
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
