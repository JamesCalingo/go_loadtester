package loadtester

import (
	"fmt"
	"log"
	"net/http"
	"slices"
	"strings"
)

func CheckStatus(url string, method string) int {
	var allowed_methods = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	if !slices.Contains(allowed_methods, strings.ToUpper(method)) {
		log.Fatal("This is not a valid HTTP method. Valid methods are \"GET\",  \"POST\", \"PUT\", \"PATCH\", and \"DELETE\" (they do not need to be in all caps)")
	}

	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err, res)
		// os.Exit(1)
	} else if res.StatusCode != 200 {
		fmt.Println(res)
	}
	return res.StatusCode

}
