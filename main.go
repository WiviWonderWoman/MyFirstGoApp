package main

import (
	"fmt"

	"github.com/WiviWonderWoman/My-first-Go-App/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Tricia",
		LastName:  "McMillan",
	}
	fmt.Println(u)
}
