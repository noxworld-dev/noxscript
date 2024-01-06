package war10c

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	obj4   [21]ns.ObjectID
	ivar5  int
	obj6   ns.ObjectID
	obj7   ns.ObjectID
	wp8    ns.WaypointID
	wp9    ns.WaypointID
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
	obj24  [6]ns.ObjectID
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
	gvar40 ns.ObjectGroupID
	gvar41 ns.ObjectGroupID
	obj42  ns.ObjectID
	wp43   [2]ns.WaypointID
	wp44   ns.WaypointID
	wp45   ns.WaypointID
	gvar46 ns.WallGroupID
	gvar47 ns.WallGroupID
	gvar48 ns.WallGroupID
	gvar49 ns.WallGroupID
	gvar50 ns.WallGroupID
	gvar51 ns.WallGroupID
	gvar52 ns.WallGroupID
	gvar53 ns.WallGroupID
	gvar54 ns.WallGroupID
	gvar55 ns.WallGroupID
	gvar56 ns.WallGroupID
	gvar57 ns.WallGroupID
	gvar58 ns.WallGroupID
	gvar59 ns.WallGroupID
	flag60 bool
	flag61 bool
	flag62 bool
	flag63 bool
	flag64 bool
	flag65 bool
	flag66 bool
	flag67 bool
	ivar68 int
	wp69   ns.WaypointID
	wp70   ns.WaypointID
	obj71  ns.ObjectID
	gvar72 ns.ObjectGroupID
	wp73   [5]ns.WaypointID
	wp74   ns.WaypointID
	obj75  [2]ns.ObjectID
	obj76  [2]ns.ObjectID
	obj77  ns.ObjectID
	wp78   ns.WaypointID
	wp79   ns.WaypointID
	wp80   ns.WaypointID
	fvar81 float32
	fvar82 float32
	flag83 bool
	flag84 bool
)

