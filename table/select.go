package table

import (
	"errors"
	"slices"
)

func Select(source TableExp, columns []Column) (TableView, error) {
	result := TableView{
		Source:  source,
		Columns: make([]int, len(columns)),
		Records: make([]int, source.RowCount()),
	}

	for i := range columns {
		if !slices.Contains(source.GetCols(), columns[i]) {
			return result, errors.New(
				"(t *Table) Select(): Invalid column name")
		}
		result.Columns[i] = i
	}

	for row := range source.RowCount() {
		result.Records[row] = row
	}

	return result, nil
}
