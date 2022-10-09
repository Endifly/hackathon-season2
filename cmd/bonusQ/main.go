package main

import (
	"fmt"
	"net/http"

	"github.com/markkj/hackathon-season2/internal/sqlite"
	"github.com/markkj/hackathon-season2/internal/visualization"
)

var monthNameMap = map[int]string{
	1:  "January",
	2:  "February",
	3:  "March",
	4:  "April",
	5:  "May",
	6:  "June",
	7:  "July",
	8:  "August",
	9:  "September",
	10: "October",
	11: "November",
	12: "December",
}
var genderMap = map[int]string{
	0: "Male",
	1: "Female",
}

var statusMap = map[int]string{
	1: "Active",
	2: "Resigned",
	3: "Retired",
}

func NewSqManager() sqlite.SqlManager {
	sql := sqlite.SqlManager{}
	sql.OpenDB("./devMountain2.sqlite")

	return sql
}

// TODO: feed the correct data and display this one
// SELECT COUNT(*),STRFTIME('%Y', hired) AS year
// FROM DevMountainAnwser
// GROUP by year
func addVisualizationForHiredCountByYear(w http.ResponseWriter) {

	chart := visualization.Chart{
		Title:    "Employee hired by year",
		Subtitle: "showing new hired employee count per year",
		XAxis:    []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"},
		Items: map[string]interface{}{
			"key":  []int{1, 3, 5, 7, 3, 4, 5},
			"key2": []int{2, 4, 2, 10, 2, 4, 5},
		},
	}
	chart.CreateLineChart(w)
}

// TODO: feed the correct data and display this one
// SELECT COUNT(*),STRFTIME('%m', birthday) AS month
// FROM DevMountainAnwser
// GROUP by month
func addVisualizationForBirithdayByMonth(w http.ResponseWriter) {
	chart := visualization.Chart{
		Title:    "Employee birithday By month",
		Subtitle: "showing employee birithday group by month",
		XAxis:    []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"},
		Items: map[string]interface{}{
			"key":  []int{1, 3, 5, 7, 3, 4, 5},
			"key2": []int{2, 4, 2, 10, 2, 4, 5},
		},
	}
	chart.CreateLineChart(w)
}

// TODO: feed the correct data and display this one
// SELECT COUNT(*), STRFTIME('%Y','now') - STRFTIME('%Y', birthday) AS age
// FROM DevMountainAnwser
// GROUP by Age
func addVisualizationEmployeeAge(w http.ResponseWriter) {
	chart := visualization.Chart{
		Title:    "Employee Age",
		Subtitle: "showing employee group by age",
		XAxis:    []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"},
		Items: map[string]interface{}{
			"key":  []int{1, 3, 5, 7, 3, 4, 5},
			"key2": []int{2, 4, 2, 10, 2, 4, 5},
		},
	}
	chart.CreateLineChart(w)
}

// TODO: feed the correct data and display this one
// SELECT COUNT(*), gender
// FROM DevMountainAnwser
// GROUP by gender
func addVisualizationEmployeeGender(w http.ResponseWriter) {
	chart := visualization.Chart{
		Title:    "Employee Gender",
		Subtitle: "showing employee group by gender",
		XAxis:    []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"},
		Items: map[string]interface{}{
			"key":  []int{1, 3, 5, 7, 3, 4, 5},
			"key2": []int{2, 4, 2, 10, 2, 4, 5},
		},
	}
	chart.CreateLineChart(w)
}

// TODO: feed the correct data and display this one
// SELECT COUNT(*), dept
// FROM DevMountainAnwser
// GROUP by dept
func addVisualizationEmployeeDepartment(w http.ResponseWriter) {
	chart := visualization.Chart{
		Title:    "Employee Department",
		Subtitle: "showing employee group by department",
		XAxis:    []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"},
		Items: map[string]interface{}{
			"key":  []int{1, 3, 5, 7, 3, 4, 5},
			"key2": []int{2, 4, 2, 10, 2, 4, 5},
		},
	}
	chart.CreateLineChart(w)
}

