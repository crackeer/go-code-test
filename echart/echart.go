package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"sort"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

// generate random data for line chart
func generateLineItems() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < 7; i++ {
		items = append(items, opts.LineData{Value: rand.Intn(300)})
	}
	return items
}

func httpserver(w http.ResponseWriter, _ *http.Request) {

	list := readCsv("exp.csv")

	data := fmtExpData(list)

	// create a new line instance
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Line example in Westeros theme",
			Subtitle: "Line chart rendered by the http server this time",
		}))

	// Put data into instance
	line.SetXAxis([]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}).
		AddSeries("Category A", generateLineItems()).
		AddSeries("Category B", generateLineItems()).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	line.Render(w)
}

func main() {
	http.HandleFunc("/", httpserver)
	http.ListenAndServe(":8081", nil)
}

func fmtExpData(list [][]string) (map[string]map[string]string, []string) {

	retData := map[string]map[string]string{}

	dateMap := map[string]string{}
	for _, val := range list {
		expID := val[1]
		date := val[0]
		click := val[7]
		dateMap[date] = "1"
		if realL, ok := retData[expID]; ok {
			realL[date] = click
			retData[expID] = realL
		} else {
			tmp := map[string]string{}
			tmp[date] = click
			retData[expID] = tmp
		}
	}

	dateList := []string{}
	for d := range dateMap {
		dateList = append(dateList, d)
	}

	sort.Strings(dateList)
	return retData, dateList
}

func readCsv(fileName string) [][]string {

	list := [][]string{}
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
		}
		list = append(list, record)
	}
	return list
}
