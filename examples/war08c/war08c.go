package war08c

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	gvar4  ns.ObjectGroupID
	gvar5  ns.ObjectGroupID
	gvar6  ns.ObjectGroupID
	wp7    ns.WaypointID
	obj8   ns.ObjectID
	obj9   ns.ObjectID
	wp10   ns.WaypointID
	wp11   ns.WaypointID
	gvar12 ns.WallGroupID
	gvar13 ns.WallGroupID
	wp14   ns.WaypointID
	obj15  ns.ObjectID
	obj16  ns.ObjectID
	obj17  ns.ObjectID
	obj18  ns.ObjectID
	obj19  ns.ObjectID
	obj20  ns.ObjectID
	obj21  ns.ObjectID
	obj22  ns.ObjectID
	wp23   ns.WaypointID
	wp24   ns.WaypointID
	wp25   ns.WaypointID
	wp26   ns.WaypointID
	wp27   ns.WaypointID
	wp28   ns.WaypointID
	wp29   ns.WaypointID
	wp30   ns.WaypointID
	wp31   ns.WaypointID
	gvar32 int
	gvar33 [10]ns.ObjectGroupID
	gvar34 int
	obj35  ns.ObjectID
	obj36  [9]ns.ObjectID
	gvar37 int
	ivar38 int
	ivar39 int
	ivar40 int
	ivar41 int
	ivar42 int
	ivar43 [10]int
	flag44 [10]bool
	wp45   [10]ns.WaypointID
	str46  [6]string
)

