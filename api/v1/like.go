package v1

import (
	"context"
	"github.com/MuhammadyusufAdhamov/medium_api_gateway/api/models"
	pbp "github.com/MuhammadyusufAdhamov/medium_api_gateway/genproto/post_service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Security ApiKeyAuth
// @Router /likes [post]
// @Summary Create Or Update like
// @Description Create Or Update like
// @Tags like
// @Accept json
// @Produce json
// @Param like body models.CreateOrUpdateLikeRequest true "like"
// @Success 201 {object} models.Like
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CreateOrUpdateLike(ctx *gin.Context) {
	var (
		req models.CreateOrUpdateLikeRequest
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	payload, err := h.GetAuthPayload(ctx)
	if err != nil {
		h.logger.WithError(err).Error("failed to get auth payload")
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp, err := h.grpcClient.LikeService().CreateOrUpdate(context.Background(), &pbp.Like{
		Status: req.Status,
		PostId: req.PostID,
		UserId: payload.UserId,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, models.Like{
		ID:     resp.Id,
		PostID: resp.PostId,
		UserID: resp.UserId,
		Status: resp.Status,
	})
}

// @Security ApiKeyAuth
// @Router /likes [get]
// @Summary Get like
// @Description Get like
// @Tags like
// @Accept json
// @Produce json
// @Param id path int true "post_id"
// @Success 201 {object} models.Like
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetLike(ctx *gin.Context) {
	postID, err := strconv.Atoi(ctx.Param("post_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	payload, err := h.GetAuthPayload(ctx)
	if err != nil {
		h.logger.WithError(err).Error("failed to get auth payload")
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp, err := h.grpcClient.LikeService().Get(context.Background(), &pbp.GetLikeRequest{
		UserId: payload.UserId,
		PostId: int64(postID),
	})
	if err != nil {
		h.logger.WithError(err).Error("failed to get like")
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, models.Like{
		ID:     resp.Id,
		PostID: resp.PostId,
		UserID: resp.UserId,
		Status: resp.Status,
	})
}

// @Router /likes/get-likes-and-dislikes [get]
// @Summary Get likes and dislike count
// @Description Get likes and dislikes count
// @Tags like
// @Accept json
// @Produce json
// @Param id path int true "post_id"
// @Success 201 {object} models.LikesAndDislikesCount
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetLikesAndDislikesCount(ctx *gin.Context) {
	h.logger.Info(ctx.Param("id"))
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		h.logger.WithError(err).Error("failed to parse id or bad request %v", id)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	count, err := h.grpcClient.LikeService().GetLikeDislikeCount(context.Background(), &pbp.GetLikeRequest{
		PostId: id,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, models.LikesAndDislikesCount{
		Likes:    count.Likes,
		Dislikes: int64(count.Dislikes),
	})
}
