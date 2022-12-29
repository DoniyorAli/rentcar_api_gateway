package handlers

import (
	"net/http"
	"strconv"

	"MyProjects/RentCar_gRPC/rentcar_api_gateway/models"
	"MyProjects/RentCar_gRPC/rentcar_api_gateway/protogen/car"

	"github.com/gin-gonic/gin"
)

// * ================== Create Car =========================
// CreateCar godoc
// @Summary     Create car
// @Description Create a new car
// @Tags        cars
// @Accept      json
// @Param       car           body   models.CreateCarModel true  "car body"
// @Param       Authorization header string                false "Authorization"
// @Produce     json
// @Success     201 {object} models.JSONResponse{data=models.PackedCarModel}
// @Failure     400 {object} models.JSONErrorResponse
// @Router      /v1/car [post]
func (h *Handler) CreateCar(ctx *gin.Context) {
	var body models.CreateCarModel
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	obj, err := h.grpcClients.Car.CreateCar(ctx.Request.Context(), &car.CreateCarRequest{
		Model: body.Model,
		Color: body.Color,
		CarType: body.CarType,
		Mileage: body.Mileage,
		Year: body.Year,
		Price: body.Price,
		BrandId: body.BrandId,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: "error in ---> CreateCar",
		})
		return
	}

	car, err := h.grpcClients.Car.GetCarByID(ctx.Request.Context(), &car.GetCarByIDRequest{
		Id: obj.CarId,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: "error in ---> GetAuthorByID",
		})
		return
	}

	ctx.JSON(http.StatusCreated, models.JSONResponse{
		Message: "Car successfully created",
		Data:    car,
	})
}

// * ==================== Get Car By Id ====================
// GetCarById godoc
// @Summary     get car by id
// @Description get a new car
// @Tags        cars
// @Accept      json
// @Param       id            path   string true  "Car ID"
// @Param       Authorization header string false "Authorization"
// @Produce     json
// @Success     200 {object} models.JSONResponse{data=models.PackedCarModel}
// @Failure     404 {object} models.JSONErrorResponse
// @Router      /v1/car/{id} [get]
func (h *Handler) GetCarById(ctx *gin.Context) {
	idStr := ctx.Param("id")

	//TODO UUID validation

	car, err := h.grpcClients.Car.GetCarByID(ctx.Request.Context(), &car.GetCarByIDRequest{
		Id: idStr,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: "error in ---> GetCarById",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.JSONResponse{
		Message: "GetCarById passed successfully",
		Data:    car,
	})
}

// * ==================== Get Car List ====================
// GetCarList godoc
// @Summary     List cars
// @Description get cars
// @Tags        cars
// @Accept      json
// @Produce     json
// @Param       offset        query    int    false "0"
// @Param       limit         query    int    false "10"
// @Param       search        query    string false "search"
// @Param       Authorization header   string false "Authorization"
// @Success     200           {object} models.JSONResponse{data=[]models.Car}
// @Router      /v1/car [get]
func (h *Handler) GetCarList(ctx *gin.Context) {
	offsetStr := ctx.DefaultQuery("offset", h.Cfg.DefaultOffset)
	limitStr := ctx.DefaultQuery("limit", h.Cfg.DefaultLimit)
	searchStr := ctx.DefaultQuery("search", "")

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: ("error in ---> offset"),
		})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: "error in --- limit",
		})
		return
	}

	carList, err := h.grpcClients.Car.GetCarList(ctx.Request.Context(), &car.GetCarListRequest{
		Offset: int32(offset),
		Limit:  int32(limit),
		Search: searchStr,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: "error in ---> GetCarList",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.JSONResponse{
		Message: "OK",
		Data:    carList,
	})
}

// * ==================== Update Car =======================
// UpdateCar godoc
// @Summary     Update car
// @Description Update a new car
// @Tags        cars
// @Accept      json
// @Param       car           body   models.UpdateCarModel true  "updating car"
// @Param       Authorization header string                false "Authorization"
// @Produce     json
// @Success     200 {object} models.JSONResponse{data=models.Car}
// @Failure     400 {object} models.JSONErrorResponse
// @Router      /v1/car [put]
func (h *Handler) UpdateCar(ctx *gin.Context) {
	var body models.UpdateCarModel
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	updated, err := h.grpcClients.Car.UpdateCar(ctx.Request.Context(), &car.UpdateCarRequest{
		Id:   body.CarId,
		Model: body.Model,
		Color: body.Color,
		CarType: body.CarType,
		Mileage: body.Mileage,
		Year: body.Year,
		Price: body.Price,
		BrandId: body.BrandId,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.JSONResponse{
		Message: "Car successfully updated",
		Data:    updated,
	})
}

// * ==================== Delete Car =======================
// DeleteCar godoc
// @Summary     Delete car
// @Description delete car
// @Tags        cars
// @Accept      json
// @Param       id            path   string true  "Car ID"
// @Param       Authorization header string false "Authorization"
// @Produce     json
// @Success     200 {object} models.JSONResponse{data=models.Car}
// @Failure     400 {object} models.JSONErrorResponse
// @Router      /v1/car/{id} [delete]
func (h *Handler) DeleteCar(ctx *gin.Context) {
	idStr := ctx.Param("id")

	car, err := h.grpcClients.Car.DeleteCar(ctx.Request.Context(), &car.DeleteCarRequest{
		Id: idStr,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: "error in DeleteCar ---> GetCarByID",
		})
		return
	}

	ctx.JSON(http.StatusNotFound, models.JSONResponse{
		Message: "Car successfully deleted",
		Data:    car,
	})
}
