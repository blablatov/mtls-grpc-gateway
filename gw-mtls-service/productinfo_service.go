package main

import (
	"context"

	pb "github.com/blablatov/mtls-grpc-gateway/gw-mtls-proto"
	"github.com/gofrs/uuid"
	wrapper "github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Сервер используется для реализации product_info
type server struct {
	productMap map[string]*pb.Product
}

// AddProduct реализует ecommerce. AddProduct, добавить товар
func (s *server) AddProduct(ctx context.Context, in *pb.Product) (*wrapper.StringValue, error) {
	out, err := uuid.NewV4()
	if err != nil {
		return nil, status.Errorf(codes.Internal, " %v\nError while generating Product ID", err)
	}
	in.Id = out.String()
	if s.productMap == nil {
		s.productMap = make(map[string]*pb.Product)
	}
	s.productMap[in.Id] = in
	return &wrapper.StringValue{Value: in.Id}, status.New(codes.OK, "").Err()
}

// GetProduct реализует ecommerce. GetProduct получить товар
func (s *server) GetProduct(ctx context.Context, in *wrapper.StringValue) (*pb.Product, error) {
	value, exists := s.productMap[in.Value]
	if exists {
		return value, status.New(codes.OK, "").Err()
	}
	return nil, status.Errorf(codes.NotFound, "%v Product does not exist.", in.Value)
}
