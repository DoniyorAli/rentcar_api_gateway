package handlers

import (
	"MyProjects/RentCar_gRPC/rentcar_api_gateway/clients"
	"MyProjects/RentCar_gRPC/rentcar_api_gateway/config"
)

type Handler struct {
	Cfg 	config.Config
	grpcClients *clients.GrpcClients
}

func NewHandler(cfg config.Config, grpcClients *clients.GrpcClients) Handler {
	return Handler{
		Cfg: cfg,
		grpcClients: grpcClients,
	}
}
