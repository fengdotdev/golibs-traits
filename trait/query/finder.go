package query

import "context"

// NOT READY FOR PRODUCTION
type Finder[Result any] interface {
	Find(matchers map[string]any) (Result, error)
	FindAll(matchers map[string]any) ([]Result, error)
	FindAllInRange(matchers map[string]any, limit, offset int) ([]Result, error)
}
// NOT READY FOR PRODUCTION
type FinderWithCTX[Result any] interface {
	Find(ctx context.Context, matchers map[string]any) (Result, error)
	FindAll(ctx context.Context, matchers map[string]any) ([]Result, error)
	FindAllInRange(ctx context.Context, matchers map[string]any, limit, offset int) ([]Result, error)
}
// NOT READY FOR PRODUCTION
type FinderByID[Q any] interface {
	FindByID(id any) (Q, error) 
	FindAllByID(id any) ([]Q, error)
	FinderByIDInRange(id any, limit, offset int) ([]Q, error)
}
// NOT READY FOR PRODUCTION
type FinderByIDWithCTX[Q any] interface {
	FindByID(ctx context.Context, id any) (Q, error)
	FindAllByID(ctx context.Context, id any) ([]Q, error)
	FinderByIDInRange(ctx context.Context, id any, limit, offset int) ([]Q, error)
}
