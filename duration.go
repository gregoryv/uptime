/*
Package uptime calculates long durations.

This package solves two problems;

  - overcomes the limitation of [time.Duration], approximately 290 years.
  - format readable [Duration]
*/
package uptime

import (
	"fmt"
	"strings"
	"time"
)

func NewDuration(d time.Duration) *Duration {
	b := time.Now()
	a := b.Add(-d)
	return Between(a, b)
}

// Between returns the absolute duration between a and b.
func Between(a, b time.Time) *Duration {
	// a should always come before b
	if b.Before(a) {
		a, b = b, a
	}
	if a.Equal(b) {
		return &Duration{}
	}
	var years, months, days int
	tmp := a
	aDay := a.Day()
	var monthDays int
	chunk := func(i int) {
		for {
			next := tmp.Add(day)
			d := next.Day()
			if d == 1 {
				monthDays = tmp.Day()
			}
			i--
			days++

			if aDay == d || d == 1 && days > 28 {
				//log.Println("aDay", aDay, monthDays)
				months++
				// remove number of days of passed month
				days -= monthDays
				if months == 12 {
					years++
					months = 0
				}

			}
			//log.Println("i", i, "days", days, "d", d)
			tmp = next
			if i == 0 {
				break
			}
		}
	}
	for j := a.Year(); j < b.Year()-2; j++ {
		chunk(365)
	}
	// remaining days
	i := int(b.Sub(tmp).Truncate(day) / day)
	chunk(i)
	d := &Duration{
		Years:  years,
		Months: months,
		Days:   days,
	}
	s := b.Sub(tmp)
	h := s.Truncate(time.Hour).Hours()
	d.Hours = int(h)
	m := time.Duration(s - s.Truncate(time.Hour)).Minutes()
	d.Minutes = int(m)
	sec := time.Duration(s - s.Truncate(time.Minute)).Seconds()
	d.Seconds = int(sec)
	return d
}

const day = time.Hour * 24

// Duration represents long duration. The duration is the total of all
// fields combined.
type Duration struct {
	Years   int
	Months  int
	Days    int
	Hours   int
	Minutes int
	Seconds int
}

// Short returns an abbreviated duration representation.
func (d *Duration) Short() string {
	return fmt.Sprintf("%vy%vm%vd %vh%vm%vs",
		d.Years,
		d.Months,
		d.Days,
		d.Hours,
		d.Minutes,
		d.Seconds,
	)
}

// String returns the duration representation as named parts excluding
// 0 values.
func (d *Duration) String() string {
	var s []string
	if v := d.Years; v > 0 {
		s = append(s, plural(v, "year"))
	}
	if v := d.Months; v > 0 {
		s = append(s, plural(v, "month"))
	}
	if v := d.Days; v > 0 {
		s = append(s, plural(v, "day"))
	}
	if v := d.Hours; v > 0 {
		s = append(s, plural(v, "hour"))
	}
	if v := d.Minutes; v > 0 {
		s = append(s, plural(v, "minute"))
	}
	if v := d.Seconds; v > 0 {
		s = append(s, plural(v, "second"))
	}
	return strings.Join(s, " ")
}

func plural(v int, txt string) string {
	if v == 1 {
		return fmt.Sprintf("%v %s", v, txt)
	}
	return fmt.Sprintf("%v %ss", v, txt)
}
