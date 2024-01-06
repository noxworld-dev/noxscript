package war01a

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	flag4   bool
	flag5   bool
	flag6   bool
	flag7   bool
	flag8   bool
	obj9    ns.ObjectID
	obj10   ns.ObjectID
	obj11   ns.ObjectID
	obj12   ns.ObjectID
	obj13   ns.ObjectID
	obj14   ns.ObjectID
	obj15   ns.ObjectID
	obj16   ns.ObjectID
	obj17   ns.ObjectID
	gvar18  ns.ObjectGroupID
	gvar19  ns.ObjectGroupID
	gvar20  ns.ObjectGroupID
	gvar21  ns.ObjectGroupID
	wp22    ns.WaypointID
	ivar23  int
	ivar24  int
	gvar25  int
	flag26  bool
	flag27  bool
	obj28   ns.ObjectID
	obj29   ns.ObjectID
	obj30   ns.ObjectID
	obj31   ns.ObjectID
	obj32   ns.ObjectID
	obj33   ns.ObjectID
	obj34   ns.ObjectID
	ivar35  int
	flag36  bool
	flag37  bool
	flag38  bool
	obj39   ns.ObjectID
	gvar40  ns.ObjectID
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
	obj56   ns.ObjectID
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
	gvar73  ns.ObjectID
	gvar74  ns.ObjectGroupID
	gvar75  ns.ObjectGroupID
	gvar76  ns.ObjectGroupID
	gvar77  ns.ObjectGroupID
	gvar78  ns.ObjectGroupID
	gvar79  ns.ObjectGroupID
	gvar80  ns.ObjectGroupID
	gvar81  ns.ObjectGroupID
	flag82  bool
	flag83  bool
	flag84  bool
	flag85  bool
	flag86  bool
	flag87  bool
	flag88  bool
	gvar89  int
	ivar90  int
	wp91    ns.WaypointID
	wp92    ns.WaypointID
	flag93  bool
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
	gvar118 ns.ObjectID
	obj119  ns.ObjectID
	obj120  ns.ObjectID
	gvar121 ns.ObjectGroupID
	gvar122 ns.ObjectGroupID
	gvar123 ns.WallGroupID
	wp124   ns.WaypointID
	gvar125 ns.WaypointGroupID
	gvar126 int
	gvar127 int
	gvar128 int
	gvar129 int
	gvar130 int
	gvar131 int
	gvar132 int
	gvar133 int
	ivar134 int
	flag135 bool
	obj136  ns.ObjectID
	obj137  ns.ObjectID
	obj138  ns.ObjectID
	obj139  ns.ObjectID
	obj140  ns.ObjectID
	obj141  ns.ObjectID
	wp142   ns.WaypointID
	wp143   ns.WaypointID
	wp144   ns.WaypointID
	wp145   ns.WaypointID
	wp146   ns.WaypointID
)

func init() {
	flag4 = true
	flag5 = false
	flag6 = false
	ivar23 = 0
	ivar24 = 0
	ivar35 = 0
	flag36 = true
	flag82 = false
	flag8 = false
	flag83 = true
	flag84 = true
	flag85 = false
	flag86 = false
	flag87 = false
	flag26 = false
	gvar89 = 1
	ivar90 = 1
	flag37 = false
	flag7 = false
	flag88 = false
	flag38 = false
	flag27 = false
	flag93 = true
	gvar25 = 1
	gvar126 = 0
	gvar127 = 1
	gvar128 = 2
	gvar129 = 3
	gvar130 = 4
	gvar131 = 5
	gvar132 = 6
	gvar133 = gvar126
	ivar134 = 0
	flag135 = false
}
func GHStart6() {
	ns.LookAtObject(obj16, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "War01A.scr:GearhartTalk06")
}
func GHEnd6() {
	ns.SetDialog(obj16, ns.NORMAL, GHStart6, GHEnd6)
}
func GHStart7() {
	ns.LookAtObject(obj16, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "War01A.scr:GearhartTalk07")
}
func GHEnd7() {
	ns.JournalEdit(ns.GetHost(), "ReturnToGearhart", 4)
	ns.PrintToAll("GeneralPrint:GainedItem")
	ns.GiveXp(ns.GetHost(), 500)
	ns.Pickup(ns.GetHost(), obj52)
	ns.CancelDialog(obj16)
	gvar25 = 0
	ns.SetDialog(obj17, ns.NORMAL, QMStart3, QMEnd3)
}
func StartFleePiece() {
	ns.ObjectGroupOff(gvar20)
	ns.SetDialog(obj16, ns.NORMAL, GHStart6, GHEnd6)
	gvar25 = 3
	flag7 = true
	ns.ObjectGroupOn(gvar18)
	ns.ObjectGroupOn(gvar19)
	ns.FrameTimer(1, StartBatAttack)
}
func StartBatAttack() {
	ns.Wander(obj14)
	ns.Wander(obj15)
	ns.Attack(obj10, obj14)
	ns.Attack(obj11, obj14)
	ns.Attack(obj12, obj15)
	ns.Attack(obj13, obj15)
	ns.SecondTimer(2, BatGroup1Peel)
}
func BatGroup1Die() {
	var v0 int
	ivar23 += 1
	v0 = ivar23
	if v0 == 1 {
		goto LABEL1
	}
	if v0 == 2 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.AggressionLevel(obj11, 0.83)
	ns.Attack(obj11, ns.GetHost())
	goto LABEL3
LABEL2:
	ns.FrameTimer(1, IdleMaiden1)
	goto LABEL3
LABEL3:
	return
}
func BatGroup2Die() {
	var v0 int
	ivar24 += 1
	v0 = ivar24
	if v0 == 1 {
		goto LABEL1
	}
	if v0 == 2 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.AggressionLevel(obj13, 0.83)
	ns.Attack(obj13, ns.GetHost())
	goto LABEL3
LABEL2:
	ns.FrameTimer(1, IdleMaiden2)
	goto LABEL3
LABEL3:
	return
}
func batToOrchard() {
	r1 := ns.IsAttackedBy(ns.GetHost(), ns.GetCaller())
	if !r1 {
		goto LABEL1
	}
	ns.GoBackHome(ns.GetCaller())
	obj9 = ns.GetCaller()
	ns.FrameTimer(90, BatReport)
LABEL1:
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL2
	}
	if !flag8 {
		goto LABEL2
	}
	ns.CreatureFollow(obj17, ns.GetHost())
