package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Name struct {
	FirstName string
	LastName  string
}

type Message struct {
	FirstKey  string
	SecondKey string
	Name
	PhoneNumber string
	ICQ         string
	LastKey     int64
}

type Music struct {
	Genre Genre
}

type Genre struct {
	Country string
	Rock    string
}

func main() {

	music := Music{Genre: Genre{Country: "Taylor Swift", Rock: "Aimee"}}

	data, err := json.Marshal(music)
	if err != nil {
		fmt.Println("Oh no! There was an error:", err)
		return
	}

	fmt.Println(string(data))

	var p int
	fmt.Scan(&p)
	fmt.Println(p)
	o := p * 2
	fmt.Println(o)

	//nameStruct := Name{"Dmitry", "Victorovich"}
	m := Message{"World", "Hello", Name{"Dmitry", "Victorovich"}, "79082706690", "393181839", 211}
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
		fmt.Fprint(w, string(data))
	})

	router.HandleFunc("/2nd", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Server", "A Go Web Server")
		w.WriteHeader(200)
	})

	http.ListenAndServe(":3000", router)

}
