package wiz03c

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	obj4   ns.ObjectID
	obj5   ns.ObjectID
	gvar6  ns.ObjectGroupID
	gvar7  ns.ObjectGroupID
	obj8   ns.ObjectID
	obj9   ns.ObjectID
	obj10  ns.ObjectID
	obj11  ns.ObjectID
	gvar12 ns.ObjectGroupID
	obj13  ns.ObjectID
	wp14   ns.WaypointID
)

func PlayerDeath() {
	ns.DeathScreen(3)
}
func PlaySubMusic() {
	ns.ObjectOff(ns.GetTrigger())
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.Music(19, 100)
LABEL1:
	return
}
func MapInitialize() {
	obj4 = ns.Object("Wiz03b:AmuletofTeleportation")
	obj5 = ns.Object("FloorPit")
	gvar6 = ns.ObjectGroup("CreatureGroup1")
	ns.ObjectGroupOff(gvar6)
	gvar7 = ns.ObjectGroup("CreatureGroup2")
	ns.ObjectGroupOff(gvar7)
	obj8 = ns.Object("WizardDesk1")
	obj9 = ns.Object("WizardDesk2")
	obj10 = ns.Object("WizardDesk3")
	obj11 = ns.Object("WizardDesk4")
	ns.Damage(obj8, 0, 100, 12)
	ns.Damage(obj9, 0, 100, 12)
	ns.Damage(obj10, 0, 100, 12)
	ns.Damage(obj11, 0, 100, 12)
	gvar12 = ns.ObjectGroupID(ns.Object("SpiderAmbushGroup")) // FIXME
	ns.ObjectGroupOff(gvar12)
	wp14 = ns.Waypoint("Wiz03bExitWP")
	obj13 = ns.Object("ExitTrig")
}
func MapEntry() {
	PlaySubMusic()
	ns.UnBlind()
	ns.ObjectOn(obj13)
}
func SpiderAmbush() {
	ns.ObjectGroupOn(gvar12)
}
func EnableCreatureGroup1() {
	ns.ObjectGroupOn(gvar6)
}
func EnableCreatureGroup2() {
	ns.ObjectGroupOn(gvar7)
}
func QuestSwitch() {
	if !ns.HasItem(ns.GetHost(), ns.Object("Wiz03b:AmuletofTeleportation")) {
		goto LABEL1
	}
	ns.ObjectOn(obj5)
LABEL1:
	return
}
func PlayerWiz03bExit() {
	var (
		v0 float32
		v1 float32
	)
	v0 = ns.GetWaypointX(wp14)
	v1 = ns.GetWaypointY(wp14)
	ns.MoveObject(ns.GetHost(), v0, v1)
}
func FreezePlayer() {
	ns.Frozen(ns.GetHost(), true)
}
func GotoWiz03b() {
	ns.ObjectOff(obj13)
	ns.Blind()
	ns.FrameTimer(30, FreezePlayer)
	ns.FrameTimer(60, PlayerWiz03bExit)
}
func MonsterGoHome() {
	r1 := ns.IsAttackedBy(ns.GetCaller(), ns.GetHost())
	if !r1 {
		goto LABEL1
	}
	ns.GoBackHome(ns.GetCaller())
LABEL1:
	return
}
func PlayWanderMusic() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.Music(22, 100)
LABEL1:
	return
}
func OnEvent(typ string) {
	switch typ {
	case "PlayerDeath":
		PlayerDeath()
	case "MapInitialize":
		MapInitialize()
	case "MapEntry":
		MapEntry()
	}
}
