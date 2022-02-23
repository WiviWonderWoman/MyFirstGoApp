package main

import "fmt"

func structs(run bool) {
	if run {
		fmt.Println("Struct")
	}
}

func main() {
	arrays(false)
	slices(false)
	maps(false)
	structs(true)
}

// arrays is a funktion to describe and print out arrays in Go
func arrays(run bool) {
	if run {

		var longArr [3]int // the "long" way to declare an array
		longArr[0] = 1
		longArr[1] = 2
		longArr[2] = 3
		fmt.Println(longArr)    // output: [1 2 3]
		fmt.Println(longArr[1]) // output: 2
		// reaching outside the array = compiletime error with message:
		//! "invalid argument: index 4 (constant of type int) is out of bound"
		//* fmt.Println(arr[4])

		arr := [3]int{1, 2, 3} // shorthand way to declare an array
		fmt.Println(arr)       // output: [1 2 3]
		fmt.Println(arr[1])    // output: 2
	}
}

// slices is a funktion to describe and print out slices in Go
func slices(run bool) {
	if run {
		arr := [3]int{1, 2, 3}
		arrSlice := arr[:]                                // arrSlice is a slice (copy) of the entire array "arr"
		fmt.Println("array:", arr, "arrSlice:", arrSlice) // Output: " array: [1 2 3] slice: [1 2 3]

		// even when changed values in either of them we get identicall output // ? the slice "POINTS" at the value that the array is keeping
		arr[1] = 42
		arrSlice[2] = 27
		fmt.Println("array:", arr, "arrSlice:", arrSlice) // Output: "array: [1 42 27] slice: [1 42 27]"

		slice := []int{1, 2, 3} // declaring a slice without specifying an underlaying array // ? the compiler will create that array for us
		fmt.Println("slice:", slice)

		// we can add one or how many values we want to a slice // ? Go will create a copy off the underlaying array and add slots for a new
		slice = append(slice, 4, 5, 6, 7, 8, 9)
		fmt.Println("slice with appended numbers:", slice)

		//creating slices of slices
		s2 := slice[1:]         // s2 starts and ICLUDES index 1
		s3 := slice[:2]         // s3 ends and EXCLUDES index 2
		s4 := slice[1:2]        // s4 starts and ICLUDES index 1 && ends and EXCLUDES index 2
		fmt.Println(s2, s3, s4) // output: "[2 3 4 5 6 7 8 9] [1 2] [2]"
	}
}

// maps is a funktion to describe and print out maps in Go
func maps(run bool) {
	if run {
		m := map[string]int{"foo": 42}                  // a map consist of keys and values, m has string key and integer values
		fmt.Println("map:", m)                          // output: "map: map[foo:42]"
		fmt.Println("value of 'foo':", m["foo"])        // output: "value of 'foo': 42"
		m["foo"] = 27                                   // change value of 'foo'
		fmt.Println("value of 'foo' is now:", m["foo"]) // output: "value of 'foo' is now: 27"
		delete(m, "foo")                                // deleting element out of a map
		fmt.Println("empty map:", m)                    // output: "empty map: map[]"
	}
}
