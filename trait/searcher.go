package trait

import "context"

type Searcher[S any, R any] interface {
	Search(term S, where func(result R) bool, from ...string) ([]R, error)
	SearchAll(term S) ([]R, error)
	Count(term S, where func(result R) bool, from ...string) (int, error)
	CountAll(term S) (int, error)
	CountAllFrom(term S, from ...string) (int, error)
}

type SearcherWithCTX[S any, R any] interface {
	Search(ctx context.Context, term S, where func(R) bool, from ...string) ([]R, error)
	SearchAll(ctx context.Context, term S) ([]R, error)
}
