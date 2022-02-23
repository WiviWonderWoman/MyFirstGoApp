package main

import "fmt"

// inorder to work with iota and constant expression the constant has to be accecible from the entire package
const iotaPi = 3.1415

// we can use a constant block to declare multiple constants and use iota
const (
	// we can replace these values with iota
	// every time iota is used it increments its value by one
	first = iota // 0
	// if not specified valu Go reuse the expression above i.e : "iota"
	second // 1

	third = iota + 6 // 2 + 6
	//! << bit shift operation
	fourth = 2 << iota // 2 << 3
)

// if we create multiple constant blocks - then iota "RESETS"
const (
	secondBlockFirst = iota
	secondBlockSecond
)

func main() {

	// * -----------------
	primitveVariables(false)
	pointer(false)
	addressOf(false)
	constants(false)
	constIota(false)
	// * -----------------
}

//primitveVariables is a function to describe and print out the primitive variables in Go
func primitveVariables(run bool) {

	if run {
		// ! Intiger, first declared and then initialized - NOT ok
		// var i int
		// i = 42

		//Floating number: float32 or float64, declared and initialized
		var f float32 = 3.14
		fmt.Println("Pi is a float32", f)

		//Integer, shorthand declared and initialized
		i := 42
		fmt.Println(i, "is an integer")

		//String, shorthand declared and initialized
		firstName := "Arthur"
		fmt.Println("Hello", firstName, "!")

		//Boolean, shorthand declared and initialized
		b := true
		fmt.Println("This is", b)

		//Complex number/datatype
		c := complex(3, 4)
		fmt.Println(c)

		//Split up complex number "c" inte real and imaginary components
		r, im := real(c), imag(c)
		fmt.Println(r, im)
	}
}

// pointer is a function to describe and print out the REFERENCE operator i.e "pointer" in Go
func pointer(run bool) {
	if run {
		//= * is an REFERENCE operator, firstName is a pointer to a string variable, this output will be "<nil>" becauce it's empty
		var firstName1 *string
		fmt.Println(firstName1)
		// to assign a valu to "firstName" we need to DEREFERENCE it whith an *,
		// *firstName1 = "Arthur"
		//! PANIC!! output: "panic: runtime error: invalid memory address or nil pointer dereference"
		// because the pointer is not pointing at anything

		// if we initialize the pointer first, to a new string
		var firstName2 *string = new(string)
		*firstName2 = "Arthur"
		// the output will be "0xc000088220" the memory adress
		fmt.Println(firstName2)
		// when we DEREFERENCE the pointer again we finaly get the output "Atrhur"
		fmt.Println(*firstName2)
	}

}

// addressOf is a function to describe and print out the address of operator in Go
func addressOf(run bool) {
	if run {
		firstName := "Arthur"
		fmt.Println("Test", firstName)

		// & is an address of operator - ptr is a pointer to firstName
		ptr := &firstName
		// output the memory adress and the DEREFERENCE value
		fmt.Println(ptr, *ptr)

		// if we change the value
		firstName = "Tricia"
		// output the SAME memory adress but the NEW value
		fmt.Println(ptr, *ptr)

	}
}

func constants(run bool) {
	if run {
		// pi is a constant is a veriable with fixed value, it has to be initialized when declared.
		// the value of a constant has to be able to be determined at COMPILE time - therefor it can not be a return value from a function
		const pi = 3.1415
		fmt.Println("const", pi)
		// if we try to reassign the value of pi we get the message:
		//! " cannot assign to pi (untyped float constant 3.1415) "
		//* pi =1.2

		// c is an IMPlICITLY typed constant
		const c = 3
		// Go is dynamically interpret the type of c
		fmt.Println(c + 3)   // output: "6"
		fmt.Println(c + 1.2) // output: "4.2"

		// tc is an EXPlICITLY typed constant, an int
		const tc int = 2
		fmt.Println(tc)
		// when trying to add a float to an integer constant message:// * "stympad"
		//! " (untyped float constant) truncated* to int ""
		//* fmt.Println(tc + 1.1)
		// if c is converted to float it will work
		fmt.Println(float32(tc) + 1.2) // output: "3.2"
	}
}

func constIota(run bool) {
	fmt.Println("iotaPi =", iotaPi)                  // output: "iotaPi = 3.1415"
	fmt.Println(first, second)                       // output: " 0 1 " when value is iota and iota
	fmt.Println(third, fourth)                       // output: " 8 16 "
	fmt.Println(secondBlockFirst, secondBlockSecond) // output: " 0 1 "
}
