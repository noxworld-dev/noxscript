package war07a

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	wp4     ns.WaypointID
	obj5    [100]ns.ObjectID
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
	gvar53  ns.ObjectGroupID
	obj54   ns.ObjectID
	obj55   ns.ObjectID
	obj56   ns.ObjectID
	obj57   ns.ObjectID
	obj58   ns.ObjectID
	obj59   ns.ObjectID
	obj60   ns.ObjectID
	obj61   ns.ObjectID
	obj62   ns.ObjectID
	gvar63  ns.WallGroupID
	gvar64  ns.WallGroupID
	flag65  bool
	flag66  bool
	flag67  bool
	flag68  bool
	flag69  bool
	flag70  bool
	flag71  bool
	flag72  bool
	flag73  bool
	flag74  bool
	flag75  bool
	flag76  bool
	flag77  bool
	flag78  bool
	flag79  bool
	flag80  bool
	flag81  bool
	flag82  bool
	flag83  bool
	flag84  bool
	flag85  bool
	flag86  bool
	flag87  bool
	flag88  bool
	flag89  bool
	flag90  bool
	flag91  bool
	flag92  bool
	flag93  bool
	gvar94  int
	gvar95  int
	gvar96  int
	gvar97  int
	gvar98  int
	gvar99  int
	gvar100 int
	gvar101 int
	gvar102 int
	ivar103 int
	gvar104 int
	ivar105 int
	ivar106 int
	ivar107 int
	gvar108 int
	gvar109 int
	fvar110 float32
	fvar111 float32
	fvar112 float32
	fvar113 float32
	fvar114 float32
	fvar115 float32
	fvar116 float32
	fvar117 float32
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
	wp131   ns.WaypointID
	wp132   ns.WaypointID
	wp133   ns.WaypointID
	gvar134 ns.WaypointGroupID
	gvar135 ns.WaypointGroupID
	wp136   ns.WaypointID
	wp137   ns.WaypointID
	wp138   ns.WaypointID
	wp139   ns.WaypointID
	wp140   ns.WaypointID
	wp141   ns.WaypointID
	gvar142 int
	gvar143 int
	gvar144 int
	gvar145 int
	gvar146 int
	gvar147 int
	gvar148 int
	gvar149 int
	gvar150 int
	gvar151 int
	ivar152 int
	ivar153 int
	ivar154 int
	flag155 bool
	gvar156 int
	flag157 bool
	flag158 bool
	flag159 bool
	flag160 bool
	flag161 bool
	flag162 bool
)

