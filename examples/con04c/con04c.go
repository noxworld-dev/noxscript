package con04c

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	gvar4  ns.WallGroupID
	obj5   ns.ObjectID
	obj6   ns.ObjectID
	ivar7  int
	ivar8  int
	fvar9  float32
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
	wp26   ns.WaypointID
	wp27   ns.WaypointID
	wp28   ns.WaypointID
	wp29   ns.WaypointID
	wp30   ns.WaypointID
	wp31   ns.WaypointID
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
	gvar42 ns.ObjectGroupID
	gvar43 ns.ObjectGroupID
	gvar44 ns.ObjectGroupID
	obj45  [13]ns.ObjectID
	obj46  [13]ns.ObjectID
	ivar47 int
	wp48   ns.WaypointID
	obj49  ns.ObjectID
	obj50  ns.ObjectID
	obj51  ns.ObjectID
	obj52  ns.ObjectID
	obj53  ns.ObjectID
	obj54  ns.ObjectID
	wp55   ns.WaypointID
	wp56   ns.WaypointID
	gvar57 ns.WallGroupID
	gvar58 ns.WallGroupID
	flag59 bool
	flag60 bool
	gvar61 int
	gvar62 int
	gvar63 int
	obj64  ns.ObjectID
	obj65  ns.ObjectID
	obj66  ns.ObjectID
	obj67  ns.ObjectID
	obj68  ns.ObjectID
	obj69  ns.ObjectID
	gvar70 ns.ObjectGroupID
	wp71   [4]ns.WaypointID
	wp72   ns.WaypointID
	wp73   ns.WaypointID
	flag74 bool
	flag75 bool
	ivar76 int
	wp77   ns.WaypointID
	gvar78 ns.ObjectGroupID
	ivar79 int
	ivar80 int
	ivar81 int
	ivar82 int
	ivar83 int
	ivar84 int
	ivar85 int
	ivar86 int
	ivar87 int
	ivar88 int
	ivar89 int
	gvar90 int
	gvar91 int
	obj92  ns.ObjectID
	obj93  ns.ObjectID
	obj94  ns.ObjectID
	obj95  ns.ObjectID
	obj96  ns.ObjectID
	obj97  ns.ObjectID
	obj98  ns.ObjectID
	obj99  ns.ObjectID
	obj100 ns.ObjectID
	obj101 ns.ObjectID
	obj102 ns.ObjectID
)

