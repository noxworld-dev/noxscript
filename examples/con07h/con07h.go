package con07h

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	obj4    ns.ObjectID
	flag5   bool
	wp6     ns.WaypointID
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
	obj29   ns.ObjectID
	obj30   ns.ObjectID
	obj31   ns.ObjectID
	obj32   ns.ObjectID
	obj33   ns.ObjectID
	obj34   ns.ObjectID
	obj35   ns.ObjectID
	obj36   ns.ObjectID
	obj37   ns.ObjectID
	obj38   ns.ObjectID
	obj39   ns.ObjectID
	obj40   ns.ObjectID
	obj41   ns.ObjectID
	obj42   ns.ObjectID
	obj43   ns.ObjectID
	obj44   ns.ObjectID
	obj45   ns.ObjectID
	obj46   ns.ObjectID
	obj47   ns.ObjectID
	obj48   ns.ObjectID
	obj49   ns.ObjectID
	obj50   ns.ObjectID
	obj51   ns.ObjectID
	obj52   ns.ObjectID
	obj53   ns.ObjectID
	gvar54  ns.ObjectGroupID
	gvar55  ns.WallGroupID
	gvar56  ns.WallGroupID
	obj57   ns.ObjectID
	obj58   ns.ObjectID
	obj59   ns.ObjectID
	obj60   ns.ObjectID
	obj61   ns.ObjectID
	obj62   ns.ObjectID
	obj63   ns.ObjectID
	obj64   ns.ObjectID
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
	gvar82  int
	gvar83  ns.ObjectGroupID
	gvar84  ns.ObjectGroupID
	gvar85  ns.ObjectGroupID
	gvar86  ns.ObjectGroupID
	gvar87  ns.ObjectGroupID
	gvar88  ns.ObjectGroupID
	gvar89  ns.ObjectGroupID
	gvar90  ns.ObjectGroupID
	gvar91  ns.ObjectGroupID
	gvar92  ns.ObjectGroupID
	flag93  bool
	wp94    ns.WaypointID
	wp95    ns.WaypointID
	wp96    ns.WaypointID
	gvar97  ns.TimerID
	ivar98  int
	ivar99  int
	fvar100 float32
	fvar101 float32
	fvar102 float32
	fvar103 float32
	fvar104 float32
	fvar105 float32
	fvar106 float32
	fvar107 float32
	fvar108 float32
	fvar109 float32
	fvar110 float32
	fvar111 float32
	fvar112 float32
	fvar113 float32
	fvar114 float32
	fvar115 float32
	fvar116 float32
	fvar117 float32
	fvar118 float32
	fvar119 float32
	flag120 bool
	flag121 bool
	flag122 bool
	flag123 bool
	flag124 bool
)

