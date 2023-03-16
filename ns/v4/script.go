package ns

import (
	"math/rand"
)

type Func = any

type Timer interface {
	TimerHandle
	// Cancel a timer. Returns true if successful.
	Cancel() bool
}

// NewTimer creates a timer that calls the given script function after a given delay.
//
// Example:
//
//	// Trigger by function reference:
//	NewTimer(Frames(10), myCallback)
//	// Trigger by function name (can call original NoxScript as well):
//	NewTimer(Seconds(10), "myCallback")
//	// Passing arguments to the callback:
//	NewTimer(Time(3*time.Second), "myCallback", obj)
func NewTimer(dt Duration, fnc Func, args ...any) Timer {
	if impl == nil {
		return nil
	}
	return impl.NewTimer(dt, fnc, args...)
}

// RandomFloat generates random float.
func RandomFloat(min float32, max float32) float32 {
	if impl == nil {
		return min + rand.Float32()*(max-min)
	}
	return impl.RandomFloat(min, max)
}

// Random generates random int.
func Random(min int, max int) int {
	if impl == nil {
		return min + rand.Intn(max-min)
	}
	return impl.Random(min, max)
}

// StopScript exits current script function.
func StopScript(value any) {
	if impl == nil {
		return
	}
	impl.StopScript(value)
}

// AutoSave triggers an autosave. Only solo games.
func AutoSave() {
	if impl == nil {
		return
	}
	impl.AutoSave()
}

// StartupScreen shows startup screen to host.
func StartupScreen(which int) {
	if impl == nil {
		return
	}
	impl.StartupScreen(which)
}

// DeathScreen shows death screen to host.
func DeathScreen(which int) {
	if impl == nil {
		return
	}
	impl.DeathScreen(which)
}
