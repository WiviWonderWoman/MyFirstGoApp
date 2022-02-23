package main

import (
	// "fmt"
	"net/http"

	"github.com/WiviWonderWoman/My-first-Go-App/controllers"
	// "github.com/WiviWonderWoman/My-first-Go-App/models"
)

func main() {

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
	/* u := models.User{
		ID:        2,
		FirstName: "Tricia",
		LastName:  "McMillan",
	}
	fmt.Println(u) */
}
