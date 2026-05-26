// Package core contains all the heavy business logic around pages
// (a daily view in the diary), timelines, tasks, etc
package core

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Page struct {
	Time   time.Time
	Events []*Event
	Errors []error
	// Date       string
	RawContent string
}

func EmptyPage(t time.Time, date string) Page {
	return Page{
		Time: t,
		// Date:       date,
		RawContent: "<file not exists>",
	}
}

type pageSection = int

const (
	sectionStart = iota
	sectionTimeline
	sectionTasks
	sectionTomorrow
)

func ParseStrToPage(t time.Time, date string, content []byte) Page {
	// extract date from content

	section := sectionStart

	events := []*Event{}
	errors := []error{}

	lines := bytes.Split(content, []byte("\n"))

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		if bytes.Equal(line, []byte("## Timeline")) {
			section = sectionTimeline
			continue
		}

		if section == sectionTimeline {
			timelineEntry, err := parseTimelineEntry(line)

			if err != nil {
				errors = append(errors, err)
			} else {
				events = append(events, timelineEntry)
			}
			continue
		}
	}

	return Page{
		Time: t,
		// Date:       date,
		Events:     events,
		RawContent: string(content),
		Errors:     errors,
	}
}

type Event struct {
	StartTime   int
	DurationMin int
	Category    string
	Title       string
}

// A timeline entry looks like this for example:
// 08:00 - 09:00 #project prepare presentation
func parseTimelineEntry(line []byte) (*Event, error) {
	re := regexp.MustCompile(`(?P<start>\w+)\s+-\s+(?P<end>\w+)\s+(?P<cat>#\w+)\s(?P<title>.*$)`)

	match := re.FindSubmatch(line)

	names := re.SubexpNames()

	result := map[string]string{}

	for i, val := range match {
		if i != 0 && names[i] != "" {
			result[names[i]] = string(val)
		}
	}

	var err error
	start := 1

	if v, ok := result["start"]; ok {
		hhmm := strings.Replace(v, ":", "", 1)
		start, err = strconv.Atoi((hhmm))

		if err != nil {
			return nil, fmt.Errorf("could not parse start into number, got: %s", (v))
		}
	}

	cat := ""
	if v, ok := result["cat"]; ok {
		cat = v
	}

	title := ""
	if v, ok := result["title"]; ok {
		title = v
	}

	return &Event{
		StartTime:   start,
		DurationMin: 2,
		Category:    cat,
		Title:       title,
	}, nil
}
