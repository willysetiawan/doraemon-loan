package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"process-loan/lib/constant"
	"process-loan/lib/log"
	"process-loan/lib/response"
	helper "process-loan/lib/validate"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ValidateSymetricSignature() gin.HandlerFunc {
	return func(c *gin.Context) {
		logName := "Middleware - ValidateSymetricSignature"
		traceID := c.GetHeader("Trace-ID")
		if traceID == "" {
			traceID = uuid.NewString()
		}

		res := response.Response{
			ResponseCode:    strconv.Itoa(http.StatusBadRequest),
			ResponseMessage: "Bad Request",
		}

		code := http.StatusBadRequest
		urlString := c.Request.URL.String()
		// clientID := c.GetHeader("X-CLIENT-KEY")
		signature := c.GetHeader("X-SIGNATURE")
		timeStamp := c.GetHeader("X-TIMESTAMP")

		if signature == "" {
			res.ResponseMessage = constant.MESSAGE_MISSING_MANDATORY_FIELD + " X-SIGNATURE"
			c.Abort()
			c.JSON(code, res)
			return
		}
		if timeStamp == "" {
			res.ResponseMessage = constant.MESSAGE_MISSING_MANDATORY_FIELD + " X-TIMESTAMP"
			c.Abort()
			c.JSON(code, res)
			return
		}

		req, _ := ioutil.ReadAll(c.Request.Body)
		dst := &bytes.Buffer{}
		json.Compact(dst, req)
		fmt.Println("data: ", dst.String())

		dataMessage := helper.StringToSignHMAC512(fmt.Sprintf("%v", c.Request.Method), urlString, fmt.Sprintf("%v", timeStamp), dst.Bytes())
		if _, err := helper.ConvertStringToDate(timeStamp, time.RFC3339); err != nil {
			log.LogData(traceID, logName, "ConvertStringToDate", constant.LEVEL_LOG_ERROR, err.Error())
			res.ResponseCode = fmt.Sprintf("%v", http.StatusBadRequest)
			res.ResponseMessage = fmt.Sprintf("Invalid Field Format {%v}", "X-TIMESTAMP")
			responseCode, _ := strconv.Atoi(res.ResponseCode[:3])
			c.Abort()
			c.JSON(responseCode, res)
			return
		}

		fmt.Println("data message", dataMessage)

		if !helper.ValidateSignatureHMAC512(os.Getenv("SECRET_KEY"), []byte(dataMessage), signature) {
			res.ResponseCode = fmt.Sprintf("%v", http.StatusUnauthorized)
			res.ResponseMessage = "Unauthorized Signature"
			responseCode, _ := strconv.Atoi(res.ResponseCode[:3])
			c.Abort()
			c.JSON(responseCode, res)
			return
		}

		c.Set("client_secret_key", os.Getenv("SECRET_KEY"))
		c.Set("body", dst.String())
		c.Next()

	}
}
