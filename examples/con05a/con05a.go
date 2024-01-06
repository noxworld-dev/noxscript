package con05a

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	gvar4   int
	gvar5   int
	gvar6   int
	gvar7   int
	gvar8   int
	gvar9   int
	gvar10  int
	gvar11  int
	gvar12  int
	gvar13  int
	gvar14  int
	gvar15  int
	gvar16  int
	gvar17  int
	gvar18  int
	gvar19  int
	gvar20  int
	gvar21  int
	gvar22  int
	gvar23  int
	gvar24  int
	gvar25  int
	gvar26  int
	gvar27  int
	gvar28  int
	gvar29  int
	wp30    ns.WaypointID
	wp31    ns.WaypointID
	wp32    ns.WaypointID
	wp33    ns.WaypointID
	wp34    ns.WaypointID
	wp35    ns.WaypointID
	wp36    ns.WaypointID
	wp37    ns.WaypointID
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
	obj54   ns.ObjectID
	obj55   ns.ObjectID
	gvar56  ns.ObjectID
	gvar57  ns.ObjectID
	obj58   ns.ObjectID
	obj59   ns.ObjectID
	obj60   ns.ObjectID
	obj61   ns.ObjectID
	obj62   ns.ObjectID
	obj63   ns.ObjectID
	gvar64  ns.ObjectGroupID
	gvar65  ns.ObjectGroupID
	gvar66  ns.ObjectID
	ivar67  int
	ivar68  int
	gvar69  int
	gvar70  int
	flag71  bool
	gvar72  int
	gvar73  int
	gvar74  int
	gvar75  int
	gvar76  int
	gvar77  int
	gvar78  int
	gvar79  int
	gvar80  int
	gvar81  int
	gvar82  int
	gvar83  int
	gvar84  int
	gvar85  int
	gvar86  int
	flag87  bool
	obj88   ns.ObjectID
	obj89   ns.ObjectID
	obj90   ns.ObjectID
	obj91   ns.ObjectID
	obj92   ns.ObjectID
	obj93   ns.ObjectID
	obj94   ns.ObjectID
	obj95   ns.ObjectID
	obj96   ns.ObjectID
	obj97   ns.ObjectID
	obj98   ns.ObjectID
	obj99   ns.ObjectID
	obj100  ns.ObjectID
	obj101  ns.ObjectID
	obj102  ns.ObjectID
	obj103  ns.ObjectID
	obj104  ns.ObjectID
	obj105  ns.ObjectID
	obj106  ns.ObjectID
	obj107  ns.ObjectID
	obj108  ns.ObjectID
	obj109  ns.ObjectID
	obj110  ns.ObjectID
	obj111  ns.ObjectID
	obj112  ns.ObjectID
	obj113  ns.ObjectID
	obj114  ns.ObjectID
	obj115  ns.ObjectID
	obj116  ns.ObjectID
	obj117  ns.ObjectID
	obj118  ns.ObjectID
	obj119  ns.ObjectID
	obj120  ns.ObjectID
	obj121  ns.ObjectID
	obj122  ns.ObjectID
	obj123  ns.ObjectID
	obj124  ns.ObjectID
	obj125  ns.ObjectID
	obj126  ns.ObjectID
	obj127  ns.ObjectID
	obj128  ns.ObjectID
	obj129  ns.ObjectID
	obj130  ns.ObjectID
	obj131  ns.ObjectID
	obj132  ns.ObjectID
	obj133  ns.ObjectID
	obj134  ns.ObjectID
	obj135  ns.ObjectID
	obj136  ns.ObjectID
	obj137  ns.ObjectID
	obj138  ns.ObjectID
	obj139  ns.ObjectID
	obj140  ns.ObjectID
	obj141  ns.ObjectID
	wp142   int
	wp143   int
	obj144  ns.ObjectID
	obj145  ns.ObjectID
	obj146  ns.ObjectID
	obj147  ns.ObjectID
	obj148  ns.ObjectID
	obj149  ns.ObjectID
	obj150  ns.ObjectID
	obj151  ns.ObjectID
	obj152  ns.ObjectID
	obj153  ns.ObjectID
	obj154  ns.ObjectID
	obj155  ns.ObjectID
	obj156  ns.ObjectID
	obj157  ns.ObjectID
	obj158  ns.ObjectID
	obj159  ns.ObjectID
	obj160  ns.ObjectID
	obj161  ns.ObjectID
	obj162  ns.ObjectID
	obj163  ns.ObjectID
	obj164  ns.ObjectID
	obj165  ns.ObjectID
	obj166  ns.ObjectID
	obj167  ns.ObjectID
	obj168  ns.ObjectID
	obj169  ns.ObjectID
	obj170  ns.ObjectID
	gvar171 int
	gvar172 int
	gvar173 int
	gvar174 int
	obj175  ns.ObjectID
	obj176  ns.ObjectID
	obj177  ns.ObjectID
	obj178  ns.ObjectID
	obj179  ns.ObjectID
	obj180  ns.ObjectID
	obj181  ns.ObjectID
	obj182  ns.ObjectID
	obj183  ns.ObjectID
	obj184  ns.ObjectID
	gvar185 int
	wp186   ns.WaypointID
	wp187   ns.WaypointID
	wp188   ns.WaypointID
	wp189   ns.WaypointID
	wp190   ns.WaypointID
	wp191   ns.WaypointID
	wp192   ns.WaypointID
	obj193  ns.ObjectID
	obj194  ns.ObjectID
	obj195  ns.ObjectID
	obj196  ns.ObjectID
	gvar197 int
	gvar198 int
	gvar199 int
	gvar200 int
	wp201   ns.WaypointID
	wp202   ns.WaypointID
	wp203   ns.WaypointID
	wp204   ns.WaypointID
	wp205   ns.WaypointID
	wp206   ns.WaypointID
	wp207   ns.WaypointID
	wp208   ns.WaypointID
	wp209   ns.WaypointID
	wp210   ns.WaypointID
	wp211   ns.WaypointID
	wp212   ns.WaypointID
	wp213   ns.WaypointID
	wp214   ns.WaypointID
	wp215   ns.WaypointID
	wp216   ns.WaypointID
	wp217   ns.WaypointID
	wp218   ns.WaypointID
	wp219   ns.WaypointID
	wp220   ns.WaypointID
	wp221   ns.WaypointID
	ivar222 int
)