func init() {
	flag65 = false
	flag66 = false
	flag67 = false
	flag68 = false
	flag69 = false
	flag70 = false
	flag71 = false
	flag72 = false
	flag73 = false
	flag74 = false
	flag75 = false
	flag76 = false
	flag77 = false
	flag78 = false
	flag79 = false
	flag80 = false
	flag81 = false
	flag82 = false
	flag83 = false
	flag84 = false
	flag85 = false
	flag86 = false
	flag87 = false
	flag88 = false
	flag89 = false
	flag90 = false
	flag92 = false
	flag93 = false
	gvar94 = 0
	gvar95 = 0
	gvar96 = 0
	gvar97 = 0
	gvar98 = 0
	gvar99 = 0
	gvar100 = 0
	gvar101 = 0
	gvar102 = 0
	ivar103 = 0
	gvar104 = 0
	ivar105 = 0
	ivar106 = 0
	gvar108 = 100
	gvar118 = 0
	gvar119 = 1
	gvar120 = 2
	gvar121 = 3
	gvar122 = 0
	gvar123 = 1
	gvar124 = 2
	gvar125 = 3
	gvar126 = 0
	gvar127 = 1
	gvar128 = 2
	gvar129 = 3
	gvar142 = 0
	gvar143 = 1
	gvar144 = 2
	gvar145 = 3
	gvar146 = 4
	gvar147 = 5
	gvar148 = gvar142
	gvar149 = 0
	gvar150 = 1
	gvar151 = gvar149
	ivar152 = 0
	ivar153 = 0
	ivar154 = 0
	flag155 = true
	flag157 = false
	flag158 = false
	flag159 = false
	flag160 = false
	flag161 = false
	flag162 = false
	flag91 = false
}
func CreatureSetup() {
	ivar107 = ns.CurrentHealth(obj38)
	ivar106 = 0
	for {
		if !(ivar106 < 100) {
			goto LABEL1
		}
		obj5[ivar106] = 0
		ivar106 += 1
	}
LABEL1:
	ns.Pickup(obj22, obj59)
	ns.SetOwner(ns.GetHost(), obj23)
	ns.SetOwner(ns.GetHost(), obj24)
	ns.SetOwner(ns.GetHost(), obj11)
	ns.SetOwner(ns.GetHost(), obj6)
	ns.SetOwner(ns.GetHost(), obj7)
	ns.SetOwner(ns.GetHost(), obj8)
	ns.SetOwner(obj22, obj33)
	ns.SetOwner(obj22, obj28)
	ns.SetOwner(obj22, obj26)
	ns.SetOwner(obj22, obj25)
	ns.SetOwner(obj22, obj27)
	ns.SetOwner(obj22, obj29)
	ns.SetOwner(obj22, obj30)
	ns.SetOwner(obj22, obj31)
	ns.SetOwner(obj22, obj21)
	ns.SetOwner(ns.GetHost(), obj36)
	ns.Wander(obj38)
	ns.CreatureGuard(obj12, ns.GetObjectX(obj12), ns.GetObjectY(obj12), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	ns.CreatureGuard(obj16, ns.GetObjectX(obj16), ns.GetObjectY(obj16), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	ns.CreatureGuard(obj31, ns.GetObjectX(obj31), ns.GetObjectY(obj31), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	fvar110 = ns.GetObjectX(obj14)
	fvar111 = ns.GetObjectY(obj14)
	fvar112 = ns.GetObjectX(obj9)
	fvar113 = ns.GetObjectY(obj9)
	fvar114 = ns.GetObjectX(obj12)
	fvar115 = ns.GetObjectY(obj12)
	fvar116 = ns.GetObjectX(obj16)
	fvar117 = ns.GetObjectY(obj16)
	ns.Music(17, 100)
}
func PlayerChurchEnter() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	flag82 = true
	goto LABEL2
LABEL1:
	ns.Wander(ns.GetCaller())
LABEL2:
	return
}
func PlayerChurchExit() {
	flag82 = false
	if !flag81 {
		goto LABEL1
	}
	flag81 = false
	ns.CreatureHunt(obj25)
	ns.CreatureHunt(obj28)
	ns.CreatureHunt(obj27)
	ns.CreatureHunt(obj26)
	ns.CreatureHunt(obj29)
	ns.CreatureHunt(obj30)
LABEL1:
	return
}
func PlayerSneakExit() {
	flag82 = false
}
func DrunkRecognize() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	if gvar100 != gvar118 {
		goto LABEL1
	}
	if !(ns.CurrentHealth(obj31) > 0) {
		goto LABEL1
	}
	ns.SetDialog(obj31, ns.NORMAL, DrunkStart, DrunkEnd)
	gvar100 = gvar119
	ns.CreatureFollow(obj31, ns.GetHost())
	ns.StartDialog(obj31, ns.GetHost())
LABEL1:
	if !ns.IsCaller(obj22) {
		goto LABEL2
	}
	if !(flag72 == false && flag68 == false) {
		goto LABEL2
	}
	if !(ns.CurrentHealth(obj22) > 0) {
		goto LABEL2
	}
	ns.CreatureFollow(obj22, ns.GetHost())
LABEL2:
	return
}
func DrunkReport() {
	if !flag93 {
		goto LABEL1
	}
	flag93 = false
	ns.Wander(obj31)
LABEL1:
	return
}
func DrunkAttacks() {
	var (
		v0 int
		v1 int
	)
	if gvar100 != gvar119 {
		goto LABEL1
	}
	if !(ns.CurrentHealth(obj31) > 0) {
		goto LABEL1
	}
	v0 = ns.Random(1, 5)
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
	if v1 == 4 {
		goto LABEL5
	}
	if v1 == 5 {
		goto LABEL6
	}
	goto LABEL7
LABEL2:
	ns.Chat(obj31, "War07A.scr:DrunkardTaunt01")
	goto LABEL7
LABEL3:
	ns.Chat(obj31, "War07A.scr:DrunkardTaunt02")
	goto LABEL7
LABEL4:
	ns.Chat(obj31, "War07A.scr:DrunkardTaunt03")
	goto LABEL7
LABEL5:
	ns.Chat(obj31, "War07A.scr:DrunkardTaunt04")
	goto LABEL7
LABEL6:
	ns.Chat(obj31, "War07A.scr:DrunkardTaunt05")
	goto LABEL7
LABEL7:
	ns.Attack(obj31, ns.GetHost())
	ns.SecondTimer(5, DrunkAttacks)
LABEL1:
	return
}
func DrunkInjured() {
	if flag72 {
		goto LABEL1
	}
	if !(ns.CurrentHealth(obj31) > 0) {
		goto LABEL2
	}
	if flag74 {
		goto LABEL3
	}
	flag74 = true
	ns.Chat(obj31, "War07A.scr:DrunkardHelp01")
LABEL3:
	gvar100 = gvar120
	ns.AggressionLevel(obj22, 0.16)
	ns.AggressionLevel(obj31, 0.16)
	if !(ns.CurrentHealth(obj22) > 0) {
		goto LABEL2
	}
	ns.CreatureFollow(obj31, obj22)
LABEL2:
	goto LABEL4
LABEL1:
	ns.AggressionLevel(obj31, 0.83)
LABEL4:
	return
}
func DrunkDied() {
	ns.DestroyEveryChat()
	flag68 = true
	ns.MoveObject(obj9, ns.GetWaypointX(wp137), ns.GetWaypointY(wp137))
	ns.MoveObject(obj12, ns.GetWaypointX(wp139), ns.GetWaypointY(wp139))
	ns.MoveObject(obj16, ns.GetWaypointX(wp138), ns.GetWaypointY(wp138))
	ns.MoveObject(obj14, ns.GetWaypointX(wp136), ns.GetWaypointY(wp136))
	ns.CancelDialog(obj22)
	if gvar102 != gvar126 {
		goto LABEL1
	}
	gvar102 = gvar127
	ns.AggressionLevel(obj28, 0.16)
	ns.AggressionLevel(obj27, 0.16)
	ns.AggressionLevel(obj25, 0.16)
	ns.AggressionLevel(obj26, 0.16)
	ns.AggressionLevel(obj29, 0.16)
	ns.AggressionLevel(obj30, 0.16)
	goto LABEL2
LABEL1:
	gvar102 = gvar128
LABEL2:
	return
}
func DrunkStart() {
	ns.Frozen(ns.GetHost(), true)
	ns.CreatureIdle(ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj31)
	ns.TellStory(ns.HumanMaleEatFood, "War07A.scr:DrunkardIntro01")
}
func DrunkEnd() {
	ns.Frozen(ns.GetHost(), false)
	ns.CancelDialog(obj31)
	ns.AggressionLevel(obj31, 0.5)
	ns.Attack(obj31, ns.GetHost())
	ns.SecondTimer(3, DrunkAttacks)
}
func DrunkLoseSight() {
	if !(gvar100 == gvar119 && flag72 == false && flag70 == false) {
		goto LABEL1
	}
	ns.AggressionLevel(obj31, 0.16)
	ns.CreatureFollow(obj31, obj22)
	gvar100 = gvar120
	flag74 = false
LABEL1:
	return
}
func PutPeopleInJail() {
	IndexPlayerInventory()
	flag71 = true
	ns.PrintToAll("War07A.scr:TimePass")
	ns.MoveObject(ns.GetHost(), 3741, 5350)
	ns.MoveObject(obj22, 3774, 5385)
	ns.AggressionLevel(obj22, 0)
	ns.LookAtObject(obj22, ns.GetHost())
	flag72 = true
	if !(ns.CurrentHealth(obj31) > 0) {
		goto LABEL1
	}
	ns.AggressionLevel(obj31, 0.16)
	ns.Wander(obj31)
LABEL1:
	ns.UnBlind()
	WardenJailCellTalk()
}
func GetRandomMobYell() {
	var v0 int
	ivar105 = ns.Random(1, ivar103)
	v0 = ivar105
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
	if v0 == 5 {
		goto LABEL5
	}
	if v0 == 6 {
		goto LABEL6
	}
	goto LABEL7
LABEL1:
	if !(flag75 == true && ns.CurrentHealth(obj25) > 0) {
		goto LABEL8
	}
	ns.Chat(obj25, "War07A.scr:MobYell")
	ns.AggressionLevel(obj25, 0.83)
	ns.AggressionLevel(obj28, 0.83)
	ns.AggressionLevel(obj27, 0.83)
	ns.AggressionLevel(obj26, 0.83)
	ns.AggressionLevel(obj29, 0.83)
	ns.AggressionLevel(obj30, 0.83)
	ns.Attack(obj25, ns.GetHost())
	ns.Attack(obj28, ns.GetHost())
	ns.Attack(obj27, ns.GetHost())
	ns.Attack(obj26, ns.GetHost())
	ns.Attack(obj29, ns.GetHost())
	ns.Attack(obj30, ns.GetHost())
	flag80 = true
	goto LABEL9
LABEL8:
	if flag80 {
		goto LABEL9
	}
	ns.FrameTimer(1, GetRandomMobYell)
LABEL9:
	goto LABEL7
LABEL2:
	if !(flag73 == true && ns.CurrentHealth(obj28) > 0) {
		goto LABEL10
	}
	ns.Chat(obj28, "War07A.scr:MobYell")
	ns.AggressionLevel(obj25, 0.83)
	ns.AggressionLevel(obj28, 0.83)
	ns.AggressionLevel(obj27, 0.83)
	ns.AggressionLevel(obj26, 0.83)
	ns.AggressionLevel(obj29, 0.83)
	ns.AggressionLevel(obj30, 0.83)
	ns.Attack(obj25, ns.GetHost())
	ns.Attack(obj28, ns.GetHost())
	ns.Attack(obj27, ns.GetHost())
	ns.Attack(obj26, ns.GetHost())
	ns.Attack(obj29, ns.GetHost())
	ns.Attack(obj30, ns.GetHost())
	flag80 = true
	goto LABEL11
LABEL10:
	if flag80 {
		goto LABEL11
	}
	ns.FrameTimer(1, GetRandomMobYell)
LABEL11:
	goto LABEL7
LABEL3:
	if !(flag76 == true && ns.CurrentHealth(obj26) > 0) {
		goto LABEL12
	}
	ns.Chat(obj26, "War07A.scr:MobYell")
	ns.AggressionLevel(obj25, 0.83)
	ns.AggressionLevel(obj28, 0.83)
	ns.AggressionLevel(obj27, 0.83)
	ns.AggressionLevel(obj26, 0.83)
	ns.AggressionLevel(obj29, 0.83)
	ns.AggressionLevel(obj30, 0.83)
	ns.Attack(obj25, ns.GetHost())
	ns.Attack(obj28, ns.GetHost())
	ns.Attack(obj27, ns.GetHost())
	ns.Attack(obj26, ns.GetHost())
	ns.Attack(obj29, ns.GetHost())
	ns.Attack(obj30, ns.GetHost())
	flag80 = true
	goto LABEL13
LABEL12:
	if flag80 {
		goto LABEL13
	}
	ns.FrameTimer(1, GetRandomMobYell)
LABEL13:
	goto LABEL7
LABEL4:
	if !(flag77 == true && ns.CurrentHealth(obj27) > 0) {
		goto LABEL14
	}
	ns.Chat(obj27, "War07A.scr:MobYell")
	ns.AggressionLevel(obj25, 0.83)
	ns.AggressionLevel(obj28, 0.83)
	ns.AggressionLevel(obj27, 0.83)
	ns.AggressionLevel(obj26, 0.83)
	ns.AggressionLevel(obj29, 0.83)
	ns.AggressionLevel(obj30, 0.83)
	ns.Attack(obj25, ns.GetHost())
	ns.Attack(obj28, ns.GetHost())
	ns.Attack(obj27, ns.GetHost())
	ns.Attack(obj26, ns.GetHost())
	ns.Attack(obj29, ns.GetHost())
	ns.Attack(obj30, ns.GetHost())
	flag80 = true
	goto LABEL15
LABEL14:
	if flag80 {
		goto LABEL15
	}
	ns.FrameTimer(1, GetRandomMobYell)
LABEL15:
	goto LABEL7
LABEL5:
	if !(flag78 == true && ns.CurrentHealth(obj29) > 0) {
		goto LABEL16
	}
	ns.Chat(obj29, "War07A.scr:MobYell")
	ns.AggressionLevel(obj25, 0.83)
	ns.AggressionLevel(obj28, 0.83)
	ns.AggressionLevel(obj27, 0.83)
	ns.AggressionLevel(obj26, 0.83)
	ns.AggressionLevel(obj29, 0.83)
	ns.AggressionLevel(obj30, 0.83)
	ns.Attack(obj25, ns.GetHost())
	ns.Attack(obj28, ns.GetHost())
	ns.Attack(obj27, ns.GetHost())
	ns.Attack(obj26, ns.GetHost())
	ns.Attack(obj29, ns.GetHost())
	ns.Attack(obj30, ns.GetHost())
	flag80 = true
	goto LABEL17
LABEL16:
	if flag80 {
		goto LABEL17
	}
	ns.FrameTimer(1, GetRandomMobYell)
LABEL17:
	goto LABEL7
LABEL6:
	if !(flag79 == true && ns.CurrentHealth(obj30) > 0) {
		goto LABEL18
	}
	ns.Chat(obj30, "War07A.scr:MobYell")
	ns.AggressionLevel(obj25, 0.83)
	ns.AggressionLevel(obj28, 0.83)
	ns.AggressionLevel(obj27, 0.83)
	ns.AggressionLevel(obj26, 0.83)
	ns.AggressionLevel(obj29, 0.83)
	ns.AggressionLevel(obj30, 0.83)
	ns.Attack(obj25, ns.GetHost())
	ns.Attack(obj28, ns.GetHost())
	ns.Attack(obj27, ns.GetHost())
	ns.Attack(obj26, ns.GetHost())
	ns.Attack(obj29, ns.GetHost())
	ns.Attack(obj30, ns.GetHost())
	flag80 = true
	goto LABEL19
LABEL18:
	if flag80 {
		goto LABEL19
	}
	ns.FrameTimer(1, GetRandomMobYell)
LABEL19:
	goto LABEL7
LABEL7:
	return
}
func AppleManRecognize() {
	ns.Chat(obj32, "War07A.scr:AppleManTalk01")
}
func PlayerPit() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.ObjectOn(obj52)
LABEL1:
	return
}
func MaxDialogStart() {
	var v0 int
	v0 = gvar148
	if v0 == gvar142 {
		goto LABEL1
	}
	if v0 == gvar143 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.TellStory(ns.HumanMaleEatFood, "War07a.scr:MaxOffer01")
	goto LABEL3
LABEL2:
	ns.TellStory(ns.HumanMaleEatFood, "War07A.scr:MaxWaiting")
	goto LABEL3
LABEL3:
	return
}
func MaxDialogEnd() {
	var (
		v0 int
		v1 int
		v2 int
		v3 int
		v4 int
		v5 int
	)
	v0 = 0
	v1 = 0
	v2 = 1
	v3 = 2
	v4 = 0
	v0 = ns.GetAnswer(obj21)
	v5 = v0
	if v5 == v4 {
		goto LABEL1
	}
	if v5 == v3 {
		goto LABEL2
	}
	if v5 == v2 {
		goto LABEL3
	}
	goto LABEL4
LABEL1:
	gvar148 = gvar143
	ResetMaxDialog()
	goto LABEL4
LABEL2:
	ns.SetDialog(obj21, ns.NORMAL, NullFunction, ResetMaxDialog)
	ns.TellStory(ns.HumanMaleEatFood, "War07A.scr:TurnMaxDown")
	gvar148 = gvar143
	goto LABEL4
LABEL3:
	v1 = ns.GetGold(ns.GetHost())
	if !(v1 >= 50000) {
		goto LABEL5
	}
	ns.SetDialog(obj21, ns.NORMAL, NullFunction, ResetMaxDialog)
	ns.TellStory(ns.HumanMaleEatFood, "War07A.scr:HaveGoldForMax")
	ns.ChangeGold(ns.GetHost(), -50000)
	gvar148 = gvar147
	ns.UnlockDoor(obj41)
	ns.UnlockDoor(obj42)
	goto LABEL4
	goto LABEL4
LABEL5:
	ns.SetDialog(obj21, ns.NORMAL, NullFunction, ResetMaxDialog)
	ns.TellStory(ns.HumanMaleEatFood, "War07A.scr:NoGoldForMax")
	gvar148 = gvar143
	goto LABEL4
LABEL4:
	return
}
func PriestGreetingStart() {
	if flag81 {
		goto LABEL1
	}
	ns.TellStory(ns.HumanMaleEatFood, "War07A.scr:PriestTalk01")
	goto LABEL2
LABEL1:
	ns.TellStory(ns.HumanMaleEatFood, "War07A.scr:PriestTalk02")
LABEL2:
	return
}
func PriestGreetingEnd() {
	ns.SetDialog(obj11, ns.NORMAL, PriestGreetingStart, PriestGreetingEnd)
}
func UndertakerTalkStart() {
	ns.TellStory(ns.HumanMaleEatFood, "War07A.scr:UndertakerTalk01")
}
func UndertakerTalkEnd() {
	ns.SetDialog(obj33, ns.NORMAL, UndertakerTalkStart, UndertakerTalkEnd)
}
func DorianRecognize() {
	var v0 int
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	flag83 = true
	if flag82 {
		goto LABEL1
	}
	v0 = gvar102
	if v0 == gvar127 {
		goto LABEL2
	}
	if v0 == gvar128 {
		goto LABEL3
	}
	goto LABEL1
LABEL2:
	ns.CreatureFollow(obj28, ns.GetHost())
	if flag73 {
		goto LABEL4
	}
	ivar103 += 1
	flag73 = true
LABEL4:
	if ivar103 != 3 {
		goto LABEL5
	}
	gvar102 = gvar128
	GetRandomMobYell()
LABEL5:
	goto LABEL1
LABEL3:
	flag81 = false
	ns.Attack(obj28, ns.GetHost())
	goto LABEL1
LABEL1:
	return
}
func GrunbarRecognize() {
	var v0 int
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	flag85 = true
	if flag82 {
		goto LABEL1
	}
	v0 = gvar102
	if v0 == gvar127 {
		goto LABEL2
	}
	if v0 == gvar128 {
		goto LABEL3
	}
	goto LABEL1
LABEL2:
	ns.CreatureFollow(obj27, ns.GetHost())
	if flag77 {
		goto LABEL4
	}
	ivar103 += 1
	flag77 = true
LABEL4:
	if ivar103 != 3 {
		goto LABEL5
	}
	gvar102 = gvar128
	GetRandomMobYell()
LABEL5:
	goto LABEL1
LABEL3:
	flag81 = false
	ns.Attack(obj27, ns.GetHost())
	goto LABEL1
LABEL1:
	return
}
func JorganRecognize() {
	var v0 int
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	flag86 = true
	if flag82 {
		goto LABEL1
	}
	v0 = gvar102
	if v0 == gvar127 {
		goto LABEL2
	}
	if v0 == gvar128 {
		goto LABEL3
	}
	goto LABEL1
LABEL2:
	ns.CreatureFollow(obj26, ns.GetHost())
	if flag76 {
		goto LABEL4
	}
	ivar103 += 1
	flag76 = true
LABEL4:
	if ivar103 != 3 {
		goto LABEL5
	}
	gvar102 = gvar128
	GetRandomMobYell()
LABEL5:
	goto LABEL1
LABEL3:
	flag81 = false
	ns.Attack(obj26, ns.GetHost())
	goto LABEL1
LABEL1:
	return
}
func AlbiRecognize() {
	var v0 int
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	flag84 = true
	if flag82 {
		goto LABEL1
	}
	v0 = gvar102
	if v0 == gvar127 {
		goto LABEL2
	}
	if v0 == gvar128 {
		goto LABEL3
	}
	goto LABEL1
LABEL2:
	ns.CreatureFollow(obj25, ns.GetHost())
	if flag75 {
		goto LABEL4
	}
	ivar103 += 1
	flag75 = true
LABEL4:
	if ivar103 != 3 {
		goto LABEL5
	}
	gvar102 = gvar128
	GetRandomMobYell()
LABEL5:
	goto LABEL1
LABEL3:
	flag81 = false
	ns.Attack(obj25, ns.GetHost())
	goto LABEL1
LABEL1:
	return
}
func Civvy3Recognize() {
	var v0 int
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	flag87 = true
	if flag82 {
		goto LABEL1
	}
	v0 = gvar102
	if v0 == gvar127 {
		goto LABEL2
	}
	if v0 == gvar128 {
		goto LABEL3
	}
	goto LABEL1
LABEL2:
	ns.CreatureFollow(obj29, ns.GetHost())
	if flag78 {
		goto LABEL4
	}
	ivar103 += 1
	flag78 = true
LABEL4:
	if ivar103 != 3 {
		goto LABEL5
	}
	gvar102 = gvar128
	GetRandomMobYell()
LABEL5:
	goto LABEL1
LABEL3:
	flag81 = false
	ns.Attack(obj29, ns.GetHost())
	goto LABEL1
LABEL1:
	return
}
func Civvy4Recognize() {
	var v0 int
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	flag88 = true
	if flag82 {
		goto LABEL1
	}
	v0 = gvar102
	if v0 == gvar127 {
		goto LABEL2
	}
	if v0 == gvar128 {
		goto LABEL3
	}
	goto LABEL1
LABEL2:
	ns.CreatureFollow(obj30, ns.GetHost())
	if flag79 {
		goto LABEL4
	}
	ivar103 += 1
	flag79 = true
LABEL4:
	if ivar103 != 3 {
		goto LABEL5
	}
	gvar102 = gvar128
	GetRandomMobYell()
LABEL5:
	goto LABEL1
LABEL3:
	flag81 = false
	ns.Attack(obj30, ns.GetHost())
	goto LABEL1
LABEL1:
	return
}
func AlbiWaiting() {
	if !ns.IsCaller(obj25) {
		goto LABEL1
	}
	if !(gvar102 != gvar126 || gvar94 == 1) {
		goto LABEL1
	}
	if !flag82 {
		goto LABEL1
	}
	flag81 = true
	ns.CreatureIdle(obj25)
	ns.AggressionLevel(obj25, 0.16)
	ns.CreatureGuard(obj25, ns.GetObjectX(obj25), ns.GetObjectY(obj25), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
LABEL1:
	return
}
func DorianWaiting() {
	if !ns.IsCaller(obj28) {
		goto LABEL1
	}
	if !(gvar102 != gvar126 || gvar95 == 1) {
		goto LABEL1
	}
	if !flag82 {
		goto LABEL1
	}
	flag81 = true
	ns.CreatureIdle(obj28)
	ns.AggressionLevel(obj28, 0.16)
	ns.CreatureGuard(obj28, ns.GetObjectX(obj28), ns.GetObjectY(obj28), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
LABEL1:
	return
}
func GrunbarWaiting() {
	if !ns.IsCaller(obj27) {
		goto LABEL1
	}
	if !(gvar102 != gvar126 || gvar96 == 1) {
		goto LABEL1
	}
	if !flag82 {
		goto LABEL1
	}
	flag81 = true
	ns.CreatureIdle(obj27)
	ns.AggressionLevel(obj27, 0.16)
	ns.CreatureGuard(obj27, ns.GetObjectX(obj27), ns.GetObjectY(obj27), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
LABEL1:
	return
}
func JorganWaiting() {
	if !ns.IsCaller(obj26) {
		goto LABEL1
	}
	if !(gvar102 != gvar126 || gvar97 == 1) {
		goto LABEL1
	}
	if !flag82 {
		goto LABEL1
	}
	flag81 = true
	ns.CreatureIdle(obj26)
	ns.AggressionLevel(obj26, 0.16)
	ns.CreatureGuard(obj26, ns.GetObjectX(obj26), ns.GetObjectY(obj26), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
LABEL1:
	return
}
func Civvy3Waiting() {
	if !ns.IsCaller(obj29) {
		goto LABEL1
	}
	if !(gvar102 != gvar126 || gvar98 == 1) {
		goto LABEL1
	}
	if !flag82 {
		goto LABEL1
	}
	flag81 = true
	ns.CreatureIdle(obj29)
	ns.AggressionLevel(obj29, 0.16)
	ns.CreatureGuard(obj29, ns.GetObjectX(obj29), ns.GetObjectY(obj29), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
LABEL1:
	return
}
func Civvy4Waiting() {
	if !ns.IsCaller(obj30) {
		goto LABEL1
	}
	if !(gvar102 != gvar126 || gvar99 == 1) {
		goto LABEL1
	}
	if !flag82 {
		goto LABEL1
	}
	flag81 = true
	ns.CreatureIdle(obj30)
	ns.AggressionLevel(obj30, 0.16)
	ns.CreatureGuard(obj30, ns.GetObjectX(obj30), ns.GetObjectY(obj30), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
LABEL1:
	return
}
func DorianLoseSight() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	flag83 = false
LABEL1:
	return
}
func AlbiLoseSight() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	flag84 = false
LABEL1:
	return
}
func JorganLoseSight() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	flag86 = false
LABEL1:
	return
}
func GrunbarLoseSight() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	flag85 = false
LABEL1:
	return
}
func Civvy3LoseSight() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	flag87 = false
LABEL1:
	return
}
func Civvy4LoseSight() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	flag88 = false
LABEL1:
	return
}
func MlurghWelcomeStart() {
	ns.TellStory(ns.HumanMaleEatFood, "War07A.scr:MlurghTalk01")
}
func MlurghWelcomeEnd() {
	ns.SetDialog(obj34, ns.NORMAL, MlurghRumorStart, MlurghRumorEnd)
}
func MlurghRumorStart() {
	ns.TellStory(ns.HumanMaleEatFood, "War07A.scr:MlurghTalk02")
}
func MlurghRumorEnd() {
	ns.SetDialog(obj34, ns.NORMAL, MlurghRumorStart, MlurghRumorEnd)
}
func FreezePhim() {
	if !ns.IsCaller(obj9) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.MoveObject(obj10, fvar112, fvar113)
	ns.Frozen(obj9, true)
	ns.Frozen(obj10, true)
LABEL1:
	return
}
func FreezeMage() {
	if !ns.IsCaller(obj12) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.MoveObject(obj13, fvar114, fvar115)
	ns.Frozen(obj12, true)
	ns.Frozen(obj13, true)
LABEL1:
	return
}
func FreezeKin() {
	if !ns.IsCaller(obj16) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.MoveObject(obj17, fvar116, fvar117)
	ns.Frozen(obj16, true)
	ns.Frozen(obj17, true)
LABEL1:
	return
}
func FreezeBB() {
	if !ns.IsCaller(obj14) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.MoveObject(obj15, fvar110, fvar111)
	ns.Frozen(obj14, true)
	ns.Frozen(obj15, true)
LABEL1:
	return
}
func WardenRecognize() {
	if !flag68 {
		goto LABEL1
	}
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.AggressionLevel(obj22, 0.83)
	ns.CreatureHunt(obj22)
	ns.Attack(obj22, ns.GetHost())
LABEL1:
	if !(flag72 == false && flag68 == false && flag70 == false) {
		goto LABEL2
	}
	if !ns.IsCaller(obj31) {
		goto LABEL3
	}
	ns.AggressionLevel(obj22, 0.16)
	ns.CreatureFollow(obj22, ns.GetHost())
LABEL3:
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL4
	}
	gvar101 = gvar124
	ns.AggressionLevel(obj22, 0.16)
	ns.CreatureFollow(obj22, ns.GetHost())
	ns.SetDialog(obj22, ns.YESNO, WardenArrestStart, WardenArrestEnd)
	StartWardenTalk()
