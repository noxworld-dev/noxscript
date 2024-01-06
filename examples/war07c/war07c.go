package war07c

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	obj4   ns.ObjectID
	obj5   ns.ObjectID
	wp6    ns.WaypointID
	wp7    ns.WaypointID
	obj8   ns.ObjectID
	obj9   ns.ObjectID
	obj10  ns.ObjectID
	obj11  ns.ObjectID
	obj12  ns.ObjectID
	obj13  ns.ObjectID
	obj14  ns.ObjectID
	obj15  ns.ObjectID
	gvar16 ns.ObjectGroupID
	gvar17 int
)

func SecondFloorPatrolMage1() {
	ns.Wander(obj8)
}
func SecondFloorPatrolMage2() {
	ns.Wander(obj9)
}
func SecondFloorPatrolMage3() {
	ns.Wander(obj10)
}
func ActivateGears() {
	if !ns.IsObjectOn(obj5) {
		goto LABEL1
	}
	ns.MoveWaypoint(wp6, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.Gear2, wp6)
	ns.ObjectGroupOff(gvar16)
	ns.ObjectOff(obj11)
	ns.ObjectOff(obj12)
	ns.ObjectOff(obj13)
	ns.ObjectOff(obj14)
	goto LABEL2
LABEL1:
	ns.MoveWaypoint(wp6, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.Gear3, wp6)
	ns.ObjectGroupOn(gvar16)
	ns.ObjectOn(obj11)
	ns.ObjectOn(obj12)
	ns.ObjectOn(obj13)
	ns.ObjectOn(obj14)
LABEL2:
	return
}
func TrainingRoomXP() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.MoveWaypoint(wp7, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretFound, wp7)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 100)
LABEL1:
	return
}
func MapInitialize() {
	obj8 = ns.Object("PatrolMage1")
	obj9 = ns.Object("PatrolMage2")
	obj10 = ns.Object("PatrolMage3")
	obj4 = ns.Object("ToLevelOneTP")
	obj11 = ns.Object("Globe1")
	obj12 = ns.Object("Globe2")
	obj13 = ns.Object("Globe3")
	obj14 = ns.Object("Globe4")
	obj5 = ns.Object("TestGear")
	obj15 = ns.Object("ResetWizTrigger")
	gvar16 = ns.ObjectGroup("Level2Gears")
	wp6 = ns.Waypoint("PlayerSounds")
	wp7 = ns.Waypoint("SecretSounds")
	ns.CreatureHunt(obj8)
	ns.CreatureHunt(obj9)
	ns.CreatureHunt(obj10)
	ns.ObjectGroupOn(gvar16)
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
