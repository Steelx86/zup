package models

import (
	"fmt"
	"strings"
	"time"
)

func CreateJournal() Journal {
	return Journal{
		Entries: nil,
		Count:   0,
	}
}

func parseJournal(journalString string) (Journal, error) {
	lines := strings.Split(journalString, "\n")
	journal := CreateJournal()

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		entry, err := parseJournalEntry(line)
		if err != nil {
			return Journal{}, fmt.Errorf("failed to parse journal entry: %v", err)
		}
		journal.Entries = append(journal.Entries, entry)
		journal.Count++
	}

	return journal, nil
}

func parseJournalEntry(entryString string) (JournalEntry, error) {
	var id int
	var timeString, content string

	// entry parse
	_, err := fmt.Sscanf(entryString, `{ ID: %d Time: %s Entry %s }`, &id, &timeString, &content)
	if err != nil {
		return JournalEntry{}, fmt.Errorf("invalid journal entry format: %v", err)
	}

	// Time parse
	parsedTime, err := time.Parse(time.RFC3339, timeString)
	if err != nil {
		return JournalEntry{}, fmt.Errorf("invalid time format: %v", err)
	}

	return JournalEntry{
		ID:      id,
		Time:    parsedTime,
		Content: content,
	}, nil
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
		`---
		ID: %d
		Time: %s
		Entry: 
		%s
		---`,
		je.ID, je.Time, je.Content)
}
