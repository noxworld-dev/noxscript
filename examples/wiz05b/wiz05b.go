package wiz05b

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	obj4   ns.ObjectID
	gvar5  ns.WallGroupID
	obj6   ns.ObjectID
	obj7   ns.ObjectID
	obj8   ns.ObjectID
	obj9   ns.ObjectID
	obj10  ns.ObjectID
	obj11  ns.ObjectID
	wp12   ns.WaypointID
	gvar13 int
	gvar14 int
	gvar15 int
	gvar16 int
	gvar17 int
	gvar18 int
	gvar19 int
	gvar20 int
	flag21 bool
	flag22 bool
	flag23 bool
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
	gvar35 ns.ObjectGroupID
	gvar36 ns.ObjectGroupID
	gvar37 ns.ObjectGroupID
	gvar38 ns.WallGroupID
	wp39   ns.WaypointID
	wp40   ns.WaypointID
	wp41   ns.WaypointID
	wp42   ns.WaypointID
	wp43   ns.WaypointID
)

func init() {
	gvar13 = 0
	gvar14 = 1
	gvar15 = 2
	gvar16 = 3
	gvar17 = 4
	gvar18 = 5
	gvar19 = gvar13
	gvar20 = gvar17
	flag21 = false
	flag22 = false
	flag23 = false
}
func StopElevator() {
	ns.ObjectOff(obj4)
	ns.FrameTimer(60, LowerKeyProtectWalls)
}
func LowerKeyProtectWalls() {
	ns.NoWallSound(true)
	ns.WallGroupOpen(gvar5)
	ns.NoWallSound(false)
}
func RaiseKey() {
	ns.ObjectOn(obj4)
	ns.FrameTimer(5, StopElevator)
}
func UnlockMainGates() {
	ns.ObjectOff(obj10)
	ns.UnlockDoor(obj6)
	ns.UnlockDoor(obj7)
}
func OpenKingGate() {
	ns.ObjectOff(obj11)
	ns.UnlockDoor(obj8)
	ns.UnlockDoor(obj9)
}
func SecretFound() {
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.MoveWaypoint(wp12, ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger()))
	ns.AudioEvent(ns.SecretFound, wp12)
	ns.GiveXp(ns.GetHost(), 40)
}
func CallForGuards() {
	ns.Frozen(obj31, true)
	ns.CreatureIdle(obj31)
	ns.SetDialog(obj31, ns.NEXT, KingStart, KingEnd)
	ns.StoryPic(obj31, "OgreWarlordPic")
	ns.StartDialog(obj31, ns.GetHost())
}
func ProtectTheKing() {
	LetterBoxOff()
	ns.Frozen(ns.GetHost(), false)
	ns.EnchantOff(obj31, ns.ENCHANT_INVULNERABLE)
	ns.Frozen(obj31, false)
	ns.ObjectOn(obj32)
	ns.ObjectOn(obj33)
	ns.AggressionLevel(obj31, 0.5)
	ns.MoveObject(obj32, ns.GetWaypointX(wp40), ns.GetWaypointY(wp40))
	ns.MoveObject(obj33, ns.GetWaypointX(wp41), ns.GetWaypointY(wp41))
	ns.Wander(obj32)
	ns.Wander(obj33)
}
func KingStart() {
	var v0 int
	v0 = gvar20
	if v0 == gvar17 {
		goto LABEL1
	}
	if v0 == gvar18 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.TellStory(ns.SwordsmanHurt, "Con05:OgreTalk07")
	ns.LookAtObject(obj31, ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj31)
	goto LABEL3
LABEL2:
	ns.TellStory(ns.SwordsmanHurt, "Con05:OgreTalk08")
	goto LABEL3
LABEL3:
	return
}
func KingEnd() {
	var v0 int
	v0 = gvar20
	if v0 == gvar17 {
		goto LABEL1
	}
	if v0 == gvar18 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	gvar20 = gvar18
	ns.StartDialog(obj31, ns.GetHost())
	goto LABEL3
LABEL2:
	ns.CancelDialog(obj31)
	ns.CreatureGuard(obj31, ns.GetWaypointX(wp39), ns.GetWaypointY(wp39), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 400)
	ProtectTheKing()
	goto LABEL3
LABEL3:
	return
}
func HorvathStart() {
	var v0 int
	v0 = gvar19
	if v0 == gvar13 {
		goto LABEL1
	}
	if v0 == gvar14 {
		goto LABEL2
	}
	if v0 == gvar15 {
		goto LABEL3
	}
	if v0 == gvar16 {
		goto LABEL4
	}
	goto LABEL5
LABEL1:
	ns.DestroyChat(obj24)
	ns.Frozen(obj24, true)
	ns.LookAtObject(obj24, ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj24)
	ns.TellStory(ns.SwordsmanHurt, "Wiz05C.scr:HorvathGreeting")
	goto LABEL5
LABEL2:
	ns.Frozen(obj24, true)
	ns.LookAtObject(obj24, ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj24)
	ns.TellStory(ns.SwordsmanHurt, "Wiz05C.scr:HorvathWaiting")
	goto LABEL5
LABEL3:
	ns.LookAtObject(obj24, ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj24)
	ns.TellStory(ns.SwordsmanHurt, "Wiz05C.scr:HorvathFreed")
	goto LABEL5
LABEL4:
	ns.LookAtObject(obj24, ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj24)
	ns.TellStory(ns.SwordsmanHurt, "Wiz05C.scr:HorvathFollowing")
	goto LABEL5
LABEL5:
	return
}
func HorvathEnd() {
	var v0 int
	v0 = gvar19
	if v0 == gvar13 {
		goto LABEL1
	}
	if v0 == gvar14 {
		goto LABEL2
	}
	if v0 == gvar15 {
		goto LABEL3
	}
	if v0 == gvar16 {
		goto LABEL4
	}
	goto LABEL5
LABEL1:
	gvar19 = gvar14
	ns.Frozen(obj24, false)
	goto LABEL5
LABEL2:
	ns.Frozen(obj24, false)
	goto LABEL5
LABEL3:
	ns.CancelDialog(obj24)
	gvar19 = gvar16
	ns.AggressionLevel(obj24, 0.5)
	ns.BecomePet(obj24)
	ns.MakeFriendly(obj24)
	ns.WideScreen(false)
	ns.Frozen(ns.GetHost(), false)
	ns.ObjectGroupOn(gvar37)
	goto LABEL5
LABEL4:
	goto LABEL5
LABEL5:
	return
}
func HorvathDeath() {
	ns.PrintToAll("Wiz05C.scr:ObjectiveFailed")
	ns.FrameTimer(30, PlayerDeath)
}
func LetterBoxOn() {
	ns.WideScreen(true)
}
func LetterBoxOff() {
	ns.WideScreen(false)
}
func PlayerDeath() {
	ns.DeathScreen(5)
}
func WhoDidWhat() {
	if !flag23 {
		goto LABEL1
	}
	if !ns.IsVisibleTo(obj25, ns.GetHost()) {
		goto LABEL2
	}
	if !ns.IsVisibleTo(obj25, obj24) {
		goto LABEL3
	}
	if flag22 {
		goto LABEL4
	}
	flag22 = true
	ns.ObjectGroupOff(gvar37)
	ns.ObjectGroupOff(gvar36)
	ns.ObjectOff(obj27)
	ns.ObjectOff(obj28)
	gvar19 = gvar15
	ns.GiveXp(ns.GetHost(), 500)
	ns.SetQuestStatus(1, "GotHorvath")
	ns.WideScreen(true)
	ns.Frozen(ns.GetHost(), true)
	ns.AggressionLevel(obj24, 0.5)
	ns.CreatureFollow(obj24, ns.GetHost())
	ns.StartDialog(obj24, ns.GetHost())
	return
LABEL4:
	goto LABEL5
LABEL3:
	if flag22 {
		goto LABEL5
	}
	ns.ObjectOn(obj27)
	ns.ObjectOn(obj28)
LABEL5:
	goto LABEL6
LABEL2:
	if !ns.IsVisibleTo(obj25, obj24) {
		goto LABEL6
	}
	ns.ObjectGroupOff(gvar36)
	ns.ObjectOff(obj27)
	ns.ObjectOff(obj28)
	gvar19 = gvar14
	ns.AggressionLevel(obj24, 0.16)
	ns.CreatureGuard(obj24, ns.GetWaypointX(wp42), ns.GetWaypointY(wp42), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
LABEL6:
	ns.FrameTimer(25, WhoDidWhat)
LABEL1:
	return
}
func MapInitialize() {
	obj24 = ns.Object("Horvath")
	obj25 = ns.Object("Spy")
	obj26 = ns.Object("CellDoor")
	obj27 = ns.Object("En")
	obj28 = ns.Object("Dis")
	obj6 = ns.Object("MainGate1")
	obj7 = ns.Object("MainGate2")
	obj29 = ns.Object("TownGate1")
	obj30 = ns.Object("TownGate2")
	obj8 = ns.Object("KingDoor1")
	obj9 = ns.Object("KingDoor2")
	obj10 = ns.Object("MainGateSwitch")
	obj11 = ns.Object("KingGateSwitch")
	obj31 = ns.Object("OgreKing")
	obj34 = ns.Object("CallForGuardsTrigger")
	obj4 = ns.Object("KeyElevator")
	obj32 = ns.Object("Guard1")
	obj33 = ns.Object("Guard2")
	gvar35 = ns.ObjectGroup("KingTriggerGroup")
	gvar36 = ns.ObjectGroup("HorvathRescueTrigger")
	gvar37 = ns.ObjectGroup("AllEnemies")
	gvar38 = ns.WallGroup("LordSafetyWalls")
	gvar5 = ns.WallGroup("KeyProtectWalls")
	wp39 = ns.Waypoint("OgreKingWaypoint")
	wp40 = ns.Waypoint("Guard1WP")
	wp41 = ns.Waypoint("Guard2WP")
	wp42 = ns.Waypoint("HorvathSpot")
	wp43 = ns.Waypoint("Spot")
	wp12 = ns.Waypoint("PlayerSounds")
	ns.SetDialog(obj24, ns.NORMAL, HorvathStart, HorvathEnd)
	ns.StoryPic(obj24, "HorvathPic")
	ns.LockDoor(obj8)
	ns.LockDoor(obj9)
	ns.LockDoor(obj6)
	ns.LockDoor(obj7)
}
func EnableRescueTrigger() {
	ns.ObjectGroupOn(gvar36)
}
func DisableRescueTrigger() {
	ns.ObjectGroupOff(gvar36)
}
func RescueHorvath() {
	flag22 = true
	ns.ObjectGroupOff(gvar37)
	ns.ObjectGroupOff(gvar36)
	ns.ObjectOff(obj27)
	ns.ObjectOff(obj28)
	gvar19 = gvar15
	ns.SetQuestStatus(1, "GotHorvath")
	LetterBoxOn()
	ns.Frozen(ns.GetHost(), true)
	ns.AggressionLevel(obj24, 0.5)
	ns.CreatureFollow(obj24, ns.GetHost())
	ns.StartDialog(obj24, ns.GetHost())
}
func PollForRescue() {
	if !(ns.IsCaller(ns.GetHost()) && flag22 == false) {
		goto LABEL1
	}
	flag23 = true
	WhoDidWhat()
LABEL1:
	return
}
func PlayerExited() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	flag23 = false
LABEL1:
	return
}
func GotoSpot() {
	if flag22 {
		goto LABEL1
	}
	if !ns.IsCaller(obj24) {
		goto LABEL2
	}
	ns.CreatureGuard(obj24, ns.GetWaypointX(wp43), ns.GetWaypointY(wp43), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
LABEL2:
	if !ns.IsAttackedBy(ns.GetCaller(), ns.GetHost()) {
		goto LABEL1
	}
	ns.CreatureGuard(ns.GetCaller(), ns.GetWaypointX(wp43), ns.GetWaypointY(wp43), ns.GetWaypointX(wp41), ns.GetWaypointY(wp41), 0)
LABEL1:
	return
}
func Activate() {
	flag21 = true
	ns.ObjectGroupOff(gvar35)
	ns.ObjectOn(obj34)
	ns.Frozen(ns.GetHost(), true)
	ns.FrameTimer(1, LetterBoxOn)
	ns.NoWallSound(true)
	ns.WallGroupOpen(gvar38)
	ns.NoWallSound(false)
	ns.Move(obj31, wp39)
}
func PummelPlayer() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.HitLocation(ns.GetTrigger(), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
LABEL1:
	return
}
func MapEntry() {
	ns.Frozen(ns.GetHost(), false)
	ns.WideScreen(false)
	ns.NoWallSound(false)
	ns.Music(7, 100)
	if ns.GetQuestStatus("Wiz05B:GotHorvath") != 1 {
		goto LABEL1
	}
	obj24 = ns.Object("Wiz05B:Horvath")
	ns.SetCallback(obj24, 5, HorvathDeath)
LABEL1:
	return
}
func FadeOut() {
	ns.Frozen(ns.GetHost(), true)
	ns.CreatureIdle(ns.GetHost())
	ns.Blind()
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
