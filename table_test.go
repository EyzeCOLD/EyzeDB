package model

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

func TestInsert(t *testing.T) {
	testCases := map[string]struct {
		values  []Value
		wantErr bool
	}{
		"matching insert": {
			values: []Value{
				{Type: ColTypeText, String: "James"},
				{Type: ColTypeInt, Integer: 100},
				{Type: ColTypeFloat, Float: 0.25},
			},
			wantErr: false,
		},
		"matching insert 2": {
			values: []Value{
				{Type: ColTypeText, String: "Boblerone"},
				{Type: ColTypeInt, Integer: 300000},
				{Type: ColTypeFloat, Float: 0.99},
			},
			wantErr: false,
		},
		"Type mismatch": {
			values: []Value{
				{Type: ColTypeInt, Integer: 661},
				{Type: ColTypeInt, Integer: 300000},
				{Type: ColTypeFloat, Float: 0.99},
			},
			wantErr: true,
		},
		"Type mismatch 2": {
			values: []Value{
				{Type: ColTypeText, String: "Boblerone"},
				{Type: ColTypeInt, Integer: 300000},
				{Type: ColTypeInt, Integer: 99},
			},
			wantErr: true,
		},
		"Type mismatch inside value": {
			values: []Value{
				{Type: ColTypeText, String: "Boblerone"},
				{Type: ColTypeInt, Float: 3000},
				{Type: ColTypeInt, Integer: 99},
			},
			wantErr: true,
		},
		"Wrong amount of columns": {
			values: []Value{
				{Type: ColTypeText, String: "Boblerone"},
				{Type: ColTypeInt, Integer: 3000},
			},
			wantErr: true,
		},
	}

	columns := []Column{
		{Name: "name", Type: ColTypeText},
		{Name: "balance", Type: ColTypeInt},
		{Name: "rate", Type: ColTypeFloat},
	}

	table := CreateTable("test", columns...)

	expectedSize := 0
	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			err := table.Insert(test.values...)
			if err != nil && !test.wantErr {
				t.Errorf("error: %v", err)
			}
			if err == nil && test.wantErr {
				t.Error("expected error, got none")
			}
			if test.wantErr {
				return
			}

			expectedSize++
			if len(table.Records) != expectedSize {
				t.Errorf("Record wasn't inserted to the table")
				return
			}

			lastIndex := len(table.Records) - 1
			lastRecord := table.Records[lastIndex]
			for col := range table.Records[lastIndex] {
				if lastRecord[col].Type != table.Columns[col].Type {
					t.Errorf("Type mismatch in column %d", col)
				}
				if lastRecord[col] != test.values[col] {
					t.Errorf("Value mismatch in column %d", col)
				}
			}
		})
	}
}
