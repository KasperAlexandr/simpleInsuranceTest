package main

import (
	"log"
	"net/http"
	"simpleInsuranceTest/handler"
	"simpleInsuranceTest/middleware"
	"simpleInsuranceTest/requestCounter"
	"time"
)

func main() {

	rc := requestCounter.NewRequestCounter(60 * time.Second)

	http.Handle("/hello", middleware.RequestCounterMiddleware(rc, http.HandlerFunc(handler.HelloHandler)))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
