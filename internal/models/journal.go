package models

import (
	"fmt"
	"time"
)

func CreateJournal() Journal {
	return Journal{
		Entries: nil,
		Count:   0,
	}
}

func (j *Journal) NewEntry(location string, content string) {
	j.Entries = append(j.Entries, JournalEntry{
		ID:      j.Count,
		Time:    time.Now(),
		Content: content,
	})

	j.Count += 1
}

func (j *Journal) String() string {
	var journalEntries string
	for _, entry := range j.Entries {
		journalEntries += entry.String()
	}

	return journalEntries
}

func (j *Journal) GetEntry(id int) string {
	return j.Entries[id].String()
}

func (je *JournalEntry) String() string {
	return fmt.Sprintf(
		`{
			ID: %d
			Time: %s
			Entry: 
			%s
		}`,
		je.ID, je.Time, je.Content)
}