LABEL2:
	return
}
func stopEscorting() {
	if !ns.IsCaller(obj17) {
		goto LABEL1
	}
	flag8 = true
	ns.CreatureGuard(obj17, 1291, 4108, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
LABEL1:
	return
}
func IdleMaiden1() {
	ns.CreatureIdle(obj14)
	flag5 = true
	ns.SetDialog(obj14, ns.NORMAL, Maiden1Start, Maiden1End)
	ns.CreatureGuard(obj14, ns.GetObjectX(obj14), ns.GetObjectY(obj14), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
}
func IdleMaiden2() {
	ns.CreatureIdle(obj15)
	flag6 = true
	ns.SetDialog(obj15, ns.NORMAL, Maiden2Start, Maiden2End)
	ns.CreatureGuard(obj15, ns.GetObjectX(obj15), ns.GetObjectY(obj15), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
}
func Maiden1See() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	if !flag4 {
		goto LABEL1
	}
	if !(ivar23 < 2) {
		goto LABEL1
	}
	ns.Chat(obj14, "War01A.scr:Maiden1Talk03")
	flag4 = false
	ns.SecondTimer(3, ResetMaiden)
LABEL1:
	return
}
func Maiden2See() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	if !flag4 {
		goto LABEL1
	}
	if !(ivar24 < 2) {
		goto LABEL1
	}
	ns.Chat(obj15, "War01A.scr:Maiden2Talk02")
	flag4 = false
	ns.SecondTimer(3, ResetMaiden)
LABEL1:
	return
}
func Maiden1Start() {
	ns.TellStory(ns.HumanMaleEatFood, "War01A.scr:Maiden1Talk04")
}
func Maiden1End() {
	ns.SetDialog(obj14, ns.NORMAL, Maiden1Start, Maiden1End)
}
func Maiden2Start() {
	ns.TellStory(ns.HumanMaleEatFood, "War01A.scr:Maiden2Talk03")
}
func Maiden2End() {
	ns.SetDialog(obj15, ns.NORMAL, Maiden2Start, Maiden2End)
}
func ResetMaiden() {
	flag4 = true
}
func Maiden1Contact() {
	if !flag5 {
		goto LABEL1
	}
	if !ns.IsAttackedBy(obj14, ns.GetCaller()) {
		goto LABEL1
	}
	ns.Wander(obj14)
	flag5 = false
	ns.CancelDialog(obj14)
	ns.FrameTimer(90, IdleMaiden1)
LABEL1:
	return
}
func Maiden2Contact() {
	if !flag6 {
		goto LABEL1
	}
	if !ns.IsAttackedBy(obj15, ns.GetCaller()) {
		goto LABEL1
	}
	ns.Wander(obj15)
	flag6 = false
	ns.CancelDialog(obj15)
	ns.FrameTimer(90, IdleMaiden2)
LABEL1:
	return
}
func BatGroup1Peel() {
	ns.Attack(obj10, ns.GetHost())
	ns.SecondTimer(2, BatGroup2Peel)
}
func BatGroup2Peel() {
	ns.Attack(obj12, ns.GetHost())
}
func CaveEntranceCheck() {
	ns.AggressionLevel(ns.GetCaller(), 0.16)
	ns.GoBackHome(ns.GetCaller())
}
func BatReport() {
	ns.AggressionLevel(obj9, 0.83)
}
func UrchinDead() {
	ns.MoveWaypoint(wp22, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.FlagCapture, wp22)
	ns.JournalEdit(ns.GetHost(), "OrchardQuest", 4)
	ns.JournalEntry(ns.GetHost(), "ReturnToGearhart", 2)
	ns.PrintToAll("Con01a:NewJournalEntry")
	ns.SetDialog(obj16, ns.NORMAL, GHStart7, GHEnd7)
	gvar25 = 4
	ns.ObjectGroupOff(gvar21)
	ns.Music(16, 100)
}
func ShortDelay() {
	ns.ObjectOff(obj32)
	ns.MoveObject(obj17, 2288, 2840)
	ns.SetQuestStatus(1, "PlayerInSewers")
	flag27 = true
	ns.FrameTimer(20, ElevatorExplode)
}
func ElevatorExplode() {
	var v0 int
	v0 = ivar35
	if v0 == 3 {
		goto LABEL1
	}
	goto LABEL2
LABEL1:
	ns.MoveWaypoint(wp22, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.MechGolemDie, wp22)
	ns.FrameTimer(30, QMTalkStart)
	goto LABEL3
LABEL2:
	if ivar35 != 0 {
		goto LABEL4
	}
	ns.Frozen(ns.GetHost(), true)
	ns.WideScreen(true)
LABEL4:
	ns.MoveWaypoint(wp22, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.HammerMissing, wp22)
	ns.AudioEvent(ns.WallDestroyed, wp22)
	ns.Effect(ns.JIGGLE, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 10, 0)
	ns.Effect(ns.SMOKE_BLAST, ns.GetObjectX(obj28), ns.GetObjectY(obj28), 0, 0)
	ns.Effect(ns.SMOKE_BLAST, ns.GetObjectX(obj29), ns.GetObjectY(obj29), 0, 0)
	ns.ObjectOff(obj28)
	ns.ObjectOff(obj29)
	ns.ObjectOff(obj32)
	ivar35 += 1
	if ivar35 != 1 {
		goto LABEL5
	}
	ns.MoveObject(obj30, 4507, 4426)
LABEL5:
	if ivar35 != 2 {
		goto LABEL6
	}
	ns.MoveObject(obj31, 4507, 4426)
LABEL6:
	ns.FrameTimer(10, ElevatorExplode)
	goto LABEL3
LABEL3:
	return
}
func QMTalkStart() {
	ns.SetDialog(obj17, ns.NORMAL, QMBlastStart, QMBlastEnd)
	ns.StartDialog(obj17, ns.GetHost())
}
func QMBlastStart() {
	ns.TellStory(ns.HumanMaleEatFood, "War02:QuarterMaster11")
}
func QMBlastEnd() {
	ns.CancelDialog(obj17)
	ns.Frozen(ns.GetHost(), false)
	ns.FrameTimer(2, BackToNormal)
}
func BackToNormal() {
	ns.WideScreen(false)
}
func TurnOffLights() {
	ns.LookAtObject(obj16, ns.GetHost())
	ns.ObjectGroupToggle(gvar75)
	ns.ObjectGroupToggle(gvar76)
	ns.ObjectGroupToggle(gvar77)
	ns.ObjectGroupToggle(gvar78)
	ns.ObjectToggle(obj53)
	ns.ObjectToggle(obj54)
	ns.ObjectToggle(obj55)
	if !flag84 {
		goto LABEL1
	}
	flag84 = false
	ns.MoveWaypoint(wp22, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.Gear2, wp22)
	ns.SetDialog(obj41, ns.NORMAL, DanStart2, DanEnd2)
	ns.SetDialog(obj42, ns.NORMAL, JesseStart2, JesseEnd2)
	ns.SetDialog(obj44, ns.NORMAL, EricStart2, EricEnd2)
	ns.SetDialog(obj16, ns.NORMAL, GHStart8, GHEnd8)
	ns.SetDialog(obj45, ns.NORMAL, JacobStart2, JacobEnd2)
	if ns.GetQuestStatus("EnteredGauntlet") == 1 {
		goto LABEL2
	}
	if !ns.IsVisibleTo(obj16, ns.GetHost()) {
		goto LABEL2
	}
	ns.FrameTimer(2, TalkToGearhart)
LABEL2:
	goto LABEL3
LABEL1:
	flag84 = true
	ns.MoveWaypoint(wp22, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.Gear3, wp22)
	ns.SetDialog(obj41, ns.NORMAL, DanStart1, DanEnd1)
	ns.SetDialog(obj42, ns.NORMAL, JesseStart1, JesseEnd1)
	ns.SetDialog(obj44, ns.NORMAL, EricStart1, EricEnd1)
	ns.SetDialog(obj45, ns.NORMAL, JacobStart1, JacobEnd1)
LABEL3:
	return
}
func TurnOnLights() {
	ns.ObjectGroupOn(gvar75)
	ns.ObjectGroupOn(gvar76)
	ns.ObjectGroupOn(gvar77)
	ns.ObjectOn(obj53)
	ns.ObjectOn(obj54)
	ns.ObjectOn(obj55)
	ns.SetDialog(obj41, ns.NORMAL, DanStart1, DanEnd1)
	ns.SetDialog(obj42, ns.NORMAL, JesseStart1, JesseEnd1)
	ns.SetDialog(obj44, ns.NORMAL, EricStart1, EricEnd1)
	ns.SetDialog(obj45, ns.NORMAL, JacobStart1, JacobEnd1)
}
func MoveToFacade() {
	ns.MoveObject(ns.GetHost(), 3473, 1234)
	ns.Frozen(obj16, false)
	ns.Frozen(ns.GetHost(), false)
	ns.MoveObject(obj16, 2215, 3389)
	ns.CreatureGuard(obj16, 2215, 3389, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	ns.UnBlind()
	ns.SetDialog(obj50, ns.NORMAL, HorStart1, HorEnd1)
	ns.StartupScreen(3)
}
func DanStart1() {
	ns.DestroyEveryChat()
	ns.Frozen(obj41, true)
	ns.CreatureIdle(obj41)
	ns.LookAtObject(obj41, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "War01A.scr:Townsman2Talk01")
}
func DanEnd1() {
	ns.Frozen(obj41, false)
	ns.Wander(obj41)
	ns.SetDialog(obj41, ns.NORMAL, DanStart1, DanEnd1)
}
func DanStart2() {
	ns.DestroyEveryChat()
	ns.Frozen(obj41, true)
	ns.CreatureIdle(obj41)
	ns.LookAtObject(obj41, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "War01A.scr:Townsman2Talk02")
}
func DanEnd2() {
	ns.Frozen(obj41, false)
	ns.Wander(obj41)
	ns.SetDialog(obj41, ns.NORMAL, DanStart2, DanEnd2)
}
func JesseStart1() {
	ns.DestroyEveryChat()
	ns.Frozen(obj42, true)
	ns.CreatureIdle(obj42)
	ns.LookAtObject(obj42, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "War01A.scr:Townsman1Talk01")
}
func JesseEnd1() {
	ns.Frozen(obj42, false)
	ns.Wander(obj42)
	ns.SetDialog(obj42, ns.NORMAL, JesseStart1, JesseEnd1)
}
func JesseStart2() {
	ns.DestroyEveryChat()
	ns.Frozen(obj42, true)
	ns.CreatureIdle(obj42)
	ns.LookAtObject(obj42, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "War01A.scr:Townsman1Talk02")
}
func JesseEnd2() {
	ns.Frozen(obj42, false)
	ns.Wander(obj42)
	ns.SetDialog(obj42, ns.NORMAL, JesseStart2, JesseEnd2)
}
func EricStart1() {
	ns.DestroyEveryChat()
	ns.Frozen(obj44, true)
	ns.CreatureIdle(obj44)
	ns.LookAtObject(obj44, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "War01A.scr:Townsman3Talk01")
}
func EricEnd1() {
	ns.Frozen(obj44, false)
	ns.Wander(obj44)
	ns.SetDialog(obj44, ns.NORMAL, EricStart1, EricEnd1)
}
func EricStart2() {
	ns.DestroyEveryChat()
	ns.Frozen(obj44, true)
	ns.CreatureIdle(obj44)
	ns.LookAtObject(obj44, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "War01A.scr:Townsman3Talk02")
}
func EricEnd2() {
	ns.Frozen(obj44, false)
	ns.Wander(obj44)
	ns.SetDialog(obj44, ns.NORMAL, EricStart2, EricEnd2)
}
func TylerStart1() {
	ns.DestroyEveryChat()
	ns.Frozen(obj46, true)
	ns.CreatureIdle(obj46)
	ns.LookAtObject(obj46, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "War01A.scr:Townsman4Talk01")
}
func TylerEnd1() {
	ns.Frozen(obj46, false)
	ns.LookAtObject(obj46, obj43)
	ns.CreatureGuard(obj46, ns.GetObjectX(obj46), ns.GetObjectY(obj46), ns.GetObjectX(obj43), ns.GetObjectY(obj43), 0)
	ns.SetDialog(obj46, ns.NORMAL, TylerStart1, TylerEnd1)
}
func JacobStart1() {
	ns.DestroyEveryChat()
	ns.Frozen(obj45, true)
	ns.CreatureIdle(obj45)
	ns.LookAtObject(obj45, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "War01A.scr:Townsman5Talk01")
}
func JacobEnd1() {
	ns.Frozen(obj45, false)
	ns.Wander(obj45)
	ns.SetDialog(obj45, ns.NORMAL, JacobStart1, JacobEnd1)
}
func JacobStart2() {
	ns.DestroyEveryChat()
	ns.Frozen(obj45, true)
	ns.CreatureIdle(obj45)
	ns.LookAtObject(obj45, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "War01A.scr:Townsman5Talk02")
}
func JacobEnd2() {
	ns.Frozen(obj45, false)
	ns.Wander(obj45)
	ns.SetDialog(obj45, ns.NORMAL, JacobStart2, JacobEnd2)
}
func BingStart1() {
	var (
		v0 int
		v1 int
	)
	ns.DestroyEveryChat()
	v0 = ns.Random(1, 6)
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
	goto LABEL7
LABEL1:
	ns.TellStory(ns.HumanMaleEatFood, "War02A:BartenderTalk1aStart")
	goto LABEL7
LABEL2:
	ns.TellStory(ns.HumanMaleEatFood, "War02A:BartenderTalk1bStart")
	goto LABEL7
LABEL3:
	ns.TellStory(ns.HumanMaleEatFood, "War02A:BartenderTalk1cStart")
	goto LABEL7
LABEL4:
	ns.TellStory(ns.HumanMaleEatFood, "War02:BartenderTalk1dStart")
	goto LABEL7
LABEL5:
	ns.TellStory(ns.HumanMaleEatFood, "War02A:BartenderTalk1eStart")
	goto LABEL7
LABEL6:
	ns.TellStory(ns.HumanMaleEatFood, "War02A:BartenderTalk1fStart")
	goto LABEL7
LABEL7:
	return
}
func BingEnd1() {
	ns.SetDialog(obj43, ns.NORMAL, BingStart1, BingEnd1)
}
func BingStart2() {
	ns.TellStory(ns.HumanMaleEatFood, "War02B.scr:BarTalk1")
}
func BingEnd2() {
	ns.SetDialog(obj43, ns.NORMAL, BingStart2, BingEnd2)
}
func LydiaStart1() {
	ns.DestroyEveryChat()
	ns.LookAtObject(obj49, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "War01A.scr:LydiaTalk02")
}
func LydiaEnd1() {
	ns.SetDialog(obj49, ns.NORMAL, LydiaStart1, LydiaEnd1)
}
func LydiaStart2() {
	ns.DestroyEveryChat()
	ns.LookAtObject(obj49, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "War01A.scr:LydiaTalk01")
}
func LydiaEnd2() {
	ns.SetDialog(obj49, ns.NORMAL, LydiaStart3, LydiaEnd3)
}
func LydiaStart3() {
	var (
		v0 int
		v1 int
	)
	v0 = ns.Random(1, 2)
	v1 = v0
	if v1 == 1 {
		goto LABEL1
	}
	if v1 == 2 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.TellStory(ns.HumanMaleEatFood, "War01A.scr:LydiaTalk01")
	goto LABEL3
LABEL2:
	ns.TellStory(ns.HumanMaleEatFood, "War01A.scr:LydiaTalk03")
	goto LABEL3
LABEL3:
	return
}
func LydiaEnd3() {
	ns.SetDialog(obj49, ns.NORMAL, LydiaStart3, LydiaEnd3)
}
func MelStart1() {
	var v0 int
	ns.DestroyEveryChat()
	ns.Frozen(obj47, true)
	ns.CreatureIdle(obj47)
	ns.LookAtObject(obj47, ns.GetHost())
	ivar90 = ns.Random(1, 2)
	v0 = ivar90
	if v0 == 1 {
		goto LABEL1
	}
	if v0 == 2 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.TellStory(ns.HumanMaleEatFood, "War01A.scr:Maiden1Talk01")
	goto LABEL3
LABEL2:
	ns.TellStory(ns.HumanMaleEatFood, "War01A.scr:Maiden1Talk02")
	goto LABEL3
LABEL3:
	return
}
func MelEnd1() {
	ns.Frozen(obj47, false)
	ns.LookAtObject(obj47, obj43)
	ns.CreatureGuard(obj47, ns.GetObjectX(obj47), ns.GetObjectY(obj47), ns.GetObjectX(obj43), ns.GetObjectY(obj43), 0)
	ns.SetDialog(obj47, ns.NORMAL, MelStart1, MelEnd1)
}
func EvelynStart1() {
	ns.DestroyEveryChat()
	ns.Frozen(obj48, true)
	ns.CreatureIdle(obj48)
	ns.LookAtObject(obj48, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "War01A.scr:Maiden2Talk01")
}
func EvelynEnd1() {
	ns.Frozen(obj48, false)
	ns.Wander(obj48)
	ns.SetDialog(obj48, ns.NORMAL, EvelynStart1, EvelynEnd1)
}
func QMProd() {
	var v0 float32
	v0 = ns.Distance(ns.GetObjectX(obj17), ns.GetObjectY(obj17), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	if !(v0 > 300) {
		goto LABEL1
	}
	if gvar89 != 1 {
		goto LABEL1
	}
	ns.CreatureFollow(obj17, ns.GetHost())
	flag8 = true
	gvar89 = 2
LABEL1:
	if !(v0 < 150) {
		goto LABEL2
	}
	if gvar89 != 2 {
		goto LABEL2
	}
	if !(ns.IsTalking() == false && ns.IsTrading() == false) {
		goto LABEL2
	}
	if !ns.IsVisibleTo(obj17, ns.GetHost()) {
		goto LABEL2
	}
	gvar89 = 1
	ns.StartDialog(obj17, ns.GetHost())
LABEL2:
	if flag26 {
		goto LABEL3
	}
	ns.SecondTimer(3, QMProd)
LABEL3:
	return
}
func QMGearTrigger() {
	if !ns.IsCaller(obj17) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.Wander(obj17)
	ns.CreatureGuard(obj17, 2280, 3101, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
LABEL1:
	return
}
func QMStart1() {
	ns.TellStory(ns.HumanMaleEatFood, "War01A.scr:QuaterMasterTalk01")
	ns.JournalEdit(ns.GetHost(), "FirstQuest", 4)
}
func QMEnd1() {
	if flag85 {
		goto LABEL1
	}
	ns.MoveObject(obj16, 1720, 3432)
	ns.CreatureFollow(obj16, ns.GetHost())
	flag37 = true
	ns.SetDialog(obj16, ns.YESNO, GHStart1, GHEnd1)
LABEL1:
	ns.SetDialog(obj17, ns.NORMAL, QMStart2, QMEnd2)
	if !flag88 {
		goto LABEL2
	}
	ns.JournalEdit(ns.GetHost(), "FirstQuest", 4)
LABEL2:
	ns.JournalEntry(ns.GetHost(), "SponsorQuest", 2)
	ns.PrintToAll("Con01a:NewJournalEntry")
}
func QMStart2() {
	ns.TellStory(ns.HumanMaleEatFood, "War01A.scr:QuarterMasterTalk02")
}
func QMEnd2() {
	ns.SetDialog(obj17, ns.NORMAL, QMStart2, QMEnd2)
}
func QMStart3() {
	ns.TellStory(ns.HumanMaleEatFood, "War01A.scr:QuarterMasterTalk03")
}
func QMEnd3() {
	ns.CancelDialog(obj17)
	ns.GiveXp(ns.GetHost(), 1000)
	ns.MoveWaypoint(wp22, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.FlagCapture, wp22)
	ns.Blind()
	ns.SetQuestStatus(1, "EnteredGauntlet")
	flag38 = true
	ns.FrameTimer(45, DropLetter)
}
func QMStart4() {
	ns.TellStory(ns.HumanMaleEatFood, "War02a:QuarterMasterTalk03")
}
func QMEnd4() {
	ns.JournalEntry(ns.GetHost(), "SaveGearhart", 2)
	ns.PrintToAll("Con01a:NewJournalEntry")
	ns.CreatureFollow(obj17, ns.GetHost())
	flag8 = true
	ns.SetDialog(obj17, ns.NORMAL, QMStart5, QMEnd5)
	ns.StartDialog(obj17, ns.GetHost())
}
func QMStart5() {
	ns.TellStory(ns.HumanMaleEatFood, "War02a:QuarterMasterTalk04")
}
func QMEnd5() {
	ns.SetDialog(obj17, ns.NORMAL, QMStart5a, QMEnd5a)
}
func QMStart5a() {
	ns.TellStory(ns.HumanMaleEatFood, "War02a:QuarterMasterTalk05")
}
func QMEnd5a() {
	ns.SetDialog(obj17, ns.NORMAL, QMStart5a, QMEnd5a)
}
func QMStart6() {
	ns.TellStory(ns.HumanMaleEatFood, "War02:QuarterMaster10")
}
func QMEnd6() {
	ns.CancelDialog(obj17)
	ns.CreatureGuard(obj17, 2280, 3101, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	if flag26 {
		goto LABEL1
	}
	ns.Walk(obj17, 2276, 3120)
LABEL1:
	return
}
func QMStart7() {
	ns.SetDialog(obj39, ns.NORMAL, KirikStart2, KirikEnd2)
	ns.TellStory(ns.HumanMaleEatFood, "War02:QuarterMaster12")
}
func QMEnd7() {
	if gvar40 == 0 {
		goto LABEL1
	}
	ns.Delete(gvar40)
LABEL1:
	if gvar73 == 0 {
		goto LABEL2
	}
	ns.Delete(gvar73)
LABEL2:
	ns.ObjectOff(obj58)
	ns.ObjectOn(obj72)
	ns.GiveXp(ns.GetHost(), 1500)
	ns.MoveWaypoint(wp22, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.FlagCapture, wp22)
	ns.JournalEdit(ns.GetHost(), "SaveGearhart", 4)
	ns.CancelDialog(obj17)
	ns.SetDialog(obj16, ns.NORMAL, GHStart9, GHEnd9)
	ns.StartDialog(obj16, ns.GetHost())
}
func TalkToGearhart() {
	ns.StartDialog(obj16, ns.GetHost())
}
func GHStart1() {
	ns.LookAtObject(obj16, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "War01A.scr:GearhartTalk01")
}
func GHEnd1() {
	if ns.GetAnswer(obj16) != 1 {
		goto LABEL1
	}
	ns.SetDialog(obj16, ns.NORMAL, GHStart2, GHEnd2)
	GHStart2()
	goto LABEL2
LABEL1:
	ns.SetDialog(obj16, ns.NORMAL, GHStart3, GHEnd3)
	GHStart3()
LABEL2:
	flag85 = true
}
func GHStart2() {
	ns.LookAtObject(obj16, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "War01A.scr:GearhartTalk02")
}
func GHEnd2() {
	ns.Frozen(obj16, false)
	flag86 = true
	gvar25 = 2
	ns.CancelDialog(obj16)
	ns.Pickup(ns.GetHost(), obj51)
	ns.JournalEdit(ns.GetHost(), "SponsorQuest", 4)
	ns.JournalEntry(ns.GetHost(), "OrchardQuest", 2)
	ns.PrintToAll("GeneralPrint:GainedKey")
	ns.PrintToAll("Con01a:NewJournalEntry")
	if flag87 {
		goto LABEL1
	}
	ns.ObjectOn(obj71)
	flag37 = false
	ns.Move(obj16, wp91)
LABEL1:
	return
}
func GHStart3() {
	ns.Frozen(obj16, true)
	ns.LookAtObject(obj16, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "War01A.scr:GearhartTalk03")
}
func GHEnd3() {
	ns.Frozen(obj16, false)
	ns.SetDialog(obj16, ns.YESNO, GHStart5, GHEnd5)
	gvar25 = 1
	if flag87 {
		goto LABEL1
	}
	ns.ObjectOn(obj71)
	flag37 = false
	ns.Move(obj16, wp91)
LABEL1:
	return
}
func GHStart4() {
	ns.LookAtObject(obj16, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "War01A.scr:GearhartTalk05")
}
func GHEnd4() {
	ns.SetDialog(obj16, ns.NORMAL, GHStart4, GHEnd4)
}
func GHStart5() {
	ns.Frozen(obj16, true)
	ns.CreatureIdle(obj16)
	ns.LookAtObject(obj16, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "War01A.scr:GearhartTalk04")
}
func GHEnd5() {
	if ns.GetAnswer(obj16) != 1 {
		goto LABEL1
	}
	ns.SetDialog(obj16, ns.NORMAL, GHStart2, GHEnd2)
	GHStart2()
	goto LABEL2
LABEL1:
	ns.SetDialog(obj16, ns.NORMAL, GHStart3, GHEnd3)
	GHStart3()
LABEL2:
	flag85 = true
}
func GHStart8() {
	ns.LookAtObject(obj16, ns.GetHost())
	ns.Chat(obj16, "War01A.scr:GearhartTalk08")
	ns.FrameTimer(60, GHEnd8)
}
func GHEnd8() {
	var v0 int
	v0 = gvar25
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
	goto LABEL6
LABEL1:
	ns.CancelDialog(obj16)
	goto LABEL6
LABEL2:
	ns.SetDialog(obj16, ns.YESNO, GHStart5, GHEnd5)
	goto LABEL6
LABEL3:
	ns.SetDialog(obj16, ns.NORMAL, GHStart4, GHEnd4)
	goto LABEL6
LABEL4:
	ns.SetDialog(obj16, ns.NORMAL, GHStart6, GHEnd6)
	goto LABEL6
LABEL5:
	ns.SetDialog(obj16, ns.NORMAL, GHStart7, GHEnd7)
	goto LABEL6
LABEL6:
	return
}
func GHStart9() {
	ns.LookAtObject(ns.GetHost(), obj16)
	ns.TellStory(ns.HumanMaleEatFood, "War02a:GearhartTalk07")
}
func GHEnd9() {
	ns.CancelDialog(obj16)
	ns.Blind()
	ns.SetOwner(ns.GetHost(), obj50)
	ns.FrameTimer(45, MoveToFacade)
}
func KirikStart2() {
	ns.TellStory(ns.HumanMaleEatFood, "War03a:Guard1Talk02")
}
func KirikEnd2() {
	ns.SetDialog(obj39, ns.NORMAL, KirikStart2, KirikEnd2)
}
func HorStart1() {
	flag36 = false
	ns.LookAtObject(obj50, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "War02B.scr:HorrendousTalk1Start")
}
func HorEnd1() {
	ns.UnlockDoor(obj56)
	ns.UnlockDoor(obj57)
	ns.JournalEntry(ns.GetHost(), "War03aIxQuest", 2)
	ns.PrintToAll("Con02a:ItemsAddedToInventory")
	ns.Pickup(ns.GetHost(), obj62)
	ns.Pickup(ns.GetHost(), obj63)
	ns.Pickup(ns.GetHost(), obj64)
	ns.Pickup(ns.GetHost(), obj70)
	ns.SetDialog(obj50, ns.NORMAL, HorStart2, HorEnd2)
}
func HorStart2() {
	ns.TellStory(ns.HumanMaleEatFood, "War02A.scr:RealTalk2Start")
}
func HorEnd2() {
	ns.SetDialog(obj50, ns.NORMAL, HorStart2, HorEnd2)
}
func HorTalkStart() {
	if !flag36 {
		goto LABEL1
	}
	flag36 = false
	ns.StartDialog(obj50, ns.GetHost())
LABEL1:
	return
}
func GHInRoom() {
	if !ns.IsCaller(obj16) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	flag87 = true
	ns.CreatureGuard(obj16, ns.GetObjectX(obj16), ns.GetObjectY(obj16), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	if !flag86 {
		goto LABEL2
	}
	ns.SetDialog(obj16, ns.NORMAL, GHStart4, GHEnd4)
	goto LABEL1
LABEL2:
	ns.SetDialog(obj16, ns.YESNO, GHStart5, GHEnd5)
LABEL1:
	return
}
func PlayerInGearRoom() {
	if !ns.IsVisibleTo(obj17, ns.GetHost()) {
		goto LABEL1
	}
	ns.ObjectGroupOff(gvar81)
	ns.SetDialog(obj17, ns.NORMAL, QMStart6, QMEnd6)
	ns.StartDialog(obj17, ns.GetHost())
LABEL1:
	return
}
func PlayerExitSewers() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.MoveObject(obj17, 2392, 3545)
	ns.MoveObject(obj16, 2457, 3476)
	ns.MoveObject(obj50, 3395, 1224)
	ns.MoveObject(obj59, 3348, 1195)
	ns.MoveObject(obj60, 3371, 1278)
	ns.LookAtObject(obj16, ns.GetHost())
	ns.LookAtObject(obj17, ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj17)
	ns.Frozen(ns.GetHost(), true)
	ns.CreatureGuard(obj17, ns.GetObjectX(obj17), ns.GetObjectY(obj17), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	ns.CreatureGuard(obj16, ns.GetObjectX(obj16), ns.GetObjectY(obj16), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	ns.CreatureGuard(obj50, ns.GetObjectX(obj50), ns.GetObjectY(obj50), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	ns.CreatureGuard(obj59, ns.GetObjectX(obj59), ns.GetObjectY(obj59), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	ns.CreatureGuard(obj60, ns.GetObjectX(obj60), ns.GetObjectY(obj60), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	ns.FrameTimer(30, Congrats)
LABEL1:
	return
}
func Congrats() {
	ns.SetDialog(obj17, ns.NEXT, QMStart7, QMEnd7)
	ns.StartDialog(obj17, ns.GetHost())
}
func DropLetter() {
	ns.Delete(obj52)
	ns.FrameTimer(1, MovePlayer)
}
func MovePlayer() {
	ns.MoveObject(ns.GetHost(), 2586, 5309)
}
func SecretTrainingRoom() {
	ns.ObjectGroupOff(gvar79)
	secretSound()
	ns.GiveXp(ns.GetHost(), 100)
	ns.PrintToAll("GeneralPrint:SecretFound")
}
func SecretRoom() {
	ns.ObjectGroupOff(gvar80)
	ns.GiveXp(ns.GetHost(), 100)
	secretSound()
	ns.PrintToAll("GeneralPrint:SecretFound")
}
func GeneralSecret() {
	ns.GiveXp(ns.GetHost(), 100)
	secretSound()
	ns.PrintToAll("GeneralPrint:SecretFound")
}
func secretSound() {
	ns.MoveWaypoint(wp22, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretFound, wp22)
}
func QMEndTrigger() {
	if !ns.IsCaller(obj17) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	flag26 = true
	ns.CreatureGuard(obj17, ns.GetObjectX(obj17), ns.GetObjectY(obj17), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
LABEL1:
	return
}
func FKG1Start() {
	ns.TellStory(ns.SwordsmanRecognize, "War01A.scr:Guard2Talk1aStart")
}
func FKG1End() {
	ns.SetDialog(obj117, ns.NORMAL, FKG1Start, FKG1End)
}
func SecretSurprise() {
	ns.ObjectOn(gvar73)
	ns.MoveObject(gvar73, 4416, 1000)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 50)
	secretSound()
}
func SecretSword() {
	secretSound()
	ns.GiveXp(ns.GetHost(), 150)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.FrameTimer(15, FlamesOut)
}
func FlamesOut() {
	ns.GroupDelete(gvar122)
	ns.AudioEvent(ns.FireExtinguish, wp124)
}
func gauntletQuest() {
	ns.Music(1, 100)
	ns.JournalEntry(ns.GetHost(), "FirstQuest", 2)
	ns.PrintToAll("Con01a:NewJournalEntry")
}
func MakeAWish() {
	var v0 int
	v0 = ns.MaxHealth(ns.GetCaller())
	ns.RestoreHealth(ns.GetCaller(), v0-ns.CurrentHealth(ns.GetCaller()))
	ns.MoveWaypoint(wp22, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.PrintToAll("GeneralPrint:WellSignRefresh")
	ns.AudioEvent(ns.RestoreHealthName, wp22)
}
func CaveMusic() {
	ns.ObjectOff(obj113)
	ns.ObjectGroupOn(gvar21)
	ns.Music(19, 100)
}
func TownMusic() {
	ns.ObjectGroupOff(gvar21)
	if gvar118 == 0 {
		goto LABEL1
	}
	ns.ObjectOn(obj113)
LABEL1:
	ns.Music(16, 100)
}
func runFromPain() {
	ns.RunAway(ns.GetTrigger(), ns.GetHost(), 60)
	obj94 = ns.GetTrigger()
	ns.RestoreHealth(obj94, 1000)
	if !flag37 {
		goto LABEL1
	}
	ns.FrameTimer(63, followPlayer)
LABEL1:
	if ns.GetTrigger() != obj17 {
		goto LABEL2
	}
	ns.FrameTimer(63, followPlayer)
LABEL2:
	return
}
func followPlayer() {
	ns.CreatureFollow(obj94, ns.GetHost())
}
func InitializeWizardSetpiece() {
	obj136 = ns.Object("Wizard")
	obj137 = ns.Object("Wizard2")
	obj138 = ns.Object("Jerry")
	obj139 = ns.Object("Bryan")
	obj140 = ns.Object("FireKnightGuard1")
	obj141 = ns.Object("Staff")
	wp142 = ns.Waypoint("WizardMark")
	wp143 = ns.Waypoint("BryanMark")
	wp144 = ns.Waypoint("AudioOrigin")
	wp145 = ns.Waypoint("WizardStorage")
	wp146 = ns.Waypoint("WizardExitPath")
	ns.Pickup(obj136, obj141)
	ns.SetOwner(ns.GetHost(), obj139)
	ns.SetOwner(ns.GetHost(), obj138)
	ns.CastSpellObjectObject(ns.SPELL_SHIELD, obj136, obj136)
	ns.Enchant(obj136, ns.ENCHANT_INVISIBLE, 0)
}
func CreatureSetup() {
	ns.SetOwner(ns.GetHost(), obj115)
	ns.SetOwner(ns.GetHost(), obj116)
	ns.SetOwner(ns.GetHost(), obj39)
	ns.SetOwner(ns.GetHost(), obj43)
	ns.SetOwner(ns.GetHost(), obj44)
	ns.SetOwner(ns.GetHost(), obj45)
	ns.SetOwner(ns.GetHost(), obj41)
	ns.SetOwner(ns.GetHost(), obj42)
	ns.SetOwner(ns.GetHost(), obj46)
	ns.SetOwner(ns.GetHost(), obj47)
	ns.SetOwner(ns.GetHost(), obj48)
	ns.SetOwner(ns.GetHost(), obj49)
	ns.SetOwner(ns.GetHost(), obj16)
	ns.SetOwner(ns.GetHost(), obj17)
	ns.SetOwner(ns.GetHost(), obj14)
	ns.SetOwner(ns.GetHost(), obj15)
	ns.SetOwner(ns.GetHost(), obj117)
	ns.SetOwner(ns.GetHost(), obj59)
	ns.SetOwner(ns.GetHost(), obj60)
	ns.SetOwner(ns.GetHost(), obj39)
	if ns.GetQuestStatus("War02A:EnteredGauntlet") != 1 {
		goto LABEL1
	}
	if !flag38 {
		goto LABEL1
	}
	flag38 = false
	ns.Delete(obj14)
	ns.Delete(obj15)
	ns.UnlockDoor(obj34)
	ns.UnlockDoor(obj120)
	ns.UnlockDoor(obj119)
	ns.CancelDialog(obj115)
	ns.CancelDialog(obj116)
	ns.MoveObject(obj42, 2724, 3161)
	ns.CreatureGuard(obj42, 2724, 3161, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	ns.SetDialog(obj49, ns.NORMAL, LydiaStart2, LydiaEnd2)
	ns.SetDialog(obj17, ns.NEXT, QMStart4, QMEnd4)
	ns.SetDialog(obj41, ns.NORMAL, DanStart1, DanEnd1)
	ns.SetDialog(obj42, ns.NORMAL, JesseStart2, JesseEnd2)
	ns.SetDialog(obj44, ns.NORMAL, EricStart1, EricEnd1)
	ns.SetDialog(obj45, ns.NORMAL, JacobStart1, JacobEnd1)
	ns.MoveObject(obj16, 2590, 3279)
	ns.CreatureGuard(obj16, ns.GetObjectX(obj16), ns.GetObjectY(obj16), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	if ns.GetQuestStatus("PlayerInSewers") == 1 {
		goto LABEL2
	}
	ns.CreatureFollow(obj17, ns.GetHost())
	ns.StartDialog(obj17, ns.GetHost())
	goto LABEL3
LABEL2:
	ns.MoveObject(obj17, 2296, 3099)
	ns.CreatureGuard(obj17, 2296, 3099, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
LABEL3:
	ns.SetShopkeeperText(obj43, "War02B.scr:BarTalk1")
	ns.SetShopkeeperText(obj114, "War03b:Shopkeeper")
	ns.ObjectGroupOn(gvar74)
	ns.ObjectGroupOn(gvar75)
	ns.ObjectGroupOn(gvar76)
	ns.ObjectGroupOn(gvar77)
	ns.ObjectOn(obj53)
	ns.ObjectOn(obj54)
	ns.ObjectOn(obj55)
	ns.MoveObject(obj41, 3829, 3524)
	ns.FrameTimer(1, DisableDudes)
LABEL1:
	if ns.GetQuestStatus("War02A:EnteredGauntlet") != 2 {
		goto LABEL4
	}
	if !flag27 {
		goto LABEL4
	}
	flag27 = false
	ns.LockDoor(obj33)
	ns.LockDoor(obj34)
	ns.SetDialog(obj41, ns.NORMAL, DanStart1, DanEnd1)
	ns.SetDialog(obj42, ns.NORMAL, JesseStart2, JesseEnd2)
	ns.SetDialog(obj44, ns.NORMAL, EricStart1, EricEnd1)
	ns.SetDialog(obj45, ns.NORMAL, JacobStart1, JacobEnd1)
	ns.ObjectGroupOn(gvar74)
	ns.ObjectGroupOn(gvar75)
	ns.ObjectGroupOn(gvar76)
	ns.ObjectGroupOn(gvar77)
	ns.ObjectOn(obj53)
	ns.ObjectOn(obj54)
	ns.ObjectOn(obj55)
	ns.LockDoor(obj34)
	ns.LockDoor(obj33)
	ns.UnlockDoor(obj57)
	ns.UnlockDoor(obj56)
	ns.SetShopkeeperText(obj114, "War03b:Shopkeeper")
	ns.CancelDialog(obj115)
	ns.CancelDialog(obj116)
	ns.ObjectGroupOn(gvar74)
	PlayerExitSewers()
LABEL4:
	return
}
func GateGuard1AStart() {
	ns.LookAtObject(obj115, ns.GetHost())
	ns.Frozen(obj115, true)
	ns.TellStory(ns.HumanMaleEatFood, "War01A.scr:GateGuard1Talk01")
}
func GateGuard1AEnd() {
	ns.Frozen(obj115, false)
	ns.SetDialog(obj115, ns.NORMAL, GateGuard1BStart, GateGuard1BEnd)
}
func GateGuard1BStart() {
	ns.LookAtObject(obj115, ns.GetHost())
	ns.Frozen(obj115, true)
	ns.TellStory(ns.HumanMaleEatFood, "War01A.scr:GateGuard1Talk02")
}
func GateGuard1BEnd() {
	ns.Frozen(obj115, false)
	ns.SetDialog(obj115, ns.NORMAL, GateGuard1BStart, GateGuard1BEnd)
}
func GateGuard2AStart() {
	ns.LookAtObject(obj116, ns.GetHost())
	ns.Frozen(obj116, true)
	ns.TellStory(ns.HumanMaleEatFood, "War01A.scr:GateGuard2Talk01")
}
func GateGuard2AEnd() {
	ns.Frozen(obj116, false)
	ns.SetDialog(obj116, ns.NORMAL, GateGuard2BStart, GateGuard2BEnd)
}
func GateGuard2BStart() {
	ns.LookAtObject(obj116, ns.GetHost())
	ns.Frozen(obj116, true)
	ns.TellStory(ns.HumanMaleEatFood, "War01A.scr:GateGuard2Talk02")
}
func GateGuard2BEnd() {
	ns.Frozen(obj116, false)
	ns.SetDialog(obj116, ns.NORMAL, GateGuard2AStart, GateGuard2AEnd)
}
func OpenGearDoor() {
	if !ns.IsCaller(obj16) {
		goto LABEL1
	}
	ns.UnlockDoor(obj119)
	ns.UnlockDoor(obj120)
	ns.ObjectOff(ns.GetTrigger())
LABEL1:
	return
}
func CaptainStartTalk() {
	ns.StartDialog(obj95, ns.GetHost())
}
func CaptainStart1() {
	ns.TellStory(ns.HumanMaleEatFood, "War01A.scr:CaptainTalkStart")
}
func CaptainEnd1() {
	flag88 = true
	ns.SetDialog(obj95, ns.NORMAL, CaptainStart2, CaptainEnd2)
}
func CaptainStart2() {
	var (
		v0 int
		v1 int
	)
	v0 = ns.Random(1, 4)
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
	goto LABEL5
LABEL1:
	ns.TellStory(ns.HumanMaleEatFood, "War01a:CaptainTalk2aStart")
	goto LABEL5
LABEL2:
	ns.TellStory(ns.HumanMaleEatFood, "War01a:CaptainTalk2bStart")
	goto LABEL5
LABEL3:
	ns.TellStory(ns.HumanMaleEatFood, "War01a:CaptainTalk2cStart")
	goto LABEL5
LABEL4:
	ns.TellStory(ns.HumanMaleEatFood, "War01a:CaptainTalk2eStart")
	goto LABEL5
LABEL5:
	return
}
func CaptainEnd2() {
	ns.SetDialog(obj95, ns.NORMAL, CaptainStart2, CaptainEnd2)
}
func KirikStart1() {
	ns.TellStory(ns.HumanMaleEatFood, "War01A.scr:Guard1Talk01")
}
func KirikEnd1() {
	ns.SetDialog(obj39, ns.NORMAL, KirikStart1, KirikEnd1)
}
func GiveXP() {
	ns.ObjectGroupOff(gvar121)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 50)
}
func DisableDudes() {
	ns.ObjectOff(obj41)
}
func IntroBatAttack() {
	if !flag93 {
		goto LABEL1
	}
	if !ns.IsTrading() {
		goto LABEL2
	}
	ns.Frozen(gvar40, true)
	ns.FrameTimer(10, ResetBat)
LABEL2:
	ns.FrameTimer(10, IntroBatAttack)
LABEL1:
	return
}
func ResetBat() {
	if !flag93 {
		goto LABEL1
	}
	ns.Frozen(gvar40, false)
LABEL1:
	return
}
func IntroBatDies() {
	flag93 = false
}
func DestroyIntroBat() {
	ns.Delete(gvar40)
}
func MapInitialize() {
	obj95 = ns.Object("AirshipCaptain")
	obj96 = ns.Object("Basket")
	obj97 = ns.Object("BasketShadow")
	obj106 = ns.Object("BarrelSign")
	obj107 = ns.Object("WallSign")
	obj108 = ns.Object("ChestSign")
	obj109 = ns.Object("JumpSign")
	obj110 = ns.Object("WellSign")
	obj111 = ns.Object("AppleSign")
	obj100 = ns.Object("DownGear")
	obj101 = ns.Object("UpGear")
	obj98 = ns.Object("UpToSurface")
	obj99 = ns.Object("ElevatorDoor")
	obj16 = ns.Object("F2Gearhart")
	obj50 = ns.Object("Horrendous")
	obj102 = ns.Object("MidCityGate1")
	obj103 = ns.Object("MidCityGate2")
	obj104 = ns.Object("GauntletDoorL")
	obj105 = ns.Object("GauntletDoorR")
	gvar73 = ns.Object("SecretSpider")
	ns.ObjectOff(gvar73)
	obj71 = ns.Object("MachRoomTrigger")
	ns.ObjectOff(obj71)
	obj72 = ns.Object("TalkWithHorTrigger")
	ns.ObjectOff(obj72)
	obj113 = ns.Object("CaveMusicTrigger")
	gvar21 = ns.ObjectGroup("TownMusicTriggers")
	ns.ObjectGroupOff(gvar21)
	obj39 = ns.Object("Kirik")
	gvar40 = ns.Object("IntroBat")
	obj114 = ns.Object("F2Bull")
	obj115 = ns.Object("Bruce")
	obj116 = ns.Object("F6FireGuard10")
	obj41 = ns.Object("Daniel")
	obj42 = ns.Object("Jesse")
	obj45 = ns.Object("Jacob")
	obj44 = ns.Object("Eric")
	obj46 = ns.Object("Tyler")
	obj43 = ns.Object("Bing")
	obj47 = ns.Object("Melissa")
	obj48 = ns.Object("Evelyn")
	obj49 = ns.Object("Lydia")
	obj117 = ns.Object("F1Rogan")
	obj59 = ns.Object("FireKnightGuard3")
	obj60 = ns.Object("FireKnightGuard4")
	obj17 = ns.Object("QuarterMaster")
	obj10 = ns.Object("AppleBat1")
	obj12 = ns.Object("AppleBat1A")
	obj13 = ns.Object("AppleBat1B")
	obj11 = ns.Object("AppleBat2")
	obj14 = ns.Object("Kristine")
	obj15 = ns.Object("Jennifer")
	gvar118 = ns.Object("MasterUrchin")
	obj53 = ns.Object("GearLight1")
	obj54 = ns.Object("GearLight2")
	obj55 = ns.Object("GearLight3")
	obj119 = ns.Object("GearDoorR")
	obj120 = ns.Object("GearDoorL")
	obj51 = ns.Object("OrchardKey")
	obj52 = ns.Object("SponsorshipLetter")
	obj28 = ns.Object("SewerGear1")
	obj29 = ns.Object("SewerGear2")
	obj30 = ns.Object("Brick1")
	obj31 = ns.Object("Brick2")
	obj32 = ns.Object("SewerBottom")
	obj33 = ns.Object("ExpressDoor")
	obj34 = ns.Object("ServiceDoor")
	obj57 = ns.Object("IxDoorR")
	obj56 = ns.Object("IxDoorL")
	obj58 = ns.Object("SewerExitTrigger")
	obj61 = ns.Object("Goodies")
	obj62 = ns.Object("HorGold")
	obj63 = ns.Object("SturdyBoots")
	obj64 = ns.Object("NiceCloak")
	obj65 = ns.Object("BatScroll")
	obj66 = ns.Object("LeechScroll")
	obj67 = ns.Object("SpiderScroll")
	obj68 = ns.Object("LargeSpiderScroll")
	obj69 = ns.Object("TrollScroll")
	obj70 = ns.Object("UrchinScroll")
	obj112 = ns.Object("LightTrigger")
	gvar74 = ns.ObjectGroup("Peds")
	gvar18 = ns.ObjectGroup("FleePiece")
	gvar19 = ns.ObjectGroup("FleePiece2")
	gvar75 = ns.ObjectGroup("StreetLamps")
	gvar76 = ns.ObjectGroup("StreetLamps2")
	gvar77 = ns.ObjectGroup("LightGears")
	gvar78 = ns.ObjectGroup("LightGears2")
	gvar20 = ns.ObjectGroup("FleeTriggers")
	gvar121 = ns.ObjectGroup("SecretAreaWow")
	gvar79 = ns.ObjectGroup("SecretTrainingTriggers")
	gvar80 = ns.ObjectGroup("SecretRoom2Triggers")
	gvar81 = ns.ObjectGroup("PlayerGearTriggers")
	gvar122 = ns.ObjectGroup("FireSwordFlames")
	gvar123 = ns.WallGroup("SetPieceWalls")
	wp22 = ns.Waypoint("PlayerSounds")
	wp91 = ns.Waypoint("GearhartWP")
	wp92 = ns.Waypoint("QMStartWP")
	wp124 = ns.Waypoint("FireSnufWay")
	gvar125 = ns.WaypointGroup("CaveWP")
	ns.StoryPic(obj117, "WarriorPic")
	ns.StoryPic(obj41, "MalePic5")
	ns.StoryPic(obj42, "Townsman1Pic")
	ns.StoryPic(obj44, "Townsman3Pic")
	ns.StoryPic(obj17, "QuarterMasterPic")
	ns.StoryPic(obj46, "IxGuard1Pic")
	ns.StoryPic(obj49, "MaidenPic3")
	ns.StoryPic(obj47, "MaidenPic2")
	ns.StoryPic(obj48, "MaidenPic")
	ns.StoryPic(obj45, "DrunkPic")
	ns.StoryPic(obj16, "GearhartPic")
	ns.StoryPic(obj95, "AirshipCaptainPic")
	ns.StoryPic(obj39, "Warrior4Pic")
	ns.StoryPic(obj14, "MaidenPic2")
	ns.StoryPic(obj15, "MaidenPic")
	ns.SetDialog(obj117, ns.NORMAL, FKG1Start, FKG1End)
	ns.SetDialog(obj115, ns.NORMAL, GateGuard1AStart, GateGuard1AEnd)
	ns.SetDialog(obj116, ns.NORMAL, GateGuard2AStart, GateGuard2AEnd)
	ns.SetDialog(obj41, ns.NORMAL, DanStart1, DanEnd1)
	ns.SetDialog(obj42, ns.NORMAL, JesseStart1, JesseEnd1)
	ns.SetDialog(obj44, ns.NORMAL, EricStart1, EricEnd1)
	ns.SetDialog(obj46, ns.NORMAL, TylerStart1, TylerEnd1)
	ns.SetDialog(obj45, ns.NORMAL, JacobStart1, JacobEnd1)
	ns.SetDialog(obj49, ns.NORMAL, LydiaStart1, LydiaEnd1)
	ns.SetDialog(obj47, ns.NORMAL, MelStart1, MelEnd1)
	ns.SetDialog(obj48, ns.NORMAL, EvelynStart1, EvelynEnd1)
	ns.SetDialog(obj17, ns.NORMAL, QMStart1, QMEnd1)
	ns.SetDialog(obj95, ns.NORMAL, CaptainStart1, CaptainEnd1)
	ns.SetDialog(obj39, ns.NORMAL, KirikStart1, KirikEnd1)
	ns.ObjectGroupOff(gvar18)
	ns.ObjectGroupOff(gvar19)
	ns.WayPointGroupOff(gvar125)
	ns.ObjectGroupOff(gvar74)
	ns.LockDoor(obj33)
	ns.LockDoor(obj34)
	ns.LockDoor(obj119)
	ns.LockDoor(obj120)
	ns.LockDoor(obj56)
	ns.LockDoor(obj57)
	ns.LockDoor(obj102)
	ns.LockDoor(obj103)
	ns.LockDoor(obj104)
	ns.LockDoor(obj105)
	ns.StoryPic(obj50, "HorrendousPic")
	ns.StoryPic(obj115, "Warrior5Pic")
	ns.StoryPic(obj116, "Warrior3Pic")
	InitializeWizardSetpiece()
	ns.StartupScreen(1)
	ns.FrameTimer(3, gauntletQuest)
	ns.FrameTimer(1, CaptainStartTalk)
}
func MapEntry() {
	if ns.GetQuestStatus("War02A:EnteredGauntlet") != 1 {
		goto LABEL1
	}
	ns.Music(16, 100)
LABEL1:
	if ns.GetQuestStatus("War02A:EnteredGauntlet") != 2 {
		goto LABEL2
	}
	ns.Music(16, 100)
LABEL2:
	ns.FrameTimer(1, CreatureSetup)
}
func PlayerDeath() {
	if ns.GetQuestStatus("War02A:EnteredGauntlet") != 1 {
		goto LABEL1
	}
	ns.DeathScreen(2)
	goto LABEL2
LABEL1:
	if ns.GetQuestStatus("War02A:EnteredGauntlet") != 2 {
		goto LABEL3
	}
	ns.DeathScreen(3)
	goto LABEL2
LABEL3:
	ns.DeathScreen(1)
LABEL2:
	return
}
func WizardEnter() {
	gvar133 = gvar127
	ns.Move(obj136, wp142)
	ns.Move(obj139, wp143)
}
func JerrySpeaks() {
	ns.LookAtObject(obj138, obj136)
	ns.Chat(obj138, "War01A.scr:SetpieceGuard1Talk01")
	ns.FrameTimer(60, WizardAttackJerry)
}
func WizardAttackJerry() {
	gvar133 = gvar129
	ns.LookAtObject(obj136, obj138)
	ns.LookAtObject(obj138, obj136)
	ns.Chat(obj138, "War01A.scr:SetpieceGuard2Talk01")
	ns.CastSpellObjectLocation(ns.SPELL_LIGHTNING, obj136, ns.GetObjectX(obj138), ns.GetObjectY(obj138))
	ns.FrameTimer(15, BryanAttackWizard)
}
func BryanAttackWizardAgain() {
	ns.Attack(obj139, obj136)
}
func BryanSpeaks() {
	ns.LookAtObject(obj139, obj137)
	ns.LookAtObject(obj138, obj137)
	ns.Chat(obj139, "War01A.scr:SetpieceGuard2Talk02")
	ns.FrameTimer(80, WizardExit)
}
func WizardExit() {
	ns.DestroyEveryChat()
	ns.Move(obj139, wp146)
	ns.CreatureFollow(obj137, obj139)
	ns.CreatureFollow(obj138, obj137)
	gvar133 = gvar131
	ns.FrameTimer(120, ReleasePlayer)
}
func ReleasePlayer() {
	ns.Music(16, 100)
	ns.WideScreen(false)
	ns.MoveObject(obj138, ns.GetWaypointX(wp145), ns.GetWaypointY(wp145))
	ns.MoveObject(obj139, ns.GetWaypointX(wp145), ns.GetWaypointY(wp145))
	ns.MoveObject(obj137, ns.GetWaypointX(wp145), ns.GetWaypointY(wp145))
	ns.Frozen(ns.GetHost(), false)
	ns.ObjectGroupOn(gvar74)
	ns.NoWallSound(true)
	ns.WallGroupOpen(gvar123)
	ns.NoWallSound(false)
	ns.FrameTimer(1, DisableGuys)
}
func DisableGuys() {
	ns.ObjectOff(obj138)
	ns.ObjectOff(obj139)
	ns.ObjectOff(obj137)
}
func WizardReport() {
	var v0 int
	v0 = gvar133
	if v0 == gvar127 {
		goto LABEL1
	}
	goto LABEL2
LABEL1:
	ns.EnchantOff(obj136, ns.ENCHANT_INVISIBLE)
	ns.LookAtObject(obj136, obj138)
	gvar133 = gvar128
	ns.FrameTimer(15, JerrySpeaks)
	goto LABEL2
LABEL2:
	return
}
func StopBryan() {
	ns.CreatureIdle(obj139)
}
func SwapWizards() {
	ns.Drop(obj136, obj141)
	ns.AudioEvent(ns.WoodenWeaponDrop, wp144)
	ns.MoveObject(obj137, ns.GetObjectX(obj136), ns.GetObjectY(obj136))
	ns.MoveObject(obj136, ns.GetWaypointX(wp145), ns.GetWaypointY(wp145))
}
func WizardInjured() {
	var v0 int
	ivar134 += 1
	if !(ivar134 > 1) {
		goto LABEL1
	}
	v0 = gvar133
	if v0 == gvar129 {
		goto LABEL2
	}
	if v0 == gvar131 {
		goto LABEL3
	}
	goto LABEL1
LABEL2:
	ns.DestroyEveryChat()
	ns.FrameTimer(5, SwapWizards)
	ns.AudioEvent(ns.NPCHurt, wp144)
	gvar133 = gvar130
	ns.CreatureIdle(obj136)
	ns.FrameTimer(10, StopBryan)
	ns.FrameTimer(20, BryanSpeaks)
	goto LABEL1
LABEL3:
	if !ns.IsCaller(obj140) {
		goto LABEL4
	}
	ns.Damage(ns.GetTrigger(), 0, 1000, 0)
LABEL4:
	goto LABEL1
LABEL1:
	return
}
func StartWizardSetpiece() {
	if !(!flag135 && ns.IsCaller(ns.GetHost())) {
		goto LABEL1
	}
	if gvar40 == 0 {
		goto LABEL2
	}
	ns.Damage(gvar40, 0, 100, 0)
	ns.FrameTimer(30, DestroyIntroBat)
LABEL2:
	ns.Music(4, 100)
	flag135 = true
	ns.Frozen(ns.GetHost(), true)
	ns.CreatureIdle(ns.GetHost())
	ns.WideScreen(true)
	ns.MoveWaypoint(wp144, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.Delete(obj95)
	ns.Delete(obj96)
	ns.Delete(obj97)
	ns.FrameTimer(45, WizardEnter)
LABEL1:
	return
}
func BryanAttackWizard() {
	ns.AudioEvent(ns.SwordsmanRecognize, wp144)
	ns.Attack(obj139, obj136)
	ns.FrameTimer(7, BryanAttackWizardAgain)
}
func Clear() {
	ns.WideScreen(false)
	ns.Frozen(ns.GetHost(), false)
}
func OnEvent(typ string) {
	switch typ {
	case "MapInitialize":
		MapInitialize()
	case "MapEntry":
		MapEntry()
	case "PlayerDeath":
		PlayerDeath()
	}
}
