package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

type SortRequest struct {
	TableToSort []int32 `json:"tableToSort"`
}

type Response struct {
	Sorted      bool    `json:"sorted"`
	SortedTable []int32 `json:"tableToSort"`
}

const (
	loop = 1
	test1 = 1
	test10 = 10
	test100 = 100
	test1000 = 1000
)

func main() {
	time.Sleep(time.Second * 30)
	var wg sync.WaitGroup

	client := &http.Client{
		Timeout: time.Duration(1000 * time.Second),
	}

	data, _ := ioutil.ReadFile("data.json")
	var sort SortRequest
	_ = json.Unmarshal(data, &sort)

	// TEST CASE #1
	/*log.Println("TEST CASE #1")
	wg.Add(loop * test1)
	for i:=0; i<loop; i++ {
		for j:=0; j<test1; j++ {
			go sendRequests(client, sort ,&wg)
		}
		time.Sleep(time.Second * 10)
	}
	wg.Wait()*/

	// TEST CASE #10
	/*log.Println("TEST CASE #10")
	wg.Add(loop * test10)
	for i:=0; i<loop; i++ {
		for j:=0; j<test10; j++ {
			go sendRequests(client, sort ,&wg)
		}
		time.Sleep(time.Second * 10)
	}
	wg.Wait()*/

	// TEST CASE #100
	/*log.Println("TEST CASE #100")
	wg.Add(loop * test100)
	for i:=0; i<loop; i++ {
		for j:=0; j<test100; j++ {
			go sendRequests(client, sort ,&wg)
			//time.Sleep(time.Millisecond* 500)
		}
		time.Sleep(time.Second * 10)
	}
	wg.Wait()
	log.Println("DONE")*/

	// TEST CASE #1000
	log.Println("TEST CASE #1000")
	wg.Add(test1000)
	for j:=0; j<test1000; j++ {
		go sendRequests(client, sort ,&wg)
	}
	wg.Wait()

	time.Sleep(time.Hour * 1)
}

func sendRequests(client *http.Client, sort SortRequest, wg *sync.WaitGroup) {
	defer wg.Done()

	startJSON := time.Now()

	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(sort)
	if err != nil{
		log.Fatalf("got error: %s", err)
	}

	start := time.Now()

	sortedTableResponse, err := client.Post("http://rest-service.default.svc:50051/sort", "application/json", buf)//nowy request do innej aplikacji
	if err != nil {
		log.Fatalf("fatal: %w", err)
	}
	passed := time.Since(start)

	defer sortedTableResponse.Body.Close()

	var response Response
	_ = json.NewDecoder(sortedTableResponse.Body).Decode(&response)
	passedJSON := time.Since(startJSON)
	log.Println("Sorting time with JSON: ", passedJSON, "Sorting time only: ", passed, " is sorted = ", response.Sorted)
}


/*
1. 1 request
 */