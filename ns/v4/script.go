package ns

import (
	"math/rand"
)

type Func = any

// MapEvent is an event type for OnMapEvent function.
type MapEvent int

const (
	MapInitialize = MapEvent(iota + 1)
	MapEntry
	MapExit
	MapShutdown
)

// MapEventFunc is a callback type for OnMapEvent.
type MapEventFunc func()

// OnMapEvent registers a callback function that will be called each server map event.
//
// See MapEvent for a list of events and MapEventFunc for callback type.
func OnMapEvent(typ MapEvent, fnc MapEventFunc) {
	if impl == nil {
		return
	}
	impl.OnMapEvent(typ, fnc)
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
