package con02a

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	flag4   bool
	obj5    ns.ObjectID
	ivar6   int
	ivar7   int
	ivar8   int
	ivar9   int
	ivar10  int
	obj11   ns.ObjectID
	obj12   ns.ObjectID
	wp13    ns.WaypointID
	wp14    ns.WaypointID
	wp15    ns.WaypointID
	wp16    ns.WaypointID
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
	gvar30  int
	gvar31  int
	gvar32  int
	gvar33  int
	gvar34  int
	gvar35  int
	gvar36  int
	gvar37  int
	gvar38  int
	gvar39  int
	gvar40  int
	gvar41  int
	gvar42  int
	gvar43  int
	gvar44  int
	flag45  bool
	flag46  bool
	flag47  bool
	flag48  bool
	flag49  bool
	gvar50  int
	gvar51  int
	gvar52  int
	gvar53  int
	gvar54  int
	gvar55  int
	gvar56  int
	gvar57  int
	ivar58  int
	ivar59  int
	ivar60  int
	ivar61  int
	obj62   ns.ObjectID
	flag63  bool
	flag64  bool
	flag65  bool
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
	obj95   ns.ObjectID
	obj96   ns.ObjectID
	obj97   ns.ObjectID
	obj98   ns.ObjectID
	obj99   ns.ObjectID
	obj100  ns.ObjectID
	obj101  ns.ObjectID
	obj102  ns.ObjectID
	gvar103 ns.ObjectGroupID
	wp104   ns.WaypointID
	wp105   ns.WaypointID
	wp106   ns.WaypointID
	wp107   ns.WaypointID
	wp108   ns.WaypointID
	wp109   ns.WaypointID
	wp110   ns.WaypointID
	wp111   ns.WaypointID
	wp112   ns.WaypointID
	wp113   ns.WaypointID
	gvar114 ns.WaypointGroupID
	gvar115 ns.WaypointGroupID
	gvar116 ns.WaypointGroupID
	flag117 bool
	flag118 bool
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
	gvar134 ns.ObjectGroupID
	gvar135 ns.ObjectGroupID
	wp136   ns.WaypointID
	wp137   ns.WaypointID
	wp138   ns.WaypointID
	wp139   ns.WaypointID
	wp140   ns.WaypointID
	wp141   ns.WaypointID
	wp142   ns.WaypointID
	wp143   ns.WaypointID
	wp144   ns.WaypointID
	ivar145 int
	flag146 bool
	flag147 bool
	flag148 bool
	flag149 bool
	gvar150 int
	gvar151 int
	gvar152 int
	gvar153 int
	gvar154 int
	fvar155 float32
	fvar156 float32
	obj157  ns.ObjectID
	flag158 bool
	flag159 bool
	ivar160 int
	wp161   ns.WaypointID
	wp162   ns.WaypointID
	wp163   ns.WaypointID
	wp164   ns.WaypointID
	obj165  ns.ObjectID
	gvar166 ns.ObjectGroupID
	gvar167 ns.ObjectGroupID
	gvar168 ns.WallGroupID
	obj169  ns.ObjectID
	gvar170 int
	obj171  ns.ObjectID
	gvar172 int
	obj173  ns.ObjectID
	obj174  ns.ObjectID
	obj175  ns.ObjectID
	obj176  ns.ObjectID
	obj177  ns.ObjectID
	obj178  ns.ObjectID
	obj179  ns.ObjectID
	obj180  [2]ns.ObjectID
	gvar181 ns.ObjectGroupID
	gvar182 ns.ObjectGroupID
	gvar183 ns.ObjectGroupID
	gvar184 ns.ObjectGroupID
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
	wp196   ns.WaypointID
	wp197   ns.WaypointID
	wp198   ns.WaypointID
	wp199   ns.WaypointID
	wp200   [2]ns.WaypointID
	obj201  ns.ObjectID
	gvar202 int
	gvar203 int
	gvar204 int
	gvar205 int
	gvar206 int
	gvar207 int
	gvar208 int
	gvar209 int
	gvar210 int
	gvar211 int
	gvar212 int
	gvar213 int
	obj214  ns.ObjectID
	obj215  ns.ObjectID
	obj216  ns.ObjectID
	obj217  ns.ObjectID
	gvar218 ns.ObjectGroupID
	wp219   ns.WaypointID
	flag220 bool
	ivar221 int
)

