package uptime

import (
	"fmt"
	"testing"
	"time"
)

func TestBetween(t *testing.T) {
	cases := []struct {
		t string // text description
		a string
		b string
		s string // short
		l string // long
	}{
		{
			t: "zero",
			a: "2021-01-01 12:00:00",
			b: "2021-01-01 12:00:00",
			s: "0y0m0d 0h0m0s",
			l: "",
		},
		{
			t: "one day",
			a: "2021-01-01 12:00:00",
			b: "2021-01-02 12:00:00",
			s: "0y0m1d 0h0m0s",
			l: "1 day",
		},
		{
			t: "two days",
			a: "2021-01-01 12:00:00",
			b: "2021-01-03 12:00:00",
			s: "0y0m2d 0h0m0s",
			l: "2 days",
		},
		{
			t: "13 months",
			a: "2021-01-01 12:00:00",
			b: "2022-02-01 12:00:00",
			s: "1y1m0d 0h0m0s",
			l: "1 year 1 month",
		},
	}
	for _, c := range cases {
		t.Run(c.t, func(t *testing.T) {
			a, err := time.Parse("2006-01-02 15:04:05", c.a)
			if err != nil {
				t.Fatal(err)
			}
			b, err := time.Parse("2006-01-02 15:04:05", c.b)
			if err != nil {
				t.Fatal(err)
			}
			dur := Between(a, b)
			if got := dur.Short(); got != c.s {
				t.Log("got", got)
				t.Error("exp", c.s)
			}
			if got := dur.String(); got != c.l {
				t.Log("got", got)
				t.Error("exp", c.l)
			}
		})
	}
}

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
