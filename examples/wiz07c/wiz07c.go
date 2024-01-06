package wiz07c

import (
	"strconv"

	"github.com/noxworld-dev/noxscript/ns/v3"
)

var (
	flag4  bool
	gvar5  int
	gvar6  int
	gvar7  int
	gvar8  int
	gvar9  int
	gvar10 int
	gvar11 int
	gvar12 int
	ivar13 int
	ivar14 int
	ivar15 int
	obj16  [20]ns.ObjectID
	wp17   [20]ns.WaypointID
	wp18   ns.WaypointID
	obj19  ns.ObjectID
	obj20  ns.ObjectID
	obj21  ns.ObjectID
	obj22  ns.ObjectID
	gvar23 ns.WallGroupID
	obj24  ns.ObjectID
	obj25  ns.ObjectID
	obj26  ns.ObjectID
	obj27  ns.ObjectID
	obj28  ns.ObjectID
	obj29  ns.ObjectID
	wp30   ns.WaypointID
	wp31   ns.WaypointID
	wp32   ns.WaypointID
	obj33  ns.ObjectID
	wp34   ns.WaypointID
	wp35   ns.WaypointID
	wp36   ns.WaypointID
	wp37   ns.WaypointID
	wp38   ns.WaypointID
	wp39   ns.WaypointID
	wp40   ns.WaypointID
	obj41  ns.ObjectID
	obj42  ns.ObjectID
	obj43  ns.ObjectID
	obj44  ns.ObjectID
	gvar45 ns.ObjectGroupID
	gvar46 ns.ObjectGroupID
	wp47   ns.WaypointID
	wp48   ns.WaypointID
	wp49   ns.WaypointID
	wp50   ns.WaypointID
	ivar51 int
	obj52  [4]ns.ObjectID
	ivar53 [4]int
	obj54  [4]ns.ObjectID
	ivar55 [4]int
	gvar56 int
	flag57 bool
	gvar58 ns.ObjectGroupID
	wp59   ns.WaypointID
	obj60  ns.ObjectID
	wp61   ns.WaypointID
	wp62   ns.WaypointID
	gvar63 ns.ObjectGroupID
	obj64  ns.ObjectID
	wp65   ns.WaypointID
	wp66   ns.WaypointID
	wp67   [9]ns.WaypointID
	wp68   [9]ns.WaypointID
	wp69   [9]ns.WaypointID
	obj70  ns.ObjectID
	obj71  ns.ObjectID
	obj72  ns.ObjectID
	obj73  ns.ObjectID
	fvar74 float32
	fvar75 float32
	fvar76 float32
	gvar77 int
	ivar78 int
)

