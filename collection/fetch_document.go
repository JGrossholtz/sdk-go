package collection

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Retrieves a Document using its provided unique id.
*/
func (dc Collection) FetchDocument(id string, options types.QueryOptions) (Document, error) {
	if id == "" {
		return Document{}, errors.New("Collection.FetchDocument: document id required")
	}

	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Collection: dc.collection,
		Index:      dc.index,
		Controller: "document",
		Action:     "get",
		Id:         id,
	}
	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return Document{}, errors.New(res.Error.Message)
	}

	document := Document{}
	json.Unmarshal(res.Result, &document)

	return document, nil
}

/*
  Get specific documents according to given IDs.
*/
func (dc Collection) MGetDocument(ids []string, options types.QueryOptions) (KuzzleSearchResult, error) {
	if len(ids) == 0 {
		return KuzzleSearchResult{}, errors.New("Collection.MGetDocument: please provide at least one id of document to retrieve")
	}

	ch := make(chan types.KuzzleResponse)

	type body struct {
		Ids []string `json:"ids"`
	}

	query := types.KuzzleRequest{
		Collection: dc.collection,
		Index:      dc.index,
		Controller: "document",
		Action:     "mGet",
		Body:       &body{Ids: ids},
	}
	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return KuzzleSearchResult{}, errors.New(res.Error.Message)
	}

	result := KuzzleSearchResult{}
	json.Unmarshal(res.Result, &result)

	for _, d := range result.Hits {
		d.collection = dc
	}

	return result, nil
}
