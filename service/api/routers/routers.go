package routers

import (
	v1 "ebay_dapp_golang/api/v1"

	"github.com/gin-gonic/gin"
)

func Load(g *gin.Engine) *gin.Engine {

	v1Api := g.Group("/v1")
	{
		v1Api.POST("/user/get", v1.GetUserInfoByID)

		//nft token
		v1Api.POST("/asset/create", v1.CreateYPTToken)
	}

	return g
}
