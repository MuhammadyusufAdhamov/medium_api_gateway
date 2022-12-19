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
// @Router /categories [post]
// @Summary Create a category
// @Description Create a category
// @Tags category
// @Accept json
// @Produce json
// @Param category body models.CreateCategoryRequest true "Category"
// @Success 201 {object} models.Category
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CreateCategory(c *gin.Context) {
	var (
		req models.CreateCategoryRequest
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	resp, err := h.grpcClient.CategoryService().Create(context.Background(), &pbp.Category{
		Title: req.Title,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, models.Category{
		ID:        resp.Id,
		Title:     resp.Title,
		CreatedAt: resp.CreatedAt,
	})
}

// @Security ApiKeyAuth
// @Router /categories/{id} [put]
// @Summary Update a category
// @Description Update a category
// @Tags category
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param category body models.CreateCategoryRequest true "category"
// @Success 200 {object} models.Category
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) UpdateCategory(c *gin.Context) {
	var (
		req models.UpdateCategoryRequest
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	category, err := h.grpcClient.CategoryService().Update(context.Background(), &pbp.Category{
		Id:    int64(id),
		Title: req.Title,
	})
	if err != nil {
		h.logger.WithError(err).Error("failed to update category")
		if s, _ := status.FromError(err); s.Code() == codes.NotFound {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, parseCategoryModel(category))
}

// @Router /categories/{id} [get]
// @Summary Get category by id
// @Description Get category by id
// @Tags category
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.Category
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetCategory(c *gin.Context) {
	h.logger.Info("get category")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	resp, err := h.grpcClient.CategoryService().Get(context.Background(), &pbp.GetCategoryRequest{Id: int64(id)})
	if err != nil {
		h.logger.WithError(err).Error("failed to get category")
		if s, _ := status.FromError(err); s.Code() == codes.NotFound {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, parseCategoryModel(resp))
}

func parseCategoryModel(category *pbp.Category) models.Category {
	return models.Category{
		ID:        category.Id,
		Title:     category.Title,
		CreatedAt: category.CreatedAt,
	}
}

// @Router /categories [get]
// @Summary Get all categories
// @Description Get all categories
// @Tags category
// @Accept json
// @Produce json
// @Param filter query models.GetAllParams false "Filter"
// @Success 200 {object} models.GetAllCategoriesResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetAllCategories(c *gin.Context) {
	req, err := validateGetAllParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	result, err := h.grpcClient.CategoryService().GetAll(context.Background(), &pbp.GetAllCategoriesRequest{
		Page:   req.Page,
		Limit:  req.Limit,
		Search: req.Search,
	})
	if err != nil {
		h.logger.WithError(err).Error("failed to get all categories")
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, getCategoriesResponse(result))
}

func getCategoriesResponse(data *pbp.GetAllCategoriesResponse) *models.GetAllCategoriesResponse {
	response := models.GetAllCategoriesResponse{
		Categories: make([]*models.Category, 0),
		Count:      data.Count,
	}

	for _, category := range data.Categories {
		u := parseCategoryModel(category)
		response.Categories = append(response.Categories, &u)
	}

	return &response
}

// @Security ApiKeyAuth
// @Router /categories/{id} [delete]
// @Summary Delete category
// @Description Delete category
// @Tags category
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.ResponseOK
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) DeleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	_, err = h.grpcClient.CategoryService().Delete(context.Background(), &pbp.GetCategoryRequest{Id: int64(id)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, models.ResponseOK{
		Message: "success",
	})
}
