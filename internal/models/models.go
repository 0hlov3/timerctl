package models

import "time"

type Options struct {
	Tick       time.Duration
	ShowMillis bool
	BellOnDone bool
	Max        time.Duration
	ShowBar    bool
	BarWidth   int
}
