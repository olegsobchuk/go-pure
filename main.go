package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name string
	Age  uint32
}

func main() {
	templates := template.Must(template.ParseFiles("app/views/index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		u := User{Name: "Oleg"}
		if name := r.FormValue("name"); name != "" {
			u.Name = name
		}
		if err := templates.ExecuteTemplate(w, "index.html", u); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// sessions
	http.HandleFunc("/session/new", func(w http.ResponseWriter, r *http.Request) {
		if err := templates.ExecuteTemplate(w, "new.html", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	http.HandleFunc("/session/destroy", destroySessionHandleFunc)

	fmt.Println("Listening server at 127.0.0.1:8000")
	fmt.Println(http.ListenAndServe(":8000", nil))
}

func destroySessionHandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, It's Destroy Session page")
}
