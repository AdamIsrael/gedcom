package date

import (
	"time"

	"github.com/araddon/dateparse"
)

/*
Dates in GEDCOM can be fuzzy, i.e., all of these are valid

Between 4 Apr 1935 and 9 Apr 1935
Btw. 4 April 1935 and 9 April 1935
Between 4 April 1935 and 9 April 1935
Abt. April 1935
About Apr 1935
After Apr. 4 1935
4 April 1935
4 Apr 1935
4 Apr. 1935
April 4, 1935



*/

// Parse will attempt to parse a date string and return
func Parse(date string) (time.Time, error) {
	var t time.Time
	var err error

	// For now, try to parse it with dateparse first
	t, err = dateparse.ParseLocal(date)
	if err != nil {
		// attempt to parse it manually

	}
	return t, err
}
