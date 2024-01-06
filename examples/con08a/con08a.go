package con08a

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
	ivar42  int
	gvar43  int
	ivar44  int
	ivar45  int
	flag46  bool
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
	wp79    ns.WaypointID
	wp80    ns.WaypointID
	wp81    ns.WaypointID
	wp82    ns.WaypointID
	obj83   ns.ObjectID
	obj84   ns.ObjectID
	obj85   ns.ObjectID
	obj86   ns.ObjectID
	obj87   ns.ObjectID
	obj88   ns.ObjectID
	obj89   [8]ns.ObjectID
	wp90    ns.WaypointID
	wp91    ns.WaypointID
	gvar92  int
	gvar93  int
	gvar94  int
	gvar95  int
	gvar96  int
	gvar97  int
	gvar98  int
	gvar99  int
	gvar100 int
	gvar101 int
	gvar102 int
	gvar103 int
	ivar104 int
	gvar105 int
	ivar106 int
	flag107 bool
	flag108 bool
	flag109 bool
	flag110 bool
	flag111 bool
	flag112 bool
	flag113 bool
	flag114 bool
	flag115 bool
	flag116 bool
	gvar117 int
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
	obj128  [4]ns.ObjectID
	gvar129 int
	wp130   ns.WaypointID
	wp131   ns.WaypointID
	wp132   ns.WaypointID
	wp133   ns.WaypointID
	gvar134 ns.ObjectGroupID
	gvar135 ns.ObjectGroupID
	gvar136 ns.WallGroupID
	ivar137 int
	flag138 bool
)

