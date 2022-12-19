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
// @Router /posts [post]
// @Summary Create a post
// @Description Create a post
// @Tags post
// @Accept json
// @Produce json
// @Param post body models.CreatePostRequest true "post"
// @Success 201 {object} models.Post
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CreatePost(c *gin.Context) {
	var (
		req models.CreatePostRequest
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	payload, err := h.GetAuthPayload(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp, err := h.grpcClient.PostService().Create(context.Background(), &pbp.Post{
		Title:       req.Title,
		Description: req.Description,
		ImageUrl:    req.ImageUrl,
		UserId:      payload.UserId,
		CategoryId:  req.CategoryID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	post := parsePostModel(resp)
	c.JSON(http.StatusCreated, post)
}

func parsePostModel(post *pbp.Post) models.Post {
	return models.Post{
		ID:          post.Id,
		Title:       post.Title,
		Description: post.Description,
		ImageUrl:    post.ImageUrl,
		UserID:      post.UserId,
		CategoryID:  post.CategoryId,
		CreatedAt:   post.CreatedAt,
		ViewsCount:  post.ViewsCount,
	}
}

// @Security ApiKeyAuth
// @Router /posts{id} [put]
// @Summary Update a post
// @Description Update a post
// @Tags post
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param post body models.CreatePostRequest true "post"
// @Success 200 {object} models.Post
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) UpdatePost(c *gin.Context) {
	var (
		req models.UpdatePostRequest
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	payload, err := h.GetAuthPayload(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	post, err := h.grpcClient.PostService().Update(context.Background(), &pbp.Post{
		Id:          int64(id),
		Title:       req.Title,
		Description: req.Description,
		ImageUrl:    req.ImageUrl,
		UserId:      payload.UserId,
		CategoryId:  req.CategoryID,
	})
	if err != nil {
		h.logger.WithError(err).Error("failed to update post")
		if s, _ := status.FromError(err); s.Code() == codes.NotFound {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, parsePostModel(post))
}

// @Router /posts/{id} [get]
// @Summary Get post by id
// @Description Get post by id
// @Tags post
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.Post
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetPost(c *gin.Context) {
	h.logger.Info("get post")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	resp, err := h.grpcClient.PostService().Get(context.Background(), &pbp.GetPostRequest{Id: int64(id)})
	if err != nil {
		h.logger.WithError(err).Error("failed to get post")
		if s, _ := status.FromError(err); s.Code() == codes.NotFound {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, parsePostModel(resp))
}

// @Router /posts [get]
// @Summary Get all posts
// @Description Get all posts
// @Tags post
// @Accept json
// @Produce json
// @Param filter query models.GetAllParams false "Filter"
// @Success 200 {object} models.GetAllCategoriesResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetAllPosts(c *gin.Context) {
	req, err := validateGetAllParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	result, err := h.grpcClient.PostService().GetAll(context.Background(), &pbp.GetAllPostsRequest{
		Page:   req.Page,
		Limit:  req.Limit,
		Search: req.Search,
	})
	if err != nil {
		h.logger.WithError(err).Error("failed to get all posts")
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, getPostsResponse(result))
}

func getPostsResponse(data *pbp.GetAllPostsResponse) *models.GetAllPostsResponse {
	response := models.GetAllPostsResponse{
		Posts: make([]*models.Post, 0),
		Count: data.Count,
	}

	for _, post := range data.Posts {
		u := parsePostModel(post)
		response.Posts = append(response.Posts, &u)
	}

	return &response
}

// @Security ApiKeyAuth
// @Router /posts/{id} [delete]
// @Summary Delete post
// @Description Delete post
// @Tags post
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.ResponseOK
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) DeletePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	_, err = h.grpcClient.PostService().Delete(context.Background(), &pbp.GetPostRequest{Id: int64(id)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, models.ResponseOK{
		Message: "success",
	})
}
