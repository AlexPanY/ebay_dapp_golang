package v1

import (
	"ebay_dapp_golang/api"
	"ebay_dapp_golang/errno"

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
		api.ErrJSONWithRawErr(c, errno.ErrParamInvalid, err)
		return
	}
	// user, _ := service.NewUserService().GetUserInfoByID(req.UserID)
	user := make([]int, 0)

	api.SuccJSONWithData(c, user)
	return
}
