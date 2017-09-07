package server

import (
	"fmt"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.
		Methods("GET").
		Path(fmt.Sprintf("/register/{%s}", usernameVar)).
		HandlerFunc(Register)
	router.
		Methods("GET").
		Path("/users/").
		HandlerFunc(List)
	router.
		Methods("GET").
		Path(fmt.Sprintf("/session/{%s}/{%s}", senderVar, receiverVar)).
		HandlerFunc(CreateSession)
	router.
		Methods("GET").
		Path(fmt.Sprintf("/session/{%s}", receiverVar)).
		HandlerFunc(ListSessions)
	router.
		Methods("POST").
		Path(fmt.Sprintf("/send/{%s}", sessionIDVar)).
		HandlerFunc(SendMessage)
	router.
		Methods("GET").
		Path(fmt.Sprintf("/list/{%s}", sessionIDVar)).
		HandlerFunc(ListMessages)
	return router
}
