package main

//need access to database to read and write data

import (
	"net/http"
	"fmt"
)

func main(){
	mux := http.NewServeMux();
	mux.Handle("/", http.FileServer(http.Dir("./templates")))
	mux.HandleFunc("/register", register)
	mux.HandleFunc("/login", login)
	http.ListenAndServe(":8080", mux)
}

func register(res http.ResponseWriter, req *http.Request){
	err := req.ParseForm()
	if err != nil{
		fmt.Println(err)
	}
	updateDB(req.FormValue("email"), req.FormValue("password"))
	fmt.Fprintf(res, "EMAIL: %v, PASSWORD: %v\n", req.FormValue("email"), req.FormValue("password"))
}

func login(res http.ResponseWriter, req *http.Request){
	err := req.ParseForm()
	if err != nil{
		fmt.Println(err)
	}
	enteredEmail := req.FormValue("email");
	
	_, password, _ := retrieveRow(enteredEmail)
	if password != req.FormValue("password"){
		fmt.Fprintf(res, "worng password")	
	}else{
		fmt.Fprintf(res, "successfully logged in! EMAIL: %v, PASSWORD: %v\n", req.FormValue("email"), req.FormValue("password"))	
	}
}
