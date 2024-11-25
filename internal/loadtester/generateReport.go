package loadtester

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type report struct {
	Successes int
	Failures  int
	//I'm keeping the times as times here, but you can change them if you think making them strings looks better
	Start   time.Time
	End     time.Time
	Elapsed time.Duration
}

func GenerateReport(url string, method string, attempts int) string {
	var report report
	httpMethod := strings.ToUpper(method)
	// start := time.Now()
	report.Start = time.Now()
	for i := 0; i < attempts; i++ {
		switch httpMethod {
		case "GET":
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
		case "POST":
			// res, err := http.Post(url, "", nil)
			return "In progress"
		case "PUT", "PATCH":
			return "Go's HTTP module is not quite for this"
		case "DELETE":
			return "DELETED!"
		}
	}
	// end := time.Now()
	report.End = time.Now()
	report.Elapsed = time.Duration((report.End.Sub(report.Start)).Milliseconds())
	// This is to make the report look "pretty" when it's logged; we turn it into JSON, then convert that back to a string
	obj, _ := json.Marshal(report)
	printedReport := string(obj)
	// Changing the commas to new lines isn't necessary - I just think it looks better
	printedReport = strings.ReplaceAll(printedReport, ",", ",\n")
	return strings.ReplaceAll(printedReport, ":", ": ")
}
