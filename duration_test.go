package uptime

import (
	"fmt"
	"time"
)

func ExampleDuration_Short() {
	a := time.Date(2021, 1, 01, 12, 00, 00, 0, time.UTC)
	b := time.Date(2022, 3, 02, 13, 10, 20, 0, time.UTC)
	fmt.Print(a, "\n", b, "\n", Between(a, b).Short())
	// output:
	// 2021-01-01 12:00:00 +0000 UTC
	// 2022-03-02 13:10:20 +0000 UTC
	// 1y2m1d 1h10m20s
}

func ExampleDuration_String() {
	a := time.Date(2022, 1, 01, 12, 00, 00, 0, time.UTC)
	b := time.Date(2021, 1, 01, 12, 00, 00, 0, time.UTC)
	fmt.Print(a, "\n", b, "\n", Between(a, b))
	// output:
	// 2022-01-01 12:00:00 +0000 UTC
	// 2021-01-01 12:00:00 +0000 UTC
	// 1 year
}

func Example_longDurationBetween() {
	a := time.Date(1021, 1, 01, 12, 00, 00, 0, time.UTC)
	b := time.Date(2022, 3, 07, 16, 00, 00, 0, time.UTC)
	fmt.Print(a, "\n", b, "\n", Between(a, b))
	// output:
	// 1021-01-01 12:00:00 +0000 UTC
	// 2022-03-07 16:00:00 +0000 UTC
	// 1001 years 2 months 6 days 4 hours
}

func ExampleBetween_january() {
	d := Between(
		time.Date(2022, 1, 30, 12, 00, 00, 0, time.UTC),
		time.Date(2022, 3, 10, 12, 00, 00, 0, time.UTC),
	)
	fmt.Println(d)
	// output:
	// 1 month 11 days
}
