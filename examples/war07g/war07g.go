package war07g

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	obj4    ns.ObjectID
	obj5    ns.ObjectID
	obj6    ns.ObjectID
	obj7    ns.ObjectID
	obj8    ns.ObjectID
	obj9    ns.ObjectID
	obj10   ns.ObjectID
	obj11   ns.ObjectID
	obj12   ns.ObjectID
	obj13   ns.ObjectID
	obj14   ns.ObjectID
	obj15   ns.ObjectID
	obj16   ns.ObjectID
	obj17   ns.ObjectID
	obj18   ns.ObjectID
	obj19   ns.ObjectID
	obj20   ns.ObjectID
	obj21   ns.ObjectID
	obj22   ns.ObjectID
	obj23   ns.ObjectID
	obj24   ns.ObjectID
	obj25   ns.ObjectID
	obj26   ns.ObjectID
	obj27   ns.ObjectID
	obj28   ns.ObjectID
	obj29   ns.ObjectID
	obj30   ns.ObjectID
	obj31   ns.ObjectID
	obj32   ns.ObjectID
	obj33   ns.ObjectID
	gvar34  ns.ObjectGroupID
	wp35    ns.WaypointID
	gvar36  ns.WallGroupID
	gvar37  ns.WallGroupID
	flag38  bool
	flag39  bool
	flag40  bool
	flag41  bool
	fvar42  float32
	fvar43  float32
	fvar44  float32
	fvar45  float32
	fvar46  float32
	fvar47  float32
	fvar48  float32
	fvar49  float32
	fvar50  float32
	fvar51  float32
	fvar52  float32
	fvar53  float32
	fvar54  float32
	fvar55  float32
	fvar56  float32
	fvar57  float32
	fvar58  float32
	fvar59  float32
	fvar60  float32
	fvar61  float32
	obj62   ns.ObjectID
	gvar63  int
	gvar64  int
	obj65   ns.ObjectID
	obj66   ns.ObjectID
	obj67   ns.ObjectID
	obj68   ns.ObjectID
	obj69   ns.ObjectID
	obj70   ns.ObjectID
	obj71   ns.ObjectID
	obj72   ns.ObjectID
	obj73   ns.ObjectID
	obj74   ns.ObjectID
	obj75   ns.ObjectID
	obj76   ns.ObjectID
	obj77   ns.ObjectID
	obj78   ns.ObjectID
	obj79   ns.ObjectID
	obj80   ns.ObjectID
	obj81   ns.ObjectID
	obj82   ns.ObjectID
	obj83   ns.ObjectID
	obj84   ns.ObjectID
	obj85   ns.ObjectID
	obj86   ns.ObjectID
	obj87   ns.ObjectID
	obj88   ns.ObjectID
	obj89   ns.ObjectID
	obj90   ns.ObjectID
	obj91   ns.ObjectID
	obj92   ns.ObjectID
	obj93   ns.ObjectID
	obj94   ns.ObjectID
	gvar95  int
	gvar96  ns.ObjectGroupID
	gvar97  ns.ObjectGroupID
	gvar98  ns.ObjectGroupID
	gvar99  ns.ObjectGroupID
	gvar100 ns.ObjectGroupID
	gvar101 ns.ObjectGroupID
	gvar102 ns.ObjectGroupID
	gvar103 ns.ObjectGroupID
	gvar104 ns.ObjectGroupID
	flag105 bool
	wp106   ns.WaypointID
	wp107   ns.WaypointID
	wp108   ns.WaypointID
	wp109   ns.WaypointID
	wp110   ns.WaypointID
	gvar111 ns.TimerID
	ivar112 int
	ivar113 int
)

