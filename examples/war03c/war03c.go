package war03c

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	obj4   ns.ObjectID
	obj5   ns.ObjectID
	obj6   ns.ObjectID
	obj7   ns.ObjectID
	gvar8  ns.ObjectGroupID
	obj9   ns.ObjectID
	gvar10 ns.WallGroupID
	gvar11 ns.ObjectGroupID
	wp12   ns.WaypointID
	obj13  ns.ObjectID
	obj14  ns.ObjectID
	obj15  ns.ObjectID
	wp16   ns.WaypointID
	wp17   ns.WaypointID
	wp18   ns.WaypointID
	obj19  ns.ObjectID
	gvar20 ns.ObjectGroupID
	gvar21 ns.ObjectGroupID
	obj22  ns.ObjectID
	obj23  ns.ObjectID
	obj24  ns.ObjectID
	obj25  ns.ObjectID
	obj26  ns.ObjectID
	obj27  ns.ObjectID
	obj28  ns.ObjectID
	obj29  ns.ObjectID
	obj30  ns.ObjectID
	wp31   [6]ns.WaypointID
	obj32  ns.ObjectID
	wp33   ns.WaypointID
	wp34   ns.WaypointID
	wp35   ns.WaypointID
	wp36   ns.WaypointID
	wp37   ns.WaypointID
	wp38   ns.WaypointID
	wp39   ns.WaypointID
	gvar40 int
	gvar41 int
	gvar42 int
	gvar43 int
	gvar44 int
	gvar45 int
	fvar46 float32
	obj47  ns.ObjectID
	obj48  ns.ObjectID
	obj49  ns.ObjectID
	obj50  ns.ObjectID
	obj51  ns.ObjectID
	obj52  ns.ObjectID
	obj53  ns.ObjectID
	obj54  ns.ObjectID
	gvar55 ns.ObjectGroupID
	flag56 bool
	flag57 bool
	flag58 bool
	flag59 bool
	wp60   ns.WaypointID
	wp61   ns.WaypointID
	wp62   ns.WaypointID
	wp63   ns.WaypointID
	wp64   ns.WaypointID
	wp65   ns.WaypointID
	wp66   ns.WaypointID
	wp67   ns.WaypointID
	wp68   ns.WaypointID
	wp69   ns.WaypointID
)