func init() {
	flag5 = false
	flag93 = false
	ivar98 = 0
	ivar99 = 0
	flag120 = false
	flag121 = false
	flag122 = false
	flag123 = false
	flag124 = false
}
func CreatureSetup() {
	ns.SetOwner(obj7, obj12)
	ns.SetOwner(obj7, obj13)
	fvar100 = ns.GetObjectX(obj33)
	fvar101 = ns.GetObjectX(obj34)
	fvar102 = ns.GetObjectX(obj35)
	fvar103 = ns.GetObjectX(obj36)
	fvar104 = ns.GetObjectX(obj37)
	fvar105 = ns.GetObjectX(obj38)
	fvar106 = ns.GetObjectX(obj39)
	fvar107 = ns.GetObjectX(obj40)
	fvar108 = ns.GetObjectX(obj41)
	fvar109 = ns.GetObjectX(obj42)
	fvar110 = ns.GetObjectY(obj33)
	fvar111 = ns.GetObjectY(obj34)
	fvar112 = ns.GetObjectY(obj35)
	fvar113 = ns.GetObjectY(obj36)
	fvar114 = ns.GetObjectY(obj37)
	fvar115 = ns.GetObjectY(obj38)
	fvar116 = ns.GetObjectY(obj39)
	fvar117 = ns.GetObjectY(obj40)
	fvar118 = ns.GetObjectY(obj41)
	fvar119 = ns.GetObjectY(obj42)
}
func StartChamber() {
	ns.ObjectGroupOn(gvar84)
	ns.ObjectGroupOn(gvar86)
	ns.ObjectGroupOff(gvar83)
	ns.ObjectGroupOff(gvar85)
	flag93 = true
	ns.SecondTimer(3, ChamberSpikes)
}
func ChamberSpikes() {
	if flag93 {
		goto LABEL1
	}
	ns.ObjectGroupOff(gvar83)
	ns.ObjectGroupOff(gvar85)
	ns.FrameTimer(10, EnableSpikes2And4)
	goto LABEL2
LABEL1:
	ns.ObjectGroupOff(gvar84)
	ns.ObjectGroupOff(gvar86)
	ns.FrameTimer(10, EnableSpikes1And3)
LABEL2:
	gvar97 = ns.SecondTimer(3, ChamberSpikes)
}
func Block1Jiggle() {
	ns.MoveWaypoint(wp6, ns.GetObjectX(obj80), ns.GetObjectY(obj80))
	ns.AudioEvent(ns.HammerMissing, wp6)
	ns.Effect(ns.JIGGLE, ns.GetObjectX(obj80), ns.GetObjectY(obj80), 10, 0)
}
func Block2Jiggle() {
	ns.AudioEvent(ns.HammerMissing, wp94)
	ns.Effect(ns.JIGGLE, ns.GetObjectX(obj81), ns.GetObjectY(obj81), 10, 0)
}
func EnableSpikes2And4() {
	ns.ObjectGroupOn(gvar84)
	ns.ObjectGroupOn(gvar86)
	flag93 = true
}
func EnableSpikes1And3() {
	ns.ObjectGroupOn(gvar83)
	ns.ObjectGroupOn(gvar85)
	flag93 = false
}
func DemonDie() {
	ivar99 += 1
	if !(ivar99 >= 3) {
		goto LABEL1
	}
	ns.GiveXp(ns.GetHost(), 1000)
	ns.Music(1, 100)
LABEL1:
	return
}
func Horvath1Start() {
	ns.TellStory(ns.HumanMaleEatFood, "Con07H.scr:HorvathTalk01")
}
func Horvath1End() {
	ns.SetDialog(obj7, ns.NORMAL, Hecubah2Start, Hecubah2End)
	ns.StartDialog(obj7, ns.GetHost())
}
func Hecubah1Start() {
	ns.TellStory(ns.HumanMaleEatFood, "Con07H.scr:HecubahTalk01")
}
func Hecubah1End() {
	ns.SetDialog(obj7, ns.NEXT, Hecubah2Start, Hecubah2End)
	HorvathSpeak()
}
func MapInitialize() {
	obj57 = ns.Object("SnakeHead1Mover")
	obj58 = ns.Object("SnakeHead2Mover")
	obj59 = ns.Object("SnakeHead3Mover")
	obj60 = ns.Object("SnakeHead4Mover")
	obj61 = ns.Object("SnakeHead5Mover")
	obj62 = ns.Object("SnakeHead6Mover")
	obj63 = ns.Object("SnakeHead7Mover")
	obj64 = ns.Object("SnakeHead8Mover")
	obj65 = ns.Object("SnakeHead9Mover")
	obj66 = ns.Object("SnakeHead10Mover")
	obj67 = ns.Object("SnakeHead11Mover")
	obj68 = ns.Object("SnakeHead12Mover")
	obj69 = ns.Object("Block1AMover")
	obj70 = ns.Object("Block1BMover")
	obj71 = ns.Object("Block2AMover")
	obj72 = ns.Object("Block2BMover")
	obj4 = ns.Object("ToLevelOneTP")
	obj79 = ns.Object("ToThreeTP")
	obj7 = ns.Object("Hecubah")
	obj12 = ns.Object("Lord1")
	obj13 = ns.Object("Lord2")
	obj17 = ns.Object("DeadGuy1")
	obj18 = ns.Object("DeadGuy2")
	obj8 = ns.Object("HecDemon1")
	obj9 = ns.Object("HecDemon2")
	obj10 = ns.Object("HecDemon3")
	obj11 = ns.Object("HecDemon4")
	obj14 = ns.Object("Horvath")
	obj15 = ns.Object("Marik")
	obj16 = ns.Object("Mage2")
	obj80 = ns.Object("Block1Trigger")
	obj81 = ns.Object("Block2Trigger")
	obj73 = ns.Object("ChamberDoor1")
	obj74 = ns.Object("ChamberDoor2")
	obj75 = ns.Object("ChamberDoor2A")
	obj76 = ns.Object("ChamberDoor2B")
	obj77 = ns.Object("ChamberExitDoor1")
	obj78 = ns.Object("ChamberExitDoor2")
	gvar83 = ns.ObjectGroup("ChamberSpikes1")
	gvar84 = ns.ObjectGroup("ChamberSpikes2")
	gvar85 = ns.ObjectGroup("ChamberSpikes3")
	gvar86 = ns.ObjectGroup("ChamberSpikes4")
	gvar87 = ns.ObjectGroup("SnakeLetterMovers")
	gvar88 = ns.ObjectGroup("Snake3Movers")
	gvar89 = ns.ObjectGroup("Snake4Movers")
	gvar90 = ns.ObjectGroup("ShowdownTriggers")
	wp6 = ns.Waypoint("PlayerSounds")
	wp94 = ns.Waypoint("Jiggle2WP")
	wp95 = ns.Waypoint("Traitor1WP")
	wp96 = ns.Waypoint("Traitor2WP")
	wp19 = ns.Waypoint("HecDemon1WP")
	wp20 = ns.Waypoint("HecDemon2WP")
	wp21 = ns.Waypoint("Lord1WP")
	wp22 = ns.Waypoint("Lord2WP")
	obj23 = ns.Object("HON")
	obj24 = ns.Object("HONBase")
	obj25 = ns.Object("PowerCell1")
	obj26 = ns.Object("PowerCell2")
	obj27 = ns.Object("PowerCell3")
	obj28 = ns.Object("PowerCell4")
	obj29 = ns.Object("PowerCell5")
	obj30 = ns.Object("PowerCell6")
	obj31 = ns.Object("PowerCell7")
	obj32 = ns.Object("PowerCell8")
	obj33 = ns.Object("HeartLight1")
	obj34 = ns.Object("HeartLight2")
	obj35 = ns.Object("HeartLight3")
	obj36 = ns.Object("HeartLight4")
	obj37 = ns.Object("HeartLight5")
	obj38 = ns.Object("HeartLight6")
	obj39 = ns.Object("HeartLight7")
	obj40 = ns.Object("HeartLight8")
	obj41 = ns.Object("HeartLight9")
	obj42 = ns.Object("HeartLight10")
	obj43 = ns.Object("NewLight1")
	obj44 = ns.Object("NewLight2")
	obj45 = ns.Object("NewLight3")
	obj46 = ns.Object("NewLight4")
	obj47 = ns.Object("NewLight5")
	obj48 = ns.Object("NewLight6")
	obj49 = ns.Object("NewLight7")
	obj50 = ns.Object("NewLight8")
	obj51 = ns.Object("NewLight9")
	obj52 = ns.Object("NewLight10")
	obj53 = ns.Object("Rain")
	gvar54 = ns.ObjectGroup("PowerCells")
	gvar91 = ns.ObjectGroup("TraitorTriggers")
	gvar92 = ns.ObjectGroup("NewGroup")
	gvar55 = ns.WallGroup("HONWalls")
	gvar56 = ns.WallGroup("HONExitWalls")
	ns.StoryPic(obj14, "HorvathPic")
	ns.StoryPic(obj7, "HecubahPic")
	ns.SetDialog(obj7, ns.NEXT, Hecubah1Start, Hecubah1End)
	ns.FrameTimer(1, CreatureSetup)
}
func PlayerDeath() {
	ns.DeathScreen(7)
}
func LightShow() {
	var (
		v0 int
		v1 int
	)
	ns.MoveWaypoint(wp6, ns.GetObjectX(obj23), ns.GetObjectY(obj23))
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
	ns.AudioEvent(ns.LightningBolt, wp6)
	ns.Effect(ns.LIGHTNING, ns.GetObjectX(obj25), ns.GetObjectY(obj25), ns.GetObjectX(obj23), ns.GetObjectY(obj23))
	goto LABEL9
LABEL2:
	ns.AudioEvent(ns.LightningBolt, wp6)
	ns.Effect(ns.LIGHTNING, ns.GetObjectX(obj26), ns.GetObjectY(obj26), ns.GetObjectX(obj23), ns.GetObjectY(obj23))
	goto LABEL9
LABEL3:
	ns.AudioEvent(ns.LightningBolt, wp6)
	ns.Effect(ns.LIGHTNING, ns.GetObjectX(obj27), ns.GetObjectY(obj27), ns.GetObjectX(obj23), ns.GetObjectY(obj23))
	goto LABEL9
LABEL4:
	ns.AudioEvent(ns.LightningBolt, wp6)
	ns.Effect(ns.LIGHTNING, ns.GetObjectX(obj28), ns.GetObjectY(obj28), ns.GetObjectX(obj23), ns.GetObjectY(obj23))
	goto LABEL9
LABEL5:
	ns.AudioEvent(ns.LightningBolt, wp6)
	ns.Effect(ns.LIGHTNING, ns.GetObjectX(obj29), ns.GetObjectY(obj29), ns.GetObjectX(obj23), ns.GetObjectY(obj23))
	goto LABEL9
LABEL6:
	ns.AudioEvent(ns.LightningBolt, wp6)
	ns.Effect(ns.LIGHTNING, ns.GetObjectX(obj30), ns.GetObjectY(obj30), ns.GetObjectX(obj23), ns.GetObjectY(obj23))
	goto LABEL9
LABEL7:
	ns.AudioEvent(ns.LightningBolt, wp6)
	ns.Effect(ns.LIGHTNING, ns.GetObjectX(obj31), ns.GetObjectY(obj31), ns.GetObjectX(obj23), ns.GetObjectY(obj23))
	goto LABEL9
LABEL8:
	ns.AudioEvent(ns.LightningBolt, wp6)
	ns.Effect(ns.LIGHTNING, ns.GetObjectX(obj32), ns.GetObjectY(obj32), ns.GetObjectX(obj23), ns.GetObjectY(obj23))
	goto LABEL9
LABEL9:
	if flag123 {
		goto LABEL10
	}
	ns.FrameTimer(30, LightShow)
LABEL10:
	return
}
func PlateOneTrigger() {
	flag120 = true
	if !(flag121 == true && flag122 == true) {
		goto LABEL1
	}
	ns.WallGroupOpen(gvar55)
LABEL1:
	return
}
func PlateTwoTrigger() {
	flag121 = true
	if !(flag120 == true && flag122 == true) {
		goto LABEL1
	}
	ns.WallGroupOpen(gvar55)
LABEL1:
	return
}
func PlateThreeTrigger() {
	flag122 = true
	if !(flag120 == true && flag121 == true) {
		goto LABEL1
	}
	ns.WallGroupOpen(gvar55)
LABEL1:
	return
}
func PlateOneRelease() {
	flag120 = false
}
func PlateTwoRelease() {
	flag121 = false
}
func PlateThreeRelease() {
	flag122 = false
}
func HONTrigger() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.ObjectOff(obj33)
	ns.ObjectOff(obj34)
	ns.ObjectOff(obj35)
	ns.ObjectOff(obj36)
	ns.ObjectOff(obj37)
	ns.ObjectOff(obj38)
	ns.ObjectOff(obj39)
	ns.ObjectOff(obj40)
	ns.ObjectOff(obj41)
	ns.ObjectOff(obj42)
	ns.MoveObject(obj43, fvar100, fvar110)
	ns.MoveObject(obj44, fvar101, fvar111)
	ns.MoveObject(obj45, fvar102, fvar112)
	ns.MoveObject(obj46, fvar103, fvar113)
	ns.MoveObject(obj47, fvar104, fvar114)
	ns.MoveObject(obj48, fvar105, fvar115)
	ns.MoveObject(obj49, fvar106, fvar116)
	ns.MoveObject(obj50, fvar107, fvar117)
	ns.MoveObject(obj51, fvar108, fvar118)
	ns.MoveObject(obj52, fvar109, fvar119)
	ns.ObjectOff(obj23)
	ns.ObjectOff(obj24)
	flag123 = true
	ns.SetHalberd(1)
	ns.PrintToAll("Con07H.scr:HONTada")
	ns.JournalEdit(ns.GetHost(), "GetHON", 4)
	ns.SecondTimer(2, FadeOut)
