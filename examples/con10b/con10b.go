package con10b

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	gvar4   ns.ObjectGroupID
	obj5    ns.ObjectID
	obj6    ns.ObjectID
	obj7    ns.ObjectID
	obj8    ns.ObjectID
	obj9    ns.ObjectID
	obj10   ns.ObjectID
	obj11   ns.ObjectID
	obj12   ns.ObjectID
	obj13   ns.ObjectID
	wp14    ns.WaypointID
	wp15    ns.WaypointID
	wp16    ns.WaypointID
	wp17    ns.WaypointID
	wp18    ns.WaypointID
	wp19    ns.WaypointID
	wp20    ns.WaypointID
	wp21    ns.WaypointID
	wp22    ns.WaypointID
	obj23   ns.ObjectID
	obj24   ns.ObjectID
	obj25   ns.ObjectID
	obj26   ns.ObjectID
	obj27   ns.ObjectID
	obj28   ns.ObjectID
	wp29    ns.WaypointID
	wp30    ns.WaypointID
	wp31    ns.WaypointID
	wp32    ns.WaypointID
	wp33    ns.WaypointID
	wp34    ns.WaypointID
	wp35    ns.WaypointID
	wp36    ns.WaypointID
	wp37    ns.WaypointID
	wp38    ns.WaypointID
	wp39    ns.WaypointID
	wp40    ns.WaypointID
	gvar41  int
	gvar42  int
	gvar43  int
	obj44   ns.ObjectID
	obj45   ns.ObjectID
	obj46   ns.ObjectID
	obj47   ns.ObjectID
	obj48   ns.ObjectID
	wp49    ns.WaypointID
	wp50    ns.WaypointID
	wp51    ns.WaypointID
	wp52    ns.WaypointID
	wp53    ns.WaypointID
	wp54    ns.WaypointID
	wp55    ns.WaypointID
	wp56    ns.WaypointID
	wp57    ns.WaypointID
	wp58    ns.WaypointID
	gvar59  int
	gvar60  int
	gvar61  int
	obj62   ns.ObjectID
	obj63   ns.ObjectID
	obj64   ns.ObjectID
	wp65    ns.WaypointID
	wp66    ns.WaypointID
	wp67    ns.WaypointID
	wp68    ns.WaypointID
	wp69    ns.WaypointID
	wp70    ns.WaypointID
	gvar71  int
	gvar72  int
	gvar73  int
	obj74   ns.ObjectID
	obj75   ns.ObjectID
	wp76    ns.WaypointID
	wp77    ns.WaypointID
	gvar78  ns.WallGroupID
	gvar79  ns.ObjectGroupID
	gvar80  ns.WallGroupID
	gvar81  ns.WallGroupID
	gvar82  ns.WallGroupID
	obj83   ns.ObjectID
	obj84   ns.ObjectID
	gvar85  ns.ObjectGroupID
	gvar86  ns.ObjectGroupID
	gvar87  ns.WallGroupID
	obj88   ns.ObjectID
	wp89    ns.WaypointID
	gvar90  ns.WallGroupID
	obj91   ns.ObjectID
	obj92   ns.ObjectID
	obj93   ns.ObjectID
	wp94    ns.WaypointID
	wp95    ns.WaypointID
	wp96    ns.WaypointID
	obj97   ns.ObjectID
	obj98   ns.ObjectID
	obj99   ns.ObjectID
	wp100   ns.WaypointID
	wp101   ns.WaypointID
	gvar102 ns.ObjectGroupID
	gvar103 ns.WallGroupID
	gvar104 ns.WallGroupID
	gvar105 ns.WallGroupID
	gvar106 ns.WallGroupID
	gvar107 ns.WallGroupID
	obj108  ns.ObjectID
	obj109  ns.ObjectID
	gvar110 ns.WallGroupID
	obj111  [2]ns.ObjectID
	obj112  [2]ns.ObjectID
	obj113  ns.ObjectID
	wp114   ns.WaypointID
	wp115   ns.WaypointID
	wp116   ns.WaypointID
	fvar117 float32
	fvar118 float32
	flag119 bool
	flag120 bool
)

