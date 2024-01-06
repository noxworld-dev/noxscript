package war08a

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
	wp78    ns.WaypointID
	wp79    ns.WaypointID
	wp80    ns.WaypointID
	wp81    ns.WaypointID
	obj82   ns.ObjectID
	obj83   ns.ObjectID
	obj84   ns.ObjectID
	obj85   ns.ObjectID
	obj86   ns.ObjectID
	obj87   ns.ObjectID
	wp88    ns.WaypointID
	wp89    ns.WaypointID
	gvar90  int
	gvar91  int
	gvar92  int
	gvar93  int
	gvar94  int
	gvar95  int
	gvar96  int
	gvar97  int
	ivar98  int
	flag99  bool
	flag100 bool
	flag101 bool
	flag102 bool
	flag103 bool
	flag104 bool
	flag105 bool
	flag106 bool
	flag107 bool
	obj108  ns.ObjectID
	gvar109 int
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
	obj120  [4]ns.ObjectID
	gvar121 int
	wp122   ns.WaypointID
	wp123   ns.WaypointID
	wp124   ns.WaypointID
	wp125   ns.WaypointID
	gvar126 ns.ObjectGroupID
	gvar127 ns.ObjectGroupID
	gvar128 ns.WallGroupID
	ivar129 int
	flag130 bool
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
	gvar43 = 2
	ivar44 = 0
	ivar45 = 0
	flag46 = false
	gvar90 = 0
	gvar91 = 1
	gvar92 = 2
	gvar93 = 3
	gvar94 = 4
	gvar95 = 5
	gvar96 = 6
	gvar97 = gvar90
	ivar98 = 0
	flag99 = false
	flag100 = false
	flag101 = false
	flag102 = false
	flag103 = false
	flag104 = false
	flag105 = true
	flag106 = false
	flag107 = false
	ivar129 = 60
	flag130 = true
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
	if v0 == gvar30 {
		goto LABEL3
	}
	goto LABEL4
LABEL1:
	ns.Frozen(ns.GetHost(), false)
	flag46 = false
	goto LABEL4
LABEL2:
	ns.Frozen(ns.GetHost(), false)
	flag46 = false
	goto LABEL4
LABEL3:
	ns.CancelDialog(obj75)
	ns.Frozen(ns.GetHost(), false)
	goto LABEL4
