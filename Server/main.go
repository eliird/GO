package main

import (
	"fmt"
	"net/http"
)

func handleFunc(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello From the global Function of the program"))
}

func main(){

	server := http.NewServeMux()
	server.HandleFunc("/hello", handleFunc)
	
	fs := http.FileServer(http.Dir("./public/"))
	server.Handle("/", fs)
	err := http.ListenAndServe(":8000", nil)//Opens a port to listen
	
	if err == nil{
		fmt.Println("Server could not be opened")
	}
}
