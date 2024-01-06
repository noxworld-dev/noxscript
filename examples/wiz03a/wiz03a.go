package wiz03a

import (
	"strconv"

	"github.com/noxworld-dev/noxscript/ns/v3"
)

var (
	obj4   ns.ObjectID
	wp5    ns.WaypointID
	wp6    [70]ns.WaypointID
	obj7   ns.ObjectID
	obj8   ns.ObjectID
	gvar9  ns.ObjectGroupID
	obj10  ns.ObjectID
	obj11  ns.ObjectID
	obj12  ns.ObjectID
	obj13  ns.ObjectID
	wp14   ns.WaypointID
	obj15  ns.ObjectID
	obj16  ns.ObjectID
	gvar17 ns.WallGroupID
	gvar18 int
	gvar19 int
	gvar20 int
	gvar21 int
	gvar22 int
	wp23   ns.WaypointID
	wp24   ns.WaypointID
	obj25  ns.ObjectID
	wp26   ns.WaypointID
	wp27   ns.WaypointID
	wp28   ns.WaypointID
	wp29   ns.WaypointID
	obj30  ns.ObjectID
	gvar31 ns.ObjectID
	gvar32 ns.ObjectID
	gvar33 ns.ObjectGroupID
	obj34  ns.ObjectID
	obj35  ns.ObjectID
	gvar36 ns.ObjectGroupID
	gvar37 ns.ObjectGroupID
	gvar38 ns.WallGroupID
	wp39   ns.WaypointID
	gvar40 int
	gvar41 int
	gvar42 int
	obj43  ns.ObjectID
	gvar44 ns.ObjectGroupID
	obj45  ns.ObjectID
	gvar46 ns.ObjectGroupID
	gvar47 int
	gvar48 [6]int
	gvar49 int
	gvar50 int
)

