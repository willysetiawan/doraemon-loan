package response

import (
	"net/http"
	"process-loan/lib/env"
	"strconv"

	"github.com/gin-gonic/gin"
)

type (
	Response struct {
		ResponseCode    string      `json:"response_code"`
		ResponseMessage string      `json:"response_message"`
		Data            interface{} `json:"data"`
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
		DebugParam string `json:"debug_param,omitempty"`
		TraceID    string `example:"97125121-ea32-4ee0-8706-5b7375e83e94" json:"trace_id,omitempty"`
	}
	MetaErrMock struct {
		DebugParam string `example:"Request Time Out" json:"debug_param"`
		TraceID    string `example:"97125121-ea32-4ee0-8706-5b7375e83e94" json:"trace_id"`
	}
)

func JsonGen(c *gin.Context, res Response) {
	code := http.StatusNotImplemented
	code, _ = strconv.Atoi(res.ResponseCode)
	// switch res.ResponseCode {
	// case constant.CODE_SUCCESS:
	// 	code = http.StatusOK
	// 	res.ResponseMessage = constant.CODE_SUCCESS_MSG
	// case constant.CODE_FAILED:
	// 	code = http.StatusBadRequest
	// 	res.ResponseMessage = constant.CODE_FAILED_MSG
	// case constant.CODE_INTERNAL_SERVER_ERR:
	// 	code = http.StatusInternalServerError
	// 	res.ResponseMessage = constant.CODE_INTERNAL_SERVER_ERR_MSG
	// case constant.CODE_PENDING:
	// 	code = http.StatusOK
	// 	res.ResponseMessage = constant.CODE_PENDING_MSG
	// }
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
