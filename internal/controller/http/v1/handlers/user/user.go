package user

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
	srv userService
}

func NewHandler(srv userService) *Handler {
	return &Handler{
		srv: srv,
	}
}

func (h Handler) GetUser(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	userID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		resp.ErrorResponse(c, resp.MapError(entity.ErrBadIDFormat))
		return
	}

	entity, err := h.srv.GetUserByID(ctx, userID)

	if err != nil {
		resp.ErrorResponse(c, resp.MapError(err))
		return
	}

	user := &resp.User{}
	user.FillFromEntity(entity)

	c.JSON(http.StatusOK, user)
}

func (h Handler) RegisterUser(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	request := &req.UserRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		resp.ErrorResponse(c, resp.ErrorResponseData{
			Code:   resp.CodeInvalidJSON,
			Status: http.StatusBadRequest,
		})
		return
	}

	userID, err := h.srv.RegisterUser(ctx, request.ToEntity())

	if err != nil {
		resp.ErrorResponse(c, resp.MapError(err))
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"id": userID,
	})
}

func (h Handler) UpdateUser(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	userID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		resp.ErrorResponse(c, resp.MapError(entity.ErrBadIDFormat))
		return
	}

	request := &req.UserRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		resp.ErrorResponse(c, resp.ErrorResponseData{
			Code: resp.CodeInvalidJSON,
		})
		return
	}

	err = h.srv.UpdateUser(ctx, userID, request.ToEntity())

	if err != nil {
		resp.ErrorResponse(c, resp.MapError(err))
		return
	}

	c.Status(http.StatusOK)
}

func (h Handler) DeleteUser(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	userID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		resp.ErrorResponse(c, resp.MapError(entity.ErrBadIDFormat))
		return
	}

	err = h.srv.DeleteUser(ctx, userID)

	if err != nil {
		resp.ErrorResponse(c, resp.MapError(err))
		return
	}

	c.Status(http.StatusOK)
}
