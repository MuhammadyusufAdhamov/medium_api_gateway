package v1

import (
	"context"
	"github.com/MuhammadyusufAdhamov/medium_api_gateway/api/models"
	pbu "github.com/MuhammadyusufAdhamov/medium_api_gateway/genproto/user_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Router /auth/register [post]
// @Summary Register a user
// @Description Register a user
// @Tags auth
// @Accept json
// @Produce json
// @Param data body models.RegisterRequest true "Data"
// @Success 200 {object} models.ResponseOK
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) Register(c *gin.Context) {
	var (
		req models.RegisterRequest
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, _ := h.grpcClient.UserService().GetByEmail(context.Background(), &pbu.GetByEmailRequest{
		Email: req.Email,
	})
	if user != nil {
		c.JSON(http.StatusBadRequest, errorResponse(ErrEmailExists))
		return
	}

	_, err = h.grpcClient.AuthService().Register(context.Background(), &pbu.RegisterRequest{
		Email:     req.Email,
		Password:  req.Password,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, models.ResponseOK{
		Message: "success",
	})
}

// @Router /auth/verify [post]
// @Summary Verify user
// @Description Verify user
// @Tags auth
// @Accept json
// @Produce json
// @Param data body models.VerifyRequest true "Data"
// @Success 200 {object} models.AuthResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) Verify(c *gin.Context) {
	var (
		req models.VerifyRequest
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	result, err := h.grpcClient.AuthService().Verify(context.Background(), &pbu.VerifyRegisterRequest{
		Email: req.Email,
		Code:  req.Code,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, models.AuthResponse{
		ID:          result.Id,
		FirstName:   result.FirstName,
		LastName:    result.LastName,
		Email:       result.Email,
		Type:        result.Type,
		CreatedAt:   result.CreatedAt,
		AccessToken: result.AccessToken,
	})
}

// @Router /auth/login [post]
// @Summary Login user
// @Description Login user
// @Tags auth
// @Accept json
// @Produce json
// @Param data body models.LoginRequest true "Data"
// @Success 200 {object} models.AuthResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) Login(c *gin.Context) {
	var (
		req models.LoginRequest
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	result, err := h.grpcClient.AuthService().Login(context.Background(), &pbu.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, models.AuthResponse{
		ID:          result.Id,
		FirstName:   result.FirstName,
		LastName:    result.LastName,
		Email:       result.Email,
		Type:        result.Type,
		CreatedAt:   result.CreatedAt,
		AccessToken: result.AccessToken,
	})
}

// @Router /auth/verify-forgot-password [post]
// @Summary Verify forgot password
// @Description Verify forgot password
// @Tags auth
// @Accept json
// @Produce json
// @Param data body models.VerifyRequest true "Data"
// @Success 200 {object} models.AuthResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) VerifyForgotPassword(c *gin.Context) {
	var (
		req models.VerifyRequest
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	_, err = h.grpcClient.AuthService().ForgotPassword(context.Background(), &pbu.ForgotPasswordRequest{
		Email: req.Email,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, models.ResponseOK{
		Message: "Validation code has been sent",
	})
}
