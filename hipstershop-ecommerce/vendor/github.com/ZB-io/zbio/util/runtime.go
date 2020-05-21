package util

import (
	"runtime"

	"github.com/ZB-io/zbio/log"
)

// CallerFunc displays the function name who called this.
func CallerFunc(depth int) string {
	uptr := make([]uintptr, 1)
	n := runtime.Callers(depth, uptr)
	if n == 0 {
		log.Infof("No caller")
		return "No Caller"
	}
	caller := runtime.FuncForPC(uptr[0] - 1)
	if caller == nil {
		log.Infof("No caller")
		return "No Caller"
	}
	return caller.Name()
}
