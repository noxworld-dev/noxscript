package wiz07e

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
	fvar22 float32
	fvar23 float32
	fvar24 float32
	fvar25 float32
	fvar26 float32
	fvar27 float32
	fvar28 float32
	fvar29 float32
	fvar30 float32
	fvar31 float32
	fvar32 float32
	fvar33 float32
	fvar34 float32
	fvar35 float32
	fvar36 float32
	fvar37 float32
	fvar38 float32
	fvar39 float32
	obj40  ns.ObjectID
	obj41  ns.ObjectID
	obj42  ns.ObjectID
	obj43  ns.ObjectID
	obj44  ns.ObjectID
	obj45  ns.ObjectID
	obj46  ns.ObjectID
	flag47 bool
	gvar48 ns.WallGroupID
	wp49   ns.WaypointID
	wp50   ns.WaypointID
	wp51   ns.WaypointID
	wp52   ns.WaypointID
	wp53   ns.WaypointID
	gvar54 int
	gvar55 int
	gvar56 int
	gvar57 int
	gvar58 int
	gvar59 int
	gvar60 int
	gvar61 int
	gvar62 int
	gvar63 int
	gvar64 int
)

func init() {
	flag47 = false
	gvar54 = 0
	gvar55 = 1
	gvar56 = 2
	gvar57 = 3
	gvar58 = 0
	gvar59 = 1
	gvar60 = 2
	gvar61 = 3
	gvar62 = 4
}
func NoxHeartInit() {
	obj4 = ns.Object("HeartLight1")
	obj5 = ns.Object("HeartLight2")
	obj6 = ns.Object("HeartLight3")
	obj7 = ns.Object("HeartLight4")
	obj8 = ns.Object("HeartLight5")
	obj9 = ns.Object("HeartLight6")
	obj10 = ns.Object("HeartLight7")
	obj11 = ns.Object("HeartLight8")
	obj12 = ns.Object("HeartLight9")
	obj13 = ns.Object("NewLight1")
	obj14 = ns.Object("NewLight2")
	obj15 = ns.Object("NewLight3")
	obj16 = ns.Object("NewLight4")
	obj17 = ns.Object("NewLight5")
	obj18 = ns.Object("NewLight6")
	obj19 = ns.Object("NewLight7")
	obj20 = ns.Object("NewLight8")
	obj21 = ns.Object("NewLight9")
	fvar22 = ns.GetObjectX(obj4)
	fvar23 = ns.GetObjectX(obj5)
	fvar24 = ns.GetObjectX(obj6)
	fvar25 = ns.GetObjectX(obj7)
	fvar26 = ns.GetObjectX(obj8)
	fvar27 = ns.GetObjectX(obj9)
	fvar28 = ns.GetObjectX(obj10)
	fvar29 = ns.GetObjectX(obj11)
	fvar30 = ns.GetObjectX(obj12)
	fvar31 = ns.GetObjectY(obj4)
	fvar32 = ns.GetObjectY(obj5)
	fvar33 = ns.GetObjectY(obj6)
	fvar34 = ns.GetObjectY(obj7)
	fvar35 = ns.GetObjectY(obj8)
	fvar36 = ns.GetObjectY(obj9)
	fvar37 = ns.GetObjectY(obj10)
	fvar38 = ns.GetObjectY(obj11)
	fvar39 = ns.GetObjectY(obj12)
}
func Sparklies() {
	ns.Frozen(ns.GetHost(), true)
	ns.CreatureIdle(ns.GetHost())
	ns.ObjectOff(obj4)
	ns.ObjectOff(obj5)
	ns.ObjectOff(obj6)
	ns.ObjectOff(obj7)
	ns.ObjectOff(obj8)
	ns.ObjectOff(obj9)
	ns.ObjectOff(obj10)
	ns.ObjectOff(obj11)
	ns.ObjectOff(obj12)
	ns.MoveObject(obj13, fvar22, fvar31)
	ns.MoveObject(obj14, fvar23, fvar32)
	ns.MoveObject(obj15, fvar24, fvar33)
	ns.MoveObject(obj16, fvar25, fvar34)
	ns.MoveObject(obj17, fvar26, fvar35)
	ns.MoveObject(obj18, fvar27, fvar36)
	ns.MoveObject(obj19, fvar28, fvar37)
	ns.MoveObject(obj20, fvar29, fvar38)
	ns.MoveObject(obj21, fvar30, fvar39)
}
func CreatureSetup() {
	ns.Music(1, 100)
	ns.SetOwner(ns.GetHost(), obj42)
	ns.LookAtObject(obj42, ns.GetHost())
	ns.FrameTimer(1, NoxHeartInit)
	ns.StartDialog(obj42, ns.GetHost())
}
func HONTrigger() {
	ns.ObjectOff(obj44)
	ns.ObjectOff(obj45)
	Sparklies()
	ns.SetHalberd(1)
	ns.MoveWaypoint(wp53, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.StaffOblivionAchieve2, wp53)
	ns.JournalEdit(ns.GetHost(), "GetHON", 4)
	ns.PrintToAll("Con07H.scr:HONTada")
	ns.WallGroupOpen(gvar48)
	ns.FrameTimer(1, EndSparklies)
}
func HorvathStart() {
	var v0 int
	v0 = gvar63
	if v0 == gvar54 {
		goto LABEL1
	}
	if v0 == gvar55 {
		goto LABEL2
	}
	if v0 == gvar56 {
		goto LABEL3
	}
	if v0 == gvar59 {
		goto LABEL4
	}
	goto LABEL5
LABEL1:
	ns.TellStory(ns.HumanMaleEatFood, "Wiz07.scr:HorvathTalk01")
	goto LABEL5
LABEL2:
	ns.TellStory(ns.HumanMaleEatFood, "Wiz07.scr:HorvathTalk02")
	goto LABEL5
LABEL3:
	ns.TellStory(ns.HumanMaleEatFood, "Wiz07.scr:HorvathTalk03")
	goto LABEL5
LABEL4:
	ns.TellStory(ns.HumanMaleEatFood, "Wiz07.scr:HorvathTalk04")
	goto LABEL5
LABEL5:
	return
}
func HorvathEnd() {
	var v0 int
	v0 = gvar63
	if v0 == gvar54 {
		goto LABEL1
	}
	if v0 == gvar55 {
		goto LABEL2
	}
	if v0 == gvar56 {
		goto LABEL3
	}
	goto LABEL4
LABEL1:
	ns.CreatureFollow(obj42, ns.GetHost())
	ns.CancelDialog(obj42)
	gvar63 = gvar55
	goto LABEL4
LABEL2:
	ns.Move(obj42, wp49)
	ns.CancelDialog(obj42)
	gvar63 = gvar56
	goto LABEL4
LABEL3:
	ns.SetDialog(obj42, ns.NORMAL, HorvathStart, HorvathEnd)
	goto LABEL4
LABEL4:
	return
}
func Horvath2Start() {
	ns.TellStory(ns.HumanMaleEatFood, "Wiz07.scr:HorvathTalk04")
}
func Horvath2End() {
	ns.SetDialog(obj43, ns.NEXT, Hecubah2Start, Hecubah2End)
	ns.StartDialog(obj43, ns.GetHost())
}
func HorvathTrigger() {
	if !ns.IsCaller(obj42) {
		goto LABEL1
	}
	flag47 = true
	ns.CreatureGuard(obj42, ns.GetWaypointX(wp49), ns.GetWaypointY(wp49), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	ns.SetDialog(obj42, ns.NORMAL, HorvathStart, HorvathEnd)
LABEL1:
	return
}
func StartSetPiece() {
	if !flag47 {
		goto LABEL1
	}
	ns.Music(7, 100)
	ns.ObjectOff(obj46)
	ns.WideScreen(true)
	ns.Frozen(ns.GetHost(), true)
	ns.Enchant(ns.GetHost(), ns.ENCHANT_INVISIBLE, 0)
	ns.LookAtObject(ns.GetHost(), obj42)
	ns.Effect(ns.LESSER_EXPLOSION, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0, 0)
	ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp50), ns.GetWaypointY(wp50), 0, 0)
	ns.Effect(ns.LESSER_EXPLOSION, ns.GetWaypointX(wp50), ns.GetWaypointY(wp50), 0, 0)
	ns.AudioEvent(ns.TeleportIn, wp50)
	ns.MoveObject(obj43, ns.GetWaypointX(wp50), ns.GetWaypointY(wp50))
	ns.CreatureGuard(obj43, ns.GetObjectX(obj43), ns.GetObjectY(obj43), ns.GetObjectX(obj42), ns.GetObjectY(obj42), 0)
	ns.CreatureGuard(obj42, ns.GetObjectX(obj42), ns.GetObjectY(obj42), ns.GetObjectX(obj43), ns.GetObjectY(obj43), 0)
	ns.StartDialog(obj43, ns.GetHost())
	goto LABEL2
