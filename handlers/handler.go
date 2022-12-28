package handlers

import (
	"MyProjects/RentCar_gRPC/rentcar_api_gateway/clients"
	"MyProjects/RentCar_gRPC/rentcar_api_gateway/config"
)

type handler struct {
	cfg 	config.Config
	grpcClients *clients.GrpcClients
}

func NewHandler(cfg config.Config, grpcClients *clients.GrpcClients) handler {
	return handler{
		cfg: cfg,
		grpcClients: grpcClients,
	}
}
