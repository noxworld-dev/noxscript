package wiz06a

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	ivar4  int
	ivar5  int
	ivar6  int
	flag7  bool
	flag8  bool
	flag9  bool
	flag10 bool
	flag11 bool
	flag12 bool
	flag13 bool
	flag14 bool
	gvar15 int
	ivar16 int
	gvar17 int
	ivar18 int
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
	obj37  ns.ObjectID
	obj38  ns.ObjectID
	obj39  ns.ObjectID
	obj40  ns.ObjectID
	obj41  ns.ObjectID
	obj42  ns.ObjectID
	obj43  ns.ObjectID
	obj44  ns.ObjectID
	obj45  ns.ObjectID
	obj46  ns.ObjectID
	obj47  ns.ObjectID
	obj48  ns.ObjectID
	obj49  ns.ObjectID
	obj50  ns.ObjectID
	obj51  ns.ObjectID
	obj52  ns.ObjectID
	gvar53 ns.ObjectGroupID
	gvar54 ns.ObjectGroupID
	gvar55 ns.ObjectGroupID
	gvar56 ns.WallGroupID
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
	obj76  ns.ObjectID
	gvar77 ns.ObjectGroupID
	wp78   ns.WaypointID
	wp79   ns.WaypointID
	wp80   ns.WaypointID
	wp81   ns.WaypointID
)

