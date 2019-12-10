package main

import (
	"context"
	"fmt"
	"github.com/makubit/grpc-vs-rest-service/grpc-sorting-service/sortLib"
	"os"

	s "github.com/makubit/grpc-vs-rest-service/grpc-sorting-service/proto/sortingService"
	"github.com/micro/go-micro"
)

func (s *service) Sort(ctx context.Context, req *s.SortRequest, res *s.Response) (error) {
	sorted, _ := sortLib.QuickSort(req.TableToSort)

	res.SortedTable = sorted
	res.Sorted = true
	return nil
}

type service struct {
}

func main() {
	srv := micro.NewService(
		micro.Name(os.Getenv("APP_NAME")),
	)

	srv.Init()
	s.RegisterSortingServiceHandler(srv.Server(), &service{})

	if err := srv.Run(); err != nil {
	fmt.Println(err)
	}
}