LABEL1:
	return
}
func MakeMapSwitch() {
	ns.MoveObject(ns.GetHost(), 2149, 2151)
}
func FadeOut() {
	ns.Blind()
	ns.Music(0, 100)
	ns.FrameTimer(45, MakeMapSwitch)
}
func BeginShowdownPiece() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	if flag5 {
		goto LABEL1
	}
	flag5 = true
	ns.WideScreen(true)
	ns.Music(7, 100)
	ns.SetDialog(obj14, ns.NEXT, Horvath1Start, Horvath1End)
	ns.MoveWaypoint(wp6, 2133, 1410)
	ns.AudioEvent(ns.TeleportIn, wp6)
	ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp6), ns.GetWaypointY(wp6), 0, 0)
	ns.Effect(ns.SPARK_EXPLOSION, ns.GetWaypointX(wp6), ns.GetWaypointY(wp6), 0, 0)
	ns.Effect(ns.LESSER_EXPLOSION, ns.GetWaypointX(wp6), ns.GetWaypointY(wp6), 0, 0)
	ns.MoveObject(obj7, ns.GetWaypointX(wp6), ns.GetWaypointY(wp6))
	ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp21), ns.GetWaypointY(wp21), 0, 0)
	ns.Effect(ns.SPARK_EXPLOSION, ns.GetWaypointX(wp21), ns.GetWaypointY(wp21), 0, 0)
	ns.Effect(ns.LESSER_EXPLOSION, ns.GetWaypointX(wp21), ns.GetWaypointY(wp21), 0, 0)
	ns.MoveObject(obj12, ns.GetWaypointX(wp21), ns.GetWaypointY(wp21))
	ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp22), ns.GetWaypointY(wp22), 0, 0)
	ns.Effect(ns.SPARK_EXPLOSION, ns.GetWaypointX(wp22), ns.GetWaypointY(wp22), 0, 0)
	ns.Effect(ns.LESSER_EXPLOSION, ns.GetWaypointX(wp22), ns.GetWaypointY(wp22), 0, 0)
	ns.MoveObject(obj13, ns.GetWaypointX(wp22), ns.GetWaypointY(wp22))
	ns.Frozen(ns.GetHost(), true)
	ns.CreatureGuard(obj7, ns.GetObjectX(obj7), ns.GetObjectY(obj7), ns.GetObjectX(obj14), ns.GetObjectY(obj14), 0)
	ns.Frozen(obj7, true)
	ns.Frozen(obj12, true)
	ns.Frozen(obj13, true)
	ns.LookAtObject(obj7, obj14)
	ns.ObjectGroupOff(gvar92)
	ns.SecondTimer(1, HorvathSpeak)
