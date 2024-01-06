package wiz08a

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
	ivar41  int
	gvar42  int
	ivar43  int
	ivar44  int
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
	wp76    ns.WaypointID
	wp77    ns.WaypointID
	wp78    ns.WaypointID
	wp79    ns.WaypointID
	obj80   ns.ObjectID
	obj81   ns.ObjectID
	obj82   ns.ObjectID
	obj83   ns.ObjectID
	obj84   ns.ObjectID
	obj85   ns.ObjectID
	wp86    ns.WaypointID
	wp87    ns.WaypointID
	gvar88  int
	gvar89  int
	gvar90  int
	gvar91  int
	gvar92  int
	gvar93  int
	gvar94  int
	gvar95  int
	ivar96  int
	flag97  bool
	flag98  bool
	flag99  bool
	flag100 bool
	flag101 bool
	flag102 bool
	flag103 bool
	flag104 bool
	flag105 bool
	obj106  ns.ObjectID
	gvar107 int
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
	obj118  [4]ns.ObjectID
	gvar119 int
	wp120   ns.WaypointID
	wp121   ns.WaypointID
	wp122   ns.WaypointID
	wp123   ns.WaypointID
	gvar124 ns.ObjectGroupID
	gvar125 ns.ObjectGroupID
	gvar126 ns.WallGroupID
	flag127 bool
	ivar128 int
	flag129 bool
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
	gvar30 = gvar13
	gvar31 = gvar10
	gvar32 = gvar7
	gvar33 = gvar14
	gvar34 = gvar18
	gvar35 = gvar22
	gvar36 = gvar23
	gvar37 = gvar24
	gvar38 = gvar25
	gvar39 = gvar26
	gvar40 = gvar4
	ivar41 = 200
	gvar42 = 2
	ivar43 = 0
	ivar44 = 0
	gvar88 = 0
	gvar89 = 1
	gvar90 = 2
	gvar91 = 3
	gvar92 = 4
	gvar93 = 5
	gvar94 = 6
	gvar95 = gvar88
	ivar96 = 0
	flag97 = false
	flag98 = false
	flag99 = false
	flag100 = false
	flag101 = false
	flag102 = false
	flag103 = true
	flag104 = false
	flag105 = false
	flag127 = false
	ivar128 = 60
	flag129 = true
}
func JandorStart() {
	var v0 int
	v0 = gvar40
	if v0 == gvar4 {
		goto LABEL1
	}
	if v0 == gvar5 {
		goto LABEL2
	}
	if v0 == gvar6 {
		goto LABEL3
	}
	goto LABEL4
LABEL1:
	ns.TellStory(ns.HumanMaleEatFood, "Wiz08a:CaptainGreet")
	goto LABEL4
LABEL2:
	ns.TellStory(ns.SwordsmanHurt, "Con03A.scr:JandorB")
	goto LABEL4
LABEL3:
	ns.Frozen(ns.GetHost(), true)
	ns.Blind()
	ns.FrameTimer(45, ExitToSwamp)
	goto LABEL4
LABEL4:
	return
}
func JandorEnd() {
	var v0 int
	v0 = gvar40
	if v0 == gvar4 {
		goto LABEL1
	}
	if v0 == gvar5 {
		goto LABEL2
	}
	if v0 == gvar6 {
		goto LABEL3
	}
	goto LABEL4
LABEL1:
	goto LABEL4
LABEL2:
	gvar40 = gvar6
	goto LABEL4
LABEL3:
	ns.Frozen(ns.GetHost(), true)
	ns.Blind()
	ns.FrameTimer(45, ExitToSwamp)
	goto LABEL4
LABEL4:
	return
}
func JacobDialogStart() {
	var (
		v0 int
		v1 int
		v2 int
	)
	v2 = gvar39
	if v2 == gvar26 {
		goto LABEL1
	}
	if v2 == gvar27 {
		goto LABEL2
	}
	if v2 == gvar28 {
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
	v0 = gvar39
	if v0 == gvar27 {
		goto LABEL1
	}
	if v0 == gvar28 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.Frozen(ns.GetHost(), false)
	goto LABEL3
LABEL2:
	ns.Frozen(ns.GetHost(), false)
	goto LABEL3
LABEL3:
	return
}
func ResetJacobDialog() {
	if !ns.IsCaller(obj73) {
		goto LABEL1
	}
	if gvar39 != gvar29 {
		goto LABEL1
	}
	gvar39 = gvar26
	ns.SetDialog(obj73, ns.NORMAL, JacobDialogStart, JacobDialogEnd)
	ivar44 = 0
	ns.ObjectOff(obj75)
LABEL1:
	return
}
func MorganInjured() {
	var v0 int
	ivar44 += 1
	v0 = ivar44
	if v0 == 1 {
		goto LABEL1
	}
	if v0 == 2 {
		goto LABEL2
	}
	if v0 == 3 {
		goto LABEL3
	}
	goto LABEL4
LABEL1:
	gvar39 = gvar27
	ns.Frozen(ns.GetHost(), true)
	ns.CreatureIdle(ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj73)
	ns.LookAtObject(obj73, ns.GetHost())
	if !(ns.CurrentHealth(obj73) > 0) {
		goto LABEL5
	}
	ns.StartDialog(obj73, ns.GetHost())
LABEL5:
	goto LABEL4
LABEL2:
	gvar39 = gvar28
	ns.Frozen(ns.GetHost(), true)
	ns.CreatureIdle(ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj73)
	ns.LookAtObject(obj73, ns.GetHost())
	if !(ns.CurrentHealth(obj73) > 0) {
		goto LABEL6
	}
	ns.StartDialog(obj73, ns.GetHost())
LABEL6:
	goto LABEL4
LABEL3:
	gvar39 = gvar29
	ns.CancelDialog(obj73)
	ns.ClearOwner(obj73)
	ns.AggressionLevel(obj73, 0.83)
	if !(ns.CurrentHealth(obj73) > 0) {
		goto LABEL7
	}
	ns.Attack(obj73, ns.GetHost())
LABEL7:
	goto LABEL4
LABEL4:
	return
}
func PlayerFlee() {
	if gvar39 != gvar29 {
		goto LABEL1
	}
	ns.ObjectOn(obj75)
	ns.SetOwner(ns.GetHost(), obj73)
	ns.AggressionLevel(obj73, 0)
	ns.CreatureGuard(obj73, ns.GetWaypointX(wp77), ns.GetWaypointY(wp77), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
LABEL1:
	return
}
func HenrickDialogStart() {
	var v0 int
	v0 = gvar31
	if v0 == gvar10 {
		goto LABEL1
	}
	if v0 == gvar11 {
		goto LABEL2
	}
	if v0 == gvar12 {
		goto LABEL3
	}
	goto LABEL4
LABEL1:
	ns.LookAtObject(obj46, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War08b:HenrickSalesPitchA")
	goto LABEL4
LABEL2:
	ns.LookAtObject(obj46, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War08b:HenrickSalesPitchB")
	goto LABEL4
LABEL3:
	ns.LookAtObject(obj46, ns.GetHost())
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
	v7 = gvar31
	if v7 == gvar10 {
		goto LABEL1
	}
	if v7 == gvar11 {
		goto LABEL2
	}
	if v7 == gvar12 {
		goto LABEL3
	}
	goto LABEL4
LABEL1:
	v1 = ns.GetAnswer(obj46)
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
	if !(v0 < ivar41) {
		goto LABEL9
	}
	ns.SetDialog(obj46, ns.NORMAL, NullDialogStart, ResetHenrickDialogChoice)
	ns.TellStory(ns.SwordsmanHurt, "War08b:HenrickNotEnoughGold")
	goto LABEL10
LABEL9:
	ns.ChangeGold(ns.GetHost(), -ivar41)
	gvar31 = gvar11
	ReleaseCharmedWolf(ivar43)
	ivar43 += 1
	ns.SetDialog(obj46, ns.NORMAL, NullDialogStart, ResetHenrickDialogChoice)
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
	v1 = ns.GetAnswer(obj46)
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
	if !(v0 < ivar41) {
		goto LABEL15
	}
	ns.SetDialog(obj46, ns.NORMAL, NullDialogStart, ResetHenrickDialogChoice)
	ns.TellStory(ns.SwordsmanHurt, "War08b:HenrickNotEnoughGold")
	goto LABEL16
LABEL15:
	ns.ChangeGold(ns.GetHost(), -ivar41)
	ReleaseCharmedWolf(ivar43)
	ivar43 += 1
	ns.SetDialog(obj46, ns.NORMAL, NullDialogStart, ResetHenrickDialogChoice)
	ns.TellStory(ns.SwordsmanHurt, "War08b:HenrickSaleSuccessful")
	if !(ivar43 >= 2) {
		goto LABEL16
	}
	gvar31 = gvar12
	ns.SetDialog(obj46, ns.NORMAL, HenrickDialogStart, HenrickDialogEnd)
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
func ResetHenrickDialogChoice() {
	ns.SetDialog(obj46, ns.YESNO, HenrickDialogStart, HenrickDialogEnd)
}
func DrunkDialogStart() {
	ns.LookAtObject(obj47, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:DrunkGreeting")
}
func DrunkDialogEnd() {
}
func GatekeeperDialogStart() {
	var v0 int
	ns.LookAtObject(obj50, ns.GetHost())
	v0 = gvar35
	if v0 == gvar22 {
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
	ns.LookAtObject(obj61, ns.GetHost())
	v0 = gvar36
	if v0 == gvar23 {
		goto LABEL1
	}
	goto LABEL2
LABEL1:
	ns.TellStory(ns.SwordsmanHurt, "Wiz08a:Guard01Greet")
	goto LABEL2
LABEL2:
	return
}
func Gatekeeper2DialogEnd() {
}
func Gatekeeper3DialogStart() {
	var v0 int
	ns.LookAtObject(obj62, ns.GetHost())
	v0 = gvar37
	if v0 == gvar24 {
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
	ns.LookAtObject(ns.GetHost(), obj53)
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
	v0 = gvar30
	if v0 == gvar13 {
		goto LABEL1
	}
	goto LABEL2
LABEL1:
	ns.LookAtObject(obj51, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War03b:MayorsGuardIntro")
	goto LABEL2
LABEL2:
	return
}
func MayorsGuardDialogEnd() {
}
func AldwinDialogStart() {
	var v0 int
	v0 = gvar32
	if v0 == gvar7 {
		goto LABEL1
	}
	if v0 == gvar8 {
		goto LABEL2
	}
	if v0 == gvar9 {
		goto LABEL3
	}
	goto LABEL4
LABEL1:
	ns.LookAtObject(obj52, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Wiz08a:AldwynGreet")
	goto LABEL4
LABEL2:
	ns.LookAtObject(obj52, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War08a:AldwynProd")
	goto LABEL4
LABEL3:
	ns.LookAtObject(obj52, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con08:AldwynProd")
	goto LABEL4
LABEL4:
	return
}
func AldwinDialogEnd() {
	var v0 int
	v0 = gvar32
	if v0 == gvar7 {
		goto LABEL1
	}
	if v0 == gvar8 {
		goto LABEL2
	}
	if v0 == gvar9 {
		goto LABEL3
	}
	goto LABEL4
LABEL1:
	ns.Pickup(ns.GetHost(), obj48)
	ns.PrintToAll("GeneralPrint:GainedKey")
	gvar32 = gvar8
	ns.JournalEdit(ns.GetHost(), "Chapter8LocateAldwyn", 4)
	ns.PrintToAll("Con02a:ObjectiveComplete")
	ns.JournalEntry(ns.GetHost(), "Chapter8IxPriest", 2)
	ns.PrintToAll("Con01a:NewJournalEntry")
	goto LABEL4
LABEL2:
	gvar32 = gvar9
	goto LABEL4
LABEL3:
	goto LABEL4
LABEL4:
	return
}
func BridgeGuardDialogStart() {
	var v0 int
	v0 = gvar38
	if v0 == gvar25 {
		goto LABEL1
	}
	goto LABEL2
LABEL1:
	ns.LookAtObject(obj63, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con02a:RandomSay")
	goto LABEL2
LABEL2:
	return
}
func BridgeGuardDialogEnd() {
	var v0 int
	v0 = gvar38
	if v0 == gvar25 {
		goto LABEL1
	}
	goto LABEL2
LABEL1:
	goto LABEL2
LABEL2:
	return
}
func Maiden1DialogStart() {
	ns.LookAtObject(obj65, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con02a:BarMaiden")
}
func Maiden1DialogEnd() {
}
func Maiden2DialogStart() {
	ns.Frozen(ns.GetHost(), true)
	ns.Frozen(obj66, true)
	ns.CreatureIdle(ns.GetHost())
	ns.CreatureIdle(obj66)
	ns.LookAtObject(ns.GetHost(), obj66)
	ns.LookAtObject(obj66, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "War06A:WomenSpeak2")
}
func Maiden2DialogEnd() {
	ns.Frozen(ns.GetHost(), false)
	ns.Frozen(obj66, false)
	ns.Wander(obj66)
}
func Maiden3DialogStart() {
	ns.Frozen(ns.GetHost(), true)
	ns.Frozen(obj67, true)
	ns.CreatureIdle(ns.GetHost())
	ns.CreatureIdle(obj67)
	ns.LookAtObject(ns.GetHost(), obj67)
	ns.LookAtObject(obj67, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02A.scr:MaidenTalk05")
}
func Maiden3DialogEnd() {
	ns.Frozen(ns.GetHost(), false)
	ns.Frozen(obj67, false)
	ns.Wander(obj67)
}
func Maiden4DialogStart() {
	var (
		v0 int
		v1 int
	)
	ns.Frozen(ns.GetHost(), true)
	ns.Frozen(obj68, true)
	ns.CreatureIdle(ns.GetHost())
	ns.CreatureIdle(obj68)
	ns.LookAtObject(ns.GetHost(), obj68)
	ns.LookAtObject(obj68, ns.GetHost())
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
	ns.Frozen(obj68, false)
	ns.Wander(obj68)
}
func Townsman1DialogStart() {
	ns.Frozen(ns.GetHost(), true)
	ns.Frozen(obj70, true)
	ns.CreatureIdle(ns.GetHost())
	ns.CreatureIdle(obj70)
	ns.LookAtObject(ns.GetHost(), obj70)
	ns.LookAtObject(obj70, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "War02a:NewTownsman3")
}
func Townsman1DialogEnd() {
	ns.Frozen(ns.GetHost(), false)
	ns.Frozen(obj70, false)
	ns.Wander(obj70)
}
func Townsman2DialogStart() {
	ns.Frozen(ns.GetHost(), true)
	ns.Frozen(obj71, true)
	ns.CreatureIdle(ns.GetHost())
	ns.CreatureIdle(obj71)
	ns.LookAtObject(ns.GetHost(), obj71)
	ns.LookAtObject(obj71, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02A.scr:AlbiTalk01")
}
func Townsman2DialogEnd() {
	ns.Frozen(ns.GetHost(), false)
	ns.Frozen(obj71, false)
	ns.Wander(obj71)
}
func Townsman3DialogStart() {
	ns.Frozen(ns.GetHost(), true)
	ns.Frozen(obj72, true)
	ns.CreatureIdle(ns.GetHost())
	ns.CreatureIdle(obj72)
	ns.LookAtObject(ns.GetHost(), obj72)
	ns.LookAtObject(obj72, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02A.scr:MaxRumor03")
}
func Townsman3DialogEnd() {
	ns.Frozen(ns.GetHost(), false)
	ns.Frozen(obj72, false)
	ns.Wander(obj72)
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
	ns.WallOpen(ns.Wall(128, 106))
	ns.EnchantOff(obj118[a1], ns.ENCHANT_INVULNERABLE)
	ns.CreatureFollow(obj118[a1], ns.GetHost())
	ns.PauseObject(obj118[a1], 30)
	ns.AggressionLevel(obj118[a1], 0.83)
	ns.MakeFriendly(obj118[a1])
	ns.BecomePet(obj118[a1])
	goto LABEL5
LABEL2:
	ns.WallOpen(ns.Wall(130, 104))
	ns.EnchantOff(obj118[a1], ns.ENCHANT_INVULNERABLE)
	ns.CreatureFollow(obj118[a1], ns.GetHost())
	ns.PauseObject(obj118[a1], 30)
	ns.AggressionLevel(obj118[a1], 0.83)
	ns.MakeFriendly(obj118[a1])
	ns.BecomePet(obj118[a1])
	goto LABEL5
LABEL3:
	ns.WallOpen(ns.Wall(132, 102))
	ns.EnchantOff(obj118[a1], ns.ENCHANT_INVULNERABLE)
	ns.CreatureFollow(obj118[a1], ns.GetHost())
	ns.PauseObject(obj118[a1], 30)
	ns.AggressionLevel(obj118[a1], 0.83)
	ns.MakeFriendly(obj118[a1])
	ns.BecomePet(obj118[a1])
	goto LABEL5
LABEL4:
	ns.WallOpen(ns.Wall(134, 100))
	ns.EnchantOff(obj118[a1], ns.ENCHANT_INVULNERABLE)
	ns.CreatureFollow(obj118[a1], ns.GetHost())
	ns.PauseObject(obj118[a1], 30)
	ns.AggressionLevel(obj118[a1], 0.83)
	ns.MakeFriendly(obj118[a1])
	ns.BecomePet(obj118[a1])
	goto LABEL5
LABEL5:
	ns.PauseObject(obj118[a1], 30)
}
func InitializeDialogs() {
	ns.StoryPic(obj46, "HenrickPic")
	ns.StoryPic(obj47, "DrunkPic")
	ns.StoryPic(obj49, "Townsman1Pic")
	ns.StoryPic(obj50, "IxGuard2Pic")
	ns.StoryPic(obj51, "Warrior3Pic")
	ns.StoryPic(obj53, "TheogrinPic")
	ns.StoryPic(obj52, "AldwynPic")
	ns.StoryPic(obj61, "IxGuard2Pic")
	ns.StoryPic(obj62, "IxGuard1Pic")
	ns.StoryPic(obj63, "Townsman3Pic")
	ns.StoryPic(obj65, "MaidenPic")
	ns.StoryPic(obj66, "MaidenPic2")
	ns.StoryPic(obj67, "MaidenPic3")
	ns.StoryPic(obj68, "MaidenPic4")
	ns.StoryPic(obj70, "MalePic4")
	ns.StoryPic(obj71, "MalePic11")
	ns.StoryPic(obj72, "MalePic12")
	ns.StoryPic(obj45, "AirshipCaptainPic")
	ns.StoryPic(obj73, "WardenPic")
	ns.SetDialog(obj73, ns.NORMAL, JacobDialogStart, JacobDialogEnd)
	ns.SetDialog(obj46, ns.YESNO, HenrickDialogStart, HenrickDialogEnd)
	ns.SetDialog(obj45, ns.NORMAL, JandorStart, JandorEnd)
	ns.SetDialog(obj47, ns.NORMAL, DrunkDialogStart, DrunkDialogEnd)
	ns.SetDialog(obj61, ns.NORMAL, Gatekeeper2DialogStart, Gatekeeper2DialogEnd)
	ns.SetDialog(obj62, ns.NORMAL, Gatekeeper3DialogStart, Gatekeeper3DialogEnd)
	ns.SetDialog(obj51, ns.NORMAL, MayorsGuardDialogStart, MayorsGuardDialogEnd)
	ns.SetDialog(obj52, ns.NORMAL, AldwinDialogStart, AldwinDialogEnd)
	ns.SetDialog(obj63, ns.NORMAL, BridgeGuardDialogStart, BridgeGuardDialogEnd)
	ns.SetDialog(obj65, ns.NORMAL, Maiden1DialogStart, Maiden1DialogEnd)
	ns.SetDialog(obj66, ns.NORMAL, Maiden2DialogStart, Maiden2DialogEnd)
	ns.SetDialog(obj67, ns.NORMAL, Maiden3DialogStart, Maiden3DialogEnd)
	ns.SetDialog(obj68, ns.NORMAL, Maiden4DialogStart, Maiden4DialogEnd)
	ns.SetDialog(obj70, ns.NORMAL, Townsman1DialogStart, Townsman1DialogEnd)
	ns.SetDialog(obj71, ns.NORMAL, Townsman2DialogStart, Townsman2DialogEnd)
	ns.SetDialog(obj72, ns.NORMAL, Townsman3DialogStart, Townsman3DialogEnd)
}
func InitializeSecondaryDialogs() {
	gvar40 = gvar5
}
func NullDialogStart() {
}
func ExitToSwamp() {
	ns.Frozen(ns.GetHost(), false)
	ns.MoveObject(ns.GetHost(), ns.GetWaypointX(ns.Waypoint("SwampExitWP")), ns.GetWaypointY(ns.Waypoint("SwampExitWP")))
}
func PlayerDeath() {
	ns.DeathScreen(8)
}
func GoToMineSEG2() {
	ns.MoveObject(ns.GetHost(), ns.GetWaypointX(wp78), ns.GetWaypointY(wp78))
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
	ns.MoveObject(ns.GetHost(), ns.GetWaypointX(wp79), ns.GetWaypointY(wp79))
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
	obj80 = ns.Object("Teacher")
	obj81 = ns.Object("Apprentice")
	obj82 = ns.Object("InvertSpell")
	obj83 = ns.Object("LessonGate1")
	obj84 = ns.Object("LessonGate2")
	obj85 = ns.Object("ApprenticeTrigger")
	wp86 = ns.Waypoint("PlayerSounds")
	wp87 = ns.Waypoint("ApprenticeWP")
	ns.StoryPic(obj80, "WizardGuard1Pic")
	ns.SetDialog(obj80, ns.YESNO, TeacherStart, TeacherEnd)
	ns.LockDoor(obj83)
	ns.LockDoor(obj84)
	ns.ObjectOff(obj85)
	ns.FrameTimer(1, InversionOwnObject)
}
func InversionOwnObject() {
	ns.SetOwner(ns.GetHost(), obj80)
	ns.Enchant(obj80, ns.ENCHANT_ANCHORED, 0)
	ns.Enchant(obj81, ns.ENCHANT_ANCHORED, 0)
}
func TeacherStart() {
	var v0 int
	v0 = gvar95
	if v0 == gvar88 {
		goto LABEL1
	}
	if v0 == gvar89 {
		goto LABEL2
	}
	if v0 == gvar90 {
		goto LABEL3
	}
	if v0 == gvar91 {
		goto LABEL4
	}
	if v0 == gvar92 {
		goto LABEL5
	}
	if v0 == gvar93 {
		goto LABEL6
	}
	if v0 == gvar94 {
		goto LABEL7
	}
	goto LABEL8
LABEL1:
	ns.Frozen(ns.GetHost(), true)
	ns.CreatureIdle(ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj80)
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
	v6 = gvar95
	if v6 == gvar88 {
		goto LABEL1
	}
	if v6 == gvar89 {
		goto LABEL2
	}
	if v6 == gvar90 {
		goto LABEL3
	}
	if v6 == gvar91 {
		goto LABEL4
	}
	if v6 == gvar92 {
		goto LABEL5
	}
	if v6 == gvar93 {
		goto LABEL6
	}
	if v6 == gvar94 {
		goto LABEL7
	}
	goto LABEL8
LABEL1:
	v1 = ns.GetAnswer(obj80)
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
	ns.SetDialog(obj80, ns.NORMAL, TeacherStart, TeacherEnd)
	gvar95 = gvar89
	ns.StartDialog(obj80, ns.GetHost())
	goto LABEL12
	goto LABEL10
LABEL13:
	ns.ChangeGold(ns.GetHost(), -100)
	ns.UnlockDoor(obj83)
	ns.UnlockDoor(obj84)
	ns.SetDialog(obj80, ns.NORMAL, TeacherStart, TeacherEnd)
	gvar95 = gvar90
	ns.StartDialog(obj80, ns.GetHost())
	goto LABEL12
LABEL10:
	ns.Frozen(ns.GetHost(), false)
	gvar95 = gvar88
	ns.SetDialog(obj80, ns.YESNO, TeacherStart, TeacherEnd)
	goto LABEL12
LABEL11:
	ns.Frozen(ns.GetHost(), false)
	gvar95 = gvar88
	ns.SetDialog(obj80, ns.YESNO, TeacherStart, TeacherEnd)
	goto LABEL12
LABEL12:
	goto LABEL8
LABEL2:
	ns.Frozen(ns.GetHost(), false)
	gvar95 = gvar88
	ns.SetDialog(obj80, ns.YESNO, TeacherStart, TeacherEnd)
	goto LABEL8
LABEL3:
	ns.CancelDialog(obj80)
	ns.Frozen(ns.GetHost(), false)
	ns.Pickup(ns.GetHost(), obj82)
	goto LABEL8
LABEL4:
	ns.CancelDialog(obj80)
	ns.SecondTimer(1, Phoneme4)
	goto LABEL8
LABEL5:
	ns.UnlockDoor(obj56)
	ns.UnlockDoor(obj57)
	ns.UnlockDoor(obj83)
	ns.UnlockDoor(obj84)
	gvar95 = gvar94
	ns.SetDialog(obj80, ns.NORMAL, TeacherStart, TeacherEnd)
	goto LABEL8
LABEL6:
	ns.UnlockDoor(obj56)
	ns.UnlockDoor(obj57)
	ns.UnlockDoor(obj83)
	ns.UnlockDoor(obj84)
	gvar95 = gvar94
	ns.SetDialog(obj80, ns.NORMAL, TeacherStart, TeacherEnd)
	goto LABEL8
LABEL7:
	gvar95 = gvar94
	ns.SetDialog(obj80, ns.NORMAL, TeacherStart, TeacherEnd)
	goto LABEL8
LABEL8:
	return
}
func Phoneme1() {
	if flag101 {
		goto LABEL1
	}
	ns.MoveWaypoint(wp86, 1326, 2493)
	ns.AudioEvent(ns.SpellPhonemeDown, wp86)
	ns.FrameTimer(10, Phoneme2)
LABEL1:
	return
}
func Phoneme2() {
	if flag101 {
		goto LABEL1
	}
	ns.AudioEvent(ns.SpellPhonemeDown, wp86)
	ns.FrameTimer(10, Phoneme3)
LABEL1:
	return
}
func Phoneme3() {
	if flag101 {
		goto LABEL1
	}
	ns.AudioEvent(ns.SpellPhonemeDown, wp86)
	ns.FrameTimer(10, CastSpells)
LABEL1:
	return
}
func Phoneme4() {
	if !flag101 {
		goto LABEL1
	}
	ns.MoveWaypoint(wp86, 1326, 2493)
	ns.AudioEvent(ns.SpellPhonemeDown, wp86)
	ns.FrameTimer(10, Phoneme5)
LABEL1:
	return
}
func Phoneme5() {
	if !flag101 {
		goto LABEL1
	}
	ns.AudioEvent(ns.SpellPhonemeDown, wp86)
	ns.FrameTimer(10, Phoneme6)
LABEL1:
	return
}
func Phoneme6() {
	if !flag101 {
		goto LABEL1
	}
	ns.AudioEvent(ns.SpellPhonemeDown, wp86)
	ns.FrameTimer(10, CastSpells2)
LABEL1:
	return
}
func LessonBegin() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	if gvar95 != gvar90 {
		goto LABEL1
	}
	if !flag103 {
		goto LABEL2
	}
	ns.LockDoor(obj83)
	ns.LockDoor(obj84)
	ns.Move(obj81, wp87)
	flag103 = false
LABEL2:
	if !(flag97 == false && flag102 == true) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.LookAtObject(obj81, ns.GetHost())
	flag97 = true
	ns.LockDoor(obj56)
	ns.LockDoor(obj57)
	ns.SecondTimer(1, Phoneme1)
LABEL1:
	return
}
func CastSpells() {
	if flag101 {
		goto LABEL1
	}
	if !(ivar96 < 4) {
		goto LABEL2
	}
	ns.CastSpellObjectLocation(ns.SPELL_SLOW, obj81, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
LABEL2:
	if ivar96 != 4 {
		goto LABEL3
	}
	ns.CastSpellObjectLocation(ns.SPELL_SLOW, obj81, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
LABEL3:
	ivar96 += 1
	if !(ivar96 < 5) {
		goto LABEL4
	}
	ns.SecondTimer(3, Phoneme1)
	goto LABEL1
LABEL4:
	flag98 = true
	if !flag100 {
		goto LABEL5
	}
	gvar95 = gvar92
	ns.SetDialog(obj80, ns.NORMAL, TeacherStart, TeacherEnd)
	ns.StartDialog(obj80, ns.GetHost())
	goto LABEL1
LABEL5:
	gvar95 = gvar93
	ns.SetDialog(obj80, ns.NORMAL, TeacherStart, TeacherEnd)
	ns.StartDialog(obj80, ns.GetHost())
LABEL1:
	return
}
func CastSpells2() {
	if !flag101 {
		goto LABEL1
	}
	if !(ivar96 < 4) {
		goto LABEL2
	}
	ns.CastSpellObjectLocation(ns.SPELL_SLOW, obj81, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
LABEL2:
	if ivar96 != 4 {
		goto LABEL3
	}
	ns.CastSpellObjectLocation(ns.SPELL_SLOW, obj81, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
LABEL3:
	ivar96 += 1
	if !(ivar96 < 5) {
		goto LABEL4
	}
	ns.SecondTimer(3, Phoneme4)
	goto LABEL1
LABEL4:
	flag98 = true
	if !flag100 {
		goto LABEL5
	}
	gvar95 = gvar92
	ns.SetDialog(obj80, ns.NORMAL, TeacherStart, TeacherEnd)
	ns.StartDialog(obj80, ns.GetHost())
	goto LABEL1
LABEL5:
	gvar95 = gvar93
	ns.SetDialog(obj80, ns.NORMAL, TeacherStart, TeacherEnd)
	ns.StartDialog(obj80, ns.GetHost())
LABEL1:
	return
}
func CheckApprentice() {
	if !ns.HasEnchant(obj81, ns.ENCHANT_SLOWED) {
		goto LABEL1
	}
	ns.ObjectOff(obj85)
	flag101 = true
	gvar95 = gvar91
	flag100 = true
	if !(ivar96 > 2) {
		goto LABEL2
	}
	ivar96 = 2
LABEL2:
	ns.SetDialog(obj80, ns.NORMAL, TeacherStart, TeacherEnd)
	ns.StartDialog(obj80, ns.GetHost())
LABEL1:
	if !(ns.CurrentHealth(obj81) <= 0) {
		goto LABEL3
	}
	ns.ObjectOff(ns.GetTrigger())
	flag100 = true
LABEL3:
	return
}
func ApprenticeDie() {
	ns.AudioEvent(ns.HumanMaleHurtHeavy, wp86)
}
func ApprenticeReport() {
	ns.CreatureGuard(obj81, ns.GetWaypointX(wp87), ns.GetWaypointY(wp87), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	ns.ObjectOn(obj85)
	flag102 = true
}
func OwnObjects() {
	ns.GroupSetOwner(ns.GetHost(), gvar124)
	ns.SetOwner(ns.GetHost(), obj46)
	ns.SetOwner(ns.GetHost(), obj106)
	ns.GroupSetOwner(ns.GetHost(), gvar125)
	ns.SetOwner(ns.GetHost(), obj47)
	ns.SetOwner(ns.GetHost(), obj65)
	ns.SetOwner(ns.GetHost(), obj66)
	ns.SetOwner(ns.GetHost(), obj67)
	ns.SetOwner(ns.GetHost(), obj68)
	ns.SetOwner(ns.GetHost(), obj45)
	ns.SetOwner(ns.GetHost(), obj73)
	ns.StoryPic(obj49, "ShopKeeperBrownPic")
}
func FirstJournalEntry() {
	if flag127 {
		goto LABEL1
	}
	flag127 = true
	ns.JournalEntry(ns.GetHost(), "Chapter8LocateAldwyn", 2)
	ns.PrintToAll("Con01a:NewJournalEntry")
LABEL1:
	return
}
func DisableElevators() {
	ns.ObjectOff(obj58)
	ns.ObjectOff(obj64)
}
func MapInitialize() {
	var v0 int
	obj46 = ns.Object("Henrick")
	obj106 = ns.Object("HenricksWolf")
	obj118[0] = ns.Object("CharmedWolf01")
	obj118[1] = ns.Object("CharmedWolf02")
	obj118[2] = ns.Object("CharmedWolf03")
	obj118[3] = ns.Object("CharmedWolf04")
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL1
		}
		ns.Enchant(obj118[v0], ns.ENCHANT_INVULNERABLE, 0)
		v0 += 1
	}
LABEL1:
	obj47 = ns.Object("Drunk")
	obj52 = ns.Object("Aldwyn")
	obj50 = ns.Object("Geoff")
	obj61 = ns.Object("Ryan")
	obj62 = ns.Object("Ed")
	obj53 = ns.Object("Mayor_Theogrin")
	obj51 = ns.Object("Mayor's_Guard")
	obj49 = ns.Object("Barkeeper")
	obj63 = ns.Object("Bridge_Guard")
	obj45 = ns.Object("Jandor")
	obj73 = ns.Object("Jacob")
	obj74 = ns.Object("Morgan")
	obj117 = ns.Object("CaveDoor")
	obj114 = ns.Object("MorganDoor")
	obj54 = ns.Object("MayorGate1")
	obj55 = ns.Object("MayorGate2")
	obj56 = ns.Object("ContestGate1")
	obj57 = ns.Object("ContestGate2")
	obj110 = ns.Object("CemeteryGate01")
	obj111 = ns.Object("CemeteryGate02")
	obj108 = ns.Object("BasementDoor")
	obj58 = ns.Object("AldwinElevator")
	obj59 = ns.Object("RightTownGate")
	obj60 = ns.Object("LeftTownGate")
	obj70 = ns.Object("Bryan")
	obj71 = ns.Object("Clyde")
	obj72 = ns.Object("Tommy")
	obj64 = ns.Object("MagicShopElevator")
	obj65 = ns.Object("Joyce")
	obj66 = ns.Object("Julie")
	obj67 = ns.Object("Tanya")
	obj68 = ns.Object("Gretchen")
	obj48 = ns.Object("BlueKey")
	obj69 = ns.Object("AldwynFireplace")
	obj75 = ns.Object("JacobTrigger")
	obj109 = ns.Object("AldwynsDoor")
	obj112 = ns.Object("TempDoorR")
	obj113 = ns.Object("TempDoorL")
	obj115 = ns.Object("Mystic")
	obj116 = ns.Object("Mystic2")
	wp76 = ns.Waypoint("AldwynHome")
	wp120 = ns.Waypoint("AldwynStorage")
	wp77 = ns.Waypoint("JacobWP")
	wp121 = ns.Waypoint("MysticWP")
	wp122 = ns.Waypoint("MysticStorage")
	wp123 = ns.Waypoint("FaceWP")
	gvar124 = ns.ObjectGroup("NPC's")
	gvar125 = ns.ObjectGroup("CharmedWolves")
	gvar126 = ns.WallGroup("MagicShopElevatorWalls")
	wp78 = ns.Waypoint("MineExitWP")
	wp79 = ns.Waypoint("CaveExitWP")
	ns.Wander(obj70)
	ns.Wander(obj71)
	ns.Wander(obj72)
	ns.Wander(obj66)
	ns.Wander(obj67)
	ns.LockDoor(obj54)
	ns.LockDoor(obj55)
	ns.LockDoor(obj110)
	ns.LockDoor(obj111)
	ns.LockDoor(obj59)
	ns.LockDoor(obj60)
	ns.LockDoor(obj114)
	ns.LockDoor(obj117)
	InversionInit()
	InitializeDialogs()
	ns.StartupScreen(8)
	ns.FrameTimer(1, OwnObjects)
	ns.FrameTimer(1, FirstJournalEntry)
	ns.FrameTimer(35, DisableElevators)
}
func MapEntry() {
	var v0 int
	v0 = ns.GetQuestStatus("Chapter8:HasWeirdling")
	ns.Music(15, 100)
	if v0 != 1 {
		goto LABEL1
	}
	flag105 = true
	ns.MoveObject(obj52, ns.GetWaypointX(wp120), ns.GetWaypointY(wp120))
	ns.LockDoor(obj109)
	ns.LockDoor(obj112)
	ns.LockDoor(obj113)
	InitializeSecondaryDialogs()
	ns.SetShopkeeperText(obj115, "Con08a:Mystic2")
	ns.CancelDialog(obj61)
	ns.CancelDialog(obj62)
LABEL1:
	ns.UnBlind()
}
func DebugOn() {
	flag104 = true
	ns.PrintToAll("Debug mode is enabled.")
}
func DebugOff() {
	flag104 = false
	ns.PrintToAll("Debug mode has been disabled.")
}
func OpenCaveDoor() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.UnlockDoor(obj117)
	ns.ObjectOff(ns.GetTrigger())
LABEL1:
	return
}
func EnableRandomBump() {
	flag129 = true
}
func RandomBump() {
	if !(ns.IsCaller(ns.GetHost()) && flag129) {
		goto LABEL1
	}
	ns.Chat(ns.GetTrigger(), "Con02a:RandomBump")
	flag129 = false
	ns.FrameTimer(ivar128, EnableRandomBump)
LABEL1:
	return
}
func RestAwhile() {
	var v0 int
	v0 = ns.Random(60, 120)
	ns.PauseObject(ns.GetCaller(), v0)
}
func OpenMagicShopElevatorWalls() {
	ns.WallGroupOpen(gvar126)
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
