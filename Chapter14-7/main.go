package main

import (
	"flag"
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
)

func HomePageHandle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	
	if vars["name"] == "" {
		name = "World"
	}
	
	/*name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}*/
	fmt.Fprintf(w, "Hello, %s!", name)
}

func UsersHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Users Page")
}

func NewRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/{name}", HomePageHandle).Methods("GET")
	r.HandleFunc("/users",UsersHandle).Methods("GET")
	return r
}

func main() {
	var port string
	flag.StringVar(&port, "port", ":3000", "default port: 3000")
	flag.Parse()

	http.ListenAndServe(port, NewRouter())
}
