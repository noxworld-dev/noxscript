package con11a

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	gvar4  int
	gvar5  int
	gvar6  int
	gvar7  int
	gvar8  int
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
	gvar31 ns.ObjectGroupID
	wp32   ns.WaypointID
	wp33   ns.WaypointID
	wp34   ns.WaypointID
	wp35   ns.WaypointID
	wp36   ns.WaypointID
	wp37   ns.WaypointID
	wp38   ns.WaypointID
	wp39   ns.WaypointID
	wp40   ns.WaypointID
	wp41   ns.WaypointID
	wp42   ns.WaypointID
	wp43   ns.WaypointID
	wp44   ns.WaypointID
	gvar45 ns.WaypointGroupID
	gvar46 ns.WaypointGroupID
	gvar47 ns.WaypointGroupID
	gvar48 int
	flag49 bool
	obj50  ns.ObjectID
	obj51  ns.ObjectID
	obj52  ns.ObjectID
	obj53  ns.ObjectID
	gvar54 ns.ObjectGroupID
	gvar55 ns.ObjectGroupID
	wp56   ns.WaypointID
	wp57   ns.WaypointID
	wp58   ns.WaypointID
	wp59   ns.WaypointID
	wp60   ns.WaypointID
	wp61   ns.WaypointID
	wp62   ns.WaypointID
	wp63   ns.WaypointID
	wp64   ns.WaypointID
	wp65   ns.WaypointID
	wp66   ns.WaypointID
	wp67   ns.WaypointID
	wp68   ns.WaypointID
	wp69   ns.WaypointID
	wp70   ns.WaypointID
	wp71   ns.WaypointID
	wp72   ns.WaypointID
	wp73   ns.WaypointID
	wp74   ns.WaypointID
	wp75   ns.WaypointID
	wp76   ns.WaypointID
	wp77   ns.WaypointID
	wp78   ns.WaypointID
	wp79   ns.WaypointID
	wp80   ns.WaypointID
	wp81   ns.WaypointID
	wp82   ns.WaypointID
	wp83   ns.WaypointID
	wp84   ns.WaypointID
	wp85   ns.WaypointID
	wp86   ns.WaypointID
	wp87   ns.WaypointID
	wp88   ns.WaypointID
	wp89   ns.WaypointID
	wp90   ns.WaypointID
	wp91   ns.WaypointID
	wp92   ns.WaypointID
	wp93   ns.WaypointID
	wp94   ns.WaypointID
	wp95   ns.WaypointID
	wp96   ns.WaypointID
	wp97   ns.WaypointID
	flag98 bool
	gvar99 ns.WaypointGroupID
)

