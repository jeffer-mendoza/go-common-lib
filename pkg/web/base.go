package web

import (
	"github.com/jeffer-mendoza/go-common-lib/pkg/domain"
	"github.com/jeffer-mendoza/go-common-lib/pkg/errors"

	"github.com/gin-gonic/gin"
)


func ReturnInfo(ctx *gin.Context, payload interface{}, err error, statusCode int) {
	var apiError errors.ApiError
	var response domain.ApiResponse
	if err != nil {
		apiError = errors.GetAPIError(err)
	} else {
		response = domain.ApiResponse{
			StatusCode:      statusCode,
			ServiceResponse: payload,
		}
	}
	respond(ctx, apiError, &response)
}

func respond(c *gin.Context, apiError errors.ApiError, apiResponse domain.HttpResponse) {
	var status int
	var response interface{}

	if apiError == nil {
		status = apiResponse.Status()
		response = apiResponse.Response()
	} else {
		status = apiError.Status()
		response = apiError
	}
	c.IndentedJSON(status, response)
}

