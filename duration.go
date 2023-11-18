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
func Since(a time.Time) Duration {
	return Between(a, time.Now())
}

// Parse period string "a to b", where format of b must match a. It is
// ok for a and b to use only part of format as long as they are
// equal.
func Parse(format, period string) (Duration, error) {
	a, b, err := parseTimes(format, period)
	if err != nil {
		return Duration{}, err
	}
	return Between(a, b), nil
}

func parseTimes(format, period string) (a, b time.Time, err error) {
	sep := " to "
	i := strings.Index(period, sep)
	// left part dictates format
	format = format[:i]

	a, err = time.Parse(format, period[:i])
	if err != nil {
		return
	}
	b, err = time.Parse(format, period[i+len(sep):])
	if err != nil {
		return
	}
	return
}

// Between returns the absolute duration between a and b.
func Between(a, b time.Time) Duration {
	// a should always come before b
	if b.Before(a) {
		a, b = b, a
	}
	a = a.Truncate(time.Second)
	b = b.Truncate(time.Second)
	if a.Equal(b) {
		return Duration{}
	}

	if years := b.Year() - a.Year(); years > 0 {
		dur := untilNewYear(a)
		Y, M, _ := a.Date()
		dur = dur.add(sinceNewYear(b), daysInMonth(Y, M))
		dur[0] += years - 1
		return dur
	}

	// less than a year
	var years, months, days int
	tmp := a
	aDay := a.Day()
	var monthDays int

	for i := int(b.Sub(tmp).Truncate(day) / day); i > 0; i-- {
		next := tmp.Add(day)
		d := next.Day()
		if d == 1 {
			monthDays = tmp.Day()
		}
		days++

		if aDay == d || d == 1 && days > 28 {
			//log.Println("aDay", aDay, monthDays)
			months++
			// remove number of days of passed month
			days -= monthDays
		}
		//log.Println("i", i, "days", days, "d", d)
		tmp = next
	}

	d := Duration{years, months, days}
	s := b.Sub(tmp)
	d = d.setHourMinSec(s)
	return d
}

const day = time.Hour * 24

// untilNewYear returns duration before new years
func untilNewYear(t time.Time) Duration {
	y, m, d := t.Date()
	dur := Duration{
		0,
		12 - int(m),
		daysInMonth(y, m) - d,
	}
	h, mm, s := t.Clock()
	hms := 24*time.Hour -
		time.Duration(h)*time.Hour -
		time.Duration(mm)*time.Minute -
		time.Duration(s)*time.Second
	dur = dur.setHourMinSec(hms)
	return dur
}

// sinceNewYear returns duration since new years
func sinceNewYear(t time.Time) Duration {
	_, m, d := t.Date()
	h, mm, s := t.Clock()
	return Duration{
		0,
		int(m) - 1,
		d - 1,
		h,
		mm,
		s,
	}
}

// Duration represents long duration. The duration is the total of all
// fields combined.
type Duration [6]int

const (
	iYears = iota
	iMonths
	iDays
	iHours
	iMinutes
	iSeconds
)

// Years returns years part of the duration
func (d Duration) Years() int   { return d[0] }
// Months returns months part of the duration
func (d Duration) Months() int  { return d[1] }
// Days returns days part of the duration
func (d Duration) Days() int    { return d[2] }
// Hour returns hour part of the duration
func (d Duration) Hours() int   { return d[3] }
// Minutes returns minutes part of the duration
func (d Duration) Minutes() int { return d[4] }
// Seconds returns seconds part of the duration
func (d Duration) Seconds() int { return d[5] }

func (d Duration) setHourMinSec(s time.Duration) Duration {
	h := s.Truncate(time.Hour).Hours()
	d[iHours] = int(h)
	m := time.Duration(s - s.Truncate(time.Hour)).Minutes()
	d[iMinutes] = int(m)
	sec := time.Duration(s - s.Truncate(time.Minute)).Seconds()
	d[iSeconds] = int(sec)
	return d
}

func (d Duration) add(v Duration, monthDays int) Duration {
	d[iYears] += v[iYears]
	d[iMonths] += v[iMonths]
	d[iDays] += v[iDays]
	d[iHours] += v[iHours]
	d[iMinutes] += v[iMinutes]
	d[iSeconds] += v[iSeconds]

	if d[iSeconds] > 59 {
		d[iMinutes]++
		d[iSeconds] -= 60
	}
	if d[iMinutes] > 59 {
		d[iHours]++
		d[iMinutes] -= 60
	}
	if d[iHours] > 23 {
		d[iDays]++
		d[iHours] -= 24
	}
	if d[iDays] >= monthDays {
		d[iMonths]++
		d[iDays] -= monthDays
	}
	if d[iMonths] > 11 {
		d[iYears]++
		d[iMonths] -= 12
	}
	return d
}

// Short returns an abbreviated duration representation.
func (d Duration) Short() string {
	return fmt.Sprintf("%vy%vm%vd %vh%vm%vs",
		d[iYears],
		d[iMonths],
		d[iDays],
		d[iHours],
		d[iMinutes],
		d[iSeconds],
	)
}

// String returns the duration representation as named parts excluding
// 0 values.
func (d Duration) String() string {
	var s []string
	if v := d.Years(); v > 0 {
		s = append(s, plural(v, "year"))
	}
	if v := d.Months(); v > 0 {
		s = append(s, plural(v, "month"))
	}
	if v := d.Days(); v > 0 {
		s = append(s, plural(v, "day"))
	}
	if v := d.Hours(); v > 0 {
		s = append(s, plural(v, "hour"))
	}
	if v := d.Minutes(); v > 0 {
		s = append(s, plural(v, "minute"))
	}
	if v := d.Seconds(); v > 0 {
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

func daysInMonth(year int, m time.Month) int {
	switch m {
	case time.February:
		d := time.Date(year, m, 1, 0, 0, 0, 0, time.UTC)
		d = d.AddDate(0, 1, -1)
		return d.Day()

	case time.April, time.June, time.September, time.November:
		return 30

	default:
		return 31
	}
}
