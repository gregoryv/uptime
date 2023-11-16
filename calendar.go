package uptime

import "time"

func NewCalendar() *Calendar {
	return &Calendar{
		cache: make(map[int]map[time.Month]int),
	}
}

type Calendar struct {
	cache map[int]map[time.Month]int
}

func (c *Calendar) Days(year int, m time.Month) int {
	return c.calc(year)[m]
}

func (c *Calendar) calc(year int) map[time.Month]int {
	if v, found := c.cache[year]; found {
		return v
	}
	for m := time.January; m <= time.December; m++ {
		d := time.Date(year, m, 1, 0, 0, 0, 0, time.UTC)
		d = d.AddDate(0, 1, -1)
		if m == time.January {
			c.cache[year] = make(map[time.Month]int)
		}
		c.cache[year][m] = d.Day()
	}
	return c.cache[year]
}
