package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DixitLukhi/GOLANG-LCO/mongoAPi/router"
)

func main() {
	fmt.Println("DB API")

	r := router.Rouer()

	r.HandleFunc("/", serveHome).Methods("GET")
	fmt.Println("server starting...")
	log.Fatal(http.ListenAndServe(":5000", r))
	fmt.Println("Listening at port 5000")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Home for API in GOLANG 5000</h1>"))
}