func init() {
	flag4 = true
	ivar6 = 1
	ivar7 = 0
	ivar8 = 0
	ivar9 = 85
	ivar10 = 2
	gvar17 = 0
	gvar18 = 1
	gvar19 = 2
	gvar20 = 3
	gvar21 = 4
	gvar22 = 5
	gvar23 = 6
	gvar24 = 7
	gvar25 = 8
	gvar26 = 9
	gvar27 = 10
	gvar28 = 11
	gvar29 = 12
	gvar30 = 13
	gvar31 = 14
	gvar32 = 15
	gvar33 = 16
	gvar34 = 17
	gvar35 = 18
	gvar36 = 19
	gvar37 = 20
	gvar38 = 21
	gvar39 = 22
	gvar40 = 23
	gvar41 = 24
	gvar42 = 25
	gvar43 = 26
	gvar44 = 27
	flag45 = false
	flag46 = false
	flag47 = false
	flag48 = true
	flag49 = true
	gvar50 = gvar25
	gvar51 = gvar22
	gvar52 = gvar17
	gvar53 = gvar28
	gvar54 = gvar32
	gvar55 = gvar36
	gvar56 = gvar40
	gvar57 = gvar42
	ivar58 = 20
	ivar59 = 30
	ivar60 = 100
	ivar61 = 0
	flag63 = false
	flag64 = false
	flag65 = false
	ivar145 = 0
	flag146 = false
	flag147 = false
	flag148 = false
	flag149 = false
	gvar150 = 0
	gvar151 = 1
	gvar152 = gvar150
	gvar153 = 0
	gvar154 = 1
	fvar155 = 1650
	fvar156 = 3055
	flag117 = false
	flag158 = false
	flag159 = true
	ivar160 = 30
	flag118 = false
	gvar202 = 0
	gvar203 = 1
	gvar204 = 2
	gvar205 = 3
	gvar206 = 4
	gvar207 = 5
	gvar208 = gvar202
	gvar209 = gvar202
	gvar210 = gvar202
	gvar211 = gvar202
	gvar212 = gvar202
	gvar213 = gvar202
	flag220 = false
	ivar221 = 0
}
func ContestWon(a1 int) {
	ns.LookAtObject(obj11, ns.GetHost())
	flag64 = true
	gvar54 = gvar34
	ns.StartDialog(obj11, ns.GetHost())
	ns.UnlockDoor(obj77)
	ns.UnlockDoor(obj78)
}
func ContestLost(a1 int) {
	var v0 int
	ns.DestroyEveryChat()
	ivar61 += 1
	ns.LookAtObject(obj11, ns.GetHost())
	if ivar61 != 1 {
		goto LABEL1
	}
	gvar54 = gvar32
LABEL1:
	if !(ivar61 > 1) {
		goto LABEL2
	}
	ns.SetDialog(obj11, ns.NORMAL, ContestOfficialDialogStart, ContestOfficialDialogEnd)
	ns.SetShopkeeperText(obj67, "Con02:BarkeeperDefault")
	gvar54 = gvar35
LABEL2:
	v0 = a1
	if v0 == 8 {
		goto LABEL3
	}
	if v0 == 7 {
		goto LABEL4
	}
	if v0 == 6 {
		goto LABEL5
	}
	if v0 == 5 {
		goto LABEL6
	}
	if v0 == 4 {
		goto LABEL7
	}
	if v0 == 3 {
		goto LABEL8
	}
	if v0 == 2 {
		goto LABEL9
	}
	if v0 == 1 {
		goto LABEL10
	}
	if v0 == 0 {
		goto LABEL11
	}
	goto LABEL12
LABEL3:
	ns.Chat(obj102, "Con02a:Contest8of10")
	goto LABEL12
LABEL4:
	ns.Chat(obj102, "Con02a:Contest7of10")
	goto LABEL12
LABEL5:
	ns.Chat(obj102, "Con02a:Contest6of10")
	goto LABEL12
LABEL6:
	ns.Chat(obj102, "Con02a:Contest5of10")
	goto LABEL12
LABEL7:
	ns.Chat(obj102, "Con02a:Contest4of10")
	goto LABEL12
LABEL8:
	ns.Chat(obj102, "Con02a:Contest3of10")
	goto LABEL12
LABEL9:
	ns.Chat(obj102, "Con02a:Contest2of10")
	goto LABEL12
LABEL10:
	ns.Chat(obj102, "Con02a:Contest1of10")
	goto LABEL12
LABEL11:
	ns.Chat(obj102, "Con02a:Contest0of10")
	goto LABEL12
LABEL12:
	ns.UnlockDoor(obj77)
	ns.UnlockDoor(obj78)
}
func WinContest() {
	gvar210 = gvar207
}
func ContestEngine() {
	var (
		v0 int
		v1 int
		v2 float32
		v3 float32
		v4 int
	)
	v1 = ns.CurrentHealth(obj5)
	if flag4 {
		goto LABEL1
	}
	v0 = ns.Random(0, 2)
	for {
		if v0 != ivar6 {
			goto LABEL2
		}
		v0 = ns.Random(0, 2)
	}
LABEL2:
	ivar6 = v0
	if !(v1 > 0) {
		goto LABEL4
	}
	ns.Delete(obj5)
	ivar8 += 1
	goto LABEL5
LABEL4:
	ivar7 += 1
	ivar9 = ivar9 - ivar10
LABEL5:
	goto LABEL6
LABEL1:
	v0 = 1
	flag4 = false
LABEL6:
	if !(ivar8+ivar7 < 10) {
		goto LABEL7
	}
	v4 = v0
	if v4 == 0 {
		goto LABEL8
	}
	if v4 == 1 {
		goto LABEL9
	}
	if v4 == 2 {
		goto LABEL10
	}
	goto LABEL11
LABEL8:
	v2 = ns.GetWaypointX(wp13)
	v3 = ns.GetWaypointY(wp13)
	ns.Effect(ns.SMOKE_BLAST, v2, v3, 0, 0)
	obj5 = ns.CreateObject("TargetBarrel1", wp13)
	ns.AudioEvent(ns.BlinkCast, wp16)
	goto LABEL11
LABEL9:
	v2 = ns.GetWaypointX(wp14)
	v3 = ns.GetWaypointY(wp14)
	ns.Effect(ns.SMOKE_BLAST, v2, v3, 0, 0)
	obj5 = ns.CreateObject("TargetBarrel1", wp14)
	ns.AudioEvent(ns.BlinkCast, wp16)
	goto LABEL11
LABEL10:
	v2 = ns.GetWaypointX(wp15)
	v3 = ns.GetWaypointY(wp15)
	ns.Effect(ns.SMOKE_BLAST, v2, v3, 0, 0)
	obj5 = ns.CreateObject("TargetBarrel1", wp15)
	ns.AudioEvent(ns.BlinkCast, wp16)
	goto LABEL11
LABEL11:
	ns.LookAtObject(obj11, obj5)
	ns.FrameTimer(ivar9, ContestEngine)
	goto LABEL12
LABEL7:
	if !(ivar7 >= 9 && ivar8 <= 1) {
		goto LABEL13
	}
	ContestWon(ivar7)
	WinContest()
	ivar7 = 0
	ivar8 = 0
	ivar6 = 1
	ivar9 = 60
	flag4 = true
	goto LABEL12
LABEL13:
	ContestLost(ivar7)
	ivar7 = 0
	ivar8 = 0
	ivar6 = 1
	ivar9 = 85
	flag4 = true
LABEL12:
	return
}
func OfficialSEG4() {
	ns.LookAtObject(obj11, ns.GetHost())
	ns.Chat(obj11, "Con02a:OfficialSEG4")
	ContestEngine()
}
func OfficialSEG3() {
	ns.LookAtObject(obj11, ns.GetHost())
	ns.Chat(obj11, "Con02a:OfficialSEG3")
	ns.FrameTimer(30, OfficialSEG4)
}
func OfficialSEG2() {
	ns.LookAtObject(obj11, ns.GetHost())
	ns.ObjectOff(obj12)
	ns.Chat(obj11, "Con02a:OfficialSEG2")
	ns.FrameTimer(30, OfficialSEG3)
}
func GatekeeperDialogStart() {
	var v0 int
	ns.LookAtObject(obj68, ns.GetHost())
	v0 = gvar56
	if v0 == gvar40 {
		goto LABEL1
	}
	if v0 == gvar41 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.TellStory(ns.SwordsmanHurt, "Con02a:GatekeeperWaiting")
	goto LABEL3
LABEL2:
	ns.TellStory(ns.SwordsmanHurt, "Con02a:GatekeeperGuarding")
	goto LABEL3
LABEL3:
	return
}
func GatekeeperDialogEnd() {
}
func BarkeeperDialogStart() {
	var v0 int
	ns.LookAtObject(obj67, ns.GetHost())
	v0 = gvar55
	if v0 == gvar36 {
		goto LABEL1
	}
	if v0 == gvar37 {
		goto LABEL2
	}
	if v0 == gvar38 {
		goto LABEL3
	}
	if v0 == gvar39 {
		goto LABEL4
	}
	goto LABEL5
LABEL1:
	ns.TellStory(ns.SwordsmanHurt, "Con02a:BarkeeperGreet")
	goto LABEL5
LABEL2:
	ns.TellStory(ns.SwordsmanHurt, "Con02a:BarkeeperQuest")
	goto LABEL5
LABEL3:
	ns.TellStory(ns.SwordsmanHurt, "Con02a:BarkeeperAfterQuest")
	goto LABEL5
LABEL4:
	ns.TellStory(ns.SwordsmanHurt, "Con02a:RandomSay")
	goto LABEL5
LABEL5:
	return
}
func BarkeeperDialogEnd() {
	if gvar55 != gvar38 {
		goto LABEL1
	}
	gvar55 = gvar39
LABEL1:
	return
}
func ContestGuardDialogStart() {
	var v0 int
	v0 = gvar51
	if v0 == gvar22 {
		goto LABEL1
	}
	if v0 == gvar23 {
		goto LABEL2
	}
	if v0 == gvar24 {
		goto LABEL3
	}
	goto LABEL4
LABEL1:
	ns.LookAtObject(obj69, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con02a:ContestGuard")
	goto LABEL4
LABEL2:
	ns.LookAtObject(obj69, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con02a:EnterContest")
	flag63 = true
	ns.UnlockDoor(obj77)
	ns.UnlockDoor(obj78)
	goto LABEL4
LABEL3:
	ns.LookAtObject(obj69, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con02a:RandomSay")
	goto LABEL4
LABEL4:
	return
}
func ContestGuardDialogEnd() {
	var (
		v0 int
		v1 int
		v2 int
		v3 int
		v4 int
		v5 bool
		v6 ns.ObjectID
		v7 int
		v8 int
	)
	v0 = 0
	v1 = 0
	v2 = 1
	v3 = 2
	v4 = 0
	v5 = false
	v6 = ns.GetLastItem(ns.GetHost())
	for {
		if v6 == 0 {
			goto LABEL1
		}
		if !(ns.HasClass(v6, ns.WEAPON) && ns.HasSubclass(v6, "BOW")) {
			goto LABEL2
		}
		v5 = true
		goto LABEL1
	LABEL2:
		v6 = ns.GetPreviousItem(v6)
	}
LABEL1:
	v8 = gvar51
	if v8 == gvar22 {
		goto LABEL4
	}
	if v8 == gvar23 {
		goto LABEL5
	}
	if v8 == gvar24 {
		goto LABEL6
	}
	goto LABEL7
LABEL4:
	v1 = ns.GetAnswer(obj69)
	v7 = v1
	if v7 == v2 {
		goto LABEL8
	}
	if v7 == v3 {
		goto LABEL9
	}
	if v7 == v4 {
		goto LABEL10
	}
	goto LABEL11
LABEL8:
	v0 = ns.GetGold(ns.GetHost())
	if !(v0 < ivar58) {
		goto LABEL12
	}
	ns.SetDialog(obj69, ns.NORMAL, NullDialogStart, ResetContestGuardDialogChoice)
	ns.TellStory(ns.SwordsmanHurt, "Con02a:NotEnoughGold")
	goto LABEL13
LABEL12:
	if v5 {
		goto LABEL14
	}
	ns.SetDialog(obj69, ns.NORMAL, NullDialogStart, ResetContestGuardDialogChoice)
	ns.TellStory(ns.SwordsmanHurt, "Con02a:NoBow")
	goto LABEL13
LABEL14:
	ns.ChangeGold(ns.GetHost(), -ivar58)
	gvar51 = gvar23
	ns.SetDialog(obj69, ns.NORMAL, ContestGuardDialogStart, ContestGuardDialogEnd)
	ContestGuardDialogStart()
LABEL13:
	goto LABEL11
LABEL9:
	goto LABEL11
LABEL10:
	goto LABEL11
LABEL11:
	goto LABEL7
LABEL5:
	ns.CancelDialog(obj69)
	goto LABEL7
LABEL6:
	goto LABEL7
LABEL7:
	return
}
func MayorsGuardDialogStart() {
	var v0 int
	v0 = gvar50
	if v0 == gvar25 {
		goto LABEL1
	}
	if v0 == gvar26 {
		goto LABEL2
	}
	if v0 == gvar27 {
		goto LABEL3
	}
	goto LABEL4
LABEL1:
	ns.LookAtObject(obj70, ns.GetHost())
	gvar55 = gvar37
	ns.TellStory(ns.SwordsmanHurt, "Con02a:MayorsGuard")
	goto LABEL4
LABEL2:
	ns.DestroyEveryChat()
	ns.LookAtObject(obj70, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con02a:MeetTheMayor")
	if !ns.IsLocked(obj75) {
		goto LABEL5
	}
	ns.UnlockDoor(obj75)
	ns.UnlockDoor(obj76)
LABEL5:
	goto LABEL4
LABEL3:
	ns.LookAtObject(obj70, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War03b:MayorsGuardIntro")
LABEL4:
	return
}
func MayorsGuardDialogEnd() {
}
func MayorDialogStart() {
	var v0 int
	ns.DestroyEveryChat()
	v0 = gvar53
	if v0 == gvar28 {
		goto LABEL1
	}
	if v0 == gvar29 {
		goto LABEL2
	}
	if v0 == gvar30 {
		goto LABEL3
	}
	if v0 == gvar31 {
		goto LABEL4
	}
	goto LABEL5
LABEL1:
	ns.LookAtObject(obj72, ns.GetHost())
	ns.DestroyEveryChat()
	ns.TellStory(ns.SwordsmanHurt, "Con02a:MayorGreeting")
	goto LABEL5
LABEL2:
	ns.LookAtObject(obj72, ns.GetHost())
	if flag117 {
		goto LABEL6
	}
	ns.TellStory(ns.SwordsmanHurt, "Con02a:MayorProd")
	goto LABEL7
LABEL6:
	ns.TellStory(ns.SwordsmanHurt, "Con02a:MayorFree")
LABEL7:
	goto LABEL5
LABEL3:
	ns.Frozen(obj72, true)
	ns.LookAtObject(obj72, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con02a:MayorsReward")
	ns.SetShopkeeperText(obj67, "Con02a:BarkeeperAfterQuest")
	goto LABEL5
LABEL4:
	ns.LookAtObject(obj72, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con02a:MayorsGuardIdle")
	goto LABEL5
LABEL5:
	return
}
func MayorDialogEnd() {
	var v0 int
	v0 = gvar53
	if v0 == gvar28 {
		goto LABEL1
	}
	if v0 == gvar29 {
		goto LABEL2
	}
	if v0 == gvar30 {
		goto LABEL3
	}
	if v0 == gvar31 {
		goto LABEL4
	}
	goto LABEL5
LABEL1:
	gvar53 = gvar29
	goto LABEL5
LABEL2:
	if !flag117 {
		goto LABEL6
	}
	ns.WayPointGroupOff(gvar116)
	ns.WayPointGroupOn(gvar115)
	ns.UnlockDoor(obj88)
	ns.CreatureGuard(obj72, ns.GetWaypointX(wp110), ns.GetWaypointY(wp110), ns.GetWaypointX(ns.Waypoint("MayorRewardLookWP")), ns.GetWaypointY(ns.Waypoint("MayorRewardLookWP")), 150)
	gvar53 = gvar30
	ns.CancelDialog(obj66)
	ns.FrameTimer(15, AwardMayorExperience)
LABEL6:
	goto LABEL5
LABEL3:
	ns.Frozen(obj72, false)
	ns.JournalEdit(ns.GetHost(), "BanishSpiders", 4)
	ns.PrintToAll("Con02a:ObjectiveComplete")
	ns.CancelDialog(obj68)
	gvar55 = gvar38
	SaveMayor()
	HelpMayor()
	ns.CancelDialog(obj72)
	gvar50 = gvar27
	ns.Pickup(ns.GetHost(), obj89)
	ns.Move(obj72, wp111)
	ns.ObjectGroupOn(gvar103)
	if !flag118 {
		goto LABEL7
	}
	ns.PrintToAll("Gate triggers are enabled.")
LABEL7:
	goto LABEL5
LABEL4:
	goto LABEL5
LABEL5:
	return
}
func AldwinDialogStart() {
	var v0 int
	v0 = gvar52
	if v0 == gvar17 {
		goto LABEL1
	}
	if v0 == gvar18 {
		goto LABEL2
	}
	if v0 == gvar19 {
		goto LABEL3
	}
	if v0 == gvar20 {
		goto LABEL4
	}
	if v0 == gvar21 {
		goto LABEL5
	}
	goto LABEL6
LABEL1:
	ns.LookAtObject(obj71, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con02a:AldwinGreeting")
	goto LABEL6
LABEL2:
	ns.LookAtObject(obj71, ns.GetHost())
	gvar52 = gvar19
	ns.TellStory(ns.SwordsmanHurt, "Con02a:BecomeConjurer2")
	goto LABEL6
LABEL3:
	ns.LookAtObject(obj71, ns.GetHost())
	gvar52 = gvar20
	ns.TellStory(ns.SwordsmanHurt, "Con02a:BecomeConjurer3")
	goto LABEL6
LABEL4:
	ns.LookAtObject(obj71, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con02a:AldwinProd")
	goto LABEL6
LABEL5:
	ns.LookAtObject(obj71, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con02a:AldwinImp")
	goto LABEL6
LABEL6:
	return
}
func AldwinDialogEnd() {
	var (
		v0 int
		v1 int
		v2 int
		v3 int
		v4 int
		v5 int
		v6 int
	)
	v0 = 0
	v1 = 0
	v2 = 1
	v3 = 2
	v4 = 0
	v6 = gvar52
	if v6 == gvar17 {
		goto LABEL1
	}
	if v6 == gvar18 {
		goto LABEL2
	}
	if v6 == gvar19 {
		goto LABEL3
	}
	if v6 == gvar20 {
		goto LABEL4
	}
	if v6 == gvar21 {
		goto LABEL5
	}
	goto LABEL6
LABEL1:
	v1 = ns.GetAnswer(obj71)
	v5 = v1
	if v5 == v2 {
		goto LABEL7
	}
	if v5 == v3 {
		goto LABEL8
	}
	if v5 == v4 {
		goto LABEL9
	}
	goto LABEL10
LABEL7:
	v0 = ns.GetGold(ns.GetHost())
	if !(v0 < ivar59) {
		goto LABEL11
	}
	ns.SetDialog(obj71, ns.NORMAL, NullDialogStart, ResetAldwinDialogChoice)
	ns.TellStory(ns.SwordsmanHurt, "Con02a:AldwinPoor")
	goto LABEL12
LABEL11:
	if !flag48 {
		goto LABEL13
	}
	flag48 = false
	ns.JournalEdit(ns.GetHost(), "LocateAldwyn", 4)
	ns.PrintToAll("Con02a:ObjectiveComplete")
LABEL13:
	ns.ChangeGold(ns.GetHost(), -ivar59)
	ns.SetDialog(obj71, ns.NEXT, AldwinDialogStart, AldwinDialogEnd)
	gvar52 = gvar18
	flag45 = true
	gvar50 = gvar26
	ns.Pickup(ns.GetHost(), obj80)
	ns.ObjectOn(obj84)
	ns.ObjectOn(obj98)
	ns.TellStory(ns.SwordsmanHurt, "Con02a:BecomeConjurer1")
	ns.JournalEntry(ns.GetHost(), "BanishSpiders", 2)
	ns.PrintToAll("Con01a:NewJournalEntry")
LABEL12:
	goto LABEL10
LABEL8:
	goto LABEL10
LABEL9:
	goto LABEL10
LABEL10:
	goto LABEL6
LABEL2:
	gvar52 = gvar18
	AldwinDialogStart()
	ns.Pickup(ns.GetHost(), obj82)
	goto LABEL6
LABEL3:
	ns.SetDialog(obj71, ns.NORMAL, AldwinDialogStart, AldwinDialogEnd)
	ns.GiveXp(ns.GetHost(), 200)
	AldwinDialogStart()
	goto LABEL6
LABEL4:
	ns.CreatureGuard(obj71, ns.GetWaypointX(wp104), ns.GetWaypointY(wp104), ns.GetObjectX(ns.Object("AldwynBookcase")), ns.GetObjectY(ns.Object("AldwynBookcase")), 0)
	goto LABEL6
LABEL5:
	ns.JournalEntry(ns.GetHost(), "LocateMineForeman", 2)
	ns.PrintToAll("Con01a:NewJournalEntry")
	ns.SetDialog(obj68, ns.NORMAL, GatekeeperDialogStart, GatekeeperDialogEnd)
	ns.Frozen(ns.GetHost(), false)
	ns.UnlockDoor(obj90)
	ns.UnlockDoor(obj91)
	ns.CancelDialog(obj71)
	ns.FrameTimer(30, ImpExit)
	goto LABEL6
LABEL6:
	return
}
func CheckForAldwyn() {
	if !ns.IsCaller(obj71) {
		goto LABEL1
	}
	ns.CreatureGuard(obj71, ns.GetWaypointX(wp104), ns.GetWaypointY(wp104), 5117, 1091, 0)
LABEL1:
	return
}
func ContestOfficialDialogStart() {
	var v0 int
	ns.DestroyEveryChat()
	v0 = gvar54
	if v0 == gvar32 {
		goto LABEL1
	}
	if v0 == gvar33 {
		goto LABEL2
	}
	if v0 == gvar34 {
		goto LABEL3
	}
	if v0 == gvar35 {
		goto LABEL4
	}
	goto LABEL5
LABEL1:
	ns.LookAtObject(obj11, ns.GetHost())
	if !(ivar61 > 0) {
		goto LABEL6
	}
	ns.SetDialog(obj11, ns.YESNO, NullDialogStart, ContestOfficialDialogEnd)
	ns.TellStory(ns.SwordsmanHurt, "Con02a:RetryContestChoice")
	goto LABEL7
LABEL6:
	ns.TellStory(ns.SwordsmanHurt, "Con02a:ContestGreeting")
LABEL7:
	goto LABEL5
LABEL2:
	ns.LookAtObject(obj11, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con02a:ContestOfficialWaiting")
	goto LABEL5
LABEL3:
	ns.LookAtObject(obj11, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con02a:AwardContestPrize")
	ns.SetShopkeeperText(obj67, "Con02:BarkeeperDefault")
	goto LABEL5
LABEL4:
	ns.LookAtObject(obj11, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con02a:NoMoreContestTries")
	goto LABEL5
LABEL5:
	return
}
func ContestOfficialDialogEnd() {
	var (
		v0 int
		v1 int
		v2 int
		v3 int
	)
	v0 = 0
	v1 = 0
	v2 = 1
	v0 = ns.GetAnswer(obj11)
	v1 = ns.GetGold(ns.GetHost())
	v3 = gvar54
	if v3 == gvar32 {
		goto LABEL1
	}
	if v3 == gvar33 {
		goto LABEL2
	}
	if v3 == gvar34 {
		goto LABEL3
	}
	if v3 == gvar35 {
		goto LABEL4
	}
	goto LABEL5
LABEL1:
	if !(ivar61 > 0) {
		goto LABEL6
	}
	if v0 != v2 {
		goto LABEL7
	}
	if !(v1 >= ivar58) {
		goto LABEL8
	}
	ns.ChangeGold(ns.GetHost(), -ivar58)
	ns.SetDialog(obj11, ns.NORMAL, ContestOfficialDialogStart, EnableFiringLineTrigger)
	ns.TellStory(ns.SwordsmanHurt, "Con02a:RetryContest")
	ns.Pickup(ns.GetHost(), obj86)
	ns.PrintToAll("Con02a:QuiverAddedToInventory")
	ns.LockDoor(obj77)
	ns.LockDoor(obj78)
	gvar54 = gvar33
	goto LABEL9
LABEL8:
	ns.SetDialog(obj11, ns.NORMAL, NullDialogStart, ResetContestOfficialDialogChoice)
	ns.TellStory(ns.SwordsmanHurt, "Con02a:NotEnoughGold")
LABEL9:
	goto LABEL10
LABEL7:
	ns.SetDialog(obj11, "FALSE", ContestOfficialDialogStart, ContestOfficialDialogEnd)
LABEL10:
	goto LABEL11
LABEL6:
	ns.Pickup(ns.GetHost(), obj85)
	ns.PrintToAll("Con02a:QuiverAddedToInventory")
	EnableFiringLineTrigger()
	ns.LockDoor(obj77)
	ns.LockDoor(obj78)
	gvar54 = gvar33
LABEL11:
	goto LABEL5
LABEL2:
	goto LABEL5
LABEL3:
	ns.Pickup(ns.GetHost(), obj87)
	gvar54 = gvar35
	ns.FrameTimer(15, AwardContestExperience)
	goto LABEL5
LABEL4:
	goto LABEL5
LABEL5:
	return
}
func BridgeGuardDialogStart() {
	var (
		v0 ns.ObjectID
		v1 int
	)
	v0 = ns.GetLastItem(ns.GetHost())
	for {
		if v0 == 0 {
			goto LABEL1
		}
		if v0 != obj94 {
			goto LABEL2
		}
		ns.Delete(v0)
		flag65 = true
		goto LABEL1
	LABEL2:
		v0 = ns.GetPreviousItem(v0)
	}
LABEL1:
	if !(flag65 == true && flag47 == false) {
		goto LABEL4
	}
	gvar57 = gvar43
LABEL4:
	v1 = gvar57
	if v1 == gvar42 {
		goto LABEL5
	}
	if v1 == gvar43 {
		goto LABEL6
	}
	if v1 == gvar44 {
		goto LABEL7
	}
	goto LABEL8
LABEL5:
	ns.LookAtObject(obj92, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con02a:BridgeGuardGuarding")
	goto LABEL8
LABEL6:
	ns.LookAtObject(obj92, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con02a:BridgeGuardReward")
	ns.Drop(ns.GetHost(), obj94)
	ns.FrameTimer(1, PickupBoots)
	flag47 = true
	goto LABEL8
LABEL7:
	ns.LookAtObject(obj92, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con02a:RandomSay")
	goto LABEL8
LABEL8:
	return
}
func BridgeGuardDialogEnd() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
		v4 int
	)
	v4 = gvar57
	if v4 == gvar42 {
		goto LABEL1
	}
	if v4 == gvar43 {
		goto LABEL2
	}
	if v4 == gvar44 {
		goto LABEL3
	}
	goto LABEL4
LABEL1:
	if !flag49 {
		goto LABEL5
	}
	flag49 = false
	ns.JournalEntry(ns.GetHost(), "ReturnBridgeGuardsBoots", 2)
	ns.PrintToAll("Con01a:NewJournalEntry")
LABEL5:
	goto LABEL4
LABEL2:
	ns.CancelDialog(obj92)
	v0 = ns.GetObjectX(obj92)
	v1 = ns.GetObjectY(obj92)
	v2 = 3116
	v3 = 3394
	ns.MoveObject(obj92, v2, v3)
	ns.MoveObject(obj93, v0, v1)
	ns.CreatureGuard(obj92, v2, v3, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	ns.CreatureGuard(obj93, v0, v1, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	ns.LookAtObject(obj93, ns.GetHost())
	ns.JournalEdit(ns.GetHost(), "ReturnBridgeGuardsBoots", 4)
	ns.PrintToAll("Con02a:ObjectiveComplete")
	ns.FrameTimer(15, AwardBootQuestExperience)
	goto LABEL4
LABEL3:
	goto LABEL4
LABEL4:
	return
}
func ConManDialogStart() {
	ns.TellStory(ns.SwordsmanHurt, "Con02a:ConManSalesPitch")
}
func ConManDialogEnd() {
	var (
		v0 int
		v1 int
		v2 int
	)
	v0 = 0
	v1 = 1
	v2 = 0
	v0 = ns.GetAnswer(obj96)
	v2 = ns.GetGold(ns.GetHost())
	if !ns.HasItem(ns.GetHost(), obj97) {
		goto LABEL1
	}
	ns.CancelDialog(obj96)
	return
LABEL1:
	if v0 != v1 {
		goto LABEL2
	}
	if !(v2 < ivar60) {
		goto LABEL3
	}
	ns.SetDialog(obj96, ns.NORMAL, NullDialogStart, ResetConManDialogChoice)
	ns.TellStory(ns.SwordsmanHurt, "Con02a:ConManNotEnoughGold")
	goto LABEL2
LABEL3:
	ns.Pickup(ns.GetHost(), obj97)
	ns.PrintToAll("Con02a:BowAddedToInventory")
	ns.ChangeGold(ns.GetHost(), -ivar60)
	ns.SetDialog(obj96, ns.NORMAL, NullDialogStart, ConManDialogEnd)
	ns.TellStory(ns.SwordsmanHurt, "Con02a:ConManSale")
	SlipIntoShadows()
LABEL2:
	return
}
func Maiden1DialogStart() {
	ns.LookAtObject(obj99, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con02a:BarMaiden")
}
func Maiden1DialogEnd() {
}
func SpiderEngine() {
	var v0 int
	if !flag118 {
		goto LABEL1
	}
	ns.PrintToAll("SpiderEngineRunning")
LABEL1:
	v0 = gvar152
	if v0 == gvar150 {
		goto LABEL2
	}
	if v0 == gvar151 {
		goto LABEL3
	}
	goto LABEL4
LABEL2:
	if !(ns.IsOwnedBy(obj165, ns.GetHost()) || ns.CurrentHealth(obj165) <= 0) {
		goto LABEL5
	}
	obj157 = ns.CreateObject("Spider", wp161)
	ns.SetCallback(obj157, 5, TransientMonsterDie)
	ns.LookAtObject(obj157, ns.GetHost())
	gvar152 = gvar151
	ns.FrameTimerWithArg(1, gvar151, SetToGuard)
	ns.FrameTimer(30, SpiderEngine)
	goto LABEL6
LABEL5:
	ns.FrameTimer(15, SpiderEngine)
LABEL6:
	goto LABEL4
LABEL3:
	if !(ns.CurrentHealth(obj157) > 0 || ns.CurrentHealth(obj165) > 0) {
		goto LABEL7
	}
	ns.FrameTimer(15, SpiderEngine)
	goto LABEL8
LABEL7:
	flag117 = true
	SpidersFinishedPatch()
LABEL8:
	goto LABEL4
LABEL4:
	return
}
func HelpMayor() {
	gvar208 = gvar204
	gvar209 = gvar204
	gvar210 = gvar204
	gvar212 = gvar204
	gvar213 = gvar204
}
func SaveMayor() {
	flag147 = true
}
func ImpEnter() {
	if !flag118 {
		goto LABEL1
	}
	ns.PrintToAll("ImpEnter()")
LABEL1:
	ns.ObjectGroupOff(gvar103)
	ns.Frozen(ns.GetHost(), true)
	ns.CreatureIdle(ns.GetHost())
	obj62 = ns.CreateObject("Imp", wp105)
	ns.SetOwner(ns.GetHost(), obj62)
	ns.Enchant(obj62, ns.ENCHANT_INVULNERABLE, 0)
	ns.FrameTimer(1, ImpEnterSEG2)
}
func AldwinWaiting() {
	if !ns.IsTalking() {
		goto LABEL1
	}
	ns.FrameTimer(1, AldwinWaiting)
	goto LABEL2
LABEL1:
	ns.StartDialog(obj71, ns.GetHost())
LABEL2:
	return
}
func InitializeDialogs() {
	ns.StoryPic(obj68, "Townsman4Pic")
	ns.StoryPic(obj69, "Townsman2Pic")
	ns.StoryPic(obj11, "MalePic10")
	ns.StoryPic(obj70, "Warrior3Pic")
	ns.StoryPic(obj72, "TheogrinPic")
	ns.StoryPic(obj71, "AldwynPic")
	ns.StoryPic(obj92, "Townsman3Pic")
	ns.StoryPic(obj96, "MorganPic")
	ns.StoryPic(obj99, "MaidenPic")
	ns.SetDialog(obj68, ns.NORMAL, GatekeeperDialogStart, GatekeeperDialogEnd)
	ns.SetDialog(obj69, ns.YESNO, ContestGuardDialogStart, ContestGuardDialogEnd)
	ns.SetDialog(obj70, ns.NORMAL, MayorsGuardDialogStart, MayorsGuardDialogEnd)
	ns.SetDialog(obj72, ns.NORMAL, MayorDialogStart, MayorDialogEnd)
	ns.SetDialog(obj71, ns.YESNO, AldwinDialogStart, AldwinDialogEnd)
	ns.SetDialog(obj11, ns.NORMAL, ContestOfficialDialogStart, ContestOfficialDialogEnd)
	ns.SetDialog(obj92, ns.NORMAL, BridgeGuardDialogStart, BridgeGuardDialogEnd)
	ns.SetDialog(obj96, ns.YESNO, ConManDialogStart, ConManDialogEnd)
	ns.SetDialog(obj99, ns.NORMAL, Maiden1DialogStart, Maiden1DialogEnd)
}
func NullDialogStart() {
}
func AwardBootQuestExperience() {
	ns.MoveWaypoint(wp113, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.FlagDrop, wp113)
	ns.GiveXp(ns.GetHost(), 200)
}
func AwardContestExperience() {
	ns.GiveXp(ns.GetHost(), 500)
}
func AwardMayorExperience() {
	ns.GiveXp(ns.GetHost(), 1500)
}
func PickupBoots() {
	ns.Pickup(obj92, obj94)
	ns.Pickup(ns.GetHost(), obj95)
}
func ResetContestGuardDialogChoice() {
	ns.SetDialog(obj69, ns.YESNO, ContestGuardDialogStart, ContestGuardDialogEnd)
}
func ResetContestOfficialDialogChoice() {
	ns.SetDialog(obj11, ns.YESNO, ContestOfficialDialogStart, ContestOfficialDialogEnd)
}
func EnableFiringLineTrigger() {
	ns.ObjectOn(obj12)
	ns.SetDialog(obj11, ns.NORMAL, ContestOfficialDialogStart, ContestOfficialDialogEnd)
}
func ImpToAldwin() {
	var (
		v0 float32
		v1 float32
	)
	v0 = ns.GetWaypointX(wp106)
	v1 = ns.GetWaypointY(wp106)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	ns.Delete(obj62)
	ns.MoveObject(obj71, v0, v1)
	gvar52 = gvar21
	ns.LookAtObject(obj71, ns.GetHost())
	ns.AudioEvent(ns.ImpRecognize, wp106)
	ns.CreatureGuard(obj71, ns.GetObjectX(obj71), ns.GetObjectY(obj71), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	AldwinWaiting()
}
func ImpEnterSEG2() {
	ns.Move(obj62, wp106)
	ns.FrameTimer(75, ImpToAldwin)
}
func DestroyImp() {
	var (
		v0 float32
		v1 float32
	)
	v0 = ns.GetObjectX(obj62)
	v1 = ns.GetObjectY(obj62)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	ns.Delete(obj62)
}
func ImpExitSEG2() {
	ns.Move(obj62, wp107)
	ns.FrameTimer(75, DestroyImp)
}
func ImpExit() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
	)
	v0 = ns.GetObjectX(obj71)
	v1 = ns.GetObjectY(obj71)
	v2 = ns.GetWaypointX(wp109)
	v3 = ns.GetWaypointY(wp109)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	ns.MoveObject(obj71, v2, v3)
	ns.ObjectOff(obj71)
	obj62 = ns.CreateObject("Imp", wp106)
	ns.AudioEvent(ns.ImpRecognize, wp106)
	ns.SetOwner(ns.GetHost(), obj62)
	ns.Enchant(obj62, ns.ENCHANT_INVULNERABLE, 0)
	ns.WayPointGroupOff(gvar114)
	ns.FrameTimer(1, ImpExitSEG2)
}
func ResetAldwinDialogChoice() {
	ns.SetDialog(obj71, ns.YESNO, AldwinDialogStart, AldwinDialogEnd)
}
func SpidersFinishedPatch() {
	gvar53 = gvar29
}
func SlipIntoShadows() {
	if ns.IsVisibleTo(ns.GetHost(), obj96) {
		goto LABEL1
	}
	ns.MoveObject(obj96, ns.GetWaypointX(ns.Waypoint("ConManEscapeWP")), ns.GetWaypointY(ns.Waypoint("ConManEscapeWP")))
	ns.ObjectOff(obj96)
	goto LABEL2
LABEL1:
	ns.FrameTimer(15, SlipIntoShadows)
LABEL2:
	return
}
func ResetConManDialogChoice() {
	ns.SetDialog(obj96, ns.YESNO, ConManDialogStart, ConManDialogEnd)
}
func MayorGoHome() {
	ns.CreatureGuard(obj72, ns.GetWaypointX(wp112), ns.GetWaypointY(wp112), ns.GetObjectX(obj101), ns.GetObjectY(obj101), 0)
}
func UnlockTownGates() {
	gvar56 = gvar41
	ns.FrameTimer(30, MayorGoHome)
}
func PlayerDeath() {
	ns.DeathScreen(2)
}
func SwitchMaps() {
	ns.MoveObject(ns.GetHost(), ns.GetWaypointX(ns.Waypoint("ExitWP")), ns.GetWaypointY(ns.Waypoint("ExitWP")))
}
func EndMission() {
	ns.Blind()
	ns.FrameTimer(60, SwitchMaps)
}
func InitEvilMage() {
	obj119 = ns.Object("Necromancer")
	obj123 = ns.Object("Wolf1")
	obj124 = ns.Object("Wolf2")
	obj129 = ns.Object("Julie2")
	obj130 = ns.Object("Tanya2")
	obj131 = ns.Object("Julie")
	obj66 = ns.Object("Tanya")
	obj132 = ns.Object("Lydia")
	obj121 = ns.Object("SummonTrigger")
	gvar134 = ns.ObjectGroup("NecroTriggers")
	wp136 = ns.Waypoint("NecroWP")
	wp137 = ns.Waypoint("SpiderWP")
	wp138 = ns.Waypoint("EscapeWP")
	wp139 = ns.Waypoint("T1WP")
	wp140 = ns.Waypoint("T2WP")
	wp141 = ns.Waypoint("M3WP")
	wp142 = ns.Waypoint("M4WP")
	wp143 = ns.Waypoint("M3RunWP")
	wp144 = ns.Waypoint("M4RunWP")
	ns.ObjectOff(obj121)
	ns.FrameTimer(1, FunctionSetup)
}
func FunctionSetup() {
	ns.SetOwner(ns.GetHost(), obj129)
	ns.SetOwner(ns.GetHost(), obj130)
	ns.SetOwner(ns.GetHost(), obj131)
	ns.SetOwner(ns.GetHost(), obj66)
	ns.SetOwner(ns.GetHost(), obj132)
	ns.SetOwner(ns.GetHost(), obj123)
	ns.SetOwner(ns.GetHost(), obj124)
	ns.StoryPic(obj130, "MaidenPic3")
	ns.StoryPic(obj132, "SnobbyGirlPic")
	ns.SetDialog(obj130, ns.NORMAL, Maiden4TempStart, Maiden4TempEnd)
	ns.SetDialog(obj132, ns.NORMAL, Maiden7TempStart, Maiden7TempEnd)
}
func T1Trigger() {
	if flag146 {
		goto LABEL1
	}
	ivar145 += 1
	ns.CreatureGuard(obj125, ns.GetObjectX(obj125), ns.GetObjectY(obj125), ns.GetObjectX(obj119), ns.GetObjectY(obj119), 0)
	if !(ivar145 >= 4) {
		goto LABEL1
	}
	NecromancerSummon()
LABEL1:
	return
}
func T2Trigger() {
	if flag146 {
		goto LABEL1
	}
	ivar145 += 1
	ns.CreatureGuard(ns.GetTrigger(), ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger()), ns.GetObjectX(obj119), ns.GetObjectY(obj119), 0)
	if !(ivar145 >= 4) {
		goto LABEL1
	}
	NecromancerSummon()
LABEL1:
	return
}
func M3Trigger() {
	if flag146 {
		goto LABEL1
	}
	ivar145 += 1
	ns.CreatureGuard(ns.GetTrigger(), ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger()), ns.GetObjectX(obj119), ns.GetObjectY(obj119), 0)
	if !(ivar145 >= 4) {
		goto LABEL1
	}
	NecromancerSummon()
LABEL1:
	return
}
func M4Trigger() {
	if flag146 {
		goto LABEL1
	}
	ivar145 += 1
	ns.CreatureGuard(ns.GetTrigger(), ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger()), ns.GetObjectX(obj119), ns.GetObjectY(obj119), 0)
	if !(ivar145 >= 4) {
		goto LABEL1
	}
	NecromancerSummon()
LABEL1:
	return
}
func StartNecroPiece() {
	ns.MusicPushEvent()
	ns.ObjectGroupOff(gvar134)
	ns.WideScreen(true)
	ns.Frozen(ns.GetHost(), true)
	ns.CancelDialog(obj130)
	ns.CreatureIdle(ns.GetHost())
	ns.LookAtObject(ns.GetHost(), ns.ObjectID(wp136)) // FIXME
	ns.Move(obj119, wp136)
	ns.ObjectOn(obj125)
	ns.ObjectOn(obj126)
}
func NecroInPosition() {
	ns.CreatureIdle(obj119)
	ns.LookAtObject(obj119, ns.GetHost())
	ns.MoveWaypoint(wp113, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.NecromancerRecognize, wp113)
	ns.MoveWaypoint(wp113, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.MaleNPC1Talkable, wp113)
	ns.AudioEvent(ns.MaleNPC2Talkable, wp113)
	ns.Move(obj125, wp139)
	ns.Move(obj126, wp140)
	ns.Move(obj129, wp141)
	ns.Move(obj130, wp142)
	ns.FrameTimer(1, NecromancerStart)
}
func NecromancerStart() {
	ns.Music(11, 100)
	ns.SetDialog(obj119, ns.NORMAL, NecroBegin, NecroDone)
	ns.StoryPic(obj119, "NecromancerPic")
	ns.StartDialog(obj119, ns.GetHost())
}
func NecromancerSummon() {
	if !flag149 {
		goto LABEL1
	}
	flag146 = true
	ns.ObjectOn(obj121)
	ns.LookAtObject(obj119, obj121)
	ns.CastSpellObjectLocation(ns.SPELL_SUMMON_SPIDER, obj119, 1374, 3207)
	ns.AudioEvent(ns.NecromancerTaunt, wp113)
	goto LABEL2
LABEL1:
	ns.FrameTimer(2, NecromancerSummon)
LABEL2:
	return
}
func SpiderRun() {
	r1 := ns.IsAttackedBy(ns.GetCaller(), ns.GetHost())
	if !r1 {
		goto LABEL1
	}
	if ns.IsCaller(obj119) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.AudioEvent(ns.NecromancerTaunt, wp113)
	obj120 = ns.GetCaller()
	ns.MoveWaypoint(wp113, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.Maiden1Retreat, wp113)
	ns.Move(obj129, wp143)
	ns.FrameTimer(1, SpiderActivate)
LABEL1:
	return
}
func SpiderActivate() {
	ns.AggressionLevel(obj120, 0)
	ns.SetOwner(ns.GetHost(), obj120)
	ns.Move(obj120, wp137)
	ns.AudioEvent(ns.Maiden2Retreat, wp113)
	ns.Move(obj130, wp144)
	HenrickSayAttack()
}
func HenrickSayAttack() {
	ns.FrameTimer(15, PepperAttack)
}
func PepperAttack() {
	ns.AudioEvent(ns.SwordsmanRecognize, wp113)
	ns.Attack(obj126, obj119)
	ns.CreatureGuard(obj119, ns.GetObjectX(obj119), ns.GetObjectY(obj119), ns.GetObjectX(obj126), ns.GetObjectY(obj126), 0)
}
func NecroHit() {
	if flag148 {
		goto LABEL1
	}
	if !ns.IsCaller(obj126) {
		goto LABEL1
	}
	flag148 = true
	ns.Chat(obj119, "Con02A:NecroTalk02")
	ns.FrameTimer(60, NecroPoof)
LABEL1:
	return
}
func NecroPoof() {
	ns.DestroyEveryChat()
	ns.CreatureIdle(obj125)
	ns.CreatureIdle(obj126)
	ns.LookAtObject(obj126, obj119)
	ns.LookAtObject(obj125, obj119)
	ns.Move(obj119, wp138)
	ns.Attack(obj126, obj119)
	ns.Attack(obj125, obj119)
	ns.Chat(obj119, "Con02A:NecroTalk03")
	ns.FrameTimer(120, SetpieceOver)
}
func SetpieceOver() {
	ns.DestroyEveryChat()
	ns.WideScreen(false)
	ns.ObjectGroupOn(gvar135)
	ns.Frozen(ns.GetHost(), false)
	ns.Delete(obj119)
	ns.MoveObject(obj125, 1678, 2623)
	ns.MoveObject(obj126, 1708, 2655)
	ns.CreatureGuard(obj125, 1678, 2623, 1708, 2655, 0)
	ns.CreatureGuard(obj126, 1708, 2655, 1678, 2623, 0)
	ns.Wander(obj127)
	ns.Wander(obj128)
	ns.Wander(obj131)
	ns.Wander(obj66)
	ns.StoryPic(obj131, "MaidenPic2")
	ns.StoryPic(obj66, "MaidenPic3")
	ns.SetDialog(obj125, ns.NORMAL, Town1TempStart, Town1TempEnd)
	ns.SetDialog(obj126, ns.NORMAL, Town2TempStart, Town2TempEnd)
	ns.SetDialog(obj131, ns.NORMAL, Maiden5TempStart, Maiden5TempEnd)
	ns.SetDialog(obj66, ns.NORMAL, Maiden6TempStart, Maiden6TempEnd)
	ns.MusicPopEvent()
}
func DestroySpider() {
	if !ns.IsCaller(obj120) {
		goto LABEL1
	}
	ns.Delete(ns.GetCaller())
LABEL1:
	return
}
func M3Swap() {
	if !ns.IsCaller(obj129) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.MoveObject(obj131, ns.GetObjectX(obj129), ns.GetObjectY(obj129))
	ns.Delete(obj129)
LABEL1:
	return
}
func M4Swap() {
	if !ns.IsCaller(obj130) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.MoveObject(obj66, ns.GetObjectX(obj130), ns.GetObjectY(obj130))
	ns.Delete(obj130)
LABEL1:
	return
}
func NecroBegin() {
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:NecroTalk01")
}
func NecroDone() {
	ns.CancelDialog(obj119)
	ns.CreatureGuard(obj119, ns.GetObjectX(obj119), ns.GetObjectY(obj119), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	flag149 = true
}
func NecroYellStart() {
	ns.Frozen(obj125, true)
	ns.Frozen(obj126, true)
	ns.CreatureIdle(obj125)
	ns.CreatureIdle(obj126)
	ns.LookAtObject(obj126, obj119)
	ns.LookAtObject(obj125, obj119)
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:NecroTalk03")
}
func NecroYellEnd() {
	ns.CancelDialog(obj119)
	ns.Frozen(obj126, false)
	ns.Frozen(obj125, false)
	ns.Move(obj119, wp138)
	ns.Attack(obj126, obj119)
	ns.Attack(obj125, obj119)
	ns.FrameTimer(120, SetpieceOver)
}
func Town1TempStart() {
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:Townsman10Talk03")
}
func Town1TempEnd() {
	ns.Wander(obj125)
	ns.SetDialog(obj125, ns.NORMAL, Town1TalkStart, Town1TalkEnd)
}
func Town2TempStart() {
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:Townsman13Talk02")
}
func Town2TempEnd() {
	ns.Wander(obj126)
	ns.SetDialog(obj126, ns.NORMAL, Town2TalkStart, Town2TalkEnd)
}
func Maiden5TempStart() {
	ns.Frozen(obj131, true)
	ns.CreatureIdle(obj131)
	ns.LookAtObject(obj131, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:Maiden6Talk03")
}
func Maiden5TempEnd() {
	ns.Frozen(obj131, false)
	ns.Wander(obj131)
	ns.SetDialog(obj131, ns.NORMAL, Maiden5TalkStart, Maiden5TalkEnd)
}
func Maiden4TempStart() {
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:Maiden3Talk01")
}
func Maiden4TempEnd() {
	ns.SetDialog(obj130, ns.NORMAL, Maiden4TempStart, Maiden4TempEnd)
}
func Maiden6TempStart() {
	ns.Frozen(obj66, true)
	ns.CreatureIdle(obj66)
	ns.LookAtObject(obj66, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:Maiden6Talk01")
}
func Maiden6TempEnd() {
	ns.Frozen(obj66, false)
	ns.Wander(obj66)
	ns.SetDialog(obj66, ns.NORMAL, Maiden6TempStart, Maiden6TempEnd)
}
func Maiden7TempStart() {
	ns.LookAtObject(obj132, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "Con02a:BarMaiden2")
}
func Maiden7TempEnd() {
	ns.SetDialog(obj132, ns.NORMAL, Maiden7TempStart, Maiden7TempEnd)
}
func Town1TalkStart() {
	var (
		v0 int
		v1 int
	)
	ns.Frozen(obj125, true)
	ns.CreatureIdle(obj125)
	ns.LookAtObject(obj125, ns.GetHost())
	if gvar208 != gvar202 {
		goto LABEL1
	}
	v0 = ns.Random(1, 2)
	v1 = v0
	if v1 == 1 {
		goto LABEL2
	}
	if v1 == 2 {
		goto LABEL3
	}
	goto LABEL4
LABEL2:
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:Townsman10Talk02")
	goto LABEL4
LABEL3:
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:Townsman10Talk03")
	goto LABEL4
LABEL4:
	goto LABEL5
LABEL1:
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:Townsman10Talk01")
LABEL5:
	return
}
func Town1TalkEnd() {
	ns.Frozen(obj125, false)
	ns.Wander(obj125)
	ns.SetDialog(obj125, ns.NORMAL, Town1TalkStart, Town1TalkEnd)
}
func Town2TalkStart() {
	var (
		v0 int
		v1 int
	)
	ns.Frozen(obj126, true)
	ns.CreatureIdle(obj126)
	ns.LookAtObject(obj126, ns.GetHost())
	if gvar209 != gvar204 {
		goto LABEL1
	}
	v0 = ns.Random(1, 2)
	v1 = v0
	if v1 == 1 {
		goto LABEL2
	}
	if v1 == 2 {
		goto LABEL3
	}
	goto LABEL4
LABEL2:
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:Townsman13Talk01")
	goto LABEL4
LABEL3:
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:Townsman13Talk03")
	goto LABEL4
LABEL4:
	goto LABEL5
LABEL1:
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:Townsman13Talk02")
LABEL5:
	return
}
func Town2TalkEnd() {
	ns.Frozen(obj126, false)
	ns.Wander(obj126)
	ns.SetDialog(obj126, ns.NORMAL, Town2TalkStart, Town2TalkEnd)
}
func Maiden5TalkStart() {
	ns.Frozen(obj131, true)
	ns.CreatureIdle(obj131)
	ns.LookAtObject(obj131, ns.GetHost())
	if flag147 {
		goto LABEL1
	}
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:Maiden6Talk03")
	goto LABEL2
LABEL1:
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:Maiden6Talk02")
LABEL2:
	return
}
func Maiden5TalkEnd() {
	ns.Frozen(obj131, false)
	ns.Wander(obj131)
	ns.SetDialog(obj131, ns.NORMAL, Maiden5TalkStart, Maiden5TalkEnd)
}
func EnableRandomBump() {
	flag159 = true
}
func RandomBump() {
	if !(ns.IsCaller(ns.GetHost()) && flag159) {
		goto LABEL1
	}
	ns.Chat(ns.GetTrigger(), "Con02a:RandomBump")
	flag159 = false
	ns.FrameTimer(ivar160, EnableRandomBump)
LABEL1:
	return
}
func RestAwhile() {
	var v0 int
	v0 = ns.Random(60, 120)
	ns.PauseObject(ns.GetCaller(), v0)
}
func EnableSpiders() {
	ns.ObjectGroupOn(gvar167)
}
func OpenMagicShopElevatorWalls() {
	ns.WallGroupOpen(gvar168)
}
func TransientMonsterDie() {
	ns.DeleteObjectTimer(ns.GetTrigger(), 120)
}
func SetToGuard(a1 int) {
	var (
		v0 float32
		v1 float32
		v2 int
	)
	v2 = a1
	if v2 == gvar150 {
		goto LABEL1
	}
	if v2 == gvar151 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	v0 = ns.GetWaypointX(wp163)
	v1 = ns.GetWaypointY(wp163)
	ns.MoveObject(obj165, ns.GetWaypointX(wp161), ns.GetWaypointY(wp161))
	ns.CreatureGuard(obj165, v0, v1, fvar155, fvar156, 100)
	goto LABEL3
LABEL2:
	v0 = ns.GetWaypointX(wp163)
	v1 = ns.GetWaypointY(wp163)
	ns.CreatureGuard(obj157, v0, v1, fvar155, fvar156, 100)
	goto LABEL3
LABEL3:
	return
}
func CallToConjurer() {
	if !(ns.IsCaller(ns.GetHost()) && flag45 == true && flag158 == false) {
		goto LABEL1
	}
	ns.Chat(ns.GetTrigger(), "Con02a:CallToConjurer")
	flag158 = true
LABEL1:
	return
}
func MayorCall() {
	ns.Chat(obj72, "Con02a:MayorCall")
	SpiderEngine()
}
func MayorsGateTriggerFilter() int {
	if !ns.IsCaller(obj72) {
		goto LABEL1
	}
	if !flag118 {
		goto LABEL2
	}
	ns.PrintToAll("Mayor detected at gown gates.")
LABEL2:
	return 1
	goto LABEL3
LABEL1:
	if !flag118 {
		goto LABEL4
	}
	ns.PrintToAll("Trigger hit by object other than Mayor Theogrin")
LABEL4:
	return 0
LABEL3:
	return 0
}
func UnlockMayorBedroomoor() {
	ns.UnlockDoor(obj88)
}
func LockMayorsBedroomDoor() {
	ns.LockDoor(obj88)
}
func FoundBoots() {
	ns.ObjectOff(ns.GetTrigger())
	ns.PrintToAll("Con02a:FoundBoots")
	ns.AudioEvent(ns.JournalEntryAdd, ns.Waypoint("BootsAudioOrigin"))
}
func Secret01Found() {
	ns.ObjectGroupOff(gvar166)
	ns.MoveWaypoint(wp164, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretFound, wp164)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 50)
}
func Secret02Found() {
	ns.ObjectOff(ns.GetTrigger())
	ns.MoveWaypoint(wp164, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretFound, wp164)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 50)
}
func EnemyGoHome() {
	r1 := ns.IsAttackedBy(ns.GetHost(), ns.GetCaller())
	if !r1 {
		goto LABEL1
	}
	ns.GoBackHome(ns.GetCaller())
LABEL1:
	return
}
func DebugOn() {
	flag118 = true
	ns.PrintToAll("Debug mode is enabled.")
}
func OwnObjects() {
	ns.GroupSetOwner(ns.GetHost(), gvar135)
	ns.SetOwner(ns.GetHost(), obj68)
	ns.SetOwner(ns.GetHost(), obj96)
	ns.SetOwner(ns.GetHost(), obj99)
	ns.SetOwner(ns.GetHost(), obj128)
	ns.SetOwner(ns.GetHost(), obj176)
	ns.SetOwner(ns.GetHost(), obj177)
	ns.SetOwner(ns.GetHost(), obj125)
	ns.SetOwner(ns.GetHost(), obj126)
	ns.SetOwner(ns.GetHost(), obj127)
	ns.SetOwner(ns.GetHost(), obj93)
	ns.ObjectGroupOff(gvar135)
}
func InitializePiece() {
	obj201 = ns.Object("Jacob")
	obj100 = ns.Object("Henrick")
	obj122 = ns.Object("Pepper")
	obj133 = ns.Object("NecroDoor")
	obj214 = ns.Object("TempDoorR")
	obj215 = ns.Object("TempDoorL")
	obj216 = ns.Object("CemDoorR")
	obj217 = ns.Object("CemDoorL")
	gvar218 = ns.ObjectGroup("WolfTriggers")
	wp113 = ns.Waypoint("PlayerSounds")
	wp219 = ns.Waypoint("WolfCreateWP")
	ns.FrameTimer(1, SetupCreatures)
}
func DialogSetup() {
	ns.StoryPic(obj125, "MalePic4")
	ns.StoryPic(obj126, "MalePic11")
	ns.StoryPic(obj127, "MalePic12")
	ns.StoryPic(obj128, "MaidenPic4")
	ns.StoryPic(obj201, "WardenPic")
	ns.StoryPic(obj100, "HenrickPic")
	ns.SetDialog(obj127, ns.NORMAL, Town3TalkStart, Town3TalkEnd)
	ns.SetDialog(obj201, ns.NORMAL, Jail1Start, Jail1End)
}
func PlayIxTownMusic() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.Music(15, 40)
LABEL1:
	return
}
func DisableElevators() {
	PlayIxTownMusic()
	ns.JournalEntry(ns.GetHost(), "LocateAldwyn", 2)
	ns.PrintToAll("Con01a:NewJournalEntry")
	ns.ObjectOff(obj84)
	ns.ObjectOff(obj171)
	ns.ObjectOff(obj98)
}
func PickupObjects() {
	ns.Pickup(obj11, obj87)
	ns.Pickup(obj11, obj85)
	ns.Pickup(obj11, obj86)
	ns.Pickup(obj96, obj97)
}
func MapInitialize() {
	obj71 = ns.Object("Aldwyn")
	obj68 = ns.Object("Geoff")
	obj72 = ns.Object("Mayor_Theogrin")
	obj69 = ns.Object("Contest_Guard")
	obj70 = ns.Object("Mayor's_Guard")
	obj67 = ns.Object("Barkeeper")
	obj11 = ns.Object("Contest_Official")
	obj92 = ns.Object("Bridge_Guard")
	obj93 = ns.Object("Guard2")
	obj176 = ns.Object("IxGuard1")
	obj177 = ns.Object("IxGuard2")
	obj165 = ns.Object("MayorsSpider1")
	obj75 = ns.Object("MayorGate1")
	obj76 = ns.Object("MayorGate2")
	obj77 = ns.Object("ContestGate1")
	obj78 = ns.Object("ContestGate2")
	obj79 = ns.Object("ContestGate3")
	obj169 = ns.Object("BasementDoor")
	obj84 = ns.Object("AldwinElevator")
	obj171 = ns.Object("MayorElevator")
	obj81 = ns.Object("RatFieldGuide")
	obj82 = ns.Object("BigSpiderFG")
	obj80 = ns.Object("CharmSpell")
	obj83 = ns.Object("ElevatorDoor")
	obj97 = ns.Object("CheapBow")
	obj173 = ns.Object("Quiver")
	obj12 = ns.Object("FiringLineTrigger")
	obj87 = ns.Object("ContestPrize")
	obj85 = ns.Object("ContestQuiver1")
	obj86 = ns.Object("ContestQuiver2")
	obj73 = ns.Object("LeatherTunic")
	obj88 = ns.Object("MayorsBedroomDoor")
	obj90 = ns.Object("RightTownGate")
	obj91 = ns.Object("LeftTownGate")
	obj174 = ns.Object("TownGate01")
	obj175 = ns.Object("TownGate02")
	obj89 = ns.Object("Gold")
	obj125 = ns.Object("Bryan")
	obj126 = ns.Object("Clyde")
	obj127 = ns.Object("Tommy")
	obj74 = ns.Object("Cloak")
	obj94 = ns.Object("BridgeGuardsBoots")
	obj95 = ns.Object("BootsReward")
	obj96 = ns.Object("Morgan")
	obj98 = ns.Object("MagicShopElevator")
	obj99 = ns.Object("Joyce")
	obj128 = ns.Object("Gretchen")
	obj101 = ns.Object("MayorsFireplace")
	obj178 = ns.Object("CaveDoor")
	obj102 = ns.Object("Heckler")
	obj179 = ns.Object("WaspNest")
	obj180[0] = ns.Object("Wasp1")
	obj180[1] = ns.Object("Wasp2")
	gvar135 = ns.ObjectGroup("NPCs")
	gvar181 = ns.ObjectGroup("SpiderDropTriggers1")
	gvar182 = ns.ObjectGroup("SpiderDropTriggers2")
	gvar103 = ns.ObjectGroup("GateTriggers")
	gvar183 = ns.ObjectGroup("WaspTriggers")
	gvar184 = ns.ObjectGroup("Urchins")
	gvar167 = ns.ObjectGroup("Spiders")
	gvar166 = ns.ObjectGroup("Secret01Triggers")
	gvar168 = ns.WallGroup("MagicShopElevatorWalls")
	wp193 = ns.Waypoint("WellWP")
	wp104 = ns.Waypoint("Study")
	wp185 = ns.Waypoint("MayorsGates")
	wp13 = ns.Waypoint("TargetWP1")
	wp14 = ns.Waypoint("TargetWP2")
	wp15 = ns.Waypoint("TargetWP3")
	wp186 = ns.Waypoint("OfficialWP")
	wp187 = ns.Waypoint("CreationWP")
	wp16 = ns.Waypoint("ContestAudioOrigin")
	wp188 = ns.Waypoint("UrchinWP")
	wp105 = ns.Waypoint("ImpWP1")
	wp106 = ns.Waypoint("ImpWP2")
	wp107 = ns.Waypoint("ImpWP3")
	wp108 = ns.Waypoint("ImpWP4")
	wp109 = ns.Waypoint("AldwinExitWP")
	wp189 = ns.Waypoint("SpiderDropPoint1")
	wp190 = ns.Waypoint("SpiderDropPoint2")
	wp110 = ns.Waypoint("MayorRewardPath")
	wp111 = ns.Waypoint("TownGatePath")
	wp112 = ns.Waypoint("MayorHome")
	wp191 = ns.Waypoint("AudioOrigin")
	wp162 = ns.Waypoint("StudySpiderWP")
	wp161 = ns.Waypoint("BigSpiderWP")
	wp163 = ns.Waypoint("SpiderGuardPoint")
	wp192 = ns.Waypoint("WaspCreationWP")
	wp164 = ns.Waypoint("SecretAudioWP")
	wp194 = ns.Waypoint("MayorsGuardWP")
	wp199 = ns.Waypoint("BridgeGuardWP")
	wp198 = ns.Waypoint("ContestGuardWP")
	wp195 = ns.Waypoint("GatekeeperWP")
	wp196 = ns.Waypoint("Gatekeeper2WP")
	wp197 = ns.Waypoint("Gatekeeper3WP")
	wp200[0] = ns.Waypoint("WaspWP1")
	wp200[1] = ns.Waypoint("WaspWP2")
	gvar114 = ns.WaypointGroup("ImpEntrancePath")
	gvar115 = ns.WaypointGroup("MayorsPathToGate")
	gvar116 = ns.WaypointGroup("MayorsHomeSpiderPath")
	ns.WayPointGroupOff(gvar115)
	ns.LockDoor(obj75)
	ns.LockDoor(obj76)
	ns.LockDoor(obj174)
	ns.LockDoor(obj175)
	ns.LockDoor(obj77)
	ns.LockDoor(obj78)
	ns.LockDoor(obj79)
	ns.LockDoor(obj169)
	ns.LockDoor(obj88)
	ns.LockDoor(obj90)
	ns.LockDoor(obj91)
	ns.LockDoor(obj178)
	InitializePiece()
	InitEvilMage()
	DialogSetup()
	InitializeDialogs()
	ns.StartupScreen(2)
	ns.FrameTimer(1, OwnObjects)
	ns.FrameTimer(35, DisableElevators)
}
func DebugOff() {
	flag118 = false
	ns.PrintToAll("Debug mode has been disabled.")
}
func AldwinCheck() {
	if !ns.HasItem(obj71, obj80) {
		goto LABEL1
	}
	ns.PrintToAll("Spell is in Aldwins inv.")
	goto LABEL2
LABEL1:
	ns.PrintToAll("Spell not found in Aldwins inv.")
LABEL2:
	if !ns.HasItem(obj71, obj81) {
		goto LABEL3
	}
	ns.PrintToAll("Field Guide is in Aldwins inv.")
	goto LABEL4
LABEL3:
	ns.PrintToAll("Field Guide not found in Aldwins inv.")
LABEL4:
	return
}
func PlayUrchinDenMusic() {
	ns.Music(17, 100)
}
func PlayTunnelMusic() {
	ns.Music(18, 100)
}
func ExitDen() {
	gvar211 = gvar205
	ns.SetDialog(obj201, ns.NORMAL, Jail2Start, Jail2End)
	ns.MoveObject(obj123, 2969, 2354)
	ns.MoveObject(obj124, 3016, 2308)
	ns.CreatureGuard(obj123, 2969, 2354, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	ns.CreatureGuard(obj124, 3016, 2308, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	ns.SetOwner(ns.GetHost(), obj123)
	ns.SetOwner(ns.GetHost(), obj124)
}
func Town3TalkStart() {
	ns.Frozen(obj127, true)
	ns.CreatureIdle(obj127)
	ns.LookAtObject(obj127, ns.GetHost())
	if gvar210 != gvar202 {
		goto LABEL1
	}
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:Townsman9Talk01")
	goto LABEL2
LABEL1:
	if gvar210 == gvar207 {
		goto LABEL3
	}
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:Townsman9Talk02")
LABEL3:
	if gvar210 != gvar207 {
		goto LABEL2
	}
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:Townsman9Talk03")
LABEL2:
	return
}
func Town3TalkEnd() {
	ns.Frozen(obj127, false)
	ns.Wander(obj127)
	ns.SetDialog(obj127, ns.NORMAL, Town3TalkStart, Town3TalkEnd)
}
func Maiden2TalkStart() {
	ns.Frozen(obj128, true)
	ns.CreatureIdle(obj128)
	ns.LookAtObject(obj128, ns.GetHost())
	if gvar211 == gvar205 {
		goto LABEL1
	}
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:Townsman2Talk02")
	goto LABEL2
LABEL1:
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:Maiden2Talk01")
LABEL2:
	return
}
func Maiden2TalkEnd() {
	ns.Frozen(obj128, false)
	ns.Wander(obj128)
	ns.SetDialog(obj128, ns.NORMAL, Maiden2TalkStart, Maiden2TalkEnd)
}
func HenrickGreetStart() {
	ns.DestroyEveryChat()
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:HenrickTalk03")
}
func HenrickGreetEnd() {
	ns.SetDialog(obj100, ns.NORMAL, HenrickStart, HenrickEnd)
}
func HenrickStart() {
	var (
		v0 int
		v1 int
		v2 int
	)
	ns.DestroyEveryChat()
	if gvar213 != gvar202 {
		goto LABEL1
	}
	v0 = ns.Random(1, 3)
	v1 = v0
	if v1 == 1 {
		goto LABEL2
	}
	if v1 == 2 {
		goto LABEL3
	}
	if v1 == 3 {
		goto LABEL4
	}
	goto LABEL1
LABEL2:
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:HenrickTalk04")
	goto LABEL1
LABEL3:
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:HenrickTalk05")
	goto LABEL1
LABEL4:
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:HenrickTalk03")
	goto LABEL1
LABEL1:
	if gvar213 != gvar204 {
		goto LABEL5
	}
	v0 = ns.Random(1, 2)
	v2 = v0
	if v2 == 1 {
		goto LABEL6
	}
	if v2 == 2 {
		goto LABEL7
	}
	goto LABEL5
LABEL6:
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:HenrickTalk04")
	goto LABEL5
LABEL7:
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:HenrickTalk06")
	goto LABEL5
LABEL5:
	return
}
func HenrickEnd() {
	ns.SetDialog(obj100, ns.NORMAL, HenrickStart, HenrickEnd)
}
func Jail1Start() {
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:JailerTalk01")
}
func Jail1End() {
	ns.SetDialog(obj201, ns.NORMAL, Jail1Start, Jail1End)
}
func Jail2Start() {
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:JailerTalk02")
}
func Jail2End() {
	ns.SetDialog(obj201, ns.NORMAL, Jail2Start, Jail2End)
}
func MakeAWish() {
	var v0 int
	v0 = ns.MaxHealth(ns.GetCaller())
	ns.RestoreHealth(ns.GetCaller(), v0-ns.CurrentHealth(ns.GetCaller()))
	ns.PrintToAll("GeneralPrint:WellSignRefresh")
	ns.AudioEvent(ns.RestoreHealthName, wp193)
}
func SetupCreatures() {
	ns.SetOwner(ns.GetHost(), obj100)
	ns.SetOwner(ns.GetHost(), obj201)
	ns.SetOwner(ns.GetHost(), obj122)
	ns.LockDoor(obj133)
	ns.LockDoor(obj214)
	ns.LockDoor(obj215)
	ns.LockDoor(obj216)
	ns.LockDoor(obj217)
}
func StartWolfAttack() {
	if flag220 {
		goto LABEL1
	}
	ns.ObjectGroupOff(gvar218)
	ns.WideScreen(true)
	ns.Frozen(ns.GetHost(), true)
	ns.CreatureIdle(ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj100)
	ns.MoveObject(obj122, ns.GetWaypointX(wp219), ns.GetWaypointY(wp219))
	ns.FrameTimer(1, ActivateWolf)
LABEL1:
	return
}
func ActivateWolf() {
	ns.CreatureIdle(obj122)
	ns.AggressionLevel(obj122, 0.83)
	ns.Attack(obj122, ns.GetHost())
	flag220 = true
	ns.FrameTimer(50, HenrickWarning)
}
func HenrickWarning() {
	ns.Chat(obj100, "Con02A:HenrickTalk01")
	ns.Frozen(obj100, true)
	ns.CreatureIdle(obj100)
	ns.LookAtObject(obj100, obj122)
	ns.Frozen(obj122, true)
	ns.LookAtObject(obj122, obj100)
	ns.MoveWaypoint(wp113, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.CharmCast, wp113)
	ns.FrameTimer(20, CharmPepper)
}
func CharmPepper() {
	ns.Effect(ns.CHARM, ns.GetObjectX(obj100), ns.GetObjectY(obj100), ns.GetObjectX(obj122), ns.GetObjectY(obj122))
	ivar221 += 1
	if !(ivar221 < 60) {
		goto LABEL1
	}
	ns.FrameTimer(1, CharmPepper)
	goto LABEL2
LABEL1:
	CharmEnd()
LABEL2:
	return
}
func CharmEnd() {
	ns.AudioEvent(ns.CharmSuccess, wp113)
	ns.WideScreen(false)
	ns.Frozen(ns.GetHost(), false)
	ns.Frozen(obj100, false)
	ns.Frozen(obj122, false)
	ns.CreatureGuard(obj100, ns.GetObjectX(obj100), ns.GetObjectY(obj100), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	ns.SetOwner(ns.GetHost(), obj122)
	ns.AggressionLevel(obj122, 0.16)
	ns.CreatureFollow(obj122, obj100)
	ns.Chat(obj100, "Con02A:HenrickTalk02")
	ns.SetDialog(obj100, ns.NORMAL, HenrickGreetStart, HenrickGreetEnd)
}
func OnEvent(typ string) {
	switch typ {
	case "PlayerDeath":
		PlayerDeath()
	case "MapInitialize":
		MapInitialize()
	}
}
