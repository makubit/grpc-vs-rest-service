package main

import (
	"context"
	//"errors"
	"fmt"
	"github.com/makubit/grpc-vs-rest-service/grpc-sorting-service/sortLib"

	//pb "github.com/makubit/grpc-vs-rest-service/grpc-sorting-service/proto/vessel"
	s "github.com/makubit/grpc-vs-rest-service/grpc-sorting-service/proto/sortingService"
	"github.com/micro/go-micro"
)

/*type repository interface {
	FindAvalilable(*pb.Specification) (*pb.Vessel, error)
	Sort(*s.SortRequest, *s.Response) (error)
}

type VesselRepository struct {
	vessels []*pb.Vessel
}

func (repo *VesselRepository) FindAvalilable(spec *pb.Specification) (*pb.Vessel, error) {
	for _, vessel := range repo.vessels {
		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
			return vessel, nil
		}
	}
	return nil, errors.New("No vessel found by that spec")
}*/

func (s *service) Sort(ctx context.Context, req *s.SortRequest, res *s.Response) (error) {
	sorted, _ := sortLib.QuickSort(req.TableToSort)

	res.SortedTable = sorted
	res.Sorted = true
	return nil
}

type service struct {
	//repo repository
}

/*func (s *service) FindAvalilable(ctx context.Context, req *pb.Specification, res *pb.Response) (error) {
	vessel, err := s.repo.FindAvalilable(req)
	if err != nil {
		return err
	}

	res.Vessel = vessel
	return nil
}*/

func main() {
	/*vessels := []*pb.Vessel {
		{Id: "vessel001", Name: "Boaty", MaxWeight: 200000, Capacity: 500},
	}*/

	//repo := &VesselRepository{vessels}

	srv := micro.NewService(
		micro.Name("grpc.sorting.service"),
	)

	srv.Init()

	//pb.RegisterVesselServiceHandler(srv.Server(), &service{})
	s.RegisterSortingServiceHandler(srv.Server(), &service{})

	if err := srv.Run(); err != nil {
	fmt.Println(err)
	}
}

