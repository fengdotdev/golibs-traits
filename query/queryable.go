package query

type Queryable[Result any] interface {
	FinderByID[Result]
	Finder[Result]
	Exister
	Searcher[Result]
	FinderByID[Result]
}

type QueryableWithCTX[Result any] interface {
	FinderByIDWithCTX[Result]
	FinderWithCTX[Result]
	ExisterWithCTX
	SearcherWithCTX[Result]
	FinderByIDWithCTX[Result]
}
