package war03b

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	gvar4  int
	gvar5  int
	gvar6  int
	obj7   ns.ObjectID
	obj8   ns.ObjectID
	ivar9  int
	gvar10 int
	ivar11 int
	gvar12 int
	obj13  [2]ns.ObjectID
	gvar14 ns.ObjectGroupID
	obj15  ns.ObjectID
	obj16  ns.ObjectID
	wp17   ns.WaypointID
	obj18  ns.ObjectID
	obj19  ns.ObjectID
	gvar20 ns.ObjectID
	gvar21 ns.ObjectID
	gvar22 ns.ObjectGroupID
	obj23  ns.ObjectID
	obj24  ns.ObjectID
	obj25  ns.ObjectID
	obj26  ns.ObjectID
	obj27  ns.ObjectID
	obj28  ns.ObjectID
	obj29  ns.ObjectID
	obj30  ns.ObjectID
	obj31  ns.ObjectID
	obj32  ns.ObjectID
	obj33  ns.ObjectID
	obj34  ns.ObjectID
	obj35  ns.ObjectID
	obj36  ns.ObjectID
	obj37  ns.ObjectID
	obj38  ns.ObjectID
	obj39  ns.ObjectID
	obj40  ns.ObjectID
	obj41  ns.ObjectID
	obj42  ns.ObjectID
	gvar43 ns.ObjectGroupID
	gvar44 ns.ObjectGroupID
	gvar45 ns.ObjectGroupID
	gvar46 ns.ObjectGroupID
	obj47  ns.ObjectID
	obj48  ns.ObjectID
	obj49  ns.ObjectID
	obj50  ns.ObjectID
	obj51  ns.ObjectID
	obj52  ns.ObjectID
	obj53  ns.ObjectID
	obj54  ns.ObjectID
	wp55   ns.WaypointID
	wp56   ns.WaypointID
	wp57   ns.WaypointID
	wp58   ns.WaypointID
	gvar59 int
	wp60   ns.WaypointID
	gvar61 int
	gvar62 int
	gvar63 int
	gvar64 int
	gvar65 int
	gvar66 int
	gvar67 int
	gvar68 int
	gvar69 int
	gvar70 int
	gvar71 int
	gvar72 int
	gvar73 int
	gvar74 int
	gvar75 int
	gvar76 int
	gvar77 int
	gvar78 int
	gvar79 int
	gvar80 int
	gvar81 int
	gvar82 int
	gvar83 int
	gvar84 int
	gvar85 int
	gvar86 int
	gvar87 int
	gvar88 int
	gvar89 int
	gvar90 int
	gvar91 int
	gvar92 int
	gvar93 int
	gvar94 int
	gvar95 int
	gvar96 int
	gvar97 int
	obj98  ns.ObjectID
	gvar99 ns.WallGroupID
	wp100  ns.WaypointID
	wp101  ns.WaypointID
	wp102  ns.WaypointID
)

