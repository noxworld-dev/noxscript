package con07e

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	obj4   ns.ObjectID
	obj5   ns.ObjectID
	obj6   ns.ObjectID
	obj7   ns.ObjectID
	obj8   ns.ObjectID
	obj9   ns.ObjectID
	obj10  ns.ObjectID
	obj11  ns.ObjectID
	obj12  ns.ObjectID
	obj13  ns.ObjectID
	obj14  ns.ObjectID
	obj15  ns.ObjectID
	obj16  ns.ObjectID
	obj17  ns.ObjectID
	obj18  ns.ObjectID
	gvar19 ns.ObjectGroupID
	gvar20 ns.ObjectGroupID
	obj21  ns.ObjectID
	gvar22 ns.ObjectGroupID
	gvar23 ns.ObjectGroupID
	gvar24 ns.ObjectGroupID
	gvar25 ns.WallGroupID
	wp26   ns.WaypointID
	wp27   ns.WaypointID
)

func CreatureSetup() {
	if !(ns.CurrentHealth(obj4) > 0) {
		goto LABEL1
	}
	ns.CreatureHunt(obj4)
LABEL1:
	if !(ns.CurrentHealth(obj5) > 0) {
		goto LABEL2
	}
	ns.CreatureHunt(obj5)
LABEL2:
	if !(ns.CurrentHealth(obj6) > 0) {
		goto LABEL3
	}
	ns.CreatureHunt(obj6)
LABEL3:
	if !(ns.CurrentHealth(obj7) > 0) {
		goto LABEL4
	}
	ns.CreatureHunt(obj7)
LABEL4:
	if !(ns.CurrentHealth(obj8) > 0) {
		goto LABEL5
	}
	ns.CreatureHunt(obj8)
LABEL5:
	if !(ns.CurrentHealth(obj9) > 0) {
		goto LABEL6
	}
	ns.CreatureHunt(obj9)
LABEL6:
	if !(ns.CurrentHealth(obj10) > 0) {
		goto LABEL7
	}
	ns.CreatureHunt(obj10)
LABEL7:
	if !(ns.CurrentHealth(obj11) > 0) {
		goto LABEL8
	}
	ns.CreatureHunt(obj11)
LABEL8:
	ns.SetOwner(ns.GetHost(), obj12)
	ns.SetOwner(ns.GetHost(), obj13)
	ns.SetOwner(ns.GetHost(), obj14)
	ns.SetOwner(ns.GetHost(), obj15)
	ns.SetOwner(ns.GetHost(), obj16)
	ns.SetOwner(ns.GetHost(), obj17)
	ns.SetOwner(ns.GetHost(), obj18)
}
func StartImpDemon() {
	ns.ObjectGroupOn(gvar23)
	ns.ObjectGroupOn(gvar24)
	ns.WayPointOn(wp26)
	ns.WayPointOn(wp27)
	ns.ObjectGroupOff(gvar19)
	ns.ObjectGroupOff(gvar20)
	ns.ObjectOff(obj21)
}
func FollowPlayer() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	if !(ns.CurrentHealth(ns.GetTrigger()) > 0) {
		goto LABEL1
	}
	ns.CreatureFollow(ns.GetTrigger(), ns.GetHost())
LABEL1:
	return
}
func EnableMuseum() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.ObjectGroupOn(gvar22)
	ns.WallGroupOpen(gvar25)
LABEL1:
	return
}
func MapInitialize() {
	obj4 = ns.Object("MuseumDemon1")
	obj5 = ns.Object("MuseumDemon2")
	obj6 = ns.Object("MuseumDemon3")
	obj7 = ns.Object("MuseumDemon4")
	obj8 = ns.Object("MuseumDemon5")
	obj9 = ns.Object("MuseumDemon6")
	obj10 = ns.Object("MuseumDemon7")
	obj11 = ns.Object("MuseumDemon8")
	obj12 = ns.Object("Kheldon")
	obj13 = ns.Object("Dalmen")
	obj14 = ns.Object("Imp1")
	obj15 = ns.Object("Imp2")
	obj16 = ns.Object("Imp3")
	obj17 = ns.Object("Imp4")
	obj18 = ns.Object("Imp5")
	gvar22 = ns.ObjectGroup("MuseumDemons")
	gvar19 = ns.ObjectGroup("ImpDemonTriggers")
	gvar20 = ns.ObjectGroup("ImpDemonTriggers2")
	obj21 = ns.Object("ImpDemonTrigger3")
	gvar23 = ns.ObjectGroupID(ns.Object("FloorThreeSetPiece")) // FIXME
	gvar24 = ns.ObjectGroupID(ns.Object("KheldonGroup"))       // FIXME
	wp26 = ns.Waypoint("ImpWP1")
	wp27 = ns.Waypoint("ImpWP2")
	ns.WayPointOff(wp26)
	ns.WayPointOff(wp27)
	ns.ObjectGroupOff(gvar23)
	ns.ObjectGroupOff(gvar24)
	ns.ObjectGroupOff(gvar22)
	ns.FrameTimer(1, CreatureSetup)
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
