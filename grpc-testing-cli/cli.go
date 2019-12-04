package main

import (
	"fmt"
	//"encoding/json"
	"github.com/micro/go-micro"
	"time"

	//"io/ioutil"
	"log"
	//"os"

	"context"
	"github.com/gin-gonic/gin"
	//pb "github.com/makubit/grpc-vs-rest-service/grpc-service/proto/consignment"
	gr "github.com/makubit/grpc-vs-rest-service/grpc-service/proto/grpcService"
)

/*const (
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(data, &consignment)
	return consignment, err //or nil, at this point err = nil
}*/

func main() {
	/*service := micro.NewService(micro.Name("grpc.testing.cli"))
	service.Init()

	client := pb.NewShippingServiceClient("grpc.service", service.Client())

	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	getAll, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	for _, v := range getAll.Consignments {
		log.Println(v)
	}*/
	time.Sleep(5*time.Second)

	srv := micro.NewService(
		micro.Name("grpc.testing.cli"),
	)

	srv.Init()

	client := gr.NewGrpcServiceClient("grpc.service", srv.Client())

	r := gin.Default()
	sortedTableResponse, _ := client.GetFromSortingService(context.Background(), &gr.SortRequest{
		TableToSort: []int32{6,5,4,3,2,1},
	})

	r.GET("/gett", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"sortedTable": sortedTableResponse,
		})
	})


	err := r.Run()
	if err != nil {
		fmt.Println("got error: %w", err)
	}

	log.Println("Got from grpc-service table: ", sortedTableResponse.SortedTable)
}

