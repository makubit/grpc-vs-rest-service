package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	//pb "github.com/makubit/grpc-vs-rest-service/grpc-service/proto/consignment"
	sCli "github.com/makubit/grpc-vs-rest-service/grpc-sorting-service/proto/sortingService"
	//vesselProto "github.com/makubit/grpc-vs-rest-service/grpc-sorting-service/proto/vessel"
	gr "github.com/makubit/grpc-vs-rest-service/grpc-service/proto/grpcService"
	"github.com/micro/go-micro"
)

const (
	port=":50051"
)

type repository interface {
	//Create(*pb.Consignment) (*pb.Consignment, error)
	//GetAll() []*pb.Consignment
	//GetFromSortingService(request *s.)
}

type Repository struct {
	mu sync.RWMutex
	//consignments []*pb.Consignment
	sortedTableRequest []int32
}

/*func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.mu.Lock()
	updated := append(repo.consignments, consignment) //adding new request to list of consignments (it's not in proto)
	repo.consignments = updated
	repo.mu.Unlock()

	return consignment, nil
}*/

type service struct {
	repo repository
	//vesselClient vesselProto.VesselServiceClient
	sortingClient sCli.SortingServiceClient
}

/*func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) (error) {
	vesselResponse, err := s.vesselClient.FindAvalilable(context.Background(), &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity: int32(len(req.Containers)),
	})
	log.Printf("Found vessel: %s\n", vesselResponse.Vessel.Name)
	if err != nil {
		return err
	}

	req.VesselId = vesselResponse.Vessel.Id

	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	res.Created = true
	res.Consignment = consignment
	return nil
}*/

func (s *service) GetFromSortingService(ctx context.Context, req *gr.SortRequest, res *gr.Response) (error) {
	sortingResponse := s.sortingClient.Sort(context.Background(), &sCli.SortRequest{
		sorted: false,
		tableToSort: req.TableToSort,
	})
	log.Println("Table sorted: ", sortingResponse.tableToSort)

	res.Sorted = sortingResponse.sorted
	res.SortedTable = sortingResponse.tableToSort

	return nil
}

/*func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) (error) {
	consignments := s.repo.GetAll()
	res.Consignments = consignments
	return nil
}

func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}*/

func main() {
	repo := &Repository{}

	srv := micro.NewService(
		micro.Name("grpc.service"),
		)
	srv.Init()

	//vesselClient := vesselProto.NewVesselServiceClient("grpc.sorting.service", srv.Client())

	//pb.RegisterShippingServiceHandler(srv.Server(), &service{repo, vesselClient})

	sortingClient := sCli.NewSortingServiceClient("grpc.sorting.service", srv.Client())
	gr.RegisterGrpcServiceHandler(srv.Server(), &service{repo, sortingClient})
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}