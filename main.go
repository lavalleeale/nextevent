package main

import (
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/apognu/gocal"
)

func DateEqual(date1, date2 time.Time) bool {
	y1, m1, d1 := date1.Date()
	y2, m2, d2 := date2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./calendar <calendar.ics>")
		os.Exit(1)
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	c := gocal.NewParser(f)
	c.SkipBounds = true
	c.Parse()
	events := make([]gocal.Event, 0)

	for _, e := range c.Events {
		for _, ev := range c.ExpandRecurringEvent(&e) {
			if DateEqual(*ev.Start, time.Now()) {
				events = append(events, ev)
			}
		}
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].Start.Before(*events[j].Start)
	})

	for _, e := range events {
		if e.Start.After(time.Now().Add(time.Minute * -5)) {
			untilStart := e.Start.Sub(time.Now()).Round(time.Minute).String()
			untilStart = untilStart[:len(untilStart)-2]
			fmt.Printf("%s in %s\n", e.Summary, untilStart)
			break
		}
	}
}
