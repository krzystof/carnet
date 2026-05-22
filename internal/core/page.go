// Package core contains all the heavy business logic around pages
// (a daily view in the diary), timelines, tasks, etc
package core

import "time"

type Page struct {
	Time time.Time
	Date string
}

func EmptyPage(t time.Time, date string) Page {
	return Page{
		Time: t,
		Date: date,
	}
}

func ParseStrToPage(t time.Time, date string, content []byte) Page {
	// extract date from contnet

	return Page{
		Time: t,
		Date: date,
	}
}
