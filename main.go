package main

import (
	"github.com/dyeduguru/segment/server"
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", server.NewRouter())
}
