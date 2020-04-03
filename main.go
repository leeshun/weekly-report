package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/leeshun/weekly-report/weekly_report"
)

var (
	sharedURL = flag.String("shared_url", "", "shared url")
)

func main() {
	flag.Parse()
	fmt.Println("***begin to fetch weekly report data***")
	data, err := weekly_report.FetchWeeklyReportData(*sharedURL)
	if err != nil {
		fmt.Println("failed to fetch weekly report data", err)
		return
	}
	now := time.Now()
	_, week := now.ISOWeek()
	data = fmt.Sprintf("**%d.%d.%d W%d\n%s**", now.Year(), now.Month(), now.Day(), week, data)
	fmt.Println("***begin to post weekly report data***")
	err = weekly_report.Post(data)
	if err != nil {
		fmt.Println("failed to post weekly report data", err)
	} else {
		fmt.Println("success!")
	}
}
