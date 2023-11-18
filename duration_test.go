package uptime

import (
	"strings"
	"testing"
	"time"
)

func TestParse(t *testing.T) {
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
			txt:    "one year",
			period: "2022-11-16 22:32:44 to 2023-11-16 22:32:44",
			exp:    "1 year",
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

func TestSince(t *testing.T) {
	d := Since(time.Now())
	got := d.String()
	if !strings.Contains(got, "") {
		t.Fail()
	}
}
