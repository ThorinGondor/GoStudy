package typeUsage

import "time"

type student struct {
	name     string
	age      int32
	birthday time.Time
	isActive bool
}
