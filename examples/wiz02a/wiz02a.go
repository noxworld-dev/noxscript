package wiz02a

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
	flag15  bool
	gvar16  ns.WallID
	wp17    ns.WaypointID
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
	obj73   ns.ObjectID
	obj74   ns.ObjectID
	obj75   ns.ObjectID
	obj76   ns.ObjectID
	obj77   ns.ObjectID
	obj78   ns.ObjectID
	obj79   ns.ObjectID
	gvar80  ns.ObjectGroupID
	gvar81  ns.ObjectGroupID
	gvar82  ns.WallGroupID
	wp83    ns.WaypointID
	wp84    ns.WaypointID
	wp85    ns.WaypointID
	wp86    ns.WaypointID
	wp87    ns.WaypointID
	wp88    ns.WaypointID
	wp89    ns.WaypointID
	gvar90  ns.WaypointGroupID
	gvar91  ns.WaypointGroupID
	flag92  bool
	flag93  bool
	flag94  bool
	flag95  bool
	flag96  bool
	flag97  bool
	flag98  bool
	flag99  bool
	flag100 bool
	flag101 bool
	flag102 bool
	flag103 bool
	flag104 bool
	flag105 bool
	flag106 bool
	flag107 bool
	gvar108 int
	ivar109 int
	ivar110 int
	ivar111 int
	ivar112 int
	ivar113 int
	gvar114 ns.TimerID
)

func init() {
	flag15 = false
	flag92 = false
	flag93 = false
	flag94 = false
	flag95 = false
	flag96 = false
	flag97 = false
	flag98 = false
	flag99 = false
	flag100 = false
	flag101 = false
	flag102 = false
	flag103 = false
	flag104 = false
	flag105 = false
	flag106 = false
	flag107 = false
	ivar111 = 0
	ivar112 = 0
	ivar113 = 0
}
func CrowdQuiet() {
	if flag15 {
		goto LABEL1
	}
	flag15 = true
	ns.ObjectOff(obj4)
	ns.ObjectOff(obj5)
	ns.ObjectOff(obj6)
	ns.Frozen(obj8, true)
	ns.Frozen(obj9, true)
	ns.Frozen(obj10, true)
	ns.Frozen(obj11, true)
	ns.Frozen(obj12, true)
	ns.Frozen(obj13, true)
	ns.Frozen(obj14, true)
	ns.CreatureIdle(obj8)
	ns.CreatureIdle(obj9)
	ns.CreatureIdle(obj10)
	ns.CreatureIdle(obj11)
	ns.CreatureIdle(obj12)
	ns.CreatureIdle(obj13)
	ns.CreatureIdle(obj14)
	ns.LookAtObject(obj8, ns.GetHost())
	ns.LookAtObject(obj9, ns.GetHost())
	ns.LookAtObject(obj10, ns.GetHost())
	ns.LookAtObject(obj11, ns.GetHost())
	ns.LookAtObject(obj12, ns.GetHost())
	ns.LookAtObject(obj13, ns.GetHost())
	ns.LookAtObject(obj14, ns.GetHost())
	CrowdFacePlayer()
LABEL1:
	return
}
func CrowdNoisy() {
	if !flag15 {
		goto LABEL1
	}
	if ns.IsObjectOn(obj4) {
		goto LABEL1
	}
	ns.ObjectOn(obj4)
	ns.ObjectOn(obj5)
	ns.ObjectOn(obj6)
	ns.LookAtObject(obj8, obj9)
	ns.LookAtObject(obj9, obj8)
	ns.LookAtObject(obj10, obj11)
	ns.LookAtObject(obj11, obj10)
	ns.LookAtObject(obj12, obj11)
	ns.LookAtObject(obj13, obj7)
	ns.LookAtObject(obj14, obj7)
	flag15 = false
LABEL1:
	return
}
func CrowdFacePlayer() {
	if !flag15 {
		goto LABEL1
	}
	ns.LookAtObject(obj8, ns.GetHost())
	ns.LookAtObject(obj9, ns.GetHost())
	ns.LookAtObject(obj10, ns.GetHost())
	ns.LookAtObject(obj11, ns.GetHost())
	ns.LookAtObject(obj12, ns.GetHost())
	ns.LookAtObject(obj13, ns.GetHost())
	ns.LookAtObject(obj14, ns.GetHost())
	ns.FrameTimer(2, CrowdFacePlayer)
LABEL1:
	return
}
func SecretsInit() {
	gvar16 = ns.Wall(97, 69)
	wp17 = ns.Waypoint("SecretSounds")
}
func FireballSecret() {
	ns.ObjectOff(ns.GetTrigger())
	ns.MoveWaypoint(wp17, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretFound, wp17)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 50)
}
func OpenPushSecret() {
	ns.WallOpen(gvar16)
}
func PushSecretXP() {
	ns.ObjectOff(ns.GetTrigger())
	ns.MoveWaypoint(wp17, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretFound, wp17)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 25)
}
func MakeAWish() {
	var v0 int
	v0 = ns.MaxHealth(ns.GetCaller())
	ns.RestoreHealth(ns.GetCaller(), v0-ns.CurrentHealth(ns.GetCaller()))
	ns.PrintToAll("GeneralPrint:WellSignRefresh")
	ns.AudioEvent(ns.RestoreHealthName, ns.Waypoint("WellWP"))
}
func PlayerEnterCell() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	flag104 = true
	ns.UnlockDoor(obj64)
