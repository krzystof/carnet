package core

type Event struct {
	StartTime   int
	DurationMin int
	Category    string
	Title       string
}

func (e Event) Equal(other *Event) bool {
	return e.StartTime == other.StartTime && e.Title == other.Title
}
