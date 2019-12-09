package main

import (
	"fmt"
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
	//http.HandleFunc("/sort", sort)
	//log.Println(http.ListenAndServe(":50052", nil))
	r := gin.Default()
	r.POST("/sort", sort)
	r.Run(":50052")
}

/*func sort(req http.ResponseWriter, r * http.Request) {
	decoder := json.NewDecoder(r.Body)
	var unsortedTable []int32
	_ = decoder.Decode(&unsortedTable)
	defer r.Body.Close()

	req.Header().Set("Content-Type", "application/json")
	sortedTable, _ := quickSort(unsortedTable)//nowy request do innej aplikacji

	_ = json.NewEncoder(req).Encode(Response{
		Sorted: true,
		SortedTable: sortedTable,
	})

	log.Println("in sorting app: ", sortedTable)
}*/

func sort(c *gin.Context) {
	var sortRequest SortRequest
	err := c.ShouldBindJSON(&sortRequest)
	if err != nil {
		log.Println("cannot bind json: ", err)
	}
	fmt.Println("received unsorted table: ", sortRequest.TableToSort)

	sortedTable, _ := quickSort(sortRequest.TableToSort)
	fmt.Println("SortedTable: ", sortedTable)

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