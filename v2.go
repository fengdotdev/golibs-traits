package golibstraits

import "context"

// Finder defines basic retrieval operations.
type Finder[Q any] interface {
	Find(ctx context.Context, matchers map[string]any) (Q, error)
	FindAll(ctx context.Context, matchers map[string]any) ([]Q, error)
}

// FinderByID defines operations for retrieving by ID.
type FinderByID[Q any] interface {
	FindByID(ctx context.Context, id any) (Q, error)
	FindAllByID(ctx context.Context, ids ...any) ([]Q, error) // Allow variadic IDs
}

// Searcher defines text-based search capabilities.
type Searcher[Q any] interface {
	Search(ctx context.Context, term string) ([]Q, error)
}

// Filterer defines filtering operations.
type Filterer[Q any] interface {
	FilterByFieldRange(ctx context.Context, field string, min, max any) ([]Q, error)
}

// Sorter defines sorting and pagination.
type Sorter[Q any] interface {
	FindSorted(ctx context.Context, matchers map[string]any, sortBy string, asc bool, limit, offset int) ([]Q, error)
}

// Exister defines existence and counting checks.
type Exister interface {
	Exists(ctx context.Context, matchers map[string]any) (bool, error)
	Count(ctx context.Context, matchers map[string]any) (int, error)
}

// Queryable combines common query interfaces.
type Queryable[Q any] interface {
	Finder[Q]
	FinderByID[Q]
	// Add other relevant interfaces here
}

// SourceQueryable extends Queryable for source-specific operations.
type SourceQueryable[Q any] interface {
	Queryable[Q]
	FindInSource(ctx context.Context, source string, matchers map[string]any) (Q, error)
	FindAllInSource(ctx context.Context, source string, matchers map[string]any) ([]Q, error)
	// Consider a FindAllInSourceRange if pagination is common for sources
}

// AdvancedQueryable combines more specialized interfaces.
type AdvancedQueryable[Q any] interface {
	SourceQueryable[Q]
	Searcher[Q]
	Filterer[Q]
	Sorter[Q]
	Exister
	FindByArgs(ctx context.Context, args ...any) (Q, error)
	FindAllByArgs(ctx context.Context, args ...any) ([]Q, error)
}


type QueryOptions struct {
	Limit  int
	Offset int
	SortBy string
	Asc    bool
	Source string
	// Add other options as needed
}

type QueryBuilder[Q any] interface {
	Where(matchers map[string]any) QueryBuilder[Q]
	WhereID(id any) QueryBuilder[Q]
	WhereIDs(ids ...any) QueryBuilder[Q]
	Search(term string) QueryBuilder[Q]
	FilterRange(field string, min, max any) QueryBuilder[Q]
	SortBy(field string, asc bool) QueryBuilder[Q]
	Limit(limit int) QueryBuilder[Q]
	Offset(offset int) QueryBuilder[Q]
	FromSource(source string) QueryBuilder[Q]
	FindByArgs(args ...any) QueryBuilder[Q] // Consider type safety here

	// Execution methods
	Find(ctx context.Context) (Q, error)
	FindAll(ctx context.Context) ([]Q, error)
	Exists(ctx context.Context) (bool, error)
	Count(ctx context.Context) (int, error)
}
