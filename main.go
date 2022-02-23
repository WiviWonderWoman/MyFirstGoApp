package main

import (
	"fmt"
	// "net/http"
	// "github.com/WiviWonderWoman/MyFirstGoApp/controllers"
	// "github.com/WiviWonderWoman/My-first-Go-App/models"
)

func infiniteLoops(run bool) {
	if run {
		var i int
		for {
			if i == 5 {
				break //! without break this loop will be  INFINITE
			}
			fmt.Println(i)
			i++
		}
	}
}

func main() {
	// call function RegisterControllers
	// controllers.RegisterControllers()
	// ListenAndServe listens on the TCP network address addr and then calls Serve with handler to handle requests on incoming connections.
	//The handler is typically nil, in which case the DefaultServeMux is used.
	// ListenAndServe always returns a non-nil error. //* func ListenAndServe(addr string, handler Handler) error
	// http.ListenAndServe(":3000", nil)

	/* u := models.User{
		ID:        2,
		FirstName: "Tricia",
		LastName:  "McMillan",
	}
	fmt.Println(u) */
	loopWithCondition(false)
	loopTillConditionWithPostClause(false)
	infiniteLoops(true)
}

// loopWithCondition is a function that describes and print out for loops in Go
func loopWithCondition(run bool) {
	if run {

		var i int
		// loop untill i is less then 5
		for i < 5 {
			fmt.Println("loop- lap:", i) //* if break output will be: "loop- lap: 0 loop- lap: 1 loop- lap: 2"
			i++                          // mutate i so we don't get an infinite loop
			if i == 3 {
				// break //* stop the enthire loop before conditions are met
				continue // continue till next lap without printing the 'continuing...'
			}
			fmt.Println("continuing...") // outputs NO 'continuing...' between 2 and 3
		}
	}
}

// loopTillConditionWithPostClause is a function that describes and print out for loops in Go
func loopTillConditionWithPostClause(run bool) {
	if run {
		for i := 0; i < 5; i++ {
			fmt.Println("loop", i)
		}
		// can be declared like this:
		var i int
		for {
			if i == 5 {
				break
			}
			fmt.Println(i)
			i++
		}
	}
}