func init() {
	gvar5 = 0
	gvar6 = 1
	gvar7 = 2
	gvar8 = 3
	gvar9 = 4
	gvar10 = 5
	gvar11 = gvar5
	gvar12 = 0
	ivar13 = 0
	gvar56 = 1
	flag57 = true
	flag4 = true
	fvar74 = 4
	fvar75 = 12
	gvar77 = 5
	ivar78 = 1000
}
func MasterDialogStart() {
	if gvar12 != 3 {
		goto LABEL1
	}
	ns.TellStory(ns.DemonRecognize, "Wiz07C.scr:DemonTalk01")
LABEL1:
	if gvar12 != 4 {
		goto LABEL2
	}
	ns.TellStory(ns.DemonRecognize, "Wiz07C.scr:DemonTalk04")
LABEL2:
	return
}
func MasterDialogEnd() {
	if gvar12 != 3 {
		goto LABEL1
	}
	ns.CancelDialog(obj19)
	ns.FrameTimer(100, StartDemonTalk)
	MoveLine()
LABEL1:
	if gvar12 != 4 {
		goto LABEL2
	}
	ns.CancelDialog(obj19)
	ns.FrameTimer(10, StartDemonTalk)
LABEL2:
	return
}
func MoveLineState(a1 ns.ObjectID, a2 int) {
	var v0 int
	v0 = a2
	if v0 == 1 {
		goto LABEL1
	}
	if v0 == 3 {
		goto LABEL2
	}
	if v0 == 5 {
		goto LABEL3
	}
	if v0 == 7 {
		goto LABEL4
	}
	if v0 == 2 {
		goto LABEL5
	}
	if v0 == 4 {
		goto LABEL6
	}
	if v0 == 6 {
		goto LABEL7
	}
	if v0 == 8 {
		goto LABEL8
	}
	goto LABEL9
LABEL1:
	ns.Move(a1, wp34)
	goto LABEL10
LABEL2:
	ns.Move(a1, wp36)
	goto LABEL10
LABEL3:
	ns.Move(a1, wp38)
	goto LABEL10
LABEL4:
	ns.Move(a1, wp40)
	goto LABEL10
LABEL5:
	ns.Move(a1, wp34)
	goto LABEL10
LABEL6:
	ns.Move(a1, wp35)
	goto LABEL10
LABEL7:
	ns.Move(a1, wp37)
	goto LABEL10
LABEL8:
	ns.Move(a1, wp39)
	goto LABEL10
LABEL9:
	goto LABEL10
LABEL10:
	return
}
func CreateFireWall() {
	var v0 int
	v0 = gvar11
	if v0 == gvar5 {
		goto LABEL1
	}
	if v0 == gvar6 {
		goto LABEL2
	}
	if v0 == gvar7 {
		goto LABEL3
	}
	if v0 == gvar8 {
		goto LABEL4
	}
	goto LABEL5
LABEL1:
	ivar51 = 0
	for {
		if !(ivar51 < 20) {
			goto LABEL6
		}
		ns.Delete(obj16[ivar51])
		obj16[ivar51] = ns.CreateObject("SmallFlame", wp17[ivar51])
		ivar51 += 1
	}
LABEL6:
	gvar11 = gvar6
	ivar14 = 2
	goto LABEL5
LABEL2:
	ivar51 = 0
	for {
		if !(ivar51 < 18) {
			goto LABEL8
		}
		ns.Delete(obj16[ivar51])
		obj16[ivar51] = ns.CreateObject("MediumFlame", wp17[ivar51])
		ivar51 += 1
	}
LABEL8:
	gvar11 = gvar7
	ivar14 = 2
	goto LABEL5
LABEL3:
	ivar51 = 0
	for {
		if !(ivar51 < 13) {
			goto LABEL10
		}
		ns.Delete(obj16[ivar51])
		obj16[ivar51] = ns.CreateObject("Flame", wp17[ivar51])
		ivar51 += 1
	}
LABEL10:
	gvar11 = gvar8
	ivar14 = 3
	goto LABEL5
LABEL4:
	ivar51 = 0
	for {
		if !(ivar51 < 10) {
			goto LABEL12
		}
		ns.Delete(obj16[ivar51])
		obj16[ivar51] = ns.CreateObject("LargeFlame", wp17[ivar51])
		ivar51 += 1
	}
LABEL12:
	gvar11 = gvar9
	goto LABEL5
LABEL5:
	if gvar11 == gvar9 {
		goto LABEL14
	}
	ns.FrameTimer(ivar14, CreateFireWall)
LABEL14:
	return
}
func ChapterEndB() {
	ns.Frozen(ns.GetHost(), false)
	ns.MoveObject(ns.GetHost(), ns.GetWaypointX(wp59), ns.GetWaypointY(wp59))
}
func DestroyFireWall() {
	var v0 int
	v0 = gvar11
	if v0 == gvar5 {
		goto LABEL1
	}
	if v0 == gvar6 {
		goto LABEL2
	}
	if v0 == gvar7 {
		goto LABEL3
	}
	if v0 == gvar9 {
		goto LABEL4
	}
	goto LABEL5
LABEL1:
	ivar51 = 0
	for {
		if !(ivar51 < 20) {
			goto LABEL6
		}
		if !(ivar51 != 2 && ivar51 != 7 && ivar51 != 17) {
			goto LABEL7
		}
		ns.Delete(obj16[ivar51])
	LABEL7:
		ivar51 += 1
	}
LABEL6:
	gvar11 = gvar10
	goto LABEL5
LABEL2:
	ivar51 = 0
	for {
		if !(ivar51 < 18) {
			goto LABEL9
		}
		ns.Delete(obj16[ivar51])
		obj16[ivar51] = ns.CreateObject("SmallFlame", wp17[ivar51])
		ivar51 += 1
	}
LABEL9:
	ns.MoveObject(obj24, 1265, 2229)
	ns.MoveObject(obj25, 1204, 2245)
	ns.MoveObject(obj26, 1248, 2290)
	gvar11 = gvar5
	ivar14 = 300
	goto LABEL5
LABEL3:
	if !(ns.CurrentHealth(obj52[2]) > 0) {
		goto LABEL11
	}
	ns.Damage(obj52[2], 0, 1000, 0)
LABEL11:
	if !(ns.CurrentHealth(obj52[3]) > 0) {
		goto LABEL12
	}
	ns.Damage(obj52[3], 0, 1000, 0)
LABEL12:
	if !(ns.CurrentHealth(obj54[2]) > 0) {
		goto LABEL13
	}
	ns.Damage(obj54[2], 0, 1000, 0)
LABEL13:
	if !(ns.CurrentHealth(obj54[3]) > 0) {
		goto LABEL14
	}
	ns.Damage(obj54[3], 0, 1000, 0)
LABEL14:
	ns.WallGroupOpen(gvar23)
	ns.ObjectOn(obj33)
	ivar51 = 0
	for {
		if !(ivar51 < 13) {
			goto LABEL15
		}
		ns.Delete(obj16[ivar51])
		obj16[ivar51] = ns.CreateObject("MediumFlame", wp17[ivar51])
		ivar51 += 1
	}
LABEL15:
	gvar11 = gvar6
	ivar14 = 30
	goto LABEL5
LABEL4:
	ivar51 = 0
	for {
		if !(ivar51 < 10) {
			goto LABEL17
		}
		ns.Delete(obj16[ivar51])
		obj16[ivar51] = ns.CreateObject("Flame", wp17[ivar51])
		ivar51 += 1
	}
LABEL17:
	gvar11 = gvar7
	ivar14 = 10
	goto LABEL5
LABEL5:
	if gvar11 == gvar10 {
		goto LABEL19
	}
	ns.FrameTimer(ivar14, DestroyFireWall)
LABEL19:
	return
}
func PlaySubMusic() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.Music(18, 100)
LABEL1:
	return
}
func PlayActionMusic() {
	ns.Music(27, 100)
}
func FollowPath() {
	ns.Wander(ns.GetTrigger())
}
func UnlockDoor1() {
	ns.UnlockDoor(obj22)
	ns.AutoSave()
}
func MoveLine() {
	var v0 int
	if !flag57 {
		goto LABEL1
	}
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL2
		}
		if gvar56 == 0 {
			goto LABEL3
		}
		MoveLineState(obj52[v0], ivar53[v0])
		ivar53[v0] -= 2
		goto LABEL4
	LABEL3:
		MoveLineState(obj54[v0], ivar55[v0])
		ivar55[v0] -= 2
	LABEL4:
		v0 += 1
	}