LABEL1:
	return
}
func PlayerExitCell() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	flag104 = false
	ns.LockDoor(obj64)
LABEL1:
	return
}
func GenericStart() {
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02A.scr:UndertakerTalk01")
}
func GenericEnd() {
	ns.SetDialog(obj30, ns.NORMAL, GenericStart, GenericEnd)
}
func GrillfStart() {
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02A.scr:GrillfTalk01")
}
func GrillfEnd() {
	ns.SetDialog(obj28, ns.NORMAL, GrillfStart, GrillfEnd)
}
func GrillfReset() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.LookAtObject(obj28, obj79)
LABEL1:
	return
}
func PriestStart() {
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02A.scr:PriestTalk01")
}
func PriestEnd() {
	ns.SetDialog(obj31, ns.NORMAL, PriestStart, PriestEnd)
}
func SetMaxWelcome() {
	flag92 = false
}
func MaxSayWelcome() {
	if flag92 {
		goto LABEL1
	}
	ns.Chat(obj27, "Wiz02A.scr:MaxTalk01")
	flag92 = true
LABEL1:
	return
}
func MaxRumor1Start() {
	ns.DestroyEveryChat()
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02A.scr:MaxRumor01")
}
func MaxRumor1End() {
	ChooseMaxRumor()
}
func MaxRumor2Start() {
	ns.DestroyEveryChat()
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02A.scr:MaxRumor02")
}
func MaxRumor2End() {
	ChooseMaxRumor()
}
func MaxRumor3Start() {
	ns.DestroyEveryChat()
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02A.scr:MaxRumor03")
}
func MaxRumor3End() {
	ChooseMaxRumor()
}
func MaxRumor4Start() {
	ns.DestroyEveryChat()
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02A.scr:MaxRumor04")
}
func MaxRumor4End() {
	ChooseMaxRumor()
}
func BarKeepRumor1Start() {
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02A.scr:LameBarKeepRumor01")
}
func BarKeepRumorEnd() {
	var v0 int
	ivar110 = ns.Random(1, 4)
	v0 = ivar110
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
	goto LABEL5
LABEL1:
	ns.SetDialog(obj7, ns.NORMAL, BarKeepRumor1Start, BarKeepRumorEnd)
	goto LABEL5
LABEL2:
	ns.SetDialog(obj7, ns.NORMAL, BarKeepRumor2Start, BarKeepRumorEnd)
	goto LABEL5
LABEL3:
	ns.SetDialog(obj7, ns.NORMAL, BarKeepRumor3Start, BarKeepRumorEnd)
	goto LABEL5
LABEL4:
	ns.SetDialog(obj7, ns.NORMAL, BarKeepRumor4Start, BarKeepRumorEnd)
	goto LABEL5
LABEL5:
	return
}
func BarKeepRumor2Start() {
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02A.scr:LameBarKeepRumor02")
}
func BarKeepRumor3Start() {
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02A.scr:LameBarKeepRumor03")
}
func BarKeepRumor4Start() {
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02A.scr:LameBarKeepRumor04")
}
func ChooseMaxRumor() {
	var v0 int
	ivar109 = ns.Random(1, 4)
	v0 = ivar109
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
	goto LABEL5
LABEL1:
	ns.SetDialog(obj27, ns.NORMAL, MaxRumor1Start, MaxRumor1End)
	goto LABEL5
LABEL2:
	ns.SetDialog(obj27, ns.NORMAL, MaxRumor2Start, MaxRumor2End)
	goto LABEL5
LABEL3:
	ns.SetDialog(obj27, ns.NORMAL, MaxRumor3Start, MaxRumor3End)
	goto LABEL5
LABEL4:
	ns.SetDialog(obj27, ns.NORMAL, MaxRumor4Start, MaxRumor4End)
	goto LABEL5
LABEL5:
	return
}
func MaxGiveRewardStart() {
	ns.DestroyEveryChat()
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02A.scr:MaxTalk03")
	ns.Pickup(ns.GetHost(), obj72)
}
func MaxGiveRewardEnd() {
	if !flag106 {
		goto LABEL1
	}
	ns.JournalEdit(ns.GetHost(), "ArrestRogues", 4)
	ns.PrintToAll("Con02a:ObjectiveComplete")
LABEL1:
	ChooseMaxRumor()
}
func MaxProblemStart() {
	ns.DestroyEveryChat()
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02A.scr:MaxTalk02")
}
func MaxProblemEnd() {
	if flag106 {
		goto LABEL1
	}
	ns.PrintToAll("Con01a:NewJournalEntry")
	ns.JournalEntry(ns.GetHost(), "ArrestRogues", 2)
	flag106 = true
LABEL1:
	ChooseMaxRumor()
}
func WardenGreetStart() {
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02A.scr:WardenTalk01")
}
func WardenGreetEnd() {
	if flag99 {
		goto LABEL1
	}
	ns.PrintToAll("Con01a:NewJournalEntry")
	ns.JournalEntry(ns.GetHost(), "ArrestMorgan", 2)
	flag99 = true
LABEL1:
	ns.SetDialog(obj29, ns.NORMAL, WardenGreetStart, WardenGreetEnd)
}
func Warden1Start() {
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02A.scr:WardenTalk02")
	flag97 = true
}
func Warden1End() {
	flag107 = true
	ns.CancelDialog(obj29)
	ns.GiveXp(ns.GetHost(), 500)
	ns.Pickup(ns.GetHost(), obj74)
	if !flag99 {
		goto LABEL1
	}
	ns.JournalEdit(ns.GetHost(), "ArrestMorgan", 4)
	ns.PrintToAll("Con02a:ObjectiveComplete")
LABEL1:
	ns.Frozen(ns.GetHost(), false)
	ns.SetDialog(obj27, ns.NORMAL, MaxGiveRewardStart, MaxGiveRewardEnd)
	ns.WideScreen(false)
	MorganEscortWarden()
}
func OpenJailDoor() {
	if !ns.IsCaller(obj29) {
		goto LABEL1
	}
	ns.UnlockDoor(obj61)
	ns.ObjectOff(obj68)
LABEL1:
	return
}
func EnableRogueElevator() {
	ns.MoveWaypoint(wp83, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.CreatureCageAppears, wp83)
	ns.AudioEvent(ns.ChangeSpellbar, wp83)
	ns.ObjectGroupOn(gvar80)
	ns.ObjectOn(obj66)
	ns.SetQuestStatus(1, "RogueElevatorEnabled")
}
func CryptChairTrigger() {
	if flag95 {
		goto LABEL1
	}
	flag95 = true
	ns.WallGroupOpen(gvar82)
	goto LABEL2
LABEL1:
	flag95 = false
	ns.WallGroupClose(gvar82)
LABEL2:
	return
}
func MorganSecretXP() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.ObjectGroupOff(gvar81)
	ns.MoveWaypoint(wp17, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretFound, wp17)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 150)
LABEL1:
	return
}
func MorganGo() {
	if !(flag93 == true && flag98 == false && flag107 == false) {
		goto LABEL1
	}
	ns.SetDialog(obj29, ns.NORMAL, Warden1Start, Warden1End)
	ns.MoveObject(obj29, 3365, 5093)
	ns.CreatureGuard(obj29, ns.GetObjectX(obj29), ns.GetObjectY(obj29), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
LABEL1:
	return
}
func MorganDied() {
	flag98 = true
	flag97 = true
	ns.SetQuestStatus(1, "MorganDead")
	ns.SetDialog(obj27, ns.NORMAL, MaxGiveRewardStart, MaxGiveRewardEnd)
	ns.SetDialog(obj29, ns.NORMAL, WardenGiveDeadRewardStart, WardenGiveDeadRewardEnd)
}
func MorganEscortWarden() {
	if !(ns.Distance(ns.GetObjectX(obj25), ns.GetObjectY(obj25), ns.GetObjectX(obj29), ns.GetObjectY(obj29)) < 250) {
		goto LABEL1
	}
	ns.WayPointGroupOn(gvar90)
	flag100 = true
	ns.Wander(obj29)
	ns.CreatureFollow(obj25, obj29)
	goto LABEL2
LABEL1:
	ns.FrameTimer(15, MorganEscortWarden)
LABEL2:
	return
}
func MoveWardenToTossTrigger() {
	if !ns.IsCaller(obj29) {
		goto LABEL1
	}
	ns.Walk(obj29, ns.GetObjectX(obj70), ns.GetObjectY(obj70))
LABEL1:
	return
}
func WardenTossMorgan() {
	if !ns.IsCaller(obj29) {
		goto LABEL1
	}
	ns.WayPointGroupOff(gvar90)
	if !(ns.IsVisibleTo(ns.GetHost(), obj25) || ns.IsVisibleTo(ns.GetHost(), obj29)) {
		goto LABEL2
	}
	ns.Chat(obj29, "Wiz02A.scr:WardenTalk03")
LABEL2:
	ns.Walk(obj25, 3588, 5426)
LABEL1:
	return
}
func WardenBack() {
	flag101 = true
	ns.SetDialog(obj29, ns.NORMAL, WardenGiveAliveRewardStart, WardenGiveAliveRewardEnd)
	ns.CreatureGuard(obj29, ns.GetObjectX(obj29), ns.GetObjectY(obj29), ns.GetObjectX(obj78), ns.GetObjectY(obj78), 0)
	ns.Frozen(obj29, true)
	ns.CreatureIdle(obj29)
	ns.LookAtObject(obj29, obj78)
}
func WardenGiveDeadRewardStart() {
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02A.scr:WardenTalk04")
	ns.Pickup(ns.GetHost(), obj73)
}
func WardenGiveDeadRewardEnd() {
	ns.GiveXp(ns.GetHost(), 300)
	ns.JournalEdit(ns.GetHost(), "ArrestMorgan", 4)
	ns.PrintToAll("Con02a:ObjectiveComplete")
	ns.CancelDialog(obj29)
}
func WardenGiveAliveRewardStart() {
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02A.scr:WardenTalk05")
}
func WardenGiveAliveRewardEnd() {
	ns.SetDialog(obj29, ns.NORMAL, WardenGiveAliveRewardStart, WardenGiveAliveRewardEnd)
}
func LockMorganCell() {
	if !ns.IsCaller(obj25) {
		goto LABEL1
	}
	ns.CreatureGuard(obj25, ns.GetObjectX(obj25), ns.GetObjectY(obj25), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	if flag104 {
		goto LABEL1
	}
	ns.LockDoor(obj64)
	ns.ObjectOff(obj69)
	ns.WayPointGroupOff(gvar90)
	ns.WayPointGroupOn(gvar91)
	ns.Wander(obj29)
	ns.ObjectOn(obj71)
	ns.SetQuestStatus(1, "MorganCaptured")
LABEL1:
	return
}
func Civvy1Start() {
	var v0 int
	ns.Frozen(obj39, true)
	ns.CreatureIdle(obj39)
	ns.LookAtObject(obj39, ns.GetHost())
	if !flag105 {
		goto LABEL1
	}
	ivar113 = 2
	goto LABEL2
LABEL1:
	ivar113 = ns.Random(1, 2)
LABEL2:
	v0 = ivar113
	if v0 == 1 {
		goto LABEL3
	}
	if v0 == 2 {
		goto LABEL4
	}
	goto LABEL5
LABEL3:
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02A.scr:JorganTalk01")
	goto LABEL5
LABEL4:
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02A.scr:MaidenTalk01")
	goto LABEL5
LABEL5:
	return
}
func Civvy1End() {
	ns.Frozen(obj39, false)
	ns.Wander(obj39)
	ns.SetDialog(obj39, ns.NORMAL, Civvy1Start, Civvy1End)
}
func Civvy2Start() {
	ns.Frozen(obj34, true)
	ns.CreatureIdle(obj34)
	ns.LookAtObject(obj34, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02A.scr:GrunbarTalk01")
}
func Civvy2End() {
	ns.Frozen(obj34, false)
	ns.Wander(obj34)
	ns.SetDialog(obj34, ns.NORMAL, Civvy2Start, Civvy2End)
}
func Civvy5Start() {
	ns.Frozen(obj37, true)
	ns.CreatureIdle(obj37)
	ns.LookAtObject(obj37, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02A.scr:AlbiTalk01")
}
func Civvy5End() {
	ns.Frozen(obj37, false)
	ns.Wander(obj37)
	ns.SetDialog(obj37, ns.NORMAL, Civvy5Start, Civvy5End)
}
func Civvy6Start() {
	ns.Frozen(obj38, true)
	ns.CreatureIdle(obj38)
	ns.LookAtObject(obj38, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02A.scr:DorianTalk01")
}
func Civvy6End() {
	ns.Frozen(obj38, false)
	ns.Wander(obj38)
	ns.SetDialog(obj38, ns.NORMAL, Civvy6Start, Civvy6End)
}
func ShariStart() {
	var v0 int
	ns.Frozen(obj40, true)
	ns.CreatureIdle(obj40)
	ns.LookAtObject(obj40, ns.GetHost())
	ivar111 = ns.Random(1, 2)
	v0 = ivar111
	if v0 == 1 {
		goto LABEL1
	}
	if v0 == 2 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02A.scr:MaidenTalk02")
	goto LABEL3
LABEL2:
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02A.scr:MaidenTalk03")
	goto LABEL3
LABEL3:
	return
}
func ShariEnd() {
	ns.Frozen(obj40, false)
	ns.Wander(obj40)
	ns.SetDialog(obj40, ns.NORMAL, ShariStart, ShariEnd)
}
func KaylaStart() {
	var v0 int
	ns.Frozen(obj41, true)
	ns.CreatureIdle(obj41)
	ns.LookAtObject(obj41, ns.GetHost())
	ivar112 = ns.Random(1, 2)
	v0 = ivar112
	if v0 == 1 {
		goto LABEL1
	}
	if v0 == 2 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02A.scr:MaidenTalk04")
	goto LABEL3
LABEL2:
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02A.scr:MaidenTalk05")
	goto LABEL3
LABEL3:
	return
}
func KaylaEnd() {
	ns.Frozen(obj41, false)
	ns.Wander(obj41)
	ns.SetDialog(obj41, ns.NORMAL, KaylaStart, KaylaEnd)
}
func JorganStart() {
	ns.Frozen(obj33, true)
	ns.CreatureIdle(obj33)
	ns.LookAtObject(obj33, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "Con03C.scr:ForemanF")
}
func JorganEnd() {
	ns.Frozen(obj33, false)
	ns.Wander(obj33)
	ns.SetDialog(obj33, ns.NORMAL, JorganStart, JorganEnd)
}
func FentonStart() {
	ns.Frozen(obj35, true)
	ns.CreatureIdle(obj35)
	ns.LookAtObject(obj35, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "War02a:NewTownsman3")
}
func FentonEnd() {
	ns.Frozen(obj35, false)
	ns.Wander(obj35)
	ns.SetDialog(obj35, ns.NORMAL, FentonStart, FentonEnd)
}
func KelvinStart() {
	ns.Frozen(obj36, true)
	ns.CreatureIdle(obj36)
	ns.LookAtObject(obj36, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "War02a:NewTownsman4")
}
func KelvinEnd() {
	ns.Frozen(obj36, false)
	ns.Wander(obj36)
	ns.SetDialog(obj36, ns.NORMAL, KelvinStart, KelvinEnd)
}
func EowynnEnterMax() {
	if !ns.IsCaller(obj39) {
		goto LABEL1
	}
	flag105 = true
LABEL1:
	return
}
func EowynnExitMax() {
	if !ns.IsCaller(obj39) {
		goto LABEL1
	}
	flag105 = false
LABEL1:
	return
}
func CreatureSetup() {
	ns.SetOwner(ns.GetHost(), obj18)
	ns.SetOwner(ns.GetHost(), obj37)
	ns.SetOwner(ns.GetHost(), obj38)
	ns.SetOwner(ns.GetHost(), obj34)
	ns.SetOwner(ns.GetHost(), obj33)
	ns.SetOwner(ns.GetHost(), obj35)
	ns.SetOwner(ns.GetHost(), obj36)
	ns.SetOwner(ns.GetHost(), obj29)
	ns.SetOwner(ns.GetHost(), obj27)
	ns.SetOwner(ns.GetHost(), obj28)
	ns.SetOwner(ns.GetHost(), obj22)
	ns.SetOwner(ns.GetHost(), obj23)
	ns.SetOwner(ns.GetHost(), obj24)
	ns.SetOwner(ns.GetHost(), obj30)
	ns.SetOwner(ns.GetHost(), obj7)
	ns.SetOwner(ns.GetHost(), obj31)
	ns.SetOwner(ns.GetHost(), obj32)
	ns.SetOwner(ns.GetHost(), obj39)
	ns.SetOwner(ns.GetHost(), obj40)
	ns.SetOwner(ns.GetHost(), obj41)
	ns.SetOwner(ns.GetHost(), obj42)
	ns.SetOwner(ns.GetHost(), obj43)
	ns.SetOwner(ns.GetHost(), obj44)
	ns.SetOwner(ns.GetHost(), obj45)
	ns.SetOwner(ns.GetHost(), obj46)
	ns.SetOwner(ns.GetHost(), obj47)
	ns.SetOwner(ns.GetHost(), obj48)
	ns.SetOwner(ns.GetHost(), obj49)
	ns.SetOwner(ns.GetHost(), obj50)
	ns.SetOwner(ns.GetHost(), obj51)
	ns.SetOwner(ns.GetHost(), obj8)
	ns.SetOwner(ns.GetHost(), obj9)
	ns.SetOwner(ns.GetHost(), obj10)
	ns.SetOwner(ns.GetHost(), obj11)
	ns.SetOwner(ns.GetHost(), obj12)
	ns.SetOwner(ns.GetHost(), obj13)
	ns.SetOwner(ns.GetHost(), obj14)
	ns.SetOwner(ns.GetHost(), obj52)
	ns.SetOwner(ns.GetHost(), obj53)
	ns.SetOwner(ns.GetHost(), obj54)
	ns.Wander(obj33)
	ns.Wander(obj34)
	ns.Wander(obj35)
	ns.Wander(obj36)
	ns.Wander(obj37)
	ns.Wander(obj38)
	ns.Wander(obj39)
	ns.Wander(obj40)
	ns.Wander(obj41)
}
func StateEngine() {
	if !(ns.CurrentHealth(obj26) <= 0) {
		goto LABEL1
	}
	if !(flag93 == false && flag98 == false) {
		goto LABEL1
	}
	if !ns.IsVisibleTo(obj25, ns.GetHost()) {
		goto LABEL1
	}
	ns.Chat(obj25, "Wiz02A.scr:MorganTalk01")
	ns.SetOwner(ns.GetHost(), obj25)
	ns.CreatureFollow(obj25, ns.GetHost())
	flag93 = true
LABEL1:
	if !flag101 {
		goto LABEL2
	}
	ns.LookAtObject(obj29, ns.GetHost())
LABEL2:
	gvar114 = ns.FrameTimer(2, StateEngine)
}
func EnemyRun() {
	r1 := ns.IsAttackedBy(ns.GetHost(), ns.GetCaller())
	if !r1 {
		goto LABEL1
	}
	ns.GoBackHome(ns.GetCaller())
LABEL1:
	return
}
func InitObjects() {
	obj18 = ns.Object("F6Guard4")
	obj19 = ns.Object("Sheriff")
	obj7 = ns.Object("Mlurgh")
	obj20 = ns.Object("Kincaid")
	obj22 = ns.Object("Mystic")
	obj21 = ns.Object("BrightBlades")
	obj23 = ns.Object("MiscVendor")
	obj24 = ns.Object("Phim")
	obj30 = ns.Object("Undertaker")
	obj31 = ns.Object("Priest")
	obj55 = ns.Object("EmptyRoomDoor")
	obj56 = ns.Object("GuardDoor1")
	obj57 = ns.Object("GuardDoor2")
	obj27 = ns.Object("Max")
	obj28 = ns.Object("Grillf")
	obj32 = ns.Object("Loproc")
	obj33 = ns.Object("Jorgan")
	obj34 = ns.Object("Grunbar")
	obj35 = ns.Object("Fenton")
	obj36 = ns.Object("Kelvin")
	obj37 = ns.Object("Albi")
	obj38 = ns.Object("Dorian")
	obj39 = ns.Object("Eowynn")
	obj40 = ns.Object("Shari")
	obj41 = ns.Object("Kayla")
	obj42 = ns.Object("Maiden1")
	obj43 = ns.Object("Maiden2")
	obj44 = ns.Object("Maiden5")
	obj45 = ns.Object("Maiden6")
	obj46 = ns.Object("Maiden7")
	obj47 = ns.Object("Maiden8")
	obj48 = ns.Object("Maiden9")
	obj49 = ns.Object("BarFly1")
	obj50 = ns.Object("BarFly2")
	obj51 = ns.Object("BarFly3")
	obj8 = ns.Object("MlurghPatron1")
	obj9 = ns.Object("MlurghPatron2")
	obj10 = ns.Object("MlurghPatron3")
	obj11 = ns.Object("MlurghPatron4")
	obj12 = ns.Object("MlurghPatron5")
	obj13 = ns.Object("MlurghPatron6")
	obj14 = ns.Object("MlurghPatron7")
	obj52 = ns.Object("TowerNPC")
	obj53 = ns.Object("TowerNPC2")
	obj54 = ns.Object("TowerMaiden")
	obj78 = ns.Object("WardenFace")
	obj72 = ns.Object("MaxReward")
	obj58 = ns.Object("MaxHomeDoor1")
	obj59 = ns.Object("MaxHomeDoor2")
	obj60 = ns.Object("GearRoomDoor")
	obj66 = ns.Object("RogueElevator")
	obj67 = ns.Object("PlayerElevatorTopTrigger")
	obj61 = ns.Object("FirstJailDoor")
	obj68 = ns.Object("JailDoorTrigger")
	obj79 = ns.Object("GrillfFace")
	obj29 = ns.Object("Warden")
	obj4 = ns.Object("AmbSound1")
	obj5 = ns.Object("AmbSound2")
	obj6 = ns.Object("AmbSound3")
	obj73 = ns.Object("DeadReward")
	obj74 = ns.Object("AliveReward")
	obj75 = ns.Object("RogueKey")
	obj65 = ns.Object("RogueDoor")
	obj25 = ns.Object("Morgan")
	obj76 = ns.Object("CryptChairMover")
	obj26 = ns.Object("MorganGuard1")
	obj69 = ns.Object("MorganCellTrigger")
	obj71 = ns.Object("WardenBackTrigger")
	obj64 = ns.Object("MorganCellDoor")
	obj77 = ns.Object("MorganKey")
	obj62 = ns.Object("MorganDoor1")
	obj63 = ns.Object("MorganDoor2")
	gvar80 = ns.ObjectGroup("RogueGears")
	gvar81 = ns.ObjectGroup("MorganSecret")
	gvar82 = ns.WallGroup("TroveWalls")
}
func SetTownDialog() {
	if ns.GetQuestStatus("MorganCaptured") != 1 {
		goto LABEL1
	}
	flag93 = true
	flag97 = true
	if ns.GetQuestStatus("MorganDead") == 1 {
		goto LABEL2
	}
	ns.ObjectOff(obj69)
	ns.ObjectOff(obj70)
	ns.ObjectOff(obj67)
	ns.ObjectOff(obj68)
	ns.ObjectOff(obj71)
	ns.LockDoor(obj64)
	ns.UnlockDoor(obj61)
	ns.UnlockDoor(obj65)
	ns.UnlockDoor(obj62)
	ns.UnlockDoor(obj63)
	ns.ObjectOn(obj66)
	ns.MoveObject(obj25, 3627, 5467)
LABEL2:
	goto LABEL3
LABEL1:
	ns.SetDialog(obj27, ns.NORMAL, MaxProblemStart, MaxProblemEnd)
	ns.SetDialog(obj29, ns.NORMAL, WardenGreetStart, WardenGreetEnd)
LABEL3:
	ns.SetDialog(obj31, ns.NORMAL, PriestStart, PriestEnd)
	ns.SetDialog(obj23, ns.NORMAL, GenericStart, GenericEnd)
	ns.SetDialog(obj30, ns.NORMAL, GenericStart, GenericEnd)
	ns.SetDialog(obj39, ns.NORMAL, Civvy1Start, Civvy1End)
	ns.SetDialog(obj34, ns.NORMAL, Civvy2Start, Civvy2End)
	ns.SetDialog(obj37, ns.NORMAL, Civvy5Start, Civvy5End)
	ns.SetDialog(obj38, ns.NORMAL, Civvy6Start, Civvy6End)
	ns.SetDialog(obj40, ns.NORMAL, ShariStart, ShariEnd)
	ns.SetDialog(obj41, ns.NORMAL, KaylaStart, KaylaEnd)
	ns.SetDialog(obj33, ns.NORMAL, JorganStart, JorganEnd)
	ns.SetDialog(obj35, ns.NORMAL, FentonStart, FentonEnd)
	ns.SetDialog(obj36, ns.NORMAL, KelvinStart, KelvinEnd)
	ns.StoryPic(obj27, "MaxPic")
	ns.StoryPic(obj30, "UndertakerPic")
	ns.StoryPic(obj38, "Townsman2Pic")
	ns.StoryPic(obj34, "MalePic7")
	ns.StoryPic(obj37, "MalePic1")
	ns.StoryPic(obj33, "Townsman3Pic")
	ns.StoryPic(obj35, "MalePic8")
	ns.StoryPic(obj36, "MalePic5")
	ns.StoryPic(obj29, "WardenPic")
	ns.StoryPic(obj31, "GalavaPriestPic")
	ns.StoryPic(obj39, "MaidenPic3")
	ns.StoryPic(obj40, "MaidenPic")
	ns.StoryPic(obj41, "MaidenPic2")
}
func LoprocSeesPlayer() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.Chat(obj32, "Wiz02A.scr:AppleManTalk01")
LABEL1:
	return
}
func MorganCapturePiece() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	if !(flag93 == true && flag107 != true) {
		goto LABEL1
	}
	if !(ns.CurrentHealth(obj25) > 0) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.WideScreen(true)
	ns.Frozen(ns.GetHost(), true)
	ns.CreatureIdle(ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj29)
	ns.SetDialog(obj29, ns.NORMAL, Warden1Start, Warden1End)
	ns.StartDialog(obj29, ns.GetHost())
LABEL1:
	return
}
func MapInitialize() {
	InitObjects()
	wp84 = ns.Waypoint("Inn1WP")
	wp85 = ns.Waypoint("Inn2WP")
	wp86 = ns.Waypoint("RogueElevatorWP")
	wp83 = ns.Waypoint("PlayerSounds")
	wp87 = ns.Waypoint("ExitTowerWP")
	wp88 = ns.Waypoint("CryptChairStartWP")
	wp89 = ns.Waypoint("CryptChairEndWP")
	gvar90 = ns.WaypointGroup("ToJailWPPath")
	gvar91 = ns.WaypointGroup("WardenGoBackWP")
	ns.ObjectOn(obj76)
	ns.ObjectGroupOff(gvar80)
	ns.ObjectOff(obj66)
	ns.ObjectOff(obj71)
	ns.WayPointGroupOff(gvar90)
	ns.WayPointGroupOff(gvar91)
	ns.LockDoor(obj56)
	ns.LockDoor(obj57)
	ns.LockDoor(obj58)
	ns.LockDoor(obj59)
	ns.LockDoor(obj60)
	ns.LockDoor(obj61)
	if ns.GetQuestStatus("MorganDead") != 1 {
		goto LABEL1
	}
	ns.Delete(obj25)
LABEL1:
	SecretsInit()
	SetTownDialog()
	ns.FrameTimer(1, CreatureSetup)
	StateEngine()
}
func MapEntry() {
	ns.UnBlind()
	ns.Music(17, 100)
}
func PlayerDeath() {
	ns.DeathScreen(2)
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