func init() {
	gvar40 = 0
	gvar41 = 1
	gvar42 = 2
	gvar43 = 3
	gvar44 = 4
	gvar45 = gvar40
	fvar46 = 50
	flag56 = false
	flag57 = false
	flag58 = false
	flag59 = false
}
func PlayerDeath() {
	ns.DeathScreen(3)
}
func KalenTalkStart() {
	ns.TellStory(ns.SwordsmanHurt, "War03c:KalenSpeech")
	ns.ObjectGroupOff(gvar20)
}
func KalenTalkEnd() {
	var (
		v0 float32
		v1 float32
	)
	v0 = ns.GetWaypointX(wp16)
	v1 = ns.GetWaypointY(wp16)
	ns.AudioEvent(ns.NPCDie, wp16)
	ns.AudioEvent(ns.MetalArmorDrop, wp17)
	ns.AudioEvent(ns.LeatherArmorDrop, wp18)
	ns.MoveObject(obj15, v0, v1)
	ns.MoveObject(obj13, ns.GetWaypointX(wp17), ns.GetWaypointY(wp17))
	ns.MoveObject(obj14, ns.GetWaypointX(wp18), ns.GetWaypointY(wp18))
	ns.Damage(obj15, 0, 200, 12)
	ns.ObjectGroupOn(gvar20)
	ZombieAttack()
	ns.Damage(obj19, 0, 200, 12)
}
func ZombieAttack() {
	ns.WallGroupBreak(gvar10)
	ns.GroupAggressionLevel(gvar11, 0.83)
	ns.ObjectGroupOn(gvar21)
	ns.MoveObject(obj22, 1848, 2581)
	ns.MoveObject(obj23, 1848, 2581)
	ns.MoveObject(obj24, 1904, 2710)
	ns.MoveObject(obj25, 1904, 2710)
	ns.MoveObject(obj26, 1988, 2624)
	ns.MoveObject(obj27, 1988, 2624)
	ns.MoveObject(obj28, 1976, 2534)
	ns.MoveObject(obj29, 1793, 2708)
}
func UrchinGoHome() {
	ns.GroupWander(gvar8)
}
func InitializeUrchinSetpiece() {
	obj47 = ns.Object("Urchin01")
	obj48 = ns.Object("Urchin02")
	obj49 = ns.Object("Urchin03")
	obj50 = ns.Object("Urchin04")
	obj51 = ns.Object("ShamanDoor1")
	obj52 = ns.Object("ShamanDoor2")
	gvar55 = ns.ObjectGroup("OtherUrchins")
	wp60 = ns.Waypoint("Urchin01SetpieceWP")
	wp61 = ns.Waypoint("Urchin02SetpieceWP")
	wp62 = ns.Waypoint("Urchin03SetpieceWP")
	wp63 = ns.Waypoint("Urchin04SetpieceWP")
	wp64 = ns.Waypoint("ShamanSetpieceWP")
	wp65 = ns.Waypoint("UrchinSetpieceAudioOrigin")
	wp66 = ns.Waypoint("ShamanExitWP")
	wp67 = ns.Waypoint("ShamanFaceWP")
	obj53 = ns.Object("ShamanBuddy1")
	obj54 = ns.Object("ShamanBuddy2")
	wp68 = ns.Waypoint("ShamanBuddy1WP")
	wp69 = ns.Waypoint("ShamanBuddy2WP")
}
func PlaySubMusic() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.Music(18, 100)
LABEL1:
	return
}
func MapInitialize() {
	gvar8 = ns.ObjectGroup("UrchinEntryGroup")
	obj4 = ns.Object("ChamberDoor1")
	obj5 = ns.Object("ChamberDoor2")
	obj6 = ns.Object("ChamberDoor3")
	obj7 = ns.Object("ZombieJailDoor")
	ns.LockDoor(obj7)
	ns.LockDoor(obj4)
	ns.LockDoor(obj5)
	ns.LockDoor(obj6)
	obj9 = ns.Object("MayorScepter")
	obj13 = ns.Object("KalenArmor")
	obj14 = ns.Object("KalenBoots")
	wp17 = ns.Waypoint("KalenArmorWP")
	wp18 = ns.Waypoint("KalenBootsWP")
	obj15 = ns.Object("KalenDead")
	ns.Damage(obj15, 0, 200, 12)
	obj19 = ns.Object("Kalen")
	gvar20 = ns.ObjectGroup("KalenMonsterGroup")
	wp16 = ns.Waypoint("KalenWP")
	ns.SetOwner(ns.GetHost(), obj19)
	ns.StoryPic(obj19, "WoundedWarriorPic")
	ns.SetDialog(obj19, ns.NORMAL, KalenTalkStart, KalenTalkEnd)
	gvar10 = ns.WallGroup("ZombieWallGroup")
	wp12 = ns.Waypoint("ZombieSound")
	gvar11 = ns.ObjectGroup("ZombieAmbushGroup")
	gvar21 = ns.ObjectGroup("MonsterGroup1")
	obj22 = ns.Object("Monster1")
	obj23 = ns.Object("Monster2")
	obj24 = ns.Object("Monster3")
	obj25 = ns.Object("Monster4")
	obj26 = ns.Object("Monster5")
	obj27 = ns.Object("Monster6")
	obj26 = ns.Object("Monster7")
	obj27 = ns.Object("Monster8")
	obj30 = ns.Object("Shaman")
	wp31[0] = ns.Waypoint("ShamanWP1")
	wp31[1] = ns.Waypoint("ShamanWP2")
	wp31[2] = ns.Waypoint("ShamanWP3")
	wp31[3] = ns.Waypoint("ShamanWP4")
	wp31[4] = ns.Waypoint("ShamanWP5")
	wp31[5] = ns.Waypoint("ShamanWP6")
	obj32 = ns.Object("ShamanElevator")
	InitializeUrchinSetpiece()
	wp33 = ns.Waypoint("Secret01WP")
	wp34 = ns.Waypoint("Secret02WP")
	wp35 = ns.Waypoint("Secret03WP")
	wp36 = ns.Waypoint("Secret04WP")
	wp37 = ns.Waypoint("Secret05WP")
	wp38 = ns.Waypoint("Secret06WP")
	wp39 = ns.Waypoint("QuestAudioWP")
}
func MapEntry() {
	ns.UnBlind()
}
func ShamanMove() {
	var (
		v0 int
		v1 float32
		v2 float32
	)
	v0 = ns.Random(0, 5)
	v1 = ns.GetWaypointX(wp31[v0])
	v2 = ns.GetWaypointY(wp31[v0])
	ns.MoveObject(obj30, v1, v2)
	ns.LookAtObject(obj30, obj32)
}
func UnlockChamberDoors() {
	ns.UnlockDoor(obj4)
	ns.UnlockDoor(obj5)
	ns.UnlockDoor(obj6)
}
func MonsterGoHomeLobo() {
	r1 := ns.IsAttackedBy(ns.GetCaller(), ns.GetHost())
	if !r1 {
		goto LABEL1
	}
	ns.AggressionLevel(ns.GetCaller(), 0)
	ns.GoBackHome(ns.GetCaller())
LABEL1:
	return
}
func MonsterGoHome() {
	PlaySubMusic()
	r1 := ns.IsAttackedBy(ns.GetCaller(), ns.GetHost())
	if !r1 {
		goto LABEL1
	}
	ns.GoBackHome(ns.GetCaller())
LABEL1:
	return
}
func MonsterHostile() {
	PlaySubMusic()
	r1 := ns.IsAttackedBy(ns.GetCaller(), ns.GetHost())
	if !r1 {
		goto LABEL1
	}
	ns.AggressionLevel(ns.GetCaller(), 0.83)
LABEL1:
	return
}
func MonsterGoHomePassive() {
	r1 := ns.IsAttackedBy(ns.GetCaller(), ns.GetHost())
	if !r1 {
		goto LABEL1
	}
	ns.GoBackHome(ns.GetCaller())
LABEL1:
	return
}
func MonsterGoHomeUrchin() {
	if !(ns.HasClass(ns.GetCaller(), ns.MONSTER) == true && ns.HasSubclass(ns.GetCaller(), ns.SMALL_MONSTER) == true) {
		goto LABEL1
	}
	ns.GoBackHome(ns.GetCaller())
LABEL1:
	return
}
func HostilizeMe() {
	ns.AggressionLevel(ns.GetTrigger(), 0.83)
}
func PlayWanderMusic() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.Music(17, 100)
LABEL1:
	return
}
func AwardQuestXP() {
	ns.AudioEvent(ns.SecretFound, wp39)
	ns.GiveXp(ns.GetHost(), 100)
}
func Secret01XP() {
	ns.AudioEvent(ns.SecretFound, wp33)
	ns.GiveXp(ns.GetHost(), 25)
	ns.PrintToAll("GeneralPrint:SecretFound")
}
func Secret02XP() {
	ns.AudioEvent(ns.SecretFound, wp34)
	ns.GiveXp(ns.GetHost(), 50)
	ns.PrintToAll("GeneralPrint:SecretFound")
}
func Secret03XP() {
	ns.AudioEvent(ns.SecretFound, wp35)
	ns.GiveXp(ns.GetHost(), 100)
	ns.PrintToAll("GeneralPrint:SecretFound")
}
func Secret04XP() {
	ns.AudioEvent(ns.SecretFound, wp36)
	ns.GiveXp(ns.GetHost(), 50)
	ns.PrintToAll("GeneralPrint:SecretFound")
}
func Secret05XP() {
	ns.AudioEvent(ns.SecretFound, wp37)
	ns.GiveXp(ns.GetHost(), 25)
	ns.PrintToAll("GeneralPrint:SecretFound")
}
func Secret06XP() {
	ns.AudioEvent(ns.SecretFound, wp38)
	ns.GiveXp(ns.GetHost(), 100)
	ns.PrintToAll("GeneralPrint:SecretFound")
}
func FreezePlayer() {
	ns.Frozen(ns.Object("War03b:Wolf1"), true)
	ns.Frozen(ns.Object("War03b:Wolf2"), true)
	ns.Frozen(ns.GetHost(), true)
	ns.CreatureIdle(ns.GetHost())
	ns.WideScreen(true)
	ns.FrameTimer(15, UrchinsEnter)
}
func UrchinsEnter() {
	gvar45 = gvar41
	ns.Move(obj47, wp60)
	ns.Move(obj48, wp61)
	ns.Move(obj49, wp62)
	ns.Move(obj50, wp63)
}
func CheckUrchins() {
	if gvar45 == gvar41 {
		goto LABEL1
	}
	return
LABEL1:
	if !(flag56 && flag57 && flag58 && flag59) {
		goto LABEL2
	}
	gvar45 = gvar42
	ns.FrameTimer(30, ShamanEnter)
LABEL2:
	return
}
func ShamanEnter() {
	ns.AudioEvent(ns.BigGong, wp65)
	ns.Move(obj30, wp64)
}
func ShamanCommand() {
	ns.AudioEvent(ns.UrchinShamanRecognize, wp65)
	ns.WideScreen(false)
	ns.Frozen(ns.Object("War03b:Wolf1"), false)
	ns.Frozen(ns.Object("War03b:Wolf2"), false)
	ns.Frozen(ns.GetHost(), false)
	gvar45 = gvar44
	ns.FrameTimer(10, UrchinsAttack)
	ns.FrameTimer(30, ShamanGuard)
}
func UrchinsAttack() {
	ns.AggressionLevel(obj47, 0.83)
	ns.Attack(obj47, ns.GetHost())
	ns.AggressionLevel(obj48, 0.83)
	ns.Attack(obj48, ns.GetHost())
	ns.AggressionLevel(obj49, 0.83)
	ns.Attack(obj49, ns.GetHost())
	ns.AggressionLevel(obj50, 0.83)
	ns.Attack(obj50, ns.GetHost())
	ns.AggressionLevel(obj30, 0.5)
	ns.ObjectGroupOn(gvar55)
}
func ShamanGuard() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
	)
	v0 = ns.GetWaypointX(wp66)
	v1 = ns.GetWaypointY(wp66)
	v2 = ns.GetWaypointX(wp67)
	v3 = ns.GetWaypointY(wp67)
	ns.CreatureGuard(obj30, v0, v1, v2, v3, 400)
}
func StartUrchinSetpiece() {
	ns.ObjectOff(ns.GetTrigger())
	ns.FrameTimer(1, FreezePlayer)
}
func Urchin01Report() {
	if gvar45 == gvar41 {
		goto LABEL1
	}
	return
LABEL1:
	ns.AudioEvent(ns.UrchinRecognize, wp65)
	flag56 = true
	CheckUrchins()
}
func Urchin02Report() {
	if gvar45 == gvar41 {
		goto LABEL1
	}
	return
LABEL1:
	flag57 = true
	ns.AudioEvent(ns.UrchinRecognize, wp65)
	CheckUrchins()
}
func Urchin03Report() {
	if gvar45 == gvar41 {
		goto LABEL1
	}
	return
LABEL1:
	flag58 = true
	ns.AudioEvent(ns.UrchinRecognize, wp65)
	CheckUrchins()
}
func Urchin04Report() {
	if gvar45 == gvar41 {
		goto LABEL1
	}
	return
LABEL1:
	flag59 = true
	ns.AudioEvent(ns.UrchinRecognize, wp65)
	CheckUrchins()
}
func ShamanReport() {
	if gvar45 == gvar42 {
		goto LABEL1
	}
	return
LABEL1:
	gvar45 = gvar43
	ns.FrameTimer(20, ShamanCommand)
	ns.LockDoor(obj51)
	ns.LockDoor(obj52)
}
func ReleaseShamans() {
	ns.AudioEvent(ns.BigGong, wp65)
	ns.UnlockDoor(obj51)
	ns.UnlockDoor(obj52)
	ns.Move(obj53, wp68)
	ns.Move(obj54, wp69)
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
