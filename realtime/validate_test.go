// Copyright 2015-2018 Kuzzle
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package realtime_test

import (
	"encoding/json"
	"testing"

	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/realtime"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestValidateIndexNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	nr := realtime.NewRealtime(k)

	_, err := nr.Validate("", "collection", "body", nil)

	assert.NotNil(t, err)
}

func TestValidateCollectionNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	nr := realtime.NewRealtime(k)

	_, err := nr.Validate("index", "", "body", nil)

	assert.NotNil(t, err)
}

func TestValidateBodyNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	nr := realtime.NewRealtime(k)

	_, err := nr.Validate("index", "collection", "", nil)

	assert.NotNil(t, err)
}

func TestValidateError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: types.NewError("Unit test error")}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	nr := realtime.NewRealtime(k)

	_, err := nr.Validate("index", "collection", "body", nil)
	assert.NotNil(t, err)
}

func TestValidate(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "realtime", parsedQuery.Controller)
			assert.Equal(t, "validate", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)
			assert.NotNil(t, parsedQuery.Body)

			res := types.KuzzleResponse{Result: []byte(`{"valid": true}`)}

			r, _ := json.Marshal(res.Result)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	nr := realtime.NewRealtime(k)

	res, err := nr.Validate("index", "collection", "body", nil)
	assert.Nil(t, err)
	assert.Equal(t, true, res)
}
