package main

import (
	"fmt"
	// "net/http"
	// "github.com/WiviWonderWoman/MyFirstGoApp/controllers"
	// "github.com/WiviWonderWoman/My-first-Go-App/models"
)

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
	infiniteLoop(false)
	loopingOverCollections(false)
	panicFunction(false)
	ifStatement(false)
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

// infiniteLoop is a function that describes and print out for loops in Go
func infiniteLoop(run bool) {
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

// loopingOverCollections is a function that describes and print out for loops in Go
func loopingOverCollections(run bool) {
	if run {
		slice := []int{1, 2, 3}
		// loop aslong as i is less than the length(len) of slice
		for i := 0; i < len(slice); i++ {
			fmt.Println("the value:", slice[i], "at index:", i)
		}

		//* Same loop in the shorthanded version syntax
		// the range keyword tells the compiler that we are going to passing in a collection type and will loop through the enthire collection
		for i, v := range slice {
			fmt.Println("i =", i, "v =", v)
		}

		// works with maps as well
		wellKnownPorts := map[string]int{"http": 80, "https": 443}
		// you can omit the v (value) and only get the k (keys) //* output would be: " http https "
		// to omit the k (keys) and only get the v (value) = replace k with _ //* output would be: " 80 443 "
		for k, v := range wellKnownPorts {
			fmt.Println(k, v)
		}
	}
}

// panicFunction is a function that describes and print out a panic function in Go
func panicFunction(run bool) {
	if run {
		fmt.Println("Starting web server")
		//! PANIC when an unrecoverable error happens
		panic("Something bad just happend")
		// the following line will be marked as: //* " unreacheble code ""
		//* fmt.Println("Web server started")
	}
}

// ifStatement is a function that describes and print out a if statements in Go
func ifStatement(run bool) {
	if run {
		type User struct {
			ID        int
			FirstName string
			LastName  string
		}

		u1 := User{
			ID:        1,
			FirstName: "Arthur",
			LastName:  "Dent",
		}

		u2 := User{
			ID:        2,
			FirstName: "Ford",
			LastName:  "Prefect",
		}

		if u1 == u2 {
			fmt.Println("Same user!")
		} else if u1.FirstName == u2.FirstName {
			fmt.Println("Similar user.")
		} else {
			fmt.Println("Diffrent user!")
		}
	}
}
