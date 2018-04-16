// Copyright 2015-2017 Kuzzle
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

package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
	"strconv"
)

// Incrbyfloat increments the number stored at key by the provided float value.
// If the key does not exist, it is set to 0 before performing the operation.
func (ms *Ms) Incrbyfloat(key string, value float64, options types.QueryOptions) (float64, error) {
	result := make(chan *types.KuzzleResponse)

	type body struct {
		Value float64 `json:"value"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "incrbyfloat",
		Id:         key,
		Body:       &body{Value: value},
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return 0, res.Error
	}

	var stringResult string
	json.Unmarshal(res.Result, &stringResult)

	return strconv.ParseFloat(stringResult, 64)
}
