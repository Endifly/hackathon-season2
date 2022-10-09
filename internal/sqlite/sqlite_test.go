package sqlite

import (
	"fmt"
	"testing"
)

func TestToSqlite(t *testing.T) {
	err := ToSqlite("../../", "./target.csv")

	if err != nil {
		// t.Errorf("error")
		fmt.Println(err)
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
		// t.Errorf("error")
	}
}

func TestSqlField(t *testing.T) {
	field1 := new(SqlField).PrimaryKey().Text()
	fmt.Println(field1.Schema())

}
