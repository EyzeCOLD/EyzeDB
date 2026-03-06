package table

import (
	"testing"
)

func TestSelect(t *testing.T) {
	testCases := map[string]struct {
		columns []Column
		expected TableView
		wantErr bool
	}{
		"One column, OK": {
			columns: Column{
				{}
			}
		},
	}
}
