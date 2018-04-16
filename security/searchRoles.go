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

package security

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

func (s *Security) SearchRoles(body json.RawMessage, options types.QueryOptions) (*RoleSearchResult, error) {
	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "searchRoles",
		Body:       body,
	}

	if options != nil {
		query.From = options.From()
		query.Size = options.Size()

		scroll := options.Scroll()
		if scroll != "" {
			query.Scroll = scroll
		}
	}

	go s.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}
	jsonSearchResult := &jsonRoleSearchResult{}
	json.Unmarshal(res.Result, jsonSearchResult)

	searchResult := &RoleSearchResult{
		Total: jsonSearchResult.Total,
	}
	for _, j := range jsonSearchResult.Hits {
		r := j.jsonRoleToRole()
		searchResult.Hits = append(searchResult.Hits, r)
	}

	return searchResult, nil
}
