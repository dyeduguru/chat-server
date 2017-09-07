package server

import (
	"net/http"

	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
)

var (
	usernameVar  = "username"
	senderVar    = "sender"
	receiverVar  = "receiver"
	sessionIDVar = "sessionID"
	userStore    = &UserStore{users: make(map[string]struct{})}
	sessions     = &Sessions{sessions: make(map[int]*Session)}
)

func Register(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars[usernameVar]
	userStore.Add(username)
	w.WriteHeader(http.StatusOK)
}

func List(w http.ResponseWriter, r *http.Request) {
	users := userStore.List()
	if err := json.NewEncoder(w).Encode(users); err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
}

func CreateSession(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sender := vars[senderVar]
	receiver := vars[receiverVar]
	id := sessions.NewSession(sender, receiver)
	if err := json.NewEncoder(w).Encode(id); err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
}

func ListSessions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	receiver := vars[receiverVar]
	ids := sessions.List(receiver)
	if err := json.NewEncoder(w).Encode(ids); err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
}

func SendMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sessionID := vars[sessionIDVar]
	decoder := json.NewDecoder(r.Body)
	var m Message
	if err := decoder.Decode(&m); err != nil {
		panic(err)
	}
	id, err := strconv.Atoi(sessionID)
	if err != nil {
		panic(err)
	}
	session := sessions.Get(id)
	session.AddMessage(m)
	w.WriteHeader(http.StatusOK)
}

func ListMessages(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sessionID := vars[sessionIDVar]
	id, err := strconv.Atoi(sessionID)
	if err != nil {
		panic(err)
	}
	//TODO: Add authz for listing messages of a session
	session := sessions.Get(id)
	if err := json.NewEncoder(w).Encode(session.GetMessages()); err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
}
