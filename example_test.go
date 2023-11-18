package uptime_test

import (
	"fmt"
	"time"

	"github.com/gregoryv/uptime"
)

func ExampleDuration() {
	a := time.Date(1022, 1, 01, 12, 00, 00, 0, time.UTC)
	b := time.Date(2022, 3, 02, 13, 10, 20, 0, time.UTC)
	d := uptime.Between(a, b)
	fmt.Println(d.String())
	fmt.Println(d.Short())
	// output:
	// 1000 years 2 months 1 day 1 hour 10 minutes 20 seconds
	// 1000y2m1d 1h10m20s
}

func ExampleParse_partialFormat() {
	dur, _ := uptime.Parse("2006-01-02 15:04:05", "1990-01-01 to 1991-01-02")
	fmt.Print(dur)
	// output:
	// 1 year 1 day
}

func ExampleDuration_Short() {
	a := time.Date(2021, 1, 01, 12, 00, 00, 0, time.UTC)
	b := time.Date(2022, 3, 02, 13, 10, 20, 0, time.UTC)
	fmt.Print(uptime.Between(a, b).Short())
	// output:
	// 1y2m1d 1h10m20s
}

func ExampleDuration_String() {
	a := time.Date(2022, 1, 01, 12, 00, 00, 0, time.UTC)
	b := time.Date(2021, 1, 01, 12, 00, 00, 0, time.UTC)
	fmt.Print(uptime.Between(a, b))
	// output:
	// 1 year
}

func Example_longDurationBetween() {
	a := time.Date(1021, 1, 01, 12, 00, 00, 0, time.UTC)
	b := time.Date(2022, 3, 07, 16, 00, 00, 0, time.UTC)
	fmt.Print(uptime.Between(a, b))
	// output:
	// 1001 years 2 months 6 days 4 hours
}

func ExampleBetween_january() {
	a := time.Date(2022, 1, 30, 12, 00, 00, 0, time.UTC)
	b := time.Date(2022, 3, 10, 12, 00, 00, 0, time.UTC)
	fmt.Println(uptime.Between(a, b))
	// output:
	// 1 month 11 days
}
