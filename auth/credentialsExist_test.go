package auth_test

import (
	"encoding/json"
	"testing"

	"github.com/kuzzleio/sdk-go/connection/websocket"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestCredentialsExistEmptyStrategy(t *testing.T) {
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)
	k.Connect()
	_, err := k.Auth.CredentialsExist("", nil)
	assert.NotNil(t, err)
}

func TestCredentialsExistQueryError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "auth", request.Controller)
			assert.Equal(t, "credentialsExist", request.Action)

			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	k.Connect()
	_, err := k.Auth.CredentialsExist("local", nil)
	assert.NotNil(t, err)
}

func TestCredentialsExists(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "auth", request.Controller)
			assert.Equal(t, "credentialsExist", request.Action)

			ret, _ := json.Marshal(true)
			return &types.KuzzleResponse{Result: ret}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	k.Connect()
	res, _ := k.Auth.CredentialsExist("local", nil)
	assert.Equal(t, true, res)
}

func ExampleAdminExists() {
	c := websocket.NewWebSocket("localhost", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	k.Connect()
	res, _ := k.Auth.CredentialsExist("local", nil)
	println(res)
}
