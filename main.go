package main

import (
	// "fmt"
	"net/http"

	"github.com/WiviWonderWoman/MyFirstGoApp/controllers"
	// "github.com/WiviWonderWoman/My-first-Go-App/models"
)

func main() {
	// call function RegisterControllers
	controllers.RegisterControllers()
	// ListenAndServe listens on the TCP network address addr and then calls Serve with handler to handle requests on incoming connections.
	//The handler is typically nil, in which case the DefaultServeMux is used.
	// ListenAndServe always returns a non-nil error. //* func ListenAndServe(addr string, handler Handler) error
	http.ListenAndServe(":3000", nil)

	/* u := models.User{
		ID:        2,
		FirstName: "Tricia",
		LastName:  "McMillan",
	}
	fmt.Println(u) */
}
