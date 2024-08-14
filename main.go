package main

import (
	"fmt"
	"time"
)

const ()

type Status string

const (
	Operational         Status = "Android"
	DegradedPerformance Status = "IOS"
	PartialOutage       Status = "MacOS"
	MajorOutage         Status = "Windows"
	Maintenance         Status = "maintenance"
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

type Incident struct {
	Error    uint
	Protocol string
	FirstLog time.Time
	LastLog  time.Time
	Length   uint
}

func main() {
	fmt.Println("Site Uptime Alerts")
}
