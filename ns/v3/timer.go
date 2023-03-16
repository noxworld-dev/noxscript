package ns

import (
	ns4 "github.com/noxworld-dev/noxscript/ns/v4"
)

type TimerID int

func (id TimerID) ScriptID() int {
	return int(id)
}

func (id TimerID) TimerScriptID() int {
	return int(id)
}

// SecondTimer creates a timer with delay in seconds.
//
// This will create a timer that calls the given script function after a
// delay given in seconds.
func SecondTimer(seconds int, callback Func) TimerID {
	return asTimer(ns4.NewTimer(ns4.Seconds(float64(seconds)), callback))
}

// FrameTimer creates a timer with delay in frames.
//
// This will create a timer that calls the given script function after a
// delay given in frames.
func FrameTimer(frames int, callback Func) TimerID {
	return asTimer(ns4.NewTimer(ns4.Frames(frames), callback))
}

// SecondTimerWithArg creates a timer with delay in seconds with an argument.
//
// This will create a timer that calls the given script function after a
// delay given in seconds. The given argument will be passed into the script
// function.
func SecondTimerWithArg(seconds int, arg any, callback Func) TimerID {
	return asTimer(ns4.NewTimer(ns4.Seconds(float64(seconds)), callback, arg))
}

// FrameTimerWithArg creates a timer with delay in frames with an argument.
//
// This will create a timer that calls the given script function after a
// delay given in frames. The given argument will be passed into the script
// function.
func FrameTimerWithArg(frames int, arg any, callback Func) TimerID {
	return asTimer(ns4.NewTimer(ns4.Frames(frames), callback, arg))
}

// CancelTimer cancel a timer. Returns true on success.
func CancelTimer(id TimerID) bool {
	return ns4.AsTimer(id).Cancel()
}