func init() {
	flag38 = false
	flag39 = false
	flag40 = false
	flag41 = false
	flag105 = false
	ivar112 = 0
	ivar113 = 0
}
func LightShow() {
	var (
		v0 int
		v1 int
	)
	ns.MoveWaypoint(wp35, ns.GetObjectX(obj4), ns.GetObjectY(obj4))
	v0 = ns.Random(1, 8)
	v1 = v0
	if v1 == 1 {
		goto LABEL1
	}
	if v1 == 2 {
		goto LABEL2
	}
	if v1 == 3 {
		goto LABEL3
	}
	if v1 == 4 {
		goto LABEL4
	}
	if v1 == 5 {
		goto LABEL5
	}
	if v1 == 6 {
		goto LABEL6
	}
	if v1 == 7 {
		goto LABEL7
	}
	if v1 == 8 {
		goto LABEL8
	}
	goto LABEL9
LABEL1:
	ns.AudioEvent(ns.LightningBolt, wp35)
	ns.Effect(ns.LIGHTNING, ns.GetObjectX(obj6), ns.GetObjectY(obj6), ns.GetObjectX(obj4), ns.GetObjectY(obj4))
	goto LABEL9
LABEL2:
	ns.AudioEvent(ns.LightningBolt, wp35)
	ns.Effect(ns.LIGHTNING, ns.GetObjectX(obj7), ns.GetObjectY(obj7), ns.GetObjectX(obj4), ns.GetObjectY(obj4))
	goto LABEL9
LABEL3:
	ns.AudioEvent(ns.LightningBolt, wp35)
	ns.Effect(ns.LIGHTNING, ns.GetObjectX(obj8), ns.GetObjectY(obj8), ns.GetObjectX(obj4), ns.GetObjectY(obj4))
	goto LABEL9
LABEL4:
	ns.AudioEvent(ns.LightningBolt, wp35)
	ns.Effect(ns.LIGHTNING, ns.GetObjectX(obj9), ns.GetObjectY(obj9), ns.GetObjectX(obj4), ns.GetObjectY(obj4))
	goto LABEL9
LABEL5:
	ns.AudioEvent(ns.LightningBolt, wp35)
	ns.Effect(ns.LIGHTNING, ns.GetObjectX(obj10), ns.GetObjectY(obj10), ns.GetObjectX(obj4), ns.GetObjectY(obj4))
	goto LABEL9
LABEL6:
	ns.AudioEvent(ns.LightningBolt, wp35)
	ns.Effect(ns.LIGHTNING, ns.GetObjectX(obj11), ns.GetObjectY(obj11), ns.GetObjectX(obj4), ns.GetObjectY(obj4))
	goto LABEL9
LABEL7:
	ns.AudioEvent(ns.LightningBolt, wp35)
	ns.Effect(ns.LIGHTNING, ns.GetObjectX(obj12), ns.GetObjectY(obj12), ns.GetObjectX(obj4), ns.GetObjectY(obj4))
	goto LABEL9
LABEL8:
	ns.AudioEvent(ns.LightningBolt, wp35)
	ns.Effect(ns.LIGHTNING, ns.GetObjectX(obj13), ns.GetObjectY(obj13), ns.GetObjectX(obj4), ns.GetObjectY(obj4))
	goto LABEL9
LABEL9:
	if flag41 {
		goto LABEL10
	}
	ns.FrameTimer(30, LightShow)
LABEL10:
	return
}
func PlateOneTrigger() {
	flag38 = true
	if !(flag39 == true && flag40 == true) {
		goto LABEL1
	}
	ns.WallGroupOpen(gvar36)
	ns.MoveWaypoint(wp35, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.BigGong, wp35)
LABEL1:
	return
}
func PlateTwoTrigger() {
	flag39 = true
	if !(flag38 == true && flag40 == true) {
		goto LABEL1
	}
	ns.WallGroupOpen(gvar36)
	ns.MoveWaypoint(wp35, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.BigGong, wp35)
LABEL1:
	return
}
func PlateThreeTrigger() {
	flag40 = true
	if !(flag38 == true && flag39 == true) {
		goto LABEL1
	}
	ns.WallGroupOpen(gvar36)
	ns.MoveWaypoint(wp35, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.BigGong, wp35)
LABEL1:
	return
}
func PlateOneRelease() {
	flag38 = false
}
func PlateTwoRelease() {
	flag39 = false
}
func PlateThreeRelease() {
	flag40 = false
}
func HONTrigger() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.ObjectOff(obj4)
	ns.ObjectOff(obj5)
	ns.ObjectOff(obj14)
	ns.ObjectOff(obj15)
	ns.ObjectOff(obj16)
	ns.ObjectOff(obj17)
	ns.ObjectOff(obj18)
	ns.ObjectOff(obj19)
	ns.ObjectOff(obj20)
	ns.ObjectOff(obj21)
	ns.ObjectOff(obj22)
	ns.ObjectOff(obj23)
	ns.MoveObject(obj24, fvar42, fvar52)
	ns.MoveObject(obj25, fvar43, fvar53)
	ns.MoveObject(obj26, fvar44, fvar54)
	ns.MoveObject(obj27, fvar45, fvar55)
	ns.MoveObject(obj28, fvar46, fvar56)
	ns.MoveObject(obj29, fvar47, fvar57)
	ns.MoveObject(obj30, fvar48, fvar58)
	ns.MoveObject(obj31, fvar49, fvar59)
	ns.MoveObject(obj32, fvar50, fvar60)
	ns.MoveObject(obj33, fvar51, fvar61)
	flag41 = true
	ns.SetHalberd(1)
	ns.PrintToAll("Con07H.scr:HONTada")
	ns.JournalEdit(ns.GetHost(), "GetHON", 4)
	ns.SecondTimer(2, FadeOut)
