package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/olegsobchuk/go-pure/database"
)

// User struct
type User struct {
	Name string
	Age  uint32
}

// SearchResult struct
type SearchResult struct {
	Title string
	Name  string
	Age   uint8
	Count uint8
}

func main() {
	database.Init()
	database.Migrate()

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

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		result := []SearchResult{
			SearchResult{"Title first", "Name First 1", 23, 2},
			SearchResult{"Title second", "Name 2", 13, 21},
			SearchResult{"Title last", "Name 33", 40, 2},
		}

		j, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		fmt.Fprintf(w, "%s", j)
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
