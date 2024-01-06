package con07f

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
	obj19  ns.ObjectID
	obj20  ns.ObjectID
	obj21  ns.ObjectID
	obj22  ns.ObjectID
	obj23  ns.ObjectID
	obj24  ns.ObjectID
	obj25  ns.ObjectID
	obj26  ns.ObjectID
	obj27  ns.ObjectID
	obj28  ns.ObjectID
	obj29  ns.ObjectID
	obj30  ns.ObjectID
	obj31  ns.ObjectID
	obj32  ns.ObjectID
	obj33  ns.ObjectID
	obj34  ns.ObjectID
	obj35  ns.ObjectID
	obj36  ns.ObjectID
	gvar37 ns.WallGroupID
	gvar38 ns.WallGroupID
	gvar39 ns.WallGroupID
	gvar40 int
	flag41 bool
	flag42 bool
	flag43 bool
	flag44 bool
	flag45 bool
	wp46   ns.WaypointID
	wp47   ns.WaypointID
	wp48   ns.WaypointID
	wp49   ns.WaypointID
	wp50   ns.WaypointID
	wp51   ns.WaypointID
	wp52   ns.WaypointID
	wp53   ns.WaypointID
	wp54   ns.WaypointID
)

func init() {
	gvar40 = 0
	flag41 = false
	flag42 = false
	flag43 = false
	flag44 = false
	flag45 = false
}
func CreatureSetup() {
	ns.SetOwner(obj35, obj4)
	ns.SetOwner(obj35, obj5)
	ns.SetOwner(obj35, obj6)
	ns.SetOwner(obj35, obj7)
	ns.SetOwner(obj35, obj8)
	ns.SetOwner(obj35, obj9)
	ns.SetOwner(obj35, obj10)
	ns.SetOwner(obj35, obj11)
	ns.SetOwner(obj35, obj12)
	ns.SetOwner(obj35, obj13)
	ns.SetOwner(obj35, obj14)
	ns.SetOwner(obj35, obj15)
	ns.SetOwner(obj35, obj16)
	ns.SetOwner(obj35, obj17)
	ns.SetOwner(obj35, obj18)
	ns.SetOwner(obj35, obj19)
	ns.SetOwner(obj35, obj20)
	ns.SetOwner(obj35, obj21)
	ns.SetOwner(obj35, obj22)
	ns.SetOwner(obj35, obj23)
	ns.SetOwner(obj35, obj24)
	ns.SetOwner(obj35, obj25)
	ns.SetOwner(obj35, obj26)
	ns.SetOwner(ns.GetHost(), obj30)
	ns.SetOwner(ns.GetHost(), obj31)
	ns.SetOwner(ns.GetHost(), obj32)
	if !(ns.CurrentHealth(obj30) > 0) {
		goto LABEL1
	}
	ns.CreatureHunt(obj30)
LABEL1:
	if !(ns.CurrentHealth(obj31) > 0) {
		goto LABEL2
	}
	ns.CreatureHunt(obj31)
LABEL2:
	if !(ns.CurrentHealth(obj32) > 0) {
		goto LABEL3
	}
	ns.CreatureHunt(obj32)
LABEL3:
	return
}
func CreateDemons() {
}
func DemonSpawnCheck() {
	flag42 = true
}
func WallSetOneButton() {
	if flag43 {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.WallGroupOpen(gvar37)
	flag43 = true
LABEL1:
	return
}
func WallSetTwoButton() {
	if flag44 {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.WallGroupOpen(gvar38)
	flag44 = true
LABEL1:
	return
}
func WallSetThreeButton() {
	if flag45 {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.WallGroupOpen(gvar39)
	flag45 = true
	ns.ObjectOn(obj36)
LABEL1:
	return
}
func BlockTrap() {
	if ns.IsObjectOn(obj33) {
		goto LABEL1
	}
	ns.ObjectOn(obj33)
	ns.Move(obj33, wp53)
LABEL1:
	return
}
func BlockAtWP1() {
	if !ns.IsCaller(obj34) {
		goto LABEL1
	}
	ns.MoveWaypoint(wp52, ns.GetObjectX(obj34), ns.GetObjectY(obj34))
	ns.AudioEvent(ns.HammerMissing, wp52)
	ns.Effect(ns.JIGGLE, ns.GetObjectX(obj34), ns.GetObjectY(obj34), 10, 0)
	ns.Move(obj33, wp54)
LABEL1:
	return
}
func BlockAtWP2() {
}
func MapInitialize() {
	obj4 = ns.Object("EmberDemon1")
	obj5 = ns.Object("EmberDemon2")
	obj6 = ns.Object("EmberDemon3")
	obj7 = ns.Object("EmberDemon4")
	obj8 = ns.Object("EmberDemon5")
	obj9 = ns.Object("EmberDemon6")
	obj10 = ns.Object("EmberDemon7")
	obj11 = ns.Object("EmberDemon8")
	obj12 = ns.Object("EmberDemon9")
	obj13 = ns.Object("EmberDemon10")
	obj14 = ns.Object("EmberDemon11")
	obj15 = ns.Object("EmberDemon12")
	obj16 = ns.Object("EmberDemon13")
	obj17 = ns.Object("EmberDemon14")
	obj18 = ns.Object("EmberDemon15")
	obj19 = ns.Object("EmberDemon16")
	obj20 = ns.Object("EmberDemon17")
	obj21 = ns.Object("EmberDemon18")
	obj22 = ns.Object("EmberDemon19")
	obj23 = ns.Object("EmberDemon20")
	obj24 = ns.Object("EmberDemon21")
	obj25 = ns.Object("EmberDemon22")
	obj26 = ns.Object("EmberDemon23")
	obj27 = ns.Object("EmberDemon24")
	obj28 = ns.Object("EmberDemon25")
	obj29 = ns.Object("EmberDemon26")
	obj30 = ns.Object("PCMage4A")
	obj31 = ns.Object("PCMage4B")
	obj32 = ns.Object("PCMage4C")
	obj33 = ns.Object("Block1Mover")
	obj34 = ns.Object("Block1")
	obj35 = ns.Object("StoolOfWonder")
	obj36 = ns.Object("MainTeleporter")
	gvar37 = ns.WallGroup("WallSetOne")
	gvar38 = ns.WallGroup("WallSetTwo")
	gvar39 = ns.WallGroup("WallSetThree")
	wp46 = ns.Waypoint("DemonSpawn1WP")
	wp47 = ns.Waypoint("DemonSpawn2WP")
	wp48 = ns.Waypoint("DemonSpawn3WP")
	wp49 = ns.Waypoint("DemonSpawn4WP")
	wp50 = ns.Waypoint("DemonSpawn5WP")
	wp51 = ns.Waypoint("DemonSpawn6WP")
	wp52 = ns.Waypoint("PlayerSounds")
	wp53 = ns.Waypoint("BlockWP1")
	wp54 = ns.Waypoint("BlockWP2")
	ns.ObjectOff(obj36)
	ns.FrameTimer(1, CreatureSetup)
	CreateDemons()
}
func MapEntry() {
	ns.Music(28, 100)
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