func init() {
	gvar4 = 0
	gvar5 = 1
	gvar6 = 2
	gvar7 = 3
	gvar8 = 4
	gvar9 = 5
	gvar10 = 6
	gvar11 = 7
	gvar12 = 8
	gvar13 = 9
	gvar14 = 10
	gvar15 = 11
	gvar16 = 12
	gvar17 = 13
	gvar18 = 14
	gvar19 = 15
	gvar20 = 16
	gvar21 = 17
	gvar22 = 18
	gvar23 = 19
	gvar24 = 20
	gvar25 = 21
	gvar26 = 22
	gvar27 = 23
	gvar28 = 24
	gvar29 = 25
	gvar30 = 26
	gvar31 = gvar14
	gvar32 = gvar11
	gvar33 = gvar8
	gvar34 = gvar15
	gvar35 = gvar19
	gvar36 = gvar23
	gvar37 = gvar24
	gvar38 = gvar25
	gvar39 = gvar26
	gvar40 = gvar27
	gvar41 = gvar4
	ivar42 = 200
	gvar43 = 4
	ivar44 = 0
	ivar45 = 0
	flag46 = false
	gvar92 = 0
	gvar93 = 1
	gvar94 = 2
	gvar95 = 3
	gvar96 = 4
	gvar97 = 5
	gvar98 = 6
	gvar99 = 7
	gvar100 = 8
	gvar101 = 9
	gvar102 = 10
	gvar103 = gvar92
	ivar104 = 0
	gvar105 = gvar99
	ivar106 = 0
	flag107 = false
	flag108 = false
	flag109 = false
	flag110 = false
	flag111 = false
	flag112 = false
	flag113 = true
	flag114 = true
	flag115 = false
	flag116 = false
	ivar137 = 60
	flag138 = true
}
func JandorStart() {
	var v0 int
	v0 = gvar41
	if v0 == gvar4 {
		goto LABEL1
	}
	if v0 == gvar5 {
		goto LABEL2
	}
	if v0 == gvar6 {
		goto LABEL3
	}
	if v0 == gvar7 {
		goto LABEL4
	}
	goto LABEL5
LABEL1:
	ns.TellStory(ns.HumanMaleEatFood, "War08a:CaptainProd")
	goto LABEL5
LABEL2:
	ns.TellStory(ns.HumanMaleEatFood, "War08a:CaptainProd")
	goto LABEL5
LABEL3:
	ns.TellStory(ns.SwordsmanHurt, "Con03A.scr:JandorB")
	goto LABEL5
LABEL4:
	ns.Frozen(ns.GetHost(), true)
	ns.Blind()
	ns.FrameTimer(45, ExitToSwamp)
	goto LABEL5
LABEL5:
	return
}
func JandorEnd() {
	var v0 int
	v0 = gvar41
	if v0 == gvar4 {
		goto LABEL1
	}
	if v0 == gvar5 {
		goto LABEL2
	}
	if v0 == gvar6 {
		goto LABEL3
	}
	if v0 == gvar7 {
		goto LABEL4
	}
	goto LABEL5
LABEL1:
	ns.JournalEntry(ns.GetHost(), "Chapter8LocateAldwyn", 2)
	ns.PrintToAll("Con01a:NewJournalEntry")
	ns.AwardSpell(ns.GetHost(), ns.SPELL_SUMMON_CREATURE)
	ns.PrintToAll("GeneralPrint:Summon")
	ns.JournalEntry(ns.GetHost(), "SummonHint", 8)
	gvar41 = gvar5
	goto LABEL5
LABEL2:
	goto LABEL5
LABEL3:
	gvar41 = gvar7
	goto LABEL5
LABEL4:
	ns.Frozen(ns.GetHost(), true)
	ns.Blind()
	ns.FrameTimer(45, ExitToSwamp)
	goto LABEL5
LABEL5:
	return
}
func JacobDialogStart() {
	var (
		v0 int
		v1 int
		v2 int
	)
	v2 = gvar40
	if v2 == gvar27 {
		goto LABEL1
	}
	if v2 == gvar28 {
		goto LABEL2
	}
	if v2 == gvar29 {
		goto LABEL3
	}
	goto LABEL4
LABEL1:
	v0 = ns.Random(1, 3)
	v1 = v0
	if v1 == 1 {
		goto LABEL5
	}
	if v1 == 2 {
		goto LABEL6
	}
	if v1 == 3 {
		goto LABEL7
	}
	goto LABEL8
LABEL5:
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:JailerTalk01")
	goto LABEL8
LABEL6:
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:JailerTalk02")
	goto LABEL8
LABEL7:
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:JailerTalk03")
	goto LABEL8
LABEL8:
	goto LABEL4
LABEL2:
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:JailerTalk04")
	goto LABEL4
LABEL3:
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:JailerTalk05")
	goto LABEL4
LABEL4:
	return
}
func JacobDialogEnd() {
	var v0 int
	v0 = gvar40
	if v0 == gvar28 {
		goto LABEL1
	}
	if v0 == gvar29 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	flag46 = false
	ns.Frozen(ns.GetHost(), false)
	goto LABEL3
LABEL2:
	flag46 = false
	ns.Frozen(ns.GetHost(), false)
	goto LABEL3
LABEL3:
	return
}
func ResetJacobDialog() {
	if !ns.IsCaller(obj75) {
		goto LABEL1
	}
	if gvar40 != gvar30 {
		goto LABEL1
	}
	flag46 = false
	gvar40 = gvar27
	ns.SetDialog(obj75, ns.NORMAL, JacobDialogStart, JacobDialogEnd)
	ivar45 = 0
	ns.ObjectOff(obj78)
LABEL1:
	return
}
func ContestGuardStart() {
	ns.TellStory(ns.HumanMaleEatFood, "War02a:NewTownsman2")
}
func ContestGuardEnd() {
	ns.SetDialog(obj77, ns.NORMAL, ContestGuardStart, ContestGuardEnd)
}
func MorganInjured() {
	var v0 int
	if flag46 {
		goto LABEL1
	}
	flag46 = true
	ivar45 += 1
	v0 = ivar45
	if v0 == 1 {
		goto LABEL2
	}
	if v0 == 2 {
		goto LABEL3
	}
	if v0 == 3 {
		goto LABEL4
	}
	goto LABEL1
LABEL2:
	gvar40 = gvar28
	ns.Frozen(ns.GetHost(), true)
	ns.CreatureIdle(ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj75)
	ns.LookAtObject(obj75, ns.GetHost())
	ns.StartDialog(obj75, ns.GetHost())
	goto LABEL1
LABEL3:
	gvar40 = gvar29
	ns.Frozen(ns.GetHost(), true)
	ns.CreatureIdle(ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj75)
	ns.LookAtObject(obj75, ns.GetHost())
	ns.StartDialog(obj75, ns.GetHost())
	goto LABEL1
LABEL4:
	gvar40 = gvar30
	ns.CancelDialog(obj75)
	ns.ClearOwner(obj75)
	ns.AggressionLevel(obj75, 0.83)
	ns.Attack(obj75, ns.GetHost())
	goto LABEL1
LABEL1:
	return
}
func PlayerFlee() {
	if gvar40 != gvar30 {
		goto LABEL1
	}
	ns.ObjectOn(obj78)
	ns.SetOwner(ns.GetHost(), obj75)
	ns.AggressionLevel(obj75, 0)
	ns.CreatureGuard(obj75, ns.GetWaypointX(wp79), ns.GetWaypointY(wp79), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
LABEL1:
	return
}
func HenrickDialogStart() {
	var v0 int
	v0 = gvar32
	if v0 == gvar11 {
		goto LABEL1
	}
	if v0 == gvar12 {
		goto LABEL2
	}
	if v0 == gvar13 {
		goto LABEL3
	}
	goto LABEL4
LABEL1:
	ns.LookAtObject(obj48, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War08b:HenrickSalesPitchA")
	goto LABEL4
LABEL2:
	ns.LookAtObject(obj48, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War08b:HenrickSalesPitchB")
	goto LABEL4
LABEL3:
	ns.LookAtObject(obj48, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War08b:HenrickNoMoreWolves")
	goto LABEL4
LABEL4:
	return
}
func HenrickDialogEnd() {
	var (
		v0 int
		v1 int
		v2 int
		v3 int
		v4 int
		v5 int
		v6 int
		v7 int
	)
	v0 = 0
	v1 = 0
	v2 = 1
	v3 = 2
	v4 = 0
	v7 = gvar32
	if v7 == gvar11 {
		goto LABEL1
	}
	if v7 == gvar12 {
		goto LABEL2
	}
	if v7 == gvar13 {
		goto LABEL3
	}
	goto LABEL4
LABEL1:
	v1 = ns.GetAnswer(obj48)
	v5 = v1
	if v5 == v2 {
		goto LABEL5
	}
	if v5 == v3 {
		goto LABEL6
	}
	if v5 == v4 {
		goto LABEL7
	}
	goto LABEL8
LABEL5:
	v0 = ns.GetGold(ns.GetHost())
	if !(v0 < ivar42) {
		goto LABEL9
	}
	ns.SetDialog(obj48, ns.NORMAL, NullDialogStart, ResetHenrickDialogChoice)
	ns.TellStory(ns.SwordsmanHurt, "War08b:HenrickNotEnoughGold")
	goto LABEL10
LABEL9:
	ns.ChangeGold(ns.GetHost(), -ivar42)
	gvar32 = gvar12
	ReleaseCharmedWolf(ivar44)
	ivar44 += 1
	ns.SetDialog(obj48, ns.NORMAL, NullDialogStart, ResetHenrickDialogChoice)
	ns.TellStory(ns.SwordsmanHurt, "War08b:HenrickSaleSuccessful")
LABEL10:
	goto LABEL8
LABEL6:
	goto LABEL8
LABEL7:
	goto LABEL8
LABEL8:
	goto LABEL4
LABEL2:
	v1 = ns.GetAnswer(obj48)
	v6 = v1
	if v6 == v2 {
		goto LABEL11
	}
	if v6 == v3 {
		goto LABEL12
	}
	if v6 == v4 {
		goto LABEL13
	}
	goto LABEL14
LABEL11:
	v0 = ns.GetGold(ns.GetHost())
	if !(v0 < ivar42) {
		goto LABEL15
	}
	ns.SetDialog(obj48, ns.NORMAL, NullDialogStart, ResetHenrickDialogChoice)
	ns.TellStory(ns.SwordsmanHurt, "War08b:HenrickNotEnoughGold")
	goto LABEL16
LABEL15:
	ns.ChangeGold(ns.GetHost(), -ivar42)
	ReleaseCharmedWolf(ivar44)
	ivar44 += 1
	ns.SetDialog(obj48, ns.NORMAL, NullDialogStart, ResetHenrickDialogChoice)
	ns.TellStory(ns.SwordsmanHurt, "War08b:HenrickSaleSuccessful")
	if !(ivar44 >= 4) {
		goto LABEL16
	}
	gvar32 = gvar13
	ns.SetDialog(obj48, ns.NORMAL, HenrickDialogStart, HenrickDialogEnd)
LABEL16:
	goto LABEL14
LABEL12:
	goto LABEL14
LABEL13:
	goto LABEL14
LABEL14:
	goto LABEL4
LABEL3:
	goto LABEL4
LABEL4:
	return
}
func DrunkDialogStart() {
	ns.LookAtObject(obj49, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:DrunkGreeting")
}
func DrunkDialogEnd() {
}
func GatekeeperDialogStart() {
	var v0 int
	ns.LookAtObject(obj52, ns.GetHost())
	v0 = gvar36
	if v0 == gvar23 {
		goto LABEL1
	}
	goto LABEL2
LABEL1:
	ns.TellStory(ns.SwordsmanHurt, "Con02a:GatekeeperWaiting")
	goto LABEL2
LABEL2:
	return
}
func GatekeeperDialogEnd() {
}
func Gatekeeper2DialogStart() {
	var v0 int
	ns.LookAtObject(obj63, ns.GetHost())
	v0 = gvar37
	if v0 == gvar24 {
		goto LABEL1
	}
	goto LABEL2
LABEL1:
	ns.TellStory(ns.SwordsmanHurt, "Con08a:GuardGreet")
	goto LABEL2
LABEL2:
	return
}
func Gatekeeper2DialogEnd() {
}
func Gatekeeper3DialogStart() {
	var v0 int
	ns.LookAtObject(obj64, ns.GetHost())
	v0 = gvar38
	if v0 == gvar25 {
		goto LABEL1
	}
	goto LABEL2
LABEL1:
	ns.TellStory(ns.SwordsmanHurt, "War08a:Guard02Greet")
	goto LABEL2
LABEL2:
	return
}
func Gatekeeper3DialogEnd() {
}
func MayorDialogStart() {
	var (
		v0 int
		v1 int
	)
	ns.Frozen(ns.GetHost(), true)
	ns.CreatureIdle(ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj55)
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
	ns.TellStory(ns.SwordsmanHurt, "War03b:MayorPre")
	goto LABEL3
LABEL2:
	ns.TellStory(ns.SwordsmanHurt, "War03b:MayorBye")
	goto LABEL3
LABEL3:
	return
}
func MayorDialogEnd() {
	ns.Frozen(ns.GetHost(), false)
}
func MayorsGuardDialogStart() {
	var v0 int
	v0 = gvar31
	if v0 == gvar14 {
		goto LABEL1
	}
	goto LABEL2
LABEL1:
	ns.LookAtObject(obj53, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War03b:MayorsGuardPre")
	goto LABEL2
LABEL2:
	return
}
func MayorsGuardDialogEnd() {
}
func AldwinDialogStart() {
	var v0 int
	v0 = gvar33
	if v0 == gvar8 {
		goto LABEL1
	}
	if v0 == gvar9 {
		goto LABEL2
	}
	if v0 == gvar10 {
		goto LABEL3
	}
	goto LABEL4
LABEL1:
	ns.LookAtObject(obj54, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con08a:AldwynGreet")
	goto LABEL4
LABEL2:
	ns.LookAtObject(obj54, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con08:AldwynProd2")
	goto LABEL4
LABEL3:
	ns.LookAtObject(obj54, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con08:AldwynProd")
	goto LABEL4
LABEL4:
	return
}
func AldwinDialogEnd() {
	var v0 int
	v0 = gvar33
	if v0 == gvar8 {
		goto LABEL1
	}
	if v0 == gvar9 {
		goto LABEL2
	}
	if v0 == gvar10 {
		goto LABEL3
	}
	goto LABEL4
LABEL1:
	ns.Pickup(ns.GetHost(), obj50)
	ns.PrintToAll("GeneralPrint:GainedKey")
	gvar33 = gvar9
	ns.JournalEdit(ns.GetHost(), "Chapter8LocateAldwyn", 4)
	ns.PrintToAll("Con02a:ObjectiveComplete")
	ns.JournalEntry(ns.GetHost(), "Chapter8IxPriest", 2)
	ns.PrintToAll("Con01a:NewJournalEntry")
	goto LABEL4
LABEL2:
	gvar33 = gvar10
	goto LABEL4
LABEL3:
	goto LABEL4
LABEL4:
	return
}
func BridgeGuardDialogStart() {
	var v0 int
	v0 = gvar39
	if v0 == gvar26 {
		goto LABEL1
	}
	goto LABEL2
LABEL1:
	ns.LookAtObject(obj65, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con02a:RandomSay")
	goto LABEL2
LABEL2:
	return
}
func BridgeGuardDialogEnd() {
	var v0 int
	v0 = gvar39
	if v0 == gvar26 {
		goto LABEL1
	}
	goto LABEL2
LABEL1:
	goto LABEL2
LABEL2:
	return
}
func Maiden1DialogStart() {
	ns.LookAtObject(obj67, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con02a:BarMaiden")
}
func Maiden1DialogEnd() {
}
func Maiden2DialogStart() {
	ns.Frozen(ns.GetHost(), true)
	ns.Frozen(obj68, true)
	ns.CreatureIdle(ns.GetHost())
	ns.CreatureIdle(obj68)
	ns.LookAtObject(ns.GetHost(), obj68)
	ns.LookAtObject(obj68, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "Con07B.scr:MaidenTalk01")
}
func Maiden2DialogEnd() {
	ns.Frozen(ns.GetHost(), false)
	ns.Frozen(obj68, false)
	ns.Wander(obj68)
}
func Maiden3DialogStart() {
	ns.Frozen(ns.GetHost(), true)
	ns.Frozen(obj69, true)
	ns.CreatureIdle(ns.GetHost())
	ns.CreatureIdle(obj69)
	ns.LookAtObject(ns.GetHost(), obj69)
	ns.LookAtObject(obj69, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02A.scr:MaidenTalk05")
}
func Maiden3DialogEnd() {
	ns.Frozen(ns.GetHost(), false)
	ns.Frozen(obj69, false)
	ns.Wander(obj69)
}
func Maiden4DialogStart() {
	var (
		v0 int
		v1 int
	)
	ns.Frozen(ns.GetHost(), true)
	ns.Frozen(obj70, true)
	ns.CreatureIdle(ns.GetHost())
	ns.CreatureIdle(obj70)
	ns.LookAtObject(ns.GetHost(), obj70)
	ns.LookAtObject(obj70, ns.GetHost())
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
	ns.TellStory(ns.HumanMaleEatFood, "War02a:NewTownswoman1")
	goto LABEL3
LABEL2:
	ns.TellStory(ns.HumanMaleEatFood, "War02a:NewTownswoman3")
	goto LABEL3
LABEL3:
	return
}
func Maiden4DialogEnd() {
	ns.Frozen(ns.GetHost(), false)
	ns.Frozen(obj70, false)
	ns.Wander(obj70)
}
func Townsman1DialogStart() {
	ns.Frozen(ns.GetHost(), true)
	ns.Frozen(obj72, true)
	ns.CreatureIdle(ns.GetHost())
	ns.CreatureIdle(obj72)
	ns.LookAtObject(ns.GetHost(), obj72)
	ns.LookAtObject(obj72, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "War02a:NewTownsman3")
}
func Townsman1DialogEnd() {
	ns.Frozen(ns.GetHost(), false)
	ns.Frozen(obj72, false)
	ns.Wander(obj72)
}
func Townsman2DialogStart() {
	ns.Frozen(ns.GetHost(), true)
	ns.Frozen(obj73, true)
	ns.CreatureIdle(ns.GetHost())
	ns.CreatureIdle(obj73)
	ns.LookAtObject(ns.GetHost(), obj73)
	ns.LookAtObject(obj73, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02A.scr:AlbiTalk01")
}
func Townsman2DialogEnd() {
	ns.Frozen(ns.GetHost(), false)
	ns.Frozen(obj73, false)
	ns.Wander(obj73)
}
func Townsman3DialogStart() {
	ns.Frozen(ns.GetHost(), true)
	ns.Frozen(obj74, true)
	ns.CreatureIdle(ns.GetHost())
	ns.CreatureIdle(obj74)
	ns.LookAtObject(ns.GetHost(), obj74)
	ns.LookAtObject(obj74, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02A.scr:MaxRumor03")
}
func Townsman3DialogEnd() {
	ns.Frozen(ns.GetHost(), false)
	ns.Frozen(obj74, false)
	ns.Wander(obj74)
}
func ReleaseCharmedWolf(a1 int) {
	var v0 int
	v0 = a1
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
	goto LABEL5
LABEL1:
	ns.WallOpen(ns.Wall(59, 113))
	ns.EnchantOff(obj128[a1], ns.ENCHANT_INVULNERABLE)
	ns.CreatureFollow(obj128[a1], ns.GetHost())
	ns.PauseObject(obj128[a1], 30)
	ns.AggressionLevel(obj128[a1], 0.83)
	ns.MakeFriendly(obj128[a1])
	ns.BecomePet(obj128[a1])
	goto LABEL5
LABEL2:
	ns.WallOpen(ns.Wall(57, 111))
	ns.EnchantOff(obj128[a1], ns.ENCHANT_INVULNERABLE)
	ns.CreatureFollow(obj128[a1], ns.GetHost())
	ns.PauseObject(obj128[a1], 30)
	ns.AggressionLevel(obj128[a1], 0.83)
	ns.MakeFriendly(obj128[a1])
	ns.BecomePet(obj128[a1])
	goto LABEL5
LABEL3:
	ns.WallOpen(ns.Wall(55, 109))
	ns.EnchantOff(obj128[a1], ns.ENCHANT_INVULNERABLE)
	ns.CreatureFollow(obj128[a1], ns.GetHost())
	ns.PauseObject(obj128[a1], 30)
	ns.AggressionLevel(obj128[a1], 0.83)
	ns.MakeFriendly(obj128[a1])
	ns.BecomePet(obj128[a1])
	goto LABEL5
LABEL4:
	ns.WallOpen(ns.Wall(53, 107))
	ns.EnchantOff(obj128[a1], ns.ENCHANT_INVULNERABLE)
	ns.CreatureFollow(obj128[a1], ns.GetHost())
	ns.PauseObject(obj128[a1], 30)
	ns.AggressionLevel(obj128[a1], 0.83)
	ns.MakeFriendly(obj128[a1])
	ns.BecomePet(obj128[a1])
	goto LABEL5
LABEL5:
	return
}
func InitializeDialogs() {
	ns.StoryPic(obj48, "Townsman2Pic")
	ns.StoryPic(obj49, "DrunkPic")
	ns.StoryPic(obj51, "Townsman1Pic")
	ns.StoryPic(obj52, "IxGuard2Pic")
	ns.StoryPic(obj53, "Warrior3Pic")
	ns.StoryPic(obj55, "TheogrinPic")
	ns.StoryPic(obj54, "AldwynPic")
	ns.StoryPic(obj63, "IxGuard2Pic")
	ns.StoryPic(obj64, "IxGuard1Pic")
	ns.StoryPic(obj67, "MaidenPic")
	ns.StoryPic(obj68, "MaidenPic2")
	ns.StoryPic(obj69, "MaidenPic3")
	ns.StoryPic(obj70, "MaidenPic4")
	ns.StoryPic(obj72, "MalePic4")
	ns.StoryPic(obj73, "MalePic11")
	ns.StoryPic(obj74, "MalePic12")
	ns.StoryPic(obj47, "AirshipCaptainPic")
	ns.StoryPic(obj75, "WardenPic")
	ns.SetDialog(obj75, ns.NORMAL, JacobDialogStart, JacobDialogEnd)
	ns.SetDialog(obj49, ns.NORMAL, DrunkDialogStart, DrunkDialogEnd)
	ns.SetDialog(obj63, ns.NORMAL, Gatekeeper2DialogStart, Gatekeeper2DialogEnd)
	ns.SetDialog(obj64, ns.NORMAL, Gatekeeper3DialogStart, Gatekeeper3DialogEnd)
	ns.SetDialog(obj53, ns.NORMAL, MayorsGuardDialogStart, MayorsGuardDialogEnd)
	ns.SetDialog(obj54, ns.NORMAL, AldwinDialogStart, AldwinDialogEnd)
	ns.SetDialog(obj67, ns.NORMAL, Maiden1DialogStart, Maiden1DialogEnd)
	ns.SetDialog(obj68, ns.NORMAL, Maiden2DialogStart, Maiden2DialogEnd)
	ns.SetDialog(obj69, ns.NORMAL, Maiden3DialogStart, Maiden3DialogEnd)
	ns.SetDialog(obj70, ns.NORMAL, Maiden4DialogStart, Maiden4DialogEnd)
	ns.SetDialog(obj72, ns.NORMAL, Townsman1DialogStart, Townsman1DialogEnd)
	ns.SetDialog(obj73, ns.NORMAL, Townsman2DialogStart, Townsman2DialogEnd)
	ns.SetDialog(obj74, ns.NORMAL, Townsman3DialogStart, Townsman3DialogEnd)
}
func InitializeSecondaryDialogs() {
	gvar41 = gvar6
}
func NullDialogStart() {
}
func ExitToSwamp() {
	ns.Frozen(ns.GetHost(), false)
	ns.MoveObject(ns.GetHost(), ns.GetWaypointX(ns.Waypoint("SwampExitWP")), ns.GetWaypointY(ns.Waypoint("SwampExitWP")))
}
func StartCaptainConversation() {
	ns.SetDialog(obj47, ns.NORMAL, JandorStart, JandorEnd)
	ns.StartDialog(obj47, ns.GetHost())
}
func ResetHenrickDialogChoice() {
	ns.SetDialog(obj48, ns.YESNO, HenrickDialogStart, HenrickDialogEnd)
}
func PlayerDeath() {
	ns.DeathScreen(8)
}
func GoToMineSEG2() {
	ns.MoveObject(ns.GetHost(), ns.GetWaypointX(wp81), ns.GetWaypointY(wp81))
}
func GoToMine() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.Blind()
	ns.FrameTimer(60, GoToMineSEG2)
LABEL1:
	return
}
func GoToCaveSEG2() {
	ns.MoveObject(ns.GetHost(), ns.GetWaypointX(wp82), ns.GetWaypointY(wp82))
}
func GoToCave() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.Blind()
	ns.FrameTimer(60, GoToCaveSEG2)
LABEL1:
	return
}
func InversionInit() {
	obj83 = ns.Object("Teacher")
	obj84 = ns.Object("Apprentice")
	obj85 = ns.Object("InvertSpell")
	obj86 = ns.Object("LessonGate1")
	obj87 = ns.Object("LessonGate2")
	obj88 = ns.Object("ApprenticeTrigger")
	wp90 = ns.Waypoint("PlayerSounds")
	wp91 = ns.Waypoint("ApprenticeWP")
	ns.StoryPic(obj83, "WizardGuard1Pic")
	ns.SetDialog(obj83, ns.YESNO, TeacherStart, TeacherEnd)
	ns.LockDoor(obj86)
	ns.LockDoor(obj87)
	ns.ObjectOff(obj88)
	ns.FrameTimer(1, InversionOwnObject)
}
func InversionOwnObject() {
	ns.SetOwner(ns.GetHost(), obj83)
	ns.SetOwner(ns.GetHost(), obj84)
	ns.Enchant(obj84, ns.ENCHANT_INVULNERABLE, 0)
	ns.Enchant(obj83, ns.ENCHANT_ANCHORED, 0)
	ns.Enchant(obj84, ns.ENCHANT_ANCHORED, 0)
}
func TeacherStart() {
	var v0 int
	v0 = gvar103
	if v0 == gvar92 {
		goto LABEL1
	}
	if v0 == gvar93 {
		goto LABEL2
	}
	if v0 == gvar94 {
		goto LABEL3
	}
	if v0 == gvar95 {
		goto LABEL4
	}
	if v0 == gvar96 {
		goto LABEL5
	}
	if v0 == gvar97 {
		goto LABEL6
	}
	if v0 == gvar98 {
		goto LABEL7
	}
	goto LABEL8
LABEL1:
	ns.Frozen(ns.GetHost(), true)
	ns.CreatureIdle(ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj83)
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02B.scr:InversionBoyTalk01")
	goto LABEL8
LABEL2:
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02B.scr:InversionBoyTalk04")
	goto LABEL8
LABEL3:
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02B.scr:InversionBoyTalk02")
	goto LABEL8
LABEL4:
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02B.scr:InversionBoyTalk03")
	goto LABEL8
LABEL5:
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02B.scr:InversionBoyTalk05")
	goto LABEL8
LABEL6:
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02B.scr:InversionBoyTalk04a")
	goto LABEL8
LABEL7:
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02B.scr:InversionBoyTalk06")
	goto LABEL8
LABEL8:
	return
}
func TeacherEnd() {
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
	v6 = gvar103
	if v6 == gvar92 {
		goto LABEL1
	}
	if v6 == gvar93 {
		goto LABEL2
	}
	if v6 == gvar94 {
		goto LABEL3
	}
	if v6 == gvar95 {
		goto LABEL4
	}
	if v6 == gvar96 {
		goto LABEL5
	}
	if v6 == gvar97 {
		goto LABEL6
	}
	if v6 == gvar98 {
		goto LABEL7
	}
	goto LABEL8
LABEL1:
	v1 = ns.GetAnswer(obj83)
	v5 = v1
	if v5 == v2 {
		goto LABEL9
	}
	if v5 == v3 {
		goto LABEL10
	}
	if v5 == v4 {
		goto LABEL11
	}
	goto LABEL12
LABEL9:
	v0 = ns.GetGold(ns.GetHost())
	if !(v0 < 100) {
		goto LABEL13
	}
	ns.SetDialog(obj83, ns.NORMAL, TeacherStart, TeacherEnd)
	gvar103 = gvar93
	ns.StartDialog(obj83, ns.GetHost())
	goto LABEL12
	goto LABEL10
LABEL13:
	ns.ChangeGold(ns.GetHost(), -100)
	ns.UnlockDoor(obj86)
	ns.UnlockDoor(obj87)
	ns.SetDialog(obj83, ns.NORMAL, TeacherStart, TeacherEnd)
	gvar103 = gvar94
	ns.StartDialog(obj83, ns.GetHost())
	goto LABEL12
LABEL10:
	ns.Frozen(ns.GetHost(), false)
	gvar103 = gvar92
	ns.SetDialog(obj83, ns.YESNO, TeacherStart, TeacherEnd)
	goto LABEL12
LABEL11:
	ns.Frozen(ns.GetHost(), false)
	gvar103 = gvar92
	ns.SetDialog(obj83, ns.YESNO, TeacherStart, TeacherEnd)
	goto LABEL12
LABEL12:
	goto LABEL8
LABEL2:
	ns.Frozen(ns.GetHost(), false)
	gvar103 = gvar92
	ns.SetDialog(obj83, ns.YESNO, TeacherStart, TeacherEnd)
	goto LABEL8
LABEL3:
	ns.CancelDialog(obj83)
	ns.Frozen(ns.GetHost(), false)
	ns.Pickup(ns.GetHost(), obj85)
	goto LABEL8
LABEL4:
	ns.CancelDialog(obj83)
	ns.SecondTimer(1, Phoneme4)
	goto LABEL8
LABEL5:
	ns.UnlockDoor(obj58)
	ns.UnlockDoor(obj59)
	ns.UnlockDoor(obj86)
	ns.UnlockDoor(obj87)
	gvar103 = gvar98
	ns.SetDialog(obj83, ns.NORMAL, TeacherStart, TeacherEnd)
	goto LABEL8
LABEL6:
	ns.UnlockDoor(obj58)
	ns.UnlockDoor(obj59)
	ns.UnlockDoor(obj86)
	ns.UnlockDoor(obj87)
	gvar103 = gvar98
	ns.SetDialog(obj83, ns.NORMAL, TeacherStart, TeacherEnd)
	goto LABEL8
LABEL7:
	gvar103 = gvar98
	ns.SetDialog(obj83, ns.NORMAL, TeacherStart, TeacherEnd)
	goto LABEL8
LABEL8:
	return
}
func Phoneme1() {
	if flag111 {
		goto LABEL1
	}
	if !(ns.CurrentHealth(obj84) > 0) {
		goto LABEL1
	}
	ns.MoveWaypoint(wp90, 1326, 2493)
	ns.AudioEvent(ns.SpellPhonemeDown, wp90)
	ns.FrameTimer(10, Phoneme2)
LABEL1:
	return
}
func Phoneme2() {
	if flag111 {
		goto LABEL1
	}
	if !(ns.CurrentHealth(obj84) > 0) {
		goto LABEL1
	}
	ns.AudioEvent(ns.SpellPhonemeDown, wp90)
	ns.FrameTimer(10, Phoneme3)
LABEL1:
	return
}
func Phoneme3() {
	if flag111 {
		goto LABEL1
	}
	if !(ns.CurrentHealth(obj84) > 0) {
		goto LABEL1
	}
	ns.AudioEvent(ns.SpellPhonemeDown, wp90)
	ns.FrameTimer(10, CastSpells)
LABEL1:
	return
}
func Phoneme4() {
	if !flag111 {
		goto LABEL1
	}
	if !(ns.CurrentHealth(obj84) > 0) {
		goto LABEL1
	}
	ns.MoveWaypoint(wp90, 1326, 2493)
	ns.AudioEvent(ns.SpellPhonemeDown, wp90)
	ns.FrameTimer(10, Phoneme5)
LABEL1:
	return
}
func Phoneme5() {
	if !flag111 {
		goto LABEL1
	}
	if !(ns.CurrentHealth(obj84) > 0) {
		goto LABEL1
	}
	ns.AudioEvent(ns.SpellPhonemeDown, wp90)
	ns.FrameTimer(10, Phoneme6)
LABEL1:
	return
}
func Phoneme6() {
	if !flag111 {
		goto LABEL1
	}
	if !(ns.CurrentHealth(obj84) > 0) {
		goto LABEL1
	}
	ns.AudioEvent(ns.SpellPhonemeDown, wp90)
	ns.FrameTimer(10, CastSpells2)
LABEL1:
	return
}
func LessonBegin() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	if gvar103 != gvar94 {
		goto LABEL1
	}
	if !flag113 {
		goto LABEL2
	}
	ns.WideScreen(true)
	ns.Frozen(ns.GetHost(), true)
	ns.CreatureIdle(ns.GetHost())
	ns.LockDoor(obj86)
	ns.LockDoor(obj87)
	ns.Move(obj84, wp91)
	flag113 = false
LABEL2:
	if !(flag107 == false && flag112 == true) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.WideScreen(false)
	ns.Frozen(ns.GetHost(), false)
	ns.LookAtObject(obj84, ns.GetHost())
	flag107 = true
	ns.LockDoor(obj58)
	ns.LockDoor(obj59)
	ns.SecondTimer(1, Phoneme1)
LABEL1:
	return
}
func CastSpells() {
	if flag111 {
		goto LABEL1
	}
	gvar105 = gvar102
	if !(ivar104 < 4) {
		goto LABEL2
	}
	ns.ClearOwner(obj84)
	ns.CastSpellObjectLocation(ns.SPELL_SLOW, obj84, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.Enchant(obj84, ns.ENCHANT_INVULNERABLE, 0)
LABEL2:
	if ivar104 != 4 {
		goto LABEL3
	}
	ns.EnchantOff(obj84, ns.ENCHANT_INVULNERABLE)
	ns.CastSpellObjectLocation(ns.SPELL_MAGIC_MISSILE, obj84, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
LABEL3:
	ivar104 += 1
	if !(ivar104 < 5) {
		goto LABEL4
	}
	ns.SecondTimer(3, Phoneme1)
	goto LABEL1
LABEL4:
	flag108 = true
	if !flag110 {
		goto LABEL5
	}
	gvar103 = gvar96
	ns.SetDialog(obj83, ns.NORMAL, TeacherStart, TeacherEnd)
	ns.StartDialog(obj83, ns.GetHost())
	goto LABEL1
LABEL5:
	gvar103 = gvar97
	ns.SetDialog(obj83, ns.NORMAL, TeacherStart, TeacherEnd)
	ns.StartDialog(obj83, ns.GetHost())
LABEL1:
	return
}
func CastSpells2() {
	if !flag111 {
		goto LABEL1
	}
	gvar105 = gvar102
	if !(ivar104 < 4) {
		goto LABEL2
	}
	ns.CastSpellObjectLocation(ns.SPELL_SLOW, obj84, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.Enchant(obj84, ns.ENCHANT_INVULNERABLE, 0)
LABEL2:
	if ivar104 != 4 {
		goto LABEL3
	}
	ns.EnchantOff(obj84, ns.ENCHANT_INVULNERABLE)
	ns.CastSpellObjectLocation(ns.SPELL_MAGIC_MISSILE, obj84, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
LABEL3:
	ivar104 += 1
	if !(ivar104 < 5) {
		goto LABEL4
	}
	ns.SecondTimer(3, Phoneme4)
	goto LABEL1
LABEL4:
	flag108 = true
	if !flag110 {
		goto LABEL5
	}
	gvar103 = gvar96
	ns.SetDialog(obj83, ns.NORMAL, TeacherStart, TeacherEnd)
	ns.StartDialog(obj83, ns.GetHost())
	goto LABEL1
LABEL5:
	gvar103 = gvar97
	ns.SetDialog(obj83, ns.NORMAL, TeacherStart, TeacherEnd)
	ns.StartDialog(obj83, ns.GetHost())
LABEL1:
	return
}
func CheckApprentice() {
	if !ns.HasEnchant(obj84, ns.ENCHANT_SLOWED) {
		goto LABEL1
	}
	ns.ObjectOff(obj88)
	flag111 = true
	gvar103 = gvar95
	flag110 = true
	if !(ivar104 > 2) {
		goto LABEL2
	}
	ivar104 = 2
LABEL2:
	ns.SetDialog(obj83, ns.NORMAL, TeacherStart, TeacherEnd)
	ns.StartDialog(obj83, ns.GetHost())
LABEL1:
	if !(ns.CurrentHealth(obj84) <= 0) {
		goto LABEL3
	}
	ns.ObjectOff(ns.GetTrigger())
	flag110 = true
LABEL3:
	return
}
func ApprenticeDie() {
	ns.AudioEvent(ns.HumanMaleHurtHeavy, wp90)
}
func ApprenticeReport() {
	ns.CreatureGuard(obj84, ns.GetWaypointX(wp91), ns.GetWaypointY(wp91), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	ns.ObjectOn(obj88)
	flag112 = true
}
func StopCreatures() {
	var v0 int
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	flag114 = false
	v0 = 0
	for {
		if !(v0 <= ivar106) {
			goto LABEL1
		}
		ns.AggressionLevel(obj89[v0], 0)
		ns.CreatureIdle(obj89[v0])
		ns.ObjectOff(obj89[v0])
		v0 += 1
	}
LABEL1:
	r3 := ns.IsAttackedBy(ns.GetHost(), ns.GetCaller())
	if r3 {
		goto LABEL3
	}
	if !ns.HasClass(ns.GetCaller(), "CLASS_MONSTER") {
		goto LABEL3
	}
	obj89[ivar106] = ns.GetCaller()
	ns.AggressionLevel(ns.GetCaller(), 0)
	ns.CreatureIdle(ns.GetCaller())
	ns.ObjectOff(ns.GetCaller())
	ivar106 += 1
LABEL3:
	return
}
func StartCreatures() {
	var v0 int
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	v0 = 0
	for {
		if !(v0 <= ivar106) {
			goto LABEL1
		}
		ns.ObjectOn(obj89[v0])
		ns.CreatureFollow(obj89[v0], ns.GetHost())
		ns.AggressionLevel(obj89[v0], 0.5)
		flag114 = true
		v0 += 1
	}
LABEL1:
	return
}
func OwnObjects() {
	ns.GroupSetOwner(ns.GetHost(), gvar134)
	ns.SetOwner(ns.GetHost(), obj48)
	ns.GroupSetOwner(ns.GetHost(), gvar135)
	ns.SetOwner(ns.GetHost(), obj49)
	ns.SetOwner(ns.GetHost(), obj67)
	ns.SetOwner(ns.GetHost(), obj68)
	ns.SetOwner(ns.GetHost(), obj69)
	ns.SetOwner(ns.GetHost(), obj70)
	ns.SetOwner(ns.GetHost(), obj47)
	ns.SetOwner(ns.GetHost(), obj75)
	ns.SetOwner(ns.GetHost(), obj77)
}
func AwardSummonSpell() {
	ns.AwardSpell(ns.GetHost(), ns.SPELL_SUMMON_CREATURE)
	ns.PrintToAll("GeneralPrint:Summon")
	ns.JournalEntry(ns.GetHost(), "SummonHint", 8)
}
func DisableElevators() {
	ns.ObjectOff(obj60)
	ns.ObjectOff(obj66)
}
func MapInitialize() {
	obj49 = ns.Object("Drunk")
	obj54 = ns.Object("Aldwyn")
	obj52 = ns.Object("Geoff")
	obj63 = ns.Object("Ryan")
	obj64 = ns.Object("Ed")
	obj55 = ns.Object("Mayor_Theogrin")
	obj53 = ns.Object("Mayor's_Guard")
	obj51 = ns.Object("Barkeeper")
	obj65 = ns.Object("Bridge_Guard")
	obj48 = ns.Object("Henrick")
	obj47 = ns.Object("Jandor")
	obj75 = ns.Object("Jacob")
	obj76 = ns.Object("Morgan")
	obj78 = ns.Object("JacobTrigger")
	obj126 = ns.Object("MorganDoor")
	obj127 = ns.Object("CaveDoor")
	wp79 = ns.Waypoint("JacobWP")
	wp133 = ns.Waypoint("FaceWP")
	obj77 = ns.Object("Contest_Guard")
	obj56 = ns.Object("MayorGate1")
	obj57 = ns.Object("MayorGate2")
	obj58 = ns.Object("ContestGate1")
	obj59 = ns.Object("ContestGate2")
	obj120 = ns.Object("CemDoorL")
	obj121 = ns.Object("CemDoorR")
	obj118 = ns.Object("BasementDoor")
	obj60 = ns.Object("AldwinElevator")
	obj61 = ns.Object("RightTownGate")
	obj62 = ns.Object("LeftTownGate")
	obj72 = ns.Object("Bryan")
	obj73 = ns.Object("Clyde")
	obj74 = ns.Object("Tommy")
	obj66 = ns.Object("MagicShopElevator")
	obj67 = ns.Object("Joyce")
	obj68 = ns.Object("Julie")
	obj69 = ns.Object("Tanya")
	obj70 = ns.Object("Gretchen")
	obj50 = ns.Object("BlueKey")
	obj71 = ns.Object("AldwynFireplace")
	obj119 = ns.Object("AldwynsDoor")
	obj122 = ns.Object("TempDoorL")
	obj123 = ns.Object("TempDoorR")
	obj124 = ns.Object("Mystic")
	obj125 = ns.Object("Mystic2")
	wp80 = ns.Waypoint("AldwynHome")
	wp130 = ns.Waypoint("AldwynStorage")
	wp131 = ns.Waypoint("MysticWP")
	wp132 = ns.Waypoint("MysticStorage")
	gvar134 = ns.ObjectGroup("NPC's")
	gvar135 = ns.ObjectGroup("CharmedWolves")
	gvar136 = ns.WallGroup("MagicShopElevatorWalls")
	wp81 = ns.Waypoint("MineExitWP")
	wp82 = ns.Waypoint("CaveExitWP")
	ns.Wander(obj72)
	ns.Wander(obj73)
	ns.Wander(obj74)
	ns.Wander(obj68)
	ns.Wander(obj69)
	ns.LockDoor(obj56)
	ns.LockDoor(obj57)
	ns.LockDoor(obj120)
	ns.LockDoor(obj121)
	ns.LockDoor(obj61)
	ns.LockDoor(obj62)
	ns.LockDoor(obj126)
	ns.LockDoor(obj127)
	ns.LockDoor(obj58)
	ns.LockDoor(obj59)
	obj128[0] = 0
	obj128[1] = 0
	obj128[2] = 0
	obj128[3] = 0
	InversionInit()
	InitializeDialogs()
	ns.StartupScreen(8)
	ns.FrameTimer(1, OwnObjects)
	ns.FrameTimer(5, StartCaptainConversation)
	ns.FrameTimer(35, DisableElevators)
}
func MapEntry() {
	var v0 int
	v0 = ns.GetQuestStatus("Chapter8:HasWeirdling")
	ns.Music(15, 100)
	if v0 != 1 {
		goto LABEL1
	}
	flag116 = true
	ns.MoveObject(obj54, ns.GetWaypointX(wp130), ns.GetWaypointY(wp130))
	ns.LockDoor(obj119)
	ns.LockDoor(obj122)
	ns.LockDoor(obj123)
	InitializeSecondaryDialogs()
	ns.SetShopkeeperText(obj124, "Con08a:Mystic2")
	ns.CancelDialog(obj63)
	ns.CancelDialog(obj64)
LABEL1:
	if !(ns.CurrentHealth(obj84) > 0) {
		goto LABEL2
	}
	ns.Enchant(obj84, ns.ENCHANT_INVULNERABLE, 0)
LABEL2:
	ns.SetDialog(obj77, ns.NORMAL, ContestGuardStart, ContestGuardEnd)
	ns.StoryPic(obj77, "Townsman2Pic")
	ns.UnBlind()
}
func DebugOn() {
	flag115 = true
	ns.PrintToAll("Debug mode is enabled.")
}
func DebugOff() {
	flag115 = false
	ns.PrintToAll("Debug mode has been disabled.")
}
func GetWeirdlingValue() {
	var v0 int
	v0 = ns.GetQuestStatus("Con08a:HasWierdling")
	if v0 != 1 {
		goto LABEL1
	}
	ns.PrintToAll("The weirdleing has been retrieved!")
	goto LABEL2
LABEL1:
	ns.PrintToAll("The weirdling value = 0!")
LABEL2:
	return
}
func OpenCaveDoor() {
	ns.UnlockDoor(obj127)
	ns.ObjectOff(ns.GetTrigger())
}
func EnableRandomBump() {
	flag138 = true
}
func RandomBump() {
	if !(ns.IsCaller(ns.GetHost()) && flag138) {
		goto LABEL1
	}
	ns.Chat(ns.GetTrigger(), "Con02a:RandomBump")
	flag138 = false
	ns.FrameTimer(ivar137, EnableRandomBump)
LABEL1:
	return
}
func RestAwhile() {
	var v0 int
	v0 = ns.Random(60, 120)
	ns.PauseObject(ns.GetCaller(), v0)
}
func OpenMagicShopElevatorWalls() {
	ns.WallGroupOpen(gvar136)
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
	case "PlayerDeath":
		PlayerDeath()
	case "MapInitialize":
		MapInitialize()
	case "MapEntry":
		MapEntry()
	}
}