func init() {
	ivar7 = 50
	ivar8 = 20
	fvar9 = 5
	ivar47 = 120
	flag59 = false
	flag60 = false
	gvar61 = 0
	gvar62 = 1
	gvar63 = gvar62
	flag74 = false
	flag75 = false
	ivar76 = 0
	ivar79 = 0
	ivar80 = 0
	ivar81 = 0
	ivar82 = 0
	ivar83 = 0
	ivar84 = 0
	ivar85 = 0
	ivar86 = 0
	ivar87 = 0
	ivar88 = 0
	ivar89 = 0
	gvar90 = 0
	gvar91 = 0
}
func InitializeSkeletonAmbush() {
	gvar4 = ns.WallGroup("SkeletonAmbushWalls")
}
func SkeletonAmbush() {
	ns.WallGroupBreak(gvar4)
}
func SkeletonAmbushTimer() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.FrameTimer(45, SkeletonAmbush)
LABEL1:
	return
}
func EnablePitTrapElevator01() {
	ns.ObjectOff(ns.GetTrigger())
	ns.ObjectOn(ns.Object("PitTrapElevator01"))
}
func EnablePitTrapElevator02() {
	ns.ObjectOff(ns.GetTrigger())
	ns.ObjectOn(ns.Object("PitTrapElevator02"))
}
func OpenPitElevatorWalls01() {
	ns.ObjectOff(ns.GetTrigger())
	ns.WallGroupOpen(ns.WallGroup("PitTrapElevatorWalls01"))
}
func OpenPitElevatorWalls02() {
	ns.ObjectOff(ns.GetTrigger())
	ns.WallGroupOpen(ns.WallGroup("PitTrapElevatorWalls02"))
}
func ZombieAmbush() {
	ns.ObjectOff(ns.GetTrigger())
	ns.WallGroupBreak(ns.WallGroup("ZombieAmbushWalls"))
}
func InitializeHallwayArrowTraps() {
	obj5 = ns.Object("HallwayArrowTrap01")
	obj6 = ns.Object("HallwayArrowTrap02")
}
func ResetHallwayArrowTrap01() {
	ns.ObjectOff(obj5)
}
func ResetHallwayArrowTrap02() {
	ns.ObjectOff(obj6)
}
func ActivateHallwayArrowTrap01() {
	ns.ObjectOn(obj5)
	ns.FrameTimer(1, ResetHallwayArrowTrap01)
}
func ActivateHallwayArrowTrap02() {
	ns.ObjectOn(obj6)
	ns.FrameTimer(1, ResetHallwayArrowTrap02)
}
func InitializeBlockTrap() {
	obj10 = ns.Object("SpikeBlock01")
	obj11 = ns.Object("SpikeBlock02")
	obj12 = ns.Object("SpikeBlock03")
	obj13 = ns.Object("SpikeBlock04")
	obj14 = ns.Object("SpikeBlock05")
	obj15 = ns.Object("SpikeBlock06")
	obj16 = ns.Object("SpikeBlock07")
	obj17 = ns.Object("SpikeBlock08")
	obj18 = ns.Object("SpikeBlock09")
	obj19 = ns.Object("SpikeBlock10")
	obj20 = ns.Object("SpikeBlock11")
	obj21 = ns.Object("SpikeBlock12")
	obj22 = ns.Object("SpikeBlock13")
	obj23 = ns.Object("SpikeBlock14")
	obj24 = ns.Object("SpikeBlock15")
	obj25 = ns.Object("SpikeBlock16")
	wp26 = ns.Waypoint("BlockWP01")
	wp27 = ns.Waypoint("BlockWP02")
	wp28 = ns.Waypoint("BlockWP03")
	wp29 = ns.Waypoint("BlockWP04")
	wp30 = ns.Waypoint("BlockWP05")
	wp31 = ns.Waypoint("BlockWP06")
	wp32 = ns.Waypoint("BlockWP07")
	wp33 = ns.Waypoint("BlockWP08")
	wp34 = ns.Waypoint("BlockWP09")
	wp35 = ns.Waypoint("BlockWP10")
	wp36 = ns.Waypoint("BlockWP11")
	wp37 = ns.Waypoint("BlockWP12")
	wp38 = ns.Waypoint("BlockWP13")
	wp39 = ns.Waypoint("BlockWP14")
	wp40 = ns.Waypoint("BlockWP15")
	wp41 = ns.Waypoint("BlockWP16")
	gvar42 = ns.ObjectGroup("BlockTrap01Triggers")
	gvar43 = ns.ObjectGroup("BlockTrap02Triggers")
	gvar44 = ns.ObjectGroup("BlockTrap03Triggers")
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
	if v0 == 3 {
		goto LABEL3
	}
	goto LABEL4
LABEL1:
	ns.AudioEvent(ns.HammerMissing, wp28)
	ns.Effect(ns.JIGGLE, ns.GetWaypointX(wp28), ns.GetWaypointY(wp28), fvar9, 0)
	goto LABEL4
LABEL2:
	ns.AudioEvent(ns.HammerMissing, wp34)
	ns.Effect(ns.JIGGLE, ns.GetWaypointX(wp34), ns.GetWaypointY(wp34), fvar9, 0)
	goto LABEL4
LABEL3:
	ns.AudioEvent(ns.HammerMissing, wp38)
	ns.Effect(ns.JIGGLE, ns.GetWaypointX(wp38), ns.GetWaypointY(wp38), fvar9, 0)
	goto LABEL4
LABEL4:
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
	if v0 == 3 {
		goto LABEL3
	}
	goto LABEL4
LABEL1:
	ns.ObjectGroupOn(gvar42)
	ns.AudioEvent(ns.TriggerReleased, wp28)
	goto LABEL4
LABEL2:
	ns.ObjectGroupOn(gvar43)
	ns.AudioEvent(ns.TriggerReleased, wp34)
	goto LABEL4
LABEL3:
	ns.ObjectGroupOn(gvar44)
	ns.AudioEvent(ns.TriggerReleased, wp38)
	goto LABEL4
LABEL4:
	return
}
func CloseBlockTrap01() {
	ns.ObjectGroupOff(gvar42)
	ns.Move(obj10, wp26)
	ns.Move(obj11, wp27)
	ns.Move(obj12, wp28)
	ns.Move(obj13, wp29)
	ns.Move(obj14, wp30)
	ns.Move(obj15, wp31)
	ns.FrameTimerWithArg(ivar8, 1, BlockTrapBoom)
	ns.FrameTimerWithArg(ivar7, 1, ResetBlockTrap)
}
func CloseBlockTrap02() {
	ns.ObjectGroupOff(gvar43)
	ns.Move(obj16, wp32)
	ns.Move(obj17, wp33)
	ns.Move(obj18, wp34)
	ns.Move(obj19, wp35)
	ns.Move(obj20, wp36)
	ns.Move(obj21, wp37)
	ns.FrameTimerWithArg(ivar8, 2, BlockTrapBoom)
	ns.FrameTimerWithArg(ivar7, 2, ResetBlockTrap)
}
func CloseBlockTrap03() {
	ns.ObjectGroupOff(gvar44)
	ns.Move(obj22, wp38)
	ns.Move(obj23, wp39)
	ns.Move(obj24, wp40)
	ns.Move(obj25, wp41)
	ns.FrameTimerWithArg(ivar8, 3, BlockTrapBoom)
	ns.FrameTimerWithArg(ivar7, 3, ResetBlockTrap)
}
func InitializeFistTraps() {
	obj46[0] = ns.Object("FistTrap01")
	obj46[1] = ns.Object("FistTrap02")
	obj46[2] = ns.Object("FistTrap03")
	obj46[3] = ns.Object("FistTrap04")
	obj46[4] = ns.Object("FistTrap05")
	obj46[5] = ns.Object("FistTrap06")
	obj46[6] = ns.Object("FistTrap07")
	obj46[7] = ns.Object("FistTrap08")
	obj46[8] = ns.Object("FistTrap09")
	obj46[9] = ns.Object("FistTrap10")
	obj46[10] = ns.Object("FistTrap11")
	obj46[11] = ns.Object("FistTrap12")
	obj46[12] = ns.Object("FistTrap13")
	obj45[0] = ns.Object("FistTrapLight01")
	obj45[1] = ns.Object("FistTrapLight02")
	obj45[2] = ns.Object("FistTrapLight03")
	obj45[3] = ns.Object("FistTrapLight04")
	obj45[4] = ns.Object("FistTrapLight05")
	obj45[5] = ns.Object("FistTrapLight06")
	obj45[6] = ns.Object("FistTrapLight07")
	obj45[7] = ns.Object("FistTrapLight08")
	obj45[8] = ns.Object("FistTrapLight09")
	obj45[9] = ns.Object("FistTrapLight10")
	obj45[10] = ns.Object("FistTrapLight11")
	obj45[11] = ns.Object("FistTrapLight12")
	obj45[12] = ns.Object("FistTrapLight13")
}
func ResetFistTrap(a1 int) {
	ns.ObjectOn(obj45[a1])
	ns.ObjectOn(obj46[a1])
	ns.MoveWaypoint(wp48, ns.GetObjectX(obj46[a1]), ns.GetObjectY(obj46[a1]))
	ns.AudioEvent(ns.TriggerReleased, wp48)
}
func DropFist(a1 int) {
	ns.CastSpellObjectLocation(ns.SPELL_FIST, obj46[a1], ns.GetObjectX(obj46[a1]), ns.GetObjectY(obj46[a1]))
	ns.ObjectOff(obj46[a1])
	ns.ObjectOff(obj45[a1])
	ns.FrameTimerWithArg(ivar47, a1, ResetFistTrap)
}
func ActivateFistTrap01() {
	DropFist(0)
}
func ActivateFistTrap02() {
	DropFist(1)
}
func ActivateFistTrap03() {
	DropFist(2)
}
func ActivateFistTrap04() {
	DropFist(3)
}
func ActivateFistTrap05() {
	DropFist(4)
}
func ActivateFistTrap06() {
	DropFist(5)
}
func ActivateFistTrap07() {
	DropFist(6)
}
func ActivateFistTrap08() {
	DropFist(7)
}
func ActivateFistTrap09() {
	DropFist(8)
}
func ActivateFistTrap10() {
	DropFist(9)
}
func ActivateFistTrap11() {
	DropFist(10)
}
func ActivateFistTrap12() {
	DropFist(11)
}
func ActivateFistTrap13() {
	DropFist(12)
}
func EnableGuardian() {
	ns.WideScreen(false)
	ns.AudioEvent(ns.SwordsmanRecognize, wp48)
	ns.Frozen(ns.GetHost(), false)
	ns.AggressionLevel(obj53, 0.83)
}
func GuardianEnterFX() {
	ns.MoveWaypoint(wp48, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.WallDestroyedStone, wp48)
	ns.Effect(ns.JIGGLE, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 10, 0)
}
func GuardianMoveToArena() {
	ns.Move(obj53, wp56)
}
func InitializeGuardian() {
	obj53 = ns.Object("Guardian")
	obj49 = ns.Object("RightGuardianGate")
	obj50 = ns.Object("LeftGuardianGate")
	obj51 = ns.Object("RightGuardianDoor")
	obj52 = ns.Object("LeftGuardianDoor")
	wp56 = ns.Waypoint("GuardianArena")
	gvar57 = ns.WallGroup("GuardianWalls")
	gvar58 = ns.WallGroup("GuardianDestructableWalls")
	obj54 = ns.Object("GuardianDoorTrigger")
	wp55 = ns.Waypoint("GuardianDoorWP")
	ns.LockDoor(obj49)
	ns.LockDoor(obj50)
}
func GuardianEnter() {
	if !(!flag59 && ns.IsCaller(ns.GetHost())) {
		goto LABEL1
	}
	ns.MusicPushEvent()
	ns.Music(26, 100)
	ns.LockDoor(obj51)
	ns.LockDoor(obj52)
	ns.ObjectOn(obj54)
	flag59 = true
	ns.WideScreen(true)
	ns.Frozen(ns.GetHost(), true)
	ns.FrameTimer(60, GuardianMoveToArena)
	ns.FrameTimer(30, GuardianEnterFX)
LABEL1:
	return
}
func BreakGuardianWalls() {
	ns.ObjectOff(ns.GetTrigger())
	ns.AudioEvent(ns.WallDestroyedStone, wp48)
	ns.Effect(ns.JIGGLE, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 10, 0)
	ns.NoWallSound(true)
	ns.WallGroupOpen(gvar57)
	ns.NoWallSound(false)
	ns.WallGroupBreak(gvar58)
}
func GuardianArrivesInArena() {
	if !ns.IsCaller(obj53) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.LookAtObject(obj53, ns.GetHost())
	ns.FrameTimer(30, EnableGuardian)
	ns.ObjectOff(obj54)
LABEL1:
	return
}
func GuardianDie() {
	if flag60 {
		goto LABEL1
	}
	flag60 = true
	ns.UnlockDoor(obj50)
	ns.UnlockDoor(obj49)
	ns.UnlockDoor(obj52)
	ns.UnlockDoor(obj51)
	ns.MoveWaypoint(wp48, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.FlagDrop, wp48)
	ns.PrintToAll("War04d:WarriorDie")
	ns.GiveXp(ns.GetHost(), 700)
	ns.MusicPopEvent()
LABEL1:
	return
}
func ClearGuardianDoor() {
	if ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.MoveObject(ns.GetCaller(), ns.GetWaypointX(wp55), ns.GetWaypointY(wp55))
LABEL1:
	return
}
func EnableKeeper() {
	ns.LookAtObject(obj64, ns.GetHost())
	ns.AggressionLevel(obj64, 0.83)
}
func InitializeKeeper() {
	obj64 = ns.Object("Keeper")
	obj65 = ns.Object("LeftKeeperEntranceDoor")
	obj66 = ns.Object("RightKeeperEntranceDoor")
	obj67 = ns.Object("LeftKeeperExitDoor")
	obj68 = ns.Object("RightKeeperExitDoor")
	obj69 = ns.Object("KeeperDoorTrigger")
	wp71[0] = ns.Waypoint("BlinkWP1")
	wp71[1] = ns.Waypoint("BlinkWP2")
	wp71[2] = ns.Waypoint("BlinkWP3")
	wp71[3] = ns.Waypoint("BlinkWP4")
	wp72 = ns.Waypoint("KeeperEnterWP")
	wp73 = ns.Waypoint("KeeperDoorWP")
	ns.LockDoor(obj67)
	ns.LockDoor(obj68)
}
func EnemyGoHome() {
	r1 := ns.IsAttackedBy(ns.GetHost(), ns.GetCaller())
	if !(r1 && ns.GetCaller() != obj64) {
		goto LABEL1
	}
	ns.GoBackHome(ns.GetCaller())
LABEL1:
	return
}
func GoInvisible() {
	var (
		v0 int
		v1 float32
		v2 float32
		v3 float32
		v4 float32
		v5 [4]float32
		v6 float32
		v7 int
	)
	v6 = 0
	v7 = 0
	ivar76 += 1
	if gvar63 != gvar62 {
		goto LABEL1
	}
	gvar63 = gvar61
	return
	goto LABEL2
LABEL1:
	gvar63 = gvar62
LABEL2:
	v3 = ns.GetObjectX(obj64)
	v4 = ns.GetObjectY(obj64)
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL3
		}
		v5[v0] = ns.Distance(v3, v4, ns.GetWaypointX(wp71[v0]), ns.GetWaypointY(wp71[v0]))
		v0 += 1
	}
