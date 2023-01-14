package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	App         string
	AppVersion  string
	Environment string // devlopment, staging, production

	HTTPPort string

	GRPCPort string

	DefaultOffset string
	DefaultLimit  string

	CarServiceGrpcHost string
	CarServiceGrpcPort string

	BrandServiceGrpcHost string
	BrandServiceGrpcPort string

	RentalServiceGrpcHost string
	RentalServiceGrpcPort string

	AuthorizationServiceGrpcHost string
	AuthorizationServiceGrpcPort string

}

// Load ...
func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.App = cast.ToString(getOrReturnDefaultValue("APP", "article"))
	config.AppVersion = cast.ToString(getOrReturnDefaultValue("APP_VERSION", "1.0.0"))
	config.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", "development"))

	config.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":6060"))
	config.GRPCPort = cast.ToString(getOrReturnDefaultValue("GRPC_PORT", ":6003"))

	config.DefaultOffset = cast.ToString(getOrReturnDefaultValue("DEFAULT_OFFSET", "0"))
	config.DefaultLimit = cast.ToString(getOrReturnDefaultValue("DEFAULT_LIMIT", "10"))

	config.CarServiceGrpcHost = cast.ToString(getOrReturnDefaultValue("CAR_SERVICE_GRPC_HOST", "localhost"))
	config.CarServiceGrpcPort = cast.ToString(getOrReturnDefaultValue("CAR_SERVICE_GRPC_PORT", ":6001"))

	config.BrandServiceGrpcHost = cast.ToString(getOrReturnDefaultValue("BRAND_SERVICE_GRPC_HOST", "localhost"))
	config.BrandServiceGrpcPort = cast.ToString(getOrReturnDefaultValue("BRAND_SERVICE_GRPC_PORT", ":6001"))

	config.AuthorizationServiceGrpcHost = cast.ToString(getOrReturnDefaultValue("AUTHORIZATION_SERVICE_GRPC_HOST", "localhost"))
	config.AuthorizationServiceGrpcPort = cast.ToString(getOrReturnDefaultValue("AUTHORIZATION_SERVICE_GRPC_PORT", ":6002"))

	config.RentalServiceGrpcHost = cast.ToString(getOrReturnDefaultValue("RENTAL_SERVICE_GRPC_HOST", "localhost"))
	config.RentalServiceGrpcPort = cast.ToString(getOrReturnDefaultValue("RENTAL_SERVICE_GRPC_PORT", ":6003"))

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)

	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
