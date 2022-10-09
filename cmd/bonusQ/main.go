package main

import (
	"fmt"
	"net/http"

	"github.com/markkj/hackathon-season2/internal/visualization"
)


//TODO: feed the correct data and display this one
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


//TODO: feed the correct data and display this one
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

//TODO: feed the correct data and display this one
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
//TODO: feed the correct data and display this one
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
//TODO: feed the correct data and display this one
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

//TODO: feed the correct data and display this one
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

//TODO: feed the correct data and display this one
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

//TODO: remove this one 
func addVis1(w http.ResponseWriter) {
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

//TODO: remove this one 
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

//TODO: remove this one 
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

//TODO: remove this one 
func addVis4(w http.ResponseWriter) {
	chart := visualization.Chart{
		Title:    "Waiting",
		Subtitle: "Stil waiting",
		Items: map[string]interface{}{
			"Thailand": 63,
			"Russia":   200,
			"Canada":   45,
		},
	}
	chart.CreateWorldMap(w)
}

//TODO: remove this one 
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
