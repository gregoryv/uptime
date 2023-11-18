package uptime

import (
	"testing"
	"time"
)

func Benchmark_Between_one_year(b *testing.B) {
	start := time.Date(2021, 11, 6, 11, 34, 13, 0, time.UTC)
	now := time.Date(2023, 3, 19, 11, 34, 13, 0, time.UTC)
	for i := 0; i < b.N; i++ {
		Between(start, now)
	}
}

func Benchmark_Between_thousand_years(b *testing.B) {
	start := time.Date(1021, 11, 6, 11, 34, 13, 0, time.UTC)
	now := time.Date(2023, 3, 19, 11, 34, 13, 0, time.UTC)
	for i := 0; i < b.N; i++ {
		Between(start, now)
	}
}

func Benchmark_Parse_thousand_years(b *testing.B) {
	format := "2006-01-02 15:04:05"
	period := "1021-01-01 to 2021-01-02"
	for i := 0; i < b.N; i++ {
		Parse(format, period)
	}
}
