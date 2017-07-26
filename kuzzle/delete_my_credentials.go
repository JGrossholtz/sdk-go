package kuzzle

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
 * Delete credentials of the specified strategy for the current user.
 */
func (k Kuzzle) DeleteMyCredentials(strategy string, options *types.Options) (*types.AckResponse, error) {
	if strategy == "" {
		return nil, errors.New("Kuzzle.DeleteMyCredentials: strategy is required")
	}

	type body struct {
		Strategy string `json:"strategy"`
	}
	result := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "auth",
		Action:     "deleteMyCredentials",
		Body:       &body{strategy},
	}

	go k.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return nil, errors.New(res.Error.Message)
	}

	ack := &types.AckResponse{}
	json.Unmarshal(res.Result, &ack)

	return ack, nil
}
