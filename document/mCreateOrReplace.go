package document

import (
	"encoding/json"
	"strconv"

	"github.com/kuzzleio/sdk-go/types"
)

func (d *Document) MCreateOrReplace(index string, collection string, body string, options *DocumentOptions) (string, error) {
	if index == "" {
		return "", types.NewError("Document.MCreateOrReplace: index required", 400)
	}

	if collection == "" {
		return "", types.NewError("Document.MCreateOrReplace: collection required", 400)
	}

	if body == "" {
		return "", types.NewError("Document.MCreateOrReplace: body required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Index:      index,
		Collection: collection,
		Controller: "document",
		Action:     "mCreateOrReplace",
		Body:       body,
	}

	queryOpts := types.NewQueryOptions()

	if options != nil {
		queryOpts.SetRefresh(strconv.FormatBool(options.WaitFor))
	}

	go d.Kuzzle.Query(query, queryOpts, ch)

	res := <-ch

	if res.Error != nil {
		return "", res.Error
	}

	var mCreated string
	json.Unmarshal(res.Result, &mCreated)

	return mCreated, nil
}
