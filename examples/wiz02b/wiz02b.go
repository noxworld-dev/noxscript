package wiz02b

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	obj4    ns.ObjectID
	obj5    ns.ObjectID
	obj6    ns.ObjectID
	obj7    ns.ObjectID
	obj8    ns.ObjectID
	gvar9   ns.ObjectID
	gvar10  ns.ObjectID
	obj11   ns.ObjectID
	obj12   ns.ObjectID
	obj13   ns.ObjectID
	gvar14  ns.WallGroupID
	wp15    ns.WaypointID
	wp16    ns.WaypointID
	wp17    ns.WaypointID
	wp18    ns.WaypointID
	flag19  bool
	flag20  bool
	flag21  bool
	flag22  bool
	flag23  bool
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
	fvar36  float32
	fvar37  float32
	fvar38  float32
	fvar39  float32
	fvar40  float32
	fvar41  float32
	fvar42  float32
	fvar43  float32
	fvar44  float32
	fvar45  float32
	fvar46  float32
	fvar47  float32
	fvar48  float32
	fvar49  float32
	fvar50  float32
	fvar51  float32
	fvar52  float32
	fvar53  float32
	flag54  bool
	flag55  bool
	flag56  bool
	flag57  bool
	flag58  bool
	flag59  bool
	flag60  bool
	flag61  bool
	flag62  bool
	flag63  bool
	flag64  bool
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
	ivar78  int
	ivar79  int
	ivar80  int
	gvar81  int
	ivar82  int
	ivar83  int
	ivar84  int
	gvar85  int
	gvar86  int
	ivar87  int
	gvar88  int
	gvar89  int
	ivar90  int
	ivar91  int
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
	gvar108 ns.ObjectGroupID
	gvar109 ns.ObjectGroupID
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
	gvar127 ns.ObjectID
	gvar128 ns.ObjectID
	gvar129 ns.ObjectID
	obj130  ns.ObjectID
	obj131  ns.ObjectID
	obj132  ns.ObjectID
	obj133  ns.ObjectID
	obj134  ns.ObjectID
	obj135  ns.ObjectID
	gvar136 ns.ObjectGroupID
	gvar137 ns.ObjectGroupID
	gvar138 ns.WallGroupID
	gvar139 ns.WallGroupID
	gvar140 ns.WallGroupID
	gvar141 ns.WallGroupID
	gvar142 ns.WallGroupID
	wp143   ns.WaypointID
	wp144   ns.WaypointID
	wp145   ns.WaypointID
	wp146   ns.WaypointID
	wp147   ns.WaypointID
	gvar148 ns.WaypointGroupID
	gvar149 ns.WaypointGroupID
	gvar150 ns.WaypointGroupID
)

func init() {
	flag19 = false
	flag20 = false
	flag21 = false
	flag22 = false
	flag23 = false
	gvar24 = 0
	gvar25 = 1
	gvar26 = 2
	gvar27 = 3
	gvar28 = 4
	gvar29 = 5
	gvar30 = 6
	gvar31 = 7
	gvar32 = 8
	gvar33 = 9
	gvar34 = 10
	gvar35 = gvar24
	fvar36 = 4586
	fvar37 = 4430
	fvar38 = 4539
	fvar39 = 4373
	fvar44 = 4471
	fvar45 = 4423
	fvar46 = 4312
	fvar47 = 4257
	flag54 = false
	flag55 = false
	flag56 = false
	flag57 = false
	flag58 = false
	flag59 = false
	flag60 = false
	flag61 = false
	flag62 = false
	flag63 = false
	flag64 = false
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
	ivar78 = 0
	ivar79 = 0
	ivar80 = 0
	gvar81 = 0
	ivar82 = 0
	ivar83 = 0
	ivar84 = 0
	gvar85 = 0
	gvar86 = 1
	ivar87 = 1
	gvar89 = 0
	ivar90 = 0
	ivar91 = 0
}
func LibrarySetup() {
	obj4 = ns.Object("Librarian1")
	obj5 = ns.Object("Librarian2")
	obj6 = ns.Object("Archivist")
	obj7 = ns.Object("Necromancer")
	obj8 = ns.Object("FakeBook")
	obj11 = ns.Object("FakeBookMover")
	obj13 = ns.Object("SwapbookMover")
	gvar14 = ns.WallGroup("SecretExitWalls")
	wp15 = ns.Waypoint("AttackWP")
	wp16 = ns.Waypoint("NecroStealWP")
	wp17 = ns.Waypoint("NecroVanishWP")
	ns.FrameTimer(1, OwnPeople)
}
func OwnPeople() {
	ns.SetOwner(ns.GetHost(), obj4)
	ns.SetOwner(ns.GetHost(), obj5)
	ns.SetOwner(ns.GetHost(), obj6)
	ns.StoryPic(obj6, "ArchivistPic")
	ns.StoryPic(obj7, "NecromancerPic")
	ns.ObjectOn(obj11)
	ns.ObjectOn(obj13)
}
func TheftPieceCheck() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	if ns.GetQuestStatus("BookStolen") == 1 {
		goto LABEL1
	}
	ns.SetQuestStatus(1, "BookStolen")
	ns.WideScreen(true)
	ns.SetQuestStatus(1, "InLetterboxMode")
	ns.Frozen(ns.GetHost(), true)
	ns.LookAtObject(ns.GetHost(), obj6)
	flag19 = true
	ns.Move(obj7, wp15)
	ns.Music(11, 100)
