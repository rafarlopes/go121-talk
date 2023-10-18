package iotaissue

import (
	"fmt"
)

type foo int

const (
	FANCY = "aaaaaa"
	XXXX  = "bbbb"

	typeA foo = iota
	typeB
	typeC
)

func main() {
	fmt.Printf("a %T\n", typeA)
	fmt.Printf("b %T\n", typeB)
	fmt.Printf("c %T\n", typeC)
}
func pointer(i int) *int {
	return &i
}
