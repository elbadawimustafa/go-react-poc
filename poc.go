package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func main(){
	router := mux.NewRouter()
	router.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			f,_  := ioutil.ReadFile("./home.html")
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "text/html")
			_,_ = w.Write(f)
		})

	log.Fatal(http.ListenAndServe(":8080",router))
}
