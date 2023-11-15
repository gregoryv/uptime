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

func Benchmark_Between_hundred_years(b *testing.B) {
	start := time.Date(1921, 11, 6, 11, 34, 13, 0, time.UTC)
	now := time.Date(2023, 3, 19, 11, 34, 13, 0, time.UTC)
	for i := 0; i < b.N; i++ {
		Between(start, now)
	}
}
