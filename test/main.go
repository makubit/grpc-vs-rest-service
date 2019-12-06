package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var gen [31250]int32
	for i:=0; i<len(gen); i++ {
		gen[i] = rand.Int31()
	}

	fmt.Println(gen)
}
