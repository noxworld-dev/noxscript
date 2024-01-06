package war02a

import (
	"strconv"

	"github.com/noxworld-dev/noxscript/ns/v3"
)

var (
	ivar4   int
	flag5   bool
	flag6   bool
	obj7    ns.ObjectID
	obj8    ns.ObjectID
	gvar9   ns.ObjectGroupID
	wp10    ns.WaypointID
	flag11  bool
	flag12  bool
	gvar13  ns.ObjectGroupID
	gvar14  ns.ObjectGroupID
	gvar15  ns.ObjectGroupID
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
	obj54   [40]ns.ObjectID
	wp55    ns.WaypointID
	wp56    ns.WaypointID
	wp57    ns.WaypointID
	wp58    ns.WaypointID
	wp59    ns.WaypointID
	wp60    ns.WaypointID
	wp61    ns.WaypointID
	wp62    ns.WaypointID
	wp63    ns.WaypointID
	wp64    ns.WaypointID
	wp65    ns.WaypointID
	wp66    ns.WaypointID
	wp67    ns.WaypointID
	wp68    ns.WaypointID
	wp69    ns.WaypointID
	wp70    ns.WaypointID
	wp71    ns.WaypointID
	wp72    ns.WaypointID
	wp73    ns.WaypointID
	wp74    ns.WaypointID
	wp75    ns.WaypointID
	wp76    ns.WaypointID
	wp77    ns.WaypointID
	wp78    ns.WaypointID
	wp79    ns.WaypointID
	wp80    ns.WaypointID
	wp81    ns.WaypointID
	wp82    ns.WaypointID
	wp83    ns.WaypointID
	wp84    ns.WaypointID
	wp85    ns.WaypointID
	wp86    ns.WaypointID
	wp87    ns.WaypointID
	wp88    ns.WaypointID
	wp89    ns.WaypointID
	ivar90  int
	obj91   ns.ObjectID
	obj92   ns.ObjectID
	obj93   ns.ObjectID
	obj94   ns.ObjectID
	obj95   ns.ObjectID
	obj96   ns.ObjectID
	obj97   ns.ObjectID
	gvar98  ns.WallGroupID
	gvar99  int
	gvar100 int
	gvar101 int
	gvar102 int
	gvar103 int
	gvar104 int
	gvar105 int
	gvar106 int
	gvar107 int
	gvar108 int
	gvar109 int
	gvar110 int
	gvar111 int
	gvar112 int
	gvar113 int
	gvar114 int
	gvar115 int
	gvar116 int
	gvar117 int
	gvar118 int
	gvar119 int
	gvar120 int
	gvar121 int
	gvar122 int
	gvar123 int
	gvar124 int
	gvar125 int
	gvar126 int
	gvar127 int
	gvar128 int
	gvar129 int
	gvar130 int
	gvar131 int
	gvar132 int
	gvar133 int
	gvar134 int
	gvar135 int
	gvar136 int
	gvar137 int
	gvar138 int
	flag139 bool
	flag140 bool
	flag141 bool
	flag142 bool
	flag143 bool
	flag144 bool
	flag145 bool
	flag146 bool
	flag147 bool
	flag148 bool
	flag149 bool
	ivar150 int
	ivar151 int
	ivar152 int
	ivar153 int
	ivar154 int
	ivar155 int
	ivar156 int
	ivar157 int
	ivar158 int
	ivar159 int
	ivar160 int
	ivar161 int
	ivar162 int
	ivar163 int
	ivar164 int
	ivar165 int
	gvar166 int
	gvar167 int
	gvar168 int
	gvar169 int
	gvar170 int
	gvar171 int
	gvar172 int
	gvar173 int
	gvar174 int
	gvar175 int
	gvar176 int
	gvar177 int
	gvar178 int
	gvar179 int
	gvar180 int
	gvar181 int
	gvar182 int
	gvar183 int
	gvar184 int
	gvar185 int
	obj186  [8]ns.ObjectID
	obj187  [8]ns.ObjectID
	obj188  [8]ns.ObjectID
	obj189  [8]ns.ObjectID
	obj190  [8]ns.ObjectID
	obj191  [8]ns.ObjectID
	obj192  [8]ns.ObjectID
	obj193  [8]ns.ObjectID
	obj194  [8]ns.ObjectID
	obj195  [8]ns.ObjectID
	wp196   [8]ns.WaypointID
	wp197   [8]ns.WaypointID
	wp198   [8]ns.WaypointID
	wp199   [8]ns.WaypointID
	wp200   [8]ns.WaypointID
	wp201   [8]ns.WaypointID
	wp202   [8]ns.WaypointID
	wp203   [8]ns.WaypointID
	wp204   [8]ns.WaypointID
	wp205   [8]ns.WaypointID
	ivar206 int
	obj207  ns.ObjectID
	obj208  ns.ObjectID
	obj209  ns.ObjectID
	obj210  ns.ObjectID
	obj211  ns.ObjectID
	obj212  ns.ObjectID
	obj213  ns.ObjectID
	gvar214 ns.WallGroupID
	flag215 bool
	flag216 bool
	gvar217 int
	gvar218 int
	ivar219 int
	obj220  ns.ObjectID
	obj221  ns.ObjectID
	obj222  ns.ObjectID
	obj223  ns.ObjectID
	obj224  ns.ObjectID
	obj225  ns.ObjectID
	obj226  ns.ObjectID
	obj227  ns.ObjectID
	gvar228 ns.WallGroupID
	ivar229 int
	obj230  ns.ObjectID
	obj231  ns.ObjectID
	obj232  ns.ObjectID
	gvar233 ns.ObjectGroupID
	gvar234 ns.ObjectGroupID
	gvar235 ns.WallGroupID
	wp236   ns.WaypointID
	flag237 bool
	flag238 bool
	gvar239 ns.ObjectGroupID
	obj240  ns.ObjectID
	obj241  ns.ObjectID
	obj242  ns.ObjectID
	obj243  ns.ObjectID
	obj244  ns.ObjectID
	obj245  ns.ObjectID
	obj246  ns.ObjectID
	obj247  ns.ObjectID
	obj248  ns.ObjectID
	obj249  ns.ObjectID
	obj250  ns.ObjectID
	obj251  ns.ObjectID
	obj252  ns.ObjectID
	obj253  ns.ObjectID
	obj254  ns.ObjectID
	obj255  ns.ObjectID
	obj256  ns.ObjectID
	obj257  ns.ObjectID
	obj258  ns.ObjectID
	obj259  ns.ObjectID
	obj260  ns.ObjectID
	obj261  ns.ObjectID
	wp262   ns.WaypointID
	wp263   ns.WaypointID
	wp264   ns.WaypointID
	wp265   ns.WaypointID
	wp266   ns.WaypointID
	wp267   ns.WaypointID
	wp268   ns.WaypointID
	wp269   ns.WaypointID
	wp270   ns.WaypointID
	wp271   ns.WaypointID
	wp272   ns.WaypointID
	wp273   ns.WaypointID
	wp274   ns.WaypointID
	wp275   ns.WaypointID
	wp276   ns.WaypointID
	wp277   ns.WaypointID
	wp278   ns.WaypointID
	wp279   ns.WaypointID
	wp280   ns.WaypointID
	wp281   ns.WaypointID
	flag282 bool
	obj283  ns.ObjectID
	obj284  ns.ObjectID
	gvar285 ns.ObjectGroupID
)

