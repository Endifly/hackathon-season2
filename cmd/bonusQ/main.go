package main

import (
	"fmt"
	"net/http"

	// "github.com/go-echarts/go-echarts/v2/charts"
	// "github.com/go-echarts/go-echarts/v2/opts"
	// "github.com/go-echarts/go-echarts/v2/types"
	"github.com/markkj/hackathon-season2/internal/visualization"
)

func addVis1(w http.ResponseWriter) {
	chart := visualization.Chart{
		XAxis: []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"},
		Items: map[string]interface{}{
			"key":  []int{1, 3, 5, 7, 3, 4, 5},
			"key2": []int{2, 4, 2, 10, 2, 4, 5},
		},
	}
	chart.CreateLineChart(w)
}

// func addVis2() {
// 	chart := visualization.Chart{
// 		XAxis: []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"},
// 	}
// }

func httpserver(w http.ResponseWriter, _ *http.Request) {
	addVis1(w)
}

func main() {
	http.HandleFunc("/", httpserver)
	fmt.Println("Starting http server at port :8081")
	http.ListenAndServe(":8081", nil)
}
