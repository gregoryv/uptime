package uptime

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"testing"
	"time"
)

func init() {
	log.SetFlags(0)
	log.SetOutput(ioutil.Discard)
}

func TestSince(t *testing.T) {
	d := Since(time.Now())
	got := d.String()
	if !strings.Contains(got, "") {
		t.Fail()
	}
}

func TestDuration_add(t *testing.T) {
	cases := []struct {
		t string // text
		a string // time
		e string // expected
	}{
		{
			t: "less than a year",
			a: "2023-11-16 22:32:44",
			e: "1 year",
		},
	}

	for _, c := range cases {
		t.Run(c.t, func(t *testing.T) {
			a, err := time.Parse("2006-01-02 15:04:05", c.a)
			if err != nil {
				t.Fatal(err)
			}
			Y, M, _ := a.Date()
			d := untilNewYear(a)
			d = d.add(sinceNewYear(a), daysInMonth(Y, M))
			got := d.String()
			if got != c.e {
				t.Log("got", got)
				t.Fatal("exp", c.e)
			}
		})
	}
}

func TestuntilNewYear(t *testing.T) {
	cases := []struct {
		t string // text
		a string // time
		e string // expected
	}{
		{
			t: "less than a year",
			a: "2023-11-16 22:32:44",
			e: "1 month 14 days 1 hour 27 minutes 16 seconds",
		},
	}

	for _, c := range cases {
		t.Run(c.t, func(t *testing.T) {
			a, err := time.Parse("2006-01-02 15:04:05", c.a)
			if err != nil {
				t.Fatal(err)
			}
			got := untilNewYear(a).String()
			if got != c.e {
				t.Log("got", got)
				t.Fatal("exp", c.e)
			}
		})
	}
}

func TestsinceNewYear(t *testing.T) {
	cases := []struct {
		t string // text
		a string // time
		e string // expected
	}{
		{
			t: "less than a year",
			a: "2023-11-16 22:32:44",
			e: "10 months 15 days 22 hours 32 minutes 44 seconds",
		},
	}

	for _, c := range cases {
		t.Run(c.t, func(t *testing.T) {
			a, err := time.Parse("2006-01-02 15:04:05", c.a)
			if err != nil {
				t.Fatal(err)
			}
			got := sinceNewYear(a).String()
			if got != c.e {
				t.Log("got", got)
				t.Fatal("exp", c.e)
			}
		})
	}
}

func TestParse(t *testing.T) {
	defer log.SetOutput(ioutil.Discard)
	cases := []struct {
		txt    string // text description
		period string
		exp    string // long
	}{
		{
			txt:    "zero",
			period: "2021-01-01 to 2021-01-01",
			exp:    "",
		},
		{
			txt:    "short",
			period: "2021-01-01 12:00:00 to 2021-01-01 12:00:01",
			exp:    "1 second",
		},
		{
			txt:    "one hour",
			period: "2021-01-01 12 to 2021-01-01 13",
			exp:    "1 hour",
		},
		{
			txt:    "one day",
			period: "2021-01-01 to 2021-01-02",
			exp:    "1 day",
		},
		{
			txt:    "two days",
			period: "2021-01-01 to 2021-01-03",
			exp:    "2 days",
		},
		{
			txt:    "13 months",
			period: "2021-01-01 to 2022-02-01",
			exp:    "1 year 1 month",
		},
		{
			txt:    "jan to march",
			period: "2022-01-30 to 2022-03-10",
			exp:    "1 month 11 days",
		},
		{
			txt:    "thousand years",
			period: "1022-01-01 to 2022-01-01",
			exp:    "1000 years",
		},
		{
			txt:    "middle of month",
			period: "2022-01-15 to 2022-03-15",
			exp:    "2 months",
		},
		{
			txt:    "middle of month",
			period: "2022-01-15 to 2022-03-14",
			exp:    "1 month 27 days",
		},
		{
			txt:    "feb",
			period: "2022-02-01 to 2024-03-01",
			exp:    "2 years 1 month",
		},
	}
	for _, c := range cases {
		t.Run(c.txt, func(t *testing.T) {
			format := "2006-01-02 15:04:05"
			dur, err := Parse(format, c.period)
			if err != nil {
				t.Fatal(err)
			}
			if got := dur.String(); got != c.exp {
				t.Log("got", got)
				t.Fatal("exp", c.exp)
			}
		})
	}
}

func TestParse_errorsLeft(t *testing.T) {
	format := "2006-01-02"
	period := "2020/01/03 to 2022/01/03"
	_, err := Parse(format, period)
	if err == nil {
		t.Fatalf("expect Parse error format %q period %q", format, period)
	}
}

func TestParse_errorsRight(t *testing.T) {
	format := "2006-01-02"
	period := "2020-01-03 to 2022/01/03"
	_, err := Parse(format, period)
	if err == nil {
		t.Fatalf("expect Parse error format %q period %q", format, period)
	}
}

func ExampleDuration_Short() {
	a := time.Date(2021, 1, 01, 12, 00, 00, 0, time.UTC)
	b := time.Date(2022, 3, 02, 13, 10, 20, 0, time.UTC)
	fmt.Print(Between(a, b).Short())
	// output:
	// 1y2m1d 1h10m20s
}

func ExampleDuration_String() {
	a := time.Date(2022, 1, 01, 12, 00, 00, 0, time.UTC)
	b := time.Date(2021, 1, 01, 12, 00, 00, 0, time.UTC)
	fmt.Print(Between(a, b))
	// output:
	// 1 year
}

func Example_longDurationBetween() {
	a := time.Date(1021, 1, 01, 12, 00, 00, 0, time.UTC)
	b := time.Date(2022, 3, 07, 16, 00, 00, 0, time.UTC)
	fmt.Print(Between(a, b))
	// output:
	// 1001 years 2 months 6 days 4 hours
}

func ExampleBetween_january() {
	a := time.Date(2022, 1, 30, 12, 00, 00, 0, time.UTC)
	b := time.Date(2022, 3, 10, 12, 00, 00, 0, time.UTC)
	fmt.Println(Between(a, b))
	// output:
	// 1 month 11 days
}
