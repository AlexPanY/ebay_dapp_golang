package v1

import (
	"ebay_dapp_golang/api"
	"ebay_dapp_golang/errno"
	"ebay_dapp_golang/token"

	"github.com/gin-gonic/gin"
)

//CreateYPTTokenRequest
type CreateYPTTokenRequest struct {
	Account     string                          `json:"Account"`
	Name        string                          `json:"name"`
	Description string                          `json:"description"`
	Image       string                          `json:"image"`
	Collection  string                          `json:"collection"`
	Properties  []*token.ERC721_YopuNFTProperty `json:"properties"`
}

//CreateYPTToken
func CreateYPTToken(c *gin.Context) {
	var req CreateYPTTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		api.ErrJSONWithRawErr(c, errno.ErrParamInvalid, err)
		return
	}

	t := token.ERC721_YopuNFT{
		Name:        req.Name,
		Description: req.Description,
		Image:       req.Image,
		Collection:  req.Collection,
		Properties:  req.Properties,
	}

	if err := t.Mint(); err != nil {
		api.ErrJSONWithRawErr(c, errno.ErrParamInvalid, err)
		return
	}
	api.SuccJSONWithData(c, 1)
	return
}
