package response

import (
	"errors"
	"local/order-service/internal/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	errFieldRequired *entity.ErrFieldRequired
)

type ErrorCode string

const (
	CodeInvalidJSON           ErrorCode = "INVALID_JSON"
	CodeBadParamValue         ErrorCode = "BAD_PARAM_VALUE"
	CodeInternalError         ErrorCode = "INTERNAL_ERROR"
	CodeItemNotFound          ErrorCode = "ITEM_NOT_FOUND"
	CodeMissedRequredField    ErrorCode = "MISSED_REQUIRED_FIELD"
	CodeInsufficientConds     ErrorCode = "INSUFFICIENT_CONDITIONS"
	CodeInsufficientResources ErrorCode = "INSUFFICIENT_RESOURCES"
)

type ErrorResponseData struct {
	Code    ErrorCode `json:"code"`
	Details string    `json:"details,omitempty"`
	Status  int       `json:"-"`
}

func ErrorResponse(c *gin.Context, resp ErrorResponseData) {
	c.JSON(resp.Status, resp)
}

func MapError(err error) ErrorResponseData {
	switch {
	case errors.Is(err, entity.ErrBadIDFormat):
		return ErrorResponseData{
			Code:    CodeBadParamValue,
			Details: err.Error(),
			Status:  http.StatusBadRequest,
		}
	case errors.Is(err, entity.ErrUserNotExist):
		return ErrorResponseData{
			Code:    CodeInsufficientConds,
			Details: err.Error(),
			Status:  http.StatusBadRequest,
		}
	case errors.As(err, &errFieldRequired):
		return ErrorResponseData{
			Code:    CodeMissedRequredField,
			Details: err.Error(),
			Status:  http.StatusBadRequest,
		}
	case errors.Is(err, entity.ErrNotFound):
		return ErrorResponseData{
			Code:   CodeItemNotFound,
			Status: http.StatusNotFound,
		}
	case errors.Is(err, entity.ErrInsufficientAge):
		return ErrorResponseData{
			Code:    CodeInsufficientConds,
			Details: err.Error(),
			Status:  http.StatusBadRequest,
		}
	case errors.Is(err, entity.ErrPasswordTooShort):
		return ErrorResponseData{
			Code:    CodeInsufficientConds,
			Details: err.Error(),
			Status:  http.StatusBadRequest,
		}
	case errors.Is(err, entity.ErrNotEnoughProducts):
		return ErrorResponseData{
			Code:    CodeInsufficientResources,
			Details: err.Error(),
			Status:  http.StatusBadRequest,
		}
	default:
		return ErrorResponseData{
			Code:   CodeInternalError,
			Status: http.StatusInternalServerError,
		}
	}
}