LABEL4:
	goto LABEL5
LABEL2:
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL5
	}
	if !flag72 {
		goto LABEL6
	}
	if flag71 {
		goto LABEL7
	}
	ns.Chat(obj22, "War07A.scr:WardenTalk02")
	ns.AggressionLevel(obj22, 0.83)
	ns.CreatureHunt(obj22)
	ns.Attack(obj22, ns.GetHost())
	gvar101 = gvar125
	ns.LockDoor(obj43)
LABEL7:
	if !flag71 {
		goto LABEL8
	}
	ns.AggressionLevel(obj22, 0.16)
LABEL8:
	goto LABEL5
LABEL6:
	ns.AggressionLevel(obj22, 0.83)
	ns.CreatureHunt(obj22)
	ns.Attack(obj22, ns.GetHost())
LABEL5:
	return
}
func WardenArrestStart() {
	ns.Frozen(obj31, true)
	ns.Frozen(obj25, true)
	ns.Frozen(obj28, true)
	ns.Frozen(obj27, true)
	ns.Frozen(obj26, true)
	ns.Frozen(obj29, true)
	ns.Frozen(obj30, true)
	ns.Frozen(ns.GetHost(), true)
	ns.CreatureIdle(obj31)
	ns.CreatureIdle(obj25)
	ns.CreatureIdle(obj28)
	ns.CreatureIdle(obj27)
	ns.CreatureIdle(obj26)
	ns.CreatureIdle(obj29)
	ns.CreatureIdle(obj30)
	ns.CreatureIdle(ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj22)
	ns.TellStory(ns.HumanMaleEatFood, "War07A.scr:WardenPreArrest04")
}
func WardenArrestEnd() {
	var (
		v0 int
		v1 int
	)
	v0 = 0
	ns.Frozen(ns.GetHost(), false)
	ns.Frozen(obj31, false)
	ns.Frozen(obj25, false)
	ns.Frozen(obj26, false)
	ns.Frozen(obj27, false)
	ns.Frozen(obj28, false)
	ns.Frozen(obj29, false)
	ns.Frozen(obj30, false)
	ns.Wander(obj25)
	ns.Wander(obj28)
	ns.Wander(obj27)
	ns.Wander(obj26)
	ns.Wander(obj29)
	ns.Wander(obj30)
	v0 = ns.GetAnswer(obj22)
	v1 = v0
	if v1 == 0 {
		goto LABEL1
	}
	if v1 == 1 {
		goto LABEL2
	}
	if v1 == 2 {
		goto LABEL3
	}
	goto LABEL4
LABEL1:
	gvar101 = gvar125
	ns.LockDoor(obj43)
	ns.AggressionLevel(obj22, 0.83)
	ns.AggressionLevel(obj31, 0.5)
	ns.Attack(obj22, ns.GetHost())
	ns.Attack(obj31, ns.GetHost())
	ns.CreatureGuard(obj23, 3622, 5460, 3737, 5579, 0)
	ns.CancelDialog(obj23)
	goto LABEL4
LABEL2:
	ns.Blind()
	ns.AggressionLevel(obj22, 0)
	ns.AggressionLevel(obj31, 0.16)
	ns.FrameTimer(60, PutPeopleInJail)
	goto LABEL4
LABEL3:
	gvar101 = gvar125
	ns.LockDoor(obj43)
	ns.AggressionLevel(obj22, 0.83)
	ns.AggressionLevel(obj31, 0.5)
	ns.Attack(obj22, ns.GetHost())
	ns.Attack(obj31, ns.GetHost())
	ns.CreatureGuard(obj23, 3622, 5460, 3737, 5579, 0)
	ns.CancelDialog(obj23)
	goto LABEL4
LABEL4:
	ns.CancelDialog(obj22)
}
func WardenUpsetStart() {
	ns.TellStory(ns.HumanMaleEatFood, "War07A.scr:WardenTalk01")
}
func WardenUpsetEnd() {
	ns.CancelDialog(obj22)
	ns.Wander(obj22)
	ns.FrameTimer(1, MorganWhisper)
}
func WardenJailCellTalk() {
	if flag65 {
		goto LABEL1
	}
	ns.AggressionLevel(obj22, 0)
	flag65 = true
	ns.FrameTimer(60, WardenJailCellTalk)
	goto LABEL2
LABEL1:
	ns.AggressionLevel(obj22, 0)
	ns.SetDialog(obj22, ns.NORMAL, WardenUpsetStart, WardenUpsetEnd)
	ns.LookAtObject(obj22, ns.GetHost())
	ns.StartDialog(obj22, ns.GetHost())
LABEL2:
	return
}
func WardenBack() {
	ns.ObjectOn(obj22)
	ns.LookAtObject(obj22, obj46)
	flag71 = false
	ns.AggressionLevel(obj22, 0.16)
	ns.ObjectOff(obj50)
}
func WardenDied() {
	if !ns.IsLocked(obj43) {
		goto LABEL1
	}
	ns.UnlockDoor(obj43)
LABEL1:
	if !(ns.CurrentHealth(obj31) > 0) {
		goto LABEL2
	}
	flag93 = true
	ns.GoBackHome(obj31)
LABEL2:
	if flag69 {
		goto LABEL3
	}
	if !flag66 {
		goto LABEL4
	}
	ns.CreatureGuard(obj24, ns.GetObjectX(obj24), ns.GetObjectY(obj24), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	ns.SetDialog(obj24, ns.NORMAL, FriendWaitStart, FriendWaitEnd)
LABEL4:
	flag69 = true
	ns.CancelDialog(obj22)
	if !(flag84 == true || flag85 == true || flag86 == true || flag83 == true || flag87 == true || flag88 == true) {
		goto LABEL3
	}
	if gvar102 == gvar128 {
		goto LABEL3
	}
	gvar102 = gvar128
	ns.AggressionLevel(obj28, 0.83)
	ns.AggressionLevel(obj27, 0.83)
	ns.AggressionLevel(obj25, 0.83)
	ns.AggressionLevel(obj26, 0.83)
	ns.AggressionLevel(obj29, 0.83)
	ns.AggressionLevel(obj30, 0.83)
	ns.MoveObject(obj9, ns.GetWaypointX(wp137), ns.GetWaypointY(wp137))
	ns.MoveObject(obj12, ns.GetWaypointX(wp139), ns.GetWaypointY(wp139))
	ns.MoveObject(obj16, ns.GetWaypointX(wp138), ns.GetWaypointY(wp138))
	ns.MoveObject(obj14, ns.GetWaypointX(wp136), ns.GetWaypointY(wp136))
LABEL3:
	return
}
func WardenMad() {
	ns.Chat(obj22, "War07A.scr:WardenTalk04")
}
func StartWardenTalk() {
	if flag70 {
		goto LABEL1
	}
	if !(ns.Distance(ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), ns.GetObjectX(obj22), ns.GetObjectY(obj22)) > 120) {
		goto LABEL2
	}
	ns.FrameTimer(10, StartWardenTalk)
	goto LABEL1
LABEL2:
	if !ns.IsTrading() {
		goto LABEL3
	}
	ns.FrameTimer(10, StartWardenTalk)
	goto LABEL1
LABEL3:
	if !(ns.CurrentHealth(ns.GetHost()) > 0) {
		goto LABEL1
	}
	flag70 = true
	ns.StartDialog(obj22, ns.GetHost())
LABEL1:
	return
}
func WardenHurt() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.AggressionLevel(obj22, 0.83)
	gvar101 = gvar125
	ns.CancelDialog(obj22)
