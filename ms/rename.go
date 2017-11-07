package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Rename renames a key to newkey. If newkey already exists, it is overwritten.
func (ms Ms) Rename(key string, newkey string, options types.QueryOptions) (string, error) {
	result := make(chan *types.KuzzleResponse)

	type body struct {
		NewKey string `json:"newkey"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "rename",
		Id:         key,
		Body:       &body{NewKey: newkey},
	}
	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return "", res.Error
	}
	var returnedResult string
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
