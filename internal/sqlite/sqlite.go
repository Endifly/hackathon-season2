package sqlite

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type SqlField struct {
	schema  []string
	colName string
}

func (f *SqlField) addSchema(s string) {
	f.schema = append(f.schema, s)
}

func (f *SqlField) Name(name string) *SqlField {
	f.colName = name
	return f
}

func (f *SqlField) PrimaryKey() *SqlField {
	f.addSchema("primary key")
	return f
}

func (f *SqlField) Text() *SqlField {
	f.addSchema("text")
	return f
}

func (f *SqlField) Int() *SqlField {
	f.addSchema("integer")
	return f
}

func (f *SqlField) Schema() string {
	return strings.Join(f.schema, " ")
}

func NewField(name string) *SqlField {
	f := &SqlField{}
	f.Name(name)
	return f
}

type SqlData struct {
	Conn *sql.DB
}

func (c *SqlData) CreateTable(tableName string, columns []*SqlField) error {
	cmd := fmt.Sprintf("create table if not exists ? (id integer primary key autoincrement, title text, author text)")
	stmt, err := c.Conn.Prepare("create table if not exists books (id integer primary key autoincrement, title text, author text)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	return nil
}

func (c *SqlData) AddRow(data []string) error {

	stmt, err := c.Conn.Prepare("insert into books(title, author) values(?, ?)")
	if err != nil {
		// log.Fatalf("insert prepare failed: %s", err)
		return err
	}

	_, err = stmt.Exec(data)
	if err != nil {
		return err
	}

	return nil
}

/*
	alway return db connection
*/
func getDb(filePath string) SqlData {
	_, err := os.Stat(filePath)

	if err != nil {
		file, createErr := os.Create(filePath)

		if createErr != nil {
			return SqlData{}
		}

		file.Close()
	}

	database, err := sql.Open("sqlite3", filePath)

	if err != nil {
		return SqlData{}
	}

	return SqlData{Conn: database}

}

func WriteSql(filePath string) error {
	db := getDb(filePath)
	defer db.Conn.Close()

	if db.Conn == nil {
		fmt.Println("erro")
	}

	err := db.CreateTable("devMountain2", []*SqlField{
		NewField("EMPID").PrimaryKey(),
		NewField("PASSPORT").Text(),
		NewField("FIRSTNAME").Text(),
		NewField("LASTNAME").Text(),
		NewField("GENDER").Int(),
		NewField("BIRTHDAY").Text(),
		NewField("NATIONALITY").Text(),
		NewField("HIRED").Text(),
		NewField("DEPT").Text(),
		NewField("POSITION").Text(),
		NewField("STATUS").Int(),
		NewField("REGION").Text(),
	})

	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func ToSqlite(filePath string, fileTarget string) error {
	print(filePath)

	_, err := os.ReadFile(filePath)

	if err != nil {
		return err
	}

	csvFile, err := os.Create(fileTarget)

	if err != nil {
		return err
	}

	defer csvFile.Close()

	return nil
}
