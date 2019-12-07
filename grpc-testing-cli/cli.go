package main

import (
	"github.com/micro/go-micro"
	"math/rand"

	//"encoding/json"
	//"github.com/micro/go-micro"
	k8s "github.com/micro/examples/kubernetes/go/micro"
	"time"

	//"io/ioutil"
	"log"
	//"os"

	"context"
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

	srv := k8s.NewService(
		micro.Name("grpc.testing.cli"),
		micro.Version("latest"),
	)

	srv.Init()

	client := gr.NewGrpcServiceClient("grpc.service", srv.Client())

	log.Println("Starting sorting table")
	//r := gin.Default()
	//var table []int32
	//table = gen1MBTable()

	start := time.Now()
	sortedTableResponse, err := client.GetFromSortingService(context.Background(), &gr.SortRequest{
		TableToSort: []int32{6,5,4,3,2,1},
		//TableToSort: table, //big table to SORT
	})
	//log.Println("Table size: ", len(sortedTableResponse.SortedTable))
	passed := time.Since(start)
	log.Println("Sorting time: ", passed)
	log.Println(err)

	//log.Println("Got from grpc-service table: ", sortedTableResponse.SortedTable)

	/*r.GET("/gett", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"sortedTable": sortedTableResponse,
		})
	})*/

	/*err := r.Run()
	if err != nil {
		fmt.Println("got error: %w", err)
	}*/

	log.Println("Got from grpc-service table: ", sortedTableResponse)
}

func gen1MBTable() []int32 {
	gen := make([]int32, 31250)
	for i:=0; i<31250; i++ {
		gen[i] = rand.Int31()
	}
	return gen
}



