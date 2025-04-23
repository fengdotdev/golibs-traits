package query

import "context"

type Searcher[Result any] interface {
	Search(term string, where string) ([]Result, error)
}

type SearcherWithCTX[Result any] interface {
	Search(ctx context.Context, term string, where string) ([]Result, error)
}