func init() {
	gvar37 = 0
	ivar38 = 1
	ivar39 = 0
	ivar40 = 180
	ivar41 = 0
	ivar42 = 10
}
func ArrowTrapAudio() {
	ns.MoveWaypoint(wp7, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.BowShoot, wp7)
	ns.AudioEvent(ns.CrossBowShoot, wp7)
}
func DisableArrowTrapGroup01() {
	ns.ObjectGroupOff(gvar4)
}
func EnableArrowTrapGroup01() {
	ns.ObjectGroupOn(gvar4)
	ns.FrameTimer(1, DisableArrowTrapGroup01)
}
func DisableArrowTrapGroup02() {
	ns.ObjectGroupOff(gvar5)
}
func ActivateArrowTrap02() {
	ns.ObjectGroupOn(gvar5)
	ns.FrameTimer(1, DisableArrowTrapGroup02)
}
func DisableArrowTrapGroup03() {
	ns.ObjectGroupOff(gvar6)
}
func ActivateArrowTrap03() {
	ns.ObjectGroupOn(gvar6)
	ns.FrameTimer(1, DisableArrowTrapGroup03)
}
func BeholderAttackExit() {
	if !(ns.CurrentHealth(obj8) > 0) {
		goto LABEL1
	}
	if ns.IsVisibleTo(obj8, ns.GetHost()) {
		goto LABEL1
	}
	ns.Move(obj8, wp11)
LABEL1:
	return
}
func BeholderStay() {
	if !ns.IsCaller(obj8) {
		goto LABEL1
	}
	ns.Wander(obj8)
LABEL1:
	if !ns.IsCaller(obj9) {
		goto LABEL2
	}
	ns.Wander(obj9)
LABEL2:
	return
}
func OpenBoulderHallWall01() {
	ns.ObjectOff(ns.GetTrigger())
	ns.WallOpen(ns.Wall(104, 162))
}
func OpenBoulderHallWall02() {
	ns.ObjectOff(ns.GetTrigger())
	ns.WallOpen(ns.Wall(213, 107))
}
func GoldKeyRemoved() {
	ns.WallGroupOpen(gvar12)
	ns.WallGroupOpen(gvar13)
}
func ReleaseGoldKeyBombers() {
	ns.WallGroupOpen(gvar13)
}
func PlayerDeath() {
	ns.DeathScreen(8)
}
func InitializeSummoningCircles() {
	var v0 int
	v0 = 0
	for {
		if !(v0 < 10) {
			goto LABEL1
		}
		flag44[v0] = true
		ivar43[v0] = 0
		v0 += 1
	}
LABEL1:
	str46[0] = "UrchinShaman"
	str46[1] = "Mimic"
	str46[2] = "Bear"
	str46[3] = "Urchin"
	str46[4] = "Imp"
	str46[5] = "WhiteWolf"
	obj35 = ns.Object("SummoningCircleTreasureLight")
	obj36[0] = ns.Object("SummoningCircleTreasure01")
	obj36[1] = ns.Object("SummoningCircleTreasure02")
	obj36[2] = ns.Object("SummoningCircleTreasure03")
	obj36[3] = ns.Object("SummoningCircleTreasure04")
	obj36[4] = ns.Object("SummoningCircleTreasure05")
	obj36[5] = ns.Object("SummoningCircleTreasure06")
	obj36[6] = ns.Object("SummoningCircleTreasure07")
	obj36[7] = ns.Object("SummoningCircleTreasure08")
	obj36[8] = ns.Object("SummoningCircleTreasure09")
	wp45[0] = ns.Waypoint("SummoningCircle01WP")
	wp45[1] = ns.Waypoint("SummoningCircle02WP")
	wp45[2] = ns.Waypoint("SummoningCircle03WP")
	wp45[3] = ns.Waypoint("SummoningCircle04WP")
	wp45[4] = ns.Waypoint("SummoningCircle05WP")
	wp45[5] = ns.Waypoint("SummoningCircle06WP")
	wp45[6] = ns.Waypoint("SummoningCircle07WP")
	wp45[7] = ns.Waypoint("SummoningCircle08WP")
	wp45[8] = ns.Waypoint("SummoningCircle09WP")
	wp45[9] = ns.Waypoint("SummoningCircle10WP")
	gvar33[0] = ns.ObjectGroup("SummonCircle01Lights")
	gvar33[1] = ns.ObjectGroup("SummonCircle02Lights")
	gvar33[2] = ns.ObjectGroup("SummonCircle03Lights")
	gvar33[3] = ns.ObjectGroup("SummonCircle04Lights")
	gvar33[4] = ns.ObjectGroup("SummonCircle05Lights")
	gvar33[5] = ns.ObjectGroup("SummonCircle06Lights")
	gvar33[6] = ns.ObjectGroup("SummonCircle07Lights")
	gvar33[7] = ns.ObjectGroup("SummonCircle08Lights")
	gvar33[8] = ns.ObjectGroup("SummonCircle09Lights")
	gvar33[9] = ns.ObjectGroup("SummonCircle10Lights")
}
func MapInitialize() {
	obj8 = ns.Object("Beholder01")
	obj9 = ns.Object("Beholder02")
	obj15 = ns.Object("Beholder03")
	obj16 = ns.Object("Bat01")
	obj17 = ns.Object("StoneBlock01")
	obj18 = ns.Object("StoneBlock02")
	obj19 = ns.Object("StoneBlock03")
	obj20 = ns.Object("StoneBlock04")
	obj21 = ns.Object("BomberPitElevator")
	obj22 = ns.Object("SpiderPitElevator")
	wp23 = ns.Waypoint("StoneBlock01WP")
	wp26 = ns.Waypoint("StoneBlock02Home")
	wp28 = ns.Waypoint("StoneBlock03Home")
	wp25 = ns.Waypoint("StoneBlock02Dest")
	wp27 = ns.Waypoint("StoneBlock03Dest")
	wp24 = ns.Waypoint("StoneBlock04WP")
	wp14 = ns.Waypoint("GoldKeyWP")
	wp30 = ns.Waypoint("BomberCreationWP")
	wp10 = ns.Waypoint("BomberAudioOrigin")
	wp29 = ns.Waypoint("Beholder01SetPieceWP")
	wp11 = ns.Waypoint("BeholderRoomExitWP")
	wp7 = ns.Waypoint("ArrowTrapAudioOrigin")
	wp31 = ns.Waypoint("SecretAudioOrigin")
	gvar4 = ns.ObjectGroup("ArrowTrapGroup01")
	gvar5 = ns.ObjectGroup("ArrowTrapGroup02")
	gvar6 = ns.ObjectGroup("ArrowTrapGroup03")
	gvar12 = ns.WallGroup("GoldKeyWalls")
	gvar13 = ns.WallGroup("BomberCageWalls")
	InitializeSummoningCircles()
}
func MapEntry() {
	ns.Music(18, 100)
	ns.UnBlind()
}
func EnableSpiderPitElevator() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.ObjectOn(obj22)
LABEL1:
	return
}
func SecretSFX() {
	ns.MoveWaypoint(wp31, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretFound, wp31)
}
func Secret01Found() {
	ns.ObjectOff(ns.GetTrigger())
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 100)
	SecretSFX()
}
func SealTheTempleFX() {
	var (
		v0 float32
		v1 float32
	)
	v0 = ns.GetObjectX(ns.GetHost())
	v1 = ns.GetObjectY(ns.GetHost())
	ns.AudioEvent(ns.HugeFoot, wp23)
	ns.Effect(ns.JIGGLE, v0, v1, 10, 0)
}
func SealTheTemple() {
	ns.ObjectOff(ns.GetTrigger())
	ns.Move(obj17, wp23)
	ns.Move(obj20, wp24)
	ns.AudioEvent(ns.BigGong, wp23)
	ns.AudioEvent(ns.SecretWallClose, wp23)
	ns.FrameTimer(10, SealTheTempleFX)
}
func EnableElevator() {
	ns.ObjectOn(obj21)
}
func ReleasePlayer() {
	ns.WideScreen(false)
	ns.ClearOwner(obj8)
	ns.Frozen(ns.GetHost(), false)
	if !(ns.CurrentHealth(obj16) > 0) {
		goto LABEL1
	}
	ns.ObjectOn(obj16)
LABEL1:
	ns.Delete(ns.Object("War08a:CharmedWolf01"))
	ns.Delete(ns.Object("War08a:CharmedWolf02"))
	ns.Delete(ns.Object("War08a:CharmedWolf03"))
	ns.Delete(ns.Object("War08a:CharmedWolf04"))
	ns.NoWallSound(true)
	ns.WallClose(ns.Wall(169, 87))
	ns.NoWallSound(false)
	ns.FrameTimer(60, EnableElevator)
}
func BeholderSetpieceSEG3() {
	ns.FrameTimer(60, ReleasePlayer)
}
func BeholderSetpieceSEG2() {
	ns.Move(obj8, wp29)
	ns.FrameTimer(180, ReleasePlayer)
}
func BeholderSetpiece() {
	ns.SetOwner(ns.GetHost(), obj8)
	ns.ObjectOff(ns.GetTrigger())
	ns.Frozen(ns.GetHost(), true)
	ns.CreatureIdle(ns.GetHost())
	ns.WideScreen(true)
	if !(ns.CurrentHealth(obj16) > 0) {
		goto LABEL1
	}
	ns.ObjectOff(obj16)
LABEL1:
	ns.FrameTimer(60, BeholderSetpieceSEG2)
}
func DisableBomberPitElevator() {
	ns.ObjectOff(obj21)
}
func BlockGatesFX() {
	var (
		v0 float32
		v1 float32
	)
	v0 = ns.GetObjectX(ns.GetHost())
	v1 = ns.GetObjectY(ns.GetHost())
	ns.AudioEvent(ns.HugeFoot, wp25)
	ns.Effect(ns.JIGGLE, v0, v1, 10, 0)
}
func CloseBlockGates() {
	ns.Move(obj18, wp25)
	ns.Move(obj19, wp27)
	ns.AudioEvent(ns.SecretWallClose, wp25)
	ns.AudioEvent(ns.SecretWallClose, wp27)
	ns.FrameTimer(10, BlockGatesFX)
}
func SealBeholderRoom() {
	ns.ObjectOff(ns.GetTrigger())
	ns.WallClose(ns.Wall(209, 109))
	ns.MoveObject(obj8, ns.GetWaypointX(wp11), ns.GetWaypointY(wp11))
	ns.AggressionLevel(obj8, 0.83)
	ns.AggressionLevel(obj9, 0.83)
	ns.Wander(obj8)
	ns.Wander(obj9)
}
func ResetSummonCircle(a1 int) {
	flag44[a1] = true
	ns.ObjectGroupOn(gvar33[a1])
}
func SummoningCircleCleared(a1 int) {
	ivar39 += 1
	if ivar39 != 9 {
		goto LABEL1
	}
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 500)
	ns.Effect(ns.BLUE_SPARKS, ns.GetWaypointX(wp45[a1]), ns.GetWaypointY(wp45[a1]), 0, 0)
	ns.MoveObject(obj35, ns.GetWaypointX(wp45[a1]), ns.GetWaypointY(wp45[a1]))
	ns.MoveObject(obj36[ivar39-1], ns.GetWaypointX(wp45[a1]), ns.GetWaypointY(wp45[a1]))
	ns.AudioEvent(ns.SmallGong, wp45[a1])
	goto LABEL2
