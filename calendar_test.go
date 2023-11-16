package uptime

import (
	"testing"
	"time"
)

func TestCalendar_Day(t *testing.T) {
	cases := []struct {
		t string
		v string // yyyy-mm
		e int
	}{
		{
			t: "go time",
			v: "2006-02",
			e: 28,
		},
		{
			t: "go time again",
			v: "2006-02",
			e: 28,
		},
		{
			t: "happy nineties",
			v: "1990-01",
			e: 31,
		},
	}
	cal := NewCalendar()
	for _, c := range cases {
		t.Run(c.t, func(t *testing.T) {
			a, err := time.Parse("2006-01", c.v)
			if err != nil {
				t.Fatal(err)
			}
			got := cal.Days(a.Year(), a.Month())
			if got != c.e {
				t.Log(a)
				t.Log("got", got)
				t.Error("exp", c.e)
			}
		})
	}
}
