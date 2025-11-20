package models

import (
	"net/url"
	"time"
)

/* Zup data representation */

type Zup struct {
	Name        string
	Generation  int 
	WebAddress *url.URL
	Journal     Journal
}

/* Journal Portion */

type Journal struct {
	Entries []JournalEntry
	Count   int
}

type JournalEntry struct {
	ID       int
	Time     time.Time
	Content  string
}