func init() {
	gvar41 = 0
	gvar42 = 1
	gvar43 = gvar41
	gvar59 = 0
	gvar60 = 1
	gvar61 = gvar59
	gvar71 = 0
	gvar72 = 1
	gvar73 = gvar71
	flag119 = false
	flag120 = false
}
func DisableArrow1() {
	ns.ObjectOff(obj5)
}
func EnableArrow1() {
	ns.ObjectOn(obj5)
	ns.AudioEvent(ns.BowShoot, wp14)
	ns.FrameTimer(1, DisableArrow1)
}
func DisableArrow2() {
	ns.ObjectOff(obj6)
}
func EnableArrow2() {
	ns.ObjectOn(obj6)
	ns.AudioEvent(ns.BowShoot, wp15)
	ns.FrameTimer(1, DisableArrow2)
}
func DisableArrow3() {
	ns.ObjectOff(obj7)
}
func EnableArrow3() {
	ns.ObjectOn(obj7)
	ns.AudioEvent(ns.BowShoot, wp16)
	ns.FrameTimer(1, DisableArrow3)
}
func DisableArrow4() {
	ns.ObjectOff(obj8)
}
func EnableArrow4() {
	ns.ObjectOn(obj8)
	ns.AudioEvent(ns.BowShoot, wp17)
	ns.FrameTimer(1, DisableArrow4)
}
func DisableArrow5() {
	ns.ObjectOff(obj9)
}
func EnableArrow5() {
	ns.ObjectOn(obj9)
	ns.AudioEvent(ns.BowShoot, wp18)
	ns.FrameTimer(1, DisableArrow5)
}
func DisableArrow6() {
	ns.ObjectOff(obj10)
}
func EnableArrow6() {
	ns.ObjectOn(obj10)
	ns.AudioEvent(ns.BowShoot, wp19)
	ns.FrameTimer(1, DisableArrow6)
}
func DisableArrow7() {
	ns.ObjectOff(obj11)
}
func EnableArrow7() {
	ns.ObjectOn(obj11)
	ns.AudioEvent(ns.BowShoot, wp20)
	ns.FrameTimer(1, DisableArrow7)
}
func DisableArrow8() {
	ns.ObjectOff(obj12)
}
func EnableArrow8() {
	ns.ObjectOn(obj12)
	ns.AudioEvent(ns.BowShoot, wp21)
	ns.FrameTimer(1, DisableArrow8)
}
func DisableArrow9() {
	ns.ObjectOff(obj13)
}
func EnableArrow9() {
	ns.ObjectOn(obj13)
	ns.AudioEvent(ns.BowShoot, wp22)
	ns.FrameTimer(1, DisableArrow9)
}
func DisableArrowTriggers() {
	ns.ObjectGroupOff(gvar4)
}
func EnableArrowTriggers() {
	ns.ObjectGroupOn(gvar4)
}
func BlockAGotoB() {
	ns.Move(obj23, wp29)
	ns.Move(obj24, wp30)
	ns.Move(obj25, wp31)
	ns.Move(obj26, wp32)
	ns.Move(obj27, wp33)
	ns.Move(obj28, wp34)
	gvar43 = gvar42
	ns.AudioEvent(ns.SpikeBlockMove, wp30)
	ns.AudioEvent(ns.BoulderMove, wp30)
}
func BlockAGotoC() {
	ns.Move(obj23, wp35)
	ns.Move(obj24, wp36)
	ns.Move(obj25, wp37)
	ns.Move(obj26, wp38)
	ns.Move(obj27, wp39)
	ns.Move(obj28, wp40)
	gvar43 = gvar41
	ns.AudioEvent(ns.SpikeBlockMove, wp30)
	ns.AudioEvent(ns.BoulderMove, wp30)
}
func BlockAGo() {
	var v0 int
	v0 = gvar43
	if v0 == gvar41 {
		goto LABEL1
	}
	if v0 == gvar42 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	BlockAGotoB()
	goto LABEL3
LABEL2:
	BlockAGotoC()
	goto LABEL3
LABEL3:
	return
}
func BlockBGotoB() {
	ns.Move(obj44, wp49)
	ns.Move(obj45, wp50)
	ns.Move(obj46, wp51)
	ns.Move(obj47, wp52)
	ns.Move(obj48, wp53)
	gvar61 = gvar60
	ns.AudioEvent(ns.SpikeBlockMove, wp50)
	ns.AudioEvent(ns.BoulderMove, wp50)
}
func BlockBGotoC() {
	ns.Move(obj44, wp54)
	ns.Move(obj45, wp55)
	ns.Move(obj46, wp56)
	ns.Move(obj47, wp57)
	ns.Move(obj48, wp58)
	gvar61 = gvar59
	ns.AudioEvent(ns.SpikeBlockMove, wp50)
	ns.AudioEvent(ns.BoulderMove, wp50)
}
func BlockBGo() {
	var v0 int
	v0 = gvar61
	if v0 == gvar59 {
		goto LABEL1
	}
	if v0 == gvar60 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	BlockBGotoB()
	goto LABEL3
LABEL2:
	BlockBGotoC()
	goto LABEL3
LABEL3:
	return
}
func BlockCGotoB() {
	ns.Move(obj62, wp65)
	ns.Move(obj63, wp66)
	ns.Move(obj64, wp67)
	gvar73 = gvar72
	ns.AudioEvent(ns.SpikeBlockMove, wp65)
	ns.AudioEvent(ns.BoulderMove, wp65)
}
func BlockCGotoC() {
	ns.Move(obj62, wp68)
	ns.Move(obj63, wp69)
	ns.Move(obj64, wp70)
	gvar73 = gvar71
	ns.AudioEvent(ns.SpikeBlockMove, wp68)
	ns.AudioEvent(ns.BoulderMove, wp68)
}
func BlockCGo() {
	var v0 int
	v0 = gvar73
	if v0 == gvar71 {
		goto LABEL1
	}
	if v0 == gvar72 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	BlockCGotoB()
	goto LABEL3
LABEL2:
	BlockCGotoC()
	goto LABEL3
LABEL3:
	return
}
func PlayerDeath() {
	ns.DeathScreen(10)
}
func OrbLightOff() {
	ns.ObjectOff(obj74)
	ns.MoveObject(obj75, ns.GetWaypointX(wp76), ns.GetWaypointY(wp76))
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
	ns.MoveObject(obj74, ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger()))
	ns.MoveObject(obj75, ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger())+6)
	ns.MoveWaypoint(wp77, ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger()))
	ns.ObjectOn(obj74)
	ns.AudioEvent(ns.BallBounce, wp77)
	ns.PushObject(ns.GetCaller(), 60, v2, v3)
	ns.FrameTimer(4, OrbLightOff)
}
func OpenBouncySecretWallGroup() {
	ns.WallGroupOpen(gvar78)
}
func InitializeVampireKnights() {
	obj111[0] = ns.Object("VampireKnight1")
	obj112[0] = ns.Object("VKBat1")
	wp114 = ns.Waypoint("VKSpot")
	wp115 = ns.Waypoint("VampireKnight1WP")
	wp116 = ns.Waypoint("BatCreate")
}
func MapInitialize() {
	gvar80 = ns.WallGroup("ZombieAmbushWalls")
	gvar79 = ns.ObjectGroup("ZombieAmbushGroup")
	gvar81 = ns.WallGroup("SkeletonWallGroup")
	gvar82 = ns.WallGroup("SkeletonWallGroup2")
	obj83 = ns.Object("SkeletonTrigger1")
	obj84 = ns.Object("SkeletonElevator")
	gvar85 = ns.ObjectGroup("SkelAmbushGroup")
	gvar86 = ns.ObjectGroup("SecretMechGroup")
	gvar87 = ns.WallGroup("SecretMechWallGroup")
	obj88 = ns.Object("SecretMechTrig")
	wp89 = ns.Waypoint("Secret1AudioWP")
	obj23 = ns.Object("BlockA1")
	obj24 = ns.Object("BlockA2")
	obj25 = ns.Object("BlockA3")
	obj26 = ns.Object("BlockA4")
	obj27 = ns.Object("BlockA5")
	obj28 = ns.Object("BlockA6")
	wp29 = ns.Waypoint("BlockWPA1b")
	wp30 = ns.Waypoint("BlockWPA2b")
	wp31 = ns.Waypoint("BlockWPA3b")
	wp32 = ns.Waypoint("BlockWPA4b")
	wp33 = ns.Waypoint("BlockWPA5b")
	wp34 = ns.Waypoint("BlockWPA6b")
	wp35 = ns.Waypoint("BlockWPA1c")
	wp36 = ns.Waypoint("BlockWPA2c")
	wp37 = ns.Waypoint("BlockWPA3c")
	wp38 = ns.Waypoint("BlockWPA4c")
	wp39 = ns.Waypoint("BlockWPA5c")
	wp40 = ns.Waypoint("BlockWPA6c")
	obj44 = ns.Object("BlockB1")
	obj45 = ns.Object("BlockB2")
	obj46 = ns.Object("BlockB3")
	obj47 = ns.Object("BlockB4")
	obj48 = ns.Object("BlockB5")
	wp49 = ns.Waypoint("BlockWPB1b")
	wp50 = ns.Waypoint("BlockWPB2b")
	wp51 = ns.Waypoint("BlockWPB3b")
	wp52 = ns.Waypoint("BlockWPB4b")
	wp53 = ns.Waypoint("BlockWPB5b")
	wp54 = ns.Waypoint("BlockWPB1c")
	wp55 = ns.Waypoint("BlockWPB2c")
	wp56 = ns.Waypoint("BlockWPB3c")
	wp57 = ns.Waypoint("BlockWPB4c")
	wp58 = ns.Waypoint("BlockWPB5c")
	obj62 = ns.Object("BlockC1")
	obj63 = ns.Object("BlockC2")
	obj64 = ns.Object("BlockC3")
	wp65 = ns.Waypoint("BlockWPC1b")
	wp66 = ns.Waypoint("BlockWPC2b")
	wp67 = ns.Waypoint("BlockWPC3b")
	wp68 = ns.Waypoint("BlockWPC1c")
	wp69 = ns.Waypoint("BlockWPC2c")
	wp70 = ns.Waypoint("BlockWPC3c")
	gvar90 = ns.WallGroup("VampireWallGroup")
	obj91 = ns.Object("VampireKnight1")
	obj92 = ns.Object("VKBat1")
	obj93 = ns.Object("RetreatBat")
	wp94 = ns.Waypoint("VampireKnight1WP")
	wp95 = ns.Waypoint("BatCreate")
	wp96 = ns.Waypoint("VKSpot")
	InitializeVampireKnights()
	gvar4 = ns.ObjectGroup("ArrowTrigGroup")
	obj5 = ns.Object("Arrow1")
	obj6 = ns.Object("Arrow2")
	obj7 = ns.Object("Arrow3")
	obj8 = ns.Object("Arrow4")
	obj9 = ns.Object("Arrow5")
	obj10 = ns.Object("Arrow6")
	obj11 = ns.Object("Arrow7")
	obj12 = ns.Object("Arrow8")
	obj13 = ns.Object("Arrow9")
	wp14 = ns.Waypoint("Arrow1WP")
	wp15 = ns.Waypoint("Arrow2WP")
	wp16 = ns.Waypoint("Arrow3WP")
	wp17 = ns.Waypoint("Arrow4WP")
	wp18 = ns.Waypoint("Arrow5WP")
	wp19 = ns.Waypoint("Arrow6WP")
	wp20 = ns.Waypoint("Arrow7WP")
	wp21 = ns.Waypoint("Arrow8WP")
	wp22 = ns.Waypoint("Arrow9WP")
	obj97 = ns.Object("MechGolem1")
	obj98 = ns.Object("MechGolem2")
	obj99 = ns.Object("MechGolem3")
	wp100 = ns.Waypoint("MechGolem2WP")
	wp101 = ns.Waypoint("MechGolemAudioWP")
	gvar102 = ns.ObjectGroup("GolemTrigGroup")
	obj108 = ns.Object("MazeTrig1")
	obj109 = ns.Object("MazeTrig2")
	gvar103 = ns.WallGroup("GolemWallGroup")
	gvar104 = ns.WallGroup("GolemWall2Group")
	gvar105 = ns.WallGroup("GolemWall3Group")
	gvar106 = ns.WallGroup("GolemGate1Group")
	gvar107 = ns.WallGroup("GolemGate2Group")
	gvar110 = ns.WallGroup("CherubBreakWallGroup")
	obj74 = ns.Object("OrbLight")
	obj75 = ns.Object("OrbEffect")
	wp76 = ns.Waypoint("OrbEffectWP")
	wp77 = ns.Waypoint("OrbSoundWP")
	gvar78 = ns.WallGroup("BouncySecretWallGroup")
}
func Secret300XP() {
	ns.GiveXp(ns.GetHost(), 300)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.AudioEvent(ns.SecretFound, wp89)
}
func PlaySubMusic() {
	ns.Music(18, 100)
}
func PlayAction1Music() {
	ns.Music(26, 100)
}
func CherubWallsBreak() {
	ns.WallGroupBreak(gvar110)
	PlaySubMusic()
}
func GolemWallDown() {
	ns.WallGroupOpen(gvar103)
	ns.ObjectOn(obj97)
	ns.AudioEvent(ns.MechGolemPowerUp, wp101)
	ns.ObjectGroupOff(gvar102)
	PlayAction1Music()
}
func GolemWall2Down() {
	ns.NoWallSound(true)
	ns.WallGroupOpen(gvar104)
	ns.NoWallSound(false)
	ns.ObjectOn(obj98)
	ns.AudioEvent(ns.MechGolemRecognize, wp100)
	ns.AudioEvent(ns.MechGolemPowerUp, wp100)
}
func GolemWall3Down() {
	ns.WallGroupOpen(gvar105)
	ns.ObjectOn(obj99)
	ns.ObjectOn(obj88)
}
func OpenGolemGate1() {
	ns.WallGroupOpen(gvar106)
	ns.ObjectOff(obj108)
}
func OpenGolemGate2() {
	ns.WallGroupOpen(gvar107)
	ns.ObjectOff(obj109)
}
func HostilizeMe() {
	ns.AggressionLevel(ns.GetTrigger(), 0.83)
}
func PlayAction3Music() {
	ns.Music(28, 100)
}
func BatToVampireKnight() {
	if flag119 {
		goto LABEL1
	}
	fvar117 = ns.GetObjectX(ns.GetTrigger())
	fvar118 = ns.GetObjectY(ns.GetTrigger())
	ns.Delete(ns.GetTrigger())
	ns.ObjectOn(obj111[0])
LABEL1:
	ns.Effect(ns.BLUE_SPARKS, fvar117, fvar118, 0, 0)
	ns.Effect(ns.SMOKE_BLAST, fvar117, fvar118, 0, 0)
	ns.Enchant(obj111[0], ns.ENCHANT_INVISIBLE, 0.25)
	ns.MoveObject(obj111[0], fvar117, fvar118)
	ns.LookAtObject(obj111[0], ns.GetHost())
}
func SetRetreatBat() {
	ns.AggressionLevel(obj113, 0.83)
	ns.SetCallback(obj113, 7, BatDie)
	ns.SetCallback(obj113, 5, BatDie)
}
func BatDie() {
	var (
		v0 ns.ObjectID
		v1 ns.ObjectID
		v2 ns.ObjectID
	)
	if !(ns.CurrentHealth(obj113) <= 0) {
		goto LABEL1
	}
	ns.Effect(ns.SMOKE_BLAST, ns.GetObjectX(obj113), ns.GetObjectY(obj113), 0, 0)
	ns.MoveWaypoint(wp116, ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger()))
	ns.Delete(obj113)
	ns.AudioEvent(ns.BurnCast, wp116)
	v0 = ns.CreateObject("Flame", wp116)
	v1 = ns.CreateObject("MediumFlame", wp116)
	v2 = ns.CreateObject("SmallFlame", wp116)
	ns.DeleteObjectTimer(v0, 80)
	ns.DeleteObjectTimer(v1, 85)
	ns.DeleteObjectTimer(v2, 90)
	ns.GiveXp(ns.GetHost(), 250)
