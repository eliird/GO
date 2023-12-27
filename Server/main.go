package main

import (
	"femm/api"
	"femm/data"
	"fmt"
	"html/template"
	"net/http"
)

func handleFunc(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello From the global Function of the program"))
}

func handleTemplateTest(w http.ResponseWriter, r *http.Request){
	html, err := template.ParseFiles("./templates/index.tmpl")
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error!"))
		return
	}
	html.Execute(w, "Test!")

}
func handleTemplate(w http.ResponseWriter, r *http.Request){
	html, err := template.ParseFiles("./templates/index.tmpl")
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error!"))
		return
	}
	//html.Execute(w, data.GetAll()[1])//In the template file you have to pass the property with the '.', if you just use the '.' it will be weirdi
	//To serve all the values in the list we need to add range in the template file
	html.Execute(w, data.GetAll())
}


func main(){

	//server := http.NewServeMux()// Using ServeMux sometimes works but sometimes does not work
	http.HandleFunc("/hello", handleFunc) //http works but server sometimes serves files sometimes it does not
	http.HandleFunc("/templateTest", handleTemplateTest)//Passes a variable to template to be displayed to the website in this case "Test"	
	http.HandleFunc("/api/ex", api.Get)
	http.HandleFunc("/template", handleTemplate)
	http.HandleFunc("/api/post", api.Post)
	fs := http.FileServer(http.Dir("./public/"))
	http.Handle("/", fs)
	err := http.ListenAndServe(":8004", nil)//Opens a port to listen
	
	if err != nil{
		fmt.Println("Server could not be opened")
	}
}
