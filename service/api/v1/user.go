package v1

import (
	"net/http"

	"ebay_dapp_golang/api"
	"ebay_dapp_golang/internal/model"
	"ebay_dapp_golang/internal/service"

	"github.com/gin-gonic/gin"
)

//GetUserInfoByIDRequest - request
type GetUserInfoByIDRequest struct {
	Request string `json:"request_id"`
	UserID  uint64 `json:"user_id"`
}

//GetUserInfoByID - Get user info by user_id
func GetUserInfoByID(c *gin.Context) {
	var req GetUserInfoByIDRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		api.JSON(c, http.StatusOK, nil, nil)
		return
	}
	user, _ := service.NewUserService().GetUserInfoByID(req.UserID)
	api.JSON(c, http.StatusOK, user, nil)
	return
}
