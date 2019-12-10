package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
)

type SortRequest struct {
	TableToSort []int32 `json:"tableToSort"`
}

func main() {
	/*var gen [31250]int32
	for i:=0; i<len(gen); i++ {
		gen[i] = rand.Int31()
	}

	fmt.Println(gen)*/
	gen := gen1MBTable()
	sort := SortRequest{
		TableToSort: gen,
	}

	file, _ := json.Marshal(sort)
	_ = ioutil.WriteFile("data.json", file, 0644)
}

func gen1MBTable() []int32 {
	gen := make([]int32, 100000)
	for i:=0; i<100000; i++ {
		gen[i] = rand.Int31()
	}
	return gen
}