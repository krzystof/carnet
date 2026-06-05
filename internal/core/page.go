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

// type pageSection = int

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

		re := regexp.MustCompile(`^## (.*$)`)
		match := re.FindSubmatch(line)

		word := []byte{}
		if len(match) == 2 {
			word = match[1]
		}

		if bytes.Equal(word, []byte("Timeline")) {
			section = sectionTimeline
			continue
		} else if bytes.Equal(word, []byte("Tasks")) {
			section = sectionTasks
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

// A timeline entry looks like this for example:
// 08:00 - 09:00 #project prepare presentation
func parseTimelineEntry(line []byte) (*Event, error) {
	re := regexp.MustCompile(`(?P<start>\w+:\w+) - (?P<end>\w+:\w+)\W+(?P<cat>#\w+)?\W(?P<title>.*$)`)

	match := re.FindSubmatch(line)

	names := re.SubexpNames()

	result := map[string]string{}

	for i, val := range match {
		if i != 0 && names[i] != "" {
			result[names[i]] = string(val)
		}
	}

	startStr, ok := result["start"]
	if !ok {
		return nil, fmt.Errorf("missing start field")
	}

	start, err := parseTimeToMinutesSinceMidnight(startStr)
	if err != nil {
		return nil, err
	}

	endStr, ok := result["end"]
	if !ok {
		return nil, fmt.Errorf("missing end field")
	}

	end, err := parseTimeToMinutesSinceMidnight(endStr)
	if err != nil {
		return nil, err
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
		DurationMin: end - start,
		Category:    cat,
		Title:       title,
	}, nil
}

func parseTimeToMinutesSinceMidnight(hhmm string) (int, error) {
	hh, mm, ok := strings.Cut(hhmm, ":")

	if !ok {
		return 0, fmt.Errorf("could not parse time into number, got: %s", hhmm)
	}

	h, err := strconv.Atoi((hh))

	if err != nil {
		return 0, fmt.Errorf("could not parse start hour into number, got: %s", hhmm)
	}

	m, err := strconv.Atoi((mm))

	if err != nil {
		return 0, fmt.Errorf("could not parse start minutes into number, got: %s", hhmm)
	}

	return h*60 + m, nil
}

// GetEventPerSlots returns a map of the page's Events, keyed by the timeslot
// An event might appear several time, if it covers multiple timeslots.
func (p Page) GetEventPerSlots(slotDuration int) map[int]*Event {
	var m = make(map[int]*Event)

	for _, e := range p.Events {
		endTime := e.StartTime + e.DurationMin

		for c := e.StartTime; c < endTime; c += slotDuration {
			m[c] = e
		}
	}

	return m
}

func (p Page) IsToday() bool {
	now := time.Now()
	return p.Time.Year() == now.Year() && p.Time.Month() == now.Month() && p.Time.Day() == now.Day()
}
