package table

import (
	"fmt"
)

type ColType int

// Enumerator for different column datatypes
const (
	ColTypeInt ColType = iota
	ColTypeFloat
	ColTypeText
	ColTypeNull
)

// Stringify column datatypes
func (ct ColType) String() string {
	switch ct {
	case ColTypeInt:
		return "Integer"
	case ColTypeFloat:
		return "Float"
	case ColTypeText:
		return "Text"
	default:
		return "NULL"
	}
}

// Column type
type Column struct {
	Name string
	Type ColType
}

// A "union" type for table values, you need to look at the Type to know
// which field to read and write.
type Value struct {
	Type    ColType
	Integer int64
	Float   float64
	String  string
}

// Table Expression interface, generalizes Tables and TableViews
type TableExp interface {
	GetCols() []Column
	RowCount() int
	GetVal(row, col int) Value
	GetValPtr(row, col int) *Value
}

// Represents a table in the database
type Table struct {
	Name    string
	Columns []Column
	Records [][]Value
}

func (t *Table) GetCols() []Column {
	return t.Columns
}

func (t *Table) RowCount() int {
	return len(t.Records)
}

func (t *Table) GetVal(row, col int) Value {
	return t.Records[row][col]
}

func (t *Table) GetValPtr(row, col int) *Value {
	return &t.Records[row][col]
}

// A view or a reference to a table or a subset of a table in the database
type TableView struct {
	Source  TableExp
	Columns []int
	Records []int
}

func (tv *TableView) GetCols() []Column {
	return tv.Source.GetCols()
}

func (tv *TableView) RowCount() int {
	return tv.Source.RowCount()
}

func (tv *TableView) GetVal(row, col int) Value {
	return tv.Source.GetVal(row, col)
}

func (tv *TableView) GetValPtr(row, col int) *Value {
	return tv.Source.GetValPtr(row, col)
}

// debug printout
func (t *Table) Print() {
	fmt.Println(t.Name, "{")
	for row := range t.Records {
		for col := range t.Columns {
			fmt.Printf("\"%s\":", t.Columns[col].Name)
			switch t.Columns[col].Type {
			case ColTypeInt:
				fmt.Printf("\"%d\"", t.Records[row][col].Integer)
			case ColTypeFloat:
				fmt.Printf("\"%f\"", t.Records[row][col].Float)
			case ColTypeText:
				fmt.Printf("\"%s\"", t.Records[row][col].String)
			case ColTypeNull:
				fmt.Printf("\"NULL\"")
			}
		}
		fmt.Println(",")
	}
	fmt.Println("}")
}
