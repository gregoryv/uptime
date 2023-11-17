package uptime

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
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
			d.add(sinceNewYear(a), daysInMonth(Y, M))
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

func TestBetween(t *testing.T) {
	defer log.SetOutput(ioutil.Discard)
	cases := []struct {
		t string // text description
		a string
		b string
		s string // short
		l string // long
		D bool   // debug
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
		{
			t: "jan to march",
			a: "2022-01-30 12:00:00",
			b: "2022-03-10 12:00:00",
			s: "0y1m11d 0h0m0s",
			l: "1 month 11 days",
		},
		{
			t: "thousand years",
			a: "1022-01-01 12:00:00",
			b: "2022-01-01 12:00:00",
			s: "1000y0m0d 0h0m0s",
			l: "1000 years",
		},
		{
			t: "middle of month",
			a: "2022-01-15 12:00:00",
			b: "2022-03-15 12:00:00",
			s: "0y2m0d 0h0m0s",
			l: "2 months",
			D: true,
		},
		{
			t: "middle of month",
			a: "2022-01-15 12:00:00",
			b: "2022-03-14 12:00:00",
			s: "0y1m27d 0h0m0s",
			l: "1 month 27 days",
		},
		{
			t: "feb",
			a: "2022-02-01 00:00:00",
			b: "2024-03-01 00:00:00",
			s: "2y1m0d 0h0m0s",
			l: "2 years 1 month",
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
			log.SetOutput(ioutil.Discard)
			if c.D {
				log.SetOutput(os.Stderr)
			}
			dur := Between(a, b)
			if got := dur.Short(); got != c.s {
				t.Log("got", got)
				t.Fatal("exp", c.s)
			}
			if got := dur.String(); got != c.l {
				t.Log("got", got)
				t.Fatal("exp", c.l)
			}
		})
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
