package core

import (
	"testing"
	"time"
)

var content = `# Wed 27.05.2026

## Timeline

06:00 - 07:00		#life make coffee
08:00 - 12:00		#work prepare something
14:00 - 18:00		#admin file taxes

## Tasks

* [ ] #work call someone
`

func TestParseStrToPage_several_events(t *testing.T) {
	page := ParseStrToPage(time.Now(), "", []byte(content))

	got := len(page.Events)
	want := 3

	if got != want {
		t.Errorf("Expected %d Events, got %d", want, got)
	}
}

func TestParseStrToPage_single_event(t *testing.T) {
	page := ParseStrToPage(time.Now(), "", []byte(content))
	firstEvent := page.Events[0]

	got := firstEvent.StartTime
	want := 6 * 60

	if got != want {
		t.Errorf("Expected StartTime %d, got %d", want, got)
	}

	got = firstEvent.DurationMin
	want = 60

	if got != want {
		t.Errorf("Expected DurationMin %d, got %d", want, got)
	}

	gotStr := firstEvent.Category
	wantStr := "#life"

	if gotStr != wantStr {
		t.Errorf("Expected Category %s, got %s", wantStr, gotStr)
	}

	gotStr = firstEvent.Title
	wantStr = "make coffee"

	if gotStr != wantStr {
		t.Errorf("Expected Title %s, got %s", wantStr, gotStr)
	}
}

func TestGetEvents(t *testing.T) {
	page := ParseStrToPage(time.Now(), "", []byte(content))

	slotDuration := 15
	eventMap := page.GetEventPerSlots(slotDuration)

	index := 6 * 60
	first := eventMap[index]

	if first.Title != "make coffee" {
		t.Errorf("Expected event with title 'make coffee' at index %d", index)
	}

	index = 6*60 + 30
	same := eventMap[index]

	if same.Title != "make coffee" {
		t.Errorf("Expected event with title 'make coffee' at index %d", index)
	}

	index = 13 * 60
	_, ok := eventMap[index]

	if ok {
		t.Errorf("Expected no events at time %d", index)
	}
}
