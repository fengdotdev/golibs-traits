package helpertypes

type Table struct {
	rows           int
	columns        int
	data           map[string]any
	headersRows    []any
	headersColumns []any
}

// public







// private
func (t *Table) add(key string, value any) error{
	// check if the key already exists
	if _, exists := t.data[key]; exists {
		return fmt.Errorf("key %s already exists", key)
	}

}


func (t *Table) createKey(row , column int) string {
	// Create a unique key for the cell based on its row and column
	return fmt.Sprintf("%d-%d", row, column)
}