package models

import (
	"fmt"
)

func CreateZup(name string, url string) Zup {
	return Zup{
		Name: name,
		Web_address: url,
		Journal: CreateJournal(),
	}
}

func ParseZup(zupString string) Zup {
	return Zup{}
}

func (z *Zup) String() string {
	

	return fmt.Sprintf("Name: %s\n Entries:\n%s}", z.Name, z.Journal.String())
}