LABEL2:
	gvar56 = 1 - gvar56
	ivar13 += 1
	if !(ivar13 == 1 || ivar13 == 3) {
		goto LABEL1
	}
	ns.FrameTimer(82, MoveLine)
LABEL1:
	return
}
func EnableLineD() {
	ns.Wander(ns.GetCaller())
}
func StartDemonTalk() {
	var v0 int
	v0 = gvar12
	if v0 == 0 {
		goto LABEL1
	}
	if v0 == 1 {
		goto LABEL2
	}
	if v0 == 2 {
		goto LABEL3
	}
	if v0 == 3 {
		goto LABEL4
	}
	if v0 == 4 {
		goto LABEL5
	}
	if v0 == 5 {
		goto LABEL6
	}
	if v0 == 6 {
		goto LABEL7
	}
	if v0 == 7 {
		goto LABEL8
	}
	goto LABEL9
LABEL1:
	flag4 = false
	ns.WideScreen(true)
	MoveLine()
	ivar15 = 3
	gvar12 = 1
	ns.FrameTimer(ivar15, StartDemonTalk)
	goto LABEL9
LABEL2:
	ns.LookAtObject(ns.GetHost(), obj19)
	ns.Frozen(ns.GetHost(), true)
	gvar12 = 2
	ivar15 = 100
	ns.FrameTimer(ivar15, StartDemonTalk)
	goto LABEL9
LABEL3:
	gvar12 = 3
	ns.StoryPic(obj19, "DemonPic")
	ns.SetDialog(obj19, ns.NEXT, MasterDialogStart, MasterDialogEnd)
	ns.StartDialog(obj19, ns.GetHost())
	goto LABEL9
LABEL4:
	gvar12 = 4
	if !ns.HasEnchant(ns.GetHost(), ns.ENCHANT_INVISIBLE) {
		goto LABEL10
	}
	ns.EnchantOff(ns.GetHost(), ns.ENCHANT_INVISIBLE)
LABEL10:
	flag57 = false
	ns.SetDialog(obj19, ns.NORMAL, MasterDialogStart, MasterDialogEnd)
	ns.StartDialog(obj19, ns.GetHost())
	ivar51 = 0
	for {
		if !(ivar51 < 4) {
			goto LABEL11
		}
		ns.LookAtObject(obj52[ivar51], ns.GetHost())
		ivar51 += 1
	}
LABEL11:
	ivar51 = 0
	for {
		if !(ivar51 < 4) {
			goto LABEL13
		}
		ns.LookAtObject(obj54[ivar51], ns.GetHost())
		ivar51 += 1
	}
LABEL13:
	ns.LookAtObject(obj19, ns.GetHost())
	goto LABEL9
LABEL5:
	ns.WideScreen(false)
	gvar12 = 5
	ns.WallGroupClose(gvar23)
	ivar15 = 30
	ns.FrameTimer(ivar15, StartDemonTalk)
	goto LABEL9
LABEL6:
	ns.Frozen(ns.GetHost(), false)
	ivar15 = 1
	ns.AudioEvent("FireBallCast", wp18)
	ns.AudioEvent(ns.FireGrate, wp18)
	CreateFireWall()
	gvar12 = 6
	ns.FrameTimer(ivar15, StartDemonTalk)
	goto LABEL9
LABEL7:
	ns.ClearOwner(obj19)
	ivar51 = 0
	for {
		if !(ivar51 < 4) {
			goto LABEL15
		}
		ns.ClearOwner(obj52[ivar51])
		ivar51 += 1
	}
LABEL15:
	ivar51 = 0
	for {
		if !(ivar51 < 4) {
			goto LABEL17
		}
		ns.ClearOwner(obj54[ivar51])
		ivar51 += 1
	}
LABEL17:
	gvar12 = 7
	ivar15 = 1
	ns.FrameTimer(ivar15, StartDemonTalk)
	goto LABEL9
LABEL8:
	PlayActionMusic()
	ns.ObjectOff(obj33)
	gvar12 = 8
	ns.AggressionLevel(obj19, 0.83)
	ns.Move(obj41, wp47)
	ns.Move(obj42, wp48)
	ns.Move(obj43, wp49)
	ns.Move(obj44, wp50)
	ns.AggressionLevel(obj41, 0.83)
	ns.AggressionLevel(obj42, 0.83)
	ns.AggressionLevel(obj43, 0.83)
	ns.AggressionLevel(obj44, 0.83)
	goto LABEL9
LABEL9:
	return
}
func DemonDeath() {
	ns.AudioEvent(ns.FlagDrop, wp18)
	PlaySubMusic()
	DestroyFireWall()
}
func ChapterEnd() {
	ns.JournalEdit(ns.GetHost(), "Wiz07EscapeUnderworld", 4)
	ns.Blind()
	ns.Frozen(ns.GetHost(), true)
	ns.FrameTimer(60, ChapterEndB)
	ns.ObjectOff(obj60)
}
func RegenA() {
	if ns.Object("RegenD1") != 0 {
		goto LABEL1
	}
	obj27 = ns.CreateObject("EmberDemon", wp30)
	ns.CreatureHunt(obj27)
LABEL1:
	if ns.Object("RegenD2") != 0 {
		goto LABEL2
	}
	obj28 = ns.CreateObject("EmberDemon", wp31)
	ns.CreatureHunt(obj28)
LABEL2:
	if ns.Object("RegenD3") != 0 {
		goto LABEL3
	}
	obj29 = ns.CreateObject("EmberDemon", wp32)
	ns.CreatureHunt(obj29)
LABEL3:
	return
}
func Quake() {
	var (
		v0 float32
		v1 float32
		v6 int
	)
	if flag4 {
		goto LABEL1
	}
	return
LABEL1:
	v0 = ns.GetObjectX(ns.GetHost())
	v1 = ns.GetObjectY(ns.GetHost())
	fvar76 = ns.RandomFloat(fvar74, fvar75)
	ns.MoveWaypoint(wp65, v0, v1)
	ns.AudioEvent(ns.BoulderMove, wp65)
	ns.AudioEvent(ns.EarthRumbleMinor, wp65)
	ns.Effect(ns.JIGGLE, v0, v1, fvar76, 0)
	v6 = ns.Random(200, ivar78)
	ns.FrameTimer(v6, Quake)
}
func LockDoors() {
	ns.LockDoor(obj70)
	ns.LockDoor(obj22)
	ns.LockDoor(obj71)
	ns.LockDoor(obj72)
}
func MapInitialize() {
	var v0 int
	wp61 = ns.Waypoint("Secret01WP")
	wp62 = ns.Waypoint("Secret02WP")
	gvar63 = ns.ObjectGroup("ImpGroup")
	obj60 = ns.Object("ChapterEndTrig")
	obj64 = ns.Object("SpikeDoor1")
	ns.LockDoor(obj64)
	gvar45 = ns.ObjectGroup("DemonGroup1")
	gvar46 = ns.ObjectGroup("DemonGroup2")
	obj41 = ns.Object("LineDR1")
	obj42 = ns.Object("LineDL2")
	obj43 = ns.Object("LineDR3")
	obj44 = ns.Object("LineDL4")
	wp18 = ns.Waypoint("FireSoundWP")
	wp47 = ns.Waypoint("Demon1WP")
	wp48 = ns.Waypoint("Demon2WP")
	wp49 = ns.Waypoint("Demon3WP")
	wp50 = ns.Waypoint("Demon4WP")
	wp59 = ns.Waypoint("ExitWP")
	wp65 = ns.Waypoint("QuakeAudioOrigin")
	wp66 = ns.Waypoint("CreationWP")
	obj24 = ns.Object("Light1")
	obj25 = ns.Object("Light2")
	obj26 = ns.Object("Light3")
	obj27 = ns.Object("RegenD1")
	obj28 = ns.Object("RegenD2")
	obj29 = ns.Object("RegenD3")
	wp30 = ns.Waypoint("RegenLoc1")
	wp31 = ns.Waypoint("RegenLoc2")
	wp32 = ns.Waypoint("RegenLoc3")
	obj20 = ns.Object("NastyPit1")
	obj21 = ns.Object("NastyPit2")
	obj33 = ns.Object("EndTeleport1")
	wp17[0] = ns.Waypoint("FireWall1")
	wp17[1] = ns.Waypoint("FireWall2")
	wp17[2] = ns.Waypoint("FireWall3")
	wp17[3] = ns.Waypoint("FireWall4")
	wp17[4] = ns.Waypoint("FireWall5")
	wp17[5] = ns.Waypoint("FireWall6")
	wp17[6] = ns.Waypoint("FireWall7")
	wp17[7] = ns.Waypoint("FireWall8")
	wp17[8] = ns.Waypoint("FireWall9")
	wp17[9] = ns.Waypoint("FireWall10")
	wp17[10] = ns.Waypoint("FireWall11")
	wp17[11] = ns.Waypoint("FireWall12")
	wp17[12] = ns.Waypoint("FireWall13")
	wp17[13] = ns.Waypoint("FireWall14")
	wp17[14] = ns.Waypoint("FireWall15")
	wp17[15] = ns.Waypoint("FireWall16")
	wp17[16] = ns.Waypoint("FireWall17")
	wp17[17] = ns.Waypoint("FireWall18")
	wp17[18] = ns.Waypoint("FireWall19")
	wp17[19] = ns.Waypoint("FireWall20")
	gvar23 = ns.WallGroup("Walls20")
	obj70 = ns.Object("BarredDoor1")
	obj22 = ns.Object("BarredDoor2")
	obj71 = ns.Object("BarredDoor3")
	obj72 = ns.Object("BarredDoor4")
	obj73 = ns.Object("UnlockDoor1")
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL1
		}
		obj52[v0] = ns.Object("LineDR" + strconv.Itoa(2*v0+1))
		v0 += 1
	}
