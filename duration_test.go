package uptime

import (
	"fmt"
	"time"
)

func ExampleDuration_Short() {
	d := Between(
		time.Date(2021, 1, 01, 12, 00, 00, 0, time.UTC),
		time.Date(2022, 3, 02, 13, 10, 20, 0, time.UTC),
	)
	fmt.Println(d.Short())
	// output:
	// 1y2m1d 1h10m20s
}

func ExampleBetween() {
	d := Between(
		time.Date(2022, 1, 01, 12, 00, 00, 0, time.UTC),
		time.Date(2021, 1, 01, 12, 00, 00, 0, time.UTC),
	)
	fmt.Println(d)
	// output:
	// 1 year
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

func ExampleApproximate_january() {
	d := Approximate(
		time.Date(2022, 1, 30, 12, 00, 00, 0, time.UTC),
		time.Date(2022, 3, 10, 12, 00, 00, 0, time.UTC),
	)
	fmt.Println(d)
	// output:
	// 1 month 9 days
}

func ExampleApproximate() {
	d := Approximate(
		time.Date(2022, 3, 10, 12, 00, 00, 0, time.UTC),
		time.Date(2022, 1, 30, 12, 01, 00, 0, time.UTC),
	)
	fmt.Println(d)
	// output:
	// 1 month 8 days 23 hours 59 minutes
}
