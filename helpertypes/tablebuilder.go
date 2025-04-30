package helpertypes

type TableBuilder struct {
	RowData    []any
	ColumnData []any
}

func (tb *TableBuilder) AddRow(rowData any) {
	tb.RowData = append(tb.RowData, rowData)
}

func (tb *TableBuilder) AddColumn(columnData any) {
	tb.ColumnData = append(tb.ColumnData, columnData)
}

func (tb *TableBuilder) Build() (Table, error) {

	rowslen := len(tb.RowData)
	columnslen := len(tb.ColumnData)
	size := rowslen * columnslen
	if size == 0 {
		return Table{}, nil
	}

	t := Table{
		rows:    rowslen,    // 1, 2, 3,4
		columns: columnslen, // A, B, C,D
		headersRows:   tb.RowData,
		headersColumns: tb.ColumnData,
		data: make(map[string]any,size), // 4 x 4 =  16 cells

		// & | A | B | C | D
		// 1 | A1| B1| C1| D1
		// 2 | A2| B2| C2| D2
		// 3 | A3| B3| C3| D3
		// 4 | A4| B4| C4| D4
	}

	t.add("A1", "A1")

	return t, nil
}
