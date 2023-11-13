Package [uptime](https://pkg.go.dev/github.com/gregoryv/uptime) 
provides an easy textual representation of long durations.

## Quick start

    $ go get -u github.com/gregoryv/uptime


Example

    d := Between(
            time.Date(2021, 1, 01, 12, 00, 00, 0, time.UTC),
            time.Date(2022, 3, 02, 13, 10, 20, 0, time.UTC),
    )
    fmt.Println(d.String())
    fmt.Println(d.Short())
    // output:
    // 1 year 2 months 1 day 1 hour 10 minutes 20 seconds
    // 1y2m1d 1h10m20s