LABEL4:
	return
}
func ResetJacobDialog() {
	if !ns.IsCaller(obj75) {
		goto LABEL1
	}
	if gvar40 != gvar30 {
		goto LABEL1
	}
	gvar40 = gvar27
	flag46 = false
	ns.SetDialog(obj75, ns.NORMAL, JacobDialogStart, JacobDialogEnd)
	ivar45 = 0
	ns.ObjectOff(obj77)
LABEL1:
	return
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
	if !(ns.CurrentHealth(obj75) > 0) {
		goto LABEL5
	}
	ns.StartDialog(obj75, ns.GetHost())
LABEL5:
	goto LABEL1
LABEL3:
	gvar40 = gvar29
	ns.Frozen(ns.GetHost(), true)
	ns.CreatureIdle(ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj75)
	ns.LookAtObject(obj75, ns.GetHost())
	if !(ns.CurrentHealth(obj75) > 0) {
		goto LABEL6
	}
	ns.StartDialog(obj75, ns.GetHost())
LABEL6:
	goto LABEL1
LABEL4:
	gvar40 = gvar30
	ns.CancelDialog(obj75)
	ns.ClearOwner(obj75)
	ns.AggressionLevel(obj75, 0.83)
	if !(ns.CurrentHealth(obj75) > 0) {
		goto LABEL7
	}
	ns.Attack(obj75, ns.GetHost())
LABEL7:
	goto LABEL1
LABEL1:
	return
}
func PlayerFlee() {
	if gvar40 != gvar30 {
		goto LABEL1
	}
	ns.ObjectOn(obj77)
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
	if !(ivar44 >= 2) {
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
func ResetHenrickDialogChoice() {
	ns.SetDialog(obj48, ns.YESNO, HenrickDialogStart, HenrickDialogEnd)
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
	ns.TellStory(ns.SwordsmanHurt, "Wiz08a:Guard01Greet")
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
	ns.TellStory(ns.SwordsmanHurt, "War03b:MayorsGuardIntro")
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
	ns.TellStory(ns.SwordsmanHurt, "War08a:AldwynGreet")
	goto LABEL4
LABEL2:
	ns.LookAtObject(obj54, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War08a:AldwynProd")
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
	gvar33 = gvar9
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
	ns.TellStory(ns.HumanMaleEatFood, "War03b:Maiden1")
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
	ns.WallOpen(ns.Wall(128, 106))
	ns.EnchantOff(obj120[a1], ns.ENCHANT_INVULNERABLE)
	ns.CreatureFollow(obj120[a1], ns.GetHost())
	ns.PauseObject(obj120[a1], 30)
	ns.AggressionLevel(obj120[a1], 0.83)
	ns.MakeFriendly(obj120[a1])
	ns.BecomePet(obj120[a1])
	goto LABEL5
LABEL2:
	ns.WallOpen(ns.Wall(130, 104))
	ns.EnchantOff(obj120[a1], ns.ENCHANT_INVULNERABLE)
	ns.CreatureFollow(obj120[a1], ns.GetHost())
	ns.PauseObject(obj120[a1], 30)
	ns.AggressionLevel(obj120[a1], 0.83)
	ns.MakeFriendly(obj120[a1])
	ns.BecomePet(obj120[a1])
	goto LABEL5
LABEL3:
	ns.WallOpen(ns.Wall(132, 102))
	ns.EnchantOff(obj120[a1], ns.ENCHANT_INVULNERABLE)
	ns.CreatureFollow(obj120[a1], ns.GetHost())
	ns.PauseObject(obj120[a1], 30)
	ns.AggressionLevel(obj120[a1], 0.83)
	ns.MakeFriendly(obj120[a1])
	ns.BecomePet(obj120[a1])
	goto LABEL5
LABEL4:
	ns.WallOpen(ns.Wall(134, 100))
	ns.EnchantOff(obj120[a1], ns.ENCHANT_INVULNERABLE)
	ns.CreatureFollow(obj120[a1], ns.GetHost())
	ns.PauseObject(obj120[a1], 30)
	ns.AggressionLevel(obj120[a1], 0.83)
	ns.MakeFriendly(obj120[a1])
	ns.BecomePet(obj120[a1])
	goto LABEL5
LABEL5:
	return
}
func InitializeDialogs() {
	ns.StoryPic(obj48, "HenrickPic")
	ns.StoryPic(obj49, "DrunkPic")
	ns.StoryPic(obj51, "Townsman1Pic")
	ns.StoryPic(obj52, "IxGuard2Pic")
	ns.StoryPic(obj53, "Warrior3Pic")
	ns.StoryPic(obj55, "TheogrinPic")
	ns.StoryPic(obj54, "AldwynPic")
	ns.StoryPic(obj63, "IxGuard2Pic")
	ns.StoryPic(obj64, "IxGuard1Pic")
	ns.StoryPic(obj65, "Townsman3Pic")
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
	ns.SetDialog(obj48, ns.YESNO, HenrickDialogStart, HenrickDialogEnd)
	ns.SetDialog(obj49, ns.NORMAL, DrunkDialogStart, DrunkDialogEnd)
	ns.SetDialog(obj63, ns.NORMAL, Gatekeeper2DialogStart, Gatekeeper2DialogEnd)
	ns.SetDialog(obj64, ns.NORMAL, Gatekeeper3DialogStart, Gatekeeper3DialogEnd)
	ns.SetDialog(obj53, ns.NORMAL, MayorsGuardDialogStart, MayorsGuardDialogEnd)
	ns.SetDialog(obj54, ns.NORMAL, AldwinDialogStart, AldwinDialogEnd)
	ns.SetDialog(obj65, ns.NORMAL, BridgeGuardDialogStart, BridgeGuardDialogEnd)
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
func PlayerDeath() {
	ns.DeathScreen(8)
}
func GoToMineSEG2() {
	ns.MoveObject(ns.GetHost(), ns.GetWaypointX(wp80), ns.GetWaypointY(wp80))
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
	ns.MoveObject(ns.GetHost(), ns.GetWaypointX(wp81), ns.GetWaypointY(wp81))
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
	obj82 = ns.Object("Teacher")
	obj83 = ns.Object("Apprentice")
	obj84 = ns.Object("InvertSpell")
	obj85 = ns.Object("LessonGate1")
	obj86 = ns.Object("LessonGate2")
	obj87 = ns.Object("ApprenticeTrigger")
	wp88 = ns.Waypoint("PlayerSounds")
	wp89 = ns.Waypoint("ApprenticeWP")
	ns.StoryPic(obj82, "WizardGuard1Pic")
	ns.SetDialog(obj82, ns.YESNO, TeacherStart, TeacherEnd)
	ns.LockDoor(obj85)
	ns.LockDoor(obj86)
	ns.ObjectOff(obj87)
	ns.FrameTimer(1, InversionOwnObject)
}
func InversionOwnObject() {
	ns.SetOwner(ns.GetHost(), obj82)
	ns.Enchant(obj83, ns.ENCHANT_INVULNERABLE, 0)
	ns.Enchant(obj82, ns.ENCHANT_ANCHORED, 0)
	ns.Enchant(obj83, ns.ENCHANT_ANCHORED, 0)
}
func TeacherStart() {
	var v0 int
	v0 = gvar97
	if v0 == gvar90 {
		goto LABEL1
	}
	if v0 == gvar91 {
		goto LABEL2
	}
	if v0 == gvar92 {
		goto LABEL3
	}
	if v0 == gvar93 {
		goto LABEL4
	}
	if v0 == gvar94 {
		goto LABEL5
	}
	if v0 == gvar95 {
		goto LABEL6
	}
	if v0 == gvar96 {
		goto LABEL7
	}
	goto LABEL8
LABEL1:
	ns.Frozen(ns.GetHost(), true)
	ns.CreatureIdle(ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj82)
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
	v6 = gvar97
	if v6 == gvar90 {
		goto LABEL1
	}
	if v6 == gvar91 {
		goto LABEL2
	}
	if v6 == gvar92 {
		goto LABEL3
	}
	if v6 == gvar93 {
		goto LABEL4
	}
	if v6 == gvar94 {
		goto LABEL5
	}
	if v6 == gvar95 {
		goto LABEL6
	}
	if v6 == gvar96 {
		goto LABEL7
	}
	goto LABEL8
LABEL1:
	v1 = ns.GetAnswer(obj82)
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
	ns.SetDialog(obj82, ns.NORMAL, TeacherStart, TeacherEnd)
	gvar97 = gvar91
	ns.StartDialog(obj82, ns.GetHost())
	goto LABEL12
	goto LABEL10
LABEL13:
	ns.ChangeGold(ns.GetHost(), -100)
	ns.UnlockDoor(obj85)
	ns.UnlockDoor(obj86)
	ns.SetDialog(obj82, ns.NORMAL, TeacherStart, TeacherEnd)
	gvar97 = gvar92
	ns.StartDialog(obj82, ns.GetHost())
	goto LABEL12
LABEL10:
	ns.Frozen(ns.GetHost(), false)
	gvar97 = gvar90
	ns.SetDialog(obj82, ns.YESNO, TeacherStart, TeacherEnd)
	goto LABEL12
LABEL11:
	ns.Frozen(ns.GetHost(), false)
	gvar97 = gvar90
	ns.SetDialog(obj82, ns.YESNO, TeacherStart, TeacherEnd)
	goto LABEL12
LABEL12:
	goto LABEL8
LABEL2:
	ns.Frozen(ns.GetHost(), false)
	gvar97 = gvar90
	ns.SetDialog(obj82, ns.YESNO, TeacherStart, TeacherEnd)
	goto LABEL8
LABEL3:
	ns.CancelDialog(obj82)
	ns.Frozen(ns.GetHost(), false)
	ns.Pickup(ns.GetHost(), obj84)
	goto LABEL8
LABEL4:
	ns.CancelDialog(obj82)
	flag103 = false
	goto LABEL8
LABEL5:
	ns.UnlockDoor(obj58)
	ns.UnlockDoor(obj59)
	gvar97 = gvar96
	ns.SetDialog(obj82, ns.NORMAL, TeacherStart, TeacherEnd)
	goto LABEL8
LABEL6:
	ns.UnlockDoor(obj58)
	ns.UnlockDoor(obj59)
	gvar97 = gvar96
	ns.SetDialog(obj82, ns.NORMAL, TeacherStart, TeacherEnd)
	goto LABEL8
LABEL7:
	gvar97 = gvar96
	ns.SetDialog(obj82, ns.NORMAL, TeacherStart, TeacherEnd)
	goto LABEL8
LABEL8:
	return
}
func Phoneme1() {
	if flag103 {
		goto LABEL1
	}
	if !(ns.CurrentHealth(obj83) > 0) {
		goto LABEL1
	}
	ns.PrintToAll("Phoneme1")
	ns.MoveWaypoint(wp88, 1326, 2493)
	ns.AudioEvent(ns.SpellPhonemeDown, wp88)
	ns.FrameTimer(10, Phoneme2)
LABEL1:
	return
}
func Phoneme2() {
	if flag103 {
		goto LABEL1
	}
	if !(ns.CurrentHealth(obj83) > 0) {
		goto LABEL1
	}
	ns.PrintToAll("Phoneme2")
	ns.AudioEvent(ns.SpellPhonemeDown, wp88)
	ns.FrameTimer(10, Phoneme3)
LABEL1:
	return
}
func Phoneme3() {
	if flag103 {
		goto LABEL1
	}
	if !(ns.CurrentHealth(obj83) > 0) {
		goto LABEL1
	}
	ns.PrintToAll("Phoneme3")
	ns.AudioEvent(ns.SpellPhonemeDown, wp88)
	ns.FrameTimer(10, CastSpells)
LABEL1:
	return
}
func LessonBegin() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	if gvar97 != gvar92 {
		goto LABEL1
	}
	if !flag105 {
		goto LABEL2
	}
	ns.Move(obj83, wp89)
	flag105 = false
LABEL2:
	if !(flag99 == false && flag104 == true) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.LookAtObject(obj83, ns.GetHost())
	flag99 = true
	ns.LockDoor(obj58)
	ns.LockDoor(obj59)
	ns.SecondTimer(1, Phoneme1)
LABEL1:
	return
}
func CastSpells() {
	if flag103 {
		goto LABEL1
	}
	if !(ivar98 < 4) {
		goto LABEL2
	}
	ns.PrintToAll("Casting Spell")
	ns.CastSpellObjectLocation(ns.SPELL_SLOW, obj83, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.Enchant(obj83, ns.ENCHANT_INVULNERABLE, 0)
LABEL2:
	if ivar98 != 4 {
		goto LABEL3
	}
	ns.PrintToAll("CastingSpell")
	ns.EnchantOff(obj83, ns.ENCHANT_INVULNERABLE)
	ns.CastSpellObjectLocation(ns.SPELL_MAGIC_MISSILE, obj83, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
LABEL3:
	ivar98 += 1
	if !(ivar98 < 5) {
		goto LABEL4
	}
	ns.SecondTimer(3, Phoneme1)
	goto LABEL1
LABEL4:
	flag100 = true
	if !flag102 {
		goto LABEL5
	}
	gvar97 = gvar94
	ns.SetDialog(obj82, ns.NORMAL, TeacherStart, TeacherEnd)
	ns.StartDialog(obj82, ns.GetHost())
	goto LABEL1
LABEL5:
	gvar97 = gvar95
	ns.SetDialog(obj82, ns.NORMAL, TeacherStart, TeacherEnd)
	ns.StartDialog(obj82, ns.GetHost())
LABEL1:
	return
}
func CheckApprentice() {
	if !ns.HasEnchant(obj83, ns.ENCHANT_SLOWED) {
		goto LABEL1
	}
	ns.ObjectOff(obj87)
	flag103 = true
	gvar97 = gvar93
	flag102 = true
	if !(ivar98 > 2) {
		goto LABEL2
	}
	ivar98 = 2
LABEL2:
	ns.SetDialog(obj82, ns.NORMAL, TeacherStart, TeacherEnd)
	ns.StartDialog(obj82, ns.GetHost())
LABEL1:
	if !(ns.CurrentHealth(obj83) <= 0) {
		goto LABEL3
	}
	ns.ObjectOff(ns.GetTrigger())
	flag102 = true
LABEL3:
	return
}
func ApprenticeDie() {
	ns.AudioEvent(ns.HumanMaleHurtHeavy, wp88)
}
func ApprenticeReport() {
	ns.CreatureGuard(obj83, ns.GetWaypointX(wp89), ns.GetWaypointY(wp89), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	ns.ObjectOn(obj87)
	flag104 = true
}
func OwnObjects() {
	ns.GroupSetOwner(ns.GetHost(), gvar126)
	ns.SetOwner(ns.GetHost(), obj48)
	ns.SetOwner(ns.GetHost(), obj108)
	ns.GroupSetOwner(ns.GetHost(), gvar127)
	ns.SetOwner(ns.GetHost(), obj49)
	ns.SetOwner(ns.GetHost(), obj67)
	ns.SetOwner(ns.GetHost(), obj68)
	ns.SetOwner(ns.GetHost(), obj69)
	ns.SetOwner(ns.GetHost(), obj70)
	ns.SetOwner(ns.GetHost(), obj47)
	ns.SetOwner(ns.GetHost(), obj75)
	ns.StoryPic(obj51, "ShopKeeperBrownPic")
}
func DisableElevators() {
	ns.ObjectOff(obj60)
	ns.ObjectOff(obj66)
}
func MapInitialize() {
	var v0 int
	obj48 = ns.Object("Henrick")
	obj108 = ns.Object("HenricksWolf")
	obj120[0] = ns.Object("CharmedWolf01")
	obj120[1] = ns.Object("CharmedWolf02")
	obj120[2] = ns.Object("CharmedWolf03")
	obj120[3] = ns.Object("CharmedWolf04")
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL1
		}
		ns.Enchant(obj120[v0], ns.ENCHANT_INVULNERABLE, 0)
		v0 += 1
	}
LABEL1:
	obj49 = ns.Object("Drunk")
	obj54 = ns.Object("Aldwyn")
	obj52 = ns.Object("Geoff")
	obj63 = ns.Object("Ryan")
	obj64 = ns.Object("Ed")
	obj55 = ns.Object("Mayor_Theogrin")
	obj53 = ns.Object("Mayor's_Guard")
	obj51 = ns.Object("Barkeeper")
	obj65 = ns.Object("Bridge_Guard")
	obj47 = ns.Object("Jandor")
	obj75 = ns.Object("Jacob")
	obj76 = ns.Object("Morgan")
	obj119 = ns.Object("CaveDoor")
	obj116 = ns.Object("MorganDoor")
	obj56 = ns.Object("MayorGate1")
	obj57 = ns.Object("MayorGate2")
	obj58 = ns.Object("ContestGate1")
	obj59 = ns.Object("ContestGate2")
	obj112 = ns.Object("CemDoorL")
	obj113 = ns.Object("CemDoorR")
	obj110 = ns.Object("BasementDoor")
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
	obj77 = ns.Object("JacobTrigger")
	obj111 = ns.Object("AldwynsDoor")
	obj114 = ns.Object("TempDoorL")
	obj115 = ns.Object("TempDoorR")
	obj117 = ns.Object("Mystic")
	obj118 = ns.Object("Mystic2")
	wp78 = ns.Waypoint("AldwynHome")
	wp122 = ns.Waypoint("AldwynStorage")
	wp79 = ns.Waypoint("JacobWP")
	wp123 = ns.Waypoint("MysticWP")
	wp124 = ns.Waypoint("MysticStorage")
	wp125 = ns.Waypoint("FaceWP")
	gvar126 = ns.ObjectGroup("NPC's")
	gvar127 = ns.ObjectGroup("CharmedWolves")
	gvar128 = ns.WallGroup("MagicShopElevatorWalls")
	wp80 = ns.Waypoint("MineExitWP")
	wp81 = ns.Waypoint("CaveExitWP")
	ns.Wander(obj72)
	ns.Wander(obj73)
	ns.Wander(obj74)
	ns.Wander(obj68)
	ns.Wander(obj69)
	ns.LockDoor(obj56)
	ns.LockDoor(obj57)
	ns.LockDoor(obj112)
	ns.LockDoor(obj113)
	ns.LockDoor(obj61)
	ns.LockDoor(obj62)
	ns.LockDoor(obj116)
	ns.LockDoor(obj119)
	InversionInit()
	InitializeDialogs()
	OwnObjects()
	ns.StartupScreen(8)
	ns.FrameTimer(5, StartCaptainConversation)
	ns.FrameTimer(35, DisableElevators)
}
func MapEntry() {
	var v0 int
	v0 = ns.GetQuestStatus("Chapter8:HasWeirdling")
	if v0 != 1 {
		goto LABEL1
	}
	flag107 = true
	ns.MoveObject(obj54, ns.GetWaypointX(wp122), ns.GetWaypointY(wp122))
	ns.LockDoor(obj111)
	ns.LockDoor(obj114)
	ns.LockDoor(obj115)
	InitializeSecondaryDialogs()
	ns.SetShopkeeperText(obj117, "War08a:Mystic")
	ns.CancelDialog(obj63)
	ns.CancelDialog(obj64)
LABEL1:
	ns.Music(15, 100)
	ns.UnBlind()
}
func DebugOn() {
	flag106 = true
	ns.PrintToAll("Debug mode is enabled.")
}
func DebugOff() {
	flag106 = false
	ns.PrintToAll("Debug mode has been disabled.")
}
func OpenCaveDoor() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.UnlockDoor(obj119)
	ns.ObjectOff(ns.GetTrigger())
LABEL1:
	return
}
func EnableRandomBump() {
	flag130 = true
}
func RandomBump() {
	if !(ns.IsCaller(ns.GetHost()) && flag130) {
		goto LABEL1
	}
	ns.Chat(ns.GetTrigger(), "Con02a:RandomBump")
	flag130 = false
	ns.FrameTimer(ivar129, EnableRandomBump)
LABEL1:
	return
}
func RestAwhile() {
	var v0 int
	v0 = ns.Random(60, 120)
	ns.PauseObject(ns.GetCaller(), v0)
}
func OpenMagicShopElevatorWalls() {
	ns.WallGroupOpen(gvar128)
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
