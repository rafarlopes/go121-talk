package main

import "fmt"

func main() {
	doStuff()
}

// GODEBUG=panicnil=1
func doStuff() error {
	defer func() {
		fmt.Println("recover")
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()
	fmt.Println("heavy stuff...")
	panic(nil)
}
