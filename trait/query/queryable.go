package query

// NOT READY FOR PRODUCTION
type Queryable[Result any] interface {
	FinderByID[Result]
	Finder[Result]
	Exister
	Searcher[Result]
	FinderByID[Result]
}

// NOT READY FOR PRODUCTION
type QueryableWithCTX[Result any] interface {
	FinderByIDWithCTX[Result]
	FinderWithCTX[Result]
	ExisterWithCTX
	SearcherWithCTX[Result]
	FinderByIDWithCTX[Result]
}
