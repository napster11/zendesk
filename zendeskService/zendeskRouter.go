package zendeskService

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//BootRouter to startup the router
func BootRouter(port string) {
	router := mux.NewRouter()
	router.Headers("Authorization")
	//API Endpoint is http://localhost:8080/ticketList
	router.HandleFunc("/ticketList", getTicketList).Methods("GET")
	log.Fatal(http.ListenAndServe(port, router))
}
