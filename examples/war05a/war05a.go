package war05a

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
	wp34    ns.WaypointID
	wp35    ns.WaypointID
	wp36    ns.WaypointID
	wp37    ns.WaypointID
	wp38    ns.WaypointID
	wp39    ns.WaypointID
	wp40    ns.WaypointID
	wp41    ns.WaypointID
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
	gvar67  ns.ObjectGroupID
	flag68  bool
	ivar69  int
	gvar70  int
	gvar71  int
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
	wp146   int
	wp147   int
	obj148  ns.ObjectID
	obj149  ns.ObjectID
	obj150  ns.ObjectID
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
	obj161  ns.ObjectID
	obj162  ns.ObjectID
	obj163  ns.ObjectID
	obj164  ns.ObjectID
	obj165  ns.ObjectID
	obj166  ns.ObjectID
	obj167  ns.ObjectID
	obj168  ns.ObjectID
	obj169  ns.ObjectID
	obj170  ns.ObjectID
	obj171  ns.ObjectID
	obj172  ns.ObjectID
	obj173  ns.ObjectID
	obj174  ns.ObjectID
	gvar175 int
	gvar176 int
	gvar177 int
	gvar178 int
	obj179  ns.ObjectID
	obj180  ns.ObjectID
	obj181  ns.ObjectID
	obj182  ns.ObjectID
	obj183  ns.ObjectID
	obj184  ns.ObjectID
	obj185  ns.ObjectID
	obj186  ns.ObjectID
	obj187  ns.ObjectID
	gvar188 int
	wp189   ns.WaypointID
	wp190   ns.WaypointID
	wp191   ns.WaypointID
	wp192   ns.WaypointID
	wp193   ns.WaypointID
	wp194   ns.WaypointID
	wp195   ns.WaypointID
	obj196  ns.ObjectID
	obj197  ns.ObjectID
	obj198  ns.ObjectID
	obj199  ns.ObjectID
	gvar200 int
	gvar201 int
	gvar202 int
	gvar203 int
	wp204   ns.WaypointID
	wp205   ns.WaypointID
	wp206   ns.WaypointID
	wp207   ns.WaypointID
	wp208   ns.WaypointID
	wp209   ns.WaypointID
	wp210   ns.WaypointID
	wp211   ns.WaypointID
	wp212   ns.WaypointID
	wp213   ns.WaypointID
	wp214   ns.WaypointID
	wp215   ns.WaypointID
	wp216   ns.WaypointID
	wp217   ns.WaypointID
	wp218   ns.WaypointID
	wp219   ns.WaypointID
	wp220   ns.WaypointID
	wp221   ns.WaypointID
	wp222   ns.WaypointID
	wp223   ns.WaypointID
	wp224   ns.WaypointID
	wp225   ns.WaypointID
	wp226   ns.WaypointID
	wp227   ns.WaypointID
	wp228   ns.WaypointID
	ivar229 int
	gvar230 int
	gvar231 int
	gvar232 int
	gvar233 int
	gvar234 int
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
	gvar26 = 2
	gvar27 = 3
	gvar28 = 0
	gvar29 = 1
	gvar30 = 2
	gvar31 = 3
	gvar32 = 0
	gvar33 = 1
	flag68 = false
	ivar69 = 0
	gvar70 = gvar4
	gvar71 = gvar10
	gvar72 = gvar12
	gvar73 = gvar20
	gvar74 = gvar14
	gvar75 = gvar24
	gvar76 = gvar28
	gvar77 = gvar32
	gvar84 = 0
	gvar85 = 0
	gvar86 = 0
	ivar229 = 0
	gvar230 = 0
	gvar231 = 1
	gvar232 = 2
	gvar233 = 3
	gvar234 = gvar230
}
func OgreDeath() {
	ivar69 += 1
	if !ns.IsTrigger(obj51) {
		goto LABEL1
	}
	ns.CreatureGuard(obj130, ns.GetWaypointX(wp191), ns.GetWaypointY(wp191), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
LABEL1:
	if !(ivar69 >= 7) {
		goto LABEL2
	}
	SetTownDialog()
	ns.Music(23, 100)
LABEL2:
	return
}
func SetTownDialog() {
	ns.SetDialog(obj179, ns.NORMAL, Villager1Start, Villager1End)
	ns.StoryPic(obj179, "MalePic8")
	ns.SetDialog(obj180, ns.NORMAL, Villager2Start, Villager2End)
	ns.StoryPic(obj180, "MalePic3")
	ns.SetDialog(obj181, ns.NORMAL, Villager3Start, Villager3End)
	ns.StoryPic(obj181, "MalePic2")
	ns.SetDialog(obj182, ns.NORMAL, Villager4Start, Villager4End)
	ns.StoryPic(obj182, "MalePic14")
	ns.SetDialog(obj183, ns.NORMAL, Villager5Start, Villager5End)
	ns.StoryPic(obj183, "MalePic9")
	ns.SetDialog(obj130, ns.NORMAL, FarmerSavedStart, FarmerSavedEnd)
}
func CheckFlame() {
	if !ns.IsObjectOn(obj66) {
		goto LABEL1
	}
	ns.PrintToAll("Flame Exists")
	goto LABEL2
LABEL1:
	ns.PrintToAll("Flame all gone :(")
LABEL2:
	return
}
func FarmerSavedStart() {
	ns.TellStory(ns.HumanMaleEatFood, "Wiz05A.scr:FarmerIdle")
}
func FarmerSavedEnd() {
	ns.SetDialog(obj130, ns.NORMAL, FarmerSavedStart, FarmerSavedEnd)
}
func BearGoAway() {
	if !ns.IsCaller(obj59) {
		goto LABEL1
	}
	flag68 = true
	ns.GoBackHome(obj59)
LABEL1:
	return
}
func BearReport() {
	if !flag68 {
		goto LABEL1
	}
	ns.AggressionLevel(obj59, 0.5)
	ns.CreatureGuard(obj59, ns.GetWaypointX(wp41), ns.GetWaypointY(wp41), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 300)
	flag68 = false
LABEL1:
	return
}
func KillCorpses() {
	ns.FrameTimer(30, KillCorpses2)
}
func DisableVillagers() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
	)
	v0 = ns.GetWaypointX(wp191)
	v1 = ns.GetWaypointY(wp191)
	v2 = ns.GetWaypointX(wp192)
	v3 = ns.GetWaypointY(wp192)
	ns.CreatureGuard(obj130, v0, v1, v2, v3, 0)
}
func LockDoors() {
	ns.LockDoor(obj198)
	ns.LockDoor(obj199)
	ns.LockDoor(obj196)
	ns.LockDoor(obj197)
}
func SwordsmanStart() {
	var v0 int
	ns.Frozen(obj144, true)
	ns.CreatureIdle(obj144)
	ns.LookAtObject(obj144, ns.GetHost())
	v0 = gvar72
	if v0 == gvar12 {
		goto LABEL1
	}
	if v0 == gvar13 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.LookAtObject(obj144, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:SwordsmanGreeting")
	goto LABEL3
LABEL2:
	ns.LookAtObject(obj144, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:SwordsmanEnd")
	goto LABEL3
LABEL3:
	return
}
func SwordsmanEnd() {
	var v0 int
	ns.Frozen(obj144, false)
	ns.CreatureGuard(obj144, ns.GetWaypointX(wp189), ns.GetWaypointY(wp189), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 240)
	v0 = gvar72
	if v0 == gvar12 {
		goto LABEL1
	}
	if v0 == gvar13 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.FrameTimer(1, SwordsmanTriggerC)
	gvar72 = gvar13
	goto LABEL3
LABEL2:
	ns.LookAtDirection(obj144, ns.NE)
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
	v2 = gvar70
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
	if !ns.HasItem(ns.GetHost(), obj127) {
		goto LABEL8
	}
	gvar70 = gvar7
	ns.SetDialog(obj130, ns.NORMAL, FarmerStart, FarmerEnd)
	FarmerStart()
	goto LABEL7
	goto LABEL2
LABEL8:
	ns.CreatureIdle(obj130)
	ns.LookAtObject(obj130, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:FarmerGreeting")
	goto LABEL7
LABEL2:
	ns.CreatureIdle(obj130)
	ns.LookAtObject(obj130, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:FarmerGreeting")
	goto LABEL7
LABEL3:
	ns.CreatureIdle(obj130)
	ns.LookAtObject(obj130, ns.GetHost())
	if !ns.HasItem(ns.GetHost(), obj127) {
		goto LABEL9
	}
	v0 = ns.GetWaypointX(wp191)
	v1 = ns.GetWaypointY(wp191)
	ns.AudioEvent(ns.CurePoisonCast, wp191)
	ns.AudioEvent(ns.CurePoisonEffect, wp191)
	ns.Effect(ns.BLUE_SPARKS, v0, v1, 0, 0)
	gvar70 = gvar7
	FarmerStart()
	goto LABEL4
LABEL9:
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:FarmerWaiting")
	goto LABEL7
LABEL4:
	ns.CreatureIdle(obj130)
	ns.LookAtObject(obj130, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:FarmerReturned")
	goto LABEL7
LABEL5:
	ns.CreatureIdle(obj130)
	ns.LookAtObject(obj130, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:FarmerIdle")
	goto LABEL7
LABEL6:
	if !ns.HasItem(ns.GetHost(), obj127) {
		goto LABEL10
	}
	gvar70 = gvar7
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
	v0 = ns.GetAnswer(obj130)
	v4 = gvar70
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
	gvar70 = gvar6
	ns.SetDialog(obj130, ns.NORMAL, FarmerStart, FarmerEnd)
	ns.PrintToAll("Con01a:NewJournalEntry")
	ns.JournalEntry(ns.GetHost(), "War05Quest8", 2)
	goto LABEL7
	goto LABEL2
LABEL8:
	if v0 != v2 {
		goto LABEL9
	}
	ns.SetDialog(obj130, ns.YESNO, FarmerStart, FarmerEnd)
	gvar70 = gvar5
	goto LABEL7
	goto LABEL2
LABEL9:
	if obj130 != v3 {
		goto LABEL2
	}
	gvar70 = gvar4
LABEL2:
	if v0 != v1 {
		goto LABEL10
	}
	gvar70 = gvar6
	ns.SetDialog(obj130, ns.NORMAL, FarmerStart, FarmerEnd)
	ns.PrintToAll("Con01a:NewJournalEntry")
	ns.JournalEntry(ns.GetHost(), "War05Quest8", 2)
	goto LABEL7
	goto LABEL3
LABEL10:
	if v0 != v2 {
		goto LABEL11
	}
	ns.SetDialog(obj130, ns.NORMAL, FarmerStart, FarmerEnd)
	gvar70 = gvar9
	goto LABEL7
	goto LABEL3
LABEL11:
	if obj130 != v3 {
		goto LABEL3
	}
	gvar70 = gvar4
LABEL3:
	goto LABEL7
LABEL4:
	ns.PrintToAll("War05A.scr:ObjectiveComplete")
	ns.JournalEdit(ns.GetHost(), "War05Quest8", 4)
	ns.GiveXp(ns.GetHost(), 250)
	ns.Drop(obj130, obj125)
	ns.Drop(obj130, obj126)
	ns.Delete(obj127)
	EndCurse()
	gvar70 = gvar8
	goto LABEL7
LABEL5:
	goto LABEL7
LABEL6:
	goto LABEL7
LABEL7:
	return
}
func HoundStart() {
	ns.CreatureIdle(obj131)
	ns.LookAtObject(obj131, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:HoundGreeting")
}
func HoundEnd() {
	ns.SetRoamFlag(obj131, 2)
	ns.Wander(obj131)
}
func FarmerWifeStart() {
	var v0 int
	v0 = gvar71
	if v0 == gvar10 {
		goto LABEL1
	}
	if v0 == gvar11 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.LookAtObject(obj132, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:FarmerWifeGreeting")
	goto LABEL3
LABEL2:
	ns.LookAtObject(obj132, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:FarmerWifeIdle")
	goto LABEL3
LABEL3:
	return
}
func FarmerWifeEnd() {
	var v0 int
	v0 = gvar71
	if v0 == gvar10 {
		goto LABEL1
	}
	if v0 == gvar11 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	gvar71 = gvar11
	LetterBoxOff()
	ns.Frozen(ns.GetHost(), false)
	goto LABEL3
LABEL2:
	goto LABEL3
LABEL3:
	return
}
func IngridStart() {
	var (
		v0 float32
		v1 float32
		v2 int
	)
	v2 = gvar76
	if v2 == gvar28 {
		goto LABEL1
	}
	if v2 == gvar29 {
		goto LABEL2
	}
	if v2 == gvar30 {
		goto LABEL3
	}
	if v2 == gvar31 {
		goto LABEL4
	}
	goto LABEL5
LABEL1:
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:IngridGreeting")
	goto LABEL5
LABEL2:
	if gvar84 != 1 {
		goto LABEL6
	}
	v0 = ns.GetWaypointX(wp204)
	v1 = ns.GetWaypointY(wp204)
	ns.AudioEvent(ns.CurePoisonCast, wp204)
	ns.AudioEvent(ns.CurePoisonEffect, wp204)
	ns.Effect(ns.BLUE_SPARKS, v0, v1, 0, 0)
	gvar76 = gvar30
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:IngridWaiting")
	goto LABEL5
	goto LABEL3
LABEL6:
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:IngridWaiting")
	goto LABEL5
LABEL3:
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:IngridReturned")
	goto LABEL5
LABEL4:
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:IngridIdle")
	goto LABEL5
LABEL5:
	return
}
func IngridEnd() {
	var v0 int
	v0 = gvar76
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
	gvar76 = gvar29
	ns.SetDialog(obj92, ns.NORMAL, IngridStart, IngridEnd)
	ns.StartDialog(obj93, ns.GetHost())
	goto LABEL5
LABEL2:
	goto LABEL5
LABEL3:
	gvar76 = gvar31
	goto LABEL5
LABEL4:
	goto LABEL5
LABEL5:
	return
}
func CaptainStart() {
	var v2 int
	v2 = gvar77
	if v2 == gvar32 {
		goto LABEL1
	}
	if v2 == gvar33 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:CaptainGreeting")
	goto LABEL3
LABEL2:
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:CaptainWaiting")
	goto LABEL3
LABEL3:
	return
}
func CaptainEnd() {
	var v0 int
	v0 = gvar77
	if v0 == gvar32 {
		goto LABEL1
	}
	if v0 == gvar33 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	gvar77 = gvar33
	ns.StartupScreen(5)
	Ingrid4()
	ns.FrameTimer(1, HorvathAppears4)
	goto LABEL3
LABEL2:
	goto LABEL3
LABEL3:
	return
}
func MaidenStart() {
	var (
		v0 float32
		v1 float32
		v2 int
	)
	_ = v1
	_ = v0
	v2 = gvar75
	if v2 == gvar24 {
		goto LABEL1
	}
	if v2 == gvar25 {
		goto LABEL2
	}
	if v2 == gvar26 {
		goto LABEL3
	}
	if v2 == gvar27 {
		goto LABEL4
	}
	goto LABEL5
LABEL1:
	if !ns.HasItem(ns.GetHost(), obj129) {
		goto LABEL6
	}
	gvar75 = gvar26
	ns.SetDialog(obj128, ns.NORMAL, MaidenStart, MaidenEnd)
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:MaidenReturned")
	goto LABEL5
	goto LABEL2
LABEL6:
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:MaidenGreeting")
	goto LABEL5
LABEL2:
	if !ns.HasItem(ns.GetHost(), obj129) {
		goto LABEL7
	}
	v0 = ns.GetWaypointX(wp209)
	v1 = ns.GetWaypointY(wp209)
	gvar75 = gvar26
	ns.SetDialog(obj128, ns.NORMAL, MaidenStart, MaidenEnd)
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:MaidenReturned")
	goto LABEL5
	goto LABEL3
LABEL7:
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:MaidenWaiting")
	goto LABEL5
LABEL3:
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:MaidenReturned")
	goto LABEL5
LABEL4:
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:MaidenIdle")
	goto LABEL5
LABEL5:
	return
}
func MaidenEnd() {
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
	v0 = ns.GetAnswer(obj128)
	v4 = gvar75
	if v4 == gvar24 {
		goto LABEL1
	}
	if v4 == gvar25 {
		goto LABEL2
	}
	if v4 == gvar26 {
		goto LABEL3
	}
	if v4 == gvar27 {
		goto LABEL4
	}
	goto LABEL5
LABEL1:
	if v0 != 1 {
		goto LABEL6
	}
	gvar75 = gvar25
	ns.SetDialog(obj128, ns.NORMAL, MaidenStart, MaidenEnd)
	ns.JournalEntry(ns.GetHost(), "War05Quest7", 2)
	ns.PrintToAll("Con01a:NewJournalEntry")
	goto LABEL5
LABEL6:
	if v0 != 2 {
		goto LABEL7
	}
	gvar75 = gvar24
	ns.SetDialog(obj128, ns.YESNO, MaidenStart, MaidenEnd)
	goto LABEL5
LABEL7:
	if v0 != 0 {
		goto LABEL2
	}
	gvar75 = gvar24
	ns.SetDialog(obj128, ns.YESNO, MaidenStart, MaidenEnd)
	goto LABEL5
LABEL2:
	ns.SetDialog(obj128, ns.NORMAL, MaidenStart, MaidenEnd)
	goto LABEL5
LABEL3:
	gvar75 = gvar27
	ns.PrintToAll("War05A.scr:ObjectiveComplete")
	ns.GiveXp(ns.GetHost(), 250)
	ns.JournalEdit(ns.GetHost(), "War05Quest7", 4)
	ns.Delete(obj129)
	ns.Pickup(ns.GetHost(), obj123)
	ns.SetDialog(obj128, ns.NORMAL, MaidenStart, MaidenEnd)
	goto LABEL5
LABEL4:
	ns.SetDialog(obj128, ns.NORMAL, MaidenStart, MaidenEnd)
	goto LABEL5
LABEL5:
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
	v8 = ns.GetWaypointX(wp191)
	v9 = ns.GetWaypointY(wp191)
	v0 = ns.GetWaypointX(wp212)
	v1 = ns.GetWaypointY(wp212)
	v2 = ns.GetWaypointX(wp210)
	v3 = ns.GetWaypointY(wp210)
	v6 = ns.GetWaypointX(wp211)
	v7 = ns.GetWaypointY(wp211)
	v4 = ns.GetObjectX(obj131)
	v5 = ns.GetObjectY(obj131)
	ns.Frozen(ns.GetHost(), true)
	LetterBoxOn()
	ns.AudioEvent(ns.CurePoisonCast, wp191)
	ns.AudioEvent(ns.CurePoisonEffect, wp191)
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
	v8 = ns.GetWaypointX(wp191)
	v9 = ns.GetWaypointY(wp191)
	v0 = ns.GetWaypointX(wp212)
	v1 = ns.GetWaypointY(wp212)
	v2 = ns.GetWaypointX(wp210)
	v3 = ns.GetWaypointY(wp210)
	v6 = ns.GetWaypointX(wp211)
	v7 = ns.GetWaypointY(wp211)
	v4 = ns.GetObjectX(obj131)
	v5 = ns.GetObjectY(obj131)
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
	v8 = ns.GetWaypointX(wp191)
	v9 = ns.GetWaypointY(wp191)
	v0 = ns.GetWaypointX(wp212)
	v1 = ns.GetWaypointY(wp212)
	v2 = ns.GetWaypointX(wp210)
	v3 = ns.GetWaypointY(wp210)
	v6 = ns.GetWaypointX(wp211)
	v7 = ns.GetWaypointY(wp211)
	v4 = ns.GetObjectX(obj131)
	v5 = ns.GetObjectY(obj131)
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
	v8 = ns.GetWaypointX(wp191)
	v9 = ns.GetWaypointY(wp191)
	v0 = ns.GetWaypointX(wp212)
	v1 = ns.GetWaypointY(wp212)
	v2 = ns.GetWaypointX(wp210)
	v3 = ns.GetWaypointY(wp210)
	v6 = ns.GetWaypointX(wp211)
	v7 = ns.GetWaypointY(wp211)
	v4 = ns.GetObjectX(obj131)
	v5 = ns.GetObjectY(obj131)
	ns.Effect(ns.BLUE_SPARKS, v4, v5, 0, 0)
	ns.MoveObject(obj131, v0, v1)
	ns.MoveObject(obj132, v4, v5)
	ns.CreatureGuard(obj132, v2, v3, v6, v7, 0)
	ns.CreatureGuard(obj130, v8, v9, v8, v9, 0)
}
func EndCurse5() {
	ns.LookAtObject(obj132, ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj132)
	ns.StartDialog(obj132, ns.GetHost())
}
func FrogStart() {
	var v0 int
	v0 = gvar73
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
	ns.LookAtObject(obj134, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:FrogGreeting")
	ns.Frozen(obj134, false)
	ns.ObjectOff(obj136)
	ns.ObjectOff(obj137)
	ns.ObjectOff(obj138)
	ns.ObjectOn(obj134)
	ns.CreatureFollow(obj134, ns.GetHost())
	goto LABEL5
LABEL2:
	ns.LookAtObject(obj134, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:FrogFollowing")
	goto LABEL5
LABEL3:
	ns.LookAtObject(obj134, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:FrogThanks")
	goto LABEL5
LABEL4:
	ns.LookAtObject(obj134, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:FrogIdle")
	goto LABEL5
LABEL5:
	return
}
func FrogEnd() {
	var v0 int
	v0 = gvar73
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
	gvar73 = gvar21
	goto LABEL5
LABEL2:
	goto LABEL5
LABEL3:
	gvar73 = gvar23
	goto LABEL5
LABEL4:
	goto LABEL5
LABEL5:
	return
}
func FrogGuard() {
	gvar73 = gvar22
	ns.Enchant(obj134, ns.ENCHANT_INVULNERABLE, 0)
	gvar86 = 1
}
func Villager1Start() {
	ns.Frozen(obj179, true)
	ns.CreatureIdle(obj179)
	ns.LookAtObject(obj179, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con05:Townsman9Talk01")
}
func Villager1End() {
	ns.Frozen(obj179, false)
	ns.SetRoamFlag(obj179, 32)
	ns.Wander(obj179)
}
func Villager2Start() {
	ns.Frozen(obj180, true)
	ns.CreatureIdle(obj180)
	ns.LookAtObject(obj180, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con05:Townsman10Talk01")
}
func Villager2End() {
	ns.Frozen(obj180, false)
	ns.SetRoamFlag(obj180, 32)
	ns.Wander(obj180)
}
func Villager3Start() {
	ns.Frozen(obj181, true)
	ns.CreatureIdle(obj181)
	ns.LookAtObject(obj181, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con05:Townsman11Talk01")
}
func Villager3End() {
	ns.Frozen(obj181, false)
	ns.SetRoamFlag(obj181, 32)
	ns.Wander(obj181)
}
func Villager4Start() {
	ns.Frozen(obj182, true)
	ns.CreatureIdle(obj182)
	ns.LookAtObject(obj182, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con05:Townsman11Talk02")
}
func Villager4End() {
	ns.Frozen(obj182, false)
	ns.SetRoamFlag(obj182, 32)
	ns.Wander(obj182)
}
func Villager5Start() {
	ns.Frozen(obj183, true)
	ns.CreatureIdle(obj183)
	ns.LookAtObject(obj183, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:VillagerGreeting")
}
func Villager5End() {
	ns.Frozen(obj183, false)
	ns.SetRoamFlag(obj183, 32)
	ns.Wander(obj183)
}
func SadVillagerStart() {
	var (
		v0 float32
		v1 float32
		v2 int
	)
	_ = v1
	_ = v0
	v2 = gvar74
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
	ns.CreatureIdle(obj133)
	ns.LookAtObject(obj133, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:SadVillagerGreeting")
	goto LABEL7
LABEL2:
	ns.CreatureIdle(obj133)
	ns.LookAtObject(obj133, ns.GetHost())
	if gvar86 != 1 {
		goto LABEL8
	}
	v0 = ns.GetWaypointX(wp206)
	v1 = ns.GetWaypointY(wp206)
	ns.AudioEvent(ns.CurePoisonCast, wp206)
	ns.AudioEvent(ns.CurePoisonEffect, wp206)
	gvar74 = gvar16
	ns.SetDialog(obj133, ns.NEXT, SadVillagerStart, SadVillagerEnd)
	ns.CreatureFollow(obj134, obj133)
	SadVillagerStart()
	goto LABEL3
LABEL8:
	ns.TellStory(ns.SwordsmanHurt, "War05A.scr:SadVillagerWaiting")
	goto LABEL7
LABEL3:
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
	_ = v0
	v0 = 0
	v1 = 1
	v2 = 2
	v3 = 0
	v0 = ns.GetAnswer(obj133)
	v4 = gvar74
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
	v0 = 1
	if 1 == 0 {
		goto LABEL8
	}
	gvar74 = gvar15
	ns.SetDialog(obj133, ns.NORMAL, SadVillagerStart, SadVillagerEnd)
	ns.JournalEntry(ns.GetHost(), "War05Quest9", 2)
	ns.PrintToAll("Con01a:NewJournalEntry")
	goto LABEL7
LABEL8:
	v0 = 2
	if 2 == 0 {
		goto LABEL9
	}
	gvar74 = gvar14
	ns.SetDialog(obj133, ns.YESNO, SadVillagerStart, SadVillagerEnd)
	goto LABEL7
LABEL9:
	v0 = 0
	if 0 == 0 {
		goto LABEL2
	}
	gvar74 = gvar14
	ns.SetDialog(obj133, ns.YESNO, SadVillagerStart, SadVillagerEnd)
	goto LABEL7
LABEL2:
	ns.SetDialog(obj133, ns.NORMAL, SadVillagerStart, SadVillagerEnd)
	goto LABEL7
LABEL3:
	gvar74 = gvar17
	ns.SetDialog(obj133, ns.NORMAL, SadVillagerStart, SadVillagerEnd)
	ns.PrintToAll("War05A.scr:ObjectiveComplete")
	ns.JournalEdit(ns.GetHost(), "War05Quest9", 4)
	ns.GiveXp(ns.GetHost(), 250)
	ns.Drop(obj133, obj124)
	ns.FrameTimer(1, FrogThanksPlayer)
	ns.FrameTimer(1, SadVillagerGuard)
	goto LABEL7
LABEL4:
	ns.SetDialog(obj133, ns.NORMAL, SadVillagerStart, SadVillagerEnd)
	goto LABEL7
LABEL5:
	gvar74 = gvar19
	ns.Move(obj133, wp206)
	ns.SetCallback(obj133, 9, DummyBump)
	ns.JournalDelete(ns.GetHost(), "War05Quest9")
	ns.PrintToAll("War05A.scr:ObjectiveFailed")
	ns.SetDialog(obj133, ns.NORMAL, SadVillagerStart, SadVillagerEnd)
	goto LABEL7
LABEL6:
	ns.CancelDialog(obj133)
	goto LABEL7
LABEL7:
	return
}
func OwnVillagers() {
	ns.FrameTimer(1, OwnVillagers2)
}
func OwnVillagers2() {
	ns.SetOwner(ns.GetHost(), obj134)
	ns.SetOwner(ns.GetHost(), obj93)
	ns.SetOwner(ns.GetHost(), obj92)
	ns.Frozen(obj179, true)
	ns.Frozen(obj180, true)
	ns.Frozen(obj181, true)
	ns.Frozen(obj182, true)
	ns.Frozen(obj183, true)
	ns.Frozen(obj184, true)
	ns.Frozen(obj185, true)
	ns.Frozen(obj186, true)
	ns.Frozen(obj187, true)
	ns.Frozen(obj50, true)
	ns.Frozen(obj51, true)
	ns.Frozen(obj130, true)
	ns.Frozen(obj48, true)
	ns.Frozen(obj49, true)
	ns.Frozen(obj54, true)
	ns.Frozen(obj55, true)
	ns.RetreatLevel(obj47, 0)
	ns.RetreatLevel(obj48, 0)
	ns.RetreatLevel(obj49, 0)
	ns.RetreatLevel(obj50, 0)
	ns.RetreatLevel(obj51, 0)
}
func KillCorpses2() {
	ns.Damage(obj148, 0, 300, 5)
	ns.Damage(obj149, 0, 300, 5)
	ns.Damage(obj150, 0, 300, 5)
	ns.Damage(obj151, 0, 300, 5)
	ns.Damage(obj152, 0, 300, 5)
	ns.Damage(obj153, 0, 300, 5)
	ns.Damage(obj154, 0, 300, 5)
	ns.Damage(obj155, 0, 300, 5)
	ns.Damage(obj156, 0, 300, 5)
	ns.Damage(obj157, 0, 300, 5)
	ns.Damage(obj158, 0, 300, 5)
	ns.Damage(obj159, 0, 300, 5)
	ns.Damage(obj160, 0, 300, 5)
	ns.Damage(obj161, 0, 300, 5)
	ns.Damage(obj162, 0, 300, 5)
	ns.Damage(obj163, 0, 300, 0)
	ns.Damage(obj164, 0, 300, 0)
	ns.Damage(obj165, 0, 300, 0)
	ns.Damage(obj143, 0, 300, 0)
	ns.Damage(obj45, 0, 300, 0)
	ns.Damage(obj46, 0, 300, 0)
	ns.Damage(obj60, 0, 300, 0)
	ns.Damage(obj61, 0, 300, 0)
	ns.Damage(obj62, 0, 300, 0)
	ns.Damage(obj63, 0, 300, 0)
	ns.Damage(obj64, 0, 300, 0)
	ns.Damage(obj65, 0, 300, 0)
}
func VillagePath() {
	ns.SetRoamFlag(obj179, 32)
	ns.SetRoamFlag(obj180, 32)
	ns.SetRoamFlag(obj181, 32)
	ns.SetRoamFlag(obj182, 32)
	ns.SetRoamFlag(obj183, 32)
	ns.Wander(obj179)
	ns.Wander(obj180)
	ns.Wander(obj181)
	ns.Wander(obj182)
	ns.Wander(obj183)
}
func PickUpObjects() {
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
	ns.AggressionLevel(obj144, 0.83)
	ns.AggressionLevel(obj56, 0.83)
	ns.AggressionLevel(obj47, 0.83)
	ns.Attack(obj47, obj144)
}
func SwordsmanTriggerB() {
	gvar72 = gvar12
	ns.SetDialog(obj144, ns.NORMAL, SwordsmanStart, SwordsmanEnd)
	ns.StoryPic(obj144, "Warrior8Pic")
	ns.FrameTimer(1, SwordsmanTriggerC)
}
func SwordsmanTriggerC() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
	)
	v0 = ns.GetWaypointX(wp189)
	v1 = ns.GetWaypointY(wp189)
	v2 = ns.GetWaypointX(wp190)
	v3 = ns.GetWaypointY(wp190)
	ns.CreatureGuard(obj144, v0, v1, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 240)
	ns.CreatureGuard(obj56, v2, v3, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 240)
}
func SwordsmanTriggerD() {
}
func ShopKeeperGuard() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
	)
	v0 = ns.GetWaypointX(wp225)
	v1 = ns.GetWaypointY(wp225)
	v2 = ns.GetWaypointX(wp226)
	v3 = ns.GetWaypointY(wp226)
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
	v0 = ns.GetWaypointX(wp227)
	v1 = ns.GetWaypointY(wp227)
	v2 = ns.GetWaypointX(wp228)
	v3 = ns.GetWaypointY(wp228)
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
	v0 = ns.GetWaypointX(wp223)
	v1 = ns.GetWaypointY(wp223)
	v2 = ns.GetWaypointX(wp224)
	v3 = ns.GetWaypointY(wp224)
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
	v0 = ns.GetWaypointX(wp206)
	v1 = ns.GetWaypointY(wp206)
	v2 = ns.GetWaypointX(wp207)
	v3 = ns.GetWaypointY(wp207)
	ns.CreatureGuard(obj133, v0, v1, v2, v3, 0)
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
func PlayerDeath() {
	ns.DeathScreen(5)
}
func MapEntry() {
	ns.Music(23, 100)
	ns.Frozen(ns.GetHost(), false)
	LetterBoxOff()
	if ns.GetQuestStatus("War05B:ChicksSaved") != 1 {
		goto LABEL1
	}
	ns.Delete(obj93)
	ns.Delete(obj92)
LABEL1:
	flag87 = true
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
	ns.StartDialog(obj134, ns.GetHost())
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
	ns.Move(obj93, wp213)
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
	v0 = ns.GetWaypointX(wp215)
	v1 = ns.GetWaypointY(wp215)
	v2 = ns.GetWaypointX(wp213)
	v3 = ns.GetWaypointY(wp213)
	ns.Frozen(obj93, false)
	ns.CreatureGuard(obj93, v0, v1, v2, v3, 0)
	LetterBoxOff()
	ns.FrameTimer(45, HorvathAppears5)
}
func HorvathAppears5() {
	ns.Frozen(ns.GetHost(), false)
}
func Ingrid1() {
	ns.ObjectOff(obj120)
	ns.ObjectOff(obj121)
	ns.ObjectOff(obj122)
	ns.Frozen(ns.GetHost(), true)
	ns.CreatureIdle(ns.GetHost())
	LetterBoxOn()
	ns.Move(obj92, wp204)
	ns.Move(obj93, wp213)
}
func Ingrid2() {
	ns.Chat(obj92, "War05A.scr:IngridScream")
}
func Ingrid3() {
	if !ns.IsCaller(obj92) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.Frozen(obj92, true)
	ns.CreatureIdle(obj92)
	ns.LookAtObject(ns.GetHost(), obj92)
	ns.LookAtObject(obj92, ns.GetHost())
	ns.Frozen(obj93, true)
	ns.CreatureIdle(obj93)
	ns.LookAtObject(obj93, ns.GetHost())
	ns.DestroyChat(obj92)
	ns.StartDialog(obj92, ns.GetHost())
LABEL1:
	return
}
func Ingrid4() {
	ns.Frozen(obj92, false)
	ns.Move(obj92, wp205)
	ns.FrameTimer(30, Ingrid5)
}
func Ingrid5() {
	ns.JournalEntry(ns.GetHost(), "War05Quest4", 2)
	ns.PrintToAll("Con01a:NewJournalEntry")
	ns.Frozen(ns.GetHost(), false)
}
func Ingrid6() {
	ns.CreatureGuard(obj92, ns.GetObjectX(obj92), ns.GetObjectY(obj92), ns.GetObjectX(ns.GetHost()), ns.GetObjectX(ns.GetHost()), 0)
}
func MapInitialize() {
	wp34 = ns.Waypoint("WoundedRefugeeWayPoint")
	wp35 = ns.Waypoint("WoundedRefugeeWayPoint2")
	wp36 = ns.Waypoint("DeadGuyWP")
	wp39 = ns.Waypoint("PlayerSounds")
	wp37 = ns.Waypoint("LunchOgre1WP")
	wp38 = ns.Waypoint("LunchOgre2WP")
	wp40 = ns.Waypoint("WellWP")
	wp41 = ns.Waypoint("SecretBearWP")
	obj42 = ns.Object("WoundedRefugee")
	obj43 = ns.Object("WoundedRefugee2")
	obj44 = ns.Object("Antidote")
	obj45 = ns.Object("Kill")
	obj46 = ns.Object("Kill2")
	obj47 = ns.Object("FirstOgre")
	obj48 = ns.Object("LunchOgre1")
	obj49 = ns.Object("LunchOgre2")
	obj50 = ns.Object("TownOgre1")
	obj51 = ns.Object("TownOgre2")
	obj55 = ns.Object("DeadGuy")
	obj56 = ns.Object("Swordsman2")
	obj57 = ns.Object("Archer1")
	obj58 = ns.Object("Archer2")
	obj54 = ns.Object("BruteBoss")
	obj52 = ns.Object("MoreOgre1")
	obj53 = ns.Object("MoreOgre2")
	obj59 = ns.Object("SecretBear")
	obj60 = ns.Object("Bubbles1")
	obj61 = ns.Object("Bubbles2")
	obj62 = ns.Object("Bubbles3")
	obj63 = ns.Object("Bubbles4")
	obj64 = ns.Object("Bubbles5")
	obj65 = ns.Object("Bubbles6")
	obj66 = ns.Object("TestFlame")
	gvar67 = ns.ObjectGroup("OgreAttackTriggers")
	obj93 = ns.Object("W5Captain")
	obj92 = ns.Object("W5Ingrid")
	obj95 = ns.Object("W5Bartender")
	obj96 = ns.Object("W5ShopKeeper")
	obj94 = ns.Object("W5Gypsy")
	obj120 = ns.Object("CaptainTrigger1")
	obj121 = ns.Object("CaptainTrigger2")
	obj122 = ns.Object("CaptainTrigger3")
	obj97 = ns.Object("BartenderTrigger1")
	obj98 = ns.Object("BartenderTrigger2")
	obj99 = ns.Object("ShopKeeperTrigger1")
	obj100 = ns.Object("ShopKeeperTrigger2")
	obj101 = ns.Object("GypsyTrigger1")
	obj102 = ns.Object("GypsyTrigger2")
	obj103 = ns.Object("GypsyTrigger3")
	obj104 = ns.Object("GypsyTrigger4")
	obj105 = ns.Object("GypsyTrigger5")
	obj106 = ns.Object("GypsyTrigger6")
	obj107 = ns.Object("GypsyTrigger7")
	obj108 = ns.Object("GypsyTrigger8")
	obj109 = ns.Object("GypsyTrigger9")
	obj110 = ns.Object("GypsyEndTrigger1")
	obj111 = ns.Object("GypsyEndTrigger2")
	obj112 = ns.Object("GypsyEndTrigger3")
	obj113 = ns.Object("GypsyEndTrigger4")
	obj114 = ns.Object("GypsyEndTrigger5")
	obj115 = ns.Object("GypsyEndTrigger6")
	obj116 = ns.Object("GypsyEndTrigger7")
	obj117 = ns.Object("GypsyEndTrigger8")
	obj118 = ns.Object("GypsyEndTrigger9")
	obj144 = ns.Object("W5Swordsman")
	obj145 = ns.Object("W5Swordsman2")
	wp146 = int(ns.Object("SwordsmanTrigger1"))
	wp147 = int(ns.Object("SwordsmanTrigger2"))
	obj143 = ns.Object("KeyGuy")
	obj140 = ns.Object("W5Drunk1")
	obj141 = ns.Object("W5Drunk2")
	obj142 = ns.Object("W5Drunk3")
	obj123 = ns.Object("MaidenReward")
	obj124 = ns.Object("SadVillagerReward")
	obj125 = ns.Object("FarmerReward1")
	obj126 = ns.Object("FarmerReward2")
	obj130 = ns.Object("W5Farmer")
	obj131 = ns.Object("W5Hound")
	obj132 = ns.Object("W5FarmerWife")
	obj133 = ns.Object("W5SadVillager")
	obj134 = ns.Object("W5Frog")
	obj135 = ns.Object("FrogTrigger")
	obj136 = ns.Object("FrogTrigger1")
	obj137 = ns.Object("FrogTrigger2")
	obj138 = ns.Object("FrogTrigger3")
	obj139 = ns.Object("FrogTrigger4")
	obj127 = ns.Object("FarmerStaff")
	obj88 = ns.Object("AutoSave1Trigger")
	obj89 = ns.Object("AutoSave2Trigger")
	obj90 = ns.Object("AutoSave3Trigger")
	obj91 = ns.Object("AutoSave4Trigger")
	obj148 = ns.Object("Corpse1")
	obj149 = ns.Object("Corpse2")
	obj150 = ns.Object("Corpse3")
	obj151 = ns.Object("Corpse4")
	obj152 = ns.Object("Corpse5")
	obj153 = ns.Object("Corpse6")
	obj154 = ns.Object("Corpse7")
	obj155 = ns.Object("Corpse8")
	obj156 = ns.Object("Corpse9")
	obj157 = ns.Object("Corpse10")
	obj158 = ns.Object("Corpse11")
	obj159 = ns.Object("Corpse12")
	obj160 = ns.Object("Corpse13")
	obj161 = ns.Object("Corpse14")
	obj162 = ns.Object("Corpse15")
	obj119 = ns.Object("JackLine5Trigger")
	obj163 = ns.Object("OgreCorpse1")
	obj164 = ns.Object("OgreCorpse2")
	obj165 = ns.Object("OgreCorpse3")
	obj128 = ns.Object("W5Maiden")
	obj129 = ns.Object("MaidenObject")
	obj123 = ns.Object("MaidenReward")
	obj166 = ns.Object("Ogre1")
	obj167 = ns.Object("Ogre2")
	obj168 = ns.Object("Ogre3")
	obj169 = ns.Object("Ogre4")
	obj170 = ns.Object("Ogre5")
	obj171 = ns.Object("Ogre6")
	obj172 = ns.Object("Ogre7")
	obj173 = ns.Object("Ogre8")
	obj174 = ns.Object("Ogre9")
	obj174 = ns.Object("Ogre10")
	obj174 = ns.Object("Ogre11")
	obj174 = ns.Object("Ogre12")
	obj174 = ns.Object("Ogre13")
	obj196 = ns.Object("HouseDoor1")
	obj197 = ns.Object("HouseDoor1")
	obj198 = ns.Object("CryptDoor1")
	obj199 = ns.Object("CryptDoor2")
	obj179 = ns.Object("W5Villager1")
	obj180 = ns.Object("W5Villager2")
	obj181 = ns.Object("W5Villager3")
	obj182 = ns.Object("W5Villager4")
	obj183 = ns.Object("W5Villager5")
	obj184 = ns.Object("W5Villager6")
	obj185 = ns.Object("W5Villager7")
	obj186 = ns.Object("W5Villager8")
	obj187 = ns.Object("W5Villager9")
	wp190 = ns.Waypoint("Swordsman2WayPoint")
	wp189 = ns.Waypoint("SwordsmanWayPoint")
	wp146 = int(ns.Waypoint("SwordsmanTrigger1"))
	wp147 = int(ns.Waypoint("SwordsmanTrigger2"))
	wp204 = ns.Waypoint("IngridWayPoint")
	wp205 = ns.Waypoint("IngridWayPoint2")
	wp213 = ns.Waypoint("HorvathWayPoint")
	wp214 = ns.Waypoint("HorvathWayPoint2")
	wp215 = ns.Waypoint("HorvathAppears")
	wp209 = ns.Waypoint("MaidenWayPoint")
	wp191 = ns.Waypoint("FarmerWayPoint")
	wp192 = ns.Waypoint("Farmer2WayPoint")
	wp210 = ns.Waypoint("FarmerWifeWayPoint")
	wp211 = ns.Waypoint("FarmerWife2WayPoint")
	wp193 = ns.Waypoint("HoundWayPoint")
	wp212 = ns.Waypoint("Hound2WayPoint")
	wp208 = ns.Waypoint("WoundedVillagerWayPoint")
	wp222 = ns.Waypoint("MaidenRewardWayPoint")
	wp206 = ns.Waypoint("SadVillagerWayPoint")
	wp207 = ns.Waypoint("SadVillagerWayPoint2")
	wp194 = ns.Waypoint("FrogWayPoint1")
	wp195 = ns.Waypoint("FrogWayPoint2")
	wp216 = ns.Waypoint("Villager1Goal")
	wp217 = ns.Waypoint("Villager2Goal")
	wp218 = ns.Waypoint("Villager3Goal")
	wp219 = ns.Waypoint("Villager4Goal")
	wp220 = ns.Waypoint("Villager5Goal")
	wp223 = ns.Waypoint("GypsyWayPoint")
	wp224 = ns.Waypoint("GypsyWayPoint2")
	wp225 = ns.Waypoint("ShopKeeperWayPoint")
	wp226 = ns.Waypoint("ShopKeeperWayPoint2")
	wp227 = ns.Waypoint("BartenderWayPoint")
	wp228 = ns.Waypoint("BartenderWayPoint2")
	wp221 = ns.Waypoint("PlayerStartWayPoint")
	ns.SetDialog(obj140, ns.NORMAL, DrunkStart, DrunkEnd)
	ns.StoryPic(obj140, "Townsman3Pic")
	ns.SetDialog(obj141, ns.NORMAL, DrunkStart, DrunkEnd)
	ns.StoryPic(obj141, "Townsman2Pic")
	ns.SetDialog(obj142, ns.NORMAL, DrunkStart, DrunkEnd)
	ns.StoryPic(obj142, "Townsman1Pic")
	ns.StoryPic(obj130, "ThaviusPic")
	ns.SetDialog(obj131, ns.NORMAL, HoundStart, HoundEnd)
	ns.StoryPic(obj131, "WolfPic")
	ns.SetDialog(obj133, ns.YESNO, SadVillagerStart, SadVillagerEnd)
	ns.StoryPic(obj133, "Townsman3Pic")
	ns.SetDialog(obj134, ns.NORMAL, FrogStart, FrogEnd)
	ns.StoryPic(obj134, "FrogPic")
	ns.SetDialog(obj92, ns.NEXT, IngridStart, IngridEnd)
	ns.StoryPic(obj92, "IngridPic")
	ns.SetDialog(obj93, ns.NORMAL, CaptainStart, CaptainEnd)
	ns.StoryPic(obj93, "AirshipCaptainPic")
	ns.SetDialog(obj132, ns.NORMAL, FarmerWifeStart, FarmerWifeEnd)
	ns.StoryPic(obj132, "MaidenPic")
	ns.SetDialog(obj128, ns.YESNO, MaidenStart, MaidenEnd)
	ns.StoryPic(obj128, "MaidenPic")
	ns.Frozen(obj134, true)
	DisableVillagers()
	KillCorpses()
	LockDoors()
	OwnVillagers()
	HoundEnd()
	ShopKeeperGuard()
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
	if !ns.IsCaller(obj134) {
		goto LABEL1
	}
	v0 = ns.GetWaypointX(wp194)
	v1 = ns.GetWaypointY(wp194)
	v2 = ns.GetWaypointX(wp195)
	v3 = ns.GetWaypointY(wp195)
	ns.ObjectOff(obj135)
	ns.CreatureGuard(obj134, v0, v1, v2, v3, 0)
	gvar74 = gvar15
LABEL1:
	return
}
func FrogDeath() {
	gvar86 = 2
	gvar74 = gvar18
	ns.CancelDialog(obj134)
}
func AutoSave1() {
	ns.AutoSave()
	ns.ObjectOff(obj88)
}
func AutoSave2() {
	ns.AutoSave()
	ns.ObjectOff(obj89)
}
func AutoSave3() {
	ns.AutoSave()
	ns.FrameTimer(1, VillagePath)
	ns.ObjectOff(obj90)
}
func AutoSave4() {
	ns.AutoSave()
	ns.ObjectOff(obj91)
}
func SecretFound() {
	ns.PrintToAll("War05A.scr:Secret")
	ns.MoveWaypoint(wp39, ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger()))
	ns.AudioEvent(ns.SecretFound, wp39)
	ns.GiveXp(ns.GetHost(), 250)
}
func Hint() {
	ns.PrintToAll("Journal:War05Hint1")
	ns.JournalEntry(ns.GetHost(), "War05Hint1", 4)
}
func CrapFile() {
}
func OgreAttackStart() {
	ns.ObjectGroupOff(gvar67)
	UnfreezeAll()
	ns.Music(26, 100)
	ns.ObjectOn(obj52)
	ns.ObjectOn(obj53)
	ns.Frozen(obj48, false)
	ns.Frozen(obj49, false)
	ns.Frozen(obj55, false)
	ns.Attack(obj55, obj48)
	ns.Attack(obj48, obj55)
	ns.Attack(obj49, obj55)
}
func DeadGuyArrive() {
	if !ns.IsCaller(obj55) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.Frozen(obj55, true)
	ns.Frozen(obj48, true)
	ns.Frozen(obj49, true)
	ns.CreatureIdle(obj55)
	ns.CreatureIdle(obj48)
	ns.CreatureIdle(obj49)
	ns.LookAtObject(obj48, obj55)
	ns.LookAtObject(obj55, obj49)
	ns.LookAtObject(obj49, obj55)
	ns.SecondTimer(2, OgresAttackDeadGuy)
LABEL1:
	return
}
func OgresAttackDeadGuy() {
	ns.Frozen(obj55, false)
	ns.Frozen(obj48, false)
	ns.Frozen(obj49, false)
	ns.Frozen(obj54, false)
	ns.Attack(obj55, obj49)
	ns.Attack(obj48, obj55)
	ns.Attack(obj49, obj55)
}
func DeadGuyDie() {
	ns.AggressionLevel(obj48, 0.83)
	ns.AggressionLevel(obj49, 0.83)
}
func LunchOgreReport() {
	if !(ns.CurrentHealth(obj55) > 0) {
		goto LABEL1
	}
	ns.Attack(ns.GetTrigger(), obj55)
LABEL1:
	return
}
func LunchOgre1Trigger() {
	if !ns.IsCaller(obj48) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ivar229 += 1
	ns.AggressionLevel(obj48, 0.83)
LABEL1:
	return
}
func LunchOgre2Trigger() {
	if !ns.IsCaller(obj49) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ivar229 += 1
	ns.AggressionLevel(obj49, 0.83)
LABEL1:
	return
}
func UnfreezeAll() {
	ns.Frozen(ns.GetHost(), false)
	ns.Frozen(obj179, false)
	ns.Frozen(obj180, false)
	ns.Frozen(obj181, false)
	ns.Frozen(obj182, false)
	ns.Frozen(obj183, false)
	ns.Frozen(obj184, false)
	ns.Frozen(obj185, false)
	ns.Frozen(obj186, false)
	ns.Frozen(obj187, false)
	ns.Frozen(obj50, false)
	ns.Frozen(obj51, false)
	ns.Frozen(obj130, false)
	ns.Frozen(obj54, false)
	ns.Wander(obj130)
	ns.CreatureFollow(obj51, obj130)
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
	ns.AudioEvent(ns.RestoreHealthName, wp40)
}
func KillWoundedRefugee2() {
	ns.Damage(obj42, 0, 15, 5)
}
func WoundedRefugeeStart() {
	var (
		v0 float32
		v1 float32
		v2 int
	)
	v2 = gvar234
	if v2 == gvar230 {
		goto LABEL1
	}
	if v2 == gvar231 {
		goto LABEL2
	}
	if v2 == gvar232 {
		goto LABEL3
	}
	if v2 == gvar233 {
		goto LABEL4
	}
	goto LABEL5
LABEL1:
	ns.TellStory(ns.SwordsmanHurt, "War05A:WoundedRefugeeGreeting")
	goto LABEL5
LABEL2:
	if !ns.HasItem(ns.GetHost(), obj44) {
		goto LABEL6
	}
	v0 = ns.GetWaypointX(wp34)
	v1 = ns.GetWaypointY(wp34)
	ns.AudioEvent(ns.CurePoisonCast, wp34)
	ns.AudioEvent(ns.CurePoisonEffect, wp34)
	ns.Delete(obj42)
	obj42 = obj43
	ns.MoveObject(obj42, v0, v1)
	gvar234 = gvar232
	WoundedRefugeeStart()
	goto LABEL3
LABEL6:
	ns.TellStory(ns.SwordsmanHurt, "War05A:WoundedRefugeeWaiting")
	goto LABEL5
LABEL3:
	ns.TellStory(ns.SwordsmanHurt, "War05A:WoundedRefugeeHealed")
	gvar234 = gvar232
	goto LABEL5
LABEL4:
	ns.TellStory(ns.SwordsmanHurt, "War05A:WoundedRefugeeIdle")
	goto LABEL5
LABEL5:
	return
}
func WoundedRefugeeEnd() {
	var v0 int
	v0 = gvar234
	if v0 == gvar230 {
		goto LABEL1
	}
	if v0 == gvar231 {
		goto LABEL2
	}
	if v0 == gvar232 {
		goto LABEL3
	}
	if v0 == gvar233 {
		goto LABEL4
	}
	goto LABEL5
LABEL1:
	gvar234 = gvar231
	goto LABEL5
LABEL2:
	goto LABEL5
LABEL3:
	gvar234 = gvar233
	ns.Move(obj42, wp35)
	goto LABEL5
LABEL4:
	goto LABEL5
LABEL5:
	return
}
func KillWoundedRefugee() {
	ns.FrameTimer(30, KillWoundedRefugee2)
}
func OnEvent(typ string) {
	switch typ {
	case "PlayerDeath":
		PlayerDeath()
	case "MapEntry":
		MapEntry()
	case "MapInitialize":
		MapInitialize()
	}
}