LABEL1:
	return
}
func ChangeOnSight() {
	if !flag120 {
		goto LABEL1
	}
	ns.AggressionLevel(ns.GetTrigger(), 0.83)
	ns.SetCallback(ns.GetTrigger(), 6, BatToVampireKnight)
LABEL1:
	return
}
func InjureVampireKnight() {
	fvar117 = ns.GetObjectX(ns.GetTrigger())
	fvar118 = ns.GetObjectY(ns.GetTrigger())
	ns.ObjectOn(obj111[0])
	ns.Damage(obj111[0], 0, 50, 0)
	ns.Effect(ns.SMOKE_BLAST, ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger()), 0, 0)
	ns.Delete(ns.GetTrigger())
	flag119 = true
	ns.FrameTimer(1, BatToVampireKnight)
}
func BatMove() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.SetCallback(ns.GetTrigger(), 4, ChangeOnSight)
	ns.Move(obj112[0], wp114)
LABEL1:
	return
}
func VKDie() {
	fvar117 = ns.GetObjectX(ns.GetTrigger())
	fvar118 = ns.GetObjectY(ns.GetTrigger())
	ns.Delete(ns.GetTrigger())
	ns.MoveWaypoint(wp116, fvar117, fvar118)
	obj113 = ns.CreateObject("Bat", wp116)
	ns.FrameTimer(1, SetRetreatBat)
}
func ZombieAmbush() {
	ns.WallGroupOpen(gvar80)
	ns.ObjectGroupOn(gvar79)
}
func SkeletonAttack() {
	ns.WallGroupOpen(gvar81)
	ns.WallGroupClose(gvar82)
	ns.ObjectOn(obj83)
	ns.ObjectGroupOn(gvar85)
	PlayAction3Music()
}
func OpenSkeletonWallGroup2() {
	ns.WallGroupOpen(gvar82)
	ns.ObjectOn(obj84)
}
func Secret300() {
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 300)
}
func OpenSecretMechWallGroup() {
	ns.WallGroupOpen(gvar87)
	ns.GroupAggressionLevel(gvar86, 0.83)
	Secret300()
}
func OnEvent(typ string) {
	switch typ {
	case "PlayerDeath":
		PlayerDeath()
	case "MapInitialize":
		MapInitialize()
	}
}
