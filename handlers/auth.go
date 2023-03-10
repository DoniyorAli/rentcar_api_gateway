package handlers

import (
	"MyProjects/RentCar_gRPC/rentcar_api_gateway/models"
	"MyProjects/RentCar_gRPC/rentcar_api_gateway/protogen/authorization"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// //* AuthMyCORSMiddleware ...
func (h Handler) AuthMiddleware(userType string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		fmt.Println(token)
		hasAccesResponse, err := h.grpcClients.Auth.HasAccess(ctx.Request.Context(), &authorization.TokenRequest{
			Token: token,
		})
		fmt.Println(hasAccesResponse)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
				Error:err.Error(),
			})
			ctx.Abort()
			return
		}

		if !hasAccesResponse.HasAccess {
			ctx.JSON(http.StatusUnauthorized, "Unauthorized")
			ctx.Abort()
			return
		}

		if userType != "*" {
			if hasAccesResponse.User.UserType != userType {
				ctx.JSON(http.StatusUnauthorized, "Permission Denied")
				ctx.Abort()
			}
		}

		ctx.Set("auth_username", hasAccesResponse.User.Username)
		ctx.Set("auth_user_id", hasAccesResponse.User.Id)

		ctx.Next()
	}
}

// * ================== Login ======================
// Login godoc
// @Summary     Login
// @Description Login
// @Tags        authorization (login)
// @Accept      json
// @Produce     json
// @Param       login body     models.LoginModel true "Login body"
// @Success     201   {object} models.JSONResponse{data=models.TokenResponse}
// @Failure     400   {object} models.JSONErrorResponse
// @Router      /v1/login [post]
func (h *Handler) Login(ctx *gin.Context) {
	var body models.LoginModel
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	//TODO need do validation

	tokenResponse, err := h.grpcClients.Auth.Login(ctx.Request.Context(), &authorization.LoginRequest{
		Username: body.Username,
		Password: body.Password,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, models.JSONResponse{
		Message: "Brand successfully created",
		Data:    tokenResponse,
	})
}
