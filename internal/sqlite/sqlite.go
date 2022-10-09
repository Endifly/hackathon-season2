package sqlite

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/markkj/hackathon-season2/internal/csv"
	_ "github.com/mattn/go-sqlite3"
)

/*
SqlField type builder
*/
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

func (f *SqlField) Date() *SqlField {
	f.addSchema("date")
	return f
}

func (f *SqlField) Schema() string {
	schem := strings.Join(f.schema, " ")
	col := fmt.Sprintf("%s ", f.colName)
	// return strings.Join(f.schema, " ")
	return col + schem
}

// SqlField builder ...
func Field(name string) *SqlField {
	f := &SqlField{}
	f.Name(name)
	return f
}

type SqlData struct {
	Conn      *sql.DB
	tableName string
	tableCols []string
}

func (c *SqlData) UseTable(tableName string) {
	c.tableName = tableName
}

func (c *SqlData) UseSchema(columns []*SqlField) error {
	cols := ""
	for _, v := range columns {
		c.tableCols = append(c.tableCols, v.colName)
		cols = cols + v.Schema() + ","
	}

	cols = strings.TrimSuffix(cols, ",")

	cmd := fmt.Sprintf("create table if not exists %s (%s)", c.tableName, cols)
	stmt, err := c.Conn.Prepare(cmd)
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
	cols := strings.Join(c.tableCols, ",")
	// vals := strings.Join(data, ",")
	// vals := []interface{}{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	// valsTemplate := ""
	// for i := 0; i < len(data); i++ {
	// 	valsTemplate
	// }

	vals := []interface{}{}
	for _, v := range data {
		vals = append(vals, v)
	}

	cmd := fmt.Sprintf("insert into %s(%s) values(?,?,?,?,?,?,?,?,?,?,?,?)", c.tableName, cols)

	fmt.Println(cmd)
	// stmt, err := c.Conn.Prepare("insert into books(title, author) values(?, ?)")
	stmt, err := c.Conn.Prepare(cmd)
	if err != nil {
		// return err
		fmt.Println(err)
	}

	_, err = stmt.Exec(vals...)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

/*
alway return db connection
*/
func OpenDB(filePath string) SqlData {
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
	db := OpenDB(filePath)
	defer db.Conn.Close()

	if db.Conn == nil {
		fmt.Println("erro")
	}

	db.UseTable("devMountain2")

	err := db.UseSchema([]*SqlField{
		Field("EMPID").Int().PrimaryKey(),
		Field("PASSPORT").Text(),
		Field("FIRSTNAME").Text(),
		Field("LASTNAME").Text(),
		Field("GENDER").Int(),
		Field("BIRTHDAY").Date(),
		Field("NATIONALITY").Text(),
		Field("HIRED").Date(),
		Field("DEPT").Text(),
		Field("POSITION").Text(),
		Field("STATUS").Int(),
		Field("REGION").Text(),
	})

	if err != nil {
		fmt.Println(err)
	}

	employees, err := csv.CSVFileToList("../../output/clean.csv")

	if err != nil {
		fmt.Println(err)
	}

	for _, employee := range employees {
		db.AddRow(employee)
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
