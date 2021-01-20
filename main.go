package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Message struct {
	FirstKey    string
	SecondKey   string
	Name        string
	PhoneNumber string
	ICQ         string
	LastKey     int64
}

func main() {

	var p int
	fmt.Scan(&p)
	fmt.Println(p)
	o := p * 2
	fmt.Println(o)

	m := Message{"World", "Hello", "Dmitry", "79082706690", "393181839", 211}
	fmt.Println(m)
	b, err := json.Marshal(m)
	fmt.Println(b)

	var Log string
	fmt.Scanf("%s\n", &Log)
	// fmt.Println(Log)

	router := mux.NewRouter()

	// router.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println(Log)
	// }
	router.HandleFunc("/main", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, Log)
		fmt.Fprint(w, err)
		fmt.Fprint(w, string(b))
	})

	router.HandleFunc("/2nd", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Server", "A Go Web Server")
		w.WriteHeader(200)
	})

	http.ListenAndServe(":3000", router)

}
