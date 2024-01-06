package wiz05a

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
	wp38    ns.WaypointID
	wp39    ns.WaypointID
	obj40   ns.ObjectID
	wp41    ns.WaypointID
	wp42    ns.WaypointID
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
	obj56   ns.ObjectID
	obj57   ns.ObjectID
	obj58   ns.ObjectID
	obj59   ns.ObjectID
	obj60   ns.ObjectID
	obj61   ns.ObjectID
	obj62   ns.ObjectID
	obj63   ns.ObjectID
	obj64   ns.ObjectID
	gvar65  ns.ObjectGroupID
	gvar66  ns.ObjectGroupID
	gvar67  ns.ObjectGroupID
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
	flag88  bool
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
	gvar111 int
	obj112  ns.ObjectID
	gvar113 int
	gvar114 int
	obj115  ns.ObjectID
	obj116  ns.ObjectID
	obj117  ns.ObjectID
	wp118   int
	wp119   int
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
	obj142  ns.ObjectID
	obj143  ns.ObjectID
	obj144  ns.ObjectID
	obj145  ns.ObjectID
	obj146  ns.ObjectID
	gvar147 int
	gvar148 int
	gvar149 int
	gvar150 int
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
	gvar161 int
	wp162   ns.WaypointID
	wp163   ns.WaypointID
	wp164   ns.WaypointID
	wp165   ns.WaypointID
	wp166   ns.WaypointID
	wp167   ns.WaypointID
	wp168   ns.WaypointID
	gvar169 int
	gvar170 int
	gvar171 int
	gvar172 int
	gvar173 int
	gvar174 int
	wp175   ns.WaypointID
	wp176   ns.WaypointID
	wp177   ns.WaypointID
	wp178   ns.WaypointID
	wp179   ns.WaypointID
	wp180   ns.WaypointID
	wp181   ns.WaypointID
	wp182   ns.WaypointID
	wp183   ns.WaypointID
	wp184   ns.WaypointID
	wp185   ns.WaypointID
	wp186   ns.WaypointID
	wp187   ns.WaypointID
	wp188   ns.WaypointID
	wp189   ns.WaypointID
	wp190   ns.WaypointID
	wp191   ns.WaypointID
	wp192   ns.WaypointID
	wp193   ns.WaypointID
	wp194   ns.WaypointID
	wp195   ns.WaypointID
	ivar196 int
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
	flag88 = false
	ivar196 = 0
}
func OgreDeath() {
	ivar68 += 1
	if !(ivar68 >= 4) {
		goto LABEL1
	}
	ns.ObjectGroupOff(gvar67)
	SetTownDialog()
	ns.CreatureGuard(obj102, ns.GetWaypointX(wp164), ns.GetWaypointY(wp164), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	ns.Music(23, 100)
LABEL1:
	return
}
func SetTownDialog() {
	ns.SetDialog(obj152, ns.NORMAL, Villager1Start, Villager1End)
	ns.StoryPic(obj152, "MalePic8")
	ns.SetDialog(obj153, ns.NORMAL, Villager2Start, Villager2End)
	ns.StoryPic(obj153, "MalePic3")
	ns.SetDialog(obj154, ns.NORMAL, Villager3Start, Villager3End)
	ns.StoryPic(obj154, "MalePic2")
	ns.SetDialog(obj155, ns.NORMAL, Villager4Start, Villager4End)
	ns.StoryPic(obj155, "MalePic14")
	ns.SetDialog(obj156, ns.NORMAL, Villager5Start, Villager5End)
	ns.StoryPic(obj156, "MalePic9")
}
func CaptainGuard() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
	)
	v0 = ns.GetWaypointX(wp32)
	v1 = ns.GetWaypointY(wp32)
	v2 = ns.GetObjectX(ns.GetHost())
	v3 = ns.GetObjectY(ns.GetHost())
	ns.CreatureGuard(obj93, v0, v1, v2, v3, 0)
}
func EnterCaptainArea() {
	if !(ns.IsCaller(ns.GetHost()) && flag71 == false) {
		goto LABEL1
	}
	flag71 = true
	LetterBoxOn()
	ns.Frozen(ns.GetHost(), true)
	ns.Move(obj93, wp33)
	ns.FrameTimer(60, ExitCaptainArea)
LABEL1:
	return
}
func ExitCaptainArea() {
	ns.LookAtObject(obj93, ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj93)
	ns.StartDialog(obj93, ns.GetHost())
}
func ExitCaptainArea2() {
	LetterBoxOff()
	CaptainGuard()
	ns.FrameTimer(30, ExitCaptainArea3)
}
func ExitCaptainArea3() {
	ns.Frozen(ns.GetHost(), false)
}
func CheckForHorvath() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	if ns.GetQuestStatus("Wiz05B:GotHorvath") != 1 {
		goto LABEL1
	}
	ns.Frozen(ns.GetHost(), true)
	ns.CreatureIdle(ns.GetHost())
	ns.LookAtObject(ns.GetHost(), ns.ObjectID(wp39)) // FIXME
	ns.WideScreen(true)
	ns.Move(obj93, wp39)
