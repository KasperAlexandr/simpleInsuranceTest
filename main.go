package main

import (
	"fmt"
	"log"
	"net/http"
	"simpleInsuranceTest/requestCounter"
	"time"
)

func main() {

	rc := requestCounter.NewRequestCounter(60 * time.Second)

	http.HandleFunc("/request-count", func(w http.ResponseWriter, r *http.Request) {
		rc.AddRequest()
		fmt.Fprintf(w, "Request window: %v\nNumber of requests: %v", rc.Window, rc.CountRequestsInWindow())
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
