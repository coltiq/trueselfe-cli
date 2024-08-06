package data

import "time"

type Pillar struct {
	Name        string
	Description string
}

type Steps struct {
	Name        string
	Description string
	Pillar      Pillar
	Created     time.Time
	Frequency   string
}
