package timeout

import (
	"time"
)

const Deadline = time.Duration(6000) * time.Millisecond

func GetDeadline(deadline time.Duration) time.Time {
	return time.Now().Add(deadline)
}

func GetDefaultDeadline() time.Time {
	return GetDeadline(Deadline)
}