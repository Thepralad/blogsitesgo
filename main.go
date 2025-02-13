package main

//need access to database to read and write data

import (
	"net/http"
	"github.com/thepralad/blogsitesgo/handlers"
)

func main(){
	//Create a new mux for routing
	mux := http.NewServeMux();

	//Managing all routes
	mux.Handle("/", http.FileServer(http.Dir("./templates")))
	mux.HandleFunc("/register", handlers.Register)
	mux.HandleFunc("/login", handlers.Login)

	//setting up a server at :8080
	http.ListenAndServe(":8080", mux)
}

