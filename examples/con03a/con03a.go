package con03a

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	fvar4  float32
	fvar5  float32
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
	wp24   ns.WaypointID
	wp25   ns.WaypointID
	obj26  ns.ObjectID
	obj27  ns.ObjectID
	wp28   ns.WaypointID
	wp29   ns.WaypointID
	gvar30 int
	gvar31 int
	gvar32 int
	gvar33 int
	gvar34 int
	gvar35 int
	gvar36 int
	gvar37 int
	gvar38 int
	gvar39 int
	gvar40 int
	gvar41 int
	gvar42 int
	gvar43 ns.ObjectGroupID
	gvar44 ns.ObjectGroupID
	gvar45 ns.ObjectGroupID
	obj46  ns.ObjectID
	gvar47 int
	obj48  ns.ObjectID
	obj49  ns.ObjectID
	gvar50 int
	gvar51 int
	gvar52 int
	gvar53 int
	gvar54 int
	wp55   ns.WaypointID
	wp56   ns.WaypointID
	wp57   ns.WaypointID
	wp58   ns.WaypointID
	obj59  ns.ObjectID
	gvar60 ns.ObjectGroupID
	obj61  ns.ObjectID
	gvar62 ns.ObjectGroupID
	flag63 bool
	obj64  ns.ObjectID
	flag65 bool
	flag66 bool
)

