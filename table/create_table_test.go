package table

import (
	"testing"
)

func TestCreateTable(t *testing.T) {
	testCases := map[string]struct {
		tableName string
		columns   []Column
	}{
		"text, int, float": {
			tableName: "accounts",
			columns: []Column{
				{"name", ColTypeText},
				{"balance", ColTypeInt},
				{"rate", ColTypeFloat},
			},
		},
		"text, text, int": {
			tableName: "pets",
			columns: []Column{
				{"name", ColTypeText},
				{"nickname", ColTypeText},
				{"age", ColTypeInt},
			},
		},
		"empty table": {
			tableName: "empty",
			columns:   nil,
		},
		"empty name": {
			tableName: "",
			columns: []Column{
				{"type", ColTypeText},
				{"language", ColTypeText},
				{"mass", ColTypeFloat},
			},
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			table := CreateTable(test.tableName, test.columns...)
			if table.Name != test.tableName {
				t.Errorf("tableName: %v wanted: %v", table.Name, test.tableName)
			}
			if len(table.Columns) != len(test.columns) {
				t.Errorf("table columns: %d wanted: %d",
					len(table.Columns), len(test.columns))
			}
			for col := range table.Columns {
				if table.Columns[col].Type != test.columns[col].Type {
					t.Errorf("table column type: %v wanted: %v",
						table.Columns[col].Type, test.columns[col].Type)
				}
				if table.Columns[col].Name != test.columns[col].Name {
					t.Errorf("table column name: %v wanted: %v",
						table.Columns[col].Name, test.columns[col].Name)
				}
			}
		})
	}
}
