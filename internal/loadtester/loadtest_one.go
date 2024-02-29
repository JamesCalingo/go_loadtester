package loadtester

import (
	"fmt"
	"net/http"
)

func CheckStatus(url string, attempts int) string {
	var successes int
	var failures int
	for i := 0; i < attempts; i++ {
		res, err := http.Get(url)
		if err != nil {
			fmt.Println(res)
			// os.Exit(1)
		} else if res.StatusCode != 200 {
			fmt.Println(res.Status)
			failures += 1
		} else {
			successes += 1
		}
	}
	return fmt.Sprintf("Successes: %d \nFailures: %d", successes, failures)
}
