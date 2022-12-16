package v1

import (
	"context"
	"github.com/MuhammadyusufAdhamov/medium_api_gateway/api/models"
	pbp "github.com/MuhammadyusufAdhamov/medium_api_gateway/genproto/post_service"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"strconv"
)

// @Security ApiKeyAuth
// @Router /likes [like]
// @Summary Create a like
// @Description Create a like
// @Tags like
// @Accept json
// @Produce json
// @Param like body models.CreateLikeRequest true "like"
// @Success 201 {object} models.Like
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CreateLike(c *gin.Context) {
	var (
		req models.CreateLikeRequest
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	resp, err := h.grpcClient.LikeService().Create(context.Background(), &pbp.Like{
		PostId: req.PostID,
		UserId: req.UserID,
		Status: req.Status,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	like := parseLikeModel(resp)
	c.JSON(http.StatusCreated, like)
}

func parseLikeModel(like *pbp.Like) models.Like {
	return models.Like{
		ID:     like.Id,
		PostID: like.PostId,
		UserID: like.UserId,
		Status: like.Status,
	}
}

// @Router /likes/{id} [get]
// @Summary Get like by id
// @Description Get like by id
// @Tags like
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.Like
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetLike(c *gin.Context) {
	h.logger.Info("get like")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	resp, err := h.grpcClient.LikeService().Get(context.Background(), &pbp.GetLikeRequest{Id: int64(id)})
	if err != nil {
		h.logger.WithError(err).Error("failed to get like")
		if s, _ := status.FromError(err); s.Code() == codes.NotFound {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, parseLikeModel(resp))
}
