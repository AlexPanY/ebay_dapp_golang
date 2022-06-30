package token

import (
	"fmt"

	"ebay_dapp_golang/pkg/ifps"
)

//ERC721_YopuNFTProperty
type ERC721_YopuNFTProperty struct {
	field string `json:"field"`
	Value string `json:"value"`
}

//ERC721_YopuNFT
type ERC721_YopuNFT struct {
	Name        string                    `json:"name"`
	Description string                    `json:"description"`
	Image       string                    `json:"image"`
	Collection  string                    `json:"collection"`
	Properties  []*ERC721_YopuNFTProperty `json:"properties"`
}

//Mint Minting a single token based on ERC721.
func (t *ERC721_YopuNFT) Mint() error {
	uri, err := t.NewMetadataURI()
	if err != nil {
		return err
	}
	fmt.Println(uri)
	return nil
}

//NewMetadataURI
func (t *ERC721_YopuNFT) NewMetadataURI() (string, error) {
	return ifps.UpdateMetadataWithJSON(t)
}
