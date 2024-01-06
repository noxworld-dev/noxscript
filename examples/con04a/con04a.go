package con04a

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	ivar4  int
	ivar5  int
	gvar6  int
	fvar7  float32
	obj8   ns.ObjectID
	obj9   ns.ObjectID
	obj10  ns.ObjectID
	obj11  ns.ObjectID
	obj12  ns.ObjectID
	obj13  ns.ObjectID
	wp14   ns.WaypointID
	wp15   ns.WaypointID
	wp16   ns.WaypointID
	wp17   ns.WaypointID
	gvar18 int
	gvar19 int
	gvar20 int
	gvar21 ns.ObjectGroupID
	gvar22 ns.ObjectGroupID
	obj23  ns.ObjectID
	wp24   ns.WaypointID
)

func init() {
	ivar4 = 50
	ivar5 = 20
	gvar6 = 1
	fvar7 = 5
	gvar18 = 0
	gvar19 = 1
	gvar20 = gvar18
}
func CloseBlockTrap01() {
	ns.ObjectOff(obj12)
	ns.Move(obj8, wp14)
	ns.Move(obj9, wp15)
	ns.FrameTimerWithArg(ivar5, 1, BlockTrapBoom)
	ns.FrameTimerWithArg(ivar4, 1, ResetBlockTrap)
}
func CloseBlockTrap02() {
	ns.ObjectOff(obj13)
	ns.Move(obj10, wp16)
	ns.Move(obj11, wp17)
	ns.FrameTimerWithArg(ivar5, 2, BlockTrapBoom)
	ns.FrameTimerWithArg(ivar4, 2, ResetBlockTrap)
}
func BlockTrapLoop() {
	var v0 int
	v0 = gvar6
	if v0 == 1 {
		goto LABEL1
	}
	if v0 == 2 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	CloseBlockTrap01()
	gvar6 = 2
	goto LABEL3
LABEL2:
	CloseBlockTrap02()
	gvar6 = 1
	goto LABEL3
LABEL3:
	ns.FrameTimer(30, BlockTrapLoop)
}
func InitializeBlockTrap() {
	obj8 = ns.Object("SpikeBlock01")
	obj9 = ns.Object("SpikeBlock02")
	obj10 = ns.Object("SpikeBlock03")
	obj11 = ns.Object("SpikeBlock04")
	wp14 = ns.Waypoint("BlockWP01")
	wp15 = ns.Waypoint("BlockWP02")
	wp16 = ns.Waypoint("BlockWP03")
	wp17 = ns.Waypoint("BlockWP04")
	obj12 = ns.Object("BlockTrap01Trigger")
	obj13 = ns.Object("BlockTrap02Trigger")
	BlockTrapLoop()
}
func BlockTrapBoom(a1 int) {
	var v0 int
	v0 = a1
	if v0 == 1 {
		goto LABEL1
	}
	if v0 == 2 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.AudioEvent(ns.HammerMissing, wp14)
	ns.Effect(ns.JIGGLE, ns.GetWaypointX(wp14), ns.GetWaypointY(wp14), fvar7, 0)
	goto LABEL3
LABEL2:
	ns.AudioEvent(ns.HammerMissing, wp16)
	ns.Effect(ns.JIGGLE, ns.GetWaypointX(wp16), ns.GetWaypointY(wp16), fvar7, 0)
	goto LABEL3
LABEL3:
	return
}
func ResetBlockTrap(a1 int) {
	var v0 int
	v0 = a1
	if v0 == 1 {
		goto LABEL1
	}
	if v0 == 2 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.ObjectOn(obj12)
	ns.AudioEvent(ns.TriggerReleased, wp14)
	goto LABEL3
LABEL2:
	ns.ObjectOn(obj13)
	ns.AudioEvent(ns.TriggerReleased, wp16)
	goto LABEL3
LABEL3:
	return
}
func PlayerDeath() {
	ns.DeathScreen(4)
}
func CaptainDialogStart() {
	var v0 int
	v0 = gvar20
	if v0 == gvar18 {
		goto LABEL1
	}
	if v0 == gvar19 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.TellStory(ns.SwordsmanHurt, "Con04a:CaptainGreet")
	goto LABEL3
LABEL2:
	ns.TellStory(ns.SwordsmanHurt, "Con04a:CaptainIdle")
	goto LABEL3
LABEL3:
	return
}
func CaptainDialogEnd() {
	var v0 int
	v0 = gvar20
	if v0 == gvar18 {
		goto LABEL1
	}
	if v0 == gvar19 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.JournalEntry(ns.GetHost(), "Chapter4SearchCrypts", 2)
	ns.PrintToAll("Con01a:NewJournalEntry")
	ns.AwardSpell(ns.GetHost(), ns.SPELL_GLYPH)
	ns.PrintToAll("GeneralPrint:Bomber")
	ns.JournalEntry(ns.GetHost(), "BomberHint", 8)
	gvar20 = gvar19
	goto LABEL3
LABEL2:
	goto LABEL3
LABEL3:
	return
}
func StartCaptainConversation() {
	ns.SetDialog(obj23, ns.NORMAL, CaptainDialogStart, CaptainDialogEnd)
	ns.StartDialog(obj23, ns.GetHost())
}
func PlayOutdoorMusic() {
	ns.Music(22, 100)
}
func MapInitialize() {
	gvar21 = ns.ObjectGroup("Secret2Triggers")
	gvar22 = ns.ObjectGroup("Fish")
	obj23 = ns.Object("Airship_Captain")
	wp24 = ns.Waypoint("SecretAudioOrigin")
	ns.StoryPic(obj23, "AirshipCaptainPic")
	ns.SetOwner(ns.GetHost(), obj23)
	ns.GroupWander(gvar22)
	InitializeBlockTrap()
	PlayOutdoorMusic()
	ns.StartupScreen(4)
	ns.FrameTimer(5, StartCaptainConversation)
}
func OpenSecretPassageWalls() {
	ns.WallOpen(ns.Wall(129, 137))
}
func ExitMessage() {
	ns.PrintToAll("War04a:ExitMessage")
}
func PlayUndergroundMusic() {
	ns.Music(20, 100)
}
func SecretSFX() {
	ns.MoveWaypoint(wp24, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretFound, wp24)
}
func FoundSecret1() {
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 25)
	SecretSFX()
}
func FoundSecret2() {
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 100)
	ns.ObjectGroupOff(gvar21)
	SecretSFX()
}
func FoundSecret3() {
	ns.ObjectOff(ns.GetTrigger())
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 100)
	SecretSFX()
}
func OnEvent(typ string) {
	switch typ {
	case "PlayerDeath":
		PlayerDeath()
	case "MapInitialize":
		MapInitialize()
	}
}
