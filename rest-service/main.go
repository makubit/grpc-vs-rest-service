package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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
	r.Run(":50051")
}

func sort(c *gin.Context) {
	var sortRequest SortRequest
	err := c.ShouldBindJSON(&sortRequest)
	if err != nil {
		log.Println("cannot bind json: ", err)
	}
	fmt.Println("received unsorted table: ", sortRequest.TableToSort)
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(sortRequest)

	client := http.Client{}
	sortResponse, err := client.Post("http://127.0.0.1:50052/sort", "application/json", buf)
	if err != nil {
		c.JSON(500, err)
	}
	var response Response
	_ = json.NewDecoder(sortResponse.Body).Decode(&response)

	c.JSON(200, response)
}

