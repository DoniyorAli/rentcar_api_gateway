package handlers

import (
	"net/http"
	"strconv"

	"MyProjects/RentCar_gRPC/rentcar_api_gateway/models"
	"MyProjects/RentCar_gRPC/rentcar_api_gateway/protogen/authorization"
	"MyProjects/RentCar_gRPC/rentcar_api_gateway/protogen/car"
	"MyProjects/RentCar_gRPC/rentcar_api_gateway/protogen/rental"

	"github.com/gin-gonic/gin"
)

// * ================== Create Rental ======================
// CreateRental godoc
// @Summary     Create rental
// @Description Create a new rental
// @Tags        rentals
// @Accept      json
// @Produce     json
// @Param       rental        body     models.CreateRentalModel true  "rental body"
// @Param       Authorization header   string                   false "Authorization"
// @Success     201           {object} models.JSONResponse{data=models.Rental}
// @Failure     400           {object} models.JSONErrorResponse
// @Router      /v1/rental [post]
func (h Handler) CreateRental(ctx *gin.Context) {
	var body models.CreateRentalModel
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	car, err := h.grpcClients.Car.GetCarByID(ctx.Request.Context(), &car.GetCarByIDRequest{
		Id: body.CarId,
	})
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.JSONErrorResponse{
			Error: "Car not found to create rental",
		})
		return
	}

	customer, err := h.grpcClients.Auth.GetUserByID(ctx.Request.Context(), &authorization.GetUserByIDRequest{
		Id: body.CustomerId,
	})

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, models.JSONErrorResponse{
			Error: "User is not registired to rental car",
		})
		return
	}

	rental, err := h.grpcClients.Rental.CreateRental(ctx.Request.Context(), &rental.CreateRentalRequest{
		CarId:      car.CarId,
		CustomerId: customer.Id,
		StartDate:  body.StartDate,
		EndDate:    body.EndDate,
		Payment:    body.Payment,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: "error in ---> CreateRental",
		})
		return
	}

	ctx.JSON(http.StatusCreated, models.JSONResponse{
		Message: "Rental successfully created",
		Data:    rental,
	})
}

// * ==================== GetRentalById ====================
// GetRentalById godoc
// @Summary     get rental by id
// @Description get a rental by id
// @Tags        rentals
// @Accept      json
// @Param       id            path   string true  "Rental ID"
// @Param       Authorization header string false "Authorization"
// @Produce     json
// @Success     200 {object} models.JSONResponse{data=models.PackedRentalModel}
// @Failure     404 {object} models.JSONErrorResponse
// @Router      /v1/rental/{id} [get]
func (h Handler) GetRentalByID(ctx *gin.Context) {
	idStr := ctx.Param("id")

	//TODO UUID validation

	rental, err := h.grpcClients.Rental.GetRentalByID(ctx.Request.Context(), &rental.GetRentalByIDRequest{
		RentalId: idStr,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: "error in ---> GetRentalByID",
		})
		return
	}

	car, err := h.grpcClients.Car.GetCarByID(ctx.Request.Context(), &car.GetCarByIDRequest{
		Id: rental.CarId,
	})
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.JSONErrorResponse{
			Error: "failed to get car",
		})
		return
	}
	rental.Car.CarId = car.CarId
	rental.Car.Model = car.Model
	rental.Car.Color = car.Color
	rental.Car.Year = car.Year
	rental.Car.Mileage = car.Mileage
	rental.Car.BrandId = car.Brand.BrandId
	rental.Car.CreatedAt = car.CreatedAt
	rental.Car.UpdatedAt = car.UpdatedAt

	customer, err := h.grpcClients.Auth.GetUserByID(ctx.Request.Context(), &authorization.GetUserByIDRequest{
		Id: rental.CustomerId,
	})
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, models.JSONErrorResponse{
			Error: "failed to get user",
		})
		return
	}
	rental.Customer.Id = customer.Id
	rental.Customer.Fname = customer.Fname
	rental.Customer.Lname = customer.Lname
	rental.Customer.Username = customer.Username
	rental.Customer.Password = customer.Password
	rental.Customer.UserType = customer.UserType
	rental.Customer.Address = customer.Address
	rental.Customer.Phone = customer.Phone
	rental.Customer.CreatedAt = customer.CreatedAt
	rental.Customer.UpdatedAt = customer.UpdatedAt

	ctx.JSON(http.StatusOK, models.JSONResponse{
		Message: "passed successfully",
		Data:    rental,
	})
}

