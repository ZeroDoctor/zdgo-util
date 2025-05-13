package zdutil

import (
	"fmt"
	"strings"
	"time"
)

const (
	day  = time.Minute * 60 * 24
	year = 365 * day
)

// FormatDuration formats a time.Duration into a human-readable string.
//
// The formatting is similar to what the standard library's
// time.Duration.String() function does, but with a few key differences:
//
//  1. Durations less than 1 day are formatted as the standard library does.
//  2. Durations greater than or equal to 1 day are formatted as a count of
//     years and days, with the years always included and the days only
//     included if the duration is not an exact count of years.
//
// The function rounds the duration to the nearest second, millisecond, or
// microsecond before formatting, depending on the duration.
func FormatDuration(d time.Duration) string {
	switch {
	case d > time.Second:
		d = d.Round(time.Second)
	case d > time.Millisecond:
		d = d.Round(time.Millisecond)
	case d > time.Microsecond:
		d = d.Round(time.Microsecond)
	}

	if d < day {
		return d.String()
	}

	var b strings.Builder

	if d >= year {
		years := d / year
		fmt.Fprintf(&b, "%dy", years)
		d -= years * year
	}

	days := d / day
	d -= days * day
	fmt.Fprintf(&b, "%dd%s", days, d)

	return b.String()
}
