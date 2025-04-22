package golibstraits

import "context"

// Queryable defines a generic interface for querying data sources.
type Queryable[Q any] interface {
	// Basic lookups
	Find(matchers map[string]any) (Q, error)
	FindAll(matchers map[string]any) ([]Q, error)
	FindAllInRange(matchers map[string]any, limit, offset int) ([]Q, error)

	// Variadic argument lookups
	FindByArgs(args ...any) (Q, error)
	FindAllByArgs(args ...any) ([]Q, error)

	// Lookups by ID
	FindByID(id any) (Q, error)
	FindAllByID(id any) ([]Q, error)

	// Source-specific lookups
	FindInSource(source string, matchers map[string]any) (Q, error)
	FindAllInSource(source string, matchers map[string]any) ([]Q, error)
	FindAllInSourceRange(source string, matchers map[string]any, limit, offset int) ([]Q, error)

	// Context-aware versions
	FindCtx(ctx context.Context, matchers map[string]any) (Q, error)
	FindAllCtx(ctx context.Context, matchers map[string]any) ([]Q, error)

	// Search by text or keywords
	Search(term string) ([]Q, error)
	SearchInSource(source string, term string) ([]Q, error)

	// Filter by date range or numerical fields
	FilterByFieldRange(field string, min, max any) ([]Q, error)

	// Sorting and pagination
	FindSorted(matchers map[string]any, sortBy string, asc bool, limit, offset int) ([]Q, error)

	// Exists checks
	Exists(matchers map[string]any) (bool, error)
	Count(matchers map[string]any) (int, error)
}