func init() {
	gvar4 = 0
	gvar5 = 1
	gvar6 = 2
	gvar7 = 3
	gvar8 = 4
	gvar9 = 5
	gvar10 = 0
	gvar11 = 1
	gvar12 = 0
	gvar13 = 1
	gvar14 = 0
	gvar15 = 1
	gvar16 = 2
	gvar17 = 3
	gvar18 = 4
	gvar19 = 5
	gvar20 = 0
	gvar21 = 1
	gvar22 = 2
	gvar23 = 3
	gvar24 = 0
	gvar25 = 1
	gvar26 = 0
	gvar27 = 1
	gvar28 = 2
	gvar29 = 3
	ivar67 = 0
	ivar68 = 0
	gvar69 = 0
	gvar70 = 0
	flag71 = false
	gvar72 = gvar4
	gvar73 = gvar10
	gvar74 = gvar12
	gvar75 = gvar20
	gvar76 = gvar14
	gvar77 = gvar24
	gvar84 = 0
	gvar85 = 0
	gvar86 = 0
	ivar222 = 0
}
func OgreDeath() {
	ivar67 += 1
	if !(ivar67 >= 4) {
		goto LABEL1
	}
	SetTownDialog()
	ns.Music(23, 100)
LABEL1:
	return
}
func SetTownDialog() {
	ns.SetDialog(obj176, ns.NORMAL, Villager1Start, Villager1End)
	ns.StoryPic(obj176, "MalePic8")
	ns.SetDialog(obj177, ns.NORMAL, Villager2Start, Villager2End)
	ns.StoryPic(obj177, "MalePic3")
	ns.SetDialog(obj178, ns.NORMAL, Villager3Start, Villager3End)
	ns.StoryPic(obj178, "MalePic2")
	ns.SetDialog(obj179, ns.NORMAL, Villager4Start, Villager4End)
	ns.StoryPic(obj179, "MalePic14")
	ns.SetDialog(obj180, ns.NORMAL, Villager5Start, Villager5End)
	ns.StoryPic(obj180, "MalePic9")
}
func BeginMission() {
	var (
		v0 float32
		v1 float32
	)
	v0 = ns.GetWaypointX(wp209)
	v1 = ns.GetWaypointY(wp209)
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.ObjectOff(obj119)
	ns.ObjectOff(obj120)
	ns.ObjectOff(obj121)
	ns.Frozen(ns.GetHost(), true)
	ns.CreatureIdle(ns.GetHost())
	LetterBoxOn()
	ns.MoveWaypoint(wp33, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.TeleportIn, wp33)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	ns.Effect(ns.SPARK_EXPLOSION, v0, v1, 0, 0)
	ns.MoveObject(obj55, v0, v1)
	ns.Frozen(obj55, true)
	ns.CreatureIdle(obj55)
	ns.LookAtObject(obj55, ns.GetHost())
	ns.FrameTimer(5, BeginMission2)
LABEL1:
	return
}
func BeginMission2() {
	var (
		v0 float32
		v1 float32
	)
	v0 = ns.GetWaypointX(wp209)
	v1 = ns.GetWaypointY(wp209)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	ns.MoveWaypoint(wp207, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.FrameTimer(25, BeginMission3)
}
func BeginMission3() {
	ns.StartDialog(obj55, ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj55)
	ns.LookAtObject(obj55, ns.GetHost())
}
func BeginMission4() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
	)
	_ = v3
	_ = v2
	v0 = ns.GetWaypointX(wp209)
	v1 = ns.GetWaypointY(wp209)
	v2 = ns.GetWaypointX(wp207)
	v3 = ns.GetWaypointY(wp207)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	ns.Effect(ns.SPARK_EXPLOSION, v0, v1, 0, 0)
	ns.DestroyChat(obj55)
	ns.AudioEvent(ns.TeleportOut, wp33)
	ns.MoveObject(obj55, 816, 1980)
	ns.CreatureIdle(obj55)
	LetterBoxOff()
	ns.FrameTimer(15, BeginMission5)
}
func BeginMission5() {
	var (
		v0 float32
		v1 float32
	)
	v0 = ns.GetWaypointX(wp209)
	v1 = ns.GetWaypointY(wp209)
	ns.Frozen(ns.GetHost(), false)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	ns.Effect(ns.SPARK_EXPLOSION, v0, v1, 0, 0)
}
func HorvathStart() {
	var (
		v0 float32
		v1 float32
		v2 int
	)
	v2 = gvar69
	if v2 == gvar26 {
		goto LABEL1
	}
	if v2 == gvar27 {
		goto LABEL2
	}
	if v2 == gvar28 {
		goto LABEL3
	}
	if v2 == gvar29 {
		goto LABEL4
	}
	goto LABEL5
LABEL1:
	ns.TellStory(ns.SwordsmanHurt, "Con05A.scr:HorvathGreeting")
	goto LABEL5
LABEL2:
	if gvar70 != 1 {
		goto LABEL6
	}
	v0 = ns.GetWaypointX(wp207)
	v1 = ns.GetWaypointY(wp207)
	ns.AudioEvent(ns.CurePoisonCast, wp207)
	ns.AudioEvent(ns.CurePoisonEffect, wp207)
	ns.Effect(ns.BLUE_SPARKS, v0, v1, 0, 0)
	gvar69 = gvar28
	ns.TellStory(ns.SwordsmanHurt, "Con05A.scr:HorvathReturned")
	goto LABEL5
	goto LABEL3
LABEL6:
	ns.TellStory(ns.SwordsmanHurt, "Con05A.scr:HorvathWaiting")
	goto LABEL5
LABEL3:
	ns.TellStory(ns.SwordsmanHurt, "Con05A.scr:HorvathReturned")
	goto LABEL5
LABEL4:
	ns.TellStory(ns.SwordsmanHurt, "Con05A.scr:HorvathIdle")
	goto LABEL5
LABEL5:
	return
}
func HorvathEnd() {
	var v0 int
	v0 = gvar69
	if v0 == gvar26 {
		goto LABEL1
	}
	if v0 == gvar27 {
		goto LABEL2
	}
	if v0 == gvar28 {
		goto LABEL3
	}
	if v0 == gvar29 {
		goto LABEL4
	}
	goto LABEL5
LABEL1:
	ns.JournalEntry(ns.GetHost(), "Con05Quest1", 2)
	ns.PrintToAll("Con01a:NewJournalEntry")
	gvar69 = gvar27
	ns.FrameTimer(1, BeginMission4)
	goto LABEL5
LABEL2:
	goto LABEL5
LABEL3:
	gvar69 = gvar29
	ns.GiveXp(ns.GetHost(), 250)
	ns.JournalEdit(ns.GetHost(), "Con05Quest1", 4)
	ns.PrintToAll("Con05A.scr:ObjectiveComplete")
	goto LABEL5
LABEL4:
	goto LABEL5
LABEL5:
	return
}
func ThaviusOgreRecognize() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.CancelDialog(obj176)
	ns.CancelDialog(obj177)
	ns.CancelDialog(obj178)
	ns.CancelDialog(obj179)
	ns.CancelDialog(obj180)
