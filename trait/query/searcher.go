package query

import "context"

// NOT READY FOR PRODUCTION
type Searcher[Result any] interface {
	Search(term string, where string) ([]Result, error)
}

// NOT READY FOR PRODUCTION
type SearcherWithCTX[Result any] interface {
	Search(ctx context.Context, term string, where string) ([]Result, error)
}
