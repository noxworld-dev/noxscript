package war07e

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	obj4   ns.ObjectID
	obj5   ns.ObjectID
	obj6   ns.ObjectID
	obj7   ns.ObjectID
	gvar8  ns.WallGroupID
	gvar9  ns.WallGroupID
	gvar10 ns.WallGroupID
	flag11 bool
	flag12 bool
	flag13 bool
	wp14   ns.WaypointID
	wp15   ns.WaypointID
)

func init() {
	flag11 = false
	flag12 = false
	flag13 = false
}
func CreatureSetup() {
	ns.CreatureHunt(obj4)
	ns.CreatureHunt(obj5)
}
func SaveGame() {
	ns.AutoSave()
	ns.ObjectOff(obj7)
}
func WallSetOneButton() {
	if flag11 {
		goto LABEL1
	}
	ns.WallGroupOpen(gvar8)
	flag11 = true
	goto LABEL2
LABEL1:
	ns.WallGroupClose(gvar8)
	flag11 = false
LABEL2:
	return
}
func WallSetTwoButton() {
	if flag12 {
		goto LABEL1
	}
	ns.WallGroupOpen(gvar9)
	flag12 = true
	goto LABEL2
LABEL1:
	ns.WallGroupClose(gvar9)
	flag12 = false
LABEL2:
	return
}
func WallSetThreeButton() {
	if flag13 {
		goto LABEL1
	}
	ns.WallGroupOpen(gvar10)
	flag13 = true
	ns.ObjectOn(obj6)
	goto LABEL2
LABEL1:
	ns.WallGroupClose(gvar10)
	flag13 = false
	ns.ObjectOff(obj6)
LABEL2:
	return
}
func HTHAttack() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.HitLocation(ns.GetTrigger(), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
LABEL1:
	return
}
func MapInitialize() {
	obj4 = ns.Object("Hunter2")
	obj5 = ns.Object("Hunter3")
	obj6 = ns.Object("MainTeleporter")
	obj7 = ns.Object("SaveGameTrigger")
	gvar8 = ns.WallGroup("WallSetOne")
	gvar9 = ns.WallGroup("WallSetTwo")
	gvar10 = ns.WallGroup("WallSetThree")
	wp14 = ns.Waypoint("PlayerSounds")
	wp15 = ns.Waypoint("FromFloor5WP")
	ns.ObjectOff(obj6)
	ns.FrameTimer(1, CreatureSetup)
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
