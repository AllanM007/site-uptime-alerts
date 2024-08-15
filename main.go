package main

import (
	"fmt"
	"net/http"
	"time"
)

const ()

type Status string

const (
	Operational         Status = "Operational"
	DegradedPerformance Status = "DegradedPerformance"
	PartialOutage       Status = "PartialOutage"
	MajorOutage         Status = "MajorOutage"
	Maintenance         Status = "Maintenance"
)

var AllStatuses = []Status{
	Operational,
	DegradedPerformance,
	PartialOutage,
	MajorOutage,
	Maintenance,
}

type Domain struct {
	Url          string
	Intervals    uint
	Status       Status
	LastInterval time.Time
}

type Response struct {
	Status  int
	Latency uint
	Time    time.Time
}

type Incident struct {
	Error    uint
	Protocol string
	FirstLog time.Time
	LastLog  time.Time
	Length   uint
}

func (t Domain) SendPing() error {
	resp, err := http.Get(t.Url)
	if err != nil {
		return err
	}

	var newResponse Response
	newResponse.Status = resp.StatusCode

	return nil
}

func main() {
	fmt.Println("Site Uptime Alerts")
}
