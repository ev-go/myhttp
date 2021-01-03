package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	var Log string
	fmt.Scanf("%s\n", &Log)
	// fmt.Println(Log)

	router := mux.NewRouter()

	// router.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println(Log)
	// }
	router.HandleFunc("/main", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, Log)
		fmt.Fprint(w, r)
	})

	http.ListenAndServe(":3000", router)

}
