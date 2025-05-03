package product

import (
	"context"
	"net/http"

	req "local/order-service/internal/controller/http/v1/request"
	resp "local/order-service/internal/controller/http/v1/response"
	"local/order-service/internal/entity"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	srv productService
}

func NewHandler(srv productService) *Handler {
	return &Handler{
		srv: srv,
	}
}

func (h Handler) GetProduct(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	productID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		resp.ErrorResponse(c, resp.MapError(entity.ErrBadIDFormat))
		return
	}

	entity, err := h.srv.GetProductByID(ctx, productID)

	if err != nil {
		resp.ErrorResponse(c, resp.MapError(err))
		return
	}

	product := &resp.Product{}
	product.FillFromEntity(entity)

	c.JSON(http.StatusOK, product)
}

func (h Handler) AddProduct(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	request := &req.ProductRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		resp.ErrorResponse(c, resp.ErrorResponseData{
			Code: resp.CodeInvalidJSON,
		})
		return
	}

	productID, err := h.srv.CreateProduct(ctx, request.ToEntity())

	if err != nil {
		resp.ErrorResponse(c, resp.MapError(err))
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"id": productID,
	})
}

func (h Handler) UpdateProduct(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	productID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		resp.ErrorResponse(c, resp.MapError(entity.ErrBadIDFormat))
		return
	}

	request := &req.ProductRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		resp.ErrorResponse(c, resp.ErrorResponseData{
			Code: resp.CodeInvalidJSON,
		})
		return
	}

	err = h.srv.UpdateProduct(ctx, productID, request.ToEntity())

	if err != nil {
		resp.ErrorResponse(c, resp.MapError(err))
		return
	}

	c.Status(http.StatusOK)
}

func (h Handler) DeleteProduct(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	productID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		resp.ErrorResponse(c, resp.ErrorResponseData{
			Code:    resp.CodeBadParamValue,
			Details: "incorrect product id",
		})
		return
	}

	err = h.srv.DeleteProduct(ctx, productID)

	if err != nil {
		resp.ErrorResponse(c, resp.MapError(err))
		return
	}

	c.Status(http.StatusOK)
}