LABEL1:
	return
}
func CaptainCongrat() {
	if !ns.IsCaller(obj93) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.SetDialog(obj93, ns.NEXT, CaptainCongratStart, CaptainCongratEnd)
	ns.Frozen(obj93, true)
	ns.CreatureIdle(obj93)
	ns.LookAtObject(obj93, ns.GetHost())
	ns.StartDialog(obj93, ns.GetHost())
LABEL1:
	return
}
func CaptainCongratStart() {
	ns.TellStory(ns.HumanMaleEatFood, "Wiz05B.scr:CaptainReturned")
}
func CaptainCongratEnd() {
	ns.CancelDialog(obj93)
	ns.StoryPic(obj57, "HorvathPic")
	ns.SetDialog(obj57, ns.NORMAL, HorvathCongratStart, HorvathCongratEnd)
	ns.Pickup(ns.GetHost(), obj58)
	ns.StartDialog(obj57, ns.GetHost())
}
func HorvathCongratStart() {
	ns.TellStory(ns.HumanMaleEatFood, "Wiz05B.scr:HorvathIdle")
}
func HorvathCongratEnd() {
	ns.CancelDialog(obj57)
	ns.GiveXp(ns.GetHost(), 1000)
	ns.WideScreen(false)
	ns.Blind()
	ns.Frozen(ns.GetHost(), false)
	ns.MakeEnemy(obj57)
	ns.FrameTimer(45, ChangeMaps)
}
func ChangeMaps() {
	ns.MoveObject(ns.GetHost(), 3955, 1051)
}
func KillCorpses() {
	ns.FrameTimer(30, KillCorpses2)
}
func SwordsmanStart() {
	var v0 int
	ns.Frozen(obj116, true)
	ns.CreatureIdle(obj116)
	ns.LookAtObject(obj116, ns.GetHost())
	v0 = gvar74
	if v0 == gvar12 {
		goto LABEL1
	}
	if v0 == gvar13 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.LookAtObject(obj116, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:SwordsmanGreeting")
	goto LABEL3
LABEL2:
	ns.LookAtObject(obj116, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:SwordsmanEnd")
	goto LABEL3
LABEL3:
	return
}
func SwordsmanEnd() {
	var v0 int
	ns.Frozen(obj116, false)
	ns.CreatureGuard(obj116, ns.GetWaypointX(wp162), ns.GetWaypointY(wp162), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 240)
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
	ns.LookAtDirection(obj116, ns.NE)
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
	if !ns.HasItem(ns.GetHost(), obj101) {
		goto LABEL8
	}
	gvar72 = gvar7
	ns.SetDialog(obj102, ns.NORMAL, FarmerStart, FarmerEnd)
	FarmerStart()
	goto LABEL7
	goto LABEL2
LABEL8:
	ns.CreatureIdle(obj102)
	ns.LookAtObject(obj102, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:FarmerGreeting")
	goto LABEL7
LABEL2:
	ns.CreatureIdle(obj102)
	ns.LookAtObject(obj102, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:FarmerGreeting")
	goto LABEL7
LABEL3:
	ns.CreatureIdle(obj102)
	ns.LookAtObject(obj102, ns.GetHost())
	if !ns.HasItem(ns.GetHost(), obj101) {
		goto LABEL9
	}
	v0 = ns.GetWaypointX(wp164)
	v1 = ns.GetWaypointY(wp164)
	gvar72 = gvar7
	FarmerStart()
	goto LABEL4
LABEL9:
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:FarmerWaiting")
	goto LABEL7
LABEL4:
	ns.CreatureIdle(obj102)
	ns.LookAtObject(obj102, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:FarmerReturned")
	goto LABEL7
LABEL5:
	ns.CreatureIdle(obj102)
	ns.LookAtObject(obj102, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:FarmerIdle")
	goto LABEL7
LABEL6:
	if !ns.HasItem(ns.GetHost(), obj101) {
		goto LABEL10
	}
	gvar72 = gvar7
	FarmerStart()
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
	v0 = ns.GetAnswer(obj102)
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
	ns.SetDialog(obj102, ns.NORMAL, FarmerStart, FarmerEnd)
	ns.PrintToAll("Con01a:NewJournalEntry")
	ns.JournalEntry(ns.GetHost(), "War05Quest8", 2)
	goto LABEL7
	goto LABEL2
LABEL8:
	if v0 != v2 {
		goto LABEL9
	}
	ns.SetDialog(obj102, ns.YESNO, FarmerStart, FarmerEnd)
	gvar72 = gvar5
	goto LABEL7
	goto LABEL2
LABEL9:
	if obj102 != v3 {
		goto LABEL2
	}
	gvar72 = gvar4
LABEL2:
	if v0 != v1 {
		goto LABEL10
	}
	gvar72 = gvar6
	ns.SetDialog(obj102, ns.NORMAL, FarmerStart, FarmerEnd)
	ns.PrintToAll("Con01a:NewJournalEntry")
	ns.JournalEntry(ns.GetHost(), "War05Quest8", 2)
	goto LABEL7
	goto LABEL3
LABEL10:
	if v0 != v2 {
		goto LABEL11
	}
	ns.SetDialog(obj102, ns.NORMAL, FarmerStart, FarmerEnd)
	gvar72 = gvar9
	goto LABEL7
	goto LABEL3
LABEL11:
	if obj102 != v3 {
		goto LABEL3
	}
	gvar72 = gvar4
LABEL3:
	goto LABEL7
LABEL4:
	ns.JournalEdit(ns.GetHost(), "War05Quest8", 4)
	ns.PrintToAll("War05A.scr:ObjectiveComplete")
	ns.GiveXp(ns.GetHost(), 250)
	ns.Pickup(ns.GetHost(), obj99)
	ns.Delete(obj101)
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
	ns.CreatureIdle(obj103)
	ns.LookAtObject(obj103, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:HoundGreeting")
}
func HoundEnd() {
	ns.SetRoamFlag(obj103, 2)
	ns.Wander(obj103)
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
	ns.LookAtObject(obj104, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:FarmerWifeGreeting")
	goto LABEL3
LABEL2:
	ns.LookAtObject(obj104, ns.GetHost())
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
	ns.LookAtObject(obj93, ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj93)
	ns.TellStory(ns.SwordsmanHurt, "Wiz05A.scr:CaptainGreeting")
	goto LABEL3
LABEL2:
	ns.LookAtObject(obj93, ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj93)
	ns.TellStory(ns.SwordsmanHurt, "Wiz05A.scr:CaptainWaiting")
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
	ns.PrintToAll("Con01a:NewJournalEntry")
	ns.JournalEntry(ns.GetHost(), "Wiz05Quest1", 2)
	ExitCaptainArea2()
	gvar77 = gvar25
	goto LABEL3
LABEL2:
	goto LABEL3
LABEL3:
	return
}
func BartenderStart() {
	ns.Frozen(obj95, true)
	ns.LookAtObject(obj95, ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj95)
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:BartenderGreeting")
}
func BartenderEnd() {
	ns.Frozen(obj95, false)
	ns.LookAtDirection(obj95, ns.SE)
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
	v8 = ns.GetWaypointX(wp164)
	v9 = ns.GetWaypointY(wp164)
	v0 = ns.GetWaypointX(wp180)
	v1 = ns.GetWaypointY(wp180)
	v2 = ns.GetWaypointX(wp178)
	v3 = ns.GetWaypointY(wp178)
	v6 = ns.GetWaypointX(wp179)
	v7 = ns.GetWaypointY(wp179)
	v4 = ns.GetObjectX(obj103)
	v5 = ns.GetObjectY(obj103)
	ns.Frozen(ns.GetHost(), true)
	LetterBoxOn()
	ns.AudioEvent(ns.CurePoisonCast, wp164)
	ns.AudioEvent(ns.CurePoisonEffect, wp164)
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
	v8 = ns.GetWaypointX(wp164)
	v9 = ns.GetWaypointY(wp164)
	v0 = ns.GetWaypointX(wp180)
	v1 = ns.GetWaypointY(wp180)
	v2 = ns.GetWaypointX(wp178)
	v3 = ns.GetWaypointY(wp178)
	v6 = ns.GetWaypointX(wp179)
	v7 = ns.GetWaypointY(wp179)
	v4 = ns.GetObjectX(obj103)
	v5 = ns.GetObjectY(obj103)
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
	v8 = ns.GetWaypointX(wp164)
	v9 = ns.GetWaypointY(wp164)
	v0 = ns.GetWaypointX(wp180)
	v1 = ns.GetWaypointY(wp180)
	v2 = ns.GetWaypointX(wp178)
	v3 = ns.GetWaypointY(wp178)
	v6 = ns.GetWaypointX(wp179)
	v7 = ns.GetWaypointY(wp179)
	v4 = ns.GetObjectX(obj103)
	v5 = ns.GetObjectY(obj103)
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
	v8 = ns.GetWaypointX(wp164)
	v9 = ns.GetWaypointY(wp164)
	v0 = ns.GetWaypointX(wp180)
	v1 = ns.GetWaypointY(wp180)
	v2 = ns.GetWaypointX(wp178)
	v3 = ns.GetWaypointY(wp178)
	v6 = ns.GetWaypointX(wp179)
	v7 = ns.GetWaypointY(wp179)
	v4 = ns.GetObjectX(obj103)
	v5 = ns.GetObjectY(obj103)
	ns.Effect(ns.BLUE_SPARKS, v4, v5, 0, 0)
	ns.MoveObject(obj103, v0, v1)
	ns.MoveObject(obj104, v4, v5)
	ns.CreatureGuard(obj104, v2, v3, v6, v7, 0)
	ns.CreatureGuard(obj102, v8, v9, v8, v9, 0)
}
func EndCurse5() {
	if !ns.IsCaller(obj104) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.LookAtObject(obj104, ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj104)
	ns.StartDialog(obj104, ns.GetHost())
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
	ns.LookAtObject(obj106, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:FrogGreeting")
	ns.Frozen(obj106, false)
	ns.ObjectOff(obj108)
	ns.ObjectOff(obj109)
	ns.ObjectOff(obj110)
	ns.ObjectOn(obj106)
	ns.CreatureFollow(obj106, ns.GetHost())
	goto LABEL5
LABEL2:
	ns.LookAtObject(obj106, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:FrogFollowing")
	goto LABEL5
LABEL3:
	ns.LookAtObject(obj106, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:FrogThanks")
	goto LABEL5
LABEL4:
	ns.LookAtObject(obj106, ns.GetHost())
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
	if !ns.IsCaller(obj106) {
		goto LABEL1
	}
	if flag88 {
		goto LABEL1
	}
	flag88 = true
	gvar75 = gvar22
	ns.Enchant(obj106, ns.ENCHANT_INVULNERABLE, 0)
	gvar76 = gvar16
	ns.SetDialog(obj105, ns.NORMAL, SadVillagerStart, SadVillagerEnd)
	ns.StartDialog(obj105, ns.GetHost())
	gvar86 = 1
LABEL1:
	return
}
func Villager1Start() {
	ns.Frozen(obj152, true)
	ns.CreatureIdle(obj152)
	ns.LookAtObject(obj152, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con05:Townsman9Talk01")
}
func Villager1End() {
	ns.Frozen(obj152, false)
	ns.Wander(obj152)
}
func Villager2Start() {
	ns.Frozen(obj153, true)
	ns.CreatureIdle(obj153)
	ns.LookAtObject(obj153, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con05:Townsman10Talk01")
}
func Villager2End() {
	ns.Frozen(obj153, false)
	ns.Wander(obj153)
}
func Villager3Start() {
	ns.Frozen(obj154, true)
	ns.CreatureIdle(obj154)
	ns.LookAtObject(obj154, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con05:Townsman11Talk01")
}
func Villager3End() {
	ns.Frozen(obj154, false)
	ns.Wander(obj154)
}
func Villager4Start() {
	ns.Frozen(obj155, true)
	ns.CreatureIdle(obj155)
	ns.LookAtObject(obj155, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con05:Townsman11Talk02")
}
func Villager4End() {
	ns.Frozen(obj155, false)
	ns.Wander(obj155)
}
func Villager5Start() {
	ns.Frozen(obj156, true)
	ns.CreatureIdle(obj156)
	ns.LookAtObject(obj156, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:VillagerGreeting")
}
func Villager5End() {
	ns.Frozen(obj156, false)
	ns.Wander(obj156)
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
	ns.LookAtObject(obj105, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:SadVillagerGreeting")
	goto LABEL7
LABEL2:
	ns.LookAtObject(obj105, ns.GetHost())
	if gvar86 != 1 {
		goto LABEL8
	}
	v0 = ns.GetWaypointX(wp175)
	v1 = ns.GetWaypointY(wp175)
	ns.AudioEvent(ns.CurePoisonCast, wp175)
	ns.AudioEvent(ns.CurePoisonEffect, wp175)
	gvar76 = gvar16
	ns.SetDialog(obj105, ns.NEXT, SadVillagerStart, SadVillagerEnd)
	ns.CreatureFollow(obj106, obj105)
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:SadVillagerReturned")
	goto LABEL3
LABEL8:
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:SadVillagerWaiting")
	goto LABEL7
LABEL3:
	ns.LookAtObject(obj105, ns.GetHost())
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
	v0 = 0
	v1 = 1
	v2 = 2
	v3 = 0
	v0 = ns.GetAnswer(obj105)
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
	if v0 != 1 {
		goto LABEL8
	}
	gvar76 = gvar15
	ns.SetDialog(obj105, ns.NORMAL, SadVillagerStart, SadVillagerEnd)
	ns.JournalEntry(ns.GetHost(), "War05Quest9", 2)
	ns.PrintToAll("Con01a:NewJournalEntry")
	goto LABEL7
LABEL8:
	if v0 != 2 {
		goto LABEL9
	}
	gvar76 = gvar14
	ns.SetDialog(obj105, ns.YESNO, SadVillagerStart, SadVillagerEnd)
	goto LABEL7
LABEL9:
	if v0 != 0 {
		goto LABEL2
	}
	gvar76 = gvar14
	ns.SetDialog(obj105, ns.YESNO, SadVillagerStart, SadVillagerEnd)
	goto LABEL7
LABEL2:
	ns.SetDialog(obj105, ns.NORMAL, SadVillagerStart, SadVillagerEnd)
	goto LABEL7
LABEL3:
	gvar76 = gvar17
	ns.SetDialog(obj105, ns.NORMAL, SadVillagerStart, SadVillagerEnd)
	ns.JournalEdit(ns.GetHost(), "War05Quest9", 4)
	ns.PrintToAll("War05A.scr:ObjectiveComplete")
	ns.GiveXp(ns.GetHost(), 250)
	ns.Pickup(ns.GetHost(), obj98)
	ns.FrameTimer(1, FrogThanksPlayer)
	ns.FrameTimer(1, SadVillagerGuard)
	goto LABEL7
LABEL4:
	ns.SetDialog(obj105, ns.NORMAL, SadVillagerStart, SadVillagerEnd)
	goto LABEL7
LABEL5:
	gvar76 = gvar19
	ns.Move(obj105, wp175)
	ns.SetCallback(obj105, 9, DummyBump)
	ns.JournalDelete(ns.GetHost(), "War05Quest9")
	ns.PrintToAll("War05A.scr:ObjectiveFailed")
	ns.SetDialog(obj105, ns.NORMAL, SadVillagerStart, SadVillagerEnd)
	goto LABEL7
LABEL6:
	ns.CancelDialog(obj105)
	goto LABEL7
LABEL7:
	return
}
func OwnVillagers() {
	ns.FrameTimer(1, OwnVillagers2)
}
func OwnVillagers2() {
	ns.Frozen(obj152, true)
	ns.Frozen(obj153, true)
	ns.Frozen(obj154, true)
	ns.Frozen(obj155, true)
	ns.Frozen(obj156, true)
	ns.Frozen(obj157, true)
	ns.Frozen(obj158, true)
	ns.Frozen(obj159, true)
	ns.Frozen(obj160, true)
	ns.Frozen(obj51, true)
	ns.Frozen(obj52, true)
	ns.Frozen(obj49, true)
	ns.Frozen(obj50, true)
	ns.Frozen(obj53, true)
}
func KillCorpses2() {
	ns.Damage(obj120, 0, 300, 5)
	ns.Damage(obj121, 0, 300, 5)
	ns.Damage(obj122, 0, 300, 5)
	ns.Damage(obj123, 0, 300, 5)
	ns.Damage(obj124, 0, 300, 5)
	ns.Damage(obj125, 0, 300, 5)
	ns.Damage(obj126, 0, 300, 5)
	ns.Damage(obj127, 0, 300, 5)
	ns.Damage(obj128, 0, 300, 5)
	ns.Damage(obj129, 0, 300, 5)
	ns.Damage(obj130, 0, 300, 5)
	ns.Damage(obj131, 0, 300, 5)
	ns.Damage(obj132, 0, 300, 5)
	ns.Damage(obj133, 0, 300, 5)
	ns.Damage(obj134, 0, 300, 5)
	ns.Damage(obj135, 0, 300, 0)
	ns.Damage(obj136, 0, 300, 0)
	ns.Damage(obj137, 0, 300, 0)
	ns.Damage(obj115, 0, 300, 0)
	ns.Damage(obj46, 0, 300, 0)
	ns.Damage(obj47, 0, 300, 0)
	ns.Damage(obj59, 0, 300, 0)
	ns.Damage(obj60, 0, 300, 0)
	ns.Damage(obj61, 0, 300, 0)
	ns.Damage(obj62, 0, 300, 0)
	ns.Damage(obj63, 0, 300, 0)
	ns.Damage(obj64, 0, 300, 0)
}
func VillagePath() {
	ns.Wander(obj152)
	ns.Wander(obj153)
	ns.Wander(obj154)
	ns.Wander(obj155)
	ns.Wander(obj156)
}
func EnableElevator() {
	ns.ObjectOn(obj151)
}
func DisableElevator() {
	ns.ObjectOff(obj151)
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
	ns.AggressionLevel(obj116, 0.83)
	ns.AggressionLevel(obj54, 0.83)
	ns.AggressionLevel(obj48, 0.83)
	ns.Attack(obj48, obj116)
	ns.CreatureHunt(obj48)
}
func SwordsmanTriggerB() {
	gvar74 = gvar12
	ns.SetDialog(obj116, ns.NORMAL, SwordsmanStart, SwordsmanEnd)
	ns.StoryPic(obj116, "Warrior8Pic")
	ns.FrameTimer(1, SwordsmanTriggerC)
}
func SwordsmanTriggerC() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
	)
	v0 = ns.GetWaypointX(wp162)
	v1 = ns.GetWaypointY(wp162)
	v2 = ns.GetWaypointX(wp163)
	v3 = ns.GetWaypointY(wp163)
	ns.CreatureGuard(obj116, v0, v1, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 240)
	ns.CreatureGuard(obj54, v2, v3, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 240)
}
func ShopKeeperGuard() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
	)
	v0 = ns.GetWaypointX(wp192)
	v1 = ns.GetWaypointY(wp192)
	v2 = ns.GetWaypointX(wp193)
	v3 = ns.GetWaypointY(wp193)
	ns.CreatureGuard(obj96, v0, v1, v2, v3, 0)
	ns.FrameTimer(1, GypsyGuard)
}
func BartenderGuard() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
	)
	v0 = ns.GetWaypointX(wp194)
	v1 = ns.GetWaypointY(wp194)
	v2 = ns.GetWaypointX(wp195)
	v3 = ns.GetWaypointY(wp195)
	ns.CreatureGuard(obj95, v0, v1, v2, v3, 0)
	ns.FrameTimer(1, SadVillagerGuard)
}
func GypsyGuard() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
	)
	v0 = ns.GetWaypointX(wp190)
	v1 = ns.GetWaypointY(wp190)
	v2 = ns.GetWaypointX(wp191)
	v3 = ns.GetWaypointY(wp191)
	ns.CreatureGuard(obj94, v0, v1, v2, v3, 0)
	ns.FrameTimer(1, BartenderGuard)
}
func SadVillagerGuard() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
	)
	v0 = ns.GetWaypointX(wp175)
	v1 = ns.GetWaypointY(wp175)
	v2 = ns.GetWaypointX(wp176)
	v3 = ns.GetWaypointY(wp176)
	ns.CreatureGuard(obj105, v0, v1, v2, v3, 0)
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
	ns.StartDialog(obj106, ns.GetHost())
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
	ns.Move(obj93, wp181)
	ns.FrameTimer(45, HorvathAppears3)
}
func HorvathAppears3() {
	ns.Frozen(obj93, true)
	ns.CreatureIdle(obj93)
	ns.CreatureIdle(ns.GetHost())
	ns.LookAtObject(obj93, ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj93)
	ns.StartDialog(obj93, ns.GetHost())
}
func HorvathAppears4() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
	)
	v0 = ns.GetWaypointX(wp183)
	v1 = ns.GetWaypointY(wp183)
	v2 = ns.GetWaypointX(wp181)
	v3 = ns.GetWaypointY(wp181)
	ns.Frozen(obj93, false)
	ns.CreatureGuard(obj93, v0, v1, v2, v3, 0)
	LetterBoxOff()
	ns.FrameTimer(45, HorvathAppears5)
}
func HorvathAppears5() {
	ns.Frozen(ns.GetHost(), false)
}
func HorvathDeath() {
	ns.PrintToAll("Wiz05C.scr:ObjectiveFailed")
	ns.FrameTimer(30, PlayerDeath)
}
func MapInitialize() {
	gvar66 = ns.ObjectGroup("ExitEnemies")
	wp30 = ns.Waypoint("WoundedRefugeeWayPoint")
	wp31 = ns.Waypoint("WoundedRefugeeWayPoint2")
	wp32 = ns.Waypoint("CaptainWayPoint")
	wp33 = ns.Waypoint("CaptainWayPoint2")
	wp34 = ns.Waypoint("DeadGuyWP")
	wp35 = ns.Waypoint("PlayerSounds")
	wp36 = ns.Waypoint("LunchOgre1WP")
	wp37 = ns.Waypoint("LunchOgre2WP")
	wp38 = ns.Waypoint("CaptainWaitingWP")
	wp39 = ns.Waypoint("CaptainCongratWP")
	obj43 = ns.Object("WoundedRefugee")
	obj44 = ns.Object("WoundedRefugee2")
	obj45 = ns.Object("Antidote")
	obj40 = ns.Object("WellWP")
	wp41 = ns.Waypoint("OgressYell1WP")
	wp42 = ns.Waypoint("OgressYell2WP")
	obj46 = ns.Object("Kill")
	obj47 = ns.Object("Kill2")
	obj48 = ns.Object("FirstOgre")
	obj49 = ns.Object("LunchOgre1")
	obj50 = ns.Object("LunchOgre2")
	obj51 = ns.Object("TownOgre1")
	obj52 = ns.Object("TownOgre2")
	obj53 = ns.Object("DeadGuy")
	obj54 = ns.Object("Swordsman2")
	obj55 = ns.Object("Archer1")
	obj56 = ns.Object("Archer2")
	obj58 = ns.Object("CapReward")
	obj59 = ns.Object("Bubbles1")
	obj60 = ns.Object("Bubbles2")
	obj61 = ns.Object("Bubbles3")
	obj62 = ns.Object("Bubbles4")
	obj63 = ns.Object("Bubbles5")
	obj64 = ns.Object("Bubbles6")
	gvar65 = ns.ObjectGroup("OgreAttackTriggers")
	obj93 = ns.Object("W5Captain")
	obj95 = ns.Object("W5Bartender")
	obj96 = ns.Object("W5ShopKeeper")
	obj94 = ns.Object("W5Gypsy")
	obj116 = ns.Object("W5Swordsman")
	obj117 = ns.Object("W5Swordsman2")
	wp118 = int(ns.Object("SwordsmanTrigger1"))
	wp119 = int(ns.Object("SwordsmanTrigger2"))
	obj115 = ns.Object("KeyGuy")
	obj112 = ns.Object("W5Drunk1")
	obj98 = ns.Object("SadVillagerReward")
	obj99 = ns.Object("FarmerReward1")
	obj100 = ns.Object("FarmerReward2")
	obj102 = ns.Object("W5Farmer")
	obj105 = ns.Object("W5SadVillager")
	obj106 = ns.Object("W5Frog")
	obj107 = ns.Object("FrogTrigger")
	obj101 = ns.Object("FarmerStaff")
	obj120 = ns.Object("Corpse1")
	obj121 = ns.Object("Corpse2")
	obj122 = ns.Object("Corpse3")
	obj123 = ns.Object("Corpse4")
	obj124 = ns.Object("Corpse5")
	obj125 = ns.Object("Corpse6")
	obj126 = ns.Object("Corpse7")
	obj127 = ns.Object("Corpse8")
	obj128 = ns.Object("Corpse9")
	obj129 = ns.Object("Corpse10")
	obj130 = ns.Object("Corpse11")
	obj131 = ns.Object("Corpse12")
	obj132 = ns.Object("Corpse13")
	obj133 = ns.Object("Corpse14")
	obj134 = ns.Object("Corpse15")
	obj97 = ns.Object("JackLine5Trigger")
	obj135 = ns.Object("OgreCorpse1")
	obj136 = ns.Object("OgreCorpse2")
	obj137 = ns.Object("OgreCorpse3")
	obj151 = ns.Object("EntranceElevator")
	obj138 = ns.Object("Ogre1")
	obj139 = ns.Object("Ogre2")
	obj140 = ns.Object("Ogre3")
	obj141 = ns.Object("Ogre4")
	obj142 = ns.Object("Ogre5")
	obj143 = ns.Object("Ogre6")
	obj144 = ns.Object("Ogre7")
	obj145 = ns.Object("Ogre8")
	obj146 = ns.Object("Ogre9")
	obj146 = ns.Object("Ogre10")
	obj146 = ns.Object("Ogre11")
	obj146 = ns.Object("Ogre12")
	obj146 = ns.Object("Ogre13")
	obj152 = ns.Object("W5Villager1")
	obj153 = ns.Object("W5Villager2")
	obj154 = ns.Object("W5Villager3")
	obj155 = ns.Object("W5Villager4")
	obj156 = ns.Object("W5Villager5")
	obj157 = ns.Object("W5Villager6")
	obj158 = ns.Object("W5Villager7")
	obj159 = ns.Object("W5Villager8")
	obj160 = ns.Object("W5Villager9")
	wp163 = ns.Waypoint("Swordsman2WayPoint")
	wp162 = ns.Waypoint("SwordsmanWayPoint")
	wp118 = int(ns.Waypoint("SwordsmanTrigger1"))
	wp119 = int(ns.Waypoint("SwordsmanTrigger2"))
	wp181 = ns.Waypoint("HorvathWayPoint")
	wp182 = ns.Waypoint("HorvathWayPoint2")
	wp183 = ns.Waypoint("HorvathAppears")
	wp164 = ns.Waypoint("FarmerWayPoint")
	wp165 = ns.Waypoint("Farmer2WayPoint")
	wp178 = ns.Waypoint("FarmerWifeWayPoint")
	wp179 = ns.Waypoint("FarmerWife2WayPoint")
	wp166 = ns.Waypoint("HoundWayPoint")
	wp180 = ns.Waypoint("Hound2WayPoint")
	wp177 = ns.Waypoint("WoundedVillagerWayPoint")
	wp175 = ns.Waypoint("SadVillagerWayPoint")
	wp176 = ns.Waypoint("SadVillagerWayPoint2")
	wp167 = ns.Waypoint("FrogWayPoint1")
	wp168 = ns.Waypoint("FrogWayPoint2")
	wp184 = ns.Waypoint("Villager1Goal")
	wp185 = ns.Waypoint("Villager2Goal")
	wp186 = ns.Waypoint("Villager3Goal")
	wp187 = ns.Waypoint("Villager4Goal")
	wp188 = ns.Waypoint("Villager5Goal")
	wp190 = ns.Waypoint("GypsyWayPoint")
	wp191 = ns.Waypoint("GypsyWayPoint2")
	wp192 = ns.Waypoint("ShopKeeperWayPoint")
	wp193 = ns.Waypoint("ShopKeeperWayPoint2")
	wp194 = ns.Waypoint("BartenderWayPoint")
	wp195 = ns.Waypoint("BartenderWayPoint2")
	gvar67 = ns.ObjectGroup("OgressYellTriggerGroup")
	wp189 = ns.Waypoint("PlayerStartWayPoint")
	ns.SetDialog(obj112, ns.NORMAL, DrunkStart, DrunkEnd)
	ns.StoryPic(obj112, "Townsman3Pic")
	ns.SetDialog(obj105, ns.YESNO, SadVillagerStart, SadVillagerEnd)
	ns.StoryPic(obj105, "GerardPic")
	ns.SetDialog(obj106, ns.NORMAL, FrogStart, FrogEnd)
	ns.StoryPic(obj106, "FrogPic")
	ns.SetDialog(obj93, ns.NORMAL, CaptainStart, CaptainEnd)
	ns.StoryPic(obj93, "AirshipCaptainPic")
	ns.Frozen(obj106, true)
	KillCorpses()
	OwnVillagers()
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
	if !ns.IsCaller(obj106) {
		goto LABEL1
	}
	v0 = ns.GetWaypointX(wp167)
	v1 = ns.GetWaypointY(wp167)
	v2 = ns.GetWaypointX(wp168)
	v3 = ns.GetWaypointY(wp168)
	ns.ObjectOff(obj107)
	ns.CreatureGuard(obj106, v0, v1, v2, v3, 0)
	gvar76 = gvar15
LABEL1:
	return
}
func FrogDeath() {
	gvar86 = 2
	gvar76 = gvar18
	ns.CancelDialog(obj106)
}
func AutoSave1() {
	ns.AutoSave()
	ns.ObjectOff(obj89)
}
func AutoSave2() {
	ns.AutoSave()
	ns.ObjectOff(obj90)
}
func AutoSave3() {
	ns.AutoSave()
	ns.FrameTimer(1, VillagePath)
	ns.ObjectOff(obj91)
}
func AutoSave4() {
	ns.AutoSave()
	ns.ObjectOff(obj92)
}
func SecretFound() {
	ns.PrintToAll("War05A.scr:Secret")
	ns.MoveWaypoint(wp35, ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger()))
	ns.AudioEvent(ns.SecretFound, wp35)
	ns.GiveXp(ns.GetHost(), 250)
}
func Hint() {
	ns.PrintToAll("Journal:War05Hint1")
	ns.JournalEntry(ns.GetHost(), "War05Hint1", 4)
}
func PlayerDeath() {
	ns.DeathScreen(5)
}
func MapEntry() {
	ns.Music(23, 100)
	ns.Frozen(ns.GetHost(), false)
	LetterBoxOff()
	if ns.GetQuestStatus("Wiz05B:GotHorvath") != 1 {
		goto LABEL1
	}
	ns.MoveObject(obj93, ns.GetWaypointX(wp38), ns.GetWaypointY(wp38))
	ns.CreatureGuard(obj93, ns.GetObjectX(obj93), ns.GetObjectY(obj93), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	obj57 = ns.Object("Wiz05B:Horvath")
	ns.StoryPic(obj57, "HorvathPic")
	ns.GroupDamage(gvar66, 0, 1000, 0)
	ns.SetCallback(obj57, 5, HorvathDeath)
LABEL1:
	flag87 = true
}
func OgressYell1() {
	ns.AudioEvent(ns.GruntRecognize, wp41)
}
func OgressYell2() {
	ns.AudioEvent(ns.GruntRecognize, wp42)
}
func CrapFile() {
}
func OgreAttackStart() {
	ns.ObjectGroupOff(gvar65)
	UnfreezeAll()
	ns.Music(26, 100)
	ns.Frozen(obj49, false)
	ns.Frozen(obj50, false)
	ns.Frozen(obj53, false)
	ns.AggressionLevel(obj49, 0.83)
	ns.AggressionLevel(obj50, 0.83)
	ns.Attack(obj53, obj49)
	ns.Attack(obj49, obj53)
	ns.Attack(obj50, obj53)
}
func DeadGuyArrive() {
	if !ns.IsCaller(obj53) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.Frozen(obj53, true)
	ns.Frozen(obj49, true)
	ns.Frozen(obj50, true)
	ns.CreatureIdle(obj53)
	ns.CreatureIdle(obj49)
	ns.CreatureIdle(obj50)
	ns.LookAtObject(obj49, obj53)
	ns.LookAtObject(obj53, obj50)
	ns.LookAtObject(obj50, obj53)
	ns.Chat(obj50, "Con05:OgreKingTalk06")
	ns.SecondTimer(2, OgresAttackDeadGuy)
LABEL1:
	return
}
func OgresAttackDeadGuy() {
	ns.Frozen(obj53, false)
	ns.Frozen(obj49, false)
	ns.Frozen(obj50, false)
	ns.Attack(obj53, obj50)
	ns.Attack(obj49, obj53)
	ns.Attack(obj50, obj53)
}
func DeadGuyDie() {
	ns.AggressionLevel(obj49, 0)
	ns.AggressionLevel(obj50, 0)
	ns.Move(obj49, wp36)
	ns.Move(obj50, wp37)
	UnfreezeAll()
}
func LunchOgreReport() {
	if !(ns.CurrentHealth(obj53) > 0) {
		goto LABEL1
	}
	ns.Attack(ns.GetTrigger(), obj53)
LABEL1:
	return
}
func LunchOgre1Trigger() {
	if !ns.IsCaller(obj49) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ivar196 += 1
	ns.AggressionLevel(obj49, 0.83)
LABEL1:
	return
}
func LunchOgre2Trigger() {
	if !ns.IsCaller(obj50) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ivar196 += 1
	ns.AggressionLevel(obj50, 0.83)
LABEL1:
	return
}
func UnfreezeAll() {
	ns.Frozen(obj152, false)
	ns.Frozen(obj153, false)
	ns.Frozen(obj154, false)
	ns.Frozen(obj155, false)
	ns.Frozen(obj156, false)
	ns.Frozen(obj157, false)
	ns.Frozen(obj158, false)
	ns.Frozen(obj159, false)
	ns.Frozen(obj160, false)
	ns.Frozen(obj51, false)
	ns.Frozen(obj52, false)
	ns.Frozen(obj102, false)
	ns.Wander(obj102)
	ns.CreatureFollow(obj51, obj102)
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
func MakeAWish() {
	var v0 int
	v0 = ns.MaxHealth(ns.GetCaller())
	ns.RestoreHealth(ns.GetCaller(), v0-ns.CurrentHealth(ns.GetCaller()))
	ns.PrintToAll("GeneralPrint:WellSignRefresh")
	ns.AudioEvent(ns.RestoreHealthName, ns.Waypoint("WellWP"))
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
