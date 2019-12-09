package main

import (
	gr "github.com/makubit/grpc-vs-rest-service/grpc-service/proto/grpcService"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"
	"testing"
	"time"
)

func BenchmarkGRPC(b *testing.B) {
	time.Sleep(5*time.Second)

	srv := micro.NewService(
		micro.Name("grpc.testing.cli"),
	)

	srv.Init()

	client := gr.NewGrpcServiceClient("grpc.service", srv.Client())

	for i := 0; i< b.N; i++ {
		grpcBench(client, b)
	}
}

func grpcBench(client gr.GrpcServiceClient, b *testing.B) {
	_, _ = client.GetFromSortingService(context.Background(), &gr.SortRequest{
		TableToSort: []int32{6, 5, 4, 3, 2, 1},
	})
}