package main

import (
	"fmt"
	"math/rand"
)

func main() {
	/*var gen [31250]int32
	for i:=0; i<len(gen); i++ {
		gen[i] = rand.Int31()
	}

	fmt.Println(gen)*/
	gen := gen1MBTable()
	fmt.Println(gen)
}

func gen1MBTable() []int32 {
	gen := make([]int32, 31250)
	for i:=0; i<31250; i++ {
		gen[i] = rand.Int31()
	}
	return gen
}