func init() {
	flag60 = false
	flag61 = false
	flag62 = false
	flag63 = false
	flag64 = false
	flag65 = false
	flag66 = false
	flag67 = false
	ivar68 = 0
	flag83 = false
	flag84 = false
}
func DisableAT(a1 int) {
	ns.ObjectOff(obj4[a1])
}
func FireAT0() {
	ns.ObjectOn(obj4[0])
	ns.FrameTimerWithArg(1, nil, DisableAT)
}
func FireAT1() {
	ns.ObjectOn(obj4[1])
	ns.FrameTimerWithArg(1, 1, DisableAT)
}
func FireAT2() {
	ns.ObjectOn(obj4[2])
	ns.FrameTimerWithArg(1, 2, DisableAT)
}
func FireAT3() {
	ns.ObjectOn(obj4[3])
	ns.FrameTimerWithArg(1, 3, DisableAT)
}
func FireAT4() {
	ns.ObjectOn(obj4[4])
	ns.FrameTimerWithArg(1, 4, DisableAT)
}
func FireAT5() {
	ns.ObjectOn(obj4[5])
	ns.FrameTimerWithArg(1, 5, DisableAT)
}
func FireAT6() {
	ns.ObjectOn(obj4[6])
	ns.FrameTimerWithArg(1, 6, DisableAT)
}
func FireAT7() {
	ns.ObjectOn(obj4[7])
	ns.FrameTimerWithArg(1, 7, DisableAT)
}
func FireAT8() {
	ns.ObjectOn(obj4[8])
	ns.FrameTimerWithArg(1, 8, DisableAT)
}
func FireAT9() {
	ns.ObjectOn(obj4[9])
	ns.FrameTimerWithArg(1, 9, DisableAT)
}
func FireAT10() {
	ns.ObjectOn(obj4[10])
	ns.FrameTimerWithArg(1, 10, DisableAT)
}
func FireAT11() {
	ns.ObjectOn(obj4[11])
	ns.FrameTimerWithArg(1, 11, DisableAT)
}
func FireAT17() {
	ns.ObjectOn(obj4[17])
	ns.FrameTimerWithArg(1, 17, DisableAT)
}
func FireAT18() {
	ns.ObjectOn(obj4[18])
	ns.FrameTimerWithArg(1, 18, DisableAT)
}
func FireAT19() {
	ns.ObjectOn(obj4[19])
	ns.FrameTimerWithArg(1, 19, DisableAT)
}
func FireAT20() {
	ns.ObjectOn(obj4[20])
	ns.FrameTimerWithArg(1, 20, DisableAT)
}
func FireATRow() {
	ns.ObjectOn(obj4[12])
	ns.ObjectOn(obj4[13])
	ns.ObjectOn(obj4[14])
	ns.ObjectOn(obj4[15])
	ns.ObjectOn(obj4[16])
	ivar5 = 12
	for {
		if !(ivar5 < 17) {
			goto LABEL1
		}
		ns.FrameTimerWithArg(1, ivar5, DisableAT)
		ivar5 += 1
	}
LABEL1:
	return
}
func PlayerDeath() {
	ns.DeathScreen(10)
}
func OrbLightOff() {
	ns.ObjectOff(obj6)
	ns.MoveObject(obj7, ns.GetWaypointX(wp8), ns.GetWaypointY(wp8))
}
func ForceOrb() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
	)
	v0 = ns.RandomFloat(-30, 30)
	v1 = ns.RandomFloat(-30, 30)
	v2 = ns.GetObjectX(ns.GetTrigger()) + v0
	v3 = ns.GetObjectY(ns.GetTrigger()) + v1
	ns.MoveObject(obj6, ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger()))
	ns.MoveObject(obj7, ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger())+6)
	ns.MoveWaypoint(wp9, ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger()))
	ns.ObjectOn(obj6)
	ns.AudioEvent(ns.BallBounce, wp9)
	ns.PushObject(ns.GetCaller(), 60, v2, v3)
	ns.FrameTimer(4, OrbLightOff)
}
func StairAccess2() {
	if !(flag62 == true && flag63 == true) {
		goto LABEL1
	}
	ns.UnlockDoor(obj12)
	ns.UnlockDoor(obj13)
LABEL1:
	return
}
func CreateCherub() {
	var v0 int
	if !(ivar68 < 6) {
		goto LABEL1
	}
	v0 = ns.Random(0, 1)
	obj24[ivar68] = ns.CreateObject("EvilCherub", wp43[v0])
	ns.Effect(ns.BLUE_SPARKS, ns.GetWaypointX(wp43[v0]), ns.GetWaypointY(wp43[v0]), 0, 0)
	ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp43[v0]), ns.GetWaypointY(wp43[v0]), 0, 0)
	ns.SetCallback(obj24[ivar68], 5, CreateCherub)
	ns.SetRoamFlag(obj24[ivar68], 128)
	ns.AggressionLevel(obj24[ivar68], 0.83)
	ivar68 += 1
	goto LABEL2
LABEL1:
	return
