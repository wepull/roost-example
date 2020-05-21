package util

import (
	"time"
)

const (
	MIN_TIMER_DURATION = 100 * time.Millisecond
	MAX_TIMER_DURATION = 24 * time.Hour
)

type UTimer struct {
	timer *time.Timer
}

//Create a new timer for Duration
func Timer(d time.Duration) *UTimer {
	if !isDurValid(d) {
		return nil
	}
	t := time.NewTimer(d)
	return &UTimer{timer: t}
}

//Stop the timer & drain any ticks from channel.
//and restart the timer to the new Duration.
func (t *UTimer) Reset(d time.Duration) bool {
	if !isDurValid(d) {
		return false
	}

	t.Stop() //stop & drain

	return t.timer.Reset(d)
}

//Stop the timer & drain any ticks from channel.
func (t *UTimer) Stop() {
	t.timer.Stop()
	t.drainChan()
}

//Channel that delivers ticks.
func (t *UTimer) Chan() <-chan time.Time {
	return t.timer.C
}

//
//internal functions.
//

func (t *UTimer) drainChan() {
	for {
		select {
		case <-t.timer.C:
			//drain channel
		default:
			//if there is nothing in the channel
			return
		}
	}
}

func isDurValid(d time.Duration) bool {
	if d < MIN_TIMER_DURATION {
		panic("Timer duration is less than MIN_TIMER_DURATION")
	}
	if d > MAX_TIMER_DURATION {
		panic("Timer duration is more than MAX_TIMER_DURATION")
	}
	return true
}
