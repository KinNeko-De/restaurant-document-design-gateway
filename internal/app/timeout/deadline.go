package timeout

import (
	"time"
)

const LongDeadline = time.Duration(60) * time.Second
const Deadline = time.Duration(100) * time.Millisecond

func GetDeadline(deadline time.Duration) time.Time {
	return time.Now().Add(deadline)
}

func GetDefaultDeadline() time.Time {
	return GetDeadline(Deadline)
}