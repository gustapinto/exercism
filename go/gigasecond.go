package gigasecond

import "time"

const GIGASECOND = 1000000000 * time.Second

func AddGigasecond(t time.Time) time.Time {
	return t.Add(GIGASECOND)
}
