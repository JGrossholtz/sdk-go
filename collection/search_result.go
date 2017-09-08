package collection

import (
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

type KuzzleSearchResult struct {
	Collection Collection
	Hits       []Document `json:"hits"`
	Total      int        `json:"total"`
	ScrollId   string     `json:"_scroll_id"`
	Options    types.QueryOptions
	Filters    types.SearchFilters
}

func (ksr KuzzleSearchResult) FetchNext() (KuzzleSearchResult, error) {
	if ksr.ScrollId != "" {
		var options = ksr.Options
		options.SetFrom(0)
		options.SetSize(0)

		return ksr.Collection.Scroll(ksr.ScrollId, options)
	}

	if ksr.Options != nil {
		if ksr.Options.GetSize() != 0 && len(ksr.Filters.Sort) > 0 {
			var filters = ksr.Filters

			for _, sortRules := range filters.Sort {
				for field := range sortRules {
					var source = ksr.Hits[len(ksr.Hits)-1].SourceToMap()
					filters.SearchAfter = append(filters.SearchAfter, source[field])
				}
			}

			var options = ksr.Options
			options.SetFrom(0)

			return ksr.Collection.Search(filters, options)
		}

		if ksr.Options.GetSize() != 0 {
			var options = ksr.Options
			options.SetFrom(ksr.Options.GetFrom() + ksr.Options.GetSize())

			if options.GetFrom() >= ksr.Total {
				return KuzzleSearchResult{}, nil
			}

			return ksr.Collection.Search(ksr.Filters, options)
		}
	}

	return KuzzleSearchResult{}, errors.New("KuzzleSearchResult.FetchNext: Unable to retrieve next results from search: missing scrollId or from/size parameters")
}