func init() {
	flag5 = true
	flag6 = false
	flag11 = true
	flag12 = true
	ivar90 = 0
	gvar99 = 0
	gvar100 = 1
	gvar101 = 2
	gvar102 = 3
	gvar103 = 0
	gvar104 = 1
	gvar105 = 2
	gvar106 = 3
	gvar107 = 0
	gvar108 = 1
	gvar109 = 2
	gvar110 = 3
	gvar111 = 0
	gvar112 = 1
	gvar113 = 2
	gvar114 = 3
	gvar115 = 0
	gvar116 = 1
	gvar117 = 2
	gvar118 = 3
	gvar119 = 0
	gvar120 = 1
	gvar121 = 2
	gvar122 = 3
	gvar123 = 0
	gvar124 = 1
	gvar125 = 2
	gvar126 = 3
	gvar127 = 0
	gvar128 = 1
	gvar129 = 2
	gvar130 = 3
	gvar131 = 0
	gvar132 = 1
	gvar133 = 2
	gvar134 = 3
	gvar135 = 0
	gvar136 = 1
	gvar137 = 2
	gvar138 = 3
	flag139 = true
	flag140 = true
	flag141 = true
	flag142 = true
	flag143 = true
	flag144 = true
	flag145 = true
	flag146 = true
	flag147 = true
	flag148 = true
	flag149 = true
	gvar166 = gvar100
	gvar167 = gvar104
	gvar168 = gvar108
	gvar169 = gvar112
	gvar170 = gvar116
	gvar171 = gvar120
	gvar172 = gvar124
	gvar173 = gvar128
	gvar174 = gvar132
	gvar175 = gvar136
	gvar176 = gvar99
	gvar177 = gvar103
	gvar178 = gvar107
	gvar179 = gvar111
	gvar180 = gvar115
	gvar181 = gvar119
	gvar182 = gvar123
	gvar183 = gvar127
	gvar184 = gvar131
	gvar185 = gvar135
	ivar206 = 0
	flag215 = true
	flag216 = true
	ivar219 = 0
	ivar229 = 0
	flag237 = true
	flag238 = false
	flag282 = false
}
func secretTrigger() {
	ns.MoveWaypoint(wp55, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretFound, wp55)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 150)
}
func PlayGauntletMusic() {
	ns.Music(29, 100)
}
func initGauntletDoor() {
	gvar14 = ns.ObjectGroup("GauntletGears")
	ns.ObjectGroupOff(gvar14)
	gvar13 = ns.ObjectGroup("GauntletDoorMovers")
	obj16 = ns.Object("B1mover")
	obj17 = ns.Object("B2mover")
	obj18 = ns.Object("B3mover")
	obj19 = ns.Object("B4mover")
	obj20 = ns.Object("B5mover")
	obj21 = ns.Object("B6mover")
	obj22 = ns.Object("B7mover")
	obj23 = ns.Object("B8mover")
	obj24 = ns.Object("B9mover")
	obj25 = ns.Object("B10mover")
	obj26 = ns.Object("F1mover")
	obj27 = ns.Object("F2mover")
	obj28 = ns.Object("F3mover")
	obj29 = ns.Object("F4mover")
	obj30 = ns.Object("F5mover")
	obj31 = ns.Object("F6mover")
	obj32 = ns.Object("F7mover")
	obj33 = ns.Object("F8mover")
	obj34 = ns.Object("F9mover")
	obj35 = ns.Object("F10mover")
	wp56 = ns.Waypoint("B1way")
	wp57 = ns.Waypoint("B2way")
	wp58 = ns.Waypoint("B3way")
	wp59 = ns.Waypoint("B4way")
	wp60 = ns.Waypoint("B5way")
	wp61 = ns.Waypoint("B6way")
	wp62 = ns.Waypoint("B7way")
	wp63 = ns.Waypoint("B8way")
	wp64 = ns.Waypoint("B9way")
	wp65 = ns.Waypoint("B10way")
	wp66 = ns.Waypoint("B11way")
	wp67 = ns.Waypoint("B12way")
	wp68 = ns.Waypoint("B13way")
	wp69 = ns.Waypoint("F1way")
	wp70 = ns.Waypoint("F2way")
	wp71 = ns.Waypoint("F3way")
	wp72 = ns.Waypoint("F4way")
	wp73 = ns.Waypoint("F5way")
	wp74 = ns.Waypoint("F6way")
	wp75 = ns.Waypoint("F7way")
	wp76 = ns.Waypoint("F8way")
	wp77 = ns.Waypoint("F9way")
	wp78 = ns.Waypoint("F10way")
	wp79 = ns.Waypoint("F11way")
	wp80 = ns.Waypoint("F12way")
	wp81 = ns.Waypoint("F13way")
	ns.ObjectGroupOff(gvar233)
	ns.LockDoor(obj283)
	ns.LockDoor(obj284)
}
func gauntletGearsStop() {
	ns.ObjectGroupOff(gvar14)
}
func openGauntlet() {
	if flag6 {
		goto LABEL1
	}
	flag6 = true
	ns.ObjectGroupOn(gvar14)
	ns.FrameTimer(75, gauntletGearsStop)
	ns.ObjectGroupOn(gvar13)
	ns.Move(obj16, wp56)
	ns.Move(obj17, wp57)
	ns.Move(obj18, wp58)
	ns.Move(obj19, wp59)
	ns.Move(obj20, wp60)
	ns.Move(obj21, wp61)
	ns.Move(obj22, wp63)
	ns.Move(obj23, wp64)
	ns.Move(obj24, wp65)
	ns.Move(obj25, wp66)
	ns.Move(obj26, wp71)
	ns.Move(obj27, wp72)
	ns.Move(obj28, wp73)
	ns.Move(obj29, wp74)
	ns.Move(obj30, wp76)
	ns.Move(obj31, wp77)
	ns.Move(obj32, wp78)
	ns.Move(obj33, wp79)
	ns.Move(obj34, wp80)
	ns.Move(obj35, wp81)
	ns.FrameTimer(3, stoneGrinder)
LABEL1:
	return
}
func closeGauntlet() {
	if !flag6 {
		goto LABEL1
	}
	flag6 = false
	ns.ObjectGroupOn(gvar14)
	ns.FrameTimer(75, gauntletGearsStop)
	ns.Move(obj16, wp58)
	ns.Move(obj17, wp59)
	ns.Move(obj18, wp60)
	ns.Move(obj19, wp61)
	ns.Move(obj20, wp62)
	ns.Move(obj21, wp63)
	ns.Move(obj22, wp65)
	ns.Move(obj23, wp66)
	ns.Move(obj24, wp67)
	ns.Move(obj25, wp68)
	ns.Move(obj26, wp69)
	ns.Move(obj27, wp70)
	ns.Move(obj28, wp71)
	ns.Move(obj29, wp72)
	ns.Move(obj30, wp74)
	ns.Move(obj31, wp75)
	ns.Move(obj32, wp76)
	ns.Move(obj33, wp77)
	ns.Move(obj34, wp78)
	ns.Move(obj35, wp79)
	ns.FrameTimer(10, stoneGrinder)
LABEL1:
	return
}
func stoneGrinder() {
	ns.MoveWaypoint(wp55, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.BoulderMove, wp55)
	ns.AudioEvent(ns.SecretWallStoneOpen, wp69)
	ns.AudioEvent(ns.BoulderMove, wp70)
	ns.AudioEvent(ns.SecretWallStoneOpen, wp71)
	ns.AudioEvent(ns.BoulderMove, wp72)
	ns.AudioEvent(ns.SecretWallStoneOpen, wp73)
	ns.AudioEvent(ns.BoulderMove, wp74)
	ns.AudioEvent(ns.SecretWallStoneOpen, wp75)
	ns.AudioEvent(ns.BoulderMove, wp76)
	ns.AudioEvent(ns.SecretWallStoneOpen, wp77)
	ns.AudioEvent(ns.BoulderMove, wp78)
	ns.AudioEvent(ns.SecretWallStoneOpen, wp79)
}
func openingGauntletPiece() {
	PlayGauntletMusic()
	ns.WideScreen(true)
	ns.Frozen(ns.GetHost(), true)
	ns.LookAtObject(ns.GetHost(), obj42)
	ns.CreatureGuard(obj42, 4982, 5036, ns.GetObjectX(obj44), ns.GetObjectY(obj44), 0)
	ns.SetDialog(obj42, ns.NEXT, guardTalk2Start, guardTalk2End)
	ns.StartDialog(obj42, ns.GetHost())
}
func realHorrendousTalk1Start() {
	ns.CreatureGuard(obj44, 4991, 4991, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	ns.TellStory(ns.DemonRecognize, "War02A.scr:RealTalk1Start")
	ns.JournalEntry(ns.GetHost(), "War2Gauntlet", 2)
}
func realHorrendousTalk1End() {
	ns.SetDialog(obj44, "FALSE", realHorrendousTalk2Start, realHorrendousTalk2End)
	ns.CreatureGuard(obj44, 5071, 4979, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	ns.CreatureGuard(obj42, 4980, 5059, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	ns.CreatureGuard(obj7, 4922, 5014, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	ns.CreatureGuard(obj8, 5014, 4922, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	ns.WideScreen(false)
	ns.Frozen(ns.GetHost(), false)
}
func realHorrendousTalk2Start() {
	ns.TellStory(ns.DemonRecognize, "War02A.scr:RealTalk2Start")
}
func realHorrendousTalk2End() {
	ns.SetDialog(obj44, "FALSE", realHorrendousTalk3Start, realHorrendousTalk3End)
}
func realHorrendousTalk3Start() {
	ns.TellStory(ns.DemonRecognize, "War02A.scr:RealTalk3Start")
}
func realHorrendousTalk3End() {
	ns.CancelDialog(obj44)
}
func guardTalk2Start() {
	ns.TellStory(ns.SwordsmanRecognize, "War02A.scr:GuardSay1Start")
}
func guardTalk2End() {
	ns.CancelDialog(obj42)
	ns.SetDialog(obj44, "FALSE", realHorrendousTalk1Start, realHorrendousTalk1End)
	ns.StartDialog(obj44, ns.GetHost())
}
func horrendousTalk1Start() {
	ns.TellStory(ns.DemonRecognize, "War02A.scr:HorrendousTalk1Start")
}
func horrendousTalk1End() {
	ns.CancelDialog(obj43)
	ns.LookAtObject(obj43, obj45)
	ns.Frozen(ns.GetHost(), false)
	ns.WideScreen(false)
	ns.FrameTimer(5, chickenShirtTalk1Start)
}
func chickenShirtTalk1Start() {
	ns.Chat(obj47, "War02A.scr:ChickenTalk1Start")
	ns.FrameTimer(120, redShirtRunGauntlet)
}
func horrendousSay1() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	if !flag11 {
		goto LABEL1
	}
	ns.Music(0, 100)
	ns.ObjectGroupOn(gvar15)
	flag11 = false
	ns.Frozen(ns.GetHost(), true)
	ns.WideScreen(true)
	ns.LockDoor(obj50)
	ns.LockDoor(obj51)
	ns.SetDialog(obj43, "FALSE", horrendousTalk1Start, horrendousTalk1End)
	ns.StartDialog(obj43, ns.GetHost())
LABEL1:
	return
}
func redShirtRunGauntlet() {
	ns.AudioEvent(ns.SwordsmanHurt, wp89)
	ns.Chat(obj45, "War02A.scr:RedShirtRunGauntlet")
	ns.FrameTimer(70, redShirtDies)
}
func redShirtDies() {
	ns.AudioEvent(ns.BerserkerChargeInvoke, wp88)
	ns.Wander(obj45)
}
func bonesFly() {
	ns.AudioEvent(ns.HammerMissing, wp89)
	ns.AudioEvent(ns.GolemHitting, wp89)
	ns.Effect(ns.JIGGLE, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 10, 0)
	ns.AudioEvent(ns.EmberDemonDie, wp88)
	ns.AudioEvent(ns.SkeletonDie, wp89)
	ivar4 = 0
	for {
		if !(ivar4 < 40) {
			goto LABEL1
		}
		obj54[ivar4] = ns.CreateObject("ArmBone", wp88)
		ns.PushObject(obj54[ivar4], -150, 1880, 3553)
		ivar4 += 1
	}
LABEL1:
	ivar4 = 0
	for {
		if !(ivar4 < 40) {
			goto LABEL3
		}
		obj54[ivar4] = ns.CreateObject("LegBone", wp88)
		ns.PushObject(obj54[ivar4], -150, 1880, 3553)
		ivar4 += 1
	}
LABEL3:
	ns.FrameTimer(30, chickenShirtRuns)
	ns.FrameTimer(10, deadHead)
}
func deadHead() {
	ns.Delete(obj45)
	ns.AudioEvent("MechGolemHit", wp86)
	ns.AudioEvent("SpikesDown", wp87)
	ns.AudioEvent(ns.SkeletonDie, wp89)
	ns.MoveObject(obj46, 1692, 3627)
	ns.PushObject(obj46, -150, 1880, 3553)
	ns.SetDialog(obj46, ns.NORMAL, deadTalk1Start, deadTalk1End)
	ns.FrameTimer(40, chickenShirtRuns)
}
func deadTalk1Start() {
	ns.TellStory("EmberDemonDies", "War02a:RedShirtDead")
}
func deadTalk1End() {
	ns.CancelDialog(obj46)
}
func chickenShirtRuns() {
	ns.Chat(obj47, "War02A.scr:ChickenShirtRuns")
	chickenShirtCries1()
	ns.CreatureGuard(obj47, 2033, 3833, 2035, 3905, 0)
	ns.CancelDialog(obj47)
	ns.SetDialog(obj47, ns.NORMAL, chickenTalk1Start, chickenTalk1End)
	ns.FrameTimer(45, horrendousSay2)
}
func chickenShirtCries1() {
	ns.AudioEvent(ns.UrchinRecognize, wp85)
}
func chickenTalk1Start() {
	ns.TellStory(ns.UrchinRecognize, "War02A.scr:ChickenShirtRuns")
}
func chickenTalk1End() {
	ns.SetDialog(obj47, ns.NORMAL, chickenTalk1Start, chickenTalk1End)
}
func horrendousSay2() {
	ns.Music(29, 100)
	ns.ObjectGroupOff(gvar15)
	ns.Chat(obj43, "War02A.scr:HorrendousSay2")
	ns.UnlockDoor(obj52)
}
func blockBash() {
	ns.MoveWaypoint(wp55, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.HammerMissing, wp55)
	ns.Effect(ns.JIGGLE, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 10, 0)
}
func blockSmash1() {
	ns.ObjectOn(obj36)
	ns.ObjectOn(obj37)
}
func blockSmash2() {
	ns.ObjectOn(obj38)
	ns.ObjectOn(obj39)
}
func blockSmash3() {
	ns.ObjectOn(obj40)
	ns.ObjectOn(obj41)
}
func releaseUrchins1() {
	ns.WallGroupOpen(gvar98)
	ns.LockDoor(obj93)
	ns.LockDoor(obj94)
	ns.ObjectOn(obj91)
	ns.ObjectOn(obj92)
	ns.ObjectOff(obj97)
}
func twoDeadUrchins() {
	if !(ivar90 > 1) {
		goto LABEL1
	}
	ns.UnlockDoor(obj95)
	ns.UnlockDoor(obj96)
LABEL1:
	return
}
func flamesOn() {
	if !flag139 {
		goto LABEL1
	}
	ivar4 = 0
	for {
		if !(ivar4 < 8) {
			goto LABEL2
		}
		obj186[ivar4] = 0
		obj187[ivar4] = 0
		obj188[ivar4] = 0
		obj189[ivar4] = 0
		obj190[ivar4] = 0
		obj191[ivar4] = 0
		obj192[ivar4] = 0
		obj193[ivar4] = 0
		obj194[ivar4] = 0
		obj195[ivar4] = 0
		ivar4 += 1
	}
LABEL2:
	flag139 = false
	flag140 = true
	flag141 = true
	flag142 = true
	flag143 = true
	flag144 = true
	flag145 = true
	flag146 = true
	flag147 = true
	flag148 = true
	flag149 = true
	flameWalk()
	ns.FrameTimer(4, flameWalk2)
	ns.FrameTimer(8, flameWalk3)
	ns.FrameTimer(12, flameWalk4)
	ns.FrameTimer(16, flameWalk5)
	ns.FrameTimer(20, flameWalk6)
	ns.FrameTimer(24, flameWalk7)
	ns.FrameTimer(28, flameWalk8)
	ns.FrameTimer(32, flameWalk9)
	ns.FrameTimer(36, flameWalk10)
	ns.LockDoor(obj52)
	ns.LockDoor(obj48)
	ns.LockDoor(obj49)
LABEL1:
	return
}
func flamesOff() {
	flag140 = false
	flag141 = false
	flag142 = false
	flag143 = false
	flag144 = false
	flag145 = false
	flag146 = false
	flag147 = false
	flag148 = false
	flag149 = false
}
func flameWalk() {
	var v0 int
	v0 = gvar166
	if v0 == gvar99 {
		goto LABEL1
	}
	if v0 == gvar100 {
		goto LABEL2
	}
	if v0 == gvar101 {
		goto LABEL3
	}
	if v0 == gvar102 {
		goto LABEL4
	}
	goto LABEL5
LABEL1:
	ivar4 = 0
	for {
		if !(ivar4 < 8) {
			goto LABEL6
		}
		ns.Delete(obj186[ivar4])
		obj186[ivar4] = 0
		ivar4 += 1
	}
LABEL6:
	gvar166 = gvar100
	gvar176 = gvar99
	ivar156 = 60
	goto LABEL5
LABEL2:
	ivar4 = 0
	for {
		if !(ivar4 < 8) {
			goto LABEL8
		}
		ns.Delete(obj186[ivar4])
		obj186[ivar4] = ns.CreateObject("SmallFlame", wp196[ivar4])
		ivar4 += 1
	}
LABEL8:
	if gvar176 != gvar99 {
		goto LABEL10
	}
	gvar166 = gvar101
	ns.AudioEvent(ns.FireballCast, wp196[4])
	goto LABEL11
LABEL10:
	gvar166 = gvar99
LABEL11:
	gvar176 = gvar100
	ivar156 = 2
	goto LABEL5
LABEL3:
	ivar4 = 0
	for {
		if !(ivar4 < 8) {
			goto LABEL12
		}
		ns.Delete(obj186[ivar4])
		obj186[ivar4] = ns.CreateObject("MediumFlame", wp196[ivar4])
		ivar4 += 1
	}
LABEL12:
	if gvar176 != gvar100 {
		goto LABEL14
	}
	gvar166 = gvar102
	ns.AudioEvent(ns.DemonBreath, wp196[5])
	goto LABEL15
LABEL14:
	gvar166 = gvar100
	ns.AudioEvent(ns.FireExtinguish, wp196[4])
LABEL15:
	gvar176 = gvar101
	ivar156 = 2
	goto LABEL5
LABEL4:
	ivar4 = 0
	for {
		if !(ivar4 < 8) {
			goto LABEL16
		}
		ns.Delete(obj186[ivar4])
		obj186[ivar4] = ns.CreateObject("Flame", wp196[ivar4])
		ivar4 += 1
	}
LABEL16:
	gvar166 = gvar101
	gvar176 = gvar102
	ivar156 = 12
	goto LABEL5
LABEL5:
	if !flag140 {
		goto LABEL18
	}
	ns.FrameTimer(ivar156, flameWalk)
	goto LABEL19
LABEL18:
	ivar154 = 0
	for {
		if !(ivar154 < 8) {
			goto LABEL19
		}
		ns.Delete(obj186[ivar154])
		obj186[ivar154] = 0
		ivar154 += 1
	}
LABEL19:
	return
}
func flameWalk2() {
	var v0 int
	v0 = gvar167
	if v0 == gvar103 {
		goto LABEL1
	}
	if v0 == gvar104 {
		goto LABEL2
	}
	if v0 == gvar105 {
		goto LABEL3
	}
	if v0 == gvar106 {
		goto LABEL4
	}
	goto LABEL5
LABEL1:
	ivar150 = 0
	for {
		if !(ivar150 < 8) {
			goto LABEL6
		}
		ns.Delete(obj187[ivar150])
		obj187[ivar150] = 0
		ivar150 += 1
	}
LABEL6:
	gvar167 = gvar104
	gvar177 = gvar103
	ivar157 = 60
	goto LABEL5
LABEL2:
	ivar150 = 0
	for {
		if !(ivar150 < 8) {
			goto LABEL8
		}
		ns.Delete(obj187[ivar150])
		obj187[ivar150] = ns.CreateObject("SmallFlame", wp197[ivar150])
		ivar150 += 1
	}
LABEL8:
	if gvar177 != gvar103 {
		goto LABEL10
	}
	gvar167 = gvar105
	ns.AudioEvent(ns.FireballCast, wp197[4])
	goto LABEL11
LABEL10:
	gvar167 = gvar103
LABEL11:
	gvar177 = gvar104
	ivar157 = 2
	goto LABEL5
LABEL3:
	ivar150 = 0
	for {
		if !(ivar150 < 8) {
			goto LABEL12
		}
		ns.Delete(obj187[ivar150])
		obj187[ivar150] = ns.CreateObject("MediumFlame", wp197[ivar150])
		ivar150 += 1
	}
LABEL12:
	if gvar177 != gvar104 {
		goto LABEL14
	}
	gvar167 = gvar106
	ns.AudioEvent(ns.DemonBreath, wp197[5])
	goto LABEL15
LABEL14:
	gvar167 = gvar104
	ns.AudioEvent(ns.FireExtinguish, wp197[4])
LABEL15:
	gvar177 = gvar105
	ivar157 = 2
	goto LABEL5
LABEL4:
	ivar150 = 0
	for {
		if !(ivar150 < 8) {
			goto LABEL16
		}
		ns.Delete(obj187[ivar150])
		obj187[ivar150] = ns.CreateObject("Flame", wp197[ivar150])
		ivar150 += 1
	}
LABEL16:
	gvar167 = gvar105
	gvar177 = gvar106
	ivar157 = 12
	goto LABEL5
LABEL5:
	if !flag141 {
		goto LABEL18
	}
	ns.FrameTimer(ivar157, flameWalk2)
	goto LABEL19
LABEL18:
	ivar154 = 0
	for {
		if !(ivar154 < 8) {
			goto LABEL19
		}
		ns.Delete(obj187[ivar154])
		obj187[ivar154] = 0
		ivar154 += 1
	}
LABEL19:
	return
}
func flameWalk3() {
	var v0 int
	v0 = gvar168
	if v0 == gvar107 {
		goto LABEL1
	}
	if v0 == gvar108 {
		goto LABEL2
	}
	if v0 == gvar109 {
		goto LABEL3
	}
	if v0 == gvar110 {
		goto LABEL4
	}
	goto LABEL5
LABEL1:
	ivar151 = 0
	for {
		if !(ivar151 < 8) {
			goto LABEL6
		}
		ns.Delete(obj188[ivar151])
		obj188[ivar151] = 0
		ivar151 += 1
	}
LABEL6:
	gvar168 = gvar108
	gvar178 = gvar107
	ivar158 = 60
	goto LABEL5
LABEL2:
	ivar151 = 0
	for {
		if !(ivar151 < 8) {
			goto LABEL8
		}
		ns.Delete(obj188[ivar151])
		obj188[ivar151] = ns.CreateObject("SmallFlame", wp198[ivar151])
		ivar151 += 1
	}
LABEL8:
	if gvar178 != gvar107 {
		goto LABEL10
	}
	gvar168 = gvar109
	ns.AudioEvent(ns.FireballCast, wp198[4])
	goto LABEL11
LABEL10:
	gvar168 = gvar107
LABEL11:
	gvar178 = gvar108
	ivar158 = 2
	goto LABEL5
LABEL3:
	ivar151 = 0
	for {
		if !(ivar151 < 8) {
			goto LABEL12
		}
		ns.Delete(obj188[ivar151])
		obj188[ivar151] = ns.CreateObject("MediumFlame", wp198[ivar151])
		ivar151 += 1
	}
LABEL12:
	if gvar178 != gvar108 {
		goto LABEL14
	}
	gvar168 = gvar110
	ns.AudioEvent(ns.DemonBreath, wp198[5])
	goto LABEL15
LABEL14:
	gvar168 = gvar108
	ns.AudioEvent(ns.FireExtinguish, wp198[4])
LABEL15:
	gvar178 = gvar109
	ivar158 = 2
	goto LABEL5
LABEL4:
	ivar151 = 0
	for {
		if !(ivar151 < 8) {
			goto LABEL16
		}
		ns.Delete(obj188[ivar151])
		obj188[ivar151] = ns.CreateObject("Flame", wp198[ivar151])
		ivar151 += 1
	}
LABEL16:
	gvar168 = gvar109
	gvar178 = gvar110
	ivar158 = 12
	goto LABEL5
LABEL5:
	if !flag142 {
		goto LABEL18
	}
	ns.FrameTimer(ivar158, flameWalk3)
	goto LABEL19
LABEL18:
	ivar154 = 0
	for {
		if !(ivar154 < 8) {
			goto LABEL19
		}
		ns.Delete(obj188[ivar154])
		obj188[ivar154] = 0
		ivar154 += 1
	}
LABEL19:
	return
}
func flameWalk4() {
	var v0 int
	v0 = gvar169
	if v0 == gvar111 {
		goto LABEL1
	}
	if v0 == gvar112 {
		goto LABEL2
	}
	if v0 == gvar113 {
		goto LABEL3
	}
	if v0 == gvar114 {
		goto LABEL4
	}
	goto LABEL5
LABEL1:
	ivar152 = 0
	for {
		if !(ivar152 < 8) {
			goto LABEL6
		}
		ns.Delete(obj189[ivar152])
		obj189[ivar152] = 0
		ivar152 += 1
	}
LABEL6:
	gvar169 = gvar112
	gvar179 = gvar111
	ivar159 = 60
	goto LABEL5
LABEL2:
	ivar152 = 0
	for {
		if !(ivar152 < 8) {
			goto LABEL8
		}
		ns.Delete(obj189[ivar152])
		obj189[ivar152] = ns.CreateObject("SmallFlame", wp199[ivar152])
		ivar152 += 1
	}
LABEL8:
	if gvar179 != gvar111 {
		goto LABEL10
	}
	gvar169 = gvar113
	ns.AudioEvent(ns.FireballCast, wp199[4])
	goto LABEL11
LABEL10:
	gvar169 = gvar111
LABEL11:
	gvar179 = gvar112
	ivar159 = 2
	goto LABEL5
LABEL3:
	ivar152 = 0
	for {
		if !(ivar152 < 8) {
			goto LABEL12
		}
		ns.Delete(obj189[ivar152])
		obj189[ivar152] = ns.CreateObject("MediumFlame", wp199[ivar152])
		ivar152 += 1
	}
LABEL12:
	if gvar179 != gvar112 {
		goto LABEL14
	}
	gvar169 = gvar114
	ns.AudioEvent(ns.DemonBreath, wp199[5])
	goto LABEL15
LABEL14:
	gvar169 = gvar112
	ns.AudioEvent(ns.FireExtinguish, wp199[4])
LABEL15:
	gvar179 = gvar113
	ivar159 = 2
	goto LABEL5
LABEL4:
	ivar152 = 0
	for {
		if !(ivar152 < 8) {
			goto LABEL16
		}
		ns.Delete(obj189[ivar152])
		obj189[ivar152] = ns.CreateObject("Flame", wp199[ivar152])
		ivar152 += 1
	}
LABEL16:
	gvar169 = gvar113
	gvar179 = gvar114
	ivar159 = 12
	goto LABEL5
LABEL5:
	if !flag143 {
		goto LABEL18
	}
	ns.FrameTimer(ivar159, flameWalk4)
	goto LABEL19
LABEL18:
	ivar154 = 0
	for {
		if !(ivar154 < 8) {
			goto LABEL19
		}
		ns.Delete(obj189[ivar154])
		obj189[ivar154] = 0
		ivar154 += 1
	}
LABEL19:
	return
}
func flameWalk5() {
	var v0 int
	v0 = gvar170
	if v0 == gvar115 {
		goto LABEL1
	}
	if v0 == gvar116 {
		goto LABEL2
	}
	if v0 == gvar117 {
		goto LABEL3
	}
	if v0 == gvar118 {
		goto LABEL4
	}
	goto LABEL5
LABEL1:
	ivar153 = 0
	for {
		if !(ivar153 < 8) {
			goto LABEL6
		}
		ns.Delete(obj190[ivar153])
		obj190[ivar153] = 0
		ivar153 += 1
	}
LABEL6:
	gvar170 = gvar116
	gvar180 = gvar115
	ivar160 = 60
	goto LABEL5
LABEL2:
	ivar153 = 0
	for {
		if !(ivar153 < 8) {
			goto LABEL8
		}
		ns.Delete(obj190[ivar153])
		obj190[ivar153] = ns.CreateObject("SmallFlame", wp200[ivar153])
		ivar153 += 1
	}
LABEL8:
	if gvar180 != gvar115 {
		goto LABEL10
	}
	gvar170 = gvar117
	ns.AudioEvent(ns.FireballCast, wp200[4])
	goto LABEL11
LABEL10:
	gvar170 = gvar115
LABEL11:
	gvar180 = gvar116
	ivar160 = 2
	goto LABEL5
LABEL3:
	ivar153 = 0
	for {
		if !(ivar153 < 8) {
			goto LABEL12
		}
		ns.Delete(obj190[ivar153])
		obj190[ivar153] = ns.CreateObject("MediumFlame", wp200[ivar153])
		ivar153 += 1
	}
LABEL12:
	if gvar180 != gvar116 {
		goto LABEL14
	}
	gvar170 = gvar118
	ns.AudioEvent(ns.DemonBreath, wp200[5])
	goto LABEL15
LABEL14:
	gvar170 = gvar116
	ns.AudioEvent(ns.FireExtinguish, wp200[4])
LABEL15:
	gvar180 = gvar117
	ivar160 = 2
	goto LABEL5
LABEL4:
	ivar153 = 0
	for {
		if !(ivar153 < 8) {
			goto LABEL16
		}
		ns.Delete(obj190[ivar153])
		obj190[ivar153] = ns.CreateObject("Flame", wp200[ivar153])
		ivar153 += 1
	}
LABEL16:
	gvar170 = gvar117
	gvar180 = gvar118
	ivar160 = 12
	goto LABEL5
LABEL5:
	if !flag144 {
		goto LABEL18
	}
	ns.FrameTimer(ivar160, flameWalk5)
	goto LABEL19
LABEL18:
	ivar154 = 0
	for {
		if !(ivar154 < 8) {
			goto LABEL19
		}
		ns.Delete(obj190[ivar154])
		obj190[ivar154] = 0
		ivar154 += 1
	}
LABEL19:
	return
}
func flameWalk6() {
	var v0 int
	v0 = gvar171
	if v0 == gvar119 {
		goto LABEL1
	}
	if v0 == gvar120 {
		goto LABEL2
	}
	if v0 == gvar121 {
		goto LABEL3
	}
	if v0 == gvar122 {
		goto LABEL4
	}
	goto LABEL5
LABEL1:
	ivar154 = 0
	for {
		if !(ivar154 < 8) {
			goto LABEL6
		}
		ns.Delete(obj191[ivar154])
		obj191[ivar154] = 0
		ivar154 += 1
	}
LABEL6:
	gvar171 = gvar120
	gvar181 = gvar119
	ivar161 = 60
	goto LABEL5
LABEL2:
	ivar154 = 0
	for {
		if !(ivar154 < 8) {
			goto LABEL8
		}
		ns.Delete(obj191[ivar154])
		obj191[ivar154] = ns.CreateObject("SmallFlame", wp201[ivar154])
		ivar154 += 1
	}
LABEL8:
	if gvar181 != gvar119 {
		goto LABEL10
	}
	gvar171 = gvar121
	ns.AudioEvent(ns.FireballCast, wp201[4])
	goto LABEL11
LABEL10:
	gvar171 = gvar119
LABEL11:
	gvar181 = gvar120
	ivar161 = 2
	goto LABEL5
LABEL3:
	ivar154 = 0
	for {
		if !(ivar154 < 8) {
			goto LABEL12
		}
		ns.Delete(obj191[ivar154])
		obj191[ivar154] = ns.CreateObject("MediumFlame", wp201[ivar154])
		ivar154 += 1
	}
LABEL12:
	if gvar181 != gvar120 {
		goto LABEL14
	}
	gvar171 = gvar122
	ns.AudioEvent(ns.DemonBreath, wp201[5])
	goto LABEL15
LABEL14:
	gvar171 = gvar120
	ns.AudioEvent(ns.FireExtinguish, wp201[4])
LABEL15:
	gvar181 = gvar121
	ivar161 = 2
	goto LABEL5
LABEL4:
	ivar154 = 0
	for {
		if !(ivar154 < 8) {
			goto LABEL16
		}
		ns.Delete(obj191[ivar154])
		obj191[ivar154] = ns.CreateObject("Flame", wp201[ivar154])
		ivar154 += 1
	}
LABEL16:
	gvar171 = gvar121
	gvar181 = gvar122
	ivar161 = 12
	goto LABEL5
LABEL5:
	if !flag145 {
		goto LABEL18
	}
	ns.FrameTimer(ivar161, flameWalk6)
	goto LABEL19
LABEL18:
	ivar154 = 0
	for {
		if !(ivar154 < 8) {
			goto LABEL19
		}
		ns.Delete(obj191[ivar154])
		obj191[ivar154] = 0
		ivar154 += 1
	}
LABEL19:
	return
}
func flameWalk7() {
	var v0 int
	v0 = gvar172
	if v0 == gvar123 {
		goto LABEL1
	}
	if v0 == gvar124 {
		goto LABEL2
	}
	if v0 == gvar125 {
		goto LABEL3
	}
	if v0 == gvar126 {
		goto LABEL4
	}
	goto LABEL5
LABEL1:
	ivar154 = 0
	for {
		if !(ivar154 < 8) {
			goto LABEL6
		}
		ns.Delete(obj192[ivar154])
		obj192[ivar154] = 0
		ivar154 += 1
	}
LABEL6:
	gvar172 = gvar124
	gvar182 = gvar123
	ivar162 = 60
	goto LABEL5
LABEL2:
	ivar154 = 0
	for {
		if !(ivar154 < 8) {
			goto LABEL8
		}
		ns.Delete(obj192[ivar154])
		obj192[ivar154] = ns.CreateObject("SmallFlame", wp202[ivar154])
		ivar154 += 1
	}
LABEL8:
	if gvar182 != gvar123 {
		goto LABEL10
	}
	gvar172 = gvar125
	ns.AudioEvent(ns.FireballCast, wp202[4])
	goto LABEL11
LABEL10:
	gvar172 = gvar123
LABEL11:
	gvar182 = gvar124
	ivar162 = 2
	goto LABEL5
LABEL3:
	ivar154 = 0
	for {
		if !(ivar154 < 8) {
			goto LABEL12
		}
		ns.Delete(obj192[ivar154])
		obj192[ivar154] = ns.CreateObject("MediumFlame", wp202[ivar154])
		ivar154 += 1
	}
LABEL12:
	if gvar182 != gvar124 {
		goto LABEL14
	}
	gvar172 = gvar126
	ns.AudioEvent(ns.DemonBreath, wp202[5])
	goto LABEL15
LABEL14:
	gvar172 = gvar124
	ns.AudioEvent(ns.FireExtinguish, wp202[4])
LABEL15:
	gvar182 = gvar125
	ivar162 = 2
	goto LABEL5
LABEL4:
	ivar154 = 0
	for {
		if !(ivar154 < 8) {
			goto LABEL16
		}
		ns.Delete(obj192[ivar154])
		obj192[ivar154] = ns.CreateObject("Flame", wp202[ivar154])
		ivar154 += 1
	}
LABEL16:
	gvar172 = gvar125
	gvar182 = gvar126
	ivar162 = 12
	goto LABEL5
LABEL5:
	if !flag146 {
		goto LABEL18
	}
	ns.FrameTimer(ivar162, flameWalk7)
	goto LABEL19
LABEL18:
	ivar154 = 0
	for {
		if !(ivar154 < 8) {
			goto LABEL19
		}
		ns.Delete(obj192[ivar154])
		obj192[ivar154] = 0
		ivar154 += 1
	}
LABEL19:
	return
}
func flameWalk8() {
	var v0 int
	v0 = gvar173
	if v0 == gvar127 {
		goto LABEL1
	}
	if v0 == gvar128 {
		goto LABEL2
	}
	if v0 == gvar129 {
		goto LABEL3
	}
	if v0 == gvar130 {
		goto LABEL4
	}
	goto LABEL5
LABEL1:
	ivar155 = 0
	for {
		if !(ivar155 < 8) {
			goto LABEL6
		}
		ns.Delete(obj193[ivar155])
		obj193[ivar155] = 0
		ivar155 += 1
	}
LABEL6:
	gvar173 = gvar128
	gvar183 = gvar127
	ivar163 = 60
	goto LABEL5
LABEL2:
	ivar155 = 0
	for {
		if !(ivar155 < 8) {
			goto LABEL8
		}
		ns.Delete(obj193[ivar155])
		obj193[ivar155] = ns.CreateObject("SmallFlame", wp203[ivar155])
		ivar155 += 1
	}
LABEL8:
	if gvar183 != gvar127 {
		goto LABEL10
	}
	gvar173 = gvar129
	ns.AudioEvent(ns.FireballCast, wp203[4])
	goto LABEL11
LABEL10:
	gvar173 = gvar127
LABEL11:
	gvar183 = gvar128
	ivar163 = 2
	goto LABEL5
LABEL3:
	ivar155 = 0
	for {
		if !(ivar155 < 8) {
			goto LABEL12
		}
		ns.Delete(obj193[ivar155])
		obj193[ivar155] = ns.CreateObject("MediumFlame", wp203[ivar155])
		ivar155 += 1
	}
LABEL12:
	if gvar183 != gvar128 {
		goto LABEL14
	}
	gvar173 = gvar130
	ns.AudioEvent(ns.DemonBreath, wp203[5])
	goto LABEL15
LABEL14:
	gvar173 = gvar128
	ns.AudioEvent(ns.FireExtinguish, wp203[4])
LABEL15:
	gvar183 = gvar129
	ivar163 = 2
	goto LABEL5
LABEL4:
	ivar155 = 0
	for {
		if !(ivar155 < 8) {
			goto LABEL16
		}
		ns.Delete(obj193[ivar155])
		obj193[ivar155] = ns.CreateObject("Flame", wp203[ivar155])
		ivar155 += 1
	}
LABEL16:
	gvar173 = gvar129
	gvar183 = gvar130
	ivar163 = 12
	goto LABEL5
LABEL5:
	if !flag147 {
		goto LABEL18
	}
	ns.FrameTimer(ivar163, flameWalk8)
	goto LABEL19
LABEL18:
	ivar154 = 0
	for {
		if !(ivar154 < 8) {
			goto LABEL19
		}
		ns.Delete(obj193[ivar154])
		obj193[ivar154] = 0
		ivar154 += 1
	}
LABEL19:
	return
}
func flameWalk9() {
	var v0 int
	v0 = gvar174
	if v0 == gvar131 {
		goto LABEL1
	}
	if v0 == gvar132 {
		goto LABEL2
	}
	if v0 == gvar133 {
		goto LABEL3
	}
	if v0 == gvar134 {
		goto LABEL4
	}
	goto LABEL5
LABEL1:
	ivar154 = 0
	for {
		if !(ivar154 < 8) {
			goto LABEL6
		}
		ns.Delete(obj194[ivar154])
		obj194[ivar154] = 0
		ivar154 += 1
	}
LABEL6:
	gvar174 = gvar132
	gvar184 = gvar131
	ivar164 = 60
	goto LABEL5
LABEL2:
	ivar154 = 0
	for {
		if !(ivar154 < 8) {
			goto LABEL8
		}
		ns.Delete(obj194[ivar154])
		obj194[ivar154] = ns.CreateObject("SmallFlame", wp204[ivar154])
		ivar154 += 1
	}
LABEL8:
	if gvar184 != gvar131 {
		goto LABEL10
	}
	gvar174 = gvar133
	ns.AudioEvent(ns.FireballCast, wp203[4])
	goto LABEL11
LABEL10:
	gvar174 = gvar131
LABEL11:
	gvar184 = gvar132
	ivar164 = 2
	goto LABEL5
LABEL3:
	ivar154 = 0
	for {
		if !(ivar154 < 8) {
			goto LABEL12
		}
		ns.Delete(obj194[ivar154])
		obj194[ivar154] = ns.CreateObject("MediumFlame", wp204[ivar154])
		ivar154 += 1
	}
LABEL12:
	if gvar184 != gvar132 {
		goto LABEL14
	}
	gvar174 = gvar134
	ns.AudioEvent(ns.DemonBreath, wp204[5])
	goto LABEL15
LABEL14:
	gvar174 = gvar132
	ns.AudioEvent(ns.FireExtinguish, wp204[4])
LABEL15:
	gvar184 = gvar133
	ivar164 = 2
	goto LABEL5
LABEL4:
	ivar154 = 0
	for {
		if !(ivar154 < 8) {
			goto LABEL16
		}
		ns.Delete(obj194[ivar154])
		obj194[ivar154] = ns.CreateObject("Flame", wp204[ivar154])
		ivar154 += 1
	}
LABEL16:
	gvar174 = gvar133
	gvar184 = gvar134
	ivar164 = 12
	goto LABEL5
LABEL5:
	if !flag148 {
		goto LABEL18
	}
	ns.FrameTimer(ivar164, flameWalk9)
	goto LABEL19
LABEL18:
	ivar154 = 0
	for {
		if !(ivar154 < 8) {
			goto LABEL19
		}
		ns.Delete(obj194[ivar154])
		obj194[ivar154] = 0
		ivar154 += 1
	}
LABEL19:
	return
}
func flameWalk10() {
	var v0 int
	v0 = gvar175
	if v0 == gvar135 {
		goto LABEL1
	}
	if v0 == gvar136 {
		goto LABEL2
	}
	if v0 == gvar137 {
		goto LABEL3
	}
	if v0 == gvar138 {
		goto LABEL4
	}
	goto LABEL5
LABEL1:
	ivar154 = 0
	for {
		if !(ivar154 < 8) {
			goto LABEL6
		}
		ns.Delete(obj195[ivar154])
		obj195[ivar154] = 0
		ivar154 += 1
	}
LABEL6:
	gvar175 = gvar136
	gvar185 = gvar135
	ivar165 = 60
	goto LABEL5
LABEL2:
	ivar154 = 0
	for {
		if !(ivar154 < 8) {
			goto LABEL8
		}
		ns.Delete(obj195[ivar154])
		obj195[ivar154] = ns.CreateObject("SmallFlame", wp205[ivar154])
		ivar154 += 1
	}
LABEL8:
	if gvar185 != gvar135 {
		goto LABEL10
	}
	gvar175 = gvar137
	ns.AudioEvent(ns.FireballCast, wp205[4])
	goto LABEL11
LABEL10:
	gvar175 = gvar135
LABEL11:
	gvar185 = gvar136
	ivar165 = 2
	goto LABEL5
LABEL3:
	ivar154 = 0
	for {
		if !(ivar154 < 8) {
			goto LABEL12
		}
		ns.Delete(obj195[ivar154])
		obj195[ivar154] = ns.CreateObject("MediumFlame", wp205[ivar154])
		ivar154 += 1
	}
LABEL12:
	if gvar185 != gvar136 {
		goto LABEL14
	}
	gvar175 = gvar138
	ns.AudioEvent(ns.DemonBreath, wp205[5])
	goto LABEL15
LABEL14:
	gvar175 = gvar136
	ns.AudioEvent(ns.FireExtinguish, wp205[4])
LABEL15:
	gvar185 = gvar137
	ivar165 = 2
	goto LABEL5
LABEL4:
	ivar154 = 0
	for {
		if !(ivar154 < 8) {
			goto LABEL16
		}
		ns.Delete(obj195[ivar154])
		obj195[ivar154] = ns.CreateObject("Flame", wp205[ivar154])
		ivar154 += 1
	}
LABEL16:
	gvar175 = gvar137
	gvar185 = gvar138
	ivar165 = 12
	goto LABEL5
LABEL5:
	if !flag149 {
		goto LABEL18
	}
	ns.FrameTimer(ivar165, flameWalk10)
	goto LABEL19
LABEL18:
	ivar154 = 0
	for {
		if !(ivar154 < 8) {
			goto LABEL19
		}
		ns.Delete(obj195[ivar154])
		obj195[ivar154] = 0
		ivar154 += 1
	}
LABEL19:
	return
}
func initFlameGrates() {
	ivar4 = 0
	for {
		if !(ivar4 < 8) {
			goto LABEL1
		}
		wp196[ivar4] = ns.Waypoint("FlameWay1-" + strconv.Itoa(ivar4+1))
		wp197[ivar4] = ns.Waypoint("FlameWay2-" + strconv.Itoa(ivar4+1))
		wp198[ivar4] = ns.Waypoint("FlameWay3-" + strconv.Itoa(ivar4+1))
		wp199[ivar4] = ns.Waypoint("FlameWay4-" + strconv.Itoa(ivar4+1))
		wp200[ivar4] = ns.Waypoint("FlameWay5-" + strconv.Itoa(ivar4+1))
		wp201[ivar4] = ns.Waypoint("FlameWay6-" + strconv.Itoa(ivar4+1))
		wp202[ivar4] = ns.Waypoint("FlameWay7-" + strconv.Itoa(ivar4+1))
		wp203[ivar4] = ns.Waypoint("FlameWay8-" + strconv.Itoa(ivar4+1))
		wp204[ivar4] = ns.Waypoint("FlameWay9-" + strconv.Itoa(ivar4+1))
		wp205[ivar4] = ns.Waypoint("FlameWay10-" + strconv.Itoa(ivar4+1))
		ivar4 += 1
	}
LABEL1:
	return
}
func releaseUrchins2() {
	ns.WallGroupOpen(gvar214)
	ns.LockDoor(obj209)
	ns.LockDoor(obj210)
	ns.ObjectOn(obj207)
	ns.ObjectOn(obj208)
	ns.ObjectOff(obj213)
}
func urchins2Die() {
	ivar206 += 1
	threeDeadUrchins()
}
func threeDeadUrchins() {
	if !(ivar206 > 1) {
		goto LABEL1
	}
	ns.UnlockDoor(obj211)
	ns.UnlockDoor(obj212)
LABEL1:
	return
}
func skullHallOn1() {
	if !flag215 {
		goto LABEL1
	}
	ns.CastSpellLocationLocation(ns.SPELL_FIREBALL, 2323, 2047, 2794, 2518)
	ns.FrameTimer(75, skullHallOn1)
LABEL1:
	return
}
func skullHallOff1() {
	flag215 = false
}
func releaseArena3() {
	ns.Music(4, 100)
	ns.WallGroupOpen(gvar228)
	ns.ObjectOn(obj224)
	ns.ObjectOn(obj225)
	ns.ObjectOn(obj226)
	ns.LockDoor(obj220)
	ns.LockDoor(obj221)
	ns.ObjectOff(obj227)
}
func arena3Die() {
	ivar219 += 1
	arena3DeathToll()
}
func arena3DeathToll() {
	if !(ivar219 > 2) {
		goto LABEL1
	}
	ns.UnlockDoor(obj222)
	ns.UnlockDoor(obj223)
LABEL1:
	return
}
func lockArena4() {
	ns.LockDoor(obj230)
	ns.LockDoor(obj231)
}
func loadArena4() {
	ns.ObjectGroupOn(gvar233)
	lockArena4()
	ns.ObjectOff(obj232)
	ns.FrameTimer(10, arena4SpikesUp)
	ns.FrameTimer(40, disableArena4Elevators)
	ns.FrameTimer(90, openArena4Cages)
}
func disableArena4Elevators() {
	ns.ObjectGroupOff(gvar233)
}
func openArena4Cages() {
	ns.WallGroupOpen(gvar235)
}
func arena4Die() {
	ivar229 += 1
	arena4DeathToll()
	ns.FrameTimer(90, arena4SpikesDown)
}
func arena4DeathToll() {
	if !(ivar229 > 0) {
		goto LABEL1
	}
	ns.Music(29, 100)
	ns.UnlockDoor(obj231)
	ns.UnlockDoor(obj240)
LABEL1:
	return
}
func arena4SpikesUp() {
	ns.ObjectGroupOn(gvar234)
	ns.MoveWaypoint(wp55, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.FloorSpikesUp, wp55)
}
func arena4SpikesDown() {
	ns.ObjectGroupOff(gvar234)
	ns.MoveWaypoint(wp55, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.FloorSpikesDown, wp55)
}
func enableZipperOfDeath() {
	flag238 = true
	cascade1()
	zipperNoise()
	zipperNoise2()
}
func zipperNoise() {
	if !flag238 {
		goto LABEL1
	}
	ns.AudioEvent(ns.SecretWallStoneOpen, wp262)
	ns.AudioEvent(ns.BoulderRoll, wp264)
	ns.AudioEvent(ns.BoulderRoll, wp265)
	ns.AudioEvent(ns.SecretWallStoneOpen, wp266)
	ns.AudioEvent(ns.EarthRumbleMajor, wp268)
	ns.AudioEvent(ns.EarthRumbleMajor, wp270)
	ns.AudioEvent(ns.EarthRumbleMajor, wp272)
	ns.AudioEvent(ns.EarthRumbleMajor, wp274)
	ns.AudioEvent(ns.SecretWallStoneOpen, wp276)
	ns.AudioEvent(ns.BoulderRoll, wp278)
	ns.AudioEvent(ns.BoulderRoll, wp279)
	ns.AudioEvent(ns.SecretWallStoneOpen, wp280)
	ns.AudioEvent(ns.SecretWallStoneOpen, wp281)
	ns.FrameTimer(30, zipperNoise)
LABEL1:
	return
}
func zipperNoise2() {
	if !flag238 {
		goto LABEL1
	}
	ns.AudioEvent(ns.MechGolemPowerUp, wp270)
	ns.AudioEvent(ns.MechGolemPowerUp, wp271)
	ns.FrameTimer(5, zipperNoise2)
LABEL1:
	return
}
func cascade1() {
	ns.ObjectOn(obj261)
	ns.ObjectOn(obj241)
	ns.ObjectOn(obj242)
	ns.FrameTimer(4, cascade2)
}
func cascade2() {
	ns.ObjectOn(obj243)
	ns.ObjectOn(obj244)
	ns.FrameTimer(4, cascade3)
}
func cascade3() {
	ns.ObjectOn(obj245)
	ns.ObjectOn(obj246)
	ns.FrameTimer(4, cascade4)
}
func cascade4() {
	ns.ObjectOn(obj247)
	ns.ObjectOn(obj248)
	ns.FrameTimer(4, cascade5)
}
func cascade5() {
	ns.ObjectOn(obj249)
	ns.ObjectOn(obj250)
	ns.FrameTimer(4, cascade6)
}
func cascade6() {
	ns.ObjectOn(obj251)
	ns.ObjectOn(obj252)
	ns.FrameTimer(4, cascade7)
}
func cascade7() {
	ns.ObjectOn(obj253)
	ns.ObjectOn(obj254)
	ns.FrameTimer(4, cascade8)
}
func cascade8() {
	ns.ObjectOn(obj255)
	ns.ObjectOn(obj256)
	ns.FrameTimer(4, cascade9)
}
func cascade9() {
	ns.ObjectOn(obj257)
	ns.ObjectOn(obj258)
	ns.FrameTimer(4, cascade10)
}
func cascade10() {
	ns.ObjectOn(obj259)
	ns.ObjectOn(obj260)
}
func endGauntletTalk1Start() {
	ns.TellStory(ns.SwordsmanRecognize, "War02A.scr:HorrendousTalk3Start")
}
func endGauntletTalk1End() {
	ns.CancelDialog(obj43)
	ns.SetQuestStatus(1, "EnteredGauntlet")
	ns.MoveWaypoint(wp55, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.FlagCapture, wp55)
	ns.JournalEdit(ns.GetHost(), "War2Gauntlet", 4)
	ns.Blind()
	ns.FrameTimer(60, backToDunMir)
}
func backToDunMir() {
	ns.WideScreen(false)
	ns.Frozen(ns.GetHost(), false)
	ns.MoveObject(ns.GetHost(), 5323, 5301)
}
func MapInitialize() {
	ns.NoWallSound(false)
	ns.Music(0, 100)
	obj7 = ns.Object("GG1")
	obj8 = ns.Object("GG2")
	gvar9 = ns.ObjectGroup("FireKnights")
	wp10 = ns.Waypoint("WarCryWay")
	ns.FrameTimer(3, initGauntletDoor)
	obj42 = ns.Object("F2Dant")
	ns.StoryPic(obj42, "WarriorPic")
	obj43 = ns.Object("F2Sargent")
	ns.StoryPic(obj43, "Warrior3Pic")
	obj44 = ns.Object("Horrendous")
	ns.StoryPic(obj44, "HorrendousPic")
	obj45 = ns.Object("RedShirt")
	obj46 = ns.Object("F2RedShirt")
	ns.StoryPic(obj46, "SkullPic")
	ns.SetOwner(ns.GetHost(), obj46)
	obj47 = ns.Object("F2ChickenShirt")
	ns.StoryPic(obj47, "Townsman2Pic")
	obj53 = ns.Object("GearRoom")
	obj48 = ns.Object("FromAcademy1")
	obj49 = ns.Object("FromAcademy2")
	obj50 = ns.Object("ToGauntlet1")
	obj51 = ns.Object("ToGauntlet2")
	obj52 = ns.Object("GauntletEntrance")
	obj36 = ns.Object("GSMover1")
	obj37 = ns.Object("GSMover2")
	obj38 = ns.Object("GSMover3")
	obj39 = ns.Object("GSMover4")
	obj40 = ns.Object("GSMover5")
	obj41 = ns.Object("GSMover6")
	wp55 = ns.Waypoint("PlayerAudio")
	wp85 = ns.Waypoint("GearWay")
	wp86 = ns.Waypoint("StereoRWay")
	wp87 = ns.Waypoint("StereoLWay")
	wp88 = ns.Waypoint("BoneCruncherWay")
	wp89 = ns.Waypoint("ChickenRunsWay")
	wp73 = ns.Waypoint("F5way")
	wp75 = ns.Waypoint("F7way")
	wp77 = ns.Waypoint("F9way")
	wp82 = ns.Waypoint("GSWay1")
	wp83 = ns.Waypoint("GSWay3")
	wp84 = ns.Waypoint("GSWay5")
	gvar15 = ns.ObjectGroup("GauntletEntrySounds")
	ns.ObjectGroupOff(gvar15)
	ns.LockDoor(obj53)
	ns.LockDoor(obj48)
	ns.LockDoor(obj49)
	ns.LockDoor(obj52)
	obj91 = ns.Object("Spider1a")
	obj92 = ns.Object("Spider1b")
	ns.ObjectOff(obj91)
	ns.ObjectOff(obj92)
	obj93 = ns.Object("Arena1Door1a")
	obj94 = ns.Object("Arena1Door1b")
	obj95 = ns.Object("Arena1Door2a")
	obj96 = ns.Object("Arena1Door2b")
	obj97 = ns.Object("Arena1Switch")
	ns.LockDoor(obj95)
	ns.LockDoor(obj96)
	gvar98 = ns.WallGroup("Arena1Walls")
	initFlameGrates()
	obj207 = ns.Object("Urchin2a")
	obj208 = ns.Object("Urchin2b")
	ns.ObjectOff(obj207)
	ns.ObjectOff(obj208)
	obj209 = ns.Object("Arena2Door1a")
	obj210 = ns.Object("Arena2Door1b")
	obj211 = ns.Object("Arena2Door2a")
	obj212 = ns.Object("Arena2Door2b")
	obj213 = ns.Object("Arena2Switch")
	ns.LockDoor(obj211)
	ns.LockDoor(obj212)
	gvar214 = ns.WallGroup("Arena2Walls")
	obj220 = ns.Object("Arena3Door1a")
	obj221 = ns.Object("Arena3Door1b")
	obj222 = ns.Object("Arena3Door2a")
	obj223 = ns.Object("Arena3Door2b")
	obj224 = ns.Object("Arena3Urchin1")
	obj225 = ns.Object("Arena3Urchin2")
	obj226 = ns.Object("Arena3Urchin3")
	obj227 = ns.Object("Arena3Switch")
	ns.ObjectOff(obj224)
	ns.ObjectOff(obj225)
	ns.ObjectOff(obj226)
	ns.LockDoor(obj222)
	ns.LockDoor(obj223)
	gvar228 = ns.WallGroup("Arena3Walls")
	obj230 = ns.Object("Arena4Door1")
	obj231 = ns.Object("Arena4Door2")
	obj232 = ns.Object("Arena4Switch")
	gvar233 = ns.ObjectGroup("Arena4Elevators")
	gvar234 = ns.ObjectGroup("Arena4Spikes")
	ns.ObjectGroupOff(gvar234)
	gvar235 = ns.WallGroup("Arena4Cages")
	wp236 = ns.Waypoint("UrchinSoundsWay")
	gvar239 = ns.ObjectGroup("ZipperOfDeath")
	ns.ObjectGroupOff(gvar239)
	obj240 = ns.Object("ZipperDoor")
	ns.LockDoor(obj240)
	obj241 = ns.Object("am1")
	obj242 = ns.Object("am2")
	obj243 = ns.Object("am3")
	obj244 = ns.Object("am4")
	obj245 = ns.Object("am5")
	obj246 = ns.Object("am6")
	obj247 = ns.Object("am7")
	obj248 = ns.Object("am8")
	obj249 = ns.Object("am9")
	obj250 = ns.Object("am10")
	obj251 = ns.Object("am11")
	obj252 = ns.Object("am12")
	obj253 = ns.Object("am13")
	obj254 = ns.Object("am14")
	obj255 = ns.Object("am15")
	obj256 = ns.Object("am16")
	obj257 = ns.Object("am17")
	obj258 = ns.Object("am18")
	obj259 = ns.Object("am19")
	obj260 = ns.Object("am20")
	obj261 = ns.Object("asbm1")
	wp262 = ns.Waypoint("arsw1")
	wp263 = ns.Waypoint("arsw2")
	wp264 = ns.Waypoint("arsw3")
	wp265 = ns.Waypoint("arsw4")
	wp266 = ns.Waypoint("arsw5")
	wp267 = ns.Waypoint("arsw6")
	wp268 = ns.Waypoint("arsw7")
	wp269 = ns.Waypoint("arsw8")
	wp270 = ns.Waypoint("arsw9")
	wp271 = ns.Waypoint("arsw10")
	wp272 = ns.Waypoint("arsw11")
	wp273 = ns.Waypoint("arsw12")
	wp274 = ns.Waypoint("arsw13")
	wp275 = ns.Waypoint("arsw14")
	wp276 = ns.Waypoint("arsw15")
	wp277 = ns.Waypoint("arsw16")
	wp278 = ns.Waypoint("arsw17")
	wp279 = ns.Waypoint("arsw18")
	wp280 = ns.Waypoint("arsw19")
	wp281 = ns.Waypoint("arsw20")
	obj283 = ns.Object("FromGauntlet1")
	obj284 = ns.Object("FromGauntlet2")
	gvar285 = ns.ObjectGroup("FinishGauntletTriggers")
	ns.StartupScreen(2)
}
func PlayerDeath() {
	ns.DeathScreen(2)
}
func openGauntSecret() {
	ns.NoWallSound(false)
	ns.WallOpen(ns.Wall(119, 191))
	ns.WallOpen(ns.Wall(120, 192))
}
func urchins1Die() {
	ivar90 += 1
	twoDeadUrchins()
}
func fireGrateReward() {
	ns.GiveXp(ns.GetHost(), 100)
}
func fireballHallReward() {
	ns.GiveXp(ns.GetHost(), 250)
}
func unlockExitDoor() {
	ns.Music(22, 100)
	ns.UnlockDoor(obj283)
	ns.UnlockDoor(obj284)
	flag282 = true
	ns.GiveXp(ns.GetHost(), 2500)
	ns.MoveObject(obj43, 5071, 4894)
	ns.CreatureGuard(obj43, 5071, 4894, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	ns.MoveObject(obj44, 4335, 3151)
	ns.MoveObject(obj42, 4335, 3151)
}
func finishGauntlet() {
	if !flag282 {
		goto LABEL1
	}
	flag282 = false
	ns.ObjectGroupOff(gvar285)
	ns.LookAtObject(obj43, ns.GetHost())
	ns.WideScreen(true)
	ns.Frozen(ns.GetHost(), true)
	ns.SetDialog(obj43, "FALSE", endGauntletTalk1Start, endGauntletTalk1End)
	ns.StartDialog(obj43, ns.GetHost())
LABEL1:
	return
}
func OnEvent(typ string) {
	switch typ {
	case "MapInitialize":
		MapInitialize()
	case "PlayerDeath":
		PlayerDeath()
	}
}
