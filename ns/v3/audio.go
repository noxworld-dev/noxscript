package ns

import (
	ns4 "github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/noxscript/ns/v4/audio"
)

// AudioEvent plays an audio event at a location.
func AudioEvent(name AudioName, wp WaypointID) {
	ns4.AudioEvent(audio.Name(name), ns4.AsWaypoint(wp))
}

// Music plays music. Volume is in range 0-100.
func Music(music int, volume int) {
	ns4.Music(music, volume)
}

func MusicPushEvent() {
	ns4.MusicPushEvent()
}

func MusicPopEvent() {
	ns4.MusicPopEvent()
}

func MusicEvent() {
	ns4.MusicEvent()
}
