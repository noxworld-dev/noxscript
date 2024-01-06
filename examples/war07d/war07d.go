package war07d

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	wp4    ns.WaypointID
	obj5   ns.ObjectID
	obj6   ns.ObjectID
	obj7   ns.ObjectID
	obj8   ns.ObjectID
	obj9   ns.ObjectID
	obj10  ns.ObjectID
	obj11  ns.ObjectID
	gvar12 ns.WallGroupID
)

func SetUpCreatures() {
	ns.CreatureHunt(obj6)
	ns.CreatureHunt(obj7)
	ns.CreatureHunt(obj8)
	ns.CreatureHunt(obj9)
	ns.CreatureHunt(obj10)
	ns.CreatureHunt(obj11)
}
func OpenMuseumWalls() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.WallGroupOpen(gvar12)
LABEL1:
	return
}
func MapInitialize() {
	wp4 = ns.Waypoint("PlayerSounds")
	obj6 = ns.Object("WizHunter1")
	obj7 = ns.Object("WizHunter2")
	obj8 = ns.Object("WizHunter3")
	obj9 = ns.Object("MuseumMage1")
	obj10 = ns.Object("MuseumMage2")
	obj11 = ns.Object("MuseumMage3")
	obj5 = ns.Object("Dalmen")
	gvar12 = ns.WallGroup("MuseumWalls")
	ns.FrameTimer(1, SetUpCreatures)
}
func MapEntry() {
	ns.Music(16, 100)
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
