package table

import (
	"errors"
)

func (t *Table) Insert(values ...Value) error {
	if len(values) != len(t.Columns) {
		return errors.New("(t *Table) Insert(): column count mismatch")
	}
	record := make([]Value, len(t.Columns))
	for col := range values {
		if values[col].Type == ColTypeNull {
			record[col].Type = ColTypeNull
		} else if t.Columns[col].Type != values[col].Type {
			return errors.New("(t *Table) Insert(): type mismatch")
		} else {
			record[col] = values[col]
		}
	}
	t.Records = append(t.Records, record)
	return nil
}
