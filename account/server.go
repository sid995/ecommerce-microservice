package account

import (
	"context"
	"fmt"
	"net"

	pb "github.com/sid995/ecommerce-microservice/account/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	service Service
	pb.UnimplementedAccountServiceServer
}

func NewGRPCServer(service Service, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	serv := grpc.NewServer()
	pb.RegisterAccountServiceServer(serv, &grpcServer{
		service:                           service,
		UnimplementedAccountServiceServer: pb.UnimplementedAccountServiceServer{},
	})
	reflection.Register(serv)
	return serv.Serve(lis)
}

func (s *grpcServer) PostAccount(ctx context.Context, r *pb.PostAccountRequest) (*pb.PostAccountResponse, error) {
	a, err := s.service.PostAccount(ctx, r.Name)
	if err != nil {
		return nil, err
	}

	return &pb.PostAccountResponse{
		Account: &pb.Account{
			Id:   a.ID,
			Name: a.Name,
		},
	}, nil
}

func (s *grpcServer) GetAccount(ctx context.Context, r *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {
	a, err := s.service.GetAccount(ctx, r.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetAccountResponse{
		Account: &pb.Account{
			Id:   a.ID,
			Name: a.Name,
		},
	}, nil
}

func (s *grpcServer) GetAccounts(ctx context.Context, r *pb.GetAccountsRequest) (*pb.GetAccountsResponse, error) {
	accounts, err := s.service.GetAccounts(ctx, r.Skip, r.Take)
	if err != nil {
		return nil, err
	}

	var pbAccounts []*pb.Account
	for _, a := range accounts {
		pbAccounts = append(pbAccounts, &pb.Account{
			Id:   a.ID,
			Name: a.Name,
		})
	}

	return &pb.GetAccountsResponse{
		Accounts: pbAccounts,
	}, nil
}
