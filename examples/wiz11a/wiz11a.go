package wiz11a

import (
	"strconv"

	"github.com/noxworld-dev/noxscript/ns/v3"
)

var (
	obj4   ns.ObjectID
	obj5   ns.ObjectID
	obj6   ns.ObjectID
	obj7   ns.ObjectID
	obj8   ns.ObjectID
	obj9   ns.ObjectID
	gvar10 ns.WaypointGroupID
	gvar11 ns.WaypointGroupID
	wp12   ns.WaypointID
	gvar13 int
	flag14 bool
	flag15 bool
	gvar16 int
	obj17  ns.ObjectID
	obj18  ns.ObjectID
	gvar19 ns.WallGroupID
	gvar20 ns.WallGroupID
	gvar21 ns.WallGroupID
	wp22   ns.WaypointID
	gvar23 ns.WaypointGroupID
	gvar24 ns.WaypointGroupID
	obj25  ns.ObjectID
	obj26  ns.ObjectID
	gvar27 ns.WallGroupID
	gvar28 ns.WallGroupID
	gvar29 ns.WallGroupID
	wp30   ns.WaypointID
	wp31   ns.WaypointID
	wp32   ns.WaypointID
	wp33   ns.WaypointID
	wp34   ns.WaypointID
	gvar35 ns.WaypointGroupID
	gvar36 ns.WaypointGroupID
	gvar37 ns.WaypointGroupID
	flag38 bool
	wp39   ns.WaypointID
	obj40  ns.ObjectID
	fvar41 float32
	fvar42 float32
	flag43 bool
	flag44 bool
	flag45 bool
	flag46 bool
)

