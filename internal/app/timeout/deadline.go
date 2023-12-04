package timeout

import (
	"time"
)

const LongDeadline = time.Duration(10 * time.Minute)
const Deadline = time.Duration(100 * time.Millisecond)

func GetDeadline(deadline time.Duration) time.Time {
	return time.Now().Add(deadline)
}
