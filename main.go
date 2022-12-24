package main

import (
	"firstApp/pkg/routes"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)


func main() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }

    log.SetOutput(file)
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	fmt.Printf("Starting server at Port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

}
