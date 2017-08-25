package main

import (
	"errors"
	"fmt"
)

func main() {
	// Make a new map that associates integers with errors
	statusCodes := make(map[int]error)

	// Available types in GO
	// bool
	// string
	// int  int8  int16  int32  int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte // alias for uint8
	// rune // alias for int32
	//      // represents a Unicode code point
	// float32 float64
	// complex64 complex128 // WHAT

	// Fill the map
	statusCodes[400] = errors.New("Bad request")
	statusCodes[403] = errors.New("Forbidden")
	statusCodes[404] = errors.New("Bad request")
	statusCodes[500] = errors.New("Internal Server Error")

	currentStatusCode := 404

	// Create a pretty message
	response := fmt.Sprintf("Status code %v: %v", currentStatusCode, statusCodes[currentStatusCode])

	fmt.Println(response)
	fmt.Println("There is a total of ", len(statusCodes), " status codes with messages")

	delete(statusCodes, 500)
	fmt.Println("There is a total of ", len(statusCodes), " status codes with messages after deleting")

	fmt.Println("Remaining statuscodes:")
	// When you loop through a map, you get `k`  and `v`
	for k, v := range statusCodes {
		fmt.Println(fmt.Sprintf("Status code %v: %v", k, v))
	}
}