LABEL1:
	return
}
func ThaviusOgreDeath() {
	ivar68 += 1
	if !(ivar68 >= 2 && ivar67 >= 4) {
		goto LABEL1
	}
	SetTownDialog()
LABEL1:
	return
}
func ChangeMaps() {
	ns.MoveObject(ns.GetHost(), 3402, 513)
}
func GetFarmerStaff() {
	ns.SetDialog(obj126, ns.NORMAL, FarmerStart, FarmerEnd)
	gvar72 = gvar7
}
func KillCorpses() {
	ns.FrameTimer(30, KillCorpses2)
}
func DisableVillagers() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
	)
	v0 = ns.GetWaypointX(wp188)
	v1 = ns.GetWaypointY(wp188)
	v2 = ns.GetWaypointX(wp189)
	v3 = ns.GetWaypointY(wp189)
	ns.CreatureGuard(obj126, v0, v1, v2, v3, 0)
}
func LockDoors() {
	ns.LockDoor(obj195)
	ns.LockDoor(obj196)
	ns.LockDoor(obj193)
	ns.LockDoor(obj194)
}
func SwordsmanStart() {
	var v0 int
	ns.Frozen(obj140, true)
	ns.CreatureIdle(obj140)
	ns.LookAtObject(obj140, ns.GetHost())
	v0 = gvar74
	if v0 == gvar12 {
		goto LABEL1
	}
	if v0 == gvar13 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.LookAtObject(obj140, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:SwordsmanGreeting")
	goto LABEL3
LABEL2:
	ns.LookAtObject(obj140, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:SwordsmanEnd")
	goto LABEL3
LABEL3:
	return
}
func SwordsmanEnd() {
	var v0 int
	ns.Frozen(obj140, false)
	ns.CreatureGuard(obj140, ns.GetWaypointX(wp186), ns.GetWaypointY(wp186), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 240)
	v0 = gvar74
	if v0 == gvar12 {
		goto LABEL1
	}
	if v0 == gvar13 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.FrameTimer(1, SwordsmanTriggerC)
	gvar74 = gvar13
	goto LABEL3
LABEL2:
	ns.LookAtDirection(obj140, ns.NE)
	goto LABEL3
LABEL3:
	return
}
func DrunkStart() {
	ns.LookAtObject(ns.GetTrigger(), ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:DrunkGreeting")
}
func DrunkEnd() {
	ns.LookAtDirection(ns.GetTrigger(), ns.NW)
}
func FarmerStart() {
	var (
		v0 float32
		v1 float32
		v2 int
	)
	_ = v1
	_ = v0
	v2 = gvar72
	if v2 == gvar4 {
		goto LABEL1
	}
	if v2 == gvar5 {
		goto LABEL2
	}
	if v2 == gvar6 {
		goto LABEL3
	}
	if v2 == gvar7 {
		goto LABEL4
	}
	if v2 == gvar8 {
		goto LABEL5
	}
	if v2 == gvar9 {
		goto LABEL6
	}
	goto LABEL7
LABEL1:
	if !ns.HasItem(ns.GetHost(), obj125) {
		goto LABEL8
	}
	gvar72 = gvar7
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:FarmerReturned")
	goto LABEL7
	goto LABEL2
LABEL8:
	ns.CreatureIdle(obj126)
	ns.LookAtObject(obj126, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:FarmerGreeting")
	goto LABEL7
LABEL2:
	ns.CreatureIdle(obj126)
	ns.LookAtObject(obj126, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:FarmerGreeting")
	goto LABEL7
LABEL3:
	ns.CreatureIdle(obj126)
	ns.LookAtObject(obj126, ns.GetHost())
	if !ns.HasItem(ns.GetHost(), obj125) {
		goto LABEL9
	}
	v0 = ns.GetWaypointX(wp188)
	v1 = ns.GetWaypointY(wp188)
	gvar72 = gvar7
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:FarmerReturned")
	goto LABEL7
	goto LABEL4
LABEL9:
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:FarmerWaiting")
	goto LABEL7
LABEL4:
	ns.CreatureIdle(obj126)
	ns.LookAtObject(obj126, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:FarmerReturned")
	goto LABEL7
LABEL5:
	ns.CreatureIdle(obj126)
	ns.LookAtObject(obj126, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:FarmerIdle")
	goto LABEL7
LABEL6:
	if !ns.HasItem(ns.GetHost(), obj125) {
		goto LABEL10
	}
	gvar72 = gvar7
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:FarmerReturned")
	goto LABEL7
	goto LABEL7
LABEL10:
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:FarmerHuffy")
	goto LABEL7
LABEL7:
	return
}
func FarmerEnd() {
	var (
		v0 int
		v1 int
		v2 int
		v3 ns.ObjectID
		v4 int
	)
	v0 = 0
	v1 = 1
	v2 = 2
	v3 = 0
	v0 = ns.GetAnswer(obj126)
	v4 = gvar72
	if v4 == gvar4 {
		goto LABEL1
	}
	if v4 == gvar5 {
		goto LABEL2
	}
	if v4 == gvar6 {
		goto LABEL3
	}
	if v4 == gvar7 {
		goto LABEL4
	}
	if v4 == gvar8 {
		goto LABEL5
	}
	if v4 == gvar9 {
		goto LABEL6
	}
	goto LABEL7
LABEL1:
	if v0 != v1 {
		goto LABEL8
	}
	gvar72 = gvar6
	ns.SetDialog(obj126, ns.NORMAL, FarmerStart, FarmerEnd)
	if flag71 {
		goto LABEL9
	}
	ns.PrintToAll("Con01a:NewJournalEntry")
	ns.JournalEntry(ns.GetHost(), "War05Quest8", 2)
	flag71 = true
LABEL9:
	goto LABEL7
	goto LABEL2
LABEL8:
	if v0 != v2 {
		goto LABEL10
	}
	ns.SetDialog(obj126, ns.YESNO, FarmerStart, FarmerEnd)
	gvar72 = gvar5
	goto LABEL7
	goto LABEL2
LABEL10:
	if obj126 != v3 {
		goto LABEL2
	}
	gvar72 = gvar4
	ns.SetDialog(obj126, ns.YESNO, FarmerStart, FarmerEnd)
	goto LABEL7
LABEL2:
	if v0 != v1 {
		goto LABEL11
	}
	gvar72 = gvar6
	ns.SetDialog(obj126, ns.NORMAL, FarmerStart, FarmerEnd)
	if flag71 {
		goto LABEL12
	}
	ns.PrintToAll("Con01a:NewJournalEntry")
	ns.JournalEntry(ns.GetHost(), "War05Quest8", 2)
	flag71 = true
LABEL12:
	goto LABEL7
	goto LABEL3
LABEL11:
	if v0 != v2 {
		goto LABEL13
	}
	ns.SetDialog(obj126, ns.NORMAL, FarmerStart, FarmerEnd)
	gvar72 = gvar9
	goto LABEL7
	goto LABEL3
LABEL13:
	if obj126 != v3 {
		goto LABEL3
	}
	gvar72 = gvar4
	ns.SetDialog(obj126, ns.YESNO, FarmerStart, FarmerEnd)
	goto LABEL7
LABEL3:
	goto LABEL7
LABEL4:
	if !flag71 {
		goto LABEL14
	}
	ns.JournalEdit(ns.GetHost(), "War05Quest8", 4)
	ns.PrintToAll("War05A.scr:ObjectiveComplete")
LABEL14:
	ns.GiveXp(ns.GetHost(), 250)
	ns.Pickup(ns.GetHost(), obj123)
	ns.Delete(obj125)
	EndCurse()
	gvar72 = gvar8
	goto LABEL7
LABEL5:
	goto LABEL7
LABEL6:
	goto LABEL7
LABEL7:
	return
}
func HoundStart() {
	ns.CreatureIdle(obj127)
	ns.LookAtObject(obj127, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:HoundGreeting")
}
func HoundEnd() {
	ns.SetRoamFlag(obj127, 2)
	ns.Wander(obj127)
}
func FarmerWifeStart() {
	var v0 int
	v0 = gvar73
	if v0 == gvar10 {
		goto LABEL1
	}
	if v0 == gvar11 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.LookAtObject(obj128, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:FarmerWifeGreeting")
	goto LABEL3
LABEL2:
	ns.LookAtObject(obj128, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:FarmerWifeIdle")
	goto LABEL3
LABEL3:
	return
}
func FarmerWifeEnd() {
	var v0 int
	v0 = gvar73
	if v0 == gvar10 {
		goto LABEL1
	}
	if v0 == gvar11 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	gvar73 = gvar11
	LetterBoxOff()
	ns.Frozen(ns.GetHost(), false)
	goto LABEL3
LABEL2:
	goto LABEL3
LABEL3:
	return
}
func CaptainStart() {
	var v2 int
	v2 = gvar77
	if v2 == gvar24 {
		goto LABEL1
	}
	if v2 == gvar25 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:CaptainGreeting")
	goto LABEL3
LABEL2:
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:CaptainWaiting")
	goto LABEL3
LABEL3:
	return
}
func CaptainEnd() {
	var v0 int
	v0 = gvar77
	if v0 == gvar24 {
		goto LABEL1
	}
	if v0 == gvar25 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	gvar77 = gvar25
	ns.FrameTimer(1, HorvathAppears4)
	goto LABEL3
LABEL2:
	goto LABEL3
LABEL3:
	return
}
func BartenderStart() {
	ns.Frozen(obj94, true)
	ns.LookAtObject(obj94, ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj94)
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:BartenderGreeting")
}
func BartenderEnd() {
	ns.Frozen(obj94, false)
	ns.LookAtDirection(obj94, ns.SE)
}
func EndCurse() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
		v4 float32
		v5 float32
		v6 float32
		v7 float32
		v8 float32
		v9 float32
	)
	_ = v7
	_ = v6
	_ = v3
	_ = v2
	_ = v1
	_ = v0
	v8 = ns.GetWaypointX(wp188)
	v9 = ns.GetWaypointY(wp188)
	v0 = ns.GetWaypointX(wp206)
	v1 = ns.GetWaypointY(wp206)
	v2 = ns.GetWaypointX(wp204)
	v3 = ns.GetWaypointY(wp204)
	v6 = ns.GetWaypointX(wp205)
	v7 = ns.GetWaypointY(wp205)
	v4 = ns.GetObjectX(obj127)
	v5 = ns.GetObjectY(obj127)
	ns.Frozen(ns.GetHost(), true)
	LetterBoxOn()
	ns.AudioEvent(ns.CurePoisonCast, wp188)
	ns.AudioEvent(ns.CurePoisonEffect, wp188)
	ns.Effect(ns.LIGHTNING, v8, v9, v4, v5)
	ns.FrameTimer(10, EndCurse2)
}
func EndCurse2() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
		v4 float32
		v5 float32
		v6 float32
		v7 float32
		v8 float32
		v9 float32
	)
	_ = v7
	_ = v6
	_ = v3
	_ = v2
	_ = v1
	_ = v0
	v8 = ns.GetWaypointX(wp188)
	v9 = ns.GetWaypointY(wp188)
	v0 = ns.GetWaypointX(wp206)
	v1 = ns.GetWaypointY(wp206)
	v2 = ns.GetWaypointX(wp204)
	v3 = ns.GetWaypointY(wp204)
	v6 = ns.GetWaypointX(wp205)
	v7 = ns.GetWaypointY(wp205)
	v4 = ns.GetObjectX(obj127)
	v5 = ns.GetObjectY(obj127)
	ns.Effect(ns.LIGHTNING, v8, v9, v4, v5)
	ns.FrameTimer(7, EndCurse3)
}
func EndCurse3() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
		v4 float32
		v5 float32
		v6 float32
		v7 float32
		v8 float32
		v9 float32
	)
	_ = v7
	_ = v6
	_ = v3
	_ = v2
	_ = v1
	_ = v0
	v8 = ns.GetWaypointX(wp188)
	v9 = ns.GetWaypointY(wp188)
	v0 = ns.GetWaypointX(wp206)
	v1 = ns.GetWaypointY(wp206)
	v2 = ns.GetWaypointX(wp204)
	v3 = ns.GetWaypointY(wp204)
	v6 = ns.GetWaypointX(wp205)
	v7 = ns.GetWaypointY(wp205)
	v4 = ns.GetObjectX(obj127)
	v5 = ns.GetObjectY(obj127)
	ns.Effect(ns.LIGHTNING, v8, v9, v4, v5)
	ns.FrameTimer(3, EndCurse4)
}
func EndCurse4() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
		v4 float32
		v5 float32
		v6 float32
		v7 float32
		v8 float32
		v9 float32
	)
	v8 = ns.GetWaypointX(wp188)
	v9 = ns.GetWaypointY(wp188)
	v0 = ns.GetWaypointX(wp206)
	v1 = ns.GetWaypointY(wp206)
	v2 = ns.GetWaypointX(wp204)
	v3 = ns.GetWaypointY(wp204)
	v6 = ns.GetWaypointX(wp205)
	v7 = ns.GetWaypointY(wp205)
	v4 = ns.GetObjectX(obj127)
	v5 = ns.GetObjectY(obj127)
	ns.Effect(ns.BLUE_SPARKS, v4, v5, 0, 0)
	ns.MoveObject(obj127, v0, v1)
	ns.MoveObject(obj128, v4, v5)
	ns.CreatureGuard(obj128, v2, v3, v6, v7, 0)
	ns.CreatureGuard(obj126, v8, v9, v8, v9, 0)
}
func EndCurse5() {
	if !ns.IsCaller(obj128) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.LookAtObject(obj128, ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj128)
	ns.StartDialog(obj128, ns.GetHost())
LABEL1:
	return
}
func FrogStart() {
	var v0 int
	v0 = gvar75
	if v0 == gvar20 {
		goto LABEL1
	}
	if v0 == gvar21 {
		goto LABEL2
	}
	if v0 == gvar22 {
		goto LABEL3
	}
	if v0 == gvar23 {
		goto LABEL4
	}
	goto LABEL5
LABEL1:
	ns.LookAtObject(obj130, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:FrogGreeting")
	ns.Frozen(obj130, false)
	ns.ObjectOff(obj132)
	ns.ObjectOff(obj133)
	ns.ObjectOff(obj134)
	ns.ObjectOn(obj130)
	ns.CreatureFollow(obj130, ns.GetHost())
	goto LABEL5
LABEL2:
	ns.LookAtObject(obj130, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:FrogFollowing")
	goto LABEL5
LABEL3:
	ns.LookAtObject(obj130, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:FrogThanks")
	goto LABEL5
LABEL4:
	ns.LookAtObject(obj130, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:FrogIdle")
	goto LABEL5
LABEL5:
	return
}
func FrogEnd() {
	var v0 int
	v0 = gvar75
	if v0 == gvar20 {
		goto LABEL1
	}
	if v0 == gvar21 {
		goto LABEL2
	}
	if v0 == gvar22 {
		goto LABEL3
	}
	if v0 == gvar23 {
		goto LABEL4
	}
	goto LABEL5
LABEL1:
	gvar75 = gvar21
	goto LABEL5
LABEL2:
	goto LABEL5
LABEL3:
	gvar75 = gvar23
	goto LABEL5
LABEL4:
	goto LABEL5
LABEL5:
	return
}
func FrogGuard() {
	gvar75 = gvar22
	ns.Enchant(obj130, ns.ENCHANT_INVULNERABLE, 0)
	gvar86 = 1
}
func Villager1Start() {
	ns.Frozen(obj176, true)
	ns.CreatureIdle(obj176)
	ns.LookAtObject(obj176, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con05:Townsman9Talk01")
}
func Villager1End() {
	ns.Frozen(obj176, false)
	ns.SetRoamFlag(obj176, 32)
	ns.Wander(obj176)
}
func Villager2Start() {
	ns.Frozen(obj177, true)
	ns.CreatureIdle(obj177)
	ns.LookAtObject(obj177, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con05:Townsman10Talk01")
}
func Villager2End() {
	ns.Frozen(obj177, false)
	ns.SetRoamFlag(obj177, 32)
	ns.Wander(obj177)
}
func Villager3Start() {
	ns.Frozen(obj178, true)
	ns.CreatureIdle(obj178)
	ns.LookAtObject(obj178, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con05:Townsman11Talk01")
}
func Villager3End() {
	ns.Frozen(obj178, false)
	ns.SetRoamFlag(obj178, 32)
	ns.Wander(obj178)
}
func Villager4Start() {
	ns.Frozen(obj179, true)
	ns.CreatureIdle(obj179)
	ns.LookAtObject(obj179, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con05:Townsman11Talk02")
}
func Villager4End() {
	ns.Frozen(obj179, false)
	ns.SetRoamFlag(obj179, 32)
	ns.Wander(obj179)
}
func Villager5Start() {
	ns.Frozen(obj180, true)
	ns.CreatureIdle(obj180)
	ns.LookAtObject(obj180, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:VillagerGreeting")
}
func Villager5End() {
	ns.Frozen(obj180, false)
	ns.SetRoamFlag(obj180, 32)
	ns.Wander(obj180)
}
func SadVillagerStart() {
	var (
		v0 float32
		v1 float32
		v2 int
	)
	_ = v1
	_ = v0
	v2 = gvar76
	if v2 == gvar14 {
		goto LABEL1
	}
	if v2 == gvar15 {
		goto LABEL2
	}
	if v2 == gvar16 {
		goto LABEL3
	}
	if v2 == gvar17 {
		goto LABEL4
	}
	if v2 == gvar18 {
		goto LABEL5
	}
	if v2 == gvar19 {
		goto LABEL6
	}
	goto LABEL7
LABEL1:
	ns.CreatureIdle(obj129)
	ns.LookAtObject(obj129, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:SadVillagerGreeting")
	goto LABEL7
LABEL2:
	ns.CreatureIdle(obj129)
	ns.LookAtObject(obj129, ns.GetHost())
	if gvar86 != 1 {
		goto LABEL8
	}
	v0 = ns.GetWaypointX(wp201)
	v1 = ns.GetWaypointY(wp201)
	ns.AudioEvent(ns.CurePoisonCast, wp201)
	ns.AudioEvent(ns.CurePoisonEffect, wp201)
	gvar76 = gvar16
	ns.SetDialog(obj129, ns.NEXT, SadVillagerStart, SadVillagerEnd)
	ns.CreatureFollow(obj130, obj129)
	SadVillagerStart()
	goto LABEL3
LABEL8:
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:SadVillagerWaiting")
	goto LABEL7
LABEL3:
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:SadVillagerReturned")
	goto LABEL7
LABEL4:
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:SadVillagerIdle")
	goto LABEL7
LABEL5:
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:SadVillagerDead")
	goto LABEL7
LABEL6:
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:SadVillagerDeadIdle")
	goto LABEL7
LABEL7:
	return
}
func SadVillagerEnd() {
	var (
		v0 int
		v1 int
		v2 int
		v3 int
		v4 int
	)
	_ = v3
	_ = v2
	_ = v1
	_ = v0
	v0 = 0
	v1 = 1
	v2 = 2
	v3 = 0
	v0 = ns.GetAnswer(obj129)
	v4 = gvar76
	if v4 == gvar14 {
		goto LABEL1
	}
	if v4 == gvar15 {
		goto LABEL2
	}
	if v4 == gvar16 {
		goto LABEL3
	}
	if v4 == gvar17 {
		goto LABEL4
	}
	if v4 == gvar18 {
		goto LABEL5
	}
	if v4 == gvar19 {
		goto LABEL6
	}
	goto LABEL7
LABEL1:
	v0 = 1
	if 1 == 0 {
		goto LABEL8
	}
	gvar76 = gvar15
	ns.SetDialog(obj129, ns.NORMAL, SadVillagerStart, SadVillagerEnd)
	ns.JournalEntry(ns.GetHost(), "War05Quest9", 2)
	ns.PrintToAll("Con01a:NewJournalEntry")
	goto LABEL7
LABEL8:
	v0 = 2
	if 2 == 0 {
		goto LABEL9
	}
	gvar76 = gvar14
	ns.SetDialog(obj129, ns.YESNO, SadVillagerStart, SadVillagerEnd)
	goto LABEL7
LABEL9:
	v0 = 0
	if 0 == 0 {
		goto LABEL2
	}
	gvar76 = gvar14
	ns.SetDialog(obj129, ns.YESNO, SadVillagerStart, SadVillagerEnd)
	goto LABEL7
LABEL2:
	ns.SetDialog(obj129, ns.NORMAL, SadVillagerStart, SadVillagerEnd)
	goto LABEL7
LABEL3:
	gvar76 = gvar17
	ns.SetDialog(obj129, ns.NORMAL, SadVillagerStart, SadVillagerEnd)
	ns.JournalEdit(ns.GetHost(), "War05Quest9", 4)
	ns.PrintToAll("War05A.scr:ObjectiveComplete")
	ns.GiveXp(ns.GetHost(), 250)
	ns.Drop(obj129, obj122)
	ns.FrameTimer(1, FrogThanksPlayer)
	ns.FrameTimer(1, SadVillagerGuard)
	goto LABEL7
LABEL4:
	ns.SetDialog(obj129, ns.NORMAL, SadVillagerStart, SadVillagerEnd)
	goto LABEL7
LABEL5:
	gvar76 = gvar19
	ns.Move(obj129, wp201)
	ns.SetCallback(obj129, 9, DummyBump)
	ns.JournalDelete(ns.GetHost(), "War05Quest9")
	ns.PrintToAll("War05A.scr:ObjectiveFailed")
	ns.SetDialog(obj129, ns.NORMAL, SadVillagerStart, SadVillagerEnd)
	goto LABEL7
LABEL6:
	ns.CancelDialog(obj129)
	goto LABEL7
LABEL7:
	return
}
func OwnVillagers() {
	ns.FrameTimer(1, OwnVillagers2)
}
func OwnVillagers2() {
	ns.SetOwner(ns.GetHost(), obj130)
	ns.SetOwner(ns.GetHost(), obj55)
	ns.SetOwner(ns.GetHost(), obj127)
	ns.SetOwner(ns.GetHost(), obj94)
	ns.SetOwner(ns.GetHost(), obj128)
	ns.Frozen(obj176, true)
	ns.Frozen(obj177, true)
	ns.Frozen(obj178, true)
	ns.Frozen(obj179, true)
	ns.Frozen(obj180, true)
	ns.Frozen(obj181, true)
	ns.Frozen(obj182, true)
	ns.Frozen(obj183, true)
	ns.Frozen(obj184, true)
	ns.Frozen(obj49, true)
	ns.Frozen(obj50, true)
	ns.Frozen(obj47, true)
	ns.Frozen(obj48, true)
	ns.Frozen(obj51, true)
	ns.Wander(obj127)
	ns.RetreatLevel(obj46, 0)
	ns.RetreatLevel(obj47, 0)
	ns.RetreatLevel(obj48, 0)
	ns.RetreatLevel(obj49, 0)
	ns.RetreatLevel(obj50, 0)
}
func KillCorpses2() {
	ns.Damage(obj144, 0, 300, 5)
	ns.Damage(obj145, 0, 300, 5)
	ns.Damage(obj146, 0, 300, 5)
	ns.Damage(obj147, 0, 300, 5)
	ns.Damage(obj148, 0, 300, 5)
	ns.Damage(obj149, 0, 300, 5)
	ns.Damage(obj150, 0, 300, 5)
	ns.Damage(obj151, 0, 300, 5)
	ns.Damage(obj152, 0, 300, 5)
	ns.Damage(obj153, 0, 300, 5)
	ns.Damage(obj154, 0, 300, 5)
	ns.Damage(obj155, 0, 300, 5)
	ns.Damage(obj156, 0, 300, 5)
	ns.Damage(obj157, 0, 300, 5)
	ns.Damage(obj158, 0, 300, 5)
	ns.Damage(obj159, 0, 300, 0)
	ns.Damage(obj160, 0, 300, 0)
	ns.Damage(obj161, 0, 300, 0)
	ns.Damage(obj139, 0, 300, 0)
	ns.Damage(obj44, 0, 300, 0)
	ns.Damage(obj45, 0, 300, 0)
	ns.Damage(obj58, 0, 300, 0)
	ns.Damage(obj59, 0, 300, 0)
	ns.Damage(obj60, 0, 300, 0)
	ns.Damage(obj61, 0, 300, 0)
	ns.Damage(obj62, 0, 300, 0)
	ns.Damage(obj63, 0, 300, 0)
}
func VillagePath() {
	ns.SetRoamFlag(obj176, 32)
	ns.SetRoamFlag(obj177, 32)
	ns.SetRoamFlag(obj178, 32)
	ns.SetRoamFlag(obj179, 32)
	ns.SetRoamFlag(obj180, 32)
	ns.Wander(obj176)
	ns.Wander(obj177)
	ns.Wander(obj178)
	ns.Wander(obj179)
	ns.Wander(obj180)
}
func PickUpObjects() {
	ns.Pickup(obj129, obj122)
	ns.Pickup(obj126, obj123)
	ns.Pickup(obj126, obj124)
}
func EnableElevator() {
	ns.ObjectOn(obj175)
}
func DisableElevator() {
	ns.ObjectOff(obj175)
}
func StartSwordsmanTrigger() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.FrameTimer(1, SwordsmanTrigger)
LABEL1:
	return
}
func SwordsmanTrigger() {
	ns.AggressionLevel(obj140, 0.83)
	ns.AggressionLevel(obj52, 0.83)
	ns.AggressionLevel(obj46, 0.83)
	ns.Attack(obj46, obj140)
	ns.CreatureHunt(obj46)
}
func SwordsmanTriggerB() {
	gvar74 = gvar12
	ns.SetDialog(obj140, ns.NORMAL, SwordsmanStart, SwordsmanEnd)
	ns.StoryPic(obj140, "Warrior8Pic")
	ns.FrameTimer(1, SwordsmanTriggerC)
}
func SwordsmanTriggerC() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
	)
	v0 = ns.GetWaypointX(wp186)
	v1 = ns.GetWaypointY(wp186)
	v2 = ns.GetWaypointX(wp187)
	v3 = ns.GetWaypointY(wp187)
	ns.CreatureGuard(obj140, v0, v1, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 240)
	ns.CreatureGuard(obj52, v2, v3, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 240)
}
func ShopKeeperGuard() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
	)
	v0 = ns.GetWaypointX(wp218)
	v1 = ns.GetWaypointY(wp218)
	v2 = ns.GetWaypointX(wp219)
	v3 = ns.GetWaypointY(wp219)
	ns.CreatureGuard(obj95, v0, v1, v2, v3, 0)
	ns.FrameTimer(1, GypsyGuard)
}
func BartenderGuard() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
	)
	v0 = ns.GetWaypointX(wp220)
	v1 = ns.GetWaypointY(wp220)
	v2 = ns.GetWaypointX(wp221)
	v3 = ns.GetWaypointY(wp221)
	ns.CreatureGuard(obj94, v0, v1, v2, v3, 0)
	ns.FrameTimer(1, SadVillagerGuard)
}
func GypsyGuard() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
	)
	v0 = ns.GetWaypointX(wp216)
	v1 = ns.GetWaypointY(wp216)
	v2 = ns.GetWaypointX(wp217)
	v3 = ns.GetWaypointY(wp217)
	ns.CreatureGuard(obj93, v0, v1, v2, v3, 0)
	ns.FrameTimer(1, BartenderGuard)
}
func SadVillagerGuard() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
	)
	v0 = ns.GetWaypointX(wp201)
	v1 = ns.GetWaypointY(wp201)
	v2 = ns.GetWaypointX(wp202)
	v3 = ns.GetWaypointY(wp202)
	ns.CreatureGuard(obj129, v0, v1, v2, v3, 0)
}
func OpenGates() {
	ns.WallOpen(ns.Wall(88, 34))
	ns.WallOpen(ns.Wall(87, 35))
	ns.WallOpen(ns.Wall(88, 36))
	ns.WallOpen(ns.Wall(89, 35))
	ns.PrintToAll("War05A.scr:CityGatesOpen")
}
func CloseGates() {
	ns.WallClose(ns.Wall(88, 34))
	ns.WallClose(ns.Wall(87, 35))
	ns.WallClose(ns.Wall(88, 36))
	ns.WallClose(ns.Wall(89, 35))
	ns.PrintToAll("War05A.scr:CityGatesClosed")
}
func PlayerDeath() {
	ns.PrintToAll("War05A.scr:PlayerDeath")
	ns.DeathScreen(5)
}
func MapEntry() {
	ns.Music(23, 100)
	ns.Frozen(ns.GetHost(), false)
	LetterBoxOff()
	ns.NoWallSound(false)
	gvar66 = ns.Object("Con05B:HorvathAmulet")
	flag87 = true
	if !ns.HasItem(ns.GetHost(), gvar66) {
		goto LABEL1
	}
	ns.GroupDamage(gvar65, 0, 1000, 0)
LABEL1:
	return
}
func NoBump() {
	flag87 = false
	ns.FrameTimer(90, YesBump)
}
func YesBump() {
	flag87 = true
}
func BumpVillager() {
}
func LetterBoxOn() {
	ns.WideScreen(true)
}
func LetterBoxOff() {
	ns.WideScreen(false)
}
func FrogThanksPlayer() {
	ns.StartDialog(obj130, ns.GetHost())
}
func DummyBump() {
}
func FreePlayer() {
	ns.Frozen(ns.GetHost(), false)
}
func HorvathAppears() {
	ns.Frozen(ns.GetHost(), true)
	LetterBoxOn()
	ns.FrameTimer(30, HorvathAppears2)
	ns.FrameTimer(1, VillagePath)
}
func HorvathAppears2() {
	ns.Move(obj92, wp207)
	ns.FrameTimer(45, HorvathAppears3)
}
func HorvathAppears3() {
	ns.Frozen(obj92, true)
	ns.CreatureIdle(obj92)
	ns.CreatureIdle(ns.GetHost())
	ns.LookAtObject(obj92, ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj92)
	ns.StartDialog(obj92, ns.GetHost())
}
func HorvathAppears4() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
	)
	v0 = ns.GetWaypointX(wp209)
	v1 = ns.GetWaypointY(wp209)
	v2 = ns.GetWaypointX(wp207)
	v3 = ns.GetWaypointY(wp207)
	ns.Frozen(obj92, false)
	ns.CreatureGuard(obj92, v0, v1, v2, v3, 0)
	LetterBoxOff()
	ns.FrameTimer(45, HorvathAppears5)
}
func HorvathAppears5() {
	ns.Frozen(ns.GetHost(), false)
}
func AmuletCheck() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	if !ns.HasItem(ns.GetHost(), gvar66) {
		goto LABEL1
	}
	ns.Frozen(ns.GetHost(), true)
	ns.WideScreen(true)
	ns.CreatureIdle(ns.GetHost())
	ns.LookAtObject(ns.GetHost(), ns.ObjectID(wp36)) // FIXME
	ns.MoveWaypoint(wp33, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.TeleportIn, wp33)
	ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp36), ns.GetWaypointY(wp36), 0, 0)
	ns.Effect(ns.SPARK_EXPLOSION, ns.GetWaypointX(wp36), ns.GetWaypointY(wp36), 0, 0)
	ns.MoveObject(obj55, ns.GetWaypointX(wp36), ns.GetWaypointY(wp36))
	ns.Frozen(obj55, true)
	ns.CreatureIdle(obj55)
	ns.LookAtObject(obj55, ns.GetHost())
	ns.SetDialog(obj55, ns.NORMAL, HorvathCongratStart, HorvathCongratEnd)
	ns.StartDialog(obj55, ns.GetHost())