func init() {
	gvar18 = 0
	gvar19 = 1
	gvar20 = 2
	gvar21 = 3
	gvar22 = gvar18
	gvar40 = 0
	gvar41 = 1
	gvar42 = gvar40
}
func ArcherIdle() {
	ns.AggressionLevel(obj4, 1)
	ns.CreatureIdle(obj4)
}
func MissileAttackBarrel() {
	var (
		v0 float32
		v1 float32
	)
	v0 = ns.GetWaypointX(wp5)
	v1 = ns.GetWaypointY(wp5)
	ns.HitFarLocation(obj4, v0, v1)
}
func AttackBarrel() {
	ns.FrameTimer(3, MissileAttackBarrel)
}
func FlamesCreate() {
	var (
		v0 int
		v1 int
	)
	v0 = 0
	for {
		if !(v0 < 70) {
			goto LABEL1
		}
		v1 = v0 % 3
		if v1 == 0 {
			goto LABEL2
		}
		if v1 == 1 {
			goto LABEL3
		}
		if v1 == 2 {
			goto LABEL4
		}
		goto LABEL5
	LABEL2:
		ns.CreateObject("MediumFlame", wp6[v0])
		goto LABEL5
	LABEL3:
		ns.CreateObject("Flame", wp6[v0])
		goto LABEL5
	LABEL4:
		ns.CreateObject("SmallFlame", wp6[v0])
		goto LABEL5
	LABEL5:
		v0 += 1
	}
LABEL1:
	return
}
func PlayerDeath() {
	ns.DeathScreen(3)
}
func HorvathTalkStart() {
	var v0 int
	if !ns.HasItem(ns.GetHost(), ns.Object("Wiz03b:AmuletofTeleportation")) {
		goto LABEL1
	}
	gvar22 = gvar21
LABEL1:
	v0 = gvar22
	if v0 == gvar18 {
		goto LABEL2
	}
	if v0 == gvar19 {
		goto LABEL3
	}
	if v0 == gvar20 {
		goto LABEL4
	}
	if v0 == gvar21 {
		goto LABEL5
	}
	goto LABEL6
LABEL2:
	ns.TellStory(ns.SwordsmanHurt, "Wiz03a:HorvathIntro")
	goto LABEL6
LABEL3:
	ns.TellStory(ns.SwordsmanHurt, "Wiz03a:HorvathPre")
	goto LABEL6
LABEL4:
	ns.TellStory(ns.SwordsmanHurt, "Wiz03a:HorvathPre2")
	goto LABEL6
LABEL5:
	ns.TellStory(ns.SwordsmanHurt, "Wiz03a:HorvathFinish")
	goto LABEL6
	ns.Blind()
	ns.FrameTimer(60, PlayerExit)
	goto LABEL6
LABEL6:
	return
}
func HorvathTalkEnd() {
	var v0 int
	v0 = gvar22
	if v0 == gvar18 {
		goto LABEL1
	}
	if v0 == gvar19 {
		goto LABEL2
	}
	if v0 == gvar21 {
		goto LABEL3
	}
	goto LABEL4
LABEL1:
	ns.ObjectOn(obj16)
	ns.JournalEntry(ns.GetHost(), "Wiz03aGetAmulet", 2)
	ns.Print("Con01a:NewJournalEntry")
	gvar22 = gvar19
	goto LABEL4
LABEL2:
	gvar22 = gvar20
	goto LABEL4
LABEL3:
	ns.Frozen(ns.GetHost(), true)
	ns.JournalEdit(ns.GetHost(), "Journal:Wiz03aGetAmulet", 4)
	gvar31 = ns.GetLastItem(ns.GetHost())
	for {
		if gvar31 == 0 {
			goto LABEL5
		}
		gvar32 = ns.GetPreviousItem(gvar31)
		if gvar31 != ns.Object("Wiz03b:AmuletofTeleportation") {
			goto LABEL6
		}
		ns.Delete(gvar31)
	LABEL6:
		gvar31 = gvar32
	}
LABEL5:
	ns.CancelDialog(obj10)
	ns.Blind()
	ns.FrameTimer(60, PlayerExit)
	goto LABEL4
LABEL4:
	return
}
func AirshipCaptainAppear() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
	)
	v0 = ns.GetWaypointX(wp26)
	v1 = ns.GetWaypointY(wp26)
	v2 = ns.GetWaypointX(wp27)
	v3 = ns.GetWaypointY(wp27)
	ns.MoveObject(obj25, v0, v1)
	ns.CreatureGuard(obj25, v0, v1, v2, v3, 500)
}
func PlayerExit() {
	var (
		v0 float32
		v1 float32
	)
	v0 = ns.GetWaypointX(wp28)
	v1 = ns.GetWaypointY(wp28)
	ns.MoveObject(ns.GetHost(), v0, v1)
	ns.Frozen(ns.GetHost(), false)
}
func MissionStart() {
	ns.ObjectOff(obj8)
}
func BrennethSpeaks() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
	)
	v0 = ns.GetObjectX(obj35)
	v1 = ns.GetObjectY(obj35)
	v2 = ns.GetObjectX(ns.GetHost())
	v3 = ns.GetObjectY(ns.GetHost())
	if !(ns.Distance(v0, v1, v2, v3) < 230) {
		goto LABEL1
	}
	ns.LookAtObject(obj35, ns.GetHost())
	ns.StartDialog(obj35, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Wiz03a:BrennethExplains")
	goto LABEL2
LABEL1:
	ns.FrameTimer(30, BrennethSpeaks)
LABEL2:
	return
}
func BrennethTalkStart() {
	var v0 int
	v0 = gvar42
	if v0 == gvar41 {
		goto LABEL1
	}
	goto LABEL2
LABEL1:
	ns.LookAtObject(obj35, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Wiz03a:BrennethEnd")
	goto LABEL2
LABEL2:
	return
}
func BrennethTalkEnd() {
	var v0 int
	v0 = gvar42
	if v0 == gvar40 {
		goto LABEL1
	}
	if v0 == gvar41 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.Move(obj35, wp39)
	ns.CreatureGuard(obj35, 4938, 1032, 4923, 1078, 500)
	gvar42 = gvar41
	goto LABEL3
LABEL2:
	goto LABEL3
LABEL3:
	return
}
func MapInitialize() {
	var v0 int
	obj8 = ns.Object("MissionTrig")
	obj10 = ns.Object("Horvath")
	gvar9 = ns.ObjectGroup("HorvathTrigGroup")
	ns.StoryPic(obj10, "HorvathPic")
	ns.SetDialog(obj10, ns.NORMAL, HorvathTalkStart, HorvathTalkEnd)
	obj11 = ns.Object("HorvathDoor1")
	obj12 = ns.Object("HorvathDoor2")
	ns.SetOwner(ns.GetHost(), obj10)
	obj13 = ns.Object("HorvathBookcaseMover1")
	obj15 = ns.Object("HorvathBookcaseMover2")
	obj16 = ns.Object("HorvathBookcaseTrigger")
	ns.LockDoor(obj11)
	ns.LockDoor(obj12)
	wp23 = ns.Waypoint("JournalSoundWP")
	wp14 = ns.Waypoint("HorvathBookcase1WP")
	obj25 = ns.Object("Airship_Captain")
	wp26 = ns.Waypoint("AirshipCaptainWP")
	wp27 = ns.Waypoint("AirshipCaptainLook")
	ns.SetOwner(ns.GetHost(), obj25)
	gvar38 = ns.WallGroup("InvisibleWallGroup")
	gvar33 = ns.ObjectGroup("ElevatorTrigGroup")
	obj34 = ns.Object("CellarElevator")
	obj35 = ns.Object("Brenneth")
	gvar36 = ns.ObjectGroup("BrennethGroup")
	ns.GroupSetOwner(ns.GetHost(), gvar36)
	ns.StoryPic(obj35, "MalePic8")
	ns.SetDialog(obj35, ns.NORMAL, BrennethTalkStart, BrennethTalkEnd)
	wp39 = ns.Waypoint("BrennethWP")
	gvar37 = ns.ObjectGroup("GearGroup")
	obj43 = ns.Object("Bandit01")
	ns.Damage(obj43, 0, 100, 5)
	gvar46 = ns.ObjectGroup("MonsterGroup1")
	ns.ObjectGroupOff(gvar46)
	gvar44 = ns.ObjectGroup("MissionExitGroup")
	obj45 = ns.Object("Wiz03b:AmuletofTeleportation")
	gvar17 = ns.WallGroup("HorvathWallGroup")
	wp28 = ns.Waypoint("ExitWP")
	wp29 = ns.Waypoint("Wiz03bExitWP")
	obj30 = ns.Object("ExitTrig")
	obj4 = ns.Object("Archer01")
	wp5 = ns.Waypoint("Barrel1WP")
	obj7 = ns.Object("Barrel01")
	wp24 = ns.Waypoint("Secret01WP")
	ns.StartupScreen(3)
	v0 = 0
	for {
		if !(v0 < 70) {
			goto LABEL1
		}
		wp6[v0] = ns.Waypoint(strconv.Itoa(v0+1) + "WP")
		v0 += 1
	}
LABEL1:
	return
}
func MapEntry() {
	ns.UnBlind()
	ns.ObjectOn(obj30)
}
func HorvathPickup() {
	ns.Pickup(obj10, ns.Object("Wiz03b:AmuletofTeleportation"))
}
func OpenHorvathWalls() {
	ns.ObjectOff(ns.GetTrigger())
	ns.ObjectOn(obj13)
	ns.ObjectOn(obj15)
	ns.WallGroupOpen(gvar17)
}
func QuestSwitch() {
	if !ns.HasItem(ns.GetHost(), ns.Object("Wiz03b:AmuletofTeleportation")) {
		goto LABEL1
	}
	ns.MoveObject(obj35, 5457, 1685)
	ns.ObjectGroupOn(gvar37)
	ns.ObjectOn(obj34)
	AirshipCaptainAppear()
LABEL1:
	return
}
func QuestXP() {
	if !ns.HasItem(ns.GetHost(), ns.Object("Wiz03b:AmuletofTeleportation")) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.GiveXp(ns.GetHost(), 1000)
	ns.AudioEvent(ns.FlagDrop, wp23)
	ns.Music(0, 100)
LABEL1:
	return
}
func DisableElevator() {
	ns.NoWallSound(true)
	ns.WallGroupOpen(gvar38)
	ns.NoWallSound(false)
	ns.ObjectOff(obj34)
	ns.ObjectGroupOff(gvar37)
	ns.ObjectGroupOff(gvar33)
	ns.CreatureFollow(obj35, ns.GetHost())
	ns.AudioEvent(ns.WoodenArmorBreak, wp39)
	ns.AudioEvent(ns.MetalArmorBreak, wp39)
	ns.AudioEvent(ns.WallDestroyed, wp39)
	ns.AudioEvent(ns.LightningWand, wp39)
	ns.Effect(ns.SMOKE_BLAST, 4953, 1047, 4953, 1047)
	ns.Effect(ns.BLUE_SPARKS, 4953, 1047, 4953, 1047)
	BrennethSpeaks()
}
func MonsterGoHome() {
	r1 := ns.IsAttackedBy(ns.GetCaller(), ns.GetHost())
	if !r1 {
		goto LABEL1
	}
	ns.GoBackHome(ns.GetCaller())
LABEL1:
	return
}
func MonsterGroup1Enable() {
	ns.ObjectGroupOn(gvar46)
}
func PlayerWiz03bExit() {
	var (
		v0 float32
		v1 float32
	)
	v0 = ns.GetWaypointX(wp29)
	v1 = ns.GetWaypointY(wp29)
	ns.MoveObject(ns.GetHost(), v0, v1)
}
func FreezePlayer() {
	ns.Frozen(ns.GetHost(), true)
}
func GotoWiz03b() {
	ns.ObjectOff(obj30)
	ns.Blind()
	ns.FrameTimer(30, FreezePlayer)
	ns.FrameTimer(60, PlayerWiz03bExit)
}
func QuestCheck() {
	if !ns.HasItem(ns.GetHost(), ns.Object("Wiz03b:AmuletofTeleportation")) {
		goto LABEL1
	}
	ns.Chat(ns.GetHost(), "I got it!")
LABEL1:
	return
}
func Secret50XP() {
	ns.GiveXp(ns.GetHost(), 50)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.AudioEvent(ns.SecretFound, wp24)
}
func Award100XP() {
	ns.GiveXp(ns.GetHost(), 100)
}
func PlaySubMusic() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.Music(19, 100)
LABEL1:
	return
}
func PlayNoMusic() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.Music(0, 100)
LABEL1:
	return
}
func PlayWanderMusic() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.Music(22, 100)
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