LABEL1:
	return
}
func HorvathSpeak() {
	ns.StartDialog(obj14, ns.GetHost())
}
func HecubahSpeak() {
	ns.StartDialog(obj7, ns.GetHost())
}
func BeginFireworks() {
	ns.ObjectOn(obj7)
	ns.ObjectOn(obj14)
	ns.ObjectOn(obj15)
	ns.ObjectOn(obj16)
}
func HorvathDied() {
	ns.Frozen(obj14, true)
	ns.SetDialog(obj14, ns.NEXT, Horvath2Start, Horvath2End)
	ns.StartDialog(obj14, ns.GetHost())
}
func Horvath2Start() {
	ns.TellStory(ns.HumanMaleEatFood, "Con07H.scr:HorvathTalk02")
}
func Horvath2End() {
	ns.Frozen(obj14, false)
	ns.CancelDialog(obj14)
	ns.CastSpellLocationLocation(ns.SPELL_TURN_UNDEAD, ns.GetObjectX(obj14), ns.GetObjectY(obj14), ns.GetObjectX(obj14), ns.GetObjectY(obj14))
	ns.Frozen(obj12, false)
	ns.Frozen(obj13, false)
	ns.ObjectOn(obj15)
	ns.ObjectOn(obj16)
	ns.ObjectOn(obj12)
	ns.ObjectOn(obj13)
	ns.HitFarLocation(obj12, ns.GetObjectX(obj15), ns.GetObjectY(obj15))
	ns.HitFarLocation(obj13, ns.GetObjectX(obj16), ns.GetObjectY(obj16))
	ns.SecondTimer(2, EraseHorvath)
	ns.SecondTimer(1, HecKill1)
}
func Hecubah2Start() {
	ns.TellStory(ns.HumanMaleEatFood, "Con07H.scr:HecubahTalk03")
}
func Hecubah2End() {
	ns.ObjectOn(obj14)
	ns.ObjectOn(obj7)
	ns.CastSpellLocationLocation(ns.SPELL_DEATH_RAY, ns.GetObjectX(obj7), ns.GetObjectY(obj7), ns.GetObjectX(obj14), ns.GetObjectY(obj14))
	ns.SecondTimer(1, SecondDeathRay)
}
func Hecubah3Start() {
	ns.TellStory(ns.HumanMaleEatFood, "Con07H.scr:HecubahTalk04")
}
func Hecubah3End() {
	ns.MoveWaypoint(wp6, ns.GetObjectX(obj7), ns.GetObjectY(obj7))
	ns.AudioEvent(ns.TeleportOut, wp6)
	ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp6), ns.GetWaypointY(wp6), 0, 0)
	ns.Effect(ns.SPARK_EXPLOSION, ns.GetWaypointX(wp6), ns.GetWaypointY(wp6), 0, 0)
	ns.Effect(ns.LESSER_EXPLOSION, ns.GetWaypointX(wp6), ns.GetWaypointY(wp6), 0, 0)
	ns.Delete(obj7)
	ns.SetDialog(obj15, ns.NORMAL, TestStart, TestEnd)
	ns.WideScreen(false)
	ns.Frozen(ns.GetHost(), false)
	ns.StoryPic(obj15, "WizardGuard1Pic")
	ns.SetDialog(obj15, ns.NORMAL, MarikTalkStart, MarikTalkEnd)
	MarikTalk()
}
func HecubahPoof() {
	ns.StartDialog(obj7, ns.GetHost())
}
func TestStart() {
	ns.Delete(obj8)
}
func TestEnd() {
	ns.CancelDialog(obj15)
}
func MarikTalk() {
	if !(ns.Distance(ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), ns.GetObjectX(obj15), ns.GetObjectY(obj15)) < 150) {
		goto LABEL1
	}
	if flag124 {
		goto LABEL2
	}
	ns.StartDialog(obj15, ns.GetHost())
	flag124 = true