func init() {
	gvar4 = 0
	gvar5 = 1
	gvar6 = 2
	gvar7 = 3
	gvar8 = 4
	gvar48 = gvar4
	flag49 = false
	flag98 = false
}
func SwampGuards() {
	ns.ObjectGroupOn(gvar31)
	ns.AggressionLevel(obj10, 0.83)
	ns.AggressionLevel(obj11, 0.83)
	ns.Move(obj12, wp35)
	ns.Move(obj13, wp36)
	ns.Move(obj14, wp37)
	ns.Frozen(ns.GetHost(), false)
	ns.WideScreen(false)
}
func HecTalk() {
	ns.WayPointGroupOn(gvar47)
	ns.SetDialog(obj9, ns.NORMAL, HecubahDialogStart, HecubahDialogEnd)
	ns.StartDialog(obj9, ns.GetHost())
}
func HecDie() {
	ns.Blind()
	ns.Frozen(obj9, true)
	ns.Frozen(ns.GetHost(), true)
	ns.Music(0, 100)
	ns.MoveWaypoint(wp41, ns.GetObjectX(obj9), ns.GetObjectY(obj9))
	ns.AudioEvent(ns.HecubahDieFrame0A, wp41)
	ns.FrameTimer(45, HecDie2)
}
func HecDie2() {
	ns.Frozen(obj9, false)
	ns.Frozen(ns.GetHost(), false)
	ns.MoveObject(obj9, ns.GetWaypointX(wp39), ns.GetWaypointY(wp39))
	ns.MoveObject(ns.GetHost(), ns.GetWaypointX(wp40), ns.GetWaypointY(wp40))
	ns.WideScreen(true)
	ns.FrameTimer(1, HecDie3)
}
func HecDie3() {
	ns.UnBlind()
	ns.LookAtObject(ns.GetHost(), obj9)
	ns.Frozen(ns.GetHost(), true)
	ns.Frozen(obj9, true)
	ns.FrameTimer(30, HecDie4)
}
func HecDie4() {
	ns.Frozen(obj9, false)
	ns.FrameTimer(67, HecubahDie1)
	ns.FrameTimer(131, HecubahDie2)
	ns.FrameTimer(171, HecubahDie3)
	ns.FrameTimer(281, HecubahDie4)
	ns.FrameTimer(440, TheEnd1)
}
func TheEnd1() {
	ns.EndGame(2)
}
func HecubahDie1() {
	ns.AudioEvent(ns.HecubahDieFrame98, wp39)
}
func HecubahDie2() {
	ns.AudioEvent(ns.HecubahDieFrame194, wp39)
}
func HecubahDie3() {
	ns.AudioEvent(ns.HecubahDieFrame283, wp39)
}
func HecubahDie4() {
	ns.AudioEvent(ns.HecubahDieFrame439, wp39)
}
func HecubahDialogStart() {
	var v0 int
	v0 = gvar48
	if v0 == gvar4 {
		goto LABEL1
	}
	if v0 == gvar5 {
		goto LABEL2
	}
	if v0 == gvar6 {
		goto LABEL3
	}
	if v0 == gvar7 {
		goto LABEL4
	}
	if v0 == gvar8 {
		goto LABEL5
	}
	goto LABEL6
LABEL1:
	ns.LookAtObject(obj9, ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj9)
	ns.TellStory(ns.SwordsmanHurt, "Con10B.scr:HecubahDialog2")
	goto LABEL6
LABEL2:
	ns.LookAtObject(obj9, ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj9)
	ns.TellStory(ns.SwordsmanHurt, "Con10B.scr:HecubahLine1")
	goto LABEL6
LABEL3:
	ns.LookAtObject(obj9, ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj9)
	ns.TellStory(ns.SwordsmanHurt, "Con10B.scr:HecubahLine3")
	goto LABEL6
LABEL4:
	ns.LookAtObject(ns.GetHost(), obj9)
	ns.TellStory(ns.SwordsmanHurt, "Con10B.scr:HecubahLine16")
	goto LABEL6
LABEL5:
	ns.TellStory(ns.SwordsmanHurt, "Con10B.scr:HecubahDefeat6")
	goto LABEL6
LABEL6:
	return
}
func HecubahDialogEnd() {
	var v0 int
	v0 = gvar48
	if v0 == gvar4 {
		goto LABEL1
	}
	if v0 == gvar5 {
		goto LABEL2
	}
	if v0 == gvar6 {
		goto LABEL3
	}
	if v0 == gvar7 {
		goto LABEL4
	}
	if v0 == gvar8 {
		goto LABEL5
	}
	goto LABEL6
LABEL1:
	gvar48 = gvar5
	ns.SetDialog(obj9, ns.NORMAL, HecubahDialogStart, HecubahDialogEnd)
	ns.StartDialog(obj9, ns.GetHost())
	goto LABEL6
LABEL2:
	gvar48 = gvar6
	ns.CancelDialog(obj9)
	ns.Enchant(obj9, ns.ENCHANT_INVISIBLE, 3)
	ns.Effect(ns.BLUE_SPARKS, ns.GetObjectX(obj9), ns.GetObjectY(obj9), 0, 0)
	ns.Effect(ns.SMOKE_BLAST, ns.GetObjectX(obj9), ns.GetObjectY(obj9), 0, 0)
	ns.MoveObject(obj9, ns.GetWaypointX(wp34), ns.GetWaypointY(wp34))
	SwampGuards()
	goto LABEL6
LABEL3:
	gvar48 = gvar7
	ns.CancelDialog(obj9)
	ns.Frozen(ns.GetHost(), false)
	ns.WideScreen(false)
	ns.AggressionLevel(obj9, 0.83)
	ns.CastSpellObjectObject(ns.SPELL_BLINK, obj9, obj9)
	goto LABEL6
LABEL4:
	ns.FrameTimer(310, TheEnd1)
	goto LABEL6
LABEL5:
	goto LABEL6
LABEL6:
	return
}
func GreetPlayer1() {
	ns.ObjectOff(ns.GetTrigger())
	ns.WideScreen(true)
	ns.Frozen(ns.GetHost(), true)
	ns.Move(obj9, wp32)
}
func GreetPlayer2() {
	if !ns.IsCaller(obj9) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.StartDialog(obj9, ns.GetHost())
LABEL1:
	return
}
func WarnPlayer() {
	if !(ns.IsCaller(ns.GetHost()) && flag49 == false) {
		goto LABEL1
	}
	flag49 = true
	ns.WayPointGroupOff(gvar45)
	ns.Frozen(ns.GetHost(), true)
	ns.WideScreen(true)
	ns.Move(obj9, wp38)
	ns.FrameTimer(45, HecTalk)
LABEL1:
	return
}
func BlinkLimit() {
	ns.WayPointGroupOff(gvar45)
	ns.WayPointGroupOn(gvar46)
}
func HecBlink() {
	if !(ns.IsCaller(obj15) || ns.IsCaller(obj16) || ns.IsCaller(obj17) || ns.IsCaller(obj18) || ns.IsCaller(obj19) || ns.IsCaller(obj20) || ns.IsCaller(obj21) || ns.IsCaller(obj22) || ns.IsCaller(obj23) || ns.IsCaller(obj24) || ns.IsCaller(obj25) || ns.IsCaller(obj26) || ns.IsCaller(obj27) || ns.IsCaller(obj28) || ns.IsCaller(obj29) || ns.IsCaller(obj30)) {
		goto LABEL1
	}
	ns.CastSpellObjectObject(ns.SPELL_BLINK, obj9, obj9)
LABEL1:
	return
}
func MoveBlocksOut() {
	ns.Move(obj15, wp57)
	ns.Move(obj16, wp58)
	ns.Move(obj17, wp59)
	ns.Move(obj18, wp60)
	ns.Move(obj19, wp61)
	ns.Move(obj20, wp62)
	ns.Move(obj21, wp63)
	ns.Move(obj22, wp64)
	ns.Move(obj23, wp65)
	ns.Move(obj24, wp66)
	ns.Move(obj25, wp67)
	ns.Move(obj26, wp68)
	ns.FrameTimer(120, MoveBlocksHome)
}
func MoveBlocksHome() {
	if !flag98 {
		goto LABEL1
	}
	ns.Move(obj15, wp69)
	ns.Move(obj16, wp70)
	ns.Move(obj17, wp71)
	ns.Move(obj18, wp72)
	ns.Move(obj19, wp73)
	ns.Move(obj20, wp74)
	ns.Move(obj21, wp75)
	ns.Move(obj22, wp76)
	ns.Move(obj23, wp77)
	ns.Move(obj24, wp78)
	ns.Move(obj25, wp79)
	ns.Move(obj26, wp80)
	ns.FrameTimer(120, MoveBlocksOut)
	goto LABEL2
LABEL1:
	return
LABEL2:
	return
}
func ReopenSet1() {
	ns.ObjectGroupOn(gvar54)
	ns.Move(obj27, wp89)
	ns.Move(obj50, wp90)
	ns.Move(obj28, wp91)
	ns.Move(obj51, wp92)
}
func ReopenSet2() {
	ns.ObjectGroupOn(gvar55)
	ns.Move(obj29, wp93)
	ns.Move(obj52, wp94)
	ns.Move(obj30, wp95)
	ns.Move(obj53, wp96)
}
func CrashAndRumble() {
	ns.MoveWaypoint(wp56, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.HammerMissing, wp56)
	ns.Effect(ns.JIGGLE, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 10, 0)
}
func Secret01Declare() {
	ns.ObjectOff(ns.GetTrigger())
	ns.MoveWaypoint(wp97, ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger()))
	ns.AudioEvent(ns.SecretFound, wp97)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 1000)
}
func CloseSet1() {
	ns.ObjectGroupOff(gvar54)
	ns.Move(obj27, wp81)
	ns.Move(obj50, wp82)
	ns.Move(obj28, wp83)
	ns.Move(obj51, wp84)
	ns.FrameTimer(16, CrashAndRumble)
	ns.FrameTimer(75, ReopenSet1)
}
func CloseSet2() {
	ns.ObjectGroupOff(gvar55)
	ns.Move(obj29, wp85)
	ns.Move(obj52, wp86)
	ns.Move(obj30, wp87)
	ns.Move(obj53, wp88)
	ns.FrameTimer(16, CrashAndRumble)
	ns.FrameTimer(75, ReopenSet2)
}
func NearBlocks() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	flag98 = true
LABEL1:
	MoveBlocksOut()
}
func FarFromBlocks() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	flag98 = false
LABEL1:
	return
}
func GoMedieval() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.AggressionLevel(ns.GetTrigger(), 0.83)
LABEL1:
	return
}
func PlayerDeath() {
	ns.DeathScreen(11)
}
func MapInitialize() {
	obj9 = ns.Object("Hecubah")
	obj10 = ns.Object("Guard1")
	obj11 = ns.Object("Guard2")
	obj12 = ns.Object("GuardA")
	obj13 = ns.Object("GuardB")
	obj14 = ns.Object("GuardC")
	obj15 = ns.Object("BlockA1")
	obj16 = ns.Object("BlockA2")
	obj17 = ns.Object("BlockA3")
	obj18 = ns.Object("BlockB1")
	obj19 = ns.Object("BlockB2")
	obj20 = ns.Object("BlockB3")
	obj21 = ns.Object("BlockC1")
	obj22 = ns.Object("BlockC2")
	obj23 = ns.Object("BlockC3")
	obj24 = ns.Object("BlockD1")
	obj25 = ns.Object("BlockD2")
	obj26 = ns.Object("BlockD3")
	obj27 = ns.Object("DoorBlock1")
	obj50 = ns.Object("DoorBlock2")
	obj28 = ns.Object("DoorBlock3")
	obj51 = ns.Object("DoorBlock4")
	obj29 = ns.Object("DoorBlock5")
	obj52 = ns.Object("DoorBlock6")
	obj30 = ns.Object("DoorBlock7")
	obj53 = ns.Object("DoorBlock8")
	gvar31 = ns.ObjectGroup("GuardGroup1")
	gvar54 = ns.ObjectGroup("TriggerSet1")
	gvar55 = ns.ObjectGroup("TriggerSet2")
	wp32 = ns.Waypoint("HecWP1")
	wp41 = ns.Waypoint("HecSounds")
	wp33 = ns.Waypoint("HecLeaveWP")
	wp34 = ns.Waypoint("HecTeleport")
	wp35 = ns.Waypoint("GuardAWP")
	wp36 = ns.Waypoint("GuardBWP")
	wp37 = ns.Waypoint("GuardCWP")
	wp38 = ns.Waypoint("HecWarning")
	wp39 = ns.Waypoint("FinalHecubahWP")
	wp40 = ns.Waypoint("FinalPlayerWP")
	wp56 = ns.Waypoint("SoundSource")
	wp97 = ns.Waypoint("PlayerSounds")
	wp57 = ns.Waypoint("A1Out")
	wp58 = ns.Waypoint("A2Out")
	wp59 = ns.Waypoint("A3Out")
	wp60 = ns.Waypoint("B1Out")
	wp61 = ns.Waypoint("B2Out")
	wp62 = ns.Waypoint("B3Out")
	wp63 = ns.Waypoint("C1Out")
	wp64 = ns.Waypoint("C2Out")
	wp65 = ns.Waypoint("C3Out")
	wp66 = ns.Waypoint("D1Out")
	wp67 = ns.Waypoint("D2Out")
	wp68 = ns.Waypoint("D3Out")
	wp69 = ns.Waypoint("A1Home")
	wp70 = ns.Waypoint("A2Home")
	wp71 = ns.Waypoint("A3Home")
	wp72 = ns.Waypoint("B1Home")
	wp73 = ns.Waypoint("B2Home")
	wp74 = ns.Waypoint("B3Home")
	wp75 = ns.Waypoint("C1Home")
	wp76 = ns.Waypoint("C2Home")
	wp77 = ns.Waypoint("C3Home")
	wp78 = ns.Waypoint("D1Home")
	wp79 = ns.Waypoint("D2Home")
	wp80 = ns.Waypoint("D3Home")
	wp81 = ns.Waypoint("Close1")
	wp82 = ns.Waypoint("Close2")
	wp83 = ns.Waypoint("Close3")
	wp84 = ns.Waypoint("Close4")
	wp85 = ns.Waypoint("Close5")
	wp86 = ns.Waypoint("Close6")
	wp87 = ns.Waypoint("Close7")
	wp88 = ns.Waypoint("Close8")
	wp89 = ns.Waypoint("Open1")
	wp90 = ns.Waypoint("Open2")
	wp91 = ns.Waypoint("Open3")
	wp92 = ns.Waypoint("Open4")
	wp93 = ns.Waypoint("Open5")
	wp94 = ns.Waypoint("Open6")
	wp95 = ns.Waypoint("Open7")
	wp96 = ns.Waypoint("Open8")
	wp42 = ns.Waypoint("EditWP1")
	wp43 = ns.Waypoint("EditWP2")
	wp44 = ns.Waypoint("EditWP3")
	gvar45 = ns.WaypointGroup("AllSearchWPs")
	gvar99 = ns.WaypointGroup("SwampSearchWPs")
	gvar46 = ns.WaypointGroup("LessonSearchWPs")
	gvar47 = ns.WaypointGroup("DungeonSearchWPs")
	ns.StoryPic(obj9, "HecubahPic")
	ns.SetDialog(obj9, ns.NEXT, HecubahDialogStart, HecubahDialogEnd)
	ns.WayPointGroupOff(gvar45)
	ns.WayPointGroupOn(gvar99)
}
func OnEvent(typ string) {
	switch typ {
	case "PlayerDeath":
		PlayerDeath()
	case "MapInitialize":
		MapInitialize()
	}
}