LABEL1:
	return
}
func HorvathCongratStart() {
	ns.TellStory(ns.HumanMaleEatFood, "Con05A.scr:HorvathReturned")
}
func HorvathCongratEnd() {
	ns.CancelDialog(obj55)
	gvar56 = ns.GetLastItem(ns.GetHost())
	for {
		if gvar56 == 0 {
			goto LABEL1
		}
		gvar57 = ns.GetPreviousItem(gvar56)
		if gvar56 != gvar66 {
			goto LABEL2
		}
		ns.Delete(gvar56)
	LABEL2:
		gvar56 = gvar57
	}
LABEL1:
	ns.Pickup(ns.GetHost(), obj41)
	ns.MoveObject(obj42, ns.GetObjectX(obj55)-15, ns.GetObjectY(obj55)+15)
	ns.Effect(ns.SPARK_EXPLOSION, ns.GetObjectX(obj42), ns.GetObjectY(obj42), 0, 0)
	ns.MoveObject(obj43, ns.GetObjectX(obj55)-15, ns.GetObjectY(obj55)-15)
	ns.Effect(ns.SPARK_EXPLOSION, ns.GetObjectX(obj43), ns.GetObjectY(obj43), 0, 0)
	ns.JournalEdit(ns.GetHost(), "Con05Quest1", 4)
	ns.GiveXp(ns.GetHost(), 1000)
	ns.Blind()
	ns.WideScreen(false)
	ns.Frozen(ns.GetHost(), false)
	ns.FrameTimer(45, ChangeMaps)
}
func MapInitialize() {
	gvar65 = ns.ObjectGroup("ExitEnemies")
	wp30 = ns.Waypoint("WoundedRefugeeWayPoint")
	wp31 = ns.Waypoint("WoundedRefugeeWayPoint2")
	wp32 = ns.Waypoint("DeadGuyWP")
	wp36 = ns.Waypoint("CongratWP")
	wp33 = ns.Waypoint("PlayerSounds")
	wp34 = ns.Waypoint("LunchOgre1WP")
	wp35 = ns.Waypoint("LunchOgre2WP")
	wp37 = ns.Waypoint("WellWP")
	obj38 = ns.Object("WoundedRefugee")
	obj39 = ns.Object("WoundedRefugee2")
	obj40 = ns.Object("Antidote")
	obj41 = ns.Object("InfravisionBook")
	obj42 = ns.Object("ManaBottle1")
	obj43 = ns.Object("ManaBottle2")
	obj44 = ns.Object("Kill")
	obj45 = ns.Object("Kill2")
	obj46 = ns.Object("FirstOgre")
	obj47 = ns.Object("LunchOgre1")
	obj48 = ns.Object("LunchOgre2")
	obj49 = ns.Object("TownOgre1")
	obj50 = ns.Object("TownOgre2")
	obj51 = ns.Object("DeadGuy")
	obj52 = ns.Object("Swordsman2")
	obj53 = ns.Object("Archer1")
	obj54 = ns.Object("Archer2")
	obj55 = ns.Object("Horvath")
	obj58 = ns.Object("Bubbles1")
	obj59 = ns.Object("Bubbles2")
	obj60 = ns.Object("Bubbles3")
	obj61 = ns.Object("Bubbles4")
	obj62 = ns.Object("Bubbles5")
	obj63 = ns.Object("Bubbles6")
	gvar64 = ns.ObjectGroup("OgreAttackTriggers")
	obj92 = ns.Object("W5Captain")
	obj94 = ns.Object("W5Bartender")
	obj95 = ns.Object("W5ShopKeeper")
	obj93 = ns.Object("W5Gypsy")
	obj119 = ns.Object("CaptainTrigger1")
	obj120 = ns.Object("CaptainTrigger2")
	obj121 = ns.Object("CaptainTrigger3")
	obj96 = ns.Object("BartenderTrigger1")
	obj97 = ns.Object("BartenderTrigger2")
	obj98 = ns.Object("ShopKeeperTrigger1")
	obj99 = ns.Object("ShopKeeperTrigger2")
	obj100 = ns.Object("GypsyTrigger1")
	obj101 = ns.Object("GypsyTrigger2")
	obj102 = ns.Object("GypsyTrigger3")
	obj103 = ns.Object("GypsyTrigger4")
	obj104 = ns.Object("GypsyTrigger5")
	obj105 = ns.Object("GypsyTrigger6")
	obj106 = ns.Object("GypsyTrigger7")
	obj107 = ns.Object("GypsyTrigger8")
	obj108 = ns.Object("GypsyTrigger9")
	obj109 = ns.Object("GypsyEndTrigger1")
	obj110 = ns.Object("GypsyEndTrigger2")
	obj111 = ns.Object("GypsyEndTrigger3")
	obj112 = ns.Object("GypsyEndTrigger4")
	obj113 = ns.Object("GypsyEndTrigger5")
	obj114 = ns.Object("GypsyEndTrigger6")
	obj115 = ns.Object("GypsyEndTrigger7")
	obj116 = ns.Object("GypsyEndTrigger8")
	obj117 = ns.Object("GypsyEndTrigger9")
	obj140 = ns.Object("W5Swordsman")
	obj141 = ns.Object("W5Swordsman2")
	wp142 = int(ns.Object("SwordsmanTrigger1"))
	wp143 = int(ns.Object("SwordsmanTrigger2"))
	obj139 = ns.Object("KeyGuy")
	obj136 = ns.Object("W5Drunk1")
	obj137 = ns.Object("W5Drunk2")
	obj138 = ns.Object("W5Drunk3")
	obj122 = ns.Object("SadVillagerReward")
	obj123 = ns.Object("FarmerReward1")
	obj124 = ns.Object("FarmerReward2")
	obj126 = ns.Object("W5Farmer")
	obj127 = ns.Object("W5Hound")
	obj128 = ns.Object("W5FarmerWife")
	obj129 = ns.Object("W5SadVillager")
	obj130 = ns.Object("W5Frog")
	obj131 = ns.Object("FrogTrigger")
	obj132 = ns.Object("FrogTrigger1")
	obj133 = ns.Object("FrogTrigger2")
	obj134 = ns.Object("FrogTrigger3")
	obj135 = ns.Object("FrogTrigger4")
	obj125 = ns.Object("FarmerStaff")
	obj88 = ns.Object("AutoSave1Trigger")
	obj89 = ns.Object("AutoSave2Trigger")
	obj90 = ns.Object("AutoSave3Trigger")
	obj91 = ns.Object("AutoSave4Trigger")
	obj144 = ns.Object("Corpse1")
	obj145 = ns.Object("Corpse2")
	obj146 = ns.Object("Corpse3")
	obj147 = ns.Object("Corpse4")
	obj148 = ns.Object("Corpse5")
	obj149 = ns.Object("Corpse6")
	obj150 = ns.Object("Corpse7")
	obj151 = ns.Object("Corpse8")
	obj152 = ns.Object("Corpse9")
	obj153 = ns.Object("Corpse10")
	obj154 = ns.Object("Corpse11")
	obj155 = ns.Object("Corpse12")
	obj156 = ns.Object("Corpse13")
	obj157 = ns.Object("Corpse14")
	obj158 = ns.Object("Corpse15")
	obj118 = ns.Object("JackLine5Trigger")
	obj159 = ns.Object("OgreCorpse1")
	obj160 = ns.Object("OgreCorpse2")
	obj161 = ns.Object("OgreCorpse3")
	obj175 = ns.Object("EntranceElevator")
	obj162 = ns.Object("Ogre1")
	obj163 = ns.Object("Ogre2")
	obj164 = ns.Object("Ogre3")
	obj165 = ns.Object("Ogre4")
	obj166 = ns.Object("Ogre5")
	obj167 = ns.Object("Ogre6")
	obj168 = ns.Object("Ogre7")
	obj169 = ns.Object("Ogre8")
	obj170 = ns.Object("Ogre9")
	obj170 = ns.Object("Ogre10")
	obj170 = ns.Object("Ogre11")
	obj170 = ns.Object("Ogre12")
	obj170 = ns.Object("Ogre13")
	obj193 = ns.Object("HouseDoor1")
	obj194 = ns.Object("HouseDoor1")
	obj195 = ns.Object("CryptDoor1")
	obj196 = ns.Object("CryptDoor2")
	obj176 = ns.Object("W5Villager1")
	obj177 = ns.Object("W5Villager2")
	obj178 = ns.Object("W5Villager3")
	obj179 = ns.Object("W5Villager4")
	obj180 = ns.Object("W5Villager5")
	obj181 = ns.Object("W5Villager6")
	obj182 = ns.Object("W5Villager7")
	obj183 = ns.Object("W5Villager8")
	obj184 = ns.Object("W5Villager9")
	wp187 = ns.Waypoint("Swordsman2WayPoint")
	wp186 = ns.Waypoint("SwordsmanWayPoint")
	wp142 = int(ns.Waypoint("SwordsmanTrigger1"))
	wp143 = int(ns.Waypoint("SwordsmanTrigger2"))
	wp207 = ns.Waypoint("HorvathWayPoint")
	wp208 = ns.Waypoint("HorvathWayPoint2")
	wp209 = ns.Waypoint("HorvathAppears")
	wp188 = ns.Waypoint("FarmerWayPoint")
	wp189 = ns.Waypoint("Farmer2WayPoint")
	wp204 = ns.Waypoint("FarmerWifeWayPoint")
	wp205 = ns.Waypoint("FarmerWife2WayPoint")
	wp190 = ns.Waypoint("HoundWayPoint")
	wp206 = ns.Waypoint("Hound2WayPoint")
	wp203 = ns.Waypoint("WoundedVillagerWayPoint")
	wp201 = ns.Waypoint("SadVillagerWayPoint")
	wp202 = ns.Waypoint("SadVillagerWayPoint2")
	wp191 = ns.Waypoint("FrogWayPoint1")
	wp192 = ns.Waypoint("FrogWayPoint2")
	wp210 = ns.Waypoint("Villager1Goal")
	wp211 = ns.Waypoint("Villager2Goal")
	wp212 = ns.Waypoint("Villager3Goal")
	wp213 = ns.Waypoint("Villager4Goal")
	wp214 = ns.Waypoint("Villager5Goal")
	wp216 = ns.Waypoint("GypsyWayPoint")
	wp217 = ns.Waypoint("GypsyWayPoint2")
	wp218 = ns.Waypoint("ShopKeeperWayPoint")
	wp219 = ns.Waypoint("ShopKeeperWayPoint2")
	wp220 = ns.Waypoint("BartenderWayPoint")
	wp221 = ns.Waypoint("BartenderWayPoint2")
	wp215 = ns.Waypoint("PlayerStartWayPoint")
	ns.SetDialog(obj136, ns.NORMAL, DrunkStart, DrunkEnd)
	ns.StoryPic(obj136, "Townsman3Pic")
	ns.SetDialog(obj137, ns.NORMAL, DrunkStart, DrunkEnd)
	ns.StoryPic(obj137, "Townsman2Pic")
	ns.SetDialog(obj138, ns.NORMAL, DrunkStart, DrunkEnd)
	ns.StoryPic(obj138, "Townsman1Pic")
	ns.SetDialog(obj126, ns.YESNO, FarmerStart, FarmerEnd)
	ns.StoryPic(obj126, "ThaviusPic")
	ns.SetDialog(obj127, ns.NORMAL, HoundStart, HoundEnd)
	ns.StoryPic(obj127, "WolfPic")
	ns.SetDialog(obj129, ns.YESNO, SadVillagerStart, SadVillagerEnd)
	ns.StoryPic(obj129, "Townsman3Pic")
	ns.SetDialog(obj130, ns.NORMAL, FrogStart, FrogEnd)
	ns.StoryPic(obj130, "FrogPic")
	ns.SetDialog(obj92, ns.NORMAL, CaptainStart, CaptainEnd)
	ns.StoryPic(obj92, "AirshipCaptainPic")
	ns.SetDialog(obj128, ns.NORMAL, FarmerWifeStart, FarmerWifeEnd)
	ns.StoryPic(obj128, "MaidenPic")
	ns.SetDialog(obj55, ns.NORMAL, HorvathStart, HorvathEnd)
	ns.StoryPic(obj55, "HorvathPic")
	ns.Frozen(obj130, true)
	DisableVillagers()
	KillCorpses()
	LockDoors()
	OwnVillagers()
	HoundEnd()
	PickUpObjects()
	EnableElevator()
	ns.FrameTimer(90, DisableElevator)
	ShopKeeperGuard()
	ns.StartupScreen(5)
}
func KeepOut() {
	r1 := ns.IsAttackedBy(ns.GetCaller(), ns.GetHost())
	if !r1 {
		goto LABEL1
	}
	ns.GoBackHome(ns.GetCaller())
LABEL1:
	return
}
func FrogReturned() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
	)
	if !ns.IsCaller(obj130) {
		goto LABEL1
	}
	v0 = ns.GetWaypointX(wp191)
	v1 = ns.GetWaypointY(wp191)
	v2 = ns.GetWaypointX(wp192)
	v3 = ns.GetWaypointY(wp192)
	ns.ObjectOff(obj131)
	ns.CreatureGuard(obj130, v0, v1, v2, v3, 0)
	gvar76 = gvar15
