package api

import (
	"net/http"

	"ebay_dapp_golang/pkg/logger"
	"ebay_dapp_golang/utils/errs"

	"github.com/gin-gonic/gin"
)

type response struct {
	RequestID string      `json:"request_id"`
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
	Next      string      `json:"next"`
}

//JSON - Public interface return
func JSON(c *gin.Context, httpCode int, data interface{}, err error) {
	if err != nil {
		logger.Error(err.Error())
	}

	c.JSON(http.StatusOK, response{
		Code: httpCode,
		Msg:  errs.ErrorMsg[httpCode],
		Data: data,
		Next: "",
	})

	return
}

//Serve
func Serve() {
	gin.SetMode(gin.ReleaseMode)

	g := gin.New()

	g = routers.Load(g)

	if err := g.Run(":8082"); err != nil {
		logger.F.Error("service start fail: " + err.Error())
	}
}
