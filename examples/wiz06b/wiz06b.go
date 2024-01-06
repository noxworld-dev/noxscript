package wiz06b

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	obj4   ns.ObjectID
	obj5   ns.ObjectID
	obj6   ns.ObjectID
	wp7    ns.WaypointID
	wp8    ns.WaypointID
	gvar9  ns.ObjectGroupID
	gvar10 ns.ObjectGroupID
	ivar11 int
	ivar12 int
	ivar13 int
)

func init() {
	ivar11 = 60
	ivar12 = 0
	ivar13 = 90
}
func DisableElevators() {
	ns.ObjectOff(obj6)
}
func MapInitialize() {
	obj4 = ns.Object("GauntletDoorR")
	obj5 = ns.Object("GauntletDoorL")
	obj6 = ns.Object("HideoutElevator")
	wp7 = ns.Waypoint("PlayerSounds")
	wp8 = ns.Waypoint("WellWP")
	gvar9 = ns.ObjectGroup("RoamingWarriors")
	gvar10 = ns.ObjectGroup("LightGears")
	ns.LockDoor(obj4)
	ns.LockDoor(obj5)
	ns.GroupWander(gvar9)
	ns.FrameTimer(45, DisableElevators)
}
func NoEnemys() {
	r1 := ns.IsAttackedBy(ns.GetCaller(), ns.GetHost())
	if !r1 {
		goto LABEL1
	}
	ns.GoBackHome(ns.GetCaller())
LABEL1:
	return
}
func PlayerDeath() {
	ns.DeathScreen(6)
}
func ToggleHideoutElevator() {
	ns.ObjectToggle(obj6)
	ns.ObjectGroupToggle(gvar10)
}
func Secret01Found() {
	ns.MoveWaypoint(wp7, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretFound, wp7)
	ns.ObjectOff(ns.GetTrigger())
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 100)
}
func Secret02Found() {
	ns.MoveWaypoint(wp7, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretFound, wp7)
	ns.ObjectOff(ns.GetTrigger())
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 100)
}
func SayCountdown() {
	ivar12 -= 1
	if !(ivar12 <= 0) {
		goto LABEL1
	}
	ivar12 = 0
	goto LABEL2
LABEL1:
	ns.FrameTimer(30, SayCountdown)
LABEL2:
	return
}
func WarriorRecognize() {
	var v0 int
	if !(ns.IsCaller(ns.GetHost()) && ivar12 == 0) {
		goto LABEL1
	}
	v0 = ns.Random(0, 3)
	ivar12 = ivar11
	SayCountdown()
	if v0 != 0 {
		goto LABEL2
	}
	ns.ChatTimer(ns.GetTrigger(), "War01A.scr:SetpieceGuard1Talk01", ivar13)
	goto LABEL1
LABEL2:
	ns.ChatTimer(ns.GetTrigger(), "Wiz06a:Guard1Recognize", ivar13)
LABEL1:
	return
}
func WarriorEngage() {
	var v0 int
	if !(ns.IsCaller(ns.GetHost()) && ivar12 == 0) {
		goto LABEL1
	}
	v0 = ns.Random(0, 3)
	ivar12 = ivar11
	SayCountdown()
	if v0 != 0 {
		goto LABEL2
	}
	ns.ChatTimer(ns.GetTrigger(), "War01A.scr:SetpieceGuard1Talk01", ivar13)
	goto LABEL1
LABEL2:
	ns.ChatTimer(ns.GetTrigger(), "Wiz06a:Guard1Recognize", ivar13)
LABEL1:
	return
}
func WarriorListen() {
	if !(ns.IsCaller(ns.GetHost()) && ivar12 == 0) {
		goto LABEL1
	}
	if ns.IsVisibleTo(ns.GetTrigger(), ns.GetHost()) {
		goto LABEL1
	}
	ivar12 = ivar11
	SayCountdown()
	ns.ChatTimer(ns.GetTrigger(), "Wiz06a:Guard2Listen", ivar13)
LABEL1:
	return
}
func MakeAWish() {
	var v0 int
	v0 = ns.MaxHealth(ns.GetCaller())
	ns.RestoreHealth(ns.GetCaller(), v0-ns.CurrentHealth(ns.GetCaller()))
	ns.PrintToAll("GeneralPrint:WellSignRefresh")
	ns.AudioEvent(ns.RestoreHealthName, wp8)
}
func OnEvent(typ string) {
	switch typ {
	case "MapInitialize":
		MapInitialize()
	case "PlayerDeath":
		PlayerDeath()
	}
}