LABEL1:
	return
}
func MakeMapSwitch() {
	ns.Frozen(ns.GetHost(), false)
	ns.MoveObject(ns.GetHost(), 1922, 2014)
}
func FadeOut() {
	ns.Blind()
	ns.Frozen(ns.GetHost(), true)
	ns.Music(0, 100)
	ns.FrameTimer(60, MakeMapSwitch)
}
func HONSetup() {
	fvar42 = ns.GetObjectX(obj14)
	fvar43 = ns.GetObjectX(obj15)
	fvar44 = ns.GetObjectX(obj16)
	fvar45 = ns.GetObjectX(obj17)
	fvar46 = ns.GetObjectX(obj18)
	fvar47 = ns.GetObjectX(obj19)
	fvar48 = ns.GetObjectX(obj20)
	fvar49 = ns.GetObjectX(obj21)
	fvar50 = ns.GetObjectX(obj22)
	fvar51 = ns.GetObjectX(obj23)
	fvar52 = ns.GetObjectY(obj14)
	fvar53 = ns.GetObjectY(obj15)
	fvar54 = ns.GetObjectY(obj16)
	fvar55 = ns.GetObjectY(obj17)
	fvar56 = ns.GetObjectY(obj18)
	fvar57 = ns.GetObjectY(obj19)
	fvar58 = ns.GetObjectY(obj20)
	fvar59 = ns.GetObjectY(obj21)
	fvar60 = ns.GetObjectY(obj22)
	fvar61 = ns.GetObjectY(obj23)
}
func CreatureSetup() {
}
func Group1Die() {
	ivar112 += 1
	if !(ivar112 >= 2) {
		goto LABEL1
	}
	ns.UnlockDoor(obj86)
	ns.UnlockDoor(obj87)
LABEL1:
	return
}
func StartChamber() {
	ns.AudioEvent(ns.FloorSpikesUp, wp110)
	ns.AudioEvent(ns.FloorSpikesDown, wp109)
	ns.ObjectGroupOn(gvar97)
	ns.ObjectGroupOn(gvar99)
	ns.ObjectGroupOff(gvar96)
	ns.ObjectGroupOff(gvar98)
	flag105 = true
	ns.SecondTimer(3, ChamberSpikes)
}
func ChamberSpikes() {
	if flag105 {
		goto LABEL1
	}
	ns.AudioEvent(ns.FloorSpikesDown, wp109)
	ns.ObjectGroupOff(gvar96)
	ns.ObjectGroupOff(gvar98)
	ns.FrameTimer(10, EnableSpikes2And4)
	goto LABEL2
LABEL1:
	ns.AudioEvent(ns.FloorSpikesDown, wp110)
	ns.ObjectGroupOff(gvar97)
	ns.ObjectGroupOff(gvar99)
	ns.FrameTimer(10, EnableSpikes1And3)
LABEL2:
	gvar111 = ns.SecondTimer(3, ChamberSpikes)
}
func Block1Jiggle() {
	ns.MoveWaypoint(wp35, ns.GetObjectX(obj93), ns.GetObjectY(obj93))
	ns.AudioEvent(ns.HammerMissing, wp35)
	ns.Effect(ns.JIGGLE, ns.GetObjectX(obj93), ns.GetObjectY(obj93), 10, 0)
}
func Block2Jiggle() {
	ns.AudioEvent(ns.HammerMissing, wp106)
	ns.Effect(ns.JIGGLE, ns.GetObjectX(obj94), ns.GetObjectY(obj94), 10, 0)
}
func EnableSpikes2And4() {
	ns.AudioEvent(ns.FloorSpikesUp, wp110)
	ns.ObjectGroupOn(gvar97)
	ns.ObjectGroupOn(gvar99)
	flag105 = true
}
func EnableSpikes1And3() {
	ns.AudioEvent(ns.FloorSpikesUp, wp109)
	ns.ObjectGroupOn(gvar96)
	ns.ObjectGroupOn(gvar98)
	flag105 = false
}
func SetTraitors() {
	ns.ObjectGroupOff(gvar104)
	ns.AudioEvent(ns.TeleportIn, wp107)
	ns.AudioEvent(ns.TeleportIn, wp108)
	ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp107), ns.GetWaypointY(wp107), 0, 0)
	ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp108), ns.GetWaypointY(wp108), 0, 0)
	ns.Effect(ns.SPARK_EXPLOSION, ns.GetWaypointX(wp107), ns.GetWaypointY(wp107), 0, 0)
	ns.Effect(ns.SPARK_EXPLOSION, ns.GetWaypointX(wp108), ns.GetWaypointY(wp108), 0, 0)
	ns.MoveObject(obj65, ns.GetWaypointX(wp107), ns.GetWaypointY(wp107))
	ns.MoveObject(obj66, ns.GetWaypointX(wp108), ns.GetWaypointY(wp108))
	ns.ObjectOn(obj65)
	ns.ObjectOn(obj66)
	ns.LockDoor(obj67)
	ns.LockDoor(obj68)
}
func HaltMusic() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.Music(0, 100)
LABEL1:
	return
}
func NewMusic() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.Music(1, 100)
LABEL1:
	return
}
func TraitorDeath() {
	ivar113 += 1
	if !(ivar113 >= 2) {
		goto LABEL1
	}
	ns.GiveXp(ns.GetHost(), 2000)
LABEL1:
	return
}
func LockStairGate() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.LockDoor(obj69)
	ns.Music(28, 100)
