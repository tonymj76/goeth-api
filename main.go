package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
	Handler "github.com/tonymj76/goeth-api/handler"
)

func main() {
	// Create a client instance to connect to our privider
	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		fmt.Println(err)
	}
	// Create a mux router
	r := mux.NewRouter()

	// we will define a single endpoint
	r.Handle("/api/v1/eth/{module}", Handler.ClientHandler{client})
	log.Fatal(http.ListenAndServe(":8080", r))
}
