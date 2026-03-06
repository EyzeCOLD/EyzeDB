package table

func CreateTable(name string, columns ...Column) Table {
	t := Table{
		Name:    name,
		Columns: make([]Column, len(columns)),
		Records: make([][]Value, 0),
	}

	copy(t.Columns, columns)

	return t
}
