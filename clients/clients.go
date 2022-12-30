package clients

import (
	"MyProjects/RentCar_gRPC/rentcar_api_gateway/config"
	"MyProjects/RentCar_gRPC/rentcar_api_gateway/protogen/authorization"
	"MyProjects/RentCar_gRPC/rentcar_api_gateway/protogen/brand"
	"MyProjects/RentCar_gRPC/rentcar_api_gateway/protogen/car"
	"MyProjects/RentCar_gRPC/rentcar_api_gateway/protogen/rental"

	"google.golang.org/grpc"
)

type GrpcClients struct {
	Car    car.CarServiceClient
	Brand  brand.BrandServiceClient
	Auth   authorization.AuthServiceClient
	Rental rental.RentalServiceClient
	conns  []*grpc.ClientConn
}

func NewGrpcClients(cfg config.Config) (*GrpcClients, error) {
	connectCar, err := grpc.Dial(cfg.CarServiceGrpcHost+cfg.CarServiceGrpcPort, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	car := car.NewCarServiceClient(connectCar)

	connectBrand, err := grpc.Dial(cfg.BrandServiceGrpcHost+cfg.BrandServiceGrpcPort, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	brand := brand.NewBrandServiceClient(connectBrand)

	connectAuth, err := grpc.Dial(cfg.AuthorizationServiceGrpcHost+cfg.AuthorizationServiceGrpcPort, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	auth := authorization.NewAuthServiceClient(connectAuth)

	connRental, err := grpc.Dial(cfg.RentalServiceGrpcHost+cfg.RentalServiceGrpcPort, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	rental := rental.NewRentalServiceClient(connRental)

	conns := make([]*grpc.ClientConn, 0)
	return &GrpcClients{
		Car:    car,
		Brand:  brand, //! nega brand qoshilmagan appendda
		Auth:   auth,
		Rental: rental,
		conns:  append(conns, connectCar, connectBrand, connectAuth, connRental),
	}, nil
}

func (c *GrpcClients) Close() {
	for _, v := range c.conns {
		v.Close()
	}
}
