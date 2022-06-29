package ifps

import (
	"fmt"
	"testing"
)

//TestMetaDataToken
type TestMetaDataToken struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Strength    int    `json:"strength"`
}

//TestUpdateMetadata
//NOTE: go test -run TestUpdateMetadata -v
func TestUpdateMetadata(t *testing.T) {

	metadata := &TestMetaDataToken{
		Name:        "Thor's hammer",
		Description: "the legendary hammer of the Norse god of thunder.",
		Image:       "https://game.example/item-id-8u5h2m.png",
		Strength:    20,
	}

	hash, err := UpdateMetadataWithJSON(metadata)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(hash)
}
