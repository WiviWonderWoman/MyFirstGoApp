package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

// RegisterControllers is a function that handels routes //? while keeping the controllers private
func RegisterControllers() {
	uc := newUserController() // calling  the constructor to initializing a new (private) instance of userController
	// Handle registers the handler for the given pattern in the DefaultServeMux //* func Handle(pattern string, handler Handler)
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}

func encodedResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
