package main

import (
	"log"

	"MyProjects/RentCar_gRPC/rentcar_api_gateway/clients"
	"MyProjects/RentCar_gRPC/rentcar_api_gateway/config"
	docs "MyProjects/RentCar_gRPC/rentcar_api_gateway/docs" // docs is generated by Swag CLI, you have to import it.
	"MyProjects/RentCar_gRPC/rentcar_api_gateway/handlers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// * @license.name  Apache 2.0
// * @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

func main() {

	cfg := config.Load()

	if cfg.Environment != "development" {
		gin.SetMode(gin.ReleaseMode)
	}

	docs.SwaggerInfo.Title = cfg.App
	docs.SwaggerInfo.Version = cfg.AppVersion

	r := gin.New()

	if cfg.Environment != "production" {
		r.Use(gin.Logger(), gin.Recovery())
	}

	//r.Use(MyCORSMiddleware())   //! * Agar opshi endpoitlarga qoyish kerak bolsa shunday ishlatiladi

	// r.GET("/ping", MyCORSMiddleware(), func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{		//! Yoki specific endpointga ham qoshishimiz mumkun
	// 		"message": "pong",
	// 	})
	// })

	grpcClients, err := clients.NewGrpcClients(cfg)
	if err != nil {
		panic(err)
	}

	defer grpcClients.Close()

	h := handlers.NewHandler(cfg, grpcClients)

	v1 := r.Group("/v1")
	{
		v1.Use(MyCORSMiddleware())
		v1.POST("/login", h.Login)

		v1.POST("/car", h.AuthMiddleware("*"), h.CreateCar)
		v1.GET("/car/:id", h.AuthMiddleware("*"), h.GetCarById)
		v1.GET("/car", h.AuthMiddleware("*"), h.GetCarList)
		v1.PUT("/car", h.AuthMiddleware("*"), h.UpdateCar)
		v1.DELETE("/car/:id", h.AuthMiddleware("ADMIN"), h.DeleteCar)

		v1.POST("brand", h.AuthMiddleware("*"), h.CreateBrand)
		v1.GET("/brand/:id", h.AuthMiddleware("*"), h.GetBrandById)
		v1.GET("/brand", h.AuthMiddleware("*"), h.GetBrandList)
		v1.PUT("/brand", h.AuthMiddleware("*"), h.UpdateBrand)
		v1.DELETE("/brand/:id", h.AuthMiddleware("ADMIN"), h.DeleteBrand)

		v1.POST("/rental", h.AuthMiddleware("*"), h.CreateRental)
		v1.GET("/rental/:id", h.AuthMiddleware("*"), h.GetRentalByID)
		v1.GET("/rental", h.AuthMiddleware("*"), h.GetRentalList)
		v1.PUT("/rental", h.AuthMiddleware("*"), h.UpdateRental)
		v1.DELETE("/rental/:id", h.AuthMiddleware("*"), h.DeleteRental)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(cfg.HTTPPort)
}

func MyCORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println("MyCORSMiddleware...")
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		ctx.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
		ctx.Header("Access-Control-Allow-HEADERS", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-TOKEN, Brandization, accept, origin, Cache-Control, X-Requested-With")
		ctx.Header("Access-Control-Max-Age", "3600")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}
		ctx.Next()
	}
}