LABEL2:
	return
}
func WakeVKBat2() {
	ns.ObjectOn(obj76[1])
}
func PlayAction2Music() {
	ns.Music(27, 100)
}
func CherubHostilize() {
	ns.GroupAggressionLevel(gvar41, 0.83)
}
func OpenSecretRoom() {
	ns.WallGroupOpen(gvar49)
}
func EnableTreasureLich1() {
	ns.ObjectOn(obj14)
}
func AccessTreasureRoom1() {
	ns.WallGroupOpen(gvar46)
}
func EnableTreasureLich2() {
	ns.ObjectOn(obj15)
}
func AccessTreasureRoom2() {
	ns.WallGroupOpen(gvar47)
}
func EnableTreasureMech() {
	ns.ObjectOn(obj16)
}
func AccessTreasureRoom3() {
	ns.WallGroupOpen(gvar48)
}
func EnableRoom() {
	if !(ns.IsCaller(ns.GetHost()) && flag60 == false) {
		goto LABEL1
	}
	flag60 = true
	ns.ObjectOn(obj27)
	ns.LockDoor(obj10)
	ns.LockDoor(obj11)
LABEL1:
	return
}
func StairAccess1() {
	ns.WallGroupOpen(gvar50)
	ns.UnlockDoor(obj10)
	ns.UnlockDoor(obj11)
}
func AllowAccess2() {
	ns.PrintToAll("War10c:SentryDisabled")
	ns.ObjectOff(obj21)
	ns.ObjectOff(obj20)
	ns.ObjectOff(obj19)
	ns.AudioEvent(ns.Gear2, wp69)
}
func EnableRoom2() {
	PlayAction2Music()
	if !(ns.IsCaller(ns.GetHost()) && flag61 == false) {
		goto LABEL1
	}
	flag61 = true
	ns.LockDoor(obj12)
	ns.LockDoor(obj13)
	ns.FrameTimer(800, CreateCherub)
LABEL1:
	return
}
func OneDown() {
	flag62 = true
	ns.FrameTimer(1, StairAccess2)
}
func TwoDown() {
	flag63 = true
	ns.FrameTimer(1, StairAccess2)
}
func SwitchOffSentries() {
	ns.ObjectOff(obj17)
	ns.ObjectOff(obj18)
	ns.UnlockDoor(obj12)
	ns.UnlockDoor(obj13)
}
func SwitchOnSentries() {
	ns.ObjectOn(obj17)
	ns.ObjectOn(obj18)
}
func UnlockAccessDoors() {
	ns.ObjectOff(obj31)
	ns.UnlockDoor(obj25)
	ns.UnlockDoor(obj26)
}
func OpenWall01() {
	ns.WallGroupOpen(gvar56)
	ns.AggressionLevel(obj42, 0.83)
}
func OpenWindowWall1() {
	ns.WallGroupOpen(gvar57)
}
func CloseWindowWall1() {
	ns.WallGroupClose(gvar57)
}
func OpenWindowWall2() {
	ns.WallGroupOpen(gvar58)
}
func CloseWindowWall2() {
	ns.WallGroupClose(gvar58)
}
func OpenCage1() {
	ns.WallGroupOpen(gvar51)
	ns.ObjectOn(obj28)
}
func OpenCage2() {
	ns.WallGroupOpen(gvar52)
	ns.ObjectOn(obj30)
}
func OpenCage3() {
	ns.WallGroupOpen(gvar53)
	ns.ObjectOn(obj29)
}
func OpenCage4() {
	ns.WallGroupOpen(gvar54)
	WakeVKBat2()
}
func OpenAccess() {
	ns.WallGroupOpen(gvar55)
}
func CheckIdols() {
	if !(flag64 == true && flag65 == true) {
		goto LABEL1
	}
	ns.ObjectGroupOn(gvar40)
LABEL1:
	return
}
func BlueIdol() {
	if !ns.IsCaller(obj32) {
		goto LABEL1
	}
	flag64 = true
	CheckIdols()
	ns.ObjectOff(obj38)
	ns.ObjectOn(obj33)
	ns.Delete(obj32)
	ns.Delete(obj36)
	ns.AudioEvent(ns.Gear2, wp70)
LABEL1:
	return
}
func RedIdol() {
	if !ns.IsCaller(obj34) {
		goto LABEL1
	}
	flag65 = true
	CheckIdols()
	ns.ObjectOff(obj39)
	ns.ObjectOn(obj35)
	ns.Delete(obj34)
	ns.Delete(obj37)
	ns.AudioEvent(ns.Gear2, wp70)
LABEL1:
	return
}
func CheckSwitches() {
	if !(flag66 == true && flag67 == true) {
		goto LABEL1
	}
	ns.WallGroupOpen(gvar59)
LABEL1:
	ns.AudioEvent("CreatureCageAppear", wp45)
}
func ExitWallSwitch1() {
	ns.AudioEvent(ns.Gear3, wp45)
	flag66 = true
	CheckSwitches()
}
func ExitWallSwitch2() {
	ns.AudioEvent(ns.Gear3, wp45)
	flag67 = true
	CheckSwitches()
}
func InitializeVampireKnights() {
	obj75[0] = ns.Object("VampireKnight1")
	obj75[1] = ns.Object("VampireKnight2")
	obj76[0] = ns.Object("VKBat1")
	obj76[1] = ns.Object("VKBat2")
	wp78 = ns.Waypoint("VKSpot")
	wp79 = ns.Waypoint("VampireKnight1WP")
	wp80 = ns.Waypoint("BatCreate")
}
func PlaySubMusic() {
	ns.Music(18, 100)
}
func MapInitialize() {
	wp70 = ns.Waypoint("LichKeyAudioWP")
	wp69 = ns.Waypoint("SentryAudioWP")
	gvar41 = ns.ObjectGroup("CherubPitGroup")
	obj6 = ns.Object("OrbLight")
	obj7 = ns.Object("OrbEffect")
	obj10 = ns.Object("ForceOrbDoorA")
	obj11 = ns.Object("ForceOrbDoorB")
	obj12 = ns.Object("ForceOrbDoorC")
	obj13 = ns.Object("ForceOrbDoorD")
	obj4[0] = ns.Object("AT0")
	obj4[1] = ns.Object("AT1")
	obj4[2] = ns.Object("AT2")
	obj4[3] = ns.Object("AT3")
	obj4[4] = ns.Object("AT4")
	obj4[5] = ns.Object("AT5")
	obj4[6] = ns.Object("AT6")
	obj4[7] = ns.Object("AT7")
	obj4[8] = ns.Object("AT8")
	obj4[9] = ns.Object("AT9")
	obj4[10] = ns.Object("AT10")
	obj4[11] = ns.Object("AT11")
	obj4[12] = ns.Object("AT12")
	obj4[13] = ns.Object("AT13")
	obj4[14] = ns.Object("AT14")
	obj4[15] = ns.Object("AT15")
	obj4[16] = ns.Object("AT16")
	obj4[17] = ns.Object("AT17")
	obj4[18] = ns.Object("AT18")
	obj4[19] = ns.Object("AT19")
	obj4[20] = ns.Object("AT20")
	obj14 = ns.Object("TreasureLich1")
	obj15 = ns.Object("TreasureLich2")
	obj16 = ns.Object("TreasureMech")
	obj27 = ns.Object("Necromancer1")
	obj28 = ns.Object("Necromancer2")
	obj29 = ns.Object("Necromancer3")
	obj42 = ns.Object("Lich1")
	obj71 = ns.Object("Lich2")
	obj30 = ns.Object("Lich3")
	obj17 = ns.Object("Sentry1")
	obj18 = ns.Object("Sentry2")
	obj19 = ns.Object("Sentry3")
	obj20 = ns.Object("Sentry4")
	obj21 = ns.Object("SentrySwitch")
	obj22 = ns.Object("MainSentry1")
	obj23 = ns.Object("MainSentry2")
	obj25 = ns.Object("AccessDoorA")
	obj26 = ns.Object("AccessDoorB")
	obj31 = ns.Object("OpenDoors")
	obj32 = ns.Object("BlueIdol")
	obj34 = ns.Object("RedIdol")
	obj33 = ns.Object("BlueIdolBase")
	obj35 = ns.Object("RedIdolBase")
	obj36 = ns.Object("BlueTrigger")
	obj37 = ns.Object("RedTrigger")
	obj38 = ns.Object("BlueIdolSentry")
	obj39 = ns.Object("RedIdolSentry")
	wp73[0] = ns.Waypoint("N1TeleWP0")
	wp73[1] = ns.Waypoint("N1TeleWP1")
	wp73[2] = ns.Waypoint("N1TeleWP2")
	wp73[3] = ns.Waypoint("N1TeleWP3")
	wp73[4] = ns.Waypoint("N1TeleWP4")
	wp9 = ns.Waypoint("OrbSoundWP")
	wp8 = ns.Waypoint("OrbEffectWP")
	wp43[0] = ns.Waypoint("CherubCreate1")
	wp43[1] = ns.Waypoint("CherubCreate2")
	gvar72 = ns.ObjectGroup("VampTriggers")
	gvar40 = ns.ObjectGroup("FinalLiches")
	gvar46 = ns.WallGroup("TreasureRoom1Walls")
	gvar47 = ns.WallGroup("TreasureRoom2Walls")
	gvar48 = ns.WallGroup("TreasureRoom3Walls")
	wp44 = ns.Waypoint("MechGolemAudioWP")
	gvar49 = ns.WallGroup("SecretRoomWall")
	gvar50 = ns.WallGroup("ForceOrbRoomWalls1")
	gvar51 = ns.WallGroup("Cage1")
	gvar52 = ns.WallGroup("Cage2")
	gvar53 = ns.WallGroup("Cage3")
	gvar54 = ns.WallGroup("Cage4")
	gvar55 = ns.WallGroup("AccessTreasureRoom3")
	gvar56 = ns.WallGroup("OpeningWall01")
	gvar57 = ns.WallGroup("WindowWall1")
	gvar58 = ns.WallGroup("WindowWall2")
	gvar59 = ns.WallGroup("LichStairAccess")
	wp45 = ns.Waypoint("ExitAudioWP")
	wp74 = ns.Waypoint("Secret1AudioWP")
	InitializeVampireKnights()
	PlaySubMusic()
	ns.LockDoor(obj25)
	ns.LockDoor(obj26)
}
func Secret100XP() {
	ns.GiveXp(ns.GetHost(), 100)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.AudioEvent(ns.SecretFound, wp74)
}
func PlayWanderMusic() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.Music(17, 100)
LABEL1:
	return
}
func Necro1TeleportInjured() {
	var (
		v0 float32
		v1 float32
		v2 int
	)
	v2 = ns.Random(0, 4)
	v0 = ns.GetWaypointX(wp73[v2])
	v1 = ns.GetWaypointY(wp73[v2])
	ns.PauseObject(ns.GetTrigger(), 15)
	ns.Enchant(ns.GetTrigger(), ns.ENCHANT_INVISIBLE, 2)
	ns.Effect(ns.BLUE_SPARKS, ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger()), 0, 0)
	ns.Effect(ns.SMOKE_BLAST, ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger()), 0, 0)
	ns.MoveObject(ns.GetTrigger(), v0, v1)
	ns.Effect(ns.BLUE_SPARKS, v0, v1, 0, 0)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
}
func SmackPlayer() {
	r0 := ns.IsAttackedBy(ns.GetCaller(), ns.GetTrigger())
	if !r0 {
		goto LABEL1
	}
	ns.LookAtObject(ns.GetTrigger(), ns.GetCaller())
	ns.HitLocation(ns.GetTrigger(), ns.GetObjectX(ns.GetCaller()), ns.GetObjectY(ns.GetCaller()))
LABEL1:
	return
}
func HuntPlayer() {
	ns.CreatureHunt(ns.GetTrigger())
}
func BatToVampireKnight() {
	if flag83 {
		goto LABEL1
	}
	fvar81 = ns.GetObjectX(ns.GetTrigger())
	fvar82 = ns.GetObjectY(ns.GetTrigger())
	ns.Delete(ns.GetTrigger())
	ns.ObjectOn(obj75[0])
LABEL1:
	ns.Effect(ns.BLUE_SPARKS, fvar81, fvar82, 0, 0)
	ns.Effect(ns.SMOKE_BLAST, fvar81, fvar82, 0, 0)
	ns.Enchant(obj75[0], ns.ENCHANT_INVISIBLE, 0.25)
	ns.MoveObject(obj75[0], fvar81, fvar82)
	ns.LookAtObject(obj75[0], ns.GetHost())
	flag83 = false
}
func BatToVampireKnight2() {
	if flag83 {
		goto LABEL1
	}
	fvar81 = ns.GetObjectX(ns.GetTrigger())
	fvar82 = ns.GetObjectY(ns.GetTrigger())
	ns.Delete(ns.GetTrigger())
	ns.ObjectOn(obj75[1])
LABEL1:
	ns.Effect(ns.BLUE_SPARKS, fvar81, fvar82, 0, 0)
	ns.Effect(ns.SMOKE_BLAST, fvar81, fvar82, 0, 0)
	ns.Enchant(obj75[1], ns.ENCHANT_INVISIBLE, 0.25)
	ns.MoveObject(obj75[1], fvar81, fvar82)
	ns.LookAtObject(obj75[1], ns.GetHost())
}
func SetRetreatBat() {
	ns.AggressionLevel(obj77, 0.83)
	ns.SetCallback(obj77, 7, BatDie)
	ns.SetCallback(obj77, 5, BatDie)
}
func BatDie() {
	var (
		v0 ns.ObjectID
		v1 ns.ObjectID
		v2 ns.ObjectID
	)
	if !(ns.CurrentHealth(ns.GetTrigger()) <= 0) {
		goto LABEL1
	}
	ns.Effect(ns.SMOKE_BLAST, ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger()), 0, 0)
	ns.MoveWaypoint(wp80, ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger()))
	ns.MoveObject(ns.GetTrigger(), ns.GetWaypointX(wp79), ns.GetWaypointY(wp79))
	ns.Delete(obj77)
	ns.AudioEvent(ns.BurnCast, wp80)
	v0 = ns.CreateObject("Flame", wp80)
	v1 = ns.CreateObject("MediumFlame", wp80)
	v2 = ns.CreateObject("SmallFlame", wp80)
	ns.DeleteObjectTimer(v0, 80)
	ns.DeleteObjectTimer(v1, 83)
	ns.DeleteObjectTimer(v2, 85)
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.GiveXp(ns.GetHost(), 250)
LABEL1:
	return
}
func ChangeOnSight() {
	if !flag84 {
		goto LABEL1
	}
	ns.AggressionLevel(ns.GetTrigger(), 0.83)
	ns.SetCallback(ns.GetTrigger(), 6, BatToVampireKnight)
LABEL1:
	return
}
func InjureVampireKnight() {
	if flag83 {
		goto LABEL1
	}
	fvar81 = ns.GetObjectX(ns.GetTrigger())
	fvar82 = ns.GetObjectY(ns.GetTrigger())
	ns.ObjectOn(obj75[0])
	ns.Damage(obj75[0], 0, 50, 0)
	ns.Effect(ns.SMOKE_BLAST, ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger()), 0, 0)
	ns.Delete(ns.GetTrigger())
	flag83 = true
	ns.FrameTimer(1, BatToVampireKnight)
