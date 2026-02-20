package model

import (
	"errors"
	"fmt"
)

type ColType int

const (
	ColTypeInt ColType = iota
	ColTypeFloat
	ColTypeText
	ColTypeNull
)

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

type Column struct {
	Name string
	Type ColType
}

type Value struct {
	Type    ColType
	Integer int64
	Float   float64
	String  string
}

type Table struct {
	Name    string
	Columns []Column
	Records [][]Value
}

func CreateTable(name string, columns ...Column) Table {
	t := Table{
		Name:    name,
		Columns: make([]Column, len(columns)),
		Records: make([][]Value, 0),
	}

	copy(t.Columns, columns)

	return t
}

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
