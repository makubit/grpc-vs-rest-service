package main

import (
	"context"
	"fmt"
	"os"
	"sync"

	gr "github.com/makubit/grpc-vs-rest-service/grpc-service/proto/grpcService"
	sCli "github.com/makubit/grpc-vs-rest-service/grpc-sorting-service/proto/sortingService"
	"github.com/micro/go-micro"
)

type repository interface {
}

type Repository struct {
	mu sync.RWMutex
	sortedTableRequest []int32
}

type service struct {
	repo repository
	sortingClient sCli.SortingServiceClient
}

func (s *service) GetFromSortingService(ctx context.Context, req *gr.SortRequest, res *gr.Response) (error) {
	sortingResponse, _ := s.sortingClient.Sort(context.Background(), &sCli.SortRequest{
		Sorted: false,
		TableToSort: req.TableToSort,
	})
	res.Sorted = sortingResponse.Sorted
	res.SortedTable = sortingResponse.SortedTable

	return nil
}

func main() {
	repo := &Repository{}

	srv := micro.NewService(
		micro.Name(os.Getenv("APP_NAME")),
		)
	srv.Init()

	sortingClient := sCli.NewSortingServiceClient(os.Getenv("SORT_APP_NAME") + ":8082", srv.Client())
	gr.RegisterGrpcServiceHandler(srv.Server(), &service{repo, sortingClient})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}