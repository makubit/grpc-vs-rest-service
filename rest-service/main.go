package main

import (
	"bytes"
	"encoding/json"
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
	if r.Run(":50051") != nil {
		log.Fatalf("cannot run gin")
	}
}

func sort(c *gin.Context) {
	var sortRequest SortRequest
	err := c.ShouldBindJSON(&sortRequest)
	if err != nil {
		log.Println("cannot bind json: ", err)
	}

	buf := new(bytes.Buffer)
	err = json.NewEncoder(buf).Encode(sortRequest)
	if err != nil {
		log.Fatalf("got error in encoding: %w", err)
	}

	client := http.Client{}
	sortResponse, err := client.Post("http://rest-sorting-service.default.svc:50052/sort", "application/json", buf)
	if err != nil {
		c.JSON(500, err)
	}
	var response Response
	_ = json.NewDecoder(sortResponse.Body).Decode(&response)

	c.JSON(200, response)
}

