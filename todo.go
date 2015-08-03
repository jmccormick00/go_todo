package main

import (
	//"html/template"
	"io/ioutil"
	"net/http"
	//"strings"
	"labix.org/v2/mgo"
    //"labix.org/v2/mgo/bson"
)

// Define the ToDo struct
type ToDo struct {
	Title string 	`bson:"title" json:"title"`
	completed bool	`bson:"completed" json:"completed"`
}

const (
	dbName string = "todo"
)

var (
	session *mgo.Session
)

func indexHandler(res http.ResponseWriter, r *http.Request) {
    body, err := ioutil.ReadFile("index.html")
    if err != nil {
		panic(err)
	}
	res.Header().Set("Content-Type", "text/html")
	res.Write(body)

}

func main() {
	var err error
	session, err = mgo.Dial("127.0.0.1")
    if err != nil {
        panic(err)
    }
    defer session.Close()

    // Optional. Switch the session to a monotonic behavior.
    session.SetMode(mgo.Monotonic, true)

	//http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/", indexHandler)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))
	http.ListenAndServe(":8080", nil)
}