LABEL1:
	ivar53[0] = 1
	ivar53[1] = 3
	ivar53[2] = 5
	ivar53[3] = 7
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL3
		}
		obj54[v0] = ns.Object("LineDL" + strconv.Itoa(2*(v0+1)))
		v0 += 1
	}
LABEL3:
	ivar55[0] = 2
	ivar55[1] = 4
	ivar55[2] = 6
	ivar55[3] = 8
	obj19 = ns.Object("C6Necromancer")
	wp34 = ns.Waypoint("A")
	wp35 = ns.Waypoint("BL")
	wp36 = ns.Waypoint("BR")
	wp37 = ns.Waypoint("CL")
	wp38 = ns.Waypoint("CR")
	wp39 = ns.Waypoint("DL")
	wp40 = ns.Waypoint("DR")
	wp67[0] = ns.Waypoint("Flamer1")
	wp67[1] = ns.Waypoint("Flamer2")
	wp67[2] = ns.Waypoint("Flamer3")
	wp67[3] = ns.Waypoint("Flamer4")
	wp67[4] = ns.Waypoint("Flamer5")
	wp67[5] = ns.Waypoint("Flamer6")
	wp67[6] = ns.Waypoint("Flamer7")
	wp67[7] = ns.Waypoint("Flamer8")
	wp67[8] = ns.Waypoint("Flamer9")
	wp68[0] = ns.Waypoint("FlamerB1")
	wp68[1] = ns.Waypoint("FlamerB2")
	wp68[2] = ns.Waypoint("FlamerB3")
	wp68[3] = ns.Waypoint("FlamerB4")
	wp68[4] = ns.Waypoint("FlamerB5")
	wp68[5] = ns.Waypoint("FlamerB6")
	wp68[6] = ns.Waypoint("FlamerB7")
	wp68[7] = ns.Waypoint("FlamerB8")
	wp68[8] = ns.Waypoint("FlamerB9")
	wp69[0] = ns.Waypoint("FlamerC1")
	wp69[1] = ns.Waypoint("FlamerC2")
	wp69[2] = ns.Waypoint("FlamerC3")
	wp69[3] = ns.Waypoint("FlamerC4")
	wp69[4] = ns.Waypoint("FlamerC5")
	wp69[5] = ns.Waypoint("FlamerC6")
	wp69[6] = ns.Waypoint("FlamerC7")
	wp69[7] = ns.Waypoint("FlamerC8")
	wp69[8] = ns.Waypoint("FlamerC9")
	gvar58 = ns.ObjectGroup("SecretT1")
	Quake()
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL5
		}
		ns.SetOwner(ns.GetHost(), obj52[v0])
		v0 += 1
	}
LABEL5:
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL7
		}
		ns.SetOwner(ns.GetHost(), obj54[v0])
		v0 += 1
	}
LABEL7:
	ns.SetOwner(ns.GetHost(), obj19)
	ns.FrameTimer(2, LockDoors)
}
func PlayerDeath() {
	ns.DeathScreen(7)
}
func MapEntry() {
	ns.ObjectOn(obj60)
	PlaySubMusic()
}
func ImpGroupHostilize() {
	ns.GroupAggressionLevel(gvar63, 0.83)
}
func MonsterGoHome() {
	if !ns.HasClass(ns.GetCaller(), "ENEMY") {
		goto LABEL1
	}
	ns.GoBackHome(ns.GetCaller())
LABEL1:
	return
}
func PlayWanderMusic() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.Music(17, 100)
LABEL1:
	return
}
func Secret01() {
	ns.AudioEvent(ns.SecretFound, wp61)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 200)
}
func Secret02() {
	ns.AudioEvent(ns.SecretFound, wp62)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 200)
}
func OnEvent(typ string) {
	switch typ {
	case "MapInitialize":
		MapInitialize()
	case "PlayerDeath":
		PlayerDeath()
	case "MapEntry":
		MapEntry()
	}
}