LABEL1:
	return
}
func InjureVampireKnight2() {
	fvar81 = ns.GetObjectX(ns.GetTrigger())
	fvar82 = ns.GetObjectY(ns.GetTrigger())
	ns.ObjectOn(obj75[1])
	ns.Damage(obj75[1], 0, 50, 0)
	ns.Effect(ns.SMOKE_BLAST, ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger()), 0, 0)
	ns.Delete(ns.GetTrigger())
	flag83 = true
	ns.FrameTimer(1, BatToVampireKnight2)
}
func BatMove() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.SetCallback(ns.GetTrigger(), 4, ChangeOnSight)
	ns.PauseObject(ns.GetTrigger(), 15)
LABEL1:
	return
}
func VKDie() {
	fvar81 = ns.GetObjectX(ns.GetTrigger())
	fvar82 = ns.GetObjectY(ns.GetTrigger())
	ns.Delete(ns.GetTrigger())
	ns.MoveWaypoint(wp80, fvar81, fvar82)
	obj77 = ns.CreateObject("Bat", wp80)
	ns.FrameTimer(1, SetRetreatBat)
}
func OnEvent(typ string) {
	switch typ {
	case "PlayerDeath":
		PlayerDeath()
	case "MapInitialize":
		MapInitialize()
	}
}
