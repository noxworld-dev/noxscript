package ns

import (
	"time"

	"github.com/noxworld-dev/opennox-lib/script"

	ns4 "github.com/noxworld-dev/noxscript/ns/v4"
)

func timeSeconds(sec float32) script.Duration {
	dur := time.Duration(float64(sec) * float64(time.Second))
	return script.Time(dur)
}

type TimerID int

func (id TimerID) TimerScriptID() int {
	return int(id)
}

// SecondTimer creates a timer with delay in seconds.
//
// This will create a timer that calls the given script function after a
// delay given in seconds.
func SecondTimer(seconds int, callback Func) TimerID {
	return asTimer(ns4.NewTimer(timeSeconds(float32(seconds)), callback))
}

// FrameTimer creates a timer with delay in frames.
//
// This will create a timer that calls the given script function after a
// delay given in frames.
func FrameTimer(frames int, callback Func) TimerID {
	return asTimer(ns4.NewTimer(script.Frames(frames), callback))
}

// SecondTimerWithArg creates a timer with delay in seconds with an argument.
//
// This will create a timer that calls the given script function after a
// delay given in seconds. The given argument will be passed into the script
// function.
func SecondTimerWithArg(seconds int, arg any, callback Func) TimerID {
	return asTimer(ns4.NewTimer(timeSeconds(float32(seconds)), callback, arg))
}

// FrameTimerWithArg creates a timer with delay in frames with an argument.
//
// This will create a timer that calls the given script function after a
// delay given in frames. The given argument will be passed into the script
// function.
func FrameTimerWithArg(frames int, arg any, callback Func) TimerID {
	return asTimer(ns4.NewTimer(script.Frames(frames), callback, arg))
}

// CancelTimer cancel a timer. Returns true on success.
func CancelTimer(id TimerID) bool {
	return ns4.AsTimer(id).Cancel()
}
