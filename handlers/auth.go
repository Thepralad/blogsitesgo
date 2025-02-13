package handlers

import(
	"net/http"
	"github.com/thepralad/blogsitesgo/models"
	"fmt"
)

//Handler for the /register route.
func Register(res http.ResponseWriter, req *http.Request){
	err := req.ParseForm()
	if err != nil{
		fmt.Println(err)
	}
	models.UpdateDB(req.FormValue("email"), req.FormValue("password"))
	fmt.Fprintf(res, "EMAIL: %v, PASSWORD: %v\n", req.FormValue("email"), req.FormValue("password"))
}

//Handler for the /login route
func Login(res http.ResponseWriter, req *http.Request){
	err := req.ParseForm()
	if err != nil{
		fmt.Println(err)
	}
	enteredEmail := req.FormValue("email");
	
	_, password, _ := models.RetrieveRow(enteredEmail)
	if password != req.FormValue("password"){
		fmt.Fprintf(res, "worng password")	
	}else{
		fmt.Fprintf(res, "successfully logged in! EMAIL: %v, PASSWORD: %v\n", req.FormValue("email"), req.FormValue("password"))	
	}
}