func init() {
	flag15 = false
	gvar16 = 0
	flag14 = false
	flag43 = false
	flag38 = false
	flag44 = false
	flag45 = false
	flag46 = false
}
func StartTTRoom() {
	ns.JournalEntry(ns.GetHost(), "War11Hecubah", 2)
	ns.PrintToAll("Con01a:NewJournalEntry")
	ns.WideScreen(true)
	ns.Frozen(ns.GetHost(), true)
	ns.Frozen(obj4, true)
	ns.Frozen(obj5, true)
	ns.Frozen(obj6, true)
	ns.Frozen(obj7, true)
	ns.Frozen(obj8, true)
	ns.Frozen(obj9, true)
	ns.CreatureIdle(obj4)
	ns.CreatureIdle(ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj4)
	ns.LookAtObject(obj4, ns.GetHost())
	ns.FrameTimer(1, HecTTTalk)
}
func HecTTTalk() {
	ns.SetDialog(obj4, ns.NORMAL, HecTTStart, HecTTEnd)
	ns.StartDialog(obj4, ns.GetHost())
}
func HecTTStart() {
	ns.TellStory(ns.HumanMaleEatFood, "War11a:CatchMe")
}
func HecTTEnd() {
	ns.CancelDialog(obj4)
	ns.Frozen(obj4, false)
	ns.Walk(obj4, 1321, 4793)
	ns.Music(27, 100)
}
func EnterBlinkRoom() {
	if !ns.IsCaller(obj4) {
		goto LABEL1
	}
	ns.MoveObject(obj4, 1052, 3866)
	ns.Frozen(obj4, true)
	ns.WideScreen(false)
	ns.Frozen(ns.GetHost(), false)
LABEL1:
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL2
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.WideScreen(true)
	ns.Frozen(ns.GetHost(), true)
	ns.CreatureIdle(obj4)
	ns.CreatureIdle(ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj4)
	ns.LookAtObject(obj4, ns.GetHost())
	ns.FrameTimer(15, HecBlinkTalk)
LABEL2:
	return
}
func HecBlinkTalk() {
	ns.SetDialog(obj4, ns.NORMAL, HecBlinkStart, HecBlinkEnd)
	ns.StartDialog(obj4, ns.GetHost())
}
func HecBlinkStart() {
	ns.TellStory(ns.HumanMaleEatFood, "Con10b.scr:HecubahDialog1")
}
func HecBlinkEnd() {
	ns.CancelDialog(obj4)
	ns.Frozen(obj4, false)
	ns.Walk(obj4, 1140, 3785)
}
func EnterArena() {
	if !ns.IsCaller(obj4) {
		goto LABEL1
	}
	ns.WideScreen(false)
	ns.Frozen(ns.GetHost(), false)
	ns.MoveObject(obj4, 4378, 1925)
	ns.CreatureGuard(obj4, 4378, 1925, 0, 0, 0)
LABEL1:
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL2
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.WayPointGroupOff(gvar10)
	ns.WayPointGroupOn(gvar11)
LABEL2:
	return
}
func TempInit() {
	ns.SetDialog(obj4, ns.NORMAL, HecTTStart, HecTTEnd)
}
func Hunting() {
	if !(ns.CurrentHealth(ns.GetTrigger()) > 0) {
		goto LABEL1
	}
	ns.CreatureHunt(ns.GetTrigger())
LABEL1:
	return
}
func Lich1Die() {
}
func Lich2Die() {
}
func Lich3Die() {
}
func Lich4Die() {
}
func Lich1Setup() {
	ns.SetCallback(obj5, 13, Hunting)
	ns.SetCallback(obj5, 5, Lich1Die)
	ns.SetCallback(obj5, 4, Hunting)
	ns.CreatureHunt(obj5)
	ns.SetOwner(obj4, obj5)
}
func Lich2Setup() {
	ns.SetCallback(obj6, 13, Hunting)
	ns.SetCallback(obj6, 5, Lich2Die)
	ns.SetCallback(obj6, 4, Hunting)
	ns.CreatureHunt(obj6)
	ns.SetOwner(obj4, obj6)
}
func Lich3Setup() {
	ns.SetCallback(obj7, 13, Hunting)
	ns.SetCallback(obj7, 5, Lich3Die)
	ns.SetCallback(obj7, 4, Hunting)
	ns.CreatureHunt(obj7)
	ns.SetOwner(obj4, obj7)
}
func Lich4Setup() {
	ns.SetCallback(obj8, 13, Hunting)
	ns.SetCallback(obj8, 5, Lich4Die)
	ns.SetCallback(obj8, 4, Hunting)
	ns.CreatureHunt(obj8)
	ns.SetOwner(obj4, obj8)
}
func HecLoseSight() {
	ns.CreatureHunt(ns.GetTrigger())
}
func OpenBlinkWalls1() {
	ns.WallGroupOpen(gvar28)
}
func CloseBlinkWalls1() {
	ns.WallGroupClose(gvar28)
}
func OpenBlinkWalls2() {
	ns.WallGroupOpen(gvar29)
}
func CloseBlinkWalls2() {
	ns.WallGroupClose(gvar29)
}
func OpeningSetPiece() {
	ns.Frozen(obj5, true)
	ns.Frozen(obj6, true)
	ns.Frozen(obj7, true)
	ns.Frozen(obj8, true)
	ns.Frozen(ns.GetHost(), true)
	ns.WideScreen(true)
	ns.LookAtObject(ns.GetHost(), obj4)
}
func Hecubah1Start() {
	ns.TellStory(ns.HumanMaleEatFood, "Wiz11A.scr:HecubahTalk01")
}
func Hecubah1End() {
	ns.CancelDialog(obj4)
	ns.WideScreen(false)
	ns.Frozen(ns.GetHost(), false)
	ns.Frozen(obj4, false)
	ns.Frozen(obj5, false)
	ns.Frozen(obj6, false)
	ns.Frozen(obj7, false)
	ns.Frozen(obj8, false)
	ns.AggressionLevel(obj4, 0.83)
	ns.SetCallback(obj4, 13, HecLoseSight)
	ns.SetCallback(obj4, 4, Hunting)
}
func SetPieceInit() {
	wp39 = ns.Waypoint("HecubahWP")
	obj40 = ns.Object("HecubahLight")
}
func OpenBlinkWallsA() {
	ns.WallGroupOpen(gvar20)
}
func OpenBlinkWallsB() {
	ns.WallGroupOpen(gvar21)
}
func MapInitialize() {
	flag38 = true
	obj25 = ns.Object("NewTeleporter")
	obj26 = ns.Object("TeleportLight")
	gvar27 = ns.WallGroup("TeleportWalls")
	gvar28 = ns.WallGroup("BlinkWalls1")
	gvar29 = ns.WallGroup("BlinkWalls2")
	gvar20 = ns.WallGroup("BlinkWallsA")
	gvar21 = ns.WallGroup("BlinkWallsB")
	wp30 = ns.Waypoint("OpenDoorWP")
	wp31 = ns.Waypoint("TeleportNewWP")
	wp32 = ns.Waypoint("SetPieceStartWP")
	wp33 = ns.Waypoint("BattleStartWP")
	wp34 = ns.Waypoint("LeaveWP")
	gvar35 = ns.WaypointGroup("EscapeWP")
	gvar36 = ns.WaypointGroup("BlinkBookWP")
	gvar10 = ns.WaypointGroup("BlinkWP")
	gvar37 = ns.WaypointGroup("Arena2WP")
	gvar11 = ns.WaypointGroup("CheeseWP")
	obj4 = ns.Object("Hecubah")
	obj5 = ns.Object("Lich1")
	obj6 = ns.Object("Lich2")
	obj7 = ns.Object("Lich3")
	obj8 = ns.Object("Lich4")
	obj9 = ns.Object("Lich5")
	obj17 = ns.Object("FinalTeleporter")
	obj18 = ns.Object("TPINTrigger")
	gvar19 = ns.WallGroup("FinaleWalls")
	wp12 = ns.Waypoint("TeleportWP")
	wp22 = ns.Waypoint("PlayerSounds")
	gvar23 = ns.WaypointGroup("ArenaWP")
	gvar24 = ns.WaypointGroup("FacadeWP")
	ns.WayPointGroupOff(gvar24)
	ns.WayPointGroupOff(gvar23)
	ns.WayPointGroupOff(gvar37)
	ns.WayPointGroupOff(gvar36)
	ns.WayPointGroupOff(gvar11)
	ns.WayPointGroupOff(gvar35)
	ns.StoryPic(obj4, "HecubahPic")
	ns.SetDialog(obj4, ns.NORMAL, Hecubah1Start, Hecubah1End)
	SetPieceInit()
	ns.StartupScreen(11)
	ns.FrameTimer(10, StartTTRoom)
}
func PlayerDeath() {
	ns.DeathScreen(11)
}
func HecubahDies() {
	fvar41 = ns.GetObjectX(obj4)
	fvar42 = ns.GetObjectY(obj4)
	ns.MoveWaypoint(wp22, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.Music(0, 100)
	ns.AudioEvent(ns.HecubahDieFrame0A, wp22)
	ns.FrameTimer(66, HecubahDie1)
	ns.FrameTimer(130, HecubahDie2)
	ns.FrameTimer(170, HecubahDie3)
	ns.FrameTimer(280, HecubahDie4)
}
func HecubahDie1() {
	ns.AudioEvent(ns.HecubahDieFrame98, wp22)
}
func HecubahDie2() {
	ns.AudioEvent(ns.HecubahDieFrame194, wp22)
}
func HecubahDie3() {
	ns.AudioEvent(ns.HecubahDieFrame283, wp22)
}
func HecubahDie4() {
	ns.AudioEvent(ns.HecubahDieFrame439, wp22)
	ns.MoveObject(obj40, fvar41, fvar42)
}
func OpenDoors() {
	if !ns.IsCaller(obj4) {
		goto LABEL1
	}
	if !flag43 {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.WallGroupOpen(gvar27)
	ns.Move(obj4, wp31)
	ns.MoveWaypoint(wp22, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.BoulderMove, wp22)
	CheckPlayerZone()
LABEL1:
	return
}
func StartTeleporter() {
	if !ns.IsCaller(obj4) {
		goto LABEL1
	}
	if flag38 {
		goto LABEL2
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.MoveObject(obj26, ns.GetWaypointX(wp31), ns.GetWaypointY(wp31))
	ns.ObjectOn(obj25)
	goto LABEL3
LABEL2:
	ns.ObjectOff(ns.GetTrigger())
	ns.ObjectOn(obj25)
	ns.MoveObject(obj26, ns.GetWaypointX(wp31), ns.GetWaypointY(wp31))
	ns.CreatureGuard(obj4, ns.GetWaypointX(wp31), ns.GetWaypointY(wp31), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
LABEL3:
	ns.ObjectOn(obj5)
	ns.ObjectOn(obj6)
	ns.ObjectOn(obj7)
	ns.ObjectOn(obj8)
	ns.ObjectOn(obj9)
	ns.WideScreen(false)
	ns.Frozen(ns.GetHost(), false)
LABEL1:
	return
}
func StartIntro() {
	ns.Frozen(obj4, true)
	ns.CreatureIdle(obj4)
	ns.LookAtObject(obj4, ns.GetHost())
	ns.StartDialog(obj4, ns.GetHost())
}
func HecubahCroak() {
	ns.ImmediateBlind()
	ns.Frozen(obj4, true)
	ns.Frozen(ns.GetHost(), true)
	ns.FrameTimer(20, EndPlaces)
}
func FreezePlayer() {
	ns.Frozen(ns.GetHost(), true)
	ns.Frozen(obj4, true)
	ns.LookAtObject(ns.GetHost(), obj4)
	ns.FrameTimer(40, DeathPieceFadeIn)
}
func SetNPCPositions() {
	ns.MoveObject(ns.GetHost(), 4142, 4364)
	ns.MoveObject(ns.GetHost(), 4249, 4342)
	ns.FrameTimer(10, FreezePlayer)
}
func DeathPieceFadeIn() {
	ns.UnBlind()
	ns.FrameTimer(60, HecubahDeathTalk)
}
func HecubahDeath1Start() {
	ns.TellStory(ns.HumanMaleEatFood, "Wiz11A.scr:HecubahTalk07")
}
func HecubahDeath1End() {
	ns.CancelDialog(obj4)
	ns.Frozen(obj4, false)
	ns.WideScreen(false)
	HecubahDies()
	ns.SecondTimer(13, FinalFadeOut)
}
func HecubahDeathTalk() {
	ns.SetDialog(obj4, ns.NORMAL, HecubahDeath1Start, HecubahDeath1End)
	ns.StartDialog(obj4, ns.GetHost())
}
func FinalFadeOut() {
	ns.Blind()
	ns.EndGame(1)
}
func HecubahRun() {
	if !(ns.CurrentHealth(obj4) <= 150) {
		goto LABEL1
	}
	if !(flag43 == false && flag38 == false) {
		goto LABEL2
	}
	flag43 = true
	ns.Enchant(obj4, ns.ENCHANT_INVULNERABLE, 0)
	ns.AggressionLevel(obj4, 0)
	ns.Move(obj4, wp30)
LABEL2:
	if !(flag43 == false && flag38 == true) {
		goto LABEL1
	}
	flag43 = true
	ns.AggressionLevel(obj4, 0)
	ns.Enchant(obj4, ns.ENCHANT_INVULNERABLE, 0)
	if !ns.HasEnchant(obj4, ns.ENCHANT_INVISIBLE) {
		goto LABEL3
	}
	ns.EnchantOff(obj4, ns.ENCHANT_INVISIBLE)
LABEL3:
	ns.MoveWaypoint(wp22, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.TeleportOut, wp22)
	ns.Effect(ns.SMOKE_BLAST, ns.GetObjectX(obj4), ns.GetObjectY(obj4), 0, 0)
	ns.Effect(ns.SPARK_EXPLOSION, ns.GetObjectX(obj4), ns.GetObjectY(obj4), 0, 0)
	ns.MoveObject(obj4, ns.GetWaypointX(wp34), ns.GetWaypointY(wp34))
	ns.Frozen(obj4, true)
	ns.CreatureIdle(obj4)
	ns.LookAtObject(obj4, ns.ObjectID(wp32)) // FIXME
	CheckPlayerZone()
LABEL1:
	return
}
func SetEveryone() {
	ns.UnBlind()
	ns.FrameTimer(1, FreezeEveryone)
}
func FreezeEveryone() {
	ns.Frozen(ns.GetHost(), true)
	ns.CreatureIdle(ns.GetHost())
	ns.Frozen(obj4, true)
	ns.Frozen(obj5, true)
	ns.Frozen(obj6, true)
	ns.Frozen(obj7, true)
	ns.Frozen(obj8, true)
	ns.LookAtObject(obj4, ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj4)
	ns.SetDialog(obj4, ns.NEXT, HecubahTeleport1Start, HecubahTeleport1End)
	ns.StartDialog(obj4, ns.GetHost())
}
func HecDeathRayMiss() {
	ns.CastSpellLocationLocation(ns.SPELL_DEATH_RAY, ns.GetObjectX(obj4), ns.GetObjectY(obj4), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost())-50)
	ns.SetDialog(obj4, ns.NEXT, HecFacade1Start, HecFacade1End)
	ns.FrameTimer(10, StartFacadeSetPiece)
}
func HecubahTeleport1Start() {
	ns.TellStory(ns.HumanMaleEatFood, "Wiz11A.scr:HecubahTalk05")
}
func HecubahTeleport1End() {
	ns.SetDialog(obj4, ns.NORMAL, HecubahTeleport2Start, HecubahTeleport2End)
	ns.StartDialog(obj4, ns.GetHost())
}
func HecubahTeleport2Start() {
	ns.TellStory(ns.HumanMaleEatFood, "Wiz11A.scr:HecubahTalk06")
}
func HecubahTeleport2End() {
	ns.ObjectOn(obj17)
	ns.Frozen(obj4, false)
	ns.Frozen(ns.GetHost(), false)
	ns.Frozen(obj5, false)
	ns.Frozen(obj6, false)
	ns.Frozen(obj7, false)
	ns.Frozen(obj8, false)
}
func SetHecubahPosition() {
	var (
		v0 int
		v1 int
	)
	if !ns.IsCaller(obj4) {
		goto LABEL1
	}
	flag14 = true
	ns.WayPointGroupOff(gvar35)
	v0 = ns.CurrentHealth(obj4)
	v1 = v0 - 300
	ns.Damage(obj4, 0, v1, 0)
	ns.MoveObject(obj4, 1462, 1322)
LABEL1:
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL2
	}
	ns.WayPointGroupOff(gvar23)
	ns.Frozen(ns.GetHost(), true)
	ns.LookAtObject(ns.GetHost(), obj4)
	ns.LookAtObject(obj4, ns.GetHost())
	ns.ObjectOff(obj18)
	ns.FrameTimer(10, HecDeathRayMiss)
LABEL2:
	return
}
func FreezeFacadeHecubah() {
	if !ns.IsCaller(obj4) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.Frozen(obj4, true)
	ns.CreatureIdle(obj4)
LABEL1:
	return
}
func StartFacadeSetPiece() {
	ns.StartDialog(obj4, ns.GetHost())
}
func HecFacade1Start() {
	ns.TellStory(ns.HumanMaleEatFood, "Wiz11A.scr:HecubahTalk02")
}
func HecFacade1End() {
	ns.SetDialog(obj4, ns.NORMAL, HecFacade2Start, HecFacade2End)
	ns.StartDialog(obj4, ns.GetHost())
}
func HecFacade2Start() {
	ns.TellStory(ns.HumanMaleEatFood, "Wiz11A.scr:HecubahTalk03")
}
func HecFacade2End() {
	ns.CancelDialog(obj4)
	ns.WayPointGroupOn(gvar24)
	ns.WayPointGroupOff(gvar37)
	ns.WayPointGroupOff(gvar36)
	ns.AggressionLevel(obj4, 0.83)
	ns.Frozen(obj4, false)
	ns.Frozen(ns.GetHost(), false)
}
func AlternateStep1() {
	ns.MoveObject(ns.GetHost(), 4841, 1183)
	ns.Frozen(ns.GetHost(), true)
	ns.CreatureIdle(ns.GetHost())
	ns.LookAtObject(ns.GetHost(), ns.ObjectID(wp34)) // FIXME
	ns.MoveObject(obj4, ns.GetWaypointX(wp30), ns.GetWaypointY(wp30))
	ns.Frozen(obj4, true)
	ns.CreatureIdle(obj4)
	ns.LookAtObject(obj4, ns.GetHost())
	ns.UnBlind()
	ns.FrameTimer(30, AlternateStep2)
}
func AlternateStep2() {
	ns.WideScreen(true)
	ns.SetDialog(obj4, ns.NEXT, AltStart, AltEnd)
	ns.StartDialog(obj4, ns.GetHost())
}
func AltStart() {
	ns.TellStory(ns.HumanMaleEatFood, "Wiz11A.scr:HecubahTalk05")
}
func AltEnd() {
	ns.SetDialog(obj4, ns.NORMAL, Alt2Start, Alt2End)
	ns.StartDialog(obj4, ns.GetHost())
}
func Alt2Start() {
	ns.TellStory(ns.HumanMaleEatFood, "Wiz11A.scr:HecubahTalk06")
}
func Alt2End() {
	flag43 = true
	ns.AggressionLevel(obj4, 0)
	ns.Frozen(obj4, false)
	ns.ObjectOn(obj25)
	ns.Walk(obj4, ns.GetWaypointX(wp30), ns.GetWaypointY(wp30))
}
func EnterZone() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	flag44 = true
LABEL1:
	return
}
func ExitZone() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	flag44 = false
LABEL1:
	return
}
func EndPlaces() {
	ns.Frozen(obj4, false)
	ns.Frozen(ns.GetHost(), false)
	ns.MoveObject(obj4, 1462, 1322)
	ns.MoveObject(ns.GetHost(), 1550, 1430)
	ns.FrameTimer(1, EndFreeze)
}
func EndFreeze() {
	ns.Frozen(obj4, true)
	ns.Frozen(ns.GetHost(), true)
	ns.LookAtObject(obj4, ns.GetHost())
	ns.CreatureIdle(ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj4)
	ns.UnBlind()
	ns.WideScreen(true)
	ns.FrameTimer(30, HecubahDeathTalk)
}
func CheckPlayerZone() {
	if !(flag43 == true && flag44 == true) {
		goto LABEL1
	}
	ns.Frozen(ns.GetHost(), true)
	ns.WideScreen(true)
	ns.CreatureIdle(ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj4)
	if !(ns.CurrentHealth(obj5) > 0) {
		goto LABEL2
	}
	ns.ObjectOff(obj5)
LABEL2:
	if !(ns.CurrentHealth(obj6) > 0) {
		goto LABEL3
	}
	ns.ObjectOff(obj6)
LABEL3:
	if !(ns.CurrentHealth(obj7) > 0) {
		goto LABEL4
	}
	ns.ObjectOff(obj7)
LABEL4:
	if !(ns.CurrentHealth(obj8) > 0) {
		goto LABEL5
	}
	ns.ObjectOff(obj8)
LABEL5:
	if !(ns.CurrentHealth(obj9) > 0) {
		goto LABEL6
	}
	ns.ObjectOff(obj9)
LABEL6:
	ns.WayPointGroupOn(gvar35)
	ns.Frozen(obj4, false)
	ns.Move(obj4, wp30)
	goto LABEL7
LABEL1:
	ns.FrameTimer(1, CheckPlayerZone)
LABEL7:
	return
}
func StartOpeningPiece() {
	flag46 = true
	ns.Frozen(ns.GetHost(), true)
	ns.WideScreen(true)
	ns.CreatureIdle(ns.GetHost())
	ns.LookAtObject(ns.GetHost(), ns.ObjectID(wp32)) // FIXME
	ns.MoveWaypoint(wp22, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.TeleportIn, wp22)
	ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp32), ns.GetWaypointY(wp32), 0, 0)
	ns.Effect(ns.SPARK_EXPLOSION, ns.GetWaypointX(wp32), ns.GetWaypointY(wp32), 0, 0)
	ns.MoveObject(obj4, ns.GetWaypointX(wp32), ns.GetWaypointY(wp32))
	ns.Frozen(obj4, true)
	ns.CreatureIdle(obj4)
	ns.LookAtObject(obj4, ns.GetHost())
	ns.SetDialog(obj4, ns.NORMAL, HecubahOpenStart, HecubahOpenEnd)
	ns.StartDialog(obj4, ns.GetHost())
}
func HecubahOpenStart() {
	ns.TellStory(ns.HumanMaleEatFood, "Wiz11A.scr:HecubahTalk01")
}
func HecubahOpenEnd() {
	ns.CancelDialog(obj4)
	ns.Frozen(obj4, false)
	ns.AudioEvent(ns.TeleportOut, wp22)
	ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp32), ns.GetWaypointY(wp32), 0, 0)
	ns.Effect(ns.SPARK_EXPLOSION, ns.GetWaypointX(wp32), ns.GetWaypointY(wp32), 0, 0)
	ns.MoveObject(obj4, ns.GetWaypointX(wp33), ns.GetWaypointY(wp33))
	ns.FrameTimer(15, ReleasePlayer)
}
func ReleasePlayer() {
	ns.Frozen(obj5, false)
	ns.Frozen(obj6, false)
	ns.Frozen(obj7, false)
	ns.Frozen(obj8, false)
	ns.Frozen(obj9, false)
	ns.AggressionLevel(obj4, 0.83)
	ns.AggressionLevel(obj5, 0.83)
	ns.AggressionLevel(obj6, 0.83)
	ns.AggressionLevel(obj7, 0.83)
	ns.AggressionLevel(obj8, 0.83)
	ns.AggressionLevel(obj9, 0.83)
	ns.CreatureHunt(obj4)
	ns.CreatureHunt(obj5)
	ns.CreatureHunt(obj6)
	ns.CreatureHunt(obj7)
	ns.CreatureHunt(obj8)
	ns.CreatureHunt(obj9)
	ns.WideScreen(false)
	ns.Frozen(ns.GetHost(), false)
	ns.Music(16, 100)
}
func StartHecTalk() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	if flag46 {
		goto LABEL1
	}
	ns.Music(11, 100)
	flag46 = true
	ns.WayPointGroupOn(gvar37)
	ns.WayPointGroupOn(gvar36)
	ns.Frozen(ns.GetHost(), true)
	ns.WideScreen(true)
	ns.CreatureIdle(ns.GetHost())
	ns.LookAtObject(ns.GetHost(), ns.ObjectID(wp32)) // FIXME
	ns.SecondTimer(1, StartOpeningPiece)
LABEL1:
	return
}
func CheckHecHealth() {
	ns.PrintToAll(strconv.Itoa(ns.CurrentHealth(obj4)))
	if !(ns.CurrentHealth(obj4) > 0) {
		goto LABEL1
	}
	ns.FrameTimer(5, CheckHecHealth)
LABEL1:
	return
}
func OnEvent(typ string) {
	switch typ {
	case "MapInitialize":
		MapInitialize()
	case "PlayerDeath":
		PlayerDeath()
	}
}