LABEL1:
	return
}
func LockBackDoor() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.LockDoor(obj46)
LABEL1:
	return
}
func MorganWhisper() {
	ns.Chat(obj23, "War07A.scr:MorganTalk01")
}
func MorganHintStart() {
	ns.TellStory(ns.HumanMaleEatFood, "War07A.scr:MorganTalk02")
}
func MorganHintEnd() {
	ns.CancelDialog(obj23)
	ns.JournalEntry(ns.GetHost(), "FindTunnel", 2)
	ns.PrintToAll("Con01a:NewJournalEntry")
	ns.Walk(obj23, 3633, 5472)
}
func MorganRequestStart() {
	ns.TellStory(ns.HumanMaleEatFood, "War07A.scr:MorganTalk04")
}
func MorganRequestEnd() {
	ns.CancelDialog(obj23)
	ns.JournalEntry(ns.GetHost(), "RescueMorgan", 2)
}
func OpenEscapeWall() {
	if !ns.IsCaller(obj23) {
		goto LABEL1
	}
	ns.WallOpen(ns.Wall(166, 220))
LABEL1:
	return
}
func MorganThanksTrigger() {
	if !ns.IsCaller(obj23) {
		goto LABEL1
	}
	ns.CreatureGuard(obj23, ns.GetObjectX(obj23), ns.GetObjectY(obj23), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	ns.SetDialog(obj23, ns.NORMAL, MorganThanksStart, MorganThanksEnd)
LABEL1:
	if !flag89 {
		goto LABEL2
	}
	ns.Frozen(ns.GetHost(), true)
	ns.Frozen(obj22, true)
	ns.Frozen(obj31, true)
	ns.CreatureIdle(ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj23)
	ns.StartDialog(obj23, ns.GetHost())
	ns.ObjectOff(obj51)
	goto LABEL3
LABEL2:
	if !(flag91 == true && flag90 == true) {
		goto LABEL4
	}
	ns.Frozen(ns.GetHost(), true)
	ns.Frozen(obj22, true)
	ns.Frozen(obj31, true)
	ns.CreatureIdle(ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj23)
	ns.StartDialog(obj23, ns.GetHost())
	ns.ObjectOff(obj51)
	ns.ObjectGroupOff(gvar53)
	goto LABEL3
LABEL4:
	ns.FrameTimer(10, MorganThanksTrigger)
LABEL3:
	return
}
func MorganThanksStart() {
	if !flag66 {
		goto LABEL1
	}
	ns.MoveObject(obj24, 3756, 4777)
	ns.CreatureFollow(obj24, ns.GetHost())
	goto LABEL2
LABEL1:
	ns.LockDoor(obj39)
LABEL2:
	ns.TellStory(ns.HumanMaleEatFood, "War07a.scr:MorganTalk06")
	ns.SetDialog(obj24, ns.NORMAL, GiveShieldStart, GiveShieldEnd)
}
func MorganThanksEnd() {
	ns.CancelDialog(obj23)
	ns.Pickup(ns.GetHost(), obj56)
	ns.Pickup(ns.GetHost(), obj57)
	ns.GiveXp(ns.GetHost(), 500)
	ns.PrintToAll("GeneralPrint:GainedKey")
	if !flag66 {
		goto LABEL1
	}
	ns.JournalEdit(ns.GetHost(), "RescueMorgan", 4)
	ns.PrintToAll("Con02a:ObjectiveComplete")
	ns.StartDialog(obj24, ns.GetHost())
	goto LABEL2
LABEL1:
	ns.WideScreen(false)
	ns.Frozen(ns.GetHost(), false)
	ns.Frozen(obj22, false)
	ns.Frozen(obj31, false)
	flag92 = true
LABEL2:
	return
}
func MorganContact() {
	r0 := ns.IsAttackedBy(obj23, ns.GetCaller())
	if !r0 {
		goto LABEL1
	}
	if !flag92 {
		goto LABEL1
	}
	ns.AggressionLevel(obj23, 0.83)
LABEL1:
	return
}
func MorganFriendStart() {
	ns.TellStory(ns.HumanMaleEatFood, "War07A.scr:MorganFriendTalk01")
}
func MorganFriendEnd() {
	ns.JournalEdit(ns.GetHost(), "FindTunnel", 4)
	ns.PrintToAll("Con02a:ObjectiveComplete")
	ns.JournalEntry(ns.GetHost(), "RescueMorgan", 2)
	ns.PrintToAll("Con01a:NewJournalEntry")
	flag66 = true
	ns.CancelDialog(obj24)
}
func UnlockGearDoor() {
	if !ns.IsLocked(obj39) {
		goto LABEL1
	}
	ns.UnlockDoor(obj39)
LABEL1:
	return
}
func CheckMorganDoor() {
	if !ns.IsLocked(obj45) {
		goto LABEL1
	}
	ns.FrameTimer(10, CheckMorganDoor)
	goto LABEL2
LABEL1:
	ns.CancelDialog(obj23)
	ns.ObjectOff(obj49)
	ns.LookAtObject(obj23, ns.GetHost())
	ns.Chat(obj23, "War07A.scr:MorganTalk05")
	ns.WayPointGroupOn(gvar134)
	ns.Wander(obj23)
LABEL2:
	return
}
func SeeMorgan() {
	if !ns.IsCaller(obj23) {
		goto LABEL1
	}
	if flag67 {
		goto LABEL1
	}
	flag67 = true
	ns.SetDialog(obj24, ns.NORMAL, GiveShieldStart, GiveShieldEnd)
LABEL1:
	return
}
func GiveShieldStart() {
	ns.Frozen(ns.GetHost(), true)
	ns.Frozen(obj22, true)
	ns.TellStory(ns.HumanMaleEatFood, "War07A.scr:MorganFriendTalk02")
}
func GiveShieldEnd() {
	ns.Frozen(ns.GetHost(), false)
	ns.Frozen(obj22, false)
	ns.Frozen(obj31, false)
	ns.WideScreen(false)
	flag92 = true
	ns.CreatureGuard(obj23, ns.GetObjectX(obj23), ns.GetObjectY(obj23), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	ns.CreatureGuard(obj24, ns.GetObjectX(obj24), ns.GetObjectY(obj24), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	ns.CancelDialog(obj24)
	ns.Pickup(ns.GetHost(), obj55)
	ns.PrintToAll("GeneralPrint:GainedItem")
	ns.JournalEdit(ns.GetHost(), "RescueMorgan", 4)
	ns.PrintToAll("Con02a:ObjectiveComplete")
}
func FriendWaitStart() {
	ns.TellStory(ns.HumanMaleEatFood, "War07A.scr:MorganFriendTalk04")
}
func FriendWaitEnd() {
	ns.SetDialog(obj24, ns.NORMAL, FriendWaitStart, FriendWaitEnd)
}
func LookForPlayer() {
	if ns.IsTalking() {
		goto LABEL1
	}
	if !(flag89 == true || flag90 == true) {
		goto LABEL1
	}
	ns.CreatureGuard(obj24, ns.GetObjectX(obj24), ns.GetObjectY(obj24), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	ns.StartDialog(obj24, ns.GetHost())
	return
LABEL1:
	ns.FrameTimer(10, LookForPlayer)
}
func RewardTrigger() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	if ns.IsLocked(obj45) {
		goto LABEL1
	}
	ns.ObjectGroupOff(gvar53)
	ns.Frozen(ns.GetHost(), true)
	ns.CreatureIdle(ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj51)
	ns.WideScreen(true)
	flag89 = true
LABEL1:
	return
}
func PlayerExitRoom() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	flag90 = false
LABEL1:
	if !ns.IsCaller(obj23) {
		goto LABEL2
	}
	ns.AggressionLevel(obj23, 0.16)
	ns.CreatureGuard(obj23, ns.GetWaypointX(wp140), ns.GetWaypointY(wp140), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
LABEL2:
	if !ns.IsCaller(obj24) {
		goto LABEL3
	}
	ns.AggressionLevel(obj24, 0.16)
	ns.CreatureGuard(obj24, ns.GetWaypointX(wp141), ns.GetWaypointY(wp141), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
LABEL3:
	return
}
func OpenTowerGate() {
	var v0 ns.ObjectID
	if !(!ns.IsLocked(obj47) || !ns.IsLocked(obj48)) {
		goto LABEL1
	}
	v0 = ns.GetLastItem(ns.GetHost())
	for {
		if v0 == 0 {
			goto LABEL1
		}
		if v0 != ns.Object("GoldKey2") {
			goto LABEL2
		}
		ns.Delete(v0)
		goto LABEL1
	LABEL2:
		if v0 != ns.Object("GoldKey1") {
			goto LABEL3
		}
		ns.Delete(v0)
		ns.ObjectOff(ns.GetTrigger())
		goto LABEL1
	LABEL3:
		v0 = ns.GetPreviousItem(v0)
	}
LABEL1:
	return
}
func BoothStart() {
	ns.LookAtObject(obj36, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "War03b:T4Post")
}
func BoothEnd() {
}
func NewMaxStart() {
	ns.DestroyEveryChat()
	ns.TellStory(ns.HumanMaleEatFood, "War07A.scr:MaxOffer01")
}
func NewMaxEnd() {
}
func IndexPlayerInventory() {
}
func WarTrigger() {
	ns.ObjectOff(ns.GetTrigger())
	ns.UnlockDoor(obj41)
	ns.UnlockDoor(obj42)
}
func Injured() {
	if flag82 {
		goto LABEL1
	}
	ns.AggressionLevel(ns.GetTrigger(), 0.83)
LABEL1:
	if !ns.IsTrigger(obj25) {
		goto LABEL2
	}
	gvar94 = 1
	if flag82 {
		goto LABEL3
	}
	if !(flag86 == true || flag83 == true || flag85 == true || flag87 == true || flag88 == true) {
		goto LABEL3
	}
	gvar102 = gvar128
	ns.AggressionLevel(obj25, 0.83)
	ns.AggressionLevel(obj26, 0.83)
	ns.AggressionLevel(obj28, 0.83)
	ns.AggressionLevel(obj27, 0.83)
	ns.AggressionLevel(obj29, 0.83)
	ns.AggressionLevel(obj30, 0.83)
LABEL3:
	goto LABEL4
LABEL2:
	if !ns.IsTrigger(obj26) {
		goto LABEL5
	}
	gvar97 = 1
	if flag82 {
		goto LABEL6
	}
	if !(flag84 == true || flag83 == true || flag85 == true || flag87 == true || flag88 == true) {
		goto LABEL6
	}
	gvar102 = gvar128
	ns.AggressionLevel(obj25, 0.83)
	ns.AggressionLevel(obj26, 0.83)
	ns.AggressionLevel(obj28, 0.83)
	ns.AggressionLevel(obj27, 0.83)
	ns.AggressionLevel(obj29, 0.83)
	ns.AggressionLevel(obj30, 0.83)
LABEL6:
	goto LABEL4
LABEL5:
	if !ns.IsTrigger(obj28) {
		goto LABEL7
	}
	gvar95 = 1
	if flag82 {
		goto LABEL8
	}
	if !(flag86 == true || flag84 == true || flag85 == true || flag87 == true || flag88 == true) {
		goto LABEL8
	}
	gvar102 = gvar128
	ns.AggressionLevel(obj25, 0.83)
	ns.AggressionLevel(obj26, 0.83)
	ns.AggressionLevel(obj28, 0.83)
	ns.AggressionLevel(obj27, 0.83)
	ns.AggressionLevel(obj29, 0.83)
	ns.AggressionLevel(obj30, 0.83)
LABEL8:
	goto LABEL4
LABEL7:
	if !ns.IsTrigger(obj27) {
		goto LABEL9
	}
	gvar96 = 1
	if flag82 {
		goto LABEL10
	}
	if !(flag86 == true || flag83 == true || flag84 == true || flag87 == true || flag88 == true) {
		goto LABEL10
	}
	gvar102 = gvar128
	ns.AggressionLevel(obj25, 0.83)
	ns.AggressionLevel(obj26, 0.83)
	ns.AggressionLevel(obj28, 0.83)
	ns.AggressionLevel(obj27, 0.83)
	ns.AggressionLevel(obj29, 0.83)
	ns.AggressionLevel(obj30, 0.83)
LABEL10:
	goto LABEL4
LABEL9:
	if !ns.IsTrigger(obj29) {
		goto LABEL11
	}
	gvar98 = 1
	if flag82 {
		goto LABEL12
	}
	if !(flag86 == true || flag83 == true || flag85 == true || flag84 == true || flag88 == true) {
		goto LABEL12
	}
	gvar102 = gvar128
	ns.AggressionLevel(obj25, 0.83)
	ns.AggressionLevel(obj26, 0.83)
	ns.AggressionLevel(obj28, 0.83)
	ns.AggressionLevel(obj27, 0.83)
	ns.AggressionLevel(obj29, 0.83)
	ns.AggressionLevel(obj30, 0.83)
LABEL12:
	goto LABEL4
LABEL11:
	if !ns.IsTrigger(obj30) {
		goto LABEL4
	}
	gvar99 = 1
	if flag82 {
		goto LABEL4
	}
	if !(flag86 == true || flag83 == true || flag85 == true || flag87 == true || flag84 == true) {
		goto LABEL4
	}
	gvar102 = gvar128
	ns.AggressionLevel(obj25, 0.83)
	ns.AggressionLevel(obj26, 0.83)
	ns.AggressionLevel(obj28, 0.83)
	ns.AggressionLevel(obj27, 0.83)
	ns.AggressionLevel(obj29, 0.83)
	ns.AggressionLevel(obj30, 0.83)
LABEL4:
	return
}
func Killed() {
	if flag82 {
		goto LABEL1
	}
	ns.AggressionLevel(ns.GetTrigger(), 0.83)
LABEL1:
	if !ns.IsTrigger(obj25) {
		goto LABEL2
	}
	if flag82 {
		goto LABEL3
	}
	if !(flag86 == true || flag83 == true || flag85 == true || flag87 == true || flag88 == true) {
		goto LABEL3
	}
	gvar102 = gvar128
	ns.AggressionLevel(obj25, 0.83)
	ns.AggressionLevel(obj26, 0.83)
	ns.AggressionLevel(obj28, 0.83)
	ns.AggressionLevel(obj27, 0.83)
	ns.AggressionLevel(obj29, 0.83)
	ns.AggressionLevel(obj30, 0.83)
	ns.MoveObject(obj9, ns.GetWaypointX(wp137), ns.GetWaypointY(wp137))
	ns.MoveObject(obj12, ns.GetWaypointX(wp139), ns.GetWaypointY(wp139))
	ns.MoveObject(obj16, ns.GetWaypointX(wp138), ns.GetWaypointY(wp138))
	ns.MoveObject(obj14, ns.GetWaypointX(wp136), ns.GetWaypointY(wp136))
LABEL3:
	goto LABEL4
LABEL2:
	if !ns.IsTrigger(obj26) {
		goto LABEL5
	}
	if flag82 {
		goto LABEL6
	}
	if !(flag84 == true || flag83 == true || flag85 == true || flag87 == true || flag88 == true) {
		goto LABEL6
	}
	gvar102 = gvar128
	ns.AggressionLevel(obj25, 0.83)
	ns.AggressionLevel(obj26, 0.83)
	ns.AggressionLevel(obj28, 0.83)
	ns.AggressionLevel(obj27, 0.83)
	ns.AggressionLevel(obj29, 0.83)
	ns.AggressionLevel(obj30, 0.83)
	ns.MoveObject(obj9, ns.GetWaypointX(wp137), ns.GetWaypointY(wp137))
	ns.MoveObject(obj12, ns.GetWaypointX(wp139), ns.GetWaypointY(wp139))
	ns.MoveObject(obj16, ns.GetWaypointX(wp138), ns.GetWaypointY(wp138))
	ns.MoveObject(obj14, ns.GetWaypointX(wp136), ns.GetWaypointY(wp136))
LABEL6:
	goto LABEL4
LABEL5:
	if !ns.IsTrigger(obj28) {
		goto LABEL7
	}
	if flag82 {
		goto LABEL8
	}
	if !(flag86 == true || flag84 == true || flag85 == true || flag87 == true || flag88 == true) {
		goto LABEL8
	}
	gvar102 = gvar128
	ns.AggressionLevel(obj25, 0.83)
	ns.AggressionLevel(obj26, 0.83)
	ns.AggressionLevel(obj28, 0.83)
	ns.AggressionLevel(obj27, 0.83)
	ns.AggressionLevel(obj29, 0.83)
	ns.AggressionLevel(obj30, 0.83)
	ns.MoveObject(obj9, ns.GetWaypointX(wp137), ns.GetWaypointY(wp137))
	ns.MoveObject(obj12, ns.GetWaypointX(wp139), ns.GetWaypointY(wp139))
	ns.MoveObject(obj16, ns.GetWaypointX(wp138), ns.GetWaypointY(wp138))
	ns.MoveObject(obj14, ns.GetWaypointX(wp136), ns.GetWaypointY(wp136))
LABEL8:
	goto LABEL4
LABEL7:
	if !ns.IsTrigger(obj27) {
		goto LABEL9
	}
	if flag82 {
		goto LABEL10
	}
	if !(flag86 == true || flag83 == true || flag84 == true || flag87 == true || flag88 == true) {
		goto LABEL10
	}
	gvar102 = gvar128
	ns.AggressionLevel(obj25, 0.83)
	ns.AggressionLevel(obj26, 0.83)
	ns.AggressionLevel(obj28, 0.83)
	ns.AggressionLevel(obj27, 0.83)
	ns.AggressionLevel(obj29, 0.83)
	ns.AggressionLevel(obj30, 0.83)
	ns.MoveObject(obj9, ns.GetWaypointX(wp137), ns.GetWaypointY(wp137))
	ns.MoveObject(obj12, ns.GetWaypointX(wp139), ns.GetWaypointY(wp139))
	ns.MoveObject(obj16, ns.GetWaypointX(wp138), ns.GetWaypointY(wp138))
	ns.MoveObject(obj14, ns.GetWaypointX(wp136), ns.GetWaypointY(wp136))
LABEL10:
	goto LABEL4
LABEL9:
	if !ns.IsTrigger(obj29) {
		goto LABEL11
	}
	if flag82 {
		goto LABEL12
	}
	if !(flag86 == true || flag83 == true || flag85 == true || flag84 == true || flag88 == true) {
		goto LABEL12
	}
	gvar102 = gvar128
	ns.AggressionLevel(obj25, 0.83)
	ns.AggressionLevel(obj26, 0.83)
	ns.AggressionLevel(obj28, 0.83)
	ns.AggressionLevel(obj27, 0.83)
	ns.AggressionLevel(obj29, 0.83)
	ns.AggressionLevel(obj30, 0.83)
	ns.MoveObject(obj9, ns.GetWaypointX(wp137), ns.GetWaypointY(wp137))
	ns.MoveObject(obj12, ns.GetWaypointX(wp139), ns.GetWaypointY(wp139))
	ns.MoveObject(obj16, ns.GetWaypointX(wp138), ns.GetWaypointY(wp138))
	ns.MoveObject(obj14, ns.GetWaypointX(wp136), ns.GetWaypointY(wp136))
LABEL12:
	goto LABEL4
LABEL11:
	if !ns.IsTrigger(obj30) {
		goto LABEL4
	}
	if flag82 {
		goto LABEL4
	}
	if !(flag86 == true || flag83 == true || flag85 == true || flag87 == true || flag84 == true) {
		goto LABEL4
	}
	gvar102 = gvar128
	ns.AggressionLevel(obj25, 0.83)
	ns.AggressionLevel(obj26, 0.83)
	ns.AggressionLevel(obj28, 0.83)
	ns.AggressionLevel(obj27, 0.83)
	ns.AggressionLevel(obj29, 0.83)
	ns.AggressionLevel(obj30, 0.83)
	ns.MoveObject(obj9, ns.GetWaypointX(wp137), ns.GetWaypointY(wp137))
	ns.MoveObject(obj12, ns.GetWaypointX(wp139), ns.GetWaypointY(wp139))
	ns.MoveObject(obj16, ns.GetWaypointX(wp138), ns.GetWaypointY(wp138))
	ns.MoveObject(obj14, ns.GetWaypointX(wp136), ns.GetWaypointY(wp136))
LABEL4:
	return
}
func EscapeHatch() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.WallGroupOpen(gvar64)
LABEL1:
	return
}
func CloseHatch() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.WallGroupClose(gvar64)
LABEL1:
	return
}
func MapInitialize() {
	obj6 = ns.Object("TowerNPC")
	obj7 = ns.Object("TowerNPC2")
	obj8 = ns.Object("TowerMaiden")
	obj9 = ns.Object("F6Townsman2")
	obj10 = ns.Object("Jesse")
	obj11 = ns.Object("Priest")
	obj14 = ns.Object("BrightBlades")
	obj15 = ns.Object("BrightBladesA")
	obj12 = ns.Object("Mystic")
	obj13 = ns.Object("Mystic2")
	obj18 = ns.Object("BookVendor")
	obj19 = ns.Object("BookVendorA")
	obj16 = ns.Object("Loreman")
	obj17 = ns.Object("Kincaid")
	obj20 = ns.Object("LameBarKeep")
	obj21 = ns.Object("Max")
	obj22 = ns.Object("Warden")
	obj23 = ns.Object("Morgan")
	obj24 = ns.Object("F6FireGuard8")
	obj32 = ns.Object("AppleMan")
	obj33 = ns.Object("Undertaker")
	obj34 = ns.Object("Mlurgh")
	obj35 = ns.Object("Grillf")
	obj36 = ns.Object("F6Elite4")
	obj37 = ns.Object("NPCMover")
	obj38 = ns.Object("DaMan")
	obj28 = ns.Object("Dorian")
	obj26 = ns.Object("Jorgan")
	obj27 = ns.Object("Grunbar")
	obj25 = ns.Object("Albi")
	obj29 = ns.Object("Civvy3")
	obj30 = ns.Object("Civvy4")
	obj31 = ns.Object("Daniel")
	obj62 = ns.Object("DaManKey")
	obj39 = ns.Object("GearRoomDoor")
	obj40 = ns.Object("GuardDoor1")
	obj41 = ns.Object("MaxHomeDoor1")
	obj42 = ns.Object("MaxHomeDoor2")
	obj43 = ns.Object("FirstJailDoor")
	obj44 = ns.Object("WarriorCellDoor")
	obj45 = ns.Object("MorganCellDoor")
	obj46 = ns.Object("BackRoomDoor")
	obj47 = ns.Object("TowerGate1")
	obj48 = ns.Object("TowerGate2")
	obj49 = ns.Object("JailCellTrigger")
	obj50 = ns.Object("WardenBackTrigger")
	obj51 = ns.Object("ThanksTrigger")
	obj54 = ns.Object("WardenFaceThis")
	obj55 = ns.Object("RogueShield")
	obj56 = ns.Object("MorganGold")
	obj57 = ns.Object("TowerKey")
	obj58 = ns.Object("GoldKey1")
	obj59 = ns.Object("GoldKey2")
	obj60 = ns.Object("SecretShirt")
	obj61 = ns.Object("SecretApple")
	gvar53 = ns.ObjectGroup("RewardTriggerGroup")
	gvar63 = ns.WallGroup("JumpWalls")
	gvar64 = ns.WallGroup("FallWalls")
	//gvar135 = ns.Waypoint("MorganFriendWP")
	wp133 = ns.Waypoint("WardenFooledWP")
	gvar134 = ns.WaypointGroup("EscapeWP")
	gvar135 = ns.WaypointGroup("MorganFriendWP")
	wp4 = ns.Waypoint("Storage")
	wp131 = ns.Waypoint("PlayerSounds")
	wp132 = ns.Waypoint("SecretSounds")
	wp136 = ns.Waypoint("BBWP")
	wp137 = ns.Waypoint("PhimWP")
	wp139 = ns.Waypoint("MVWP")
	wp138 = ns.Waypoint("KinWP")
	wp140 = ns.Waypoint("MorganGuardWP")
	wp141 = ns.Waypoint("MickGuardWP")
	ns.WayPointGroupOff(gvar134)
	ns.WayPointGroupOff(gvar135)
	ns.LockDoor(obj39)
	ns.LockDoor(obj40)
	ns.LockDoor(obj44)
	ns.LockDoor(obj41)
	ns.LockDoor(obj42)
	ns.SetOwner(ns.GetHost(), obj11)
	ns.SetDialog(obj11, ns.NORMAL, PriestGreetingStart, PriestGreetingEnd)
	ns.SetDialog(obj23, ns.NORMAL, MorganHintStart, MorganHintEnd)
	ns.SetDialog(obj24, ns.NORMAL, MorganFriendStart, MorganFriendEnd)
	ns.SetDialog(obj21, ns.NORMAL, NewMaxStart, NewMaxEnd)
	ns.SetDialog(obj33, ns.NORMAL, UndertakerTalkStart, UndertakerTalkEnd)
	ns.SetDialog(obj36, ns.NORMAL, BoothStart, BoothEnd)
	ns.StoryPic(obj21, "MaxPic")
	ns.StoryPic(obj33, "UndertakerPic")
	ns.StoryPic(obj28, "Townsman3Pic")
	ns.StoryPic(obj27, "Townsman2Pic")
	ns.StoryPic(obj25, "Townsman1Pic")
	ns.StoryPic(obj26, "Warrior2Pic")
	ns.StoryPic(obj22, "WardenPic")
	ns.StoryPic(obj11, "GalavaPriestPic")
	ns.StoryPic(obj23, "MorganPic")
	ns.StoryPic(obj24, "FentonPic")
	ns.StoryPic(obj31, "DrunkPic")
	ns.StoryPic(obj36, "MalePic4")
	gvar100 = gvar118
	gvar101 = gvar122
	gvar102 = gvar126
	ns.FrameTimer(1, CreatureSetup)
}
func PlayerDeath() {
	ns.DeathScreen(7)
}
func MickContact() {
	r0 := ns.IsAttackedBy(obj24, ns.GetCaller())
	if !r0 {
		goto LABEL1
	}
	if !flag92 {
		goto LABEL1
	}
	ns.AggressionLevel(obj24, 0.83)
LABEL1:
	return
}
func WardenWaiting() {
	if !ns.IsCaller(obj22) {
		goto LABEL1
	}
	if !(gvar102 != gvar126 || gvar101 == gvar125) {
		goto LABEL1
	}
	if !flag82 {
		goto LABEL1
	}
	ns.CreatureIdle(obj22)
	ns.AggressionLevel(obj22, 0.16)
	ns.CreatureGuard(obj22, ns.GetObjectX(obj22), ns.GetObjectY(obj22), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
LABEL1:
	return
}
func ResetMaxDialog() {
	if gvar148 != gvar147 {
		goto LABEL1
	}
	ns.CancelDialog(obj21)
	goto LABEL2
LABEL1:
	ns.SetDialog(obj21, ns.YESNO, MaxDialogStart, MaxDialogEnd)
LABEL2:
	return
}
func NullFunction() {
}
func MaxDie() {
	ns.CancelDialog(obj21)
}
func DaManDragged() {
	if !ns.IsCaller(obj38) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.UnlockDoor(obj41)
	ns.UnlockDoor(obj42)
LABEL1:
	return
}
func EmptyPlayerInventory() {
	var v0 int
	v0 = gvar151
	if v0 == gvar149 {
		goto LABEL1
	}
	if v0 == gvar150 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.Drop(ns.GetHost(), obj5[ivar153])
	ivar154 += 1
	gvar151 = gvar150
	ns.FrameTimer(1, EmptyPlayerInventory)
LABEL2:
	if !(ivar153 <= ivar152) {
		goto LABEL3
	}
	ns.MoveObject(obj5[ivar153], ns.GetWaypointX(wp4), ns.GetWaypointY(wp4))
	gvar151 = gvar149
	ivar153 += 1
	ns.FrameTimer(1, EmptyPlayerInventory)
LABEL3:
	return
}
func MakeAWish() {
	var v0 int
	v0 = ns.MaxHealth(ns.GetCaller())
	ns.MoveWaypoint(wp131, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.RestoreHealth(ns.GetCaller(), v0-ns.CurrentHealth(ns.GetCaller()))
	ns.AudioEvent(ns.RestoreHealthName, wp131)
	ns.PrintToAll("GeneralPrint:WellSignRefresh")
	ns.AudioEvent(ns.RestoreHealthName, ns.Waypoint("WellWP"))
}
func Jump() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	if !ns.HasItem(ns.GetHost(), obj60) {
		goto LABEL1
	}
	if !ns.HasItem(ns.GetHost(), obj61) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.NoWallSound(true)
	ns.WallGroupOpen(gvar63)
	ns.Chat(obj8, "Con02a:BarMaiden2")
	ns.SecondTimer(3, Upset)
LABEL1:
	return
}
func Upset() {
	ns.DestroyEveryChat()
	ns.Chat(obj6, "Wiz02B.scr:LewisTalk06")
	ns.SecondTimer(3, StartJump)
}
func StartJump() {
	ns.WallGroupOpen(gvar64)
	ns.Wander(obj6)
}
func MoveJumper() {
	if !ns.IsCaller(obj6) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.CreatureIdle(obj6)
	ns.ObjectOn(obj37)
	ns.Chat(obj6, "Wiz02B.scr:LewisTalk03")
LABEL1:
	return
}
func KillJumper() {
	if !ns.IsCaller(obj6) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.DestroyEveryChat()
	ns.Damage(obj6, 0, 500, 0)
	ns.WallGroupClose(gvar64)
	ns.WallGroupClose(gvar63)
	ns.NoWallSound(false)
LABEL1:
	return
}
func DaManHurt() {
	if !(!ns.IsLocked(obj41) || !ns.IsLocked(obj42)) {
		goto LABEL1
	}
	ns.AggressionLevel(obj38, 0.83)
	ns.Attack(obj38, ns.GetHost())
LABEL1:
	return
}
func DaManDie() {
	if !(ns.IsLocked(obj41) && ns.IsLocked(obj42)) {
		goto LABEL1
	}
	ns.Drop(obj38, obj62)
	ns.Delete(obj62)
LABEL1:
	return
}
func MorganRecognize() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	flag91 = true
LABEL1:
	return
}
func MorganLoseSight() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	flag91 = false
LABEL1:
	return
}
func ScorpionAway() {
	r1 := ns.IsAttackedBy(ns.GetHost(), ns.GetCaller())
	if !r1 {
		goto LABEL1
	}
	ns.GoBackHome(ns.GetCaller())
LABEL1:
	return
}
func CivvyJailCheck() {
	if !ns.IsCaller(obj26) {
		goto LABEL1
	}
	flag160 = true
	ns.GoBackHome(obj26)
LABEL1:
	if !ns.IsCaller(obj27) {
		goto LABEL2
	}
	flag159 = true
	ns.GoBackHome(obj27)
LABEL2:
	if !ns.IsCaller(obj25) {
		goto LABEL3
	}
	flag158 = true
	ns.GoBackHome(obj25)
LABEL3:
	if !ns.IsCaller(obj28) {
		goto LABEL4
	}
	flag157 = true
	ns.GoBackHome(obj28)
LABEL4:
	if !ns.IsCaller(obj29) {
		goto LABEL5
	}
	flag161 = true
	ns.GoBackHome(obj29)
LABEL5:
	if !ns.IsCaller(obj30) {
		goto LABEL6
	}
	flag162 = true
	ns.GoBackHome(obj30)
LABEL6:
	return
}
func FireballSecret() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.MoveWaypoint(wp132, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretFound, wp132)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 50)
LABEL1:
	return
}
func PlayerEnterRoom() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	flag90 = true
LABEL1:
	if !ns.IsCaller(obj26) {
		goto LABEL2
	}
	flag160 = true
	ns.GoBackHome(obj26)
