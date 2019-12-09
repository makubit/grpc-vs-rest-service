package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
)

type SortRequest struct {
	TableToSort []int32 `json:"tableToSort"`
}

type Response struct {
	Sorted      bool    `json:"sorted"`
	SortedTable []int32 `json:"tableToSort"`
}

func main() {
	r := gin.Default()
	r.POST("/sort", sort)
	if r.Run(":50052") != nil {
		log.Fatalf("cannot run gin")
	}
}

func sort(c *gin.Context) {
	var sortRequest SortRequest
	err := c.ShouldBindJSON(&sortRequest)
	if err != nil {
		log.Println("cannot bind json: ", err)
	}

	sortedTable, _ := quickSort(sortRequest.TableToSort)
	sort := &Response{
		SortedTable: sortedTable,
		Sorted: true,
	}

	c.JSON(200, sort)
}

func quickSort(unsortedTable []int32) ([]int32, error) {
	if len(unsortedTable) < 2 {
		return unsortedTable, nil
	}

	left, right := 0, len(unsortedTable)-1

	pivot := rand.Int() % len(unsortedTable)

	unsortedTable[pivot], unsortedTable[right] = unsortedTable[right], unsortedTable[pivot]

	for i, _ := range unsortedTable {
		if unsortedTable[i] < unsortedTable[right] {
			unsortedTable[left], unsortedTable[i] = unsortedTable[i], unsortedTable[left]
			left++
		}
	}

	unsortedTable[left], unsortedTable[right] = unsortedTable[right], unsortedTable[left]

	_, _ = quickSort(unsortedTable[:left])
	_, _ = quickSort(unsortedTable[left+1:])

	return unsortedTable, nil
}