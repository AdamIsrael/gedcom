package date

import (
	"time"

	"github.com/adamisrael/gedcom/types"
)

// isSameDay checks if the event occured on today's month/day
func IsSameDay(event types.Event) bool {
	if len(event.Date) > 0 {
		t, err := Parse(event.Date)
		if err == nil {
			_, month, day := DateDiff(t, time.Now())
			if month == 0 && day == 0 {
				return true
			}
		}
	}
	return false
}

func DateDiff(a, b time.Time) (year, month, day int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	year = int(y2 - y1)
	month = int(M2 - M1)
	day = int(d2 - d1)

	// Normalize negative values
	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return
}