LABEL1:
	ns.Effect(ns.BLUE_SPARKS, ns.GetWaypointX(wp45[a1]), ns.GetWaypointY(wp45[a1]), 0, 0)
	ns.MoveObject(obj35, ns.GetWaypointX(wp45[a1]), ns.GetWaypointY(wp45[a1]))
	ns.MoveObject(obj36[ivar39-1], ns.GetWaypointX(wp45[a1]), ns.GetWaypointY(wp45[a1]))
	ns.AudioEvent(ns.BigGong, wp45[a1])
LABEL2:
	return
}
func SummonCircle01Activate() {
	var (
		v0 ns.ObjectID
		v1 int
	)
	gvar37 = 0
	if !(ns.IsCaller(ns.GetHost()) && flag44[gvar37] == true) {
		goto LABEL1
	}
	flag44[gvar37] = false
	ns.ObjectGroupOff(gvar33[gvar37])
	ns.AudioEvent(ns.SummonCast, wp45[gvar37])
	ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp45[gvar37]), ns.GetWaypointY(wp45[gvar37]), 0, 0)
	if r4 := ns.Random(0, 100); !(r4 <= ivar42) {
		goto LABEL2
	}
	v1 = ns.Random(0, 1)
	if v1 == 0 {
		goto LABEL3
	}
	if v1 == 1 {
		goto LABEL4
	}
	goto LABEL5
