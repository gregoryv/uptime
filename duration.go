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

// Since returns the duration between a and now.
func Since(a time.Time) *Duration {
	return Between(a, time.Now())
}

// Between returns the absolute duration between a and b.
func Between(a, b time.Time) *Duration {
	// a should always come before b
	if b.Before(a) {
		a, b = b, a
	}
	a = a.Truncate(time.Second)
	b = b.Truncate(time.Second)
	if a.Equal(b) {
		return &Duration{}
	}

	if years := b.Year() - a.Year(); years > 0 {
		dur := untilNewYear(a)
		Y, M, _ := a.Date()
		dur.add(sinceNewYear(b), cal.Days(Y, M))
		dur.Years += years - 1
		return dur
	}

	// less than a year
	var years, months, days int
	tmp := a
	aDay := a.Day()
	var monthDays int
	i := int(b.Sub(tmp).Truncate(day) / day)
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
		}
		//log.Println("i", i, "days", days, "d", d)
		tmp = next
		if i == 0 {
			break
		}
	}

	d := &Duration{
		Years:  years,
		Months: months,
		Days:   days,
	}
	s := b.Sub(tmp)
	d.setHourMinSec(s)
	return d
}

const day = time.Hour * 24

var cal = newCalendar()

// untilNewYear returns duration before new years
func untilNewYear(t time.Time) *Duration {
	y, m, d := t.Date()
	dur := &Duration{
		Months: 12 - int(m),
		Days:   cal.Days(y, m) - d,
	}
	h, mm, s := t.Clock()
	hms := 24*time.Hour -
		time.Duration(h)*time.Hour -
		time.Duration(mm)*time.Minute -
		time.Duration(s)*time.Second
	dur.setHourMinSec(hms)
	return dur
}

// sinceNewYear returns duration since new years
func sinceNewYear(t time.Time) *Duration {
	_, m, d := t.Date()
	dur := &Duration{
		Months: int(m) - 1,
		Days:   d - 1,
	}
	h, mm, s := t.Clock()
	hms := time.Duration(h)*time.Hour +
		time.Duration(mm)*time.Minute +
		time.Duration(s)*time.Second
	dur.setHourMinSec(hms)
	return dur
}

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

func (d *Duration) setHourMinSec(s time.Duration) {
	h := s.Truncate(time.Hour).Hours()
	d.Hours = int(h)
	m := time.Duration(s - s.Truncate(time.Hour)).Minutes()
	d.Minutes = int(m)
	sec := time.Duration(s - s.Truncate(time.Minute)).Seconds()
	d.Seconds = int(sec)
}

func (d *Duration) add(v *Duration, monthDays int) {
	d.Years += v.Years
	d.Months += v.Months
	d.Days += v.Days
	d.Hours += v.Hours
	d.Minutes += v.Minutes
	d.Seconds += v.Seconds

	if d.Seconds > 59 {
		d.Minutes++
		d.Seconds -= 60
	}
	if d.Minutes > 59 {
		d.Hours++
		d.Minutes -= 60
	}
	if d.Hours > 23 {
		d.Days++
		d.Hours -= 24
	}
	if d.Days >= monthDays {
		d.Months++
		d.Days -= monthDays
	}
	if d.Months > 11 {
		d.Years++
		d.Months -= 12
	}
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
