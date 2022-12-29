package handlers

import (
	"net/http"
	"strconv"

	"MyProjects/RentCar_gRPC/rentcar_api_gateway/models"
	"MyProjects/RentCar_gRPC/rentcar_api_gateway/protogen/brand"

	"github.com/gin-gonic/gin"
)

// * ================== Create Brand ======================
// CreateBrand godoc
// @Summary     Create brand
// @Description Create a new brand
// @Tags        brands
// @Accept      json
// @Produce     json
// @Param       brand         body     models.CreateBrandModel true  "brand body"
// @Param       Authorization header   string                  false "Authorization"
// @Success     201           {object} models.JSONResponse{data=models.Brand}
// @Failure     400           {object} models.JSONErrorResponse
// @Router      /v1/brand [post]
func (h *Handler) CreateBrand(ctx *gin.Context) {
	var body models.CreateBrandModel
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	//TODO need do validation

	brand, err := h.grpcClients.Brand.CreateBrand(ctx.Request.Context(), &brand.CreateBrandRequest{
		Name:         body.Name,
		Country:      body.Country,
		Manufacturer: body.Manufacturer,
		AboutBrand:   body.AboutBrand,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: "This is not working ---> CreateBrand",
		})
		return
	}

	ctx.JSON(http.StatusCreated, models.JSONResponse{
		Message: "Brand successfully created",
		Data:    brand,
	})
}

// * ==================== GetBrandById ====================
// GetBrandById godoc
// @Summary     get brand by id
// @Description get a new brand
// @Tags        brands
// @Accept      json
// @Param       id            path   string true  "Brand ID"
// @Param       Authorization header string false "Authorization"
// @Produce     json
// @Success     200 {object} models.JSONResponse{data=models.Brand}
// @Failure     404 {object} models.JSONErrorResponse
// @Router      /v1/brand/{id} [get]
func (h *Handler) GetBrandById(ctx *gin.Context) {
	idStr := ctx.Param("id")

	//TODO UUID validation

	brand, err := h.grpcClients.Brand.GetBrandByID(ctx.Request.Context(), &brand.GetBrandByIDRequest{
		BrandId: idStr,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: "This is not working ---> GetBrandByID",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.JSONResponse{
		Message: "passed successfully",
		Data:    brand,
	})
}

// * ==================== GetBrandList ====================
// GetBrandList godoc
// @Summary     List brands
// @Description get brands
// @Tags        brands
// @Accept      json
// @Produce     json
// @Param       offset        query    int    false "0"
// @Param       limit         query    int    false "10"
// @Param       search        query    string false "search"
// @Param       Authorization header   string false "Authorization"
// @Success     200           {object} models.JSONResponse{data=[]models.Brand}
// @Router      /v1/brand [get]
func (h *Handler) GetBrandList(ctx *gin.Context) {

	offsetStr := ctx.DefaultQuery("offset", h.Cfg.DefaultOffset)
	limitStr := ctx.DefaultQuery("limit", h.Cfg.DefaultLimit)
	searchStr := ctx.DefaultQuery("search", "")

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: "error in offset",
		})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: "error in limit",
		})
		return
	}

	brandList, err := h.grpcClients.Brand.GetBrandList(ctx.Request.Context(), &brand.GetBrandListRequest{
		Offset: int32(offset),
		Limit:  int32(limit),
		Search: searchStr,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: "error in ---> h.grpcClients.Brand.GetBrandList",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.JSONResponse{
		Message: "OK",
		Data:    brandList,
	})
}

// * ==================== UpdateBrand ====================
// UpdateBrand godoc
// @Summary     Update brand
// @Description Update a new brand
// @Tags        brands
// @Accept      json
// @Param       brand         body   models.UpdateBrandModel true  "updating brand"
// @Param       Authorization header string                  false "Authorization"
// @Produce     json
// @Success     200 {object} models.JSONResponse{data=models.Brand}
// @Failure     400 {object} models.JSONErrorResponse
// @Router      /v1/brand [put]
func (h *Handler) UpdateBrand(ctx *gin.Context) {
	var body models.UpdateBrandModel
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	brand, err := h.grpcClients.Brand.UpdateBrand(ctx.Request.Context(), &brand.UpdateBrandRequest{
		Id: body.BrandId,
		Name:         body.Name,
		Country:      body.Country,
		Manufacturer: body.Manufacturer,
		AboutBrand:   body.AboutBrand,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: "error in updating",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.JSONResponse{
		Message: "Brand successfully updated",
		Data:    brand,
	})
}

// * ==================== DeleteBrand ====================
// DeleteBrand godoc
// @Summary     Delete brand
// @Description delete brand
// @Tags        brands
// @Accept      json
// @Param       id            path   string true  "Brand ID"
// @Param       Authorization header string false "Authorization"
// @Produce     json
// @Success     200 {object} models.JSONResponse{data=models.Brand}
// @Failure     400 {object} models.JSONErrorResponse
// @Router      /v1/brand/{id} [delete]
func (h *Handler) DeleteBrand(ctx *gin.Context) {
	idStr := ctx.Param("id")

	brand, err := h.grpcClients.Brand.DeleteBrand(ctx.Request.Context(), &brand.DeleteBrandRequest{
		BrandId: idStr,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: "Brand have been already deleted!",
		})
		return
	}

	ctx.JSON(http.StatusNotFound, models.JSONResponse{
		Message: "Brand suucessfully deleted",
		Data:    brand,
	})
}
