package con05b

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
	wp13   ns.WaypointID
	gvar14 int
	gvar15 int
	gvar16 int
	flag17 bool
	obj18  ns.ObjectID
	obj19  ns.ObjectID
	obj20  ns.ObjectID
	obj21  ns.ObjectID
	obj22  ns.ObjectID
	obj23  ns.ObjectID
	gvar24 ns.ObjectGroupID
	gvar25 ns.WallGroupID
	wp26   ns.WaypointID
	wp27   ns.WaypointID
	wp28   ns.WaypointID
)

func init() {
	gvar14 = 0
	gvar15 = 1
	gvar16 = gvar14
	flag17 = false
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
	ns.MoveWaypoint(wp13, ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger()))
	ns.AudioEvent(ns.SecretFound, wp13)
	ns.GiveXp(ns.GetHost(), 40)
}
func PickupAmulet() {
	ns.GiveXp(ns.GetHost(), 50)
	ns.AudioEvent(ns.JournalEntryAdd, wp12)
}
func CallForGuards() {
	ns.Frozen(obj21, true)
	ns.CreatureIdle(obj21)
	ns.SetDialog(obj21, ns.NEXT, KingStart, KingEnd)
	ns.StoryPic(obj21, "OgreWarlordPic")
	ns.StartDialog(obj21, ns.GetHost())
}
func ProtectTheKing() {
	LetterBoxOff()
	ns.Frozen(ns.GetHost(), false)
	ns.EnchantOff(obj21, ns.ENCHANT_INVULNERABLE)
	ns.ClearOwner(obj21)
	ns.Frozen(obj21, false)
	ns.ObjectOn(obj22)
	ns.ObjectOn(obj23)
	ns.AggressionLevel(obj21, 0.83)
	ns.MoveObject(obj22, ns.GetWaypointX(wp27), ns.GetWaypointY(wp27))
	ns.MoveObject(obj23, ns.GetWaypointX(wp28), ns.GetWaypointY(wp28))
	ns.Wander(obj22)
	ns.Wander(obj23)
}
func KingStart() {
	var v0 int
	v0 = gvar16
	if v0 == gvar14 {
		goto LABEL1
	}
	if v0 == gvar15 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.TellStory(ns.SwordsmanHurt, "Con05:OgreTalk07")
	ns.LookAtObject(obj21, ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj21)
	goto LABEL3
LABEL2:
	ns.TellStory(ns.SwordsmanHurt, "Con05:OgreTalk08")
	goto LABEL3
LABEL3:
	return
}
func KingEnd() {
	var v0 int
	v0 = gvar16
	if v0 == gvar14 {
		goto LABEL1
	}
	if v0 == gvar15 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	gvar16 = gvar15
	ns.StartDialog(obj21, ns.GetHost())
	goto LABEL3
LABEL2:
	ns.CancelDialog(obj21)
	ns.CreatureGuard(obj21, ns.GetWaypointX(wp26), ns.GetWaypointY(wp26), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 500)
	ProtectTheKing()
	goto LABEL3
LABEL3:
	return
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
func MapInitialize() {
	obj6 = ns.Object("MainGate1")
	obj7 = ns.Object("MainGate2")
	obj18 = ns.Object("TownGate1")
	obj19 = ns.Object("TownGate2")
	obj8 = ns.Object("KingDoor1")
	obj9 = ns.Object("KingDoor2")
	obj10 = ns.Object("MainGateSwitch")
	obj11 = ns.Object("KingDoorSwitch")
	obj21 = ns.Object("OgreKing")
	obj20 = ns.Object("CallForGuardsTrigger")
	obj4 = ns.Object("KeyElevator")
	obj22 = ns.Object("Guard1")
	obj23 = ns.Object("Guard2")
	gvar24 = ns.ObjectGroup("KingTriggerGroup")
	gvar25 = ns.WallGroup("LordSafetyWalls")
	gvar5 = ns.WallGroup("KeyProtectWalls")
	wp26 = ns.Waypoint("OgreKingWaypoint")
	wp27 = ns.Waypoint("Guard1WP")
	wp28 = ns.Waypoint("Guard2WP")
	wp12 = ns.Waypoint("AmuletSoundWP")
	wp13 = ns.Waypoint("PlayerSounds")
	ns.LockDoor(obj8)
	ns.LockDoor(obj9)
	ns.LockDoor(obj6)
	ns.LockDoor(obj7)
}
func Activate() {
	var v0 bool
	_ = v0
	v0 = true
	ns.ObjectGroupOff(gvar24)
	ns.NoWallSound(true)
	ns.WallGroupOpen(gvar25)
	ns.NoWallSound(false)
	ns.ObjectOn(obj20)
	ns.Frozen(ns.GetHost(), true)
	ns.FrameTimer(1, LetterBoxOn)
	ns.Move(obj21, wp26)
}
func PummelPlayer() {
	ns.HitLocation(ns.GetTrigger(), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
}
func MapEntry() {
	ns.Frozen(ns.GetHost(), false)
	ns.WideScreen(false)
	ns.Music(7, 100)
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
