package response

import (
	"net/http"
	"process-loan/lib/env"
	"strconv"

	"github.com/gin-gonic/gin"
)

type (
	Response struct {
		ResponseCode    string      `json:"responseCode"`
		ResponseMessage string      `json:"responseMessage"`
		Data            interface{} `json:"data,omitempty"`
		Meta            Meta        `json:"meta,omitempty"`
	}
	ResponsePagination struct {
		ResponseCode    string      `json:"response_code"`
		ResponseMessage string      `json:"response_message"`
		Data            interface{} `json:"data"`
		Pagination      Pagination  `json:"pagination"`
		Meta            Meta        `json:"meta"`
	}
	Pagination struct {
		Page  int   `json:"page"`
		Limit int   `json:"limit"`
		Total int64 `json:"total"`
	}
	Meta struct {
		DebugParam string `json:"debugParam,omitempty"`
		TraceID    string `example:"97125121-ea32-4ee0-8706-5b7375e83e94" json:"traceId,omitempty"`
	}
)

func JsonGen(c *gin.Context, res Response) {
	code := http.StatusNotImplemented
	code, _ = strconv.Atoi(res.ResponseCode)
	res.ResponseCode = env.String("MainSetup.PrefixCode", "") + res.ResponseCode
	c.JSON(code, res)
}

func Json(c *gin.Context, code int, data interface{}) {
	c.JSON(code, Response{Data: data, ResponseMessage: "OK"})
}

func JsonPagination(c *gin.Context, code int, data interface{}, page, limit int, total int64) {
	c.JSON(code, Response{
		Data: ResponsePagination{
			Data: data,
			Pagination: Pagination{
				Page:  page,
				Limit: limit,
				Total: total,
			},
		},
		ResponseMessage: "OK",
	})
}