LABEL3:
	v0 = ns.CreateObject("BomberYellow", wp45[gvar37])
	ns.TrapSpells(v0, ns.SPELL_METEOR, "", "")
	goto LABEL5
LABEL4:
	v0 = ns.CreateObject("Bomber", wp45[gvar37])
	goto LABEL5
LABEL5:
	goto LABEL6
LABEL2:
	ivar41 = ns.Random(0, 5)
	ns.CreateObject(str46[ivar41], wp45[gvar37])
LABEL6:
	ivar43[gvar37] += 1
	if !(ivar43[gvar37] > ivar38) {
		goto LABEL7
	}
	ns.FrameTimerWithArg(30, gvar37, SummoningCircleCleared)
	goto LABEL1
LABEL7:
	ns.FrameTimerWithArg(ivar40, gvar37, ResetSummonCircle)
LABEL1:
	return
}
func SummonCircle02Activate() {
	var (
		v0 ns.ObjectID
		v1 int
	)
	gvar37 = 1
	if !(ns.IsCaller(ns.GetHost()) && flag44[gvar37] == true) {
		goto LABEL1
	}
	flag44[gvar37] = false
	ns.ObjectGroupOff(gvar33[gvar37])
	ns.AudioEvent(ns.SummonCast, wp45[gvar37])
	ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp45[gvar37]), ns.GetWaypointY(wp45[gvar37]), 0, 0)
	if r4 := ns.Random(0, 100); !(r4 <= ivar42) {
		goto LABEL2
	}
	v1 = ns.Random(0, 1)
	if v1 == 0 {
		goto LABEL3
	}
	if v1 == 1 {
		goto LABEL4
	}
	goto LABEL5
LABEL3:
	v0 = ns.CreateObject("BomberYellow", wp45[gvar37])
	ns.TrapSpells(v0, ns.SPELL_METEOR, "", "")
	goto LABEL5
LABEL4:
	v0 = ns.CreateObject("Bomber", wp45[gvar37])
	goto LABEL5
LABEL5:
	goto LABEL6
LABEL2:
	ivar41 = ns.Random(0, 5)
	ns.CreateObject(str46[ivar41], wp45[gvar37])
LABEL6:
	ivar43[gvar37] += 1
	if !(ivar43[gvar37] > ivar38) {
		goto LABEL7
	}
	ns.FrameTimerWithArg(30, gvar37, SummoningCircleCleared)
	goto LABEL1
LABEL7:
	ns.FrameTimerWithArg(ivar40, gvar37, ResetSummonCircle)
LABEL1:
	return
}
func SummonCircle03Activate() {
	var (
		v0 ns.ObjectID
		v1 int
	)
	gvar37 = 2
	if !(ns.IsCaller(ns.GetHost()) && flag44[gvar37] == true) {
		goto LABEL1
	}
	flag44[gvar37] = false
	ns.ObjectGroupOff(gvar33[gvar37])
	ns.AudioEvent(ns.SummonCast, wp45[gvar37])
	ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp45[gvar37]), ns.GetWaypointY(wp45[gvar37]), 0, 0)
	if r4 := ns.Random(0, 100); !(r4 <= ivar42) {
		goto LABEL2
	}
	v1 = ns.Random(0, 1)
	if v1 == 0 {
		goto LABEL3
	}
	if v1 == 1 {
		goto LABEL4
	}
	goto LABEL5
