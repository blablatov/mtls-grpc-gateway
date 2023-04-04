package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	rs "grpc-redis"

	pb "github.com/blablatov/mtls-grpc-gateway/gw-mtls-proto"
	"github.com/gofrs/uuid"
	wrapper "github.com/golang/protobuf/ptypes/wrappers"
	epb "google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Implements server.
// Сервер используется для реализации productinfo_service
type server struct {
	productMap map[string]*pb.Product
}

// Method add of product. Метод сервера AddProduct, добавить товар
func (s *server) AddProduct(ctx context.Context, in *pb.Product) (*wrapper.StringValue, error) {
	// Bad request, generate and sends of error to client.
	// Некорректный запрос. Сгенерировать и отправить клиенту ошибку.
	if in.Name == "-1" {
		log.Printf("Order ID is invalid! -> Received Order Name %s", in.Id)
		// Creates state with code of error. Создаем состояние с кодом ошибки InvalidArgument.
		errorStatus := status.New(codes.InvalidArgument, "Invalid information received")
		// Describes type of error. Описываем тип ошибки BadRequest_FieldViolation
		ds, err := errorStatus.WithDetails(
			&epb.BadRequest_FieldViolation{
				Field:       "Name",
				Description: fmt.Sprintf("Order Name received is not valid %s : %s", in.Id, in.Description),
			},
		)
		if err != nil {
			return nil, errorStatus.Err()
		}
		return nil, ds.Err()
	}
	out, err := uuid.NewV4()
	if err != nil {
		return nil, status.Errorf(codes.Internal, " %v\nError while generating Product ID", err)
	}
	in.Id = out.String()
	if s.productMap == nil {
		s.productMap = make(map[string]*pb.Product)
	}
	s.productMap[in.Id] = in

	/*dtr := rs.RedisData{
		Name:        in.Name,
		Description: in.Description,
	}*/

	chs := make(chan string, 3)
	chb := make(chan string, 1)

	var wg sync.WaitGroup
	wg.Add(1) // Counter of goroutines. Значение счетчика.

	go rs.RedisDsn(in.Name, in.Description, wg, chs, chb)

	go func() {
		wg.Wait()
		close(chs)
		close(chb)
	}()

	return &wrapper.StringValue{Value: in.Id}, status.New(codes.OK, <-chb).Err()
}

// Method get of product. Метод сервера GetProduct получить товар
func (s *server) GetProduct(ctx context.Context, in *wrapper.StringValue) (*pb.Product, error) {
	value, exists := s.productMap[in.Value]
	if exists {
		return value, status.New(codes.OK, "").Err()
	}
	return nil, status.Errorf(codes.NotFound, "%v Product does not exist.", in.Value)
}
