package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"

	"context"
	"github.com/micro/go-micro"
	"log"
	"time"

	gr "github.com/makubit/grpc-vs-rest-service/grpc-service/proto/grpcService"
)

const (
	loop = 100
	test1 = 1
	test10 = 10
	test100 = 100
	test1000 = 1000
)

func main() {
	time.Sleep(10*time.Second)
	var wg sync.WaitGroup

	srv := micro.NewService(
		micro.Name(os.Getenv("APP_NAME")),
	)

	srv.Init()
	cli := gr.NewGrpcServiceClient(os.Getenv("SERV_APP_NAME") + ":8081", srv.Client())

	data, _ := ioutil.ReadFile("data.json")
	var table []int32
	_ = json.Unmarshal(data, &table)

	// TEST CASE #1
	log.Println("TEST CASE #1")
	wg.Add(loop * test1)
	for i:=0; i<loop; i++ {
		for j:=0; j<test1; j++ {
			go sendRequests(cli, table, &wg)
		}
		time.Sleep(time.Second * 10)
	}
	wg.Wait()

	// TEST CASE #10
	/*log.Println("TEST CASE #10")
	wg.Add(loop * test10)
	for i:=0; i<loop; i++ {
		for j:=0; j<test10; j++ {
			go sendRequests(cli, table, &wg)
		}
		time.Sleep(time.Second * 10)
	}
	wg.Wait()*/

	// TEST CASE #100
	/*log.Println("TEST CASE #100")
	wg.Add(loop * test100)
	for i:=0; i<loop; i++ {
		for j:=0; j<test100; j++ {
			go sendRequests(cli, table, &wg)
		}
		time.Sleep(time.Second * 10)
	}
	wg.Wait()
	log.Println("DONE")*/

	// TEST CASE #1000
	/*log.Println("TEST CASE #1000")
	wg.Add(test1000)
	for j:=0; j<test1000; j++ {
		go sendRequests(client, sort ,&wg)
	}
	wg.Wait()*/

	time.Sleep(time.Hour * 1)
}

func sendRequests(client gr.GrpcServiceClient, table []int32, wg *sync.WaitGroup) {
	defer wg.Done()

	start := time.Now()

	sortedTableResponse, err := client.GetFromSortingService(context.Background(), &gr.SortRequest{
		TableToSort: table, //big table to SORT
	})
	if err != nil {
		log.Println(err)
		return
	}

	passed := time.Since(start)
	log.Println("Sorting time: ", passed, " is sorted = ", sortedTableResponse.Sorted)
}


