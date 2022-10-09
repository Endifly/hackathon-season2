package main

import (
	"fmt"
	"net/http"

	"github.com/markkj/hackathon-season2/internal/visualization"
)

func addVisualizationForHiredPerYear(w http.ResponseWriter) {
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
func addVisualizationForBirithdayByYear(w http.ResponseWriter) {
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
