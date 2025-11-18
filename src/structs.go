package src

import ()

/* Zup data representation */

type Zup struct {
	name    string
	journal Journal
	planner Planner
}

/* Journal Portion */

type Journal struct {
	entries []JournalEntry
	count   uint16
}

type JournalEntry struct {
	id       int
	time     string
	date     string
	location string
	content  string
}

/* Planner Portion */

type Planner struct {

}

type TaskList struct {
	id        int
	name      string
	tasks     []Task
	time      string
	count     uint16
	countDone uint16
}

type Task struct {
	id   int
	name string
	done bool
}
