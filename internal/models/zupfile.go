package models

import (
	"fmt"
	"net/url"
	"strings"
)

func CreateZup(name string, rawAddress string) (Zup, error) {
	webAddress, err := url.Parse(rawAddress)
	if err != nil {
		return Zup{}, err
	}

	return Zup{
		Name:       name,
		Generation: 0,
		WebAddress: webAddress,
		Journal:    CreateJournal(),
	}, nil
}

func ParseZup(zupString string) (Zup, error) {
	lines := strings.Split(zupString, "\n")
	if len(lines) < 2 {
		return Zup{}, fmt.Errorf("invalid Zup string format")
	}

	// name parse
	nameLine := strings.TrimSpace(lines[0])
	if !strings.HasPrefix(nameLine, "Name: ") {
		return Zup{}, fmt.Errorf("invalid Zup string format: missing Name")
	}
	name := strings.TrimPrefix(nameLine, "Name: ")

	journalEntries := strings.Join(lines[2:], "\n")
	journal, err := parseJournal(journalEntries)
	if err != nil {
		return Zup{}, fmt.Errorf("failed to parse journal: %v", err)
	}

	return Zup{
		Name:       name,
		Generation: ,
		WebAddress: nil,
		Journal:    journal,
	}, nil
}

func (z *Zup) String() string {
	return fmt.Sprintf("Name: %s\nGeneration: %i\nWebAddress: %s\nEntries:\n%s}",
		z.Name, z.Generation, z.WebAddress.String(), z.Journal.String())
}