func init() {
	fvar4 = 184
	fvar5 = -772
	gvar30 = 1
	gvar31 = 1
	gvar32 = 0
	gvar33 = 0
	gvar34 = 1
	gvar35 = 2
	gvar36 = gvar33
	gvar37 = 0
	gvar38 = 1
	gvar39 = gvar37
	gvar40 = 0
	gvar41 = 1
	gvar42 = gvar40
	gvar50 = 0
	gvar51 = 1
	gvar52 = 2
	gvar53 = 3
	gvar54 = gvar50
	flag63 = true
	flag65 = false
	flag66 = false
}
func DunMirGuard1TalkStart() {
	ns.TellStory(ns.SwordsmanHurt, "Con03A.scr:DunMirGuard1")
}
func DunMirGuard1TalkEnd() {
}
func DunMirGuard2TalkStart() {
	ns.TellStory(ns.SwordsmanHurt, "Con03A.scr:DunMirGuard2")
}
func DunMirGuard2TalkEnd() {
}
func MineGuardDialogStart() {
	var v0 int
	v0 = gvar31
	if v0 == 1 {
		goto LABEL1
	}
	if v0 == 2 {
		goto LABEL2
	}
	if v0 == 4 {
		goto LABEL3
	}
	goto LABEL4
LABEL1:
	ns.TellStory(ns.DemonRecognize, "Con03A.scr:MineGuardA")
	goto LABEL5
LABEL2:
	ns.TellStory(ns.DemonRecognize, "Con03A.scr:MineGuardB")
	goto LABEL5
LABEL3:
	ns.TellStory(ns.DemonRecognize, "Con03A.scr:MineGuardC")
	goto LABEL5
LABEL4:
	goto LABEL5
LABEL5:
	return
}
func MineGuardDialogEnd() {
	var v0 int
	v0 = gvar31
	if v0 == 1 {
		goto LABEL1
	}
	goto LABEL2
LABEL1:
	gvar31 = 2
	goto LABEL3
LABEL2:
	goto LABEL3
LABEL3:
	return
}
func GalavaGuard1TalkStart() {
	ns.TellStory(ns.SwordsmanHurt, "Con03A.scr:GalavaGuard1")
}
func GalavaGuard1TalkEnd() {
}
func GalavaGuard2TalkStart() {
	ns.TellStory(ns.SwordsmanHurt, "Con03A.scr:GalavaGuard2")
}
func GalavaGuard2TalkEnd() {
}
func IxGuard1TalkStart() {
	ns.TellStory(ns.SwordsmanHurt, "Con03A.scr:IxGuard1")
}
func IxGuard1TalkEnd() {
}
func IxGuard2TalkStart() {
	ns.TellStory(ns.SwordsmanHurt, "Con03A.scr:IxGuard2")
}
func IxGuard2TalkEnd() {
}
func CaptainDialogStart() {
	var v0 int
	ns.LookAtObject(obj21, ns.GetHost())
	v0 = gvar30
	if v0 == 1 {
		goto LABEL1
	}
	if v0 == 2 {
		goto LABEL2
	}
	if v0 == 3 {
		goto LABEL3
	}
	goto LABEL4
LABEL1:
	ns.TellStory(ns.SwordsmanHurt, "Con03A.scr:JandorA")
	goto LABEL5
LABEL2:
	ns.TellStory(ns.SwordsmanHurt, "Con03A.scr:JandorB")
	goto LABEL5
LABEL3:
	ns.CancelDialog(obj21)
	ChapterEnd()
	goto LABEL5
LABEL4:
	goto LABEL5
LABEL5:
	return
}
func CaptainDialogEnd() {
	var v0 int
	v0 = gvar30
	if v0 == 1 {
		goto LABEL1
	}
	if v0 == 2 {
		goto LABEL2
	}
	if v0 == 3 {
		goto LABEL3
	}
	goto LABEL4
LABEL1:
	ns.JournalEdit(ns.GetHost(), "Con03FindAirshipCap", 4)
	ns.SetDialog(obj21, ns.NORMAL, CaptainDialogStart, CaptainDialogEnd)
	ns.PrintToAll("Con03C.scr:TaskComplete")
	gvar30 = 2
	CaptainDialogStart()
	goto LABEL5
LABEL2:
	gvar30 = 3
	goto LABEL5
LABEL3:
	goto LABEL5
LABEL4:
	goto LABEL5
LABEL5:
	return
}
func InitializeRobbery() {
	obj61 = ns.Object("Robber")
	gvar62 = ns.ObjectGroup("RobberyTriggers")
}
func HermitTalkStart() {
	var v0 int
	ns.LookAtObject(obj46, ns.GetHost())
	v0 = gvar54
	if v0 == gvar50 {
		goto LABEL1
	}
	if v0 == gvar51 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	if !ns.HasItem(ns.GetHost(), obj48) {
		goto LABEL4
	}
	ns.TellStory(ns.SwordsmanHurt, "Con03A.scr:HermitHappy")
	gvar54 = gvar52
	goto LABEL5
LABEL4:
	ns.TellStory(ns.SwordsmanHurt, "Con03A.scr:HermitMeet01")
LABEL5:
	goto LABEL3
LABEL2:
	if !ns.HasItem(ns.GetHost(), obj48) {
		goto LABEL6
	}
	ns.TellStory(ns.SwordsmanHurt, "Con03A.scr:HermitHappy")
	gvar54 = gvar52
	goto LABEL7
LABEL6:
	ns.TellStory(ns.SwordsmanHurt, "Con03A.scr:HermitMeet02")
LABEL7:
	goto LABEL3
LABEL3:
	return
}
func HermitTalkEnd() {
	var v0 int
	v0 = gvar54
	if v0 == gvar50 {
		goto LABEL1
	}
	if v0 == gvar52 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	gvar32 = 1
	gvar54 = gvar51
	ns.JournalEntry(ns.GetHost(), "Con03AHermitQuest", 2)
	ns.PrintToAll("Con03C.scr:NewTask")
	goto LABEL3
LABEL2:
	ns.CancelDialog(obj46)
	ns.GiveXp(ns.GetHost(), 500)
	ns.Pickup(ns.GetHost(), obj49)
	ns.Delete(obj48)
	if gvar32 != 1 {
		goto LABEL4
	}
	ns.JournalEdit(ns.GetHost(), "Con03AHermitQuest", 4)
	ns.PrintToAll("Con03C.scr:TaskComplete")
LABEL4:
	goto LABEL3
LABEL3:
	return
}
func OwnDudes() {
	ns.LockDoor(obj10)
	ns.LockDoor(obj11)
	ns.LockDoor(obj12)
	ns.LockDoor(obj13)
}
func ChapterEnd() {
	ns.Blind()
	ns.Delete(obj48)
	ns.Delete(ns.Object("Con02a:BridgeGuardsBoots"))
	ns.FrameTimer(60, ChapterEndB)
}
func ChapterEndB() {
	ns.MoveObject(ns.GetHost(), 650, 3421)
}
func InitializeWaspNest() {
	obj26 = ns.Object("Wasp01")
	obj27 = ns.Object("Wasp02")
	obj64 = ns.Object("WaspNest")
	wp28 = ns.Waypoint("WaspCreateWP01")
	wp29 = ns.Waypoint("WaspCreateWP02")
	ns.Wander(obj26)
	ns.Wander(obj27)
}
func PlayOutdoorMusic() {
	ns.Music(21, 100)
}
func InitializeMusicStuff() {
	obj59 = ns.Object("BanditCampEntranceTrigger")
	gvar60 = ns.ObjectGroup("BanditCampExitTrigger")
}
func MineJournalEntry() {
	ns.JournalEntry(ns.GetHost(), "LocateMineForeman", 2)
}
func MapInitialize() {
	obj6 = ns.Object("Kenneth")
	obj7 = ns.Object("Lance")
	obj8 = ns.Object("Rastur")
	obj9 = ns.Object("Kirik")
	obj10 = ns.Object("GalavaGate1")
	obj11 = ns.Object("GalavaGate2")
	obj17 = ns.Object("IxGate1")
	obj18 = ns.Object("IxGate2")
	obj12 = ns.Object("DunMirDoor1")
	obj13 = ns.Object("DunMirDoor2")
	obj19 = ns.Object("Loproc")
	obj20 = ns.Object("BiffordByzanti")
	obj14 = ns.Object("Millard")
	obj15 = ns.Object("Janero")
	obj16 = ns.Object("Horst")
	obj17 = ns.Object("IxGate1")
	obj18 = ns.Object("IxGate2")
	gvar43 = ns.ObjectGroup("WaspGroup1")
	gvar44 = ns.ObjectGroup("WaspGroup2")
	gvar45 = ns.ObjectGroup("FishGroup")
	obj21 = ns.Object("Captain")
	obj22 = ns.Object("AirshipBasket")
	obj23 = ns.Object("AirshipShadow")
	wp55 = ns.Waypoint("JandorWP")
	wp56 = ns.Waypoint("JandorLookWP")
	wp57 = ns.Waypoint("BasketWP")
	wp58 = ns.Waypoint("BasketShadowWP")
	wp24 = ns.Waypoint("JournalAudioOrigin")
	wp25 = ns.Waypoint("SecretAudioOrigin")
	ns.GroupWander(gvar43)
	ns.GroupWander(gvar44)
	ns.GroupWander(gvar45)
	obj46 = ns.Object("Osborn")
	obj48 = ns.Object("Specs")
	obj49 = ns.Object("BatGuide")
	ns.StoryPic(obj6, "Warrior8Pic")
	ns.SetDialog(obj6, ns.NORMAL, DunMirGuard1TalkStart, DunMirGuard1TalkEnd)
	ns.StoryPic(obj7, "Warrior9Pic")
	ns.SetDialog(obj7, ns.NORMAL, DunMirGuard2TalkStart, DunMirGuard2TalkEnd)
	ns.StoryPic(obj8, "WizardGuard1Pic")
	ns.SetDialog(obj8, ns.NORMAL, GalavaGuard1TalkStart, GalavaGuard1TalkEnd)
	ns.StoryPic(obj9, "WizardGuard2Pic")
	ns.SetDialog(obj9, ns.NORMAL, GalavaGuard2TalkStart, GalavaGuard2TalkEnd)
	ns.StoryPic(obj14, "Townsman2Pic")
	ns.SetDialog(obj14, ns.NORMAL, MineGuardDialogStart, MineGuardDialogEnd)
	ns.StoryPic(obj21, "AirshipCaptainPic")
	ns.SetDialog(obj21, ns.NEXT, CaptainDialogStart, CaptainDialogEnd)
	ns.StoryPic(obj46, "OsbornPic")
	ns.SetDialog(obj46, ns.NORMAL, HermitTalkStart, HermitTalkEnd)
	ns.StoryPic(obj15, "Townsman2Pic")
	ns.SetDialog(obj15, ns.NORMAL, IxGuard1TalkStart, IxGuard1TalkEnd)
	ns.StoryPic(obj16, "IxGuard2Pic")
	ns.SetDialog(obj16, ns.NORMAL, IxGuard2TalkStart, IxGuard2TalkEnd)
	InitializeRobbery()
	ns.FrameTimer(2, OwnDudes)
	InitializeWaspNest()
	InitializeMusicStuff()
	PlayOutdoorMusic()
	ns.StartupScreen(3)
	ns.FrameTimer(15, MineJournalEntry)
}
func SpecsFound() {
	ns.MoveWaypoint(wp24, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.JournalEntryAdd, wp24)
	ns.PrintToAll("thing.db:SpectaclesDescription")
}
func WaspIdle() {
	ns.CreatureGroupIdle(gvar44)
}
func SecretArea1() {
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 20)
}
func MapEntry() {
	ns.UnBlind()
	if ns.GetQuestStatus("Con03B:BeatMine") != 1 {
		goto LABEL1
	}
	gvar31 = 4
	ns.MoveObject(obj21, ns.GetWaypointX(wp55), ns.GetWaypointY(wp55))
	ns.ObjectOn(obj21)
	ns.CreatureGuard(obj21, ns.GetWaypointX(wp55), ns.GetWaypointY(wp55), ns.GetWaypointX(wp56), ns.GetWaypointY(wp56), 10)
	ns.MoveObject(obj22, ns.GetWaypointX(wp57), ns.GetWaypointY(wp57))
	ns.MoveObject(obj23, ns.GetWaypointX(wp58), ns.GetWaypointY(wp58))
	ns.CancelDialog(obj8)
	ns.CancelDialog(obj9)
	ns.CancelDialog(obj15)
	ns.CancelDialog(obj16)
	ns.CancelDialog(obj7)
LABEL1:
	return
}
func CheckForBoots() {
	if !ns.HasItem(ns.GetHost(), ns.Object("Con02a:BridgeGuardsBoots")) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.JournalEntry(ns.GetHost(), "ReturnBridgeGuardsBoots", 2)
LABEL1:
	return
}
func yobutt() {
	gvar31 = 4
	ns.MoveObject(obj21, 1851, 3388)
	ns.ObjectOn(obj21)
	ns.CreatureGuard(obj21, 1851, 3388, 1700, 3600, 10)
	ns.MoveObject(obj22, ns.GetObjectX(obj22)+fvar4, ns.GetObjectY(obj22)+fvar5)
	ns.MoveObject(obj23, ns.GetObjectX(obj23)+fvar4, ns.GetObjectY(obj23)+fvar5)
}
func PlayerDeath() {
	ns.DeathScreen(3)
}
func MonsterGoHome() {
	r1 := ns.IsAttackedBy(ns.GetHost(), ns.GetCaller())
	if !r1 {
		goto LABEL1
	}
	ns.GoBackHome(ns.GetCaller())
LABEL1:
	return
}
func ToggleBanditMusic() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	if !ns.IsObjectOn(obj59) {
		goto LABEL2
	}
	ns.ObjectOff(obj59)
	ns.ObjectGroupOn(gvar60)
	ns.Music(5, 100)
	goto LABEL1
LABEL2:
	ns.ObjectOn(obj59)
	ns.ObjectGroupOff(gvar60)
	ns.Music(21, 100)
LABEL1:
	return
}
func RobberRecognize() {
	if !(ns.IsCaller(ns.GetHost()) && flag63) {
		goto LABEL1
	}
	flag63 = false
	ns.Attack(ns.GetTrigger(), ns.GetHost())
	ns.Chat(ns.GetTrigger(), "War01A.scr:Bully1")
LABEL1:
	return
}
func RobberDie() {
	ns.DestroyEveryChat()
	ns.ObjectGroupOff(gvar62)
}
func CancelRobbery() {
	flag63 = false
	ns.ObjectGroupOff(gvar62)
}
func SecretSFX() {
	ns.MoveWaypoint(wp25, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretFound, wp25)
}
func Secret01Found() {
	ns.ObjectOff(ns.GetTrigger())
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 100)
	SecretSFX()
}
func Secret02Found() {
	ns.ObjectOff(ns.GetTrigger())
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 100)
	SecretSFX()
}
func Wasp01FollowPath() {
	ns.Move(obj26, wp28)
	ns.AggressionLevel(obj26, 0.83)
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