LABEL2:
	goto LABEL3
LABEL1:
	if flag124 {
		goto LABEL3
	}
	ns.FrameTimer(8, MarikTalk)
LABEL3:
	return
}
func MarikTalkStart() {
	ns.TellStory(ns.HumanMaleEatFood, "Con07H.scr:MarikTalk02")
	flag124 = true
}
func MarikTalkEnd() {
	ns.SetDialog(obj15, ns.NORMAL, MarikTalkStart, MarikTalkEnd)
}
func SecondDeathRay() {
	ns.CastSpellLocationLocation(ns.SPELL_DEATH_RAY, ns.GetObjectX(obj7), ns.GetObjectY(obj7), ns.GetObjectX(obj14), ns.GetObjectY(obj14))
	ns.SecondTimer(1, ThirdDeathRay)
}
func ThirdDeathRay() {
	ns.CastSpellLocationLocation(ns.SPELL_DEATH_RAY, ns.GetObjectX(obj7), ns.GetObjectY(obj7), ns.GetObjectX(obj14), ns.GetObjectY(obj14))
}
func Lord1Missile() {
	ns.HitFarLocation(obj12, ns.GetObjectX(obj15), ns.GetObjectY(obj15))
}
func Lord2Missile() {
	ns.HitFarLocation(obj13, ns.GetObjectX(obj15), ns.GetObjectY(obj15))
}
func OgreDie() {
	ivar98 += 1
	if !(ivar98 >= 2) {
		goto LABEL1
	}
	ns.AggressionLevel(obj15, 0)
	ns.AggressionLevel(obj16, 0)
	ns.CreatureGuard(obj15, ns.GetObjectX(obj15), ns.GetObjectY(obj15), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	ns.CreatureGuard(obj16, ns.GetObjectX(obj16), ns.GetObjectY(obj16), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	ns.SetDialog(obj7, ns.NEXT, Hecubah3Start, Hecubah3End)
	ns.LookAtObject(obj7, ns.GetHost())
	ns.SecondTimer(1, HecubahPoof)
LABEL1:
	return
}
func HecKill1() {
	ns.CastSpellLocationLocation(ns.SPELL_DEATH_RAY, ns.GetObjectX(obj7), ns.GetObjectY(obj7), ns.GetObjectX(obj18), ns.GetObjectY(obj18))
	ns.SecondTimer(2, HecKill2)
}
func HecKill2() {
	ns.CastSpellLocationLocation(ns.SPELL_DEATH_RAY, ns.GetObjectX(obj7), ns.GetObjectY(obj7), ns.GetObjectX(obj17), ns.GetObjectY(obj17))
}
func EraseHorvath() {
	ns.Delete(obj14)
}
func OnEvent(typ string) {
	switch typ {
	case "MapInitialize":
		MapInitialize()
	case "PlayerDeath":
		PlayerDeath()
	}
}