// * ==================== GetRentalList ====================
// GetRentalList godoc
// @Summary     List rentals
// @Description get rentals
// @Tags        rentals
// @Accept      json
// @Produce     json
// @Param       offset        query    int    false "0"
// @Param       limit         query    int    false "10"
// @Param       search        query    string false "search"
// @Param       Authorization header   string false "Authorization"
// @Success     200           {object} models.JSONResponse{data=[]models.Rental}
// @Router      /v1/rental [get]
func (h Handler) GetRentalList(ctx *gin.Context) {

	offsetStr := ctx.DefaultQuery("offset", h.Cfg.DefaultOffset)
	limitStr := ctx.DefaultQuery("limit", h.Cfg.DefaultLimit)
	searchStr := ctx.DefaultQuery("search", "")

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: "error in offset",
		})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: "error in limit",
		})
		return
	}

	rentalList, err := h.grpcClients.Rental.GetRentalList(ctx.Request.Context(), &rental.GetRentalListRequest{
		Offset: int32(offset),
		Limit:  int32(limit),
		Search: searchStr,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: "error in ---> GetRentalList",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.JSONResponse{
		Message: "OK",
		Data:    rentalList,
	})
}

// * ==================== UpdateRental ====================
// UpdateRental godoc
// @Summary     Update rental
// @Description Update a new rental
// @Tags        rentals
// @Accept      json
// @Param       rental        body   models.UpdateRentalModel true  "updating rental"
// @Param       Authorization header string                   false "Authorization"
// @Produce     json
// @Success     200 {object} models.JSONResponse{data=models.Rental}
// @Failure     400 {object} models.JSONErrorResponse
// @Router      /v1/rental [put]
func (h Handler) UpdateRental(ctx *gin.Context) {
	var body models.UpdateRentalModel
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	car, err := h.grpcClients.Car.GetCarByID(ctx.Request.Context(), &car.GetCarByIDRequest{
		Id: body.CarId,
	})
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.JSONErrorResponse{
			Error: "Car not found to create rental",
		})
		return
	}

	customer, err := h.grpcClients.Auth.GetUserByID(ctx.Request.Context(), &authorization.GetUserByIDRequest{
		Id: body.CustomerId,
	})
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, models.JSONErrorResponse{
			Error: "User is not registired to rental car",
		})
		return
	}

	updated, err := h.grpcClients.Rental.UpdateRental(ctx.Request.Context(), &rental.UpdateRentalRequest{
		RentalId:   body.RentalId,
		CarId:      car.CarId,
		CustomerId: customer.Id,
		StartDate:  body.StartDate,
		EndDate:    body.EndDate,
		Payment:    body.Payment,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: "error in ---> UpdateRental",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.JSONResponse{
		Message: "Rental successfully updated",
		Data:    updated,
	})
}

// * ==================== DeleteRental ====================
// DeleteRental godoc
// @Summary     delete rental by id
// @Description delete rental by id
// @Tags        rentals
// @Accept      json
// @Param       id            path   string true  "Rental ID"
// @Param       Authorization header string false "Authorization"
// @Produce     json
// @Success     200 {object} models.JSONResponse{data=models.Rental}
// @Failure     400 {object} models.JSONErrorResponse
// @Router      /v1/rental/{id} [delete]
func (h Handler) DeleteRental(ctx *gin.Context) {
	idStr := ctx.Param("id")

	deleted, err := h.grpcClients.Rental.DeleteRental(ctx.Request.Context(), &rental.DeleteRentalRequest{
		RentalId: idStr,
	})
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.JSONErrorResponse{
			Error: "Rental have been already deleted!",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.JSONResponse{
		Message: "Rental suucessfully deleted",
		Data:    deleted,
	})
}
