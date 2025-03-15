package main // should be named main, so that Go knows where to start the execution

import "fmt" // (This is a package. But a predefined one. Check Go Standard Library)

func main() { // should be named main for the same reason. There should be only 1 main function across the code.
	fmt.Println("Hello world!")
	fmt.Println("Hello world!")
}