LABEL3:
	v0 = ns.CreateObject("BomberYellow", wp45[gvar37])
	ns.TrapSpells(v0, ns.SPELL_METEOR, "", "")
	goto LABEL5
LABEL4:
	v0 = ns.CreateObject("Bomber", wp45[gvar37])
	goto LABEL5
LABEL5:
	goto LABEL6
LABEL2:
	ivar41 = ns.Random(0, 5)
	ns.CreateObject(str46[ivar41], wp45[gvar37])
LABEL6:
	ivar43[gvar37] += 1
	if !(ivar43[gvar37] > ivar38) {
		goto LABEL7
	}
	ns.FrameTimerWithArg(30, gvar37, SummoningCircleCleared)
	goto LABEL1
LABEL7:
	ns.FrameTimerWithArg(ivar40, gvar37, ResetSummonCircle)
LABEL1:
	return
}
func SummonCircle04Activate() {
	var (
		v0 ns.ObjectID
		v1 int
	)
	gvar37 = 3
	if !(ns.IsCaller(ns.GetHost()) && flag44[gvar37] == true) {
		goto LABEL1
	}
	flag44[gvar37] = false
	ns.ObjectGroupOff(gvar33[gvar37])
	ns.AudioEvent(ns.SummonCast, wp45[gvar37])
	ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp45[gvar37]), ns.GetWaypointY(wp45[gvar37]), 0, 0)
	if r4 := ns.Random(0, 100); !(r4 <= ivar42) {
		goto LABEL2
	}
	v1 = ns.Random(0, 1)
	if v1 == 0 {
		goto LABEL3
	}
	if v1 == 1 {
		goto LABEL4
	}
	goto LABEL5
LABEL3:
	v0 = ns.CreateObject("BomberYellow", wp45[gvar37])
	ns.TrapSpells(v0, ns.SPELL_METEOR, "", "")
	goto LABEL5
LABEL4:
	v0 = ns.CreateObject("Bomber", wp45[gvar37])
	goto LABEL5
LABEL5:
	goto LABEL6
LABEL2:
	ivar41 = ns.Random(0, 5)
	ns.CreateObject(str46[ivar41], wp45[gvar37])
LABEL6:
	ivar43[gvar37] += 1
	if !(ivar43[gvar37] > ivar38) {
		goto LABEL7
	}
	ns.FrameTimerWithArg(30, gvar37, SummoningCircleCleared)
	goto LABEL1
LABEL7:
	ns.FrameTimerWithArg(ivar40, gvar37, ResetSummonCircle)
LABEL1:
	return
}
func SummonCircle05Activate() {
	var (
		v0 ns.ObjectID
		v1 int
	)
	gvar37 = 4
	if !(ns.IsCaller(ns.GetHost()) && flag44[gvar37] == true) {
		goto LABEL1
	}
	flag44[gvar37] = false
	ns.ObjectGroupOff(gvar33[gvar37])
	ns.AudioEvent(ns.SummonCast, wp45[gvar37])
	ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp45[gvar37]), ns.GetWaypointY(wp45[gvar37]), 0, 0)
	if r4 := ns.Random(0, 100); !(r4 <= ivar42) {
		goto LABEL2
	}
	v1 = ns.Random(0, 1)
	if v1 == 0 {
		goto LABEL3
	}
	if v1 == 1 {
		goto LABEL4
	}
	goto LABEL5
LABEL3:
	v0 = ns.CreateObject("BomberYellow", wp45[gvar37])
	ns.TrapSpells(v0, ns.SPELL_METEOR, "", "")
	goto LABEL5
LABEL4:
	v0 = ns.CreateObject("Bomber", wp45[gvar37])
	goto LABEL5
LABEL5:
	goto LABEL6
LABEL2:
	ivar41 = ns.Random(0, 5)
	ns.CreateObject(str46[ivar41], wp45[gvar37])
LABEL6:
	ivar43[gvar37] += 1
	if !(ivar43[gvar37] > ivar38) {
		goto LABEL7
	}
	ns.FrameTimerWithArg(30, gvar37, SummoningCircleCleared)
	goto LABEL1
