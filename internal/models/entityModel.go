package models

import "time"

// key = column name and value is value
type Entity struct {
	EntityName        string // must match the table name
	StringParameters  map[string]string
	IntegerParameters map[string]int
	TimeParameters    map[string]time.Time
}