func init() {
	gvar4 = 0
	gvar5 = 1
	gvar6 = 2
	ivar9 = 200
	gvar10 = 2
	ivar11 = 0
	gvar12 = gvar4
	gvar61 = 0
	gvar62 = 1
	gvar63 = 2
	gvar64 = 3
	gvar65 = gvar61
	gvar66 = 0
	gvar67 = 1
	gvar68 = 2
	gvar69 = 3
	gvar70 = gvar66
	gvar71 = 0
	gvar72 = 1
	gvar73 = gvar71
	gvar74 = 0
	gvar75 = 1
	gvar76 = gvar74
	gvar77 = 0
	gvar78 = 1
	gvar79 = gvar77
	gvar80 = 0
	gvar81 = 1
	gvar82 = gvar80
	gvar83 = 0
	gvar84 = 1
	gvar85 = gvar83
	gvar86 = 0
	gvar87 = 1
	gvar88 = gvar86
	gvar89 = 0
	gvar90 = 1
	gvar91 = 2
	gvar92 = 3
	gvar93 = gvar89
	gvar94 = 0
	gvar95 = 1
	gvar96 = 2
	gvar97 = gvar94
}
func PlayerDeath() {
	ns.DeathScreen(3)
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
	ns.CreatureFollow(obj13[a1], ns.GetHost())
	ns.PauseObject(obj13[a1], 30)
	ns.AggressionLevel(obj13[a1], 0.83)
	ns.MakeFriendly(obj13[a1])
	ns.BecomePet(obj13[a1])
	goto LABEL5
LABEL2:
	ns.WallOpen(ns.Wall(130, 104))
	ns.CreatureFollow(obj13[a1], ns.GetHost())
	ns.PauseObject(obj13[a1], 30)
	ns.AggressionLevel(obj13[a1], 0.83)
	ns.MakeFriendly(obj13[a1])
	ns.BecomePet(obj13[a1])
	goto LABEL5
LABEL3:
	ns.WallOpen(ns.Wall(55, 109))
	ns.EnchantOff(obj13[a1], ns.ENCHANT_INVULNERABLE)
	ns.CreatureFollow(obj13[a1], ns.GetHost())
	ns.AggressionLevel(obj13[a1], 0.83)
	ns.MakeFriendly(obj13[a1])
	ns.BecomePet(obj13[a1])
	goto LABEL5
LABEL4:
	ns.WallOpen(ns.Wall(53, 107))
	ns.EnchantOff(obj13[a1], ns.ENCHANT_INVULNERABLE)
	ns.CreatureFollow(obj13[a1], ns.GetHost())
	ns.AggressionLevel(obj13[a1], 0.83)
	ns.MakeFriendly(obj13[a1])
	ns.BecomePet(obj13[a1])
	goto LABEL5
LABEL5:
	return
}
func HenrickTalkStart() {
	var v0 int
	v0 = gvar12
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
	ns.LookAtObject(obj7, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War08b:HenrickSalesPitchA")
	goto LABEL4
LABEL2:
	ns.LookAtObject(obj7, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War08b:HenrickSalesPitchB")
	goto LABEL4
LABEL3:
	ns.LookAtObject(obj7, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War08b:HenrickNoMoreWolves")
	goto LABEL4
LABEL4:
	return
}
func HenrickTalkEnd() {
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
	v7 = gvar12
	if v7 == gvar4 {
		goto LABEL1
	}
	if v7 == gvar5 {
		goto LABEL2
	}
	if v7 == gvar6 {
		goto LABEL3
	}
	goto LABEL4
LABEL1:
	v1 = ns.GetAnswer(obj7)
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
	if !(v0 < ivar9) {
		goto LABEL9
	}
	ns.SetDialog(obj7, ns.NORMAL, NullDialogStart, ResetHenrickDialogChoice)
	ns.TellStory(ns.SwordsmanHurt, "War08b:HenrickNotEnoughGold")
	goto LABEL10
LABEL9:
	ns.ChangeGold(ns.GetHost(), -ivar9)
	gvar12 = gvar5
	ReleaseCharmedWolf(ivar11)
	ivar11 += 1
	ns.SetDialog(obj7, ns.NORMAL, NullDialogStart, ResetHenrickDialogChoice)
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
	v1 = ns.GetAnswer(obj7)
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
	if !(v0 < ivar9) {
		goto LABEL15
	}
	ns.SetDialog(obj7, ns.NORMAL, NullDialogStart, ResetHenrickDialogChoice)
	ns.TellStory(ns.SwordsmanHurt, "War08b:HenrickNotEnoughGold")
	goto LABEL16
LABEL15:
	ns.ChangeGold(ns.GetHost(), -ivar9)
	ReleaseCharmedWolf(ivar11)
	ivar11 += 1
	ns.SetDialog(obj7, ns.NORMAL, NullDialogStart, ResetHenrickDialogChoice)
	ns.TellStory(ns.SwordsmanHurt, "War08b:HenrickSaleSuccessful")
	if !(ivar11 >= 2) {
		goto LABEL16
	}
	gvar12 = gvar6
	ns.SetDialog(obj7, ns.NORMAL, HenrickTalkStart, HenrickTalkEnd)
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
func NullDialogStart() {
}
func InitializeHenrick() {
	ns.StoryPic(obj7, "HenrickPic")
	ns.SetDialog(obj7, ns.YESNO, HenrickTalkStart, HenrickTalkEnd)
}
func ResetHenrickDialogChoice() {
	ns.SetDialog(obj7, ns.YESNO, HenrickTalkStart, HenrickTalkEnd)
}
func AirshipCaptainAppear() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
	)
	_ = v3
	_ = v2
	v0 = ns.GetWaypointX(wp55)
	v1 = ns.GetWaypointY(wp55)
	v2 = ns.GetWaypointX(wp57)
	v3 = ns.GetWaypointY(wp57)
	ns.MoveObject(obj54, v0, v1)
	ns.FrameTimer(1, CaptainGuard)
}
func MayorPickup() {
	ns.MoveObject(ns.Object("War03c:MayorScepter"), 4963, 453)
	ns.Pickup(ns.GetHost(), obj19)
}
func MayorSpeaks() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
	)
	v0 = ns.GetObjectX(obj15)
	v1 = ns.GetObjectY(obj15)
	v2 = ns.GetObjectX(ns.GetHost())
	v3 = ns.GetObjectY(ns.GetHost())
	if !(ns.Distance(v0, v1, v2, v3) < 150) {
		goto LABEL1
	}
	ns.LookAtObject(obj15, ns.GetHost())
	ns.StartDialog(obj15, ns.GetHost())
	goto LABEL2
LABEL1:
	ns.FrameTimer(22, MayorSpeaks)
LABEL2:
	return
}
func MayorGuardTime() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
	)
	v0 = ns.GetWaypointX(wp58)
	v1 = ns.GetWaypointY(wp58)
	v2 = ns.GetWaypointX(wp17)
	v3 = ns.GetWaypointY(wp17)
	ns.CreatureGuard(obj15, v0, v1, v2, v3, 0)
	ns.StoryPic(obj15, "TheogrinPic")
	ns.SetDialog(obj15, ns.NORMAL, MayorTalkStart, MayorTalkEnd)
}
func MayorTalkStart() {
	var v0 int
	if !ns.HasItem(ns.GetHost(), ns.Object("War03c:MayorScepter")) {
		goto LABEL1
	}
	gvar65 = gvar63
LABEL1:
	v0 = gvar65
	if v0 == gvar61 {
		goto LABEL2
	}
	if v0 == gvar62 {
		goto LABEL3
	}
	if v0 == gvar63 {
		goto LABEL4
	}
	if v0 == gvar64 {
		goto LABEL5
	}
	goto LABEL6
LABEL2:
	ns.LookAtObject(obj15, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War03b:MayorIntro")
	goto LABEL6
LABEL3:
	ns.CreatureIdle(obj15)
	ns.LookAtObject(obj15, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War03b:MayorPre")
	goto LABEL6
LABEL4:
	ns.LookAtObject(obj15, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War03b:MayorPost")
	goto LABEL6
LABEL5:
	ns.LookAtObject(obj15, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War03b:MayorPost2")
	goto LABEL6
LABEL6:
	return
}
func MayorsGuardTalkStart() {
	var v0 int
	ns.LookAtObject(obj16, ns.GetHost())
	v0 = gvar93
	if v0 == gvar89 {
		goto LABEL1
	}
	if v0 == gvar90 {
		goto LABEL2
	}
	if v0 == gvar91 {
		goto LABEL3
	}
	if v0 == gvar92 {
		goto LABEL4
	}
	goto LABEL5
LABEL1:
	ns.TellStory(ns.SwordsmanHurt, "War03b:MayorsGuardIntro")
	goto LABEL5
LABEL2:
	ns.TellStory(ns.SwordsmanHurt, "War03b:MayorsGuardPre")
	goto LABEL5
LABEL3:
	ns.TellStory(ns.SwordsmanHurt, "War03b:MayorsGuardPost")
	goto LABEL5
LABEL4:
	ns.TellStory(ns.SwordsmanHurt, "War03b:MayorsGuardPost2")
	goto LABEL5
LABEL5:
	return
}
func GatekeeperTalkStart() {
	var v0 int
	ns.LookAtObject(obj23, ns.GetHost())
	v0 = gvar88
	if v0 == gvar86 {
		goto LABEL1
	}
	if v0 == gvar87 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.TellStory(ns.SwordsmanHurt, "War03b:GatekeeperPre")
	goto LABEL3
LABEL2:
	ns.TellStory(ns.SwordsmanHurt, "War03b:GatekeeperPost")
	goto LABEL3
LABEL3:
	return
}
func BarkeeperTalkStart() {
	var v0 int
	ns.LookAtObject(obj24, ns.GetHost())
	v0 = gvar85
	if v0 == gvar83 {
		goto LABEL1
	}
	if v0 == gvar84 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.TellStory(ns.SwordsmanHurt, "War03b:BarkeeperPre")
	goto LABEL3
LABEL2:
	ns.TellStory(ns.SwordsmanHurt, "War03b:BarkeeperPost")
	goto LABEL3
LABEL3:
	return
}
func ContestGuardTalkStart() {
	ns.LookAtObject(obj25, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War03b:ContestGuard")
}
func BridgeGuardTalkStart() {
	var v0 int
	ns.LookAtObject(obj36, ns.GetHost())
	v0 = gvar97
	if v0 == gvar94 {
		goto LABEL1
	}
	if v0 == gvar95 {
		goto LABEL2
	}
	if v0 == gvar96 {
		goto LABEL3
	}
	goto LABEL4
LABEL1:
	ns.TellStory(ns.SwordsmanHurt, "War03b:BridgeGuardIntro")
	goto LABEL4
LABEL2:
	ns.TellStory(ns.SwordsmanHurt, "War03b:BridgeGuardPre")
	goto LABEL4
LABEL3:
	ns.TellStory(ns.SwordsmanHurt, "War03b:BridgeGuardPost")
	goto LABEL4
LABEL4:
	return
}
func Townsman1TalkStart() {
	var v0 int
	ns.CreatureIdle(obj26)
	ns.LookAtObject(obj26, ns.GetHost())
	v0 = gvar73
	if v0 == gvar71 {
		goto LABEL1
	}
	if v0 == gvar72 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.TellStory(ns.SwordsmanHurt, "War03b:T1Pre")
	goto LABEL3
LABEL2:
	ns.TellStory(ns.SwordsmanHurt, "War03b:T1Post")
	goto LABEL3
LABEL3:
	return
}
func Townsman2TalkStart() {
	var v0 int
	ns.CreatureIdle(obj27)
	ns.LookAtObject(obj27, ns.GetHost())
	v0 = gvar76
	if v0 == gvar74 {
		goto LABEL1
	}
	if v0 == gvar75 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.TellStory(ns.SwordsmanHurt, "War03b:T2Pre")
	goto LABEL3
LABEL2:
	ns.TellStory(ns.SwordsmanHurt, "War03b:T2Post")
	goto LABEL3
LABEL3:
	return
}
func Townsman3TalkStart() {
	var v0 int
	ns.LookAtObject(obj28, ns.GetHost())
	v0 = gvar79
	if v0 == gvar77 {
		goto LABEL1
	}
	if v0 == gvar78 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.Chat(obj28, "War02a:NewTownsman1")
	goto LABEL3
LABEL2:
	ns.Chat(obj28, "War02a:NewTownsman2")
	goto LABEL3
LABEL3:
	return
}
func Townsman4TalkStart() {
	var v0 int
	ns.LookAtObject(obj29, ns.GetHost())
	v0 = gvar82
	if v0 == gvar80 {
		goto LABEL1
	}
	if v0 == gvar81 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.TellStory(ns.SwordsmanHurt, "War03b:T4Pre")
	goto LABEL3
LABEL2:
	ns.TellStory(ns.SwordsmanHurt, "War03b:T4Post")
	goto LABEL3
LABEL3:
	return
}
func AldwynTalkStart() {
	var v0 int
	ns.LookAtObject(obj33, ns.GetHost())
	if !ns.HasItem(ns.GetHost(), ns.Object("War03c:MayorScepter")) {
		goto LABEL1
	}
	gvar70 = gvar68
LABEL1:
	v0 = gvar70
	if v0 == gvar66 {
		goto LABEL2
	}
	if v0 == gvar67 {
		goto LABEL3
	}
	if v0 == gvar68 {
		goto LABEL4
	}
	if v0 == gvar69 {
		goto LABEL5
	}
	goto LABEL6
LABEL2:
	ns.TellStory(ns.SwordsmanHurt, "War03b:AldwynIntro")
	goto LABEL6
LABEL3:
	ns.TellStory(ns.SwordsmanHurt, "War03b:AldwynPre")
	goto LABEL6
LABEL4:
	ns.TellStory(ns.SwordsmanHurt, "War03b:AldwynPost")
	goto LABEL6
LABEL5:
	ns.TellStory(ns.SwordsmanHurt, "War03b:AldwynPost2")
	goto LABEL6
LABEL6:
	return
}
func AirshipCaptainTalkStart() {
	ns.TellStory(ns.SwordsmanHurt, "War03b:AirshipCaptainIxSpeech")
}
func PlayerExit() {
	var (
		v0 float32
		v1 float32
	)
	v0 = ns.GetWaypointX(wp56)
	v1 = ns.GetWaypointY(wp56)
	ns.MoveObject(ns.GetHost(), v0, v1)
	ns.Frozen(ns.GetHost(), false)
}
func Maiden1TalkStart() {
	ns.CreatureIdle(obj30)
	ns.LookAtObject(obj30, ns.GetHost())
	ns.LookAtObject(obj30, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War03b:Maiden1")
}
func Maiden2TalkStart() {
	ns.CreatureIdle(obj31)
	ns.LookAtObject(obj31, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War03b:Maiden2")
}
func MayorTalkEnd() {
	var v0 int
	v0 = gvar65
	if v0 == gvar61 {
		goto LABEL1
	}
	if v0 == gvar62 {
		goto LABEL2
	}
	if v0 == gvar63 {
		goto LABEL3
	}
	goto LABEL4
LABEL1:
	ns.Move(obj15, wp58)
	ns.Print("Con01a:NewJournalEntry")
	ns.JournalEntry(ns.GetHost(), "War03bGetScepter", 2)
	ns.JournalEdit(ns.GetHost(), "War03aIxQuest", 4)
	gvar65 = gvar62
	gvar93 = gvar90
	gvar97 = gvar95
	gvar88 = gvar86
	ns.CancelDialog(obj15)
	goto LABEL4
LABEL2:
	ns.CreatureGuard(obj15, ns.GetWaypointX(wp58), ns.GetWaypointY(wp58), ns.GetWaypointX(wp17), ns.GetWaypointY(wp17), 0)
	goto LABEL4
LABEL3:
	gvar65 = gvar64
	gvar93 = gvar92
	gvar70 = gvar69
	gvar20 = ns.GetLastItem(ns.GetHost())
	for {
		if gvar20 == 0 {
			goto LABEL5
		}
		gvar21 = ns.GetPreviousItem(gvar20)
		if gvar20 != ns.Object("War03c:MayorScepter") {
			goto LABEL6
		}
		ns.Delete(gvar20)
	LABEL6:
		gvar20 = gvar21
	}
LABEL5:
	ns.Drop(obj15, obj19)
	ns.FrameTimer(1, MayorPickup)
	ns.JournalEdit(ns.GetHost(), "War03bGetScepter", 4)
	ns.GiveXp(ns.GetHost(), 1000)
	AirshipCaptainAppear()
	goto LABEL4
LABEL4:
	return
}
func MayorsGuardTalkEnd() {
}
func GatekeeperTalkEnd() {
}
func BarkeeperTalkEnd() {
}
func ContestGuardTalkEnd() {
}
func BridgeGuardTalkEnd() {
}
func Townsman1TalkEnd() {
	ns.Wander(obj26)
}
func Townsman2TalkEnd() {
	ns.Wander(obj27)
}
func Townsman3TalkEnd() {
	ns.Wander(obj28)
}
func Townsman4TalkEnd() {
}
func AldwynTalkEnd() {
	var v0 int
	v0 = gvar70
	if v0 == gvar66 {
		goto LABEL1
	}
	goto LABEL2
LABEL1:
	ns.Drop(obj33, obj34)
	ns.Drop(obj33, obj35)
	ns.FrameTimer(1, FieldGuidePickup)
	gvar70 = gvar67
	goto LABEL2
LABEL2:
	return
}
func AirshipCaptainTalkEnd() {
	ns.Frozen(ns.GetHost(), true)
	ns.Blind()
	ns.FrameTimer(60, PlayerExit)
	ns.CancelDialog(obj54)
}
func PlayTownMusic() {
	ns.Music(15, 100)
}
func Maiden1TalkEnd() {
	ns.Wander(obj30)
}
func Maiden2TalkEnd() {
}
func JacobDialogStart() {
	ns.LookAtObject(obj41, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con02A:JailerTalk01")
}
func JacobDialogEnd() {
}
func MapInitialize() {
	obj42 = ns.Object("CaveDoor")
	ns.LockDoor(obj42)
	obj41 = ns.Object("Jacob")
	wp60 = ns.Waypoint("WellWP")
	gvar22 = ns.ObjectGroup("QuestTrigGroup")
	obj15 = ns.Object("MayorTheogrin")
	wp17 = ns.Waypoint("MayorLook")
	obj18 = ns.Object("MayorScepter")
	obj19 = ns.Object("MayorGift")
	ns.Pickup(obj15, obj19)
	wp58 = ns.Waypoint("MayorHome")
	ns.Pickup(obj15, obj19)
	gvar14 = ns.ObjectGroup("MayorTrigs")
	obj37 = ns.Object("IxGuard1")
	obj38 = ns.Object("IxGuard2")
	obj23 = ns.Object("Geoff")
	obj16 = ns.Object("Mayor's_Guard")
	obj24 = ns.Object("Barkeeper")
	obj39 = ns.Object("BelforByzanti")
	obj40 = ns.Object("Mystic")
	obj36 = ns.Object("BridgeGuard")
	obj25 = ns.Object("ContestGuard")
	obj26 = ns.Object("Seth")
	obj27 = ns.Object("Rayner")
	obj28 = ns.Object("Townsman3")
	obj29 = ns.Object("Floyd")
	obj7 = ns.Object("Henrick")
	obj8 = ns.Object("Wolf1")
	obj32 = ns.Object("Wolf2")
	obj30 = ns.Object("F6Townswoman1")
	obj31 = ns.Object("F6Townswoman4")
	obj13[0] = ns.Object("Wolf1")
	obj13[1] = ns.Object("Wolf2")
	ns.Wander(obj26)
	ns.Wander(obj27)
	ns.Wander(obj28)
	gvar43 = ns.ObjectGroupID(ns.Object("WaitTrig1")) // FIXME
	gvar44 = ns.ObjectGroupID(ns.Object("WaitTrig2")) // FIXME
	gvar45 = ns.ObjectGroupID(ns.Object("WaitTrig3")) // FIXME
	gvar46 = ns.ObjectGroup("NPCs")
	obj47 = ns.Object("IxGate01")
	obj48 = ns.Object("IxGate02")
	obj49 = ns.Object("IxGate03")
	obj50 = ns.Object("IxGate04")
	obj51 = ns.Object("TempDoorL")
	obj52 = ns.Object("TempDoorR")
	obj53 = ns.Object("AldwynDoor")
	ns.LockDoor(obj49)
	ns.LockDoor(obj50)
	ns.LockDoor(obj51)
	ns.LockDoor(obj52)
	ns.LockDoor(obj53)
	obj54 = ns.Object("AirshipCaptain")
	wp55 = ns.Waypoint("AirshipCaptainWP")
	wp57 = ns.Waypoint("AirshipCaptainLook")
	wp56 = ns.Waypoint("ExitWP")
	ns.GroupSetOwner(ns.GetHost(), gvar46)
	ns.SetOwner(ns.GetHost(), obj15)
	ns.SetOwner(ns.GetHost(), obj16)
	ns.SetOwner(ns.GetHost(), obj36)
	ns.SetOwner(ns.GetHost(), obj25)
	ns.SetOwner(ns.GetHost(), obj24)
	ns.SetOwner(ns.GetHost(), obj23)
	ns.SetOwner(ns.GetHost(), obj33)
	ns.SetOwner(ns.GetHost(), obj26)
	ns.SetOwner(ns.GetHost(), obj27)
	ns.SetOwner(ns.GetHost(), obj28)
	ns.SetOwner(ns.GetHost(), obj29)
	ns.SetOwner(ns.GetHost(), obj54)
	ns.SetOwner(ns.GetHost(), obj37)
	ns.SetOwner(ns.GetHost(), obj38)
	ns.SetOwner(ns.GetHost(), obj39)
	ns.SetOwner(ns.GetHost(), obj40)
	ns.SetOwner(ns.GetHost(), obj8)
	ns.SetOwner(ns.GetHost(), obj32)
	ns.SetOwner(ns.GetHost(), obj7)
	ns.SetOwner(ns.GetHost(), obj30)
	ns.SetOwner(ns.GetHost(), obj31)
	ns.SetOwner(ns.GetHost(), obj41)
	ns.StoryPic(obj23, "MalePic10")
	ns.SetDialog(obj23, ns.NORMAL, GatekeeperTalkStart, GatekeeperTalkEnd)
	ns.StoryPic(obj15, "TheogrinPic")
	ns.SetDialog(obj15, ns.NORMAL, MayorTalkStart, MayorTalkEnd)
	ns.StoryPic(obj16, "Warrior3Pic")
	ns.SetDialog(obj16, ns.NORMAL, MayorsGuardTalkStart, MayorsGuardTalkEnd)
	ns.StoryPic(obj25, "MalePic4")
	ns.SetDialog(obj25, ns.NORMAL, ContestGuardTalkStart, ContestGuardTalkEnd)
	ns.StoryPic(obj36, "MalePic9")
	ns.SetDialog(obj36, ns.NORMAL, BridgeGuardTalkStart, BridgeGuardTalkEnd)
	ns.StoryPic(obj26, "Townsman1Pic")
	ns.SetDialog(obj26, ns.NORMAL, Townsman1TalkStart, Townsman1TalkEnd)
	ns.StoryPic(obj27, "Townsman2Pic")
	ns.SetDialog(obj27, ns.NORMAL, Townsman2TalkStart, Townsman2TalkEnd)
	ns.StoryPic(obj28, "Townsman3Pic")
	ns.SetDialog(obj28, ns.NORMAL, Townsman3TalkStart, Townsman3TalkEnd)
	ns.StoryPic(obj29, "Townsman3Pic")
	ns.SetDialog(obj29, ns.NORMAL, Townsman4TalkStart, Townsman4TalkEnd)
	ns.StoryPic(obj33, "AldwynPic")
	ns.SetDialog(obj33, ns.NORMAL, AldwynTalkStart, AldwynTalkEnd)
	ns.StoryPic(obj54, "AirshipCaptainPic")
	ns.SetDialog(obj54, ns.NORMAL, AirshipCaptainTalkStart, AirshipCaptainTalkEnd)
	ns.StoryPic(obj30, "MaidenPic2")
	ns.SetDialog(obj30, ns.NORMAL, Maiden1TalkStart, Maiden1TalkEnd)
	ns.StoryPic(obj31, "MaidenPic6")
	ns.SetDialog(obj31, ns.NORMAL, Maiden2TalkStart, Maiden2TalkEnd)
	ns.StoryPic(obj41, "WardenPic")
	ns.SetDialog(obj41, ns.NORMAL, JacobDialogStart, JacobDialogEnd)
	obj98 = ns.Object("Bookcase")
	wp100 = ns.Waypoint("BookcaseWP")
	gvar99 = ns.WallGroup("MayorSecretWallGroup")
	wp101 = ns.Waypoint("Secret01WP")
	wp102 = ns.Waypoint("Secret02WP")
	InitializeHenrick()
}
func MapEntry() {
	obj18 = ns.Object("War03c:MayorScepter")
	ns.UnBlind()
	PlayTownMusic()
}
func QuestSwitch() {
	if !ns.HasItem(ns.GetHost(), ns.Object("War03c:MayorScepter")) {
		goto LABEL1
	}
	ns.ObjectGroupOff(gvar22)
	gvar93 = gvar91
	ns.SetShopkeeperText(obj24, "War03b:BarkeeperPost")
	gvar88 = gvar87
	gvar97 = gvar96
	gvar73 = gvar72
	gvar76 = gvar75
	gvar79 = gvar78
	gvar82 = gvar81
	ns.CancelDialog(obj31)
	ns.SetQuestStatus(1, "QuestComplete")
	ns.LockDoor(obj47)
	ns.LockDoor(obj48)
LABEL1:
	return
}
func MayorGreet() {
	ns.CreatureFollow(obj15, ns.GetHost())
	ns.ObjectGroupOff(gvar14)
	MayorSpeaks()
}
func EnableWaitTrigGroup1() {
	ns.ObjectGroupOn(gvar43)
}
func Wait1() {
	if !ns.IsCaller(obj26) {
		goto LABEL1
	}
	ns.PauseObject(obj26, 500)
	ns.ObjectGroupOff(gvar43)
LABEL1:
	return
}
func EnableWaitTrigGroup2() {
	ns.ObjectGroupOn(gvar44)
}
func Wait2() {
	if !ns.IsCaller(obj27) {
		goto LABEL1
	}
	ns.PauseObject(obj27, 500)
	ns.ObjectGroupOff(gvar44)
LABEL1:
	return
}
func EnableWaitTrigGroup3() {
	ns.ObjectGroupOn(gvar45)
}
func Wait3() {
	if !ns.IsCaller(obj28) {
		goto LABEL1
	}
	ns.PauseObject(obj28, 500)
	ns.ObjectGroupOff(gvar45)
LABEL1:
	return
}
func MonsterGoHome() {
	if !ns.HasClass(ns.GetCaller(), ns.MONSTER) {
		goto LABEL1
	}
	ns.AggressionLevel(ns.GetCaller(), 0.16)
	ns.GoBackHome(ns.GetCaller())
LABEL1:
	return
}
func FieldGuidePickup() {
	ns.Pickup(ns.GetHost(), obj34)
	ns.Pickup(ns.GetHost(), obj35)
}
func CaptainGuard() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
	)
	v0 = ns.GetWaypointX(wp55)
	v1 = ns.GetWaypointY(wp55)
	v2 = ns.GetWaypointX(wp57)
	v3 = ns.GetWaypointY(wp57)
	ns.CreatureGuard(obj54, v0, v1, v2, v3, 500)
}
func PlaySubMusic() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.Music(18, 100)
LABEL1:
	return
}
func PlayWanderMusic() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.Music(17, 100)
LABEL1:
	return
}
func OpenMayorSecretWallGroup() {
	ns.Move(obj98, wp100)
	ns.WallGroupOpen(gvar99)
}
func Secret01XP() {
	ns.AudioEvent(ns.SecretFound, wp101)
	ns.GiveXp(ns.GetHost(), 50)
	ns.PrintToAll("GeneralPrint:SecretFound")
}
func Secret02XP() {
	ns.AudioEvent(ns.SecretFound, wp102)
	ns.GiveXp(ns.GetHost(), 50)
	ns.PrintToAll("GeneralPrint:SecretFound")
}
func MakeAWish() {
	var v0 int
	v0 = ns.MaxHealth(ns.GetCaller())
	ns.RestoreHealth(ns.GetCaller(), v0-ns.CurrentHealth(ns.GetCaller()))
	ns.PrintToAll("GeneralPrint:WellSignRefresh")
	ns.AudioEvent(ns.RestoreHealthName, wp60)
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
