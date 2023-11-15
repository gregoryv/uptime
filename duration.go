package uptime

import (
	"fmt"
	"strings"
	"time"
)

// Approximate returns the duration between a and b. Any year is
// considered to be 365 days, any month is 30 days. This method is
// faster though not as accurate.
func Approximate(a, b time.Time) *Duration {
	// a should always come before b
	if b.Before(a) {
		a, b = b, a
	}
	s := b.Sub(a)
	var d Duration
	// years
	year := 365 * day
	years := s.Truncate(year) / year
	d.Years = int(years)
	s -= years * year

	// months
	month := 30 * day
	months := s.Truncate(month) / month
	d.Months = int(months)
	s -= months * month

	// days
	days := s.Truncate(day) / day
	d.Days = int(days)
	s -= days * day

	// hours
	h := s.Truncate(time.Hour).Hours()
	d.Hours = int(h)
	m := time.Duration(s - s.Truncate(time.Hour)).Minutes()
	d.Minutes = int(m)
	sec := time.Duration(s - s.Truncate(time.Minute)).Seconds()
	d.Seconds = int(sec)
	return &d
}

// Between returns the absolute duration between a and b.
func Between(a, b time.Time) *Duration {
	// a should always come before b
	if b.Before(a) {
		a, b = b, a
	}
	var years, months, days int
	tmp := a
	aDay := a.Day()
	chunk := func(i int) {
		for {
			next := tmp.Add(day)
			d := next.Day()
			i--
			days++

			if aDay == d || d == 1 && days > 28 {
				months++
				// remove number of days of passed month
				days -= tmp.Day()
				if months == 12 {
					years++
					months = 0
				}

				// skip ahead if there are enough days
				if v := 27; i > v {
					i -= v
					days += v
					next = next.Add(time.Duration(v) * day)
				}
			}
			tmp = next
			if i == 0 {
				break
			}
		}
	}
	for j := a.Year(); j < b.Year()-100; j++ {
		chunk(365)
	}
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
