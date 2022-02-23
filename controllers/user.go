package controllers

import (
	"net/http"
	"regexp"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

//* newUserController is a CONSTUCTOR function
func newUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile(`^/useers/(\d+)/?`),
	}
}

func (uc userController) ServHTTP(write http.ResponseWriter, req *http.Request) {
	write.Write([]byte("Hello from User Controller!"))
}
