package zendeskService

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//BootRouter to startup the router
func BootRouter(port string) {

	//To handle CORS
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// start server listen with error handling
	router := mux.NewRouter()
	router.HandleFunc("/ticketList", getTicketList).Methods("GET")
	log.Fatal(http.ListenAndServe(port, handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
