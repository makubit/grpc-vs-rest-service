package main

import (
	"math/rand"
	"sync"

	"context"
	"github.com/micro/go-micro"
	"log"
	"time"

	gr "github.com/makubit/grpc-vs-rest-service/grpc-service/proto/grpcService"
)

const loop  = 1000

func main() {
	time.Sleep(5*time.Second)
	var wg sync.WaitGroup

	srv := micro.NewService(
		micro.Name("grpc.testing.cli"),
	)

	srv.Init()

	client := gr.NewGrpcServiceClient("grpc.service", srv.Client())

	var table []int32
	table = gen1MBTable()

	wg.Add(loop)
	for i:=0; i<loop; i++ {
		go sendRequests(client, table, &wg)
	}
	wg.Wait()
}

func gen1MBTable() []int32 {
	gen := make([]int32, 31250)
	for i:=0; i<31250; i++ {
		gen[i] = rand.Int31()
	}
	return gen
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