LABEL2:
	if !ns.IsCaller(obj27) {
		goto LABEL3
	}
	flag159 = true
	ns.GoBackHome(obj27)
LABEL3:
	if !ns.IsCaller(obj25) {
		goto LABEL4
	}
	flag158 = true
	ns.GoBackHome(obj25)
LABEL4:
	if !ns.IsCaller(obj28) {
		goto LABEL5
	}
	flag157 = true
	ns.GoBackHome(obj28)
LABEL5:
	if !ns.IsCaller(obj29) {
		goto LABEL6
	}
	flag161 = true
	ns.GoBackHome(obj29)
LABEL6:
	if !ns.IsCaller(obj30) {
		goto LABEL7
	}
	flag162 = true
	ns.GoBackHome(obj30)
LABEL7:
	return
}
func JorganReport() {
	if gvar97 != 1 {
		goto LABEL1
	}
	ns.AggressionLevel(obj26, 0.83)
	ns.CreatureHunt(obj26)
	goto LABEL2
LABEL1:
	ns.AggressionLevel(obj26, 0.16)
	ns.Wander(obj26)
LABEL2:
	flag160 = false
}
func AlbiReport() {
	if gvar94 != 1 {
		goto LABEL1
	}
	ns.AggressionLevel(obj25, 0.83)
	ns.CreatureHunt(obj25)
	goto LABEL2
LABEL1:
	ns.AggressionLevel(obj25, 0.16)
	ns.Wander(obj25)
LABEL2:
	flag158 = false
}
func GrunbarReport() {
	if gvar96 != 1 {
		goto LABEL1
	}
	ns.AggressionLevel(obj27, 0.83)
	ns.CreatureHunt(obj27)
	goto LABEL2
LABEL1:
	ns.AggressionLevel(obj27, 0.16)
	ns.Wander(obj27)
LABEL2:
	flag159 = false
}
func DorianReport() {
	if gvar95 != 1 {
		goto LABEL1
	}
	ns.AggressionLevel(obj28, 0.83)
	ns.CreatureHunt(obj28)
	goto LABEL2
LABEL1:
	ns.AggressionLevel(obj28, 0.16)
	ns.Wander(obj28)
LABEL2:
	flag157 = false
}
func Civvy3Report() {
	if gvar98 != 1 {
		goto LABEL1
	}
	ns.AggressionLevel(obj29, 0.83)
	ns.CreatureHunt(obj29)
	goto LABEL2
LABEL1:
	ns.AggressionLevel(obj29, 0.16)
	ns.Wander(obj29)
LABEL2:
	flag161 = false
}
func Civvy4Report() {
	if gvar99 != 1 {
		goto LABEL1
	}
	ns.AggressionLevel(obj30, 0.83)
	ns.CreatureHunt(obj30)
	goto LABEL2
LABEL1:
	ns.AggressionLevel(obj30, 0.16)
	ns.Wander(obj30)
LABEL2:
	flag162 = false
}
func OnEvent(typ string) {
	switch typ {
	case "MapInitialize":
		MapInitialize()
	case "PlayerDeath":
		PlayerDeath()
	}
}
