package requestCounter

import (
	"bytes"
	"encoding/gob"
	"log"
	"os"
	"sync"
	"time"
)

type RequestCounter struct {
	RequestsTime []time.Time
	Window       time.Duration
	Mutex        sync.Mutex
}

func NewRequestCounter(window time.Duration) *RequestCounter {
	requestHistory := getRequestHistoryFromFile()
	return &RequestCounter{Window: window, RequestsTime: requestHistory}
}

func (rc *RequestCounter) AddRequest() {
	rc.Mutex.Lock()
	defer rc.Mutex.Unlock()

	rc.RequestsTime = append(rc.RequestsTime, time.Now())
}

func (rc *RequestCounter) CountRequestsInWindow() int {
	rc.Mutex.Lock()
	defer rc.Mutex.Unlock()

	now := time.Now()
	windowBoundary := now.Add(-rc.Window)

	var requestsInWindow []time.Time
	for _, reqTime := range rc.RequestsTime {
		if reqTime.After(windowBoundary) {
			requestsInWindow = append(requestsInWindow, reqTime)
		}
	}
	rc.RequestsTime = requestsInWindow

	writeToRequestHistoryAFile(rc.RequestsTime)

	return len(rc.RequestsTime)
}

func writeToRequestHistoryAFile(timeArray []time.Time) {
	f, err := os.OpenFile("requestTimeStorage", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	enc := gob.NewEncoder(f)

	err = enc.Encode(timeArray)
	if err != nil {
		log.Println(err)
	}
}

func getRequestHistoryFromFile() (requestHistory []time.Time) {
	rf, err := os.ReadFile("requestTimeStorage")
	if err != nil {
		log.Println(err)
	}

	reader := bytes.NewReader(rf)

	dec := gob.NewDecoder(reader)

	err = dec.Decode(&requestHistory)
	if err != nil {
		log.Println(err)
	}

	return
}
