package main

import (
	"fmt"
	"log"
	"net/http"

	"gorm.io/gorm"
)

var DB *gorm.DB

func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

func main() {
	DB = ConnectToDB()
	fmt.Println("the db has been obtained", DB)
	r := HandlerRouting()

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
