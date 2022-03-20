package main

import (
	"log"
	"net/http"
	"html/template"
)

// check - Error handling function to
// use everywhere in order to check errors.
func check(err error){
	if err != nil{
		log.Fatal(err)
	}
}

// viewHandler - Handle browser requests for a certain path.
func viewHandler(writer http.ResponseWriter, request *http.Request){
	html, err := template.ParseFiles("view.html")
	check(err)
	err = html.Execute(writer, nil)
}

// The app will responding to requests.
func main(){
	http.HandleFunc("/guestbook", viewHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}