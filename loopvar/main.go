package main

import (
	"fmt"
	"time"
)

// GOEXPERIMENT=loopvar
func main() {
	numbers := []int{1, 2, 3, 4, 5}
	for _, i := range numbers {
		go func() {
			fancyPrint(i)
		}()
	}
	time.Sleep(10 * time.Millisecond)
}

func fancyPrint(i int) {
	fmt.Printf("number %d\n", i)
}
