package war07f

import "github.com/noxworld-dev/noxscript/ns/v3"

var wp4 ns.WaypointID

func MapInitialize() {
	wp4 = ns.Waypoint("FromFloor6WP")
}
func MapEntry() {
	ns.Music(27, 100)
}
func PlayerDeath() {
	ns.DeathScreen(7)
}
func OnEvent(typ string) {
	switch typ {
	case "MapInitialize":
		MapInitialize()
	case "MapEntry":
		MapEntry()
	case "PlayerDeath":
		PlayerDeath()
	}
}
