Package [uptime](https://pkg.go.dev/github.com/gregoryv/uptime) 
provides an easy textual representation of long durations.

## Quick start

    $ go get -u github.com/gregoryv/uptime


Example

    a := time.Date(1821, 1, 01, 12, 00, 00, 0, time.UTC)
    b := time.Date(2022, 3, 07, 16, 00, 00, 0, time.UTC)
    fmt.Println(uptime.Approximate(b.Sub(a)))
    fmt.Println(uptime.Between(a, b))
    // output:
    // 201 years 3 months 24 days 4 hours
    // 201 years 2 months 6 days 4 hours
