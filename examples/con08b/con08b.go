package con08b

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	obj4   ns.ObjectID
	obj5   ns.ObjectID
	obj6   ns.ObjectID
	obj7   ns.ObjectID
	wp8    ns.WaypointID
	wp9    ns.WaypointID
	gvar10 ns.ObjectGroupID
	wp11   ns.WaypointID
	wp12   ns.WaypointID
	wp13   ns.WaypointID
	wp14   ns.WaypointID
	gvar15 int
	gvar16 int
	gvar17 int
	gvar18 int
	gvar19 int
	flag20 bool
	flag21 bool
)

func init() {
	gvar15 = 0
	gvar16 = 1
	gvar17 = 2
	gvar18 = 3
	gvar19 = gvar15
	flag20 = false
	flag21 = false
}
func PlayerDeath() {
	ns.DeathScreen(8)
}
func GoToIxSEG2() {
	ns.Frozen(ns.GetHost(), false)
	ns.Frozen(obj4, false)
	ns.Frozen(obj5, false)
	ns.Frozen(obj6, false)
	ns.MoveObject(ns.GetHost(), ns.GetWaypointX(wp8), ns.GetWaypointY(wp8))
}
func GoToIx() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.Frozen(ns.GetHost(), true)
	ns.Frozen(obj4, true)
	ns.Frozen(obj5, true)
	ns.Frozen(obj6, true)
	ns.Blind()
	ns.FrameTimer(60, GoToIxSEG2)
LABEL1:
	return
}
func GoToTempleSEG2() {
	ns.Frozen(ns.GetHost(), false)
	ns.MoveObject(ns.GetHost(), ns.GetWaypointX(wp9), ns.GetWaypointY(wp9))
}
func GoToTemple() {
	ns.Frozen(ns.GetHost(), true)
	ns.CancelDialog(obj7)
	ns.Blind()
	ns.FrameTimer(60, GoToTempleSEG2)
}
func OwnObjects() {
	ns.SetOwner(ns.GetHost(), obj7)
}
func MapInitialize() {
	obj4 = ns.Object("Bear01")
	obj5 = ns.Object("Bear02")
	obj6 = ns.Object("Bear03")
	obj7 = ns.Object("IxPriest")
	gvar10 = ns.ObjectGroup("Fish")
	wp8 = ns.Waypoint("IxExitWP")
	wp11 = ns.Waypoint("TempleExtWP")
	wp12 = ns.Waypoint("TempleIntWP")
	wp9 = ns.Waypoint("TempleExitWP")
	wp13 = ns.Waypoint("PriestWP")
	wp14 = ns.Waypoint("PriestLookWP")
	ns.GroupWander(gvar10)
	ns.FrameTimer(1, OwnObjects)
}
func MapEntry() {
	ns.UnBlind()
	ns.Music(22, 100)
}
func EnterTempleSEG2() {
	ns.Frozen(ns.GetHost(), false)
	ns.Frozen(obj4, false)
	ns.Frozen(obj5, false)
	ns.Frozen(obj6, false)
	ns.MoveObject(ns.GetHost(), ns.GetWaypointX(wp12), ns.GetWaypointY(wp12))
	ns.UnBlind()
}
func EnterTemple() {
	ns.Frozen(ns.GetHost(), true)
	ns.Frozen(obj4, true)
	ns.Frozen(obj5, true)
	ns.Frozen(obj6, true)
	ns.Blind()
	ns.FrameTimer(60, EnterTempleSEG2)
}
func LeaveTempleSEG2() {
	ns.Frozen(ns.GetHost(), false)
	ns.MoveObject(ns.GetHost(), ns.GetWaypointX(wp11), ns.GetWaypointY(wp11))
	ns.UnBlind()
}
func LeaveTemple() {
	ns.Frozen(ns.GetHost(), true)
	ns.Blind()
	ns.FrameTimer(60, LeaveTempleSEG2)
}
func StayAway() {
	r1 := ns.IsAttackedBy(ns.GetCaller(), ns.GetHost())
	if !r1 {
		goto LABEL1
	}
	ns.GoBackHome(ns.GetCaller())
LABEL1:
	return
}
func NoMonsters() {
	r1 := ns.IsAttackedBy(ns.GetCaller(), ns.GetHost())
	if !r1 {
		goto LABEL1
	}
	ns.GoBackHome(ns.GetCaller())
LABEL1:
	return
}
func BearAttack() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.CreatureFollow(obj5, obj4)
	ns.CreatureFollow(obj6, obj4)
