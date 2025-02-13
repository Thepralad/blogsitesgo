package handlers

import(
	"net/http"
	"github.com/Thepralad/blogsites/models"
	"fmt"
)

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
