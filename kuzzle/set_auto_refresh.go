package kuzzle

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// SetAutoRefresh autorefresh status setter for the provided data index name
func (k *Kuzzle) SetAutoRefresh(index string, autoRefresh bool, options types.QueryOptions) (bool, error) {
	if index == "" {
		if k.defaultIndex == "" {
			return false, types.NewError("Kuzzle.SetAutoRefresh: index required", 400)
		}
		index = k.defaultIndex
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "index",
		Action:     "setAutoRefresh",
		Index:      index,
		Body: struct {
			AutoRefresh bool `json:"autoRefresh"`
		}{autoRefresh},
	}
	go k.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return false, res.Error
	}

	type autoRefreshResponse struct {
		Response bool `json:"response"`
	}

	var r autoRefreshResponse
	json.Unmarshal(res.Result, &r)

	return r.Response, nil
}
