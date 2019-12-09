package main

import (
	"bytes"
	"encoding/json"
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
	client := &http.Client{}
	req(client)
}

func req(client *http.Client) {
	sort := &SortRequest{
		TableToSort: []int32{4,3,2,4,3,2},
	}
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(sort)
	if err != nil{
		log.Fatalf("got error: %s", err)
	}
	log.Println(sort)

	sortedTableResponse, _ := client.Post("http://127.0.0.1:50051/sort", "application/json", buf)//nowy request do innej aplikacji

	defer sortedTableResponse.Body.Close()

	var response Response
	_ = json.NewDecoder(sortedTableResponse.Body).Decode(&response)

	log.Println("received sorted table: ", response)
}
