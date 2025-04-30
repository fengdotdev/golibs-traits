package query

import "context"

// NOT READY FOR PRODUCTION
type Exister interface {
	Exists(matchers map[string]any) (bool, error)
	Count(matchers map[string]any) (int, error)
}
// NOT READY FOR PRODUCTION
type ExisterWithCTX interface {
	Exists(ctx context.Context, matchers map[string]any) (bool, error)
	Count(ctx context.Context, matchers map[string]any) (int, error)
}