LABEL1:
	return
}
func FrogDeath() {
	gvar86 = 2
	gvar76 = gvar18
	ns.CancelDialog(obj130)
}
func AutoSave1() {
	ns.AutoSave()
	ns.ObjectOff(obj88)
}
func AutoSave2() {
	ns.AutoSave()
	ns.ObjectOff(obj89)
}
func AutoSave3() {
	ns.AutoSave()
	ns.FrameTimer(1, VillagePath)
	ns.ObjectOff(obj90)
}
func AutoSave4() {
	ns.AutoSave()
	ns.ObjectOff(obj91)
}
func SecretFound() {
	ns.PrintToAll("War05A.scr:Secret")
	ns.MoveWaypoint(wp33, ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger()))
	ns.AudioEvent(ns.SecretFound, wp33)
	ns.GiveXp(ns.GetHost(), 250)
}
func Hint() {
	ns.PrintToAll("Journal:War05Hint1")
	ns.JournalEntry(ns.GetHost(), "War05Hint1", 4)
}
func OgreAttackStart() {
	ns.ObjectGroupOff(gvar64)
	ns.Music(26, 100)
	UnfreezeAll()
	ns.Frozen(obj51, false)
	ns.Frozen(obj47, false)
	ns.Frozen(obj48, false)
	ns.Attack(obj51, obj47)
	ns.Attack(obj47, obj51)
	ns.Attack(obj48, obj51)
}
func DeadGuyArrive() {
	if !ns.IsCaller(obj51) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.Frozen(obj51, true)
	ns.Frozen(obj47, true)
	ns.Frozen(obj48, true)
	ns.CreatureIdle(obj51)
	ns.CreatureIdle(obj47)
	ns.CreatureIdle(obj48)
	ns.LookAtObject(obj47, obj51)
	ns.LookAtObject(obj51, obj48)
	ns.LookAtObject(obj48, obj51)
	ns.Chat(obj48, "Con05:OgreKingTalk06")
	ns.SecondTimer(2, OgresAttackDeadGuy)
LABEL1:
	return
}
func OgresAttackDeadGuy() {
	ns.Frozen(obj51, false)
	ns.Frozen(obj47, false)
	ns.Frozen(obj48, false)
	ns.Attack(obj51, obj48)
	ns.Attack(obj47, obj51)
	ns.Attack(obj48, obj51)
}
func DeadGuyDie() {
	ns.AggressionLevel(obj47, 0.83)
	ns.AggressionLevel(obj48, 0.83)
}
func LunchOgreReport() {
	if !(ns.CurrentHealth(obj51) > 0) {
		goto LABEL1
	}
	ns.Attack(ns.GetTrigger(), obj51)
LABEL1:
	return
}
func LunchOgre1Trigger() {
	if !ns.IsCaller(obj47) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ivar222 += 1
	ns.AggressionLevel(obj47, 0.83)
LABEL1:
	return
}
func LunchOgre2Trigger() {
	if !ns.IsCaller(obj48) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ivar222 += 1
	ns.AggressionLevel(obj48, 0.83)
LABEL1:
	return
}
func TownOgre2Die() {
	ns.CreatureGuard(obj126, ns.GetObjectX(obj126), ns.GetObjectY(obj126), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
}
func UnfreezeAll() {
	ns.Frozen(ns.GetHost(), false)
	ns.Frozen(obj176, false)
	ns.Frozen(obj177, false)
	ns.Frozen(obj178, false)
	ns.Frozen(obj179, false)
	ns.Frozen(obj180, false)
	ns.Frozen(obj181, false)
	ns.Frozen(obj182, false)
	ns.Frozen(obj183, false)
	ns.Frozen(obj184, false)
	ns.Frozen(obj49, false)
	ns.Frozen(obj50, false)
	ns.Frozen(obj126, false)
}
func GoAway() {
	r1 := ns.IsAttackedBy(ns.GetCaller(), ns.GetHost())
	if !r1 {
		goto LABEL1
	}
	if !(ns.CurrentHealth(ns.GetCaller()) > 0) {
		goto LABEL1
	}
	ns.GoBackHome(ns.GetCaller())
LABEL1:
	return
}
func StayAway() {
	r1 := ns.IsAttackedBy(ns.GetCaller(), ns.GetHost())
	if r1 {
		goto LABEL1
	}
	if !(ns.CurrentHealth(ns.GetCaller()) > 0) {
		goto LABEL1
	}
	ns.GoBackHome(ns.GetCaller())
LABEL1:
	return
}
func MakeAWish() {
	var v0 int
	v0 = ns.MaxHealth(ns.GetCaller())
	ns.RestoreHealth(ns.GetCaller(), v0-ns.CurrentHealth(ns.GetCaller()))
	ns.PrintToAll("GeneralPrint:WellSignRefresh")
	ns.AudioEvent(ns.RestoreHealthName, ns.Waypoint("WellWP"))
}
func CrapFile() {
}
func OnEvent(typ string) {
	switch typ {
	case "PlayerDeath":
		PlayerDeath()
	case "MapEntry":
		MapEntry()
	case "MapInitialize":
		MapInitialize()
	}
}
