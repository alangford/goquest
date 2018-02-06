package main


import (
	"log"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"./createIssue"
	"./auth"
)

func heartBeat(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main () {
	ec := make(chan error)
	go func() {
		fmt.Println("Started up this server on port 3000")
		log.Fatalf("Error occurred:\n%v", <-ec)
	}()






	router := mux.NewRouter()
	router.HandleFunc("/api/wow", createIssue.Create()).Methods("POST")
	router.HandleFunc("/api/auth", auth.Login()).Methods("POST")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		ec <- err
	}
}