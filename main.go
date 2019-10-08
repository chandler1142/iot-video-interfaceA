package main

import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "get called"}`))
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "post called"}`))
	case "PUT":
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{"message": "put called"}`))
	case "DELETE":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "delete called"}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	go register()
	log.Fatal(http.ListenAndServe(":8080", r))

}

func register() {
	var updatedStatus = false

	for {
		ch := time.After(5 * time.Second)
		select {
		case <-ch:
			if (!updatedStatus) {
				fmt.Println("Try to register status")
				response, err := http.Post("http://localhost:8081/Request_Update_Status", "application/xml", bytes.NewBuffer([]byte(`<?xml version="1.0" encoding="UTF-8"?>
<HTTP_XML EventType="Request_Update_Status" >
	<Item From="RequestCode" To="TargetCode" Status="1" HttpURL="127.0.0.1:8080" SipURL="127.0.0.1:5060" />
</HTTP_XML>`)))
				if err != nil {
					fmt.Printf("err occur when request to update status : %v \n ", err)
					continue
				}
				fmt.Printf("request response: %v \n", response)
				updatedStatus = true
			} else {
				fmt.Println("registered...")
			}
		}
	}
}