LABEL3:
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL5
		}
		if !(v5[v0] > v6) {
			goto LABEL6
		}
		v6 = v5[v0]
		v7 = v0
	LABEL6:
		v0 += 1
	}
LABEL5:
	if !(ns.CurrentHealth(obj64) > 0) {
		goto LABEL8
	}
	v1 = ns.GetWaypointX(wp71[v7])
	v2 = ns.GetWaypointY(wp71[v7])
	if !(ivar76 >= 4) {
		goto LABEL9
	}
	ivar76 = 0
	ns.Enchant(obj64, ns.ENCHANT_INVISIBLE, 3)
LABEL9:
	ns.Effect(ns.TELEPORT, v3, v4, 0, 0)
	ns.MoveWaypoint(wp48, v3, v4)
	ns.AudioEvent(ns.BlinkCast, wp48)
	ns.Effect(ns.SMOKE_BLAST, v1, v2, 0, 0)
	ns.MoveObject(obj64, v1, v2)
LABEL8:
	return
}
func DisableKeeperDoorTrigger() {
	ns.ObjectOff(obj69)
}
func KeeperEnter() {
	var (
		v0 float32
		v1 float32
	)
	v0 = ns.GetWaypointX(wp72)
	v1 = ns.GetWaypointY(wp72)
	if !(ns.IsCaller(ns.GetHost()) && flag75 == false) {
		goto LABEL1
	}
	ns.ObjectOn(obj69)
	ns.MusicPushEvent()
	ns.Music(4, 100)
	flag75 = true
	ns.Effect(ns.TELEPORT, v0, v1, 0, 0)
	ns.AudioEvent(ns.TeleportIn, wp72)
	ns.MoveObject(obj64, v0, v1)
	ns.LockDoor(obj65)
	ns.LockDoor(obj66)
	ns.FrameTimer(1, EnableKeeper)
	ns.FrameTimer(30, DisableKeeperDoorTrigger)
LABEL1:
	return
}
func KeeperDie() {
	if flag74 {
		goto LABEL1
	}
	flag74 = true
	ns.MoveWaypoint(wp48, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.FlagCapture, wp48)
	ns.PrintToAll("War04f:WizardDie")
	ns.GiveXp(ns.GetHost(), 1000)
	ns.UnlockDoor(obj68)
	ns.UnlockDoor(obj67)
	ns.UnlockDoor(obj66)
	ns.UnlockDoor(obj65)
	ns.MusicPopEvent()
LABEL1:
	return
}
func ClearKeeperDoor() {
	if ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.MoveObject(ns.GetCaller(), ns.GetWaypointX(wp73), ns.GetWaypointY(wp73))
LABEL1:
	return
}
func InitializeFlameGrates() {
	gvar70 = ns.ObjectGroup("FlameGrateFlames")
}
func ToggleFlameGrates() {
	ns.ObjectGroupToggle(gvar70)
}
func InitializeSkullGuns() {
	obj92 = ns.Object("ArrowTrap01")
	obj93 = ns.Object("ArrowTrap02")
	obj94 = ns.Object("ArrowTrap03")
	obj95 = ns.Object("ArrowTrap04")
	obj96 = ns.Object("ArrowTrap05")
	obj97 = ns.Object("ArrowTrap06")
	obj98 = ns.Object("ArrowTrap07")
	obj99 = ns.Object("ArrowTrap08")
	obj100 = ns.Object("ArrowTrap09")
	obj101 = ns.Object("ArrowTrap10")
	obj102 = ns.Object("ArrowTrap11")
}
func InitializeMonsterGroups() {
	gvar78 = ns.ObjectGroup("MonsterGroup01")
}
func MapInitialize() {
	wp48 = ns.Waypoint("AudioOrigin")
	wp77 = ns.Waypoint("SecretAudioOrigin")
	InitializeSkullGuns()
	InitializeSkeletonAmbush()
	InitializeGuardian()
	InitializeKeeper()
	InitializeFistTraps()
	InitializeHallwayArrowTraps()
	InitializeMonsterGroups()
	InitializeBlockTrap()
	InitializeFlameGrates()
	ns.Music(20, 100)
}
func PlayerDeath() {
	ns.DeathScreen(4)
}
func EnableMonsterGroup01() {
	ns.ObjectOff(ns.GetTrigger())
	ns.ObjectGroupOn(gvar78)
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
func SecretSFX() {
	ns.MoveWaypoint(wp77, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretFound, wp77)
}
func Secret01Found() {
	ns.ObjectOff(ns.GetTrigger())
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 100)
	ns.AggressionLevel(ns.Object("Secret01Spider"), 0.83)
	SecretSFX()
}
func Secret02Found() {
	ns.ObjectOff(ns.GetTrigger())
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 50)
	SecretSFX()
}
func CheckRows(a1 int) {
	var v0 int
	v0 = a1
	if v0 == 1 {
		goto LABEL1
	}
	if v0 == 2 {
		goto LABEL2
	}
	if v0 == 3 {
		goto LABEL3
	}
	if v0 == 4 {
		goto LABEL4
	}
	if v0 == 5 {
		goto LABEL5
	}
	if v0 == 6 {
		goto LABEL6
	}
	if v0 == 7 {
		goto LABEL7
	}
	if v0 == 8 {
		goto LABEL8
	}
	if v0 == 9 {
		goto LABEL9
	}
	if v0 == 10 {
		goto LABEL10
	}
	if v0 == 11 {
		goto LABEL11
	}
	goto LABEL12
LABEL1:
	if !(ivar79 > 0) {
		goto LABEL13
	}
	ns.ObjectOn(obj92)
	goto LABEL14
LABEL13:
	ns.ObjectOff(obj92)
LABEL14:
	goto LABEL12
LABEL2:
	if !(ivar80 > 0) {
		goto LABEL15
	}
	ns.ObjectOn(obj93)
	goto LABEL16
LABEL15:
	ns.ObjectOff(obj93)
LABEL16:
	goto LABEL12
LABEL3:
	if !(ivar81 > 0) {
		goto LABEL17
	}
	ns.ObjectOn(obj94)
	goto LABEL18
LABEL17:
	ns.ObjectOff(obj94)
LABEL18:
	goto LABEL12
LABEL4:
	if !(ivar82 > 0) {
		goto LABEL19
	}
	ns.ObjectOn(obj95)
	goto LABEL20
LABEL19:
	ns.ObjectOff(obj95)
LABEL20:
	goto LABEL12
LABEL5:
	if !(ivar83 > 0) {
		goto LABEL21
	}
	ns.ObjectOn(obj96)
	goto LABEL22
LABEL21:
	ns.ObjectOff(obj96)
LABEL22:
	goto LABEL12
LABEL6:
	if !(ivar84 > 0) {
		goto LABEL23
	}
	ns.ObjectOn(obj97)
	goto LABEL24
LABEL23:
	ns.ObjectOff(obj97)
LABEL24:
	goto LABEL12
LABEL7:
	if !(ivar85 > 0) {
		goto LABEL25
	}
	ns.ObjectOn(obj98)
	goto LABEL26
LABEL25:
	ns.ObjectOff(obj98)
LABEL26:
	goto LABEL12
LABEL8:
	if !(ivar86 > 0) {
		goto LABEL27
	}
	ns.ObjectOn(obj99)
	goto LABEL28
LABEL27:
	ns.ObjectOff(obj99)
LABEL28:
	goto LABEL12
LABEL9:
	if !(ivar87 > 0) {
		goto LABEL29
	}
	ns.ObjectOn(obj100)
	goto LABEL30
LABEL29:
	ns.ObjectOff(obj100)
LABEL30:
	goto LABEL12
LABEL10:
	if !(ivar88 > 0) {
		goto LABEL31
	}
	ns.ObjectOn(obj101)
	goto LABEL32
LABEL31:
	ns.ObjectOff(obj101)
LABEL32:
	goto LABEL12
LABEL11:
	if !(ivar89 > 0) {
		goto LABEL33
	}
	ns.ObjectOn(obj102)
	goto LABEL34
LABEL33:
	ns.ObjectOff(obj102)
LABEL34:
	goto LABEL12
LABEL12:
	return
}
func ActivateRow01() {
	ivar79 += 1
	CheckRows(1)
}
func DisableRow01() {
	ivar79 -= 1
	CheckRows(1)
}
func ActivateRow02() {
	ivar80 += 1
	CheckRows(2)
}
func DisableRow02() {
	ivar80 -= 1
	CheckRows(2)
}
func ActivateRow03() {
	ivar81 += 1
	CheckRows(3)
}
func DisableRow03() {
	ivar81 -= 1
	CheckRows(3)
}
func ActivateRow04() {
	ivar82 += 1
	CheckRows(4)
}
func DisableRow04() {
	ivar82 -= 1
	CheckRows(4)
}
func ActivateRow05() {
	ivar83 += 1
	CheckRows(5)
}
func DisableRow05() {
	ivar83 -= 1
	CheckRows(5)
}
func ActivateRow06() {
	ivar84 += 1
	CheckRows(6)
}
func DisableRow06() {
	ivar84 -= 1
	CheckRows(6)
}
func ActivateRow07() {
	ivar85 += 1
	CheckRows(7)
}
func DisableRow07() {
	ivar85 -= 1
	CheckRows(7)
}
func ActivateRow08() {
	ivar86 += 1
	CheckRows(8)
}
func DisableRow08() {
	ivar86 -= 1
	CheckRows(8)
}
func ActivateRow09() {
	ivar87 += 1
	CheckRows(9)
}
func DisableRow09() {
	ivar87 -= 1
	CheckRows(9)
}
func ActivateRow10() {
	ivar88 += 1
	CheckRows(10)
}
func DisableRow10() {
	ivar88 -= 1
	CheckRows(10)
}
func ActivateRow11() {
	ivar89 += 1
	CheckRows(11)
}
func DisableRow11() {
	ivar89 -= 1
	CheckRows(11)
}
func OnEvent(typ string) {
	switch typ {
	case "MapInitialize":
		MapInitialize()
	case "PlayerDeath":
		PlayerDeath()
	}
}
