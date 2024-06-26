package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/strivesolutions/go-gin-framework/pkg/serialization"
	"github.com/strivesolutions/go-strive-utils/pkg/striveexceptions"
	"github.com/strivesolutions/logger-go/pkg/logging"
)

type ApiResponse struct {
	Data  interface{} `json:"data"`
	Error *ApiError   `json:"error"`
}

func (r ApiResponse) ToString() string {
	return fmt.Sprintf("Data: %T\nError: %T", r.Data, r.Error)
}

func (r ApiResponse) Serialize() (*string, *string) {
	return serialization.ToJson(r)
}

// Deprecated: Will be removed in future versions
func NewError(message, path string, statusCode int) *ApiError {
	result := ApiError{}
	result.Message = message
	result.Path = path
	result.Code = statusCode
	return &result
}

func NewErrorCode(message, path string, err error) *ApiError {
	result := ApiError{}
	result.Message = message
	result.Path = path

	code, codeErr := strconv.Atoi(fmt.Sprint(err))

	if codeErr != nil {
		code = 400
	}

	result.Code = code
	result.Detail = err.Error()
	return &result
}

func HandleError(c *gin.Context, err striveexceptions.Exception) {
	detail := err.Details
	if detail == "" {
		if err.FullError != nil {
			detail = err.FullError.Error()
		}
	}

	resp := ApiResponse{
		Error: &ApiError{
			Message: err.Message,
			Code:    err.Code,
			Path:    c.Request.RequestURI,
			Detail:  detail,
		},
	}

	c.JSON(err.StatusCode, resp)
	c.Abort()
}

func AbortBadRequest(c *gin.Context, err error) {
	logging.Error(fmt.Sprintf("%s: %s", c.Request.RequestURI, err))

	resp := ApiResponse{}

	resp.Error = NewErrorCode("Bad Request", c.Request.RequestURI, err)

	c.JSON(http.StatusBadRequest, resp)
	c.Abort()
}

func AbortUnauthorized(c *gin.Context) {
	resp := ApiResponse{
		Error: &ApiError{
			Message: "Unauthorized",
			Code:    401,
			Path:    c.Request.RequestURI,
		},
	}
	c.JSON(http.StatusUnauthorized, resp)
	c.Abort()
}

func AbortForbidden(c *gin.Context) {
	err := ApiError{
		Message: "Forbidden",
		Code:    403,
		Path:    c.Request.RequestURI,
	}
	resp := ApiResponse{Error: &err}

	c.JSON(http.StatusForbidden, resp)
	c.Abort()
}

func AbortInternalServerError(c *gin.Context, err error) {
	logging.Error("Internal Server Error: %s", err)

	apiError := ApiError{
		Message: "Internal Server Error",
		Code:    500,
		Path:    c.Request.RequestURI,
	}
	resp := ApiResponse{
		Error: &apiError,
	}

	c.JSON(http.StatusInternalServerError, resp)
	c.Abort()
}

func AbortNotFound(c *gin.Context) {
	resp := ApiError{
		Message: "Not Found",
		Code:    404,
		Path:    c.Request.RequestURI,
	}

	c.JSON(http.StatusNotFound, resp)
	c.Abort()
}

func OkResponse(c *gin.Context, data interface{}) {
	resp := ApiResponse{Data: data}
	c.JSON(http.StatusOK, resp)
}

func NoContentResponse(c *gin.Context) {
	c.Status(http.StatusNoContent)
}