LABEL7:
	ns.FrameTimerWithArg(ivar40, gvar37, ResetSummonCircle)
LABEL1:
	return
}
func SummonCircle06Activate() {
	var (
		v0 ns.ObjectID
		v1 int
	)
	gvar37 = 5
	if !(ns.IsCaller(ns.GetHost()) && flag44[gvar37] == true) {
		goto LABEL1
	}
	flag44[gvar37] = false
	ns.ObjectGroupOff(gvar33[gvar37])
	ns.AudioEvent(ns.SummonCast, wp45[gvar37])
	ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp45[gvar37]), ns.GetWaypointY(wp45[gvar37]), 0, 0)
	if r4 := ns.Random(0, 100); !(r4 <= ivar42) {
		goto LABEL2
	}
	v1 = ns.Random(0, 1)
	if v1 == 0 {
		goto LABEL3
	}
	if v1 == 1 {
		goto LABEL4
	}
	goto LABEL5
LABEL3:
	v0 = ns.CreateObject("BomberYellow", wp45[gvar37])
	ns.TrapSpells(v0, ns.SPELL_METEOR, "", "")
	goto LABEL5
LABEL4:
	v0 = ns.CreateObject("Bomber", wp45[gvar37])
	goto LABEL5
LABEL5:
	goto LABEL6
LABEL2:
	ivar41 = ns.Random(0, 5)
	ns.CreateObject(str46[ivar41], wp45[gvar37])
LABEL6:
	ivar43[gvar37] += 1
	if !(ivar43[gvar37] > ivar38) {
		goto LABEL7
	}
	ns.FrameTimerWithArg(30, gvar37, SummoningCircleCleared)
	goto LABEL1
LABEL7:
	ns.FrameTimerWithArg(ivar40, gvar37, ResetSummonCircle)
LABEL1:
	return
}
func SummonCircle07Activate() {
	var (
		v0 ns.ObjectID
		v1 int
	)
	gvar37 = 6
	if !(ns.IsCaller(ns.GetHost()) && flag44[gvar37] == true) {
		goto LABEL1
	}
	flag44[gvar37] = false
	ns.ObjectGroupOff(gvar33[gvar37])
	ns.AudioEvent(ns.SummonCast, wp45[gvar37])
	ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp45[gvar37]), ns.GetWaypointY(wp45[gvar37]), 0, 0)
	if r4 := ns.Random(0, 100); !(r4 <= ivar42) {
		goto LABEL2
	}
	v1 = ns.Random(0, 1)
	if v1 == 0 {
		goto LABEL3
	}
	if v1 == 1 {
		goto LABEL4
	}
	goto LABEL5
LABEL3:
	v0 = ns.CreateObject("BomberYellow", wp45[gvar37])
	ns.TrapSpells(v0, ns.SPELL_METEOR, "", "")
	goto LABEL5
LABEL4:
	v0 = ns.CreateObject("Bomber", wp45[gvar37])
	goto LABEL5
LABEL5:
	goto LABEL6
LABEL2:
	ivar41 = ns.Random(0, 5)
	ns.CreateObject(str46[ivar41], wp45[gvar37])
LABEL6:
	ivar43[gvar37] += 1
	if !(ivar43[gvar37] > ivar38) {
		goto LABEL7
	}
	ns.FrameTimerWithArg(30, gvar37, SummoningCircleCleared)
	goto LABEL1
LABEL7:
	ns.FrameTimerWithArg(ivar40, gvar37, ResetSummonCircle)
LABEL1:
	return
}
func SummonCircle08Activate() {
	var (
		v0 ns.ObjectID
		v1 int
	)
	gvar37 = 7
	if !(ns.IsCaller(ns.GetHost()) && flag44[gvar37] == true) {
		goto LABEL1
	}
	flag44[gvar37] = false
	ns.ObjectGroupOff(gvar33[gvar37])
	ns.AudioEvent(ns.SummonCast, wp45[gvar37])
	ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp45[gvar37]), ns.GetWaypointY(wp45[gvar37]), 0, 0)
	if r4 := ns.Random(0, 100); !(r4 <= ivar42) {
		goto LABEL2
	}
	v1 = ns.Random(0, 1)
	if v1 == 0 {
		goto LABEL3
	}
	if v1 == 1 {
		goto LABEL4
	}
	goto LABEL5
