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
// @Router /comments [post]
// @Summary Create a comment
// @Description Create a comment
// @Tags comment
// @Accept json
// @Produce json
// @Param comment body models.CreateCommentRequest true "comment"
// @Success 201 {object} models.Comment
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CreateComment(c *gin.Context) {
	var (
		req models.CreateCommentRequest
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	resp, err := h.grpcClient.CommentService().Create(context.Background(), &pbp.Comment{
		PostId:      req.PostID,
		UserId:      req.UserID,
		Description: req.Description,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	comment := parseCommentModel(resp)
	c.JSON(http.StatusCreated, comment)
}

func parseCommentModel(comment *pbp.Comment) models.Comment {
	return models.Comment{
		ID:          comment.Id,
		UserID:      comment.UserId,
		PostID:      comment.PostId,
		Description: comment.Description,
		CreatedAt:   comment.CreatedAt,
	}
}

// @Router /comments [get]
// @Summary Get all comments
// @Description Get all comments
// @Tags comment
// @Accept json
// @Produce json
// @Param filter query models.GetAllCommentsParams false "Filter"
// @Success 200 {object} models.GetAllCommentsResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetAllComments(c *gin.Context) {
	req, err := validateGetAllParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	result, err := h.grpcClient.CommentService().GetAll(context.Background(), &pbp.GetAllCommentsRequest{
		Page:   req.Page,
		Limit:  req.Limit,
		Search: req.Search,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, getCommentsResponse(result))
}

func getCommentsResponse(data *pbp.GetAllCommentsResponse) *models.GetAllCommentsResponse {
	response := models.GetAllCommentsResponse{
		Comments: make([]*models.Comment, 0),
		Count:    data.Count,
	}

	for _, comment := range data.Comments {
		u := parseCommentModel(comment)
		response.Comments = append(response.Comments, &u)
	}

	return &response
}

// @Router /comments/{id} [delete]
// @Summary Delete comment
// @Description Delete comment
// @Tags comment
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.ResponseOK
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) DeleteComment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	_, err = h.grpcClient.CommentService().Delete(context.Background(), &pbp.GetCommentRequest{Id: int64(id)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, models.ResponseOK{
		Message: "success",
	})
}