LABEL1:
	ns.FrameTimer(2, StartSetPiece)
LABEL2:
	return
}
func TeleportEffect() {
}
func Hecubah1Start() {
	ns.TellStory(ns.HumanMaleEatFood, "Wiz07.scr:HecubahTalk01")
}
func Hecubah1End() {
	ns.SetDialog(obj42, ns.NEXT, Horvath2Start, Horvath2End)
	ns.StartDialog(obj42, ns.GetHost())
}
func Hecubah2Start() {
	ns.TellStory(ns.HumanMaleEatFood, "Wiz07.scr:HecubahTalk02")
}
func Hecubah2End() {
	ns.CancelDialog(obj43)
	ns.CancelDialog(obj42)
	ns.CastSpellLocationLocation(ns.SPELL_DEATH_RAY, ns.GetObjectX(obj43), ns.GetObjectY(obj43), ns.GetObjectX(obj42), ns.GetObjectY(obj42))
	ns.AggressionLevel(obj42, 0.83)
	ns.SecondTimer(2, SecondAttack)
}
func Hecubah3Start() {
	ns.TellStory(ns.HumanMaleEatFood, "Wiz07.scr:HecubahTalk03")
}
func Hecubah3End() {
	ns.Frozen(obj43, false)
	ns.CreatureGuard(obj43, ns.GetObjectX(obj43), ns.GetObjectY(obj43), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	ns.CancelDialog(obj43)
	ns.Blind()
	ns.FrameTimer(60, SetPlayerPosition)
}
func HorvathDeath() {
	ns.Chat(obj42, "Wiz07.scr:HorvathTalk05")
	ns.CastSpellLocationLocation(ns.SPELL_TURN_UNDEAD, ns.GetObjectX(obj42), ns.GetObjectY(obj42), ns.GetObjectX(obj42), ns.GetObjectY(obj42))
	ns.AggressionLevel(obj43, 0.16)
	ns.Move(obj43, wp51)
	ns.SecondTimer(2, EraseHorvath)
}
func HecubahGloat() {
	if !ns.IsCaller(obj43) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.Frozen(obj43, true)
	ns.CreatureIdle(obj43)
	ns.LookAtObject(obj43, ns.GetHost())
	ns.SetDialog(obj43, ns.NORMAL, Hecubah3Start, Hecubah3End)
	ns.StartDialog(obj43, ns.GetHost())
LABEL1:
	return
}
func SetPlayerPosition() {
	ns.SetQuestStatus(1, "HorvathDied")
	ns.Music(0, 100)
	ns.WideScreen(false)
	ns.Frozen(ns.GetHost(), false)
	ns.MoveObject(ns.GetHost(), 3160, 4613)
}
func EndSparklies() {
	ns.Frozen(ns.GetHost(), false)
	ns.SetDialog(obj42, ns.NORMAL, HorvathStart, HorvathEnd)
	ns.StartDialog(obj42, ns.GetHost())
}
func SecondAttack() {
	ns.CastSpellLocationLocation(ns.SPELL_DEATH_RAY, ns.GetObjectX(obj43), ns.GetObjectY(obj43), ns.GetObjectX(obj42), ns.GetObjectY(obj42))
	if !(ns.CurrentHealth(obj42) > 0) {
		goto LABEL1
	}
	ns.SecondTimer(1, SecondAttack)
LABEL1:
	return
}
func SaveGameFix() {
	ns.Music(0, 100)
	ns.WideScreen(false)
	ns.Frozen(ns.GetHost(), false)
	ns.MoveObject(ns.GetHost(), 3160, 4613)
}
func SaveGameFade() {
	ns.Blind()
	ns.FrameTimer(60, SaveGameFix)
}
func EraseHorvath() {
	ns.Delete(obj42)
}
func MapInitialize() {
	wp52 = ns.Waypoint("HONWP")
	obj40 = ns.Object("HONDoor1")
	obj41 = ns.Object("HONDoor2")
	obj42 = ns.Object("Horvath")
	obj43 = ns.Object("Hecubah")
	obj44 = ns.Object("HON")
	obj45 = ns.Object("HONBase")
	obj46 = ns.Object("SetPieceTrigger")
	gvar48 = ns.WallGroup("HONWalls")
	wp49 = ns.Waypoint("HorvathWP")
	wp50 = ns.Waypoint("HecubahWP")
	wp51 = ns.Waypoint("HecubahGloatWP")
	wp53 = ns.Waypoint("PlayerSounds")
	gvar63 = gvar54
	gvar64 = gvar58
	ns.LockDoor(obj40)
	ns.LockDoor(obj41)
	ns.SetDialog(obj42, ns.NORMAL, HorvathStart, HorvathEnd)
	ns.SetDialog(obj43, ns.NEXT, Hecubah1Start, Hecubah1End)
	ns.StoryPic(obj42, "HorvathPic")
	ns.StoryPic(obj43, "HecubahPic")
	ns.FrameTimer(10, CreatureSetup)
}
func MapEntry() {
	if ns.GetQuestStatus("HorvathDied") != 1 {
		goto LABEL1
	}
	ns.Frozen(ns.GetHost(), true)
	ns.CancelDialog(obj43)
	ns.FrameTimer(30, SaveGameFade)
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
