package order

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
	srv orderService
}

func NewHandler(srv orderService) *Handler {
	return &Handler{
		srv: srv,
	}
}

func (h Handler) GetOrder(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	orderID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		resp.ErrorResponse(c, resp.MapError(entity.ErrBadIDFormat))
		return
	}

	entity, err := h.srv.GetOrderByID(ctx, orderID)

	if err != nil {
		resp.ErrorResponse(c, resp.MapError(err))
		return
	}

	order := &resp.Order{}
	order.FillFromEntity(entity)

	c.JSON(http.StatusOK, order)
}

func (h Handler) CreateOrder(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	request := &req.OrderRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		resp.ErrorResponse(c, resp.ErrorResponseData{
			Code: resp.CodeInvalidJSON,
		})
		return
	}

	orderID, err := h.srv.CreateOrder(ctx, request.ToEntity())

	if err != nil {
		resp.ErrorResponse(c, resp.MapError(err))
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"id": orderID,
	})
}
