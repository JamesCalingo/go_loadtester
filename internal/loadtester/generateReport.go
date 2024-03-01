package loadtester

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"slices"
	"strings"
	"time"
)

type Report struct {
	Successes int
	Failures  int
	// having start and end as times is almost useless IMO, so I make them strings
	Start   string
	End     string
	Elapsed time.Duration
}

func GenerateReport(url string, method string, attempts int) string {
	var report Report
	var allowed_methods = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	if !slices.Contains(allowed_methods, strings.ToUpper(method)) {
		log.Fatal("This is not a valid HTTP method.")
	}
	start := time.Now()
	report.Start = start.String()
	for i := 0; i < attempts; i++ {
		res, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			// break
			//Conveniently, all of our failed statuses are 400 and above
		} else if res.StatusCode >= 400 {
			fmt.Println(res)
			report.Failures += 1
		} else {
			report.Successes += 1
		}
	}
	end := time.Now()
	report.End = end.String()
	report.Elapsed = time.Duration((end.Sub(start)).Milliseconds())
	obj, _ := json.Marshal(report)
	printedReport := string(obj)
	return printedReport
}
