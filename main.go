package main

import (
	"fmt"
	gmux "github.com/gorilla/mux"
	_ "github.com/lovohh/lovohhwebapi/infrastructure/deps"
	"net/http"
)

func main() {
	router := gmux.NewRouter()

	router.HandleFunc("/test", func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("Hello world"))
	})

	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	if serverErr := server.ListenAndServe(); serverErr != nil {
		fmt.Println("Server failed to boot: " + serverErr.Error())
	}
}
