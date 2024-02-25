package middleware

import (
	"fmt"
	"net/http"
	"simpleInsuranceTest/requestCounter"
)

func RequestCounterMiddleware(rc *requestCounter.RequestCounter, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rc.AddRequest()
		w.Write([]byte(fmt.Sprintf("Request window: %v\nNumber of requests: %v\n", rc.Window, rc.CountRequestsInWindow())))
		next.ServeHTTP(w, r)
	})
}
