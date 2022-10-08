package sqlite

import (
	"fmt"
	"testing"
)

func TestToSqlite(t *testing.T) {
	err := ToSqlite("../../", "./target.csv")

	if err != nil {
		t.Errorf("error")
	}
}

func TestToSqliteInvalidPath(t *testing.T) {
	err := ToSqlite("nowhere", "./target.csv")

	if err == nil {
		t.Errorf("error")
	}
}

func TestWriteSql(t *testing.T) {
	err := WriteSql("./target.sqlite")

	if err != nil {
		t.Errorf("error")
	}
}

func TestSqlField(t *testing.T) {
	field1 := new(SqlField).Id().Text()
	fmt.Println(field1.Schema())

}