// TODO: feed the correct data and display this one
// SELECT COUNT(*), region
// FROM DevMountainAnwser
// GROUP by region
func addVisualizationEmployeeRegion(w http.ResponseWriter) {
	chart := visualization.Chart{
		Title:    "Employee Region",
		Subtitle: "showing employee group by Region",
		XAxis:    []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"},
		Items: map[string]interface{}{
			"key":  []int{1, 3, 5, 7, 3, 4, 5},
			"key2": []int{2, 4, 2, 10, 2, 4, 5},
		},
	}
	chart.CreateLineChart(w)
}

// TODO: feed the correct data and display this one
// SELECT COUNT(*),status, dept
// FROM DevMountainAnwser
// GROUP by dept
func addVisualizationEmployeeStatusByDept(w http.ResponseWriter) {
	chart := visualization.Chart{
		Title:    "Employee Status count by Department",
		Subtitle: "showing employee status count group by department",
		XAxis:    []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"},
		Items: map[string]interface{}{
			"key":  []int{1, 3, 5, 7, 3, 4, 5},
			"key2": []int{2, 4, 2, 10, 2, 4, 5},
		},
	}
	chart.CreateLineChart(w)
}

// TODO: remove this one
func addVis1(w http.ResponseWriter) {
	// sql := NewSqManager()

	// data := sql.DB.Query("")

	// fmt.Pr
	// sql := sqlite.OpenDB("./devMountain2.sqlite")

	// sql.QueryVis1()

	chart := visualization.Chart{
		Title:    "Waiting",
		Subtitle: "Stil waiting",
		XAxis:    []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"},
		Items: map[string]interface{}{
			"key":  []int{1, 3, 5, 7, 3, 4, 5},
			"key2": []int{2, 4, 2, 10, 2, 4, 5},
		},
	}
	chart.CreateLineChart(w)
}

// TODO: remove this one
func addVis2(w http.ResponseWriter) {
	chart := visualization.Chart{
		Title:    "Waiting",
		Subtitle: "Stil waiting",
		Items: map[string]interface{}{
			"key": map[string]int{
				"test":  20,
				"test2": 32,
				"test3": 12,
				"test4": 52,
			},
		},
	}
	chart.CreatePieChart(w)
}

// TODO: remove this one
func addVis3(w http.ResponseWriter) {
	chart := visualization.Chart{
		Title:    "Waiting",
		Subtitle: "Stil waiting",
		XAxis:    []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"},
		Items: map[string]interface{}{
			"key":  []int{1, 3, 5, 7, 3, 4, 5},
			"key2": []int{2, 4, 2, 10, 2, 4, 5},
		},
	}
	chart.CreateScatter(w)
}

// TODO: remove this one
func addVis4(w http.ResponseWriter) {
	sql := sqlite.OpenDB("./devMountain2.sqlite")

	fmt.Println(sql)

	data := sql.QueryVis5()

	fmt.Println(data)

	chart := visualization.Chart{
		Title:    "Waiting",
		Subtitle: "Stil waiting",
		Items: map[string]interface{}{
			"Thailand": 63,
			"Russia":   200,
			"Canada":   45,
		},
		// Items: data,
	}
	chart.CreateWorldMap(w)
}

// TODO: remove this one
func addVis5(w http.ResponseWriter) {

	chart := visualization.Chart{
		Title:    "Waiting",
		Subtitle: "Stil waiting",
		XAxis:    []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"},
		Items: map[string]interface{}{
			"key":  []int{1, 3, 5, 7, 3, 4, 5},
			"key2": []int{2, 4, 2, 10, 2, 4, 5},
		},
	}
	chart.CreatebarChart(w)
}

func httpserver(w http.ResponseWriter, _ *http.Request) {
	addVis1(w)
	addVis2(w)
	addVis3(w)
	addVis4(w)
	addVis5(w)
}

func main() {
	http.HandleFunc("/", httpserver)
	fmt.Println("Starting http server at port :8081")
	http.ListenAndServe(":8081", nil)
}