func init() {
	ivar4 = 60
	ivar5 = 0
	ivar6 = 90
	flag7 = true
	flag8 = true
	flag9 = true
	flag10 = true
	flag11 = true
	flag12 = true
	flag13 = true
	flag14 = true
	gvar15 = 0
	ivar16 = 0
	gvar17 = 0
	ivar18 = 60
}
func SayCountdown() {
	ivar5 -= 1
	if !(ivar5 <= 0) {
		goto LABEL1
	}
	ivar5 = 0
	goto LABEL2
LABEL1:
	ns.FrameTimer(30, SayCountdown)
LABEL2:
	return
}
func WarriorRecognize() {
	var v0 int
	if !(ns.IsCaller(ns.GetHost()) && ivar5 == 0) {
		goto LABEL1
	}
	v0 = ns.Random(0, 3)
	ivar5 = ivar4
	SayCountdown()
	if v0 != 0 {
		goto LABEL2
	}
	ns.ChatTimer(ns.GetTrigger(), "Wiz06a:Guard2Recognize", ivar6)
	goto LABEL1
LABEL2:
	ns.ChatTimer(ns.GetTrigger(), "Wiz06a:Guard1Recognize", ivar6)
LABEL1:
	return
}
func WarriorEngage() {
	var v0 int
	if !(ns.IsCaller(ns.GetHost()) && ivar5 == 0) {
		goto LABEL1
	}
	v0 = ns.Random(0, 1)
	ivar5 = ivar4
	SayCountdown()
	if v0 != 0 {
		goto LABEL2
	}
	ns.ChatTimer(ns.GetTrigger(), "Wiz06a:Guard1Recognize", ivar6)
	goto LABEL1
LABEL2:
	ns.ChatTimer(ns.GetTrigger(), "Wiz06a:Guard2Recognize", ivar6)
LABEL1:
	return
}
func WarriorListen() {
	if !(ns.IsCaller(ns.GetHost()) && ivar5 == 0) {
		goto LABEL1
	}
	if ns.IsVisibleTo(ns.GetTrigger(), ns.GetHost()) {
		goto LABEL1
	}
	ivar5 = ivar4
	SayCountdown()
	ns.ChatTimer(ns.GetTrigger(), "Wiz06a:Guard2Listen", ivar6)
LABEL1:
	return
}
func delwinTalkStart() {
	ns.TellStory(ns.UrchinRecognize, "Wiz06b:Delwin1")
	ns.Frozen(ns.GetHost(), true)
	ns.ObjectGroupOff(gvar53)
}
func delwinTalkEnd() {
	ns.Frozen(ns.GetHost(), false)
	ns.ObjectGroupOn(gvar53)
	ns.CancelDialog(obj37)
}
func delwinLeaves() {
	if !flag13 {
		goto LABEL1
	}
	if !ns.IsVisibleTo(obj37, ns.GetHost()) {
		goto LABEL1
	}
	flag13 = false
	ns.Walk(obj37, ns.GetObjectX(ns.GetHost())+50, ns.GetObjectY(ns.GetHost()))
	ns.SetDialog(obj37, ns.NORMAL, delwin1TalkStart, delwin1TalkEnd)
	ns.StartDialog(obj37, ns.GetHost())
LABEL1:
	return
}
func delwin1TalkStart() {
	ns.TellStory(ns.UrchinRecognize, "Wiz06b:Delwin2")
	ns.AggressionLevel(obj37, 0)
	ns.ObjectGroupOff(gvar53)
}
func delwin1TalkEnd() {
	ns.CancelDialog(obj37)
	ns.JournalEdit(ns.GetHost(), "Wiz6SneakIn", 4)
	ns.JournalEntry(ns.GetHost(), "Wiz6Bull", 2)
	ns.JournalEntry(ns.GetHost(), "Wiz6Fortress", 2)
	ns.CreateObject("BluePotion", wp75)
	ns.Frozen(ns.GetHost(), false)
	ns.CreatureIdle(obj37)
	ns.AggressionLevel(obj37, 0.16)
	ns.ObjectGroupOn(gvar53)
	ns.GiveXp(ns.GetHost(), 2500)
	delwinGoesHome()
}
func delwinGoesHome() {
	if !flag14 {
		goto LABEL1
	}
	if ns.IsVisibleTo(ns.GetHost(), obj37) {
		goto LABEL2
	}
	flag14 = false
	ns.BecomeEnemy(obj37)
	ns.MoveObject(obj37, 3749, 3497)
LABEL2:
	ns.FrameTimer(10, delwinGoesHome)
LABEL1:
	return
}
func heardYou() {
	if !flag12 {
		goto LABEL1
	}
	flag12 = false
	ns.HitLocation(ns.GetTrigger(), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.Chat(ns.GetTrigger(), "Wiz06b:Guard1")
	ns.FrameTimer(90, hearingRestored)
LABEL1:
	return
}
func hearingRestored() {
	flag12 = true
}
func sawYou() {
	if !ns.IsVisibleTo(ns.GetTrigger(), ns.GetHost()) {
		goto LABEL1
	}
	if !flag11 {
		goto LABEL1
	}
	flag11 = false
	ns.Chat(ns.GetTrigger(), "Wiz06b:Guard2")
	ns.FrameTimer(360, sightRestored)
LABEL1:
	return
}
func sightRestored() {
	flag11 = true
}
func march() {
	ns.Wander(ns.GetTrigger())
}
func freezePlayer() {
	ns.WideScreen(true)
	ns.Frozen(ns.GetHost(), true)
	ns.FrameTimer(15, captainTalks)
}
func captainTalks() {
	ns.StartDialog(obj19, ns.GetHost())
}
func helloCaptain() {
	ns.TellStory(ns.ArcherRecognize, "Wiz06a:Captain1")
}
func goodbyeCaptain() {
	ns.MoveObject(obj24, 3989, 4542)
	ns.MoveObject(obj25, 4014, 4541)
	ns.MoveObject(obj26, 3954, 4542)
	ns.SetDialog(obj19, ns.NORMAL, helloCaptain2, goodbyeCaptain2)
	ns.StartDialog(obj19, ns.GetHost())
}
func helloCaptain2() {
	ns.TellStory(ns.ArcherRecognize, "Wiz06a:Captain2")
}
func goodbyeCaptain2() {
	ns.CancelDialog(obj19)
	ns.WideScreen(false)
	ns.Frozen(ns.GetHost(), false)
	ns.JournalEntry(ns.GetHost(), "Wiz6SneakIn", 2)
	ns.FrameTimer(3, captainVanishes)
}
func captainVanishes() {
	if !flag10 {
		goto LABEL1
	}
	if ns.IsVisibleTo(ns.GetHost(), obj19) {
		goto LABEL2
	}
	flag10 = false
	ns.Delete(obj19)
	ns.Delete(obj20)
	ns.Delete(obj21)
LABEL2:
	ns.FrameTimer(10, captainVanishes)
LABEL1:
	return
}
func guard1Listen() {
	if !flag8 {
		goto LABEL1
	}
	flag8 = false
	if !flag7 {
		goto LABEL2
	}
	ivar16 += 1
	flag7 = false
	ns.StoryPic(obj22, "Warrior3Pic")
	ns.SetDialog(obj22, ns.NORMAL, guard1TalkStart, guard1TalkEnd)
	ns.StartDialog(obj22, ns.GetHost())
	goto LABEL3
LABEL2:
	ns.Chat(obj22, "Wiz06a:Guard1Listen")
	ivar16 += 1
	if !(ivar16 > 2) {
		goto LABEL3
	}
	ivar16 = 3
	ns.Walk(obj22, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
LABEL3:
	ns.FrameTimer(180, unpaused1)
LABEL1:
	return
}
func guard1Recognize() {
	if !flag7 {
		goto LABEL1
	}
	flag7 = false
LABEL1:
	if !flag8 {
		goto LABEL2
	}
	flag8 = false
	if flag7 {
		goto LABEL3
	}
	ns.Chat(obj22, "Wiz06a:Guard1Recognize")
LABEL3:
	ns.FrameTimer(180, unpaused1)
LABEL2:
	return
}
func guard1TalkStart() {
	ns.Frozen(obj22, true)
	ns.Frozen(obj23, true)
	ns.Frozen(ns.GetHost(), true)
	ns.TellStory(ns.SwordsmanRecognize, "Wiz06a:Guard1Listen")
}
func guard1TalkEnd() {
	ns.Frozen(obj22, false)
	ns.Frozen(obj23, false)
	ns.Frozen(ns.GetHost(), false)
	ns.CancelDialog(obj22)
}
func unpaused1() {
	flag8 = true
}
func guard2Listen() {
	if !flag9 {
		goto LABEL1
	}
	flag9 = false
	ns.Chat(obj23, "Wiz06a:Guard2Listen")
	if !(ivar16 > 2) {
		goto LABEL2
	}
	ivar16 = 3
	ns.Walk(obj23, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
LABEL2:
	ns.FrameTimer(210, unpaused2)
LABEL1:
	return
}
func guard2Recognize() {
	if !flag7 {
		goto LABEL1
	}
	flag7 = false
LABEL1:
	if !flag9 {
		goto LABEL2
	}
	flag9 = false
	ns.Chat(obj23, "Wiz06a:Guard2Recognize")
	ns.FrameTimer(210, unpaused2)
LABEL2:
	return
}
func unpaused2() {
	flag9 = true
}
func goHome1() {
	ns.GoBackHome(obj22)
}
func goHome2() {
	ns.GoBackHome(obj23)
}
func ResetSquishTrapSEG1() {
	ns.Move(obj29, wp66)
	ns.Move(obj30, wp67)
	ns.Move(obj31, wp68)
	ns.Move(obj32, wp69)
	ns.Move(obj33, wp70)
	ns.Move(obj34, wp71)
	ns.Move(obj35, wp72)
	ns.Move(obj36, wp73)
	ns.AudioEvent(ns.BoulderMove, wp57)
	ns.FrameTimer(20, ResetSquishTrapSEG2)
}
func ResetSquishTrapSEG2() {
	ns.ObjectOn(obj27)
	ns.ObjectOn(obj28)
}
func InitializeSquishBlocks() {
	obj27 = ns.Object("SquishTrigger")
	obj28 = ns.Object("BoomTrigger")
	obj29 = ns.Object("SquishBlock01")
	obj30 = ns.Object("SquishBlock02")
	obj31 = ns.Object("SquishBlock03")
	obj32 = ns.Object("SquishBlock04")
	obj33 = ns.Object("SquishBlock05")
	obj34 = ns.Object("SquishBlock06")
	obj35 = ns.Object("SquishBlock07")
	obj36 = ns.Object("SquishBlock08")
	wp57 = ns.Waypoint("SquishAudioOrigin")
	wp58 = ns.Waypoint("SquishWP01")
	wp59 = ns.Waypoint("SquishWP02")
	wp60 = ns.Waypoint("SquishWP03")
	wp61 = ns.Waypoint("SquishWP04")
	wp62 = ns.Waypoint("SquishWP05")
	wp63 = ns.Waypoint("SquishWP06")
	wp64 = ns.Waypoint("SquishWP07")
	wp65 = ns.Waypoint("SquishWP08")
	wp66 = ns.Waypoint("SquishWPHome01")
	wp67 = ns.Waypoint("SquishWPHome02")
	wp68 = ns.Waypoint("SquishWPHome03")
	wp69 = ns.Waypoint("SquishWPHome04")
	wp70 = ns.Waypoint("SquishWPHome05")
	wp71 = ns.Waypoint("SquishWPHome06")
	wp72 = ns.Waypoint("SquishWPHome07")
	wp73 = ns.Waypoint("SquishWPHome08")
}
func MapInitialize() {
	ns.StartupScreen(6)
	obj19 = ns.Object("Captain")
	obj20 = ns.Object("Basket")
	obj21 = ns.Object("BasketShadow")
	obj22 = ns.Object("Guard1")
	obj23 = ns.Object("Guard2")
	obj24 = ns.Object("Invisibility")
	obj25 = ns.Object("Lock")
	obj26 = ns.Object("Trap")
	obj37 = ns.Object("Delwin")
	obj38 = ns.Object("SweeperMover1")
	obj39 = ns.Object("SweeperMover2")
	obj40 = ns.Object("SweeperMover3")
	obj41 = ns.Object("SweeperMover4")
	obj42 = ns.Object("SweeperMover5")
	obj43 = ns.Object("SweeperMover6")
	obj44 = ns.Object("SweeperMover7")
	obj45 = ns.Object("SweeperMover8")
	obj46 = ns.Object("SweeperMover9")
	obj47 = ns.Object("SweeperMover10")
	obj48 = ns.Object("SweeperMover11")
	obj49 = ns.Object("SweeperMover12")
	obj50 = ns.Object("SweeperMover13")
	obj51 = ns.Object("SweeperMover14")
	obj52 = ns.Object("SewerElevator")
	gvar55 = ns.ObjectGroup("SewerMonsters")
	gvar56 = ns.WallGroup("SewerMonsterWalls")
	gvar53 = ns.ObjectGroup("EveryMonsterOnMap")
	wp74 = ns.Waypoint("PlayerSounds")
	wp75 = ns.Waypoint("PotionWay")
	ns.StoryPic(obj37, "Townsman1Pic")
	ns.ObjectOn(obj38)
	ns.ObjectOn(obj39)
	ns.ObjectOn(obj40)
	ns.ObjectOn(obj41)
	ns.ObjectOn(obj42)
	ns.ObjectOn(obj43)
	ns.ObjectOn(obj44)
	ns.ObjectOn(obj45)
	ns.ObjectOn(obj46)
	ns.StoryPic(obj19, "AirshipCaptainPic")
	ns.SetDialog(obj19, ns.NEXT, helloCaptain, goodbyeCaptain)
	obj76 = ns.Object("TrapSpellBook")
	gvar54 = ns.ObjectGroup("RoamingGuards")
	gvar77 = ns.ObjectGroup("DelwinTriggers")
	wp78 = ns.Waypoint("DelwinGreetWP")
	wp81 = ns.Waypoint("DelwinGuardWP")
	wp79 = ns.Waypoint("DelwinLookWP")
	wp80 = ns.Waypoint("TrapSpellWP")
	ns.FrameTimer(3, InitializeSquishBlocks)
	ns.FrameTimer(5, freezePlayer)
	ns.GroupWander(gvar54)
	ns.SetOwner(ns.GetHost(), obj19)
	ns.Music(18, 100)
}
func PlayerDeath() {
	ns.DeathScreen(6)
}
func Squish() {
	ns.ObjectOff(obj27)
	ns.Move(obj29, wp58)
	ns.Move(obj30, wp59)
	ns.Move(obj31, wp60)
	ns.Move(obj32, wp61)
	ns.Move(obj33, wp62)
	ns.Move(obj34, wp63)
	ns.Move(obj35, wp64)
	ns.Move(obj36, wp65)
	ns.AudioEvent(ns.BoulderMove, wp57)
	ns.FrameTimer(ivar18, ResetSquishTrapSEG1)
}
func SquishBoom() {
	if ns.IsCaller(obj31) {
		goto LABEL1
	}
	return
LABEL1:
	ns.ObjectOff(obj28)
	ns.AudioEvent(ns.HammerMissing, wp57)
	ns.Effect(ns.JIGGLE, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 10, 0)
}
func OpenBlockSecretExit() {
	ns.WallOpen(ns.Wall(178, 46))
}
func SewerSecretFound() {
	ns.MoveWaypoint(wp74, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretFound, wp74)
	ns.ObjectGroupOff(ns.ObjectGroup("SewerSecretTriggers"))
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 100)
}
func Secret01Found() {
	ns.MoveWaypoint(wp74, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretFound, wp74)
	ns.ObjectOff(ns.GetTrigger())
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 100)
}
func ActivateSewerElevator() {
	ns.ObjectOff(ns.GetTrigger())
	ns.ObjectOn(obj52)
	ns.ObjectGroupOn(gvar55)
	ns.WallGroupOpen(gvar56)
}
func OnEvent(typ string) {
	switch typ {
	case "MapInitialize":
		MapInitialize()
	case "PlayerDeath":
		PlayerDeath()
	}
}
