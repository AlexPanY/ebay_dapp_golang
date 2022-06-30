package api

import (
	"fmt"
	"net/http"

	"ebay_dapp_golang/errno"
	log "ebay_dapp_golang/pkg/logger"

	"github.com/gin-gonic/gin"
)

//日志格式
var (
	LogErrFormatTemplate  = `req_id="%v" method="%+v" uri="%+v" msg="%+v" err="%v"`
	LogInfoFormatTemplate = `req_id="%v" method="%+v" uri="%+v" msg="%+v"`
)

const (
	X_REQUEST_ID = "X-REQUEST-ID"
)

//ApiResponse API response format
type ApiResponse struct {
	RequestID string `json:"requestId"`
	Code      string `json:"code"`
	Message   string `json:"message"`
}

type ApiResponseData struct {
	RequestID string      `json:"requestId"`
	Code      string      `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

func SuccJSONWithData(c *gin.Context, data interface{}) {
	if data == nil {
		data = make([]int, 0)
	}

	rep := ApiResponseData{
		RequestID: c.Request.Header.Get(X_REQUEST_ID),
		Code:      errno.Success.Code,
		Message:   errno.Success.Message,
		Data:      data,
	}

	c.JSON(http.StatusOK, rep)
	return
}

//ErrJSONWithRawErr
func ErrJSONWithRawErr(c *gin.Context, err *errno.Err, rawErr error) {

	rep := ApiResponse{
		RequestID: c.Request.Header.Get(X_REQUEST_ID),
		Code:      err.Code,
		Message:   err.Message,
	}

	if rawErr != nil {
		log.Error(fmt.Sprintf(LogErrFormatTemplate,
			rep.RequestID,
			c.Request.Method,
			c.Request.URL.Path,
			err.Message,
			rawErr.Error()),
		)
	}

	c.JSON(http.StatusOK, rep)
	return
}
