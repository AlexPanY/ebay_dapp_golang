package ifps

import (
	"encoding/json"
	"strings"

	shell "github.com/ipfs/go-ipfs-api"
)

var sh *shell.Shell

var (
	defaultShellURL = "localhost:5001"
)

func init() {
	sh = shell.NewShell(defaultShellURL)
}

//UpdateMetadataWithJSON
func UpdateMetadataWithJSON(data interface{}) (string, error) {
	metadataBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return sh.Add(strings.NewReader(string(metadataBytes)))
}
