package models

import (
	"time"
)

func createJournal() Journal {
	return Journal{
		entries: nil,
		count: 0,
	}
}

func (j *Journal) newEntry(location string, content string) {
	now := time.Now()
	j.count += 1
	j.entries = append(j.entries, JournalEntry{
		id:       j.count,
		time:     now.Format("15:04:05"),
		date:     now.Format("2025-01-02"),
		location: location,
		content:  content,
	})
}