LABEL1:
	return
}
func BearIdle() {
	ns.GoBackHome(ns.GetTrigger())
}
func PriestSetPieceSEG2() {
	ns.CreatureFollow(obj7, ns.GetHost())
	ns.FrameTimer(75, PriestSetPieceSEG3)
}
func PriestSetPieceSEG3() {
	ns.StoryPic(obj7, "IxPriestPic")
	ns.SetDialog(obj7, ns.NEXT, PriestDialogStart, PriestDialogEnd)
	ns.StartDialog(obj7, ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj7)
}
func PriestDialogStart() {
	var v0 int
	v0 = gvar19
	if v0 == gvar15 {
		goto LABEL1
	}
	if v0 == gvar16 {
		goto LABEL2
	}
	if v0 == gvar17 {
		goto LABEL3
	}
	if v0 == gvar18 {
		goto LABEL4
	}
	goto LABEL5
LABEL1:
	ns.TellStory(ns.SwordsmanHurt, "Con08a:PriestGreet")
	goto LABEL5
LABEL2:
	ns.TellStory(ns.SwordsmanHurt, "Con08a:PriestProd")
	goto LABEL5
LABEL3:
	ns.TellStory(ns.SwordsmanHurt, "Con08a:PriestProd2")
	goto LABEL5
LABEL4:
	ns.TellStory(ns.SwordsmanHurt, "Con08a:PriestProd3")
	goto LABEL5
LABEL5:
	return
}
func PriestDialogEnd() {
	var v0 int
	v0 = gvar19
	if v0 == gvar15 {
		goto LABEL1
	}
	if v0 == gvar16 {
		goto LABEL2
	}
	if v0 == gvar17 {
		goto LABEL3
	}
	if v0 == gvar18 {
		goto LABEL4
	}
	goto LABEL5
LABEL1:
	gvar19 = gvar16
	PriestDialogStart()
	goto LABEL5
LABEL2:
	gvar19 = gvar17
	PriestDialogStart()
	goto LABEL5
LABEL3:
	gvar19 = gvar18
	ns.SetDialog(obj7, ns.NORMAL, PriestDialogStart, PriestDialogEnd)
	PriestDialogStart()
	goto LABEL5
LABEL4:
	if flag21 {
		goto LABEL6
	}
	flag21 = true
	ns.JournalEdit(ns.GetHost(), "Chapter8IxPriest", 4)
	ns.PrintToAll("Con02a:ObjectiveComplete")
	ns.JournalEntry(ns.GetHost(), "Chapter8Wierdling", 2)
	ns.PrintToAll("Con01a:NewJournalEntry")
	ns.WideScreen(false)
	ReleasePlayer()
LABEL6:
	ns.CreatureGuard(obj7, ns.GetWaypointX(wp13), ns.GetWaypointY(wp13), ns.GetWaypointX(wp14), ns.GetWaypointY(wp14), 0)
	goto LABEL5
LABEL5:
	return
}
func ReleasePlayer() {
	ns.MusicPopEvent()
	ns.Frozen(ns.GetHost(), false)
}
func StartPriestSetPiece() {
	if flag20 {
		goto LABEL1
	}
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.MusicPushEvent()
	ns.Music(10, 100)
	flag20 = true
	ns.AudioEvent(ns.BigGong, wp12)
	ns.Frozen(ns.GetHost(), true)
	ns.WideScreen(true)
	ns.FrameTimer(60, PriestSetPieceSEG2)
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