LABEL1:
	return
}
func MapInitialize() {
	obj70 = ns.Object("SnakeHead1Mover")
	obj71 = ns.Object("SnakeHead2Mover")
	obj72 = ns.Object("SnakeHead3Mover")
	obj73 = ns.Object("SnakeHead4Mover")
	obj74 = ns.Object("SnakeHead5Mover")
	obj75 = ns.Object("SnakeHead6Mover")
	obj76 = ns.Object("SnakeHead7Mover")
	obj77 = ns.Object("SnakeHead8Mover")
	obj78 = ns.Object("SnakeHead9Mover")
	obj79 = ns.Object("SnakeHead10Mover")
	obj80 = ns.Object("SnakeHead11Mover")
	obj81 = ns.Object("SnakeHead12Mover")
	obj82 = ns.Object("Block1AMover")
	obj83 = ns.Object("Block1BMover")
	obj84 = ns.Object("Block2AMover")
	obj85 = ns.Object("Block2BMover")
	obj62 = ns.Object("ToLevelOneTP")
	obj92 = ns.Object("ToThreeTP")
	obj65 = ns.Object("Traitor1")
	obj66 = ns.Object("Traitor2")
	obj67 = ns.Object("TraitorDoor1")
	obj68 = ns.Object("TraitorDoor2")
	obj93 = ns.Object("Block1Trigger")
	obj94 = ns.Object("Block2Trigger")
	obj69 = ns.Object("StairGate")
	obj86 = ns.Object("ChamberDoor1")
	obj87 = ns.Object("ChamberDoor2")
	obj88 = ns.Object("ChamberDoor2A")
	obj89 = ns.Object("ChamberDoor2B")
	obj90 = ns.Object("ChamberExitDoor1")
	obj91 = ns.Object("ChamberExitDoor2")
	gvar96 = ns.ObjectGroup("ChamberSpikes1")
	gvar97 = ns.ObjectGroup("ChamberSpikes2")
	gvar98 = ns.ObjectGroup("ChamberSpikes3")
	gvar99 = ns.ObjectGroup("ChamberSpikes4")
	gvar100 = ns.ObjectGroup("SnakeLetterMovers")
	gvar101 = ns.ObjectGroup("Snake3Movers")
	gvar102 = ns.ObjectGroup("Snake4Movers")
	gvar103 = ns.ObjectGroup("ShowdownTriggers")
	wp35 = ns.Waypoint("PlayerSounds")
	wp106 = ns.Waypoint("Jiggle2WP")
	wp107 = ns.Waypoint("Traitor1WP")
	wp108 = ns.Waypoint("Traitor2WP")
	wp109 = ns.Waypoint("Spikes1WP")
	wp110 = ns.Waypoint("Spikes2WP")
	obj4 = ns.Object("HON")
	obj5 = ns.Object("HONBase")
	obj6 = ns.Object("PowerCell1")
	obj7 = ns.Object("PowerCell2")
	obj8 = ns.Object("PowerCell3")
	obj9 = ns.Object("PowerCell4")
	obj10 = ns.Object("PowerCell5")
	obj11 = ns.Object("PowerCell6")
	obj12 = ns.Object("PowerCell7")
	obj13 = ns.Object("PowerCell8")
	obj14 = ns.Object("HeartLight1")
	obj15 = ns.Object("HeartLight2")
	obj16 = ns.Object("HeartLight3")
	obj17 = ns.Object("HeartLight4")
	obj18 = ns.Object("HeartLight5")
	obj19 = ns.Object("HeartLight6")
	obj20 = ns.Object("HeartLight7")
	obj21 = ns.Object("HeartLight8")
	obj22 = ns.Object("HeartLight9")
	obj23 = ns.Object("HeartLight10")
	obj24 = ns.Object("NewLight11")
	obj25 = ns.Object("NewLight12")
	obj26 = ns.Object("NewLight13")
	obj27 = ns.Object("NewLight14")
	obj28 = ns.Object("NewLight15")
	obj29 = ns.Object("NewLight16")
	obj30 = ns.Object("NewLight17")
	obj31 = ns.Object("NewLight18")
	obj32 = ns.Object("NewLight19")
	obj33 = ns.Object("NewLight20")
	gvar34 = ns.ObjectGroup("PowerCells")
	gvar104 = ns.ObjectGroup("TraitorTriggers")
	gvar36 = ns.WallGroup("HONWalls")
	gvar37 = ns.WallGroup("HONExitWalls")
	ns.ObjectGroupOff(gvar96)
	ns.ObjectGroupOff(gvar97)
	ns.ObjectGroupOff(gvar98)
	ns.ObjectGroupOff(gvar99)
	ns.ObjectOn(obj70)
	ns.ObjectOn(obj71)
	ns.ObjectOn(obj72)
	ns.ObjectOn(obj73)
	ns.ObjectOn(obj74)
	ns.ObjectOn(obj75)
	ns.ObjectOn(obj76)
	ns.ObjectOn(obj77)
	ns.ObjectOn(obj78)
	ns.ObjectOn(obj79)
	ns.ObjectOn(obj80)
	ns.ObjectOn(obj81)
	ns.ObjectGroupOn(gvar100)
	ns.ObjectOn(obj82)
	ns.ObjectOn(obj83)
	ns.ObjectOn(obj84)
	ns.ObjectOn(obj85)
	ns.ObjectGroupOn(gvar101)
	ns.ObjectGroupOn(gvar102)
	ns.LockDoor(obj86)
	ns.LockDoor(obj87)
	ns.FrameTimer(2, HONSetup)
	ns.FrameTimer(1, CreatureSetup)
}
func PlayerDeath() {
	ns.DeathScreen(7)
}
func OnEvent(typ string) {
	switch typ {
	case "MapInitialize":
		MapInitialize()
	case "PlayerDeath":
		PlayerDeath()
	}
}