LABEL1:
	return
}
func TheftPieceBegin() {
	if !ns.IsCaller(obj7) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.Frozen(obj4, true)
	ns.Frozen(obj5, true)
	ns.Frozen(obj7, true)
	ns.Frozen(obj6, true)
	ns.CreatureIdle(obj4)
	ns.CreatureIdle(obj5)
	ns.CreatureIdle(obj7)
	ns.CreatureIdle(obj6)
	ns.LookAtObject(obj4, obj7)
	ns.LookAtObject(obj5, obj7)
	ns.LookAtObject(obj6, obj7)
	ns.LookAtObject(obj7, obj6)
	ns.SetDialog(obj6, ns.NEXT, ArchivistTalkStart, ArchivistTalkEnd)
	ns.StartDialog(obj6, ns.GetHost())
LABEL1:
	return
}
func Lib1Talk() {
	ns.DestroyEveryChat()
	ns.Frozen(obj4, false)
	ns.Frozen(obj5, false)
	ns.Frozen(obj7, false)
	ns.SecondTimer(1, StartTheFight)
}
func StartTheFight() {
	ns.Frozen(obj7, true)
	ns.CreatureIdle(obj7)
	ns.LookAtObject(obj7, obj4)
	ns.CastSpellObjectLocation(ns.SPELL_FIREBALL, obj7, ns.GetObjectX(obj4), ns.GetObjectY(obj4))
}
func Lib1Die() {
	ns.MoveWaypoint(wp18, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.NecromancerTaunt, wp18)
	ns.SecondTimer(1, ShootLib2)
}
func Lib2Die() {
	ns.AudioEvent(ns.NecromancerTaunt, wp18)
	ns.Frozen(obj7, false)
	ns.Chat(obj5, "Wiz02:MiscWizard02")
	ns.Move(obj7, wp16)
}
func ShootLib2() {
	if !(ns.CurrentHealth(obj5) > 0) {
		goto LABEL1
	}
	ns.LookAtObject(obj7, obj5)
	ns.CastSpellObjectLocation(ns.SPELL_FIREBALL, obj7, ns.GetObjectX(obj5), ns.GetObjectY(obj5))
	ns.FrameTimer(1, ShootLib2)
LABEL1:
	return
}
func NecroExit() {
	if !ns.IsCaller(obj7) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	gvar35 = gvar29
	ns.Frozen(obj7, true)
	ns.CreatureIdle(obj7)
	ns.LookAtObject(obj7, obj6)
	ns.SetDialog(obj7, ns.NEXT, NecromancerStart, ArchivistTalkEnd)
	flag23 = true
	ns.StartDialog(obj7, ns.GetHost())
LABEL1:
	return
}
func NecroWalkAway() {
	ns.Move(obj7, wp17)
	gvar35 = gvar31
	ns.SetDialog(obj6, ns.NORMAL, ArchivistTalkStart, ArchivistTalkEnd)
	ns.StartDialog(obj6, ns.GetHost())
}
func RemoveNecro() {
	if !ns.IsCaller(obj7) {
		goto LABEL1
	}
	if !flag23 {
		goto LABEL1
	}
	flag20 = true
	ns.ObjectOff(ns.GetTrigger())
	ns.MoveObject(obj7, 4004, 651)
	ns.ObjectOff(obj7)
	ns.SetQuestStatus(0, "InLetterboxMode")
	if !flag22 {
		goto LABEL1
	}
	ns.WideScreen(false)
	ns.Frozen(ns.GetHost(), false)
	ns.Frozen(obj6, false)
	ns.CreatureGuard(obj6, ns.GetObjectX(obj6), ns.GetObjectY(obj6), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	EnableGroupA()
	ns.Music(18, 100)
LABEL1:
	return
}
func OpenExitWalls() {
	ns.WallGroupOpen(gvar14)
	CheckArchivist()
}
func ArchivistTalkStart() {
	var v0 int
	ns.Frozen(ns.GetHost(), true)
	ns.CreatureIdle(ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj6)
	if !ns.HasItem(ns.GetHost(), gvar9) {
		goto LABEL1
	}
	if gvar35 == gvar34 {
		goto LABEL1
	}
	gvar35 = gvar33
LABEL1:
	v0 = gvar35
	if v0 == gvar24 {
		goto LABEL2
	}
	if v0 == gvar26 {
		goto LABEL3
	}
	if v0 == gvar31 {
		goto LABEL4
	}
	if v0 == gvar32 {
		goto LABEL5
	}
	if v0 == gvar33 {
		goto LABEL6
	}
	if v0 == gvar34 {
		goto LABEL7
	}
	goto LABEL8
LABEL2:
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02:ArchivistSurprise")
	gvar35 = gvar25
	goto LABEL8
LABEL3:
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02:ArchivistSnob")
	gvar35 = gvar27
	goto LABEL8
LABEL4:
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02:OldArchivist01")
	ns.CreatureGuard(obj6, 3650, 1056, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	goto LABEL8
LABEL5:
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02:OldArchivist02")
	ns.CreatureGuard(obj6, 3650, 1056, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	goto LABEL8
LABEL6:
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02:OldArchivist03")
	ns.CreatureGuard(obj6, 3650, 1056, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	goto LABEL8
LABEL7:
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02:OldArchivistTalk04")
	ns.CreatureGuard(obj6, 3650, 1056, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	goto LABEL8
LABEL8:
	return
}
func ArchivistTalkEnd() {
	var v0 int
	v0 = gvar35
	if v0 == gvar25 {
		goto LABEL1
	}
	if v0 == gvar26 {
		goto LABEL2
	}
	if v0 == gvar27 {
		goto LABEL3
	}
	if v0 == gvar28 {
		goto LABEL4
	}
	if v0 == gvar29 {
		goto LABEL5
	}
	if v0 == gvar30 {
		goto LABEL6
	}
	if v0 == gvar31 {
		goto LABEL7
	}
	if v0 == gvar32 {
		goto LABEL8
	}
	if v0 == gvar33 {
		goto LABEL9
	}
	if v0 == gvar34 {
		goto LABEL10
	}
	goto LABEL11
LABEL1:
	ns.SetDialog(obj7, ns.NEXT, NecromancerStart, ArchivistTalkEnd)
	ns.StartDialog(obj7, ns.GetHost())
	goto LABEL11
LABEL2:
	ns.SetDialog(obj6, ns.NEXT, ArchivistTalkStart, ArchivistTalkEnd)
	ns.StartDialog(obj6, ns.GetHost())
	goto LABEL11
LABEL3:
	ns.SetDialog(obj7, ns.NEXT, NecromancerStart, ArchivistTalkEnd)
	ns.StartDialog(obj7, ns.GetHost())
	goto LABEL11
LABEL4:
	ns.CancelDialog(obj7)
	ns.CancelDialog(obj6)
	Lib1Talk()
	goto LABEL11
LABEL5:
	ns.SetDialog(obj7, ns.NEXT, NecromancerStart, ArchivistTalkEnd)
	ns.StartDialog(obj7, ns.GetHost())
	goto LABEL11
LABEL6:
	ns.CancelDialog(obj7)
	ns.Pickup(obj7, obj8)
	ns.Frozen(obj7, false)
	ns.FrameTimer(1, NecroWalkAway)
	goto LABEL11
LABEL7:
	ns.SetDialog(obj6, ns.NORMAL, ArchivistTalkStart, ArchivistTalkEnd)
	if flag21 {
		goto LABEL12
	}
	ns.JournalEntry(ns.GetHost(), "GetNecromancer", 2)
	ns.PrintToAll("Con01a:NewJournalEntry")
	flag21 = true
LABEL12:
	gvar35 = gvar32
	flag22 = true
	if !flag20 {
		goto LABEL13
	}
	ns.WideScreen(false)
	ns.Frozen(ns.GetHost(), false)
	ns.Frozen(obj6, false)
	EnableGroupA()
	ns.Music(18, 100)
LABEL13:
	goto LABEL11
LABEL8:
	ns.Frozen(ns.GetHost(), false)
	ns.SetDialog(obj6, ns.NORMAL, ArchivistTalkStart, ArchivistTalkEnd)
	EnableGroupA()
	goto LABEL11
LABEL9:
	if ns.HasItem(ns.GetHost(), gvar10) {
		goto LABEL14
	}
	ns.Pickup(ns.GetHost(), gvar10)
	ns.JournalEdit(ns.GetHost(), "ArchivistBook", 4)
	ns.PrintToAll("Con02a:ObjectiveComplete")
	ns.GiveXp(ns.GetHost(), 300)
	ns.PrintToAll("GeneralPrint:GainedItem")
LABEL14:
	ns.Frozen(ns.GetHost(), false)
	ns.SetDialog(obj6, ns.NORMAL, ArchivistTalkStart, ArchivistTalkEnd)
	gvar35 = gvar34
	ns.Music(15, 100)
	goto LABEL11
LABEL10:
	ns.Frozen(ns.GetHost(), false)
	ns.SetDialog(obj6, ns.NORMAL, ArchivistTalkStart, ArchivistTalkEnd)
	goto LABEL11
LABEL11:
	return
}
func NecromancerStart() {
	var v0 int
	v0 = gvar35
	if v0 == gvar25 {
		goto LABEL1
	}
	if v0 == gvar27 {
		goto LABEL2
	}
	if v0 == gvar29 {
		goto LABEL3
	}
	goto LABEL4
LABEL1:
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02:Necromancer01")
	gvar35 = gvar26
	goto LABEL4
LABEL2:
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02:NecroTaunt")
	gvar35 = gvar28
	goto LABEL4
LABEL3:
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02:NecroHappy")
	gvar35 = gvar30
	goto LABEL4
LABEL4:
	return
}
func CheckArchivist() {
	if !(ns.Distance(ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), ns.GetObjectX(obj6), ns.GetObjectY(obj6)) <= 100) {
		goto LABEL1
	}
	ns.Frozen(ns.GetHost(), true)
	ns.CreatureIdle(ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj6)
	ns.LookAtObject(obj6, ns.GetHost())
	ns.StartDialog(obj6, ns.GetHost())
	goto LABEL2
LABEL1:
	ns.FrameTimer(15, CheckArchivist)
LABEL2:
	return
}
func ReleasePlayer() {
}
func EnableGroupA() {
	ns.ObjectGroupOn(gvar108)
}
func MainReceptionistGreetingStart() {
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02B.scr:MRTalk01")
	ns.UnlockDoor(obj110)
	ns.UnlockDoor(obj111)
	ns.SetQuestStatus(1, "MainReceptionistTalked")
}
func MainReceptionistGreetingEnd() {
	ns.CancelDialog(obj92)
}
func HorvathGivesTaskStart() {
	ns.JournalEdit(ns.GetHost(), "MeetHorvath", 4)
	ns.PrintToAll("Con02a:ObjectiveComplete")
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02B.scr:HorvathTalk01")
	ns.SetQuestStatus(1, "HorvathGaveAmuletQuest")
}
func HorvathGivesTaskEnd() {
	ns.SetDialog(obj95, ns.NEXT, HorvathGivesTask1aStart, HorvathGivesTask1aEnd)
	ns.StartDialog(obj95, ns.GetHost())
}
func HorvathGivesTask1aStart() {
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02B.scr:HorvathTalk01a")
}
func HorvathGivesTask1aEnd() {
	ns.SetDialog(obj95, ns.NORMAL, HorvathGivesTask1bStart, HorvathGivesTask1bEnd)
	ns.StartDialog(obj95, ns.GetHost())
}
func HorvathGivesTask1bStart() {
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02B.scr:HorvathTalk01b")
}
func HorvathGivesTask1bEnd() {
	ns.PrintToAll("Con01a:NewJournalEntry")
	ns.JournalEntry(ns.GetHost(), "ReturnToHorvath", 2)
	ns.SetDialog(obj95, ns.NORMAL, HorvathWaitingStart, HorvathWaitingEnd)
	ns.SetDialog(obj101, ns.NORMAL, ArchivistHappyStart, ArchivistHappyEnd)
}
func HorvathWaitingStart() {
	if ns.HasItem(ns.GetHost(), gvar10) {
		goto LABEL1
	}
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02B.scr:HorvathTalk02")
	goto LABEL2
LABEL1:
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02B.scr:HorvathTalk03")
LABEL2:
	return
}
func HorvathWaitingEnd() {
	if ns.HasItem(ns.GetHost(), gvar10) {
		goto LABEL1
	}
	ns.SetDialog(obj95, ns.NORMAL, HorvathWaitingStart, HorvathWaitingEnd)
	goto LABEL2
LABEL1:
	ns.CancelDialog(obj95)
	ns.GiveXp(ns.GetHost(), 1000)
	flag72 = true
	ns.JournalEdit(ns.GetHost(), "ReturnToHorvath", 4)
	ns.PrintToAll("Con02a:ObjectiveComplete")
	gvar127 = ns.GetLastItem(ns.GetHost())
	for {
		if gvar127 == 0 {
			goto LABEL3
		}
		gvar129 = ns.GetPreviousItem(gvar127)
		if gvar127 != gvar9 {
			goto LABEL4
		}
		ns.Delete(gvar127)
	LABEL4:
		gvar127 = gvar129
	}
LABEL3:
	ns.FrameTimer(1, DropAmulet)
LABEL2:
	return
}
func LewisIntroStart() {
	ns.LookAtObject(obj96, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02B.scr:LewisTalk01")
}
func LewisIntroEnd() {
	ns.Walk(obj96, 4582, 4363)
	ns.CancelDialog(obj96)
}
func LewisInPositionTrigger() {
	if !ns.IsCaller(obj96) {
		goto LABEL1
	}
	flag56 = true
	ns.Chat(obj96, "Wiz02B.scr:LewisTalk02")
	flag57 = true
	ns.Frozen(obj96, true)
	ns.Frozen(obj97, true)
	ns.CreatureIdle(obj96)
	ns.LookAtObject(obj96, ns.GetHost())
LABEL1:
	return
}
func LewisBookQuest() {
	ns.SetDialog(obj97, ns.NEXT, FrogTalk3Start, FrogTalk3End)
	ns.StartDialog(obj97, ns.GetHost())
}
func ReadyForFirstSwap() {
}
func ReadyForFirstReturn() {
	var v0 int
	v0 = ivar82
	if v0 == 30 {
		goto LABEL1
	}
	if v0 == 150 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.Chat(obj97, "Wiz02B.scr:LewisTalk04")
	goto LABEL3
LABEL2:
	ns.Chat(obj97, "Wiz02B.scr:LewisTalk05")
	flag58 = true
	goto LABEL3
LABEL3:
	ivar82 += 2
}
func ShowOnOff() {
	if !(flag59 == false && flag57 == true && flag60 == false) {
		goto LABEL1
	}
	ns.WideScreen(true)
	ns.Frozen(ns.GetHost(), true)
	ns.CreatureIdle(ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj96)
	flag59 = true
	FrogRoomStuff()
LABEL1:
	if !(flag59 == false && flag58 == true && flag61 == false) {
		goto LABEL2
	}
	ns.WideScreen(true)
	ns.Frozen(ns.GetHost(), true)
	ns.CreatureIdle(ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj97)
	flag59 = true
	FrogRoomStuff()
LABEL2:
	return
}
func GetArchivistButton() {
	flag76 = true
	flag77 = false
	ns.MoveWaypoint(wp18, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.ShortBellsUp, wp18)
	ns.WayPointGroupOff(gvar150)
	ns.WayPointGroupOn(gvar149)
	ns.Wander(obj101)
	if ns.GetQuestStatus("HorvathGaveAmuletQuest") != 1 {
		goto LABEL1
	}
	if ns.GetQuestStatus("LockLibraryGates") == 1 {
		goto LABEL2
	}
	if !ns.IsLocked(obj112) {
		goto LABEL2
	}
	ns.SetDialog(obj101, ns.NORMAL, ArchivistHappyStart, ArchivistHappyEnd)
LABEL2:
	goto LABEL3
LABEL1:
	ns.SetDialog(obj101, ns.NORMAL, ArchivistMadStart, ArchivistMadEnd)
LABEL3:
	return
}
func ArchivistAtGate() {
	if !ns.IsCaller(obj101) {
		goto LABEL1
	}
	ns.WayPointGroupOff(gvar149)
	if ns.IsTalking() {
		goto LABEL1
	}
	if !(flag76 == true && flag77 == false) {
		goto LABEL1
	}
	ns.Chat(obj101, "Wiz02B.scr:ArchivistTalk01")
	flag77 = true
	flag76 = false
LABEL1:
	return
}
func ArchivistGuard() {
	if !ns.IsCaller(obj101) {
		goto LABEL1
	}
	ns.CreatureGuard(obj101, ns.GetObjectX(obj101), ns.GetObjectY(obj101), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
LABEL1:
	return
}
func ArchivistHappyStart() {
	ns.DestroyEveryChat()
	ns.Frozen(obj101, true)
	ns.CreatureIdle(obj101)
	ns.LookAtObject(obj101, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02B.scr:ArchivistTalk05")
	flag76 = false
	flag77 = true
}
func ArchivistHappyEnd() {
	ns.Frozen(obj101, false)
	ns.ObjectOff(obj135)
	ns.UnlockDoor(obj112)
	ns.UnlockDoor(obj113)
	ns.ObjectOff(obj130)
	ns.WayPointGroupOff(gvar149)
	ns.WayPointGroupOn(gvar150)
	ns.CreatureIdle(obj101)
	ns.Wander(obj101)
	ns.CancelDialog(obj101)
}
func ArchivistMadStart() {
	var v0 int
	ns.DestroyEveryChat()
	ns.Frozen(obj101, true)
	ns.CreatureIdle(obj101)
	ns.LookAtObject(obj101, ns.GetHost())
	flag76 = false
	flag77 = true
	v0 = ivar87
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
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02B.scr:ArchivistTalk02")
	ivar87 += 1
	goto LABEL4
LABEL2:
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02B.scr:ArchivistTalk03")
	ivar87 += 1
	goto LABEL4
LABEL3:
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02B.scr:ArchivistTalk04")
	ivar78 += 1
	goto LABEL4
LABEL4:
	return
}
func ArchivistMadEnd() {
	ns.WayPointGroupOff(gvar149)
	ns.WayPointGroupOn(gvar150)
	ns.Frozen(obj101, false)
	ns.CreatureIdle(obj101)
	ns.Wander(obj101)
}
func OpenLibrarySecret1() {
	ns.MoveWaypoint(wp18, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretWallWoodOpen, wp18)
	ns.NoWallSound(true)
	ns.ObjectOn(obj131)
	ns.ObjectOn(obj132)
	ns.WallOpen(ns.Wall(141, 99))
	ns.NoWallSound(false)
	ns.ObjectGroupOn(gvar109)
	SwapRoomCheck()
}
func CloseLibrarySecret1() {
	if flag64 {
		goto LABEL1
	}
	ns.WallClose(ns.Wall(141, 99))
	flag64 = true
LABEL1:
	return
}
func OpenClassSecretWalls() {
	ns.WallGroupOpen(gvar141)
	ns.UnlockDoor(obj121)
	flag71 = true
}
func OpenSwapWalls() {
	ns.WallGroupOpen(gvar140)
}
func CloseSpiderWalls() {
	ivar90 += 1
	ns.WallGroupClose(gvar139)
	flag65 = true
	OpenSwapWalls()
	if !(ns.CurrentHealth(obj102) > 0) {
		goto LABEL1
	}
	ns.CreatureIdle(obj102)
LABEL1:
	if !(ns.CurrentHealth(obj103) > 0) {
		goto LABEL2
	}
	ns.CreatureIdle(obj103)
LABEL2:
	if !(ns.CurrentHealth(obj104) > 0) {
		goto LABEL3
	}
	ns.CreatureIdle(obj104)
LABEL3:
	return
}
func SpiderSwapCountDown() {
	ivar90 -= 1
}
func LibrarianTalkStart() {
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02B.scr:Librarian3")
}
func LibrarianTalkEnd() {
	ns.CancelDialog(obj107)
}
func GiveLightBook() {
	ns.Pickup(ns.GetHost(), obj126)
}
func CheckIfReturningToMap() {
	flag67 = true
	if ns.GetQuestStatus("FirstSwapComplete") != 1 {
		goto LABEL1
	}
	ns.ObjectOff(obj98)
	ns.MoveObject(obj97, 4564, 4382)
	ns.MoveObject(obj96, 4496, 4684)
	ns.WayPointGroupOn(gvar148)
	ns.Wander(obj96)
	flag55 = true
	flag56 = true
	flag57 = true
	flag58 = true
	flag59 = false
	flag60 = true
	flag61 = true
	flag62 = true
	flag63 = true
	ns.SetOwner(ns.GetHost(), obj97)
	ns.SetOwner(ns.GetHost(), obj96)
	ns.UnlockDoor(obj110)
	ns.UnlockDoor(obj111)
LABEL1:
	if ns.GetQuestStatus("MainReceptionistTalked") != 1 {
		goto LABEL2
	}
	if !(ns.IsLocked(obj110) && ns.IsLocked(obj111)) {
		goto LABEL2
	}
	ns.UnlockDoor(obj110)
	ns.UnlockDoor(obj111)
LABEL2:
	if ns.GetQuestStatus("SwapBookTaken") != 1 {
		goto LABEL3
	}
	ns.Delete(obj12)
LABEL3:
	if ns.GetQuestStatus("LockLibraryGates") != 1 {
		goto LABEL4
	}
	ns.LockDoor(obj112)
	ns.LockDoor(obj113)
LABEL4:
	return
}
func FinalSwapRoomCheck() {
	flag68 = true
}
func PlayerMadeFinalSwap() {
	ivar91 += 1
}
func FinalSwapCountDown() {
	ivar91 -= 1
}
func FrogRoomStuff() {
	if !(flag56 == true && flag57 == true && flag59 == false && flag61 == false) {
		goto LABEL1
	}
	ns.LookAtObject(obj96, obj99)
LABEL1:
	if !(flag59 == true && flag60 == false) {
		goto LABEL2
	}
	if ns.GetQuestStatus("FirstSwapComplete") == 1 {
		goto LABEL2
	}
	ns.Effect(ns.LIGHTNING, fvar40, fvar41, fvar42, fvar43)
	ns.Effect(ns.LIGHTNING, fvar44, fvar48, fvar45, fvar49)
	ns.MoveWaypoint(wp18, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	if ivar79 != 0 {
		goto LABEL3
	}
	ns.AudioEvent(ns.LightningBolt, wp18)
	ivar79 += 2
	goto LABEL4
LABEL3:
	ivar79 += 2
LABEL4:
	if !(ivar79 > 8) {
		goto LABEL5
	}
	ivar79 = 0
LABEL5:
	if ivar80 != 0 {
		goto LABEL6
	}
	ns.Effect(ns.SMOKE_BLAST, ns.GetObjectX(obj96), ns.GetObjectY(obj96), 0, 0)
	ns.Effect(ns.SMOKE_BLAST, ns.GetObjectX(obj97), ns.GetObjectY(obj97), 0, 0)
	if !(ivar83 < 92) {
		goto LABEL7
	}
	if flag73 {
		goto LABEL7
	}
	ns.Chat(obj96, "Wiz02B.scr:LewisTalk03")
	ns.Chat(obj97, "Wiz02B.scr:FrogTalk01")
	flag73 = true
LABEL7:
	if ivar83 != 92 {
		goto LABEL8
	}
	flag73 = false
LABEL8:
	if !(ivar83 > 92 && ivar83 < 182) {
		goto LABEL9
	}
	if flag73 {
		goto LABEL9
	}
	ns.Chat(obj96, "Wiz02B.scr:FrogTalk01")
	ns.Chat(obj97, "Wiz02B.scr:LewisTalk03")
	flag73 = true
LABEL9:
	ivar80 += 2
	goto LABEL10
LABEL6:
	ivar80 += 2
LABEL10:
	if !(ivar80 > 14) {
		goto LABEL11
	}
	ivar80 = 0
LABEL11:
	if !(fvar43 < 4333) {
		goto LABEL12
	}
	fvar41 = fvar37
	fvar43 = fvar39
	goto LABEL13
LABEL12:
	fvar41 -= 2
	fvar43 -= 2
LABEL13:
	if !(fvar49 < 4217) {
		goto LABEL14
	}
	fvar48 = fvar46
	fvar49 = fvar47
	goto LABEL15
LABEL14:
	fvar48 -= 2
	fvar49 -= 2
LABEL15:
	if ivar83 != 92 {
		goto LABEL16
	}
	fvar52 = ns.GetObjectX(obj97)
	fvar53 = ns.GetObjectY(obj97)
	fvar50 = ns.GetObjectX(obj96)
	fvar51 = ns.GetObjectY(obj96)
	ns.ObjectOff(obj98)
	ns.Frozen(obj96, false)
	ns.Frozen(obj97, false)
	ns.MoveObject(obj96, fvar52, fvar53)
	ns.MoveObject(obj97, fvar50, fvar51)
	ns.CreatureGuard(obj97, ns.GetObjectX(obj97), ns.GetObjectY(obj97), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	ns.CreatureGuard(obj96, ns.GetObjectX(obj96), ns.GetObjectY(obj96), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
LABEL16:
	ivar83 += 2
	if !(ivar83 > 152) {
		goto LABEL17
	}
	flag60 = true
	flag59 = false
	ns.SetQuestStatus(1, "FirstSwapComplete")
	flag73 = false
	ns.SetDialog(obj97, ns.NEXT, FrogTalk1Start, FrogTalk1End)
	ns.StartDialog(obj97, ns.GetHost())
	goto LABEL2
LABEL17:
	ns.FrameTimer(2, FrogRoomStuff)
LABEL2:
	if !(flag58 == true && flag59 == true) {
		goto LABEL18
	}
	ns.Effect(ns.LIGHTNING, fvar40, fvar41, fvar42, fvar43)
	ns.Effect(ns.LIGHTNING, fvar44, fvar48, fvar45, fvar49)
	ns.MoveWaypoint(wp18, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	if ivar79 != 0 {
		goto LABEL19
	}
	ns.AudioEvent(ns.LightningBolt, wp18)
	ivar79 += 2
	goto LABEL20
LABEL19:
	ivar79 += 2
LABEL20:
	if !(ivar79 > 8) {
		goto LABEL21
	}
	ivar79 = 0
LABEL21:
	if ivar80 != 0 {
		goto LABEL22
	}
	ns.Effect(ns.SMOKE_BLAST, ns.GetObjectX(obj96), ns.GetObjectY(obj96), 0, 0)
	ns.Effect(ns.SMOKE_BLAST, ns.GetObjectX(obj97), ns.GetObjectY(obj97), 0, 0)
	if !(ivar84 < 92) {
		goto LABEL23
	}
	if flag73 {
		goto LABEL23
	}
	ns.Chat(obj97, "Wiz02B.scr:LewisTalk03")
	ns.Chat(obj96, "Wiz02B.scr:FrogTalk01")
	flag73 = true
LABEL23:
	if ivar84 != 92 {
		goto LABEL24
	}
	flag73 = false
LABEL24:
	if !(ivar84 > 92 && ivar83 < 182) {
		goto LABEL25
	}
	if flag73 {
		goto LABEL25
	}
	ns.Chat(obj96, "Wiz02B.scr:FrogTalk01")
	ns.Chat(obj97, "Wiz02B.scr:LewisTalk03")
	flag73 = true
LABEL25:
	ivar80 += 2
	goto LABEL26
LABEL22:
	ivar80 += 2
LABEL26:
	if !(ivar80 > 14) {
		goto LABEL27
	}
	ivar80 = 0
LABEL27:
	if !(fvar43 < 4333) {
		goto LABEL28
	}
	fvar41 = fvar37
	fvar43 = fvar39
	goto LABEL29
LABEL28:
	fvar41 -= 2
	fvar43 -= 2
LABEL29:
	if !(fvar49 < 4217) {
		goto LABEL30
	}
	fvar48 = fvar46
	fvar49 = fvar47
	goto LABEL31
LABEL30:
	fvar48 -= 2
	fvar49 -= 2
LABEL31:
	ivar84 += 2
	if !(ivar84 > 152) {
		goto LABEL32
	}
	flag61 = true
	flag59 = false
	ns.SetQuestStatus(1, "FirstReverseComplete")
	goto LABEL18
LABEL32:
	ns.FrameTimer(2, FrogRoomStuff)
LABEL18:
	if !(flag61 == true && flag63 == false) {
		goto LABEL33
	}
	LewisBookQuest()
LABEL33:
	return
}
func FrogTalk1Start() {
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02B.scr:LewisTalk04")
}
func FrogTalk1End() {
	ns.SetDialog(obj97, ns.NORMAL, FrogTalk2Start, FrogTalk2End)
	ns.StartDialog(obj97, ns.GetHost())
}
func FrogTalk2Start() {
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02B.scr:LewisTalk05")
}
func FrogTalk2End() {
	ns.WideScreen(false)
	ns.Frozen(ns.GetHost(), false)
	flag58 = true
	ns.CancelDialog(obj97)
}
func FrogTalk3Start() {
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02B.scr:LewisTalk06")
}
func FrogTalk3End() {
	ns.SetDialog(obj97, ns.NEXT, FrogTalk4Start, FrogTalk4End)
	ns.WayPointGroupOn(gvar148)
	ns.Wander(obj96)
	ns.StartDialog(obj97, ns.GetHost())
}
func FrogTalk4Start() {
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02B.scr:LewisTalk07")
}
func FrogTalk4End() {
	ns.SetDialog(obj97, ns.NORMAL, FrogTalk5Start, FrogTalk5End)
	ns.StartDialog(obj97, ns.GetHost())
}
func FrogTalk5Start() {
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02B.scr:LewisTalk08")
}
func FrogTalk5End() {
	if flag75 {
		goto LABEL1
	}
	flag75 = true
	ns.WideScreen(false)
	ns.Frozen(ns.GetHost(), false)
LABEL1:
	ns.SetDialog(obj97, ns.NORMAL, FrogTalk5Start, FrogTalk5End)
}
func TeacherBegin() {
	ns.SetDialog(obj94, ns.NORMAL, TeacherTalkStart, TeacherTalkEnd)
	ns.StartDialog(obj94, ns.GetHost())
}
func TeacherTalkStart() {
	ns.LookAtObject(obj94, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02B.scr:TeacherTalk03")
}
func TeacherTalkEnd() {
	ns.CancelDialog(obj94)
}
func MapSwitch() {
	ns.MoveObject(ns.GetHost(), 4682, 5302)
}
func CreatureSetup() {
	ns.SetOwner(ns.GetHost(), obj101)
	ns.SetOwner(ns.GetHost(), obj96)
	ns.SetOwner(ns.GetHost(), obj92)
	ns.SetOwner(ns.GetHost(), obj97)
	ns.SetOwner(ns.GetHost(), obj107)
	ns.SetOwner(ns.GetHost(), obj95)
	ns.SetOwner(ns.GetHost(), obj94)
	LibrarySetup()
}
func TrapSpiderCheck() {
	if !(ns.CurrentHealth(obj106) <= 0) {
		goto LABEL1
	}
	if !(ivar91 > 0) {
		goto LABEL1
	}
	ns.LookAtObject(obj106, ns.GetHost())
LABEL1:
	if !(ns.CurrentHealth(obj106) <= 0) {
		goto LABEL2
	}
	if !(ivar91 <= 0) {
		goto LABEL2
	}
	obj106 = ns.CreateObject("SmallAlbinoSpider", wp147)
LABEL2:
	ns.FrameTimer(8, TrapSpiderCheck)
}
func SwapRoomCheck() {
	if !(ns.CurrentHealth(obj105) > 0) {
		goto LABEL1
	}
	if !(ivar90 > 0) {
		goto LABEL1
	}
	ns.LookAtObject(obj105, ns.GetHost())
LABEL1:
	if !flag65 {
		goto LABEL2
	}
	if !(ns.CurrentHealth(obj105) <= 0) {
		goto LABEL2
	}
	if !(ivar90 > 0) {
		goto LABEL2
	}
	if gvar86 != 1 {
		goto LABEL3
	}
	obj105 = ns.CreateObject("SmallAlbinoSpider", wp145)
	gvar86 = 2
	goto LABEL2
LABEL3:
	obj105 = ns.CreateObject("SmallAlbinoSpider", wp146)
	gvar86 = 1
LABEL2:
	if ns.GetQuestStatus("SwapBookTaken") == 1 {
		goto LABEL4
	}
	if !(ivar90 <= 0) {
		goto LABEL4
	}
	ns.WallGroupOpen(gvar139)
LABEL4:
	if ns.GetQuestStatus("SwapBookTaken") == 1 {
		goto LABEL5
	}
	ns.FrameTimer(15, SwapRoomCheck)
LABEL5:
	return
}
func DropAmulet() {
	gvar128 = ns.GetLastItem(ns.GetHost())
	for {
		if gvar128 == 0 {
			goto LABEL1
		}
		gvar129 = ns.GetPreviousItem(gvar128)
		if gvar128 != gvar10 {
			goto LABEL2
		}
		ns.Delete(gvar128)
	LABEL2:
		gvar128 = gvar129
	}
LABEL1:
	ns.Blind()
	ns.FrameTimer(60, MapSwitch)
}
func NoBubbles() {
	ns.DestroyEveryChat()
}
func MapEntry() {
	ns.UnBlind()
	gvar10 = ns.Object("AmuletOfClarity")
	gvar9 = ns.Object("Wiz02C:BookOfOblivion")
	if ns.GetQuestStatus("InLetterboxMode") != 1 {
		goto LABEL1
	}
	ns.WideScreen(true)
LABEL1:
	ns.Music(15, 100)
	CheckIfReturningToMap()
}
func MapInitialize() {
	obj92 = ns.Object("Main_Receptionist")
	obj93 = ns.Object("Mystic")
	obj95 = ns.Object("Horvath")
	obj96 = ns.Object("Lewis")
	obj97 = ns.Object("Frog")
	obj98 = ns.Object("FrogMover")
	obj99 = ns.Object("LewisFaceThis")
	obj100 = ns.Object("LabEntrance")
	obj101 = ns.Object("Zaris")
	obj102 = ns.Object("BigSpider1")
	obj103 = ns.Object("BigSpider2")
	obj104 = ns.Object("BigSpider3")
	gvar108 = ns.ObjectGroup("LibraryGroupA")
	gvar109 = ns.ObjectGroup("LibraryGroupB")
	obj130 = ns.Object("ElevatorChain")
	obj110 = ns.Object("SDoor1")
	obj111 = ns.Object("SDoor2")
	obj112 = ns.Object("ArchiveGate1")
	obj113 = ns.Object("ArchiveGate2")
	obj114 = ns.Object("ArchiveDoor1")
	obj135 = ns.Object("ArchiveDoorChain")
	obj115 = ns.Object("SouthHallDoor1")
	obj116 = ns.Object("SouthHallDoor2")
	obj117 = ns.Object("SouthDoorUp1")
	obj118 = ns.Object("SouthDoorUp2")
	obj119 = ns.Object("NorthDoor1")
	obj120 = ns.Object("NorthDoor2")
	obj122 = ns.Object("VaultDoor")
	obj121 = ns.Object("ClassDoor")
	obj123 = ns.Object("BookCase1")
	obj124 = ns.Object("BookCase2")
	obj131 = ns.Object("Mover1")
	obj132 = ns.Object("Mover2")
	obj107 = ns.Object("Receptionist1")
	obj125 = ns.Object("VaultKey")
	obj126 = ns.Object("LightBook")
	obj94 = ns.Object("Teacher")
	obj12 = ns.Object("SwapBook")
	obj133 = ns.Object("HorvathBookcaseMover1")
	obj134 = ns.Object("HorvathBookcaseMover2")
	gvar136 = ns.ObjectGroup("HecubahTriggers")
	gvar137 = ns.ObjectGroup("SwapTriggers")
	gvar140 = ns.WallGroup("SwapWalls")
	gvar138 = ns.WallGroup("LabWalls")
	gvar139 = ns.WallGroup("SpiderWalls")
	gvar141 = ns.WallGroup("ClassSecretWalls")
	gvar142 = ns.WallGroup("HorvathWalls")
	wp18 = ns.Waypoint("PlayerSounds")
	wp143 = ns.Waypoint("TeacherSounds")
	wp144 = ns.Waypoint("LibraryExitWP")
	wp145 = ns.Waypoint("SpiderWP1")
	wp146 = ns.Waypoint("SpiderWP2")
	wp147 = ns.Waypoint("TrapSpiderWP")
	gvar148 = ns.WaypointGroup("FrogWaypoints")
	gvar149 = ns.WaypointGroup("ArchivistToGateWP")
	gvar150 = ns.WaypointGroup("ArchivistGoBackWP")
	fvar40 = fvar36
	fvar41 = fvar37
	fvar42 = fvar38
	fvar43 = fvar39
	fvar48 = fvar46
	fvar49 = fvar47
	ns.WayPointGroupOff(gvar148)
	ns.WayPointGroupOff(gvar149)
	ns.WayPointGroupOff(gvar150)
	ns.ObjectOn(obj98)
	obj105 = ns.CreateObject("SmallAlbinoSpider", wp145)
	obj106 = ns.CreateObject("SmallAlbinoSpider", wp147)
	ns.LockDoor(obj100)
	ns.LockDoor(obj110)
	ns.LockDoor(obj111)
	ns.LockDoor(obj115)
	ns.LockDoor(obj116)
	ns.LockDoor(obj117)
	ns.LockDoor(obj118)
	ns.LockDoor(obj119)
	ns.LockDoor(obj120)
	ns.LockDoor(obj121)
	ns.LockDoor(obj112)
	ns.LockDoor(obj113)
	if ns.GetQuestStatus("MainReceptionistTalked") == 1 {
		goto LABEL1
	}
	ns.SetQuestStatus(0, "MainReceptionistTalked")
	ns.SetDialog(obj92, ns.NORMAL, MainReceptionistGreetingStart, MainReceptionistGreetingEnd)
LABEL1:
	ns.SetDialog(obj107, ns.NORMAL, LibrarianTalkStart, LibrarianTalkEnd)
	ns.SetDialog(obj96, ns.NORMAL, LewisIntroStart, LewisIntroEnd)
	if ns.GetQuestStatus("HorvathGaveAmuletQuest") != 1 {
		goto LABEL2
	}
	ns.SetDialog(obj95, ns.NORMAL, HorvathWaitingStart, HorvathWaitingEnd)
	goto LABEL3
LABEL2:
	ns.SetDialog(obj95, ns.NEXT, HorvathGivesTaskStart, HorvathGivesTaskEnd)
LABEL3:
	ns.StoryPic(obj92, "WizardGuard1Pic")
	ns.StoryPic(obj96, "WizardGuard2Pic")
	ns.StoryPic(obj101, "WizardGuard1Pic")
	ns.StoryPic(obj97, "FrogPic")
	ns.StoryPic(obj94, "WizardGuard2Pic")
	ns.StoryPic(obj95, "HorvathPic")
	ns.ObjectGroupOff(gvar108)
	ns.ObjectGroupOff(gvar109)
	ns.FrameTimer(1, CreatureSetup)
	CheckIfReturningToMap()
}
func PlayerDeath() {
	ns.DeathScreen(2)
}
func OnEvent(typ string) {
	switch typ {
	case "MapEntry":
		MapEntry()
	case "MapInitialize":
		MapInitialize()
	case "PlayerDeath":
		PlayerDeath()
	}
}
