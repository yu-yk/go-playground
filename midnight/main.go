package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(isMidNight())
}

func isMidNight() bool {
	layout := "15:04"
	now := time.Now().Format(layout)

	check, _ := time.Parse(layout, now)
	start, _ := time.Parse(layout, "00:00")
	end, _ := time.Parse(layout, "06:30")

	if start.Before(end) {
		return !check.Before(start) && !check.After(end)
	}
	if start.Equal(end) {
		return check.Equal(start)
	}
	return !start.After(check) || !end.Before(check)
}
