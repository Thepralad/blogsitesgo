package main

import (
	"fmt"
	"net/http"
	"html/template"
)

//temporary data struct
type Data struct{
	Title string
	Author string
	Views int
}
func main() {
	//temporary data slice
	data := []Data{
		{
			"This is happening, software engineers are being replaced by AI 🤯",
			"Ramujan Banoth",
			30,
		},
		{
			"The reason why can't you study, and its not your phone.",
			"Michael Russell",
			78,
		},
		{
			"How to focus for 5 hours strainght, Scientifily proven 💡",
			"Ryan Rha",
			30,
		},
	}
	tmpl := template.Must(template.ParseFiles("./templates/index.html"))
	
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request){
		tmpl.Execute(res, data)
	})

	fmt.Println("Starting Server")
	http.ListenAndServe(":8080", nil)
}