LABEL3:
	v0 = ns.CreateObject("BomberYellow", wp45[gvar37])
	ns.TrapSpells(v0, ns.SPELL_METEOR, "", "")
	goto LABEL5
LABEL4:
	v0 = ns.CreateObject("Bomber", wp45[gvar37])
	goto LABEL5
LABEL5:
	goto LABEL6
LABEL2:
	ivar41 = ns.Random(0, 5)
	ns.CreateObject(str46[ivar41], wp45[gvar37])
LABEL6:
	ivar43[gvar37] += 1
	if !(ivar43[gvar37] > ivar38) {
		goto LABEL7
	}
	ns.FrameTimerWithArg(30, gvar37, SummoningCircleCleared)
	goto LABEL1
LABEL7:
	ns.FrameTimerWithArg(ivar40, gvar37, ResetSummonCircle)
LABEL1:
	return
}
func SummonCircle09Activate() {
	var (
		v0 ns.ObjectID
		v1 int
	)
	gvar37 = 8
	if !(ns.IsCaller(ns.GetHost()) && flag44[gvar37] == true) {
		goto LABEL1
	}
	flag44[gvar37] = false
	ns.ObjectGroupOff(gvar33[gvar37])
	ns.AudioEvent(ns.SummonCast, wp45[gvar37])
	ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp45[gvar37]), ns.GetWaypointY(wp45[gvar37]), 0, 0)
	if r4 := ns.Random(0, 100); !(r4 <= ivar42) {
		goto LABEL2
	}
	v1 = ns.Random(0, 1)
	if v1 == 0 {
		goto LABEL3
	}
	if v1 == 1 {
		goto LABEL4
	}
	goto LABEL5
LABEL3:
	v0 = ns.CreateObject("BomberYellow", wp45[gvar37])
	ns.TrapSpells(v0, ns.SPELL_METEOR, "", "")
	goto LABEL5
LABEL4:
	v0 = ns.CreateObject("Bomber", wp45[gvar37])
	goto LABEL5
LABEL5:
	goto LABEL6
LABEL2:
	ivar41 = ns.Random(0, 5)
	ns.CreateObject(str46[ivar41], wp45[gvar37])
LABEL6:
	ivar43[gvar37] += 1
	if !(ivar43[gvar37] > ivar38) {
		goto LABEL7
	}
	ns.FrameTimerWithArg(30, gvar37, SummoningCircleCleared)
	goto LABEL1
LABEL7:
	ns.FrameTimerWithArg(ivar40, gvar37, ResetSummonCircle)
LABEL1:
	return
}
func SummonCircle10Activate() {
	var (
		v0 ns.ObjectID
		v1 int
	)
	gvar37 = 9
	if !(ns.IsCaller(ns.GetHost()) && flag44[gvar37] == true) {
		goto LABEL1
	}
	flag44[gvar37] = false
	ns.ObjectGroupOff(gvar33[gvar37])
	ns.AudioEvent(ns.SummonCast, wp45[gvar37])
	ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp45[gvar37]), ns.GetWaypointY(wp45[gvar37]), 0, 0)
	if r4 := ns.Random(0, 100); !(r4 <= ivar42) {
		goto LABEL2
	}
	v1 = ns.Random(0, 1)
	if v1 == 0 {
		goto LABEL3
	}
	if v1 == 1 {
		goto LABEL4
	}
	goto LABEL5
LABEL3:
	v0 = ns.CreateObject("BomberYellow", wp45[gvar37])
	ns.TrapSpells(v0, ns.SPELL_METEOR, "", "")
	goto LABEL5
LABEL4:
	v0 = ns.CreateObject("Bomber", wp45[gvar37])
	goto LABEL5
LABEL5:
	goto LABEL6
LABEL2:
	ivar41 = ns.Random(0, 5)
	ns.CreateObject(str46[ivar41], wp45[gvar37])
LABEL6:
	ivar43[gvar37] += 1
	if !(ivar43[gvar37] > ivar38) {
		goto LABEL7
	}
	ns.FrameTimerWithArg(30, gvar37, SummoningCircleCleared)
	goto LABEL1
LABEL7:
	ns.FrameTimerWithArg(ivar40, gvar37, ResetSummonCircle)
LABEL1:
	return
}
func EnableBomberPitElevator() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.ObjectOn(obj21)
LABEL1:
	return
}
func OpenBomberPitExit() {
	ns.ObjectOff(ns.GetTrigger())
	ns.WallOpen(ns.Wall(159, 109))
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
