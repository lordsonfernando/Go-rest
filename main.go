package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// declare info template with struct
type networkInfo struct {
	Id   string `json:"Id"`
	Name string `json:"Name"`
	Desc string `json:"desc"`
}

//Declare array of network info globally
var info []networkInfo

// function to handle request
func handleRequests() {
	//create router
	infoRouter := mux.NewRouter().StrictSlash(true)
	infoRouter.HandleFunc("/", home)
	infoRouter.HandleFunc("/returninfo/{Id}", returnInfo)
	log.Fatal(http.ListenAndServe(":8080", infoRouter))

}

// home page function
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Network info!")

}

//function to return single id
func returnInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["Id"]
	for _, infoarray := range info {
		if infoarray.Id == key {
			json.NewEncoder(w).Encode(infoarray)

		}
	}
}
func main() {
	fmt.Println("info app")
	info = []networkInfo{
		networkInfo{Id: "1", Name: "router", Desc: "A router is responsible for organising communication between computer networks."},
		networkInfo{Id: "2", Name: "ip address", Desc: "An IP address is a unique address that identifies a device on the internet or a local network."},
	}
	handleRequests()

}
