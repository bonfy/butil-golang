package main

import (
	"net/http"

	butil "github.com/bonfy/butil-golang"
)

// Person struct
type Person struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func handleOK(w http.ResponseWriter, r *http.Request) {
	butil.WriteJsonOK(w)
}

func handleErr(w http.ResponseWriter, r *http.Request) {
	butil.WriteJsonErr(w, "error message")
}

func handleObj(w http.ResponseWriter, r *http.Request) {
	user := Person{ID: 1, Name: "bonfy", Age: 18}
	butil.WriteJsonObject(w, user)
}

func handleList(w http.ResponseWriter, r *http.Request) {
	users := []Person{
		Person{ID: 1, Name: "bonfy", Age: 18},
		Person{ID: 2, Name: "rene", Age: 16},
	}
	butil.WriteJsonObject(w, users)
}

func main() {
	http.HandleFunc("/ok", handleOK)     // http://127.0.0.1:8888/ok
	http.HandleFunc("/err", handleErr)   // http://127.0.0.1:8888/err
	http.HandleFunc("/obj", handleObj)   // http://127.0.0.1:8888/obj
	http.HandleFunc("/list", handleList) // http://127.0.0.1:8888/list

	port := ":8888"
	http.ListenAndServe(port, nil)
}
