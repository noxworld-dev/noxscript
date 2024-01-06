package con08e

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	obj4   ns.ObjectID
	obj5   ns.ObjectID
	obj6   ns.ObjectID
	obj7   ns.ObjectID
	obj8   ns.ObjectID
	wp9    ns.WaypointID
	wp10   ns.WaypointID
	wp11   ns.WaypointID
	wp12   ns.WaypointID
	wp13   ns.WaypointID
	obj14  ns.ObjectID
	obj15  ns.ObjectID
	obj16  ns.ObjectID
	obj17  ns.ObjectID
	obj18  ns.ObjectID
	obj19  ns.ObjectID
	obj20  ns.ObjectID
	obj21  ns.ObjectID
	obj22  ns.ObjectID
	obj23  ns.ObjectID
	obj24  ns.ObjectID
	gvar25 ns.ObjectGroupID
	gvar26 ns.ObjectGroupID
	gvar27 ns.WallGroupID
	gvar28 ns.WallGroupID
	gvar29 ns.WallGroupID
	gvar30 ns.WallGroupID
	wp31   ns.WaypointID
	wp32   ns.WaypointID
	wp33   ns.WaypointID
	wp34   ns.WaypointID
	wp35   ns.WaypointID
	wp36   ns.WaypointID
	wp37   ns.WaypointID
	wp38   ns.WaypointID
	wp39   ns.WaypointID
	wp40   ns.WaypointID
	gvar41 int
	gvar42 int
	gvar43 int
	gvar44 int
	gvar45 int
	gvar46 int
	gvar47 int
	gvar48 int
	gvar49 int
	obj50  ns.ObjectID
	obj51  ns.ObjectID
	flag52 bool
	flag53 bool
	flag54 bool
	flag55 bool
	flag56 bool
	wp57   ns.WaypointID
	flag58 bool
)

func init() {
	gvar41 = 0
	gvar42 = 1
	gvar43 = 2
	gvar44 = 3
	gvar45 = 4
	gvar46 = 5
	gvar47 = 6
	gvar48 = gvar41
	gvar49 = gvar46
	flag52 = false
	flag53 = false
	flag54 = false
	flag55 = false
	flag56 = false
	flag58 = false
}
func PlayerDeath() {
	ns.DeathScreen(8)
}
func FireFONTrap01() {
	ns.CastSpellObjectLocation(ns.SPELL_FORCE_OF_NATURE, obj4, ns.GetWaypointX(wp9), ns.GetWaypointY(wp9))
	ns.CastSpellObjectLocation(ns.SPELL_FORCE_OF_NATURE, obj5, ns.GetWaypointX(wp10), ns.GetWaypointY(wp10))
	ns.CastSpellObjectLocation(ns.SPELL_FORCE_OF_NATURE, obj6, ns.GetWaypointX(wp11), ns.GetWaypointY(wp11))
	ns.CastSpellObjectLocation(ns.SPELL_FORCE_OF_NATURE, obj7, ns.GetWaypointX(wp12), ns.GetWaypointY(wp12))
	ns.SecondTimer(5, FireFONTrap01)
}
func InitializeFONtraps() {
	obj4 = ns.Object("FON_Origin01")
	obj5 = ns.Object("FON_Origin02")
	obj6 = ns.Object("FON_Origin03")
	obj7 = ns.Object("FON_Origin04")
	obj8 = ns.Object("FON_Origin05")
	wp9 = ns.Waypoint("FON_Target01")
	wp10 = ns.Waypoint("FON_Target02")
	wp11 = ns.Waypoint("FON_Target03")
	wp12 = ns.Waypoint("FON_Target04")
	wp13 = ns.Waypoint("FON_Target05")
	FireFONTrap01()
}
func OpenFONSecret() {
	ns.ObjectOff(ns.GetTrigger())
	ns.AudioEvent(ns.SecretWallOpen, ns.Waypoint("FonSecretWP"))
	ns.WallOpen(ns.Wall(44, 122))
}
func MapInitialize() {
	obj14 = ns.Object("MoverA1")
	obj15 = ns.Object("MoverA2")
	obj16 = ns.Object("MoverA3")
	obj17 = ns.Object("MoverB1")
	obj18 = ns.Object("MoverB2")
	obj19 = ns.Object("MoverB3")
	obj20 = ns.Object("Priest01")
	obj21 = ns.Object("Priest02")
	obj22 = ns.Object("Imp01")
	obj23 = ns.Object("Imp02")
	obj24 = ns.Object("WierdlingPool")
	wp31 = ns.Waypoint("GolemWP_A")
	wp32 = ns.Waypoint("GolemWP_B")
	wp33 = ns.Waypoint("Priest01WP")
	wp34 = ns.Waypoint("Priest02WP")
	wp35 = ns.Waypoint("Priest1Home")
	wp36 = ns.Waypoint("Priest2Home")
	wp37 = ns.Waypoint("Priest1ExitWP")
	wp38 = ns.Waypoint("Priest2ExitWP")
	wp39 = ns.Waypoint("ImpCreationWP")
	wp40 = ns.Waypoint("WierdlingPoolWP")
	gvar25 = ns.ObjectGroup("GolemLights")
	gvar26 = ns.ObjectGroup("GolemHandleTriggers")
	gvar27 = ns.WallGroup("GolemArenaEntranceWalls")
	gvar28 = ns.WallGroup("GolemArenaExitWalls")
	gvar29 = ns.WallGroup("WierdlingExitWalls")
	gvar30 = ns.WallGroup("AldwynElevatorWalls")
	ns.SetOwner(ns.GetHost(), obj20)
	ns.SetOwner(ns.GetHost(), obj21)
	InitializeFONtraps()
}
func OpenElevatorWalls() {
	ns.ObjectOff(ns.GetTrigger())
	ns.WallGroupOpen(gvar30)
}
func NoMonsters() {
	r1 := ns.IsAttackedBy(ns.GetCaller(), ns.GetHost())
	if !r1 {
		goto LABEL1
	}
	ns.GoBackHome(ns.GetCaller())
LABEL1:
	return
}
func PlaySub1Music() {
	ns.Music(18, 100)
}
func PlaySub2Music() {
	ns.Music(19, 100)
}
func SecretFound() {
	ns.ObjectOff(ns.GetTrigger())
	ns.MoveWaypoint(ns.Waypoint("FonSecretWP"), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretFound, ns.Waypoint("FonSecretWP"))
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 200)
}
func PriestReport() {
	var v0 int
	v0 = gvar48
	if v0 == gvar42 {
		goto LABEL1
	}
	if v0 == gvar43 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	if !ns.IsTrigger(obj22) {
		goto LABEL4
	}
	flag53 = true
	ns.Effect(ns.SMOKE_BLAST, ns.GetObjectX(obj22), ns.GetObjectY(obj22), 0, 0)
	ns.AudioEvent(ns.ImpRecognize, wp57)
	ns.MoveObject(obj20, ns.GetObjectX(obj22), ns.GetObjectY(obj22))
	ns.Delete(obj22)
LABEL4:
	if !ns.IsTrigger(obj23) {
		goto LABEL5
	}
	flag54 = true
	ns.Effect(ns.SMOKE_BLAST, ns.GetObjectX(obj23), ns.GetObjectY(obj23), 0, 0)
	ns.AudioEvent(ns.ImpRecognize, wp57)
	ns.MoveObject(obj21, ns.GetObjectX(obj23), ns.GetObjectY(obj23))
	ns.Delete(obj23)
LABEL5:
	if !flag53 {
		goto LABEL6
	}
	if !flag54 {
		goto LABEL6
	}
	gvar48 = gvar44
	ns.FrameTimer(30, SummonGolems)
LABEL6:
	goto LABEL3
LABEL2:
	if !ns.IsTrigger(obj22) {
		goto LABEL7
	}
	flag55 = true
LABEL7:
	if !ns.IsTrigger(obj23) {
		goto LABEL8
	}
	flag56 = true
LABEL8:
	if !(flag55 && flag56) {
		goto LABEL9
	}
	ns.WallGroupClose(gvar28)
	gvar48 = gvar45
	ns.Delete(obj22)
	ns.Delete(obj23)
LABEL9:
	goto LABEL3
LABEL3:
	return
}
func GolemDie() {
	if !(ns.CurrentHealth(obj50) <= 0 && ns.CurrentHealth(obj51) <= 0) {
		goto LABEL1
	}
	ns.GiveXp(ns.GetHost(), 1000)
	ns.WallGroupOpen(gvar28)
	ns.AudioEvent(ns.BigGong, wp33)
	ns.AudioEvent(ns.BigGong, wp34)
LABEL1:
	return
}
func GolemsAttack() {
	ns.Attack(obj50, ns.GetHost())
	ns.Attack(obj51, ns.GetHost())
}
func RemoveLetterbox() {
	ns.WideScreen(false)
	ns.Delete(obj22)
	ns.Delete(obj23)
	ns.Frozen(ns.GetHost(), false)
	ns.FrameTimer(45, GolemsAttack)
	ns.Music(27, 100)
}
func ImpsExit() {
	ns.Move(obj22, wp37)
	ns.Move(obj23, wp38)
}
func PriestsExit() {
	gvar48 = gvar43
	ns.Effect(ns.SMOKE_BLAST, ns.GetObjectX(obj20), ns.GetObjectY(obj20), 0, 0)
	ns.MoveWaypoint(wp39, ns.GetObjectX(obj20), ns.GetObjectY(obj20))
	obj22 = ns.CreateObject("Imp", wp39)
	ns.MoveObject(obj20, ns.GetWaypointX(wp35), ns.GetWaypointY(wp35))
	ns.AudioEvent(ns.ImpRecognize, wp57)
	ns.Effect(ns.SMOKE_BLAST, ns.GetObjectX(obj21), ns.GetObjectY(obj21), 0, 0)
	ns.MoveWaypoint(wp39, ns.GetObjectX(obj21), ns.GetObjectY(obj21))
	obj23 = ns.CreateObject("Imp", wp39)
	ns.MoveObject(obj21, ns.GetWaypointX(wp36), ns.GetWaypointY(wp36))
	ns.SetOwner(ns.GetHost(), obj22)
	ns.SetOwner(ns.GetHost(), obj23)
	ns.AggressionLevel(obj22, 0)
	ns.AggressionLevel(obj23, 0)
	ns.SetCallback(obj22, 11, PriestReport)
	ns.SetCallback(obj23, 11, PriestReport)
	ns.FrameTimer(1, ImpsExit)
}
func GetGolemHandles() {
	var v0 int
	if !(flag52 || ns.IsCaller(obj22) || ns.IsCaller(obj23)) {
		goto LABEL1
	}
	return
LABEL1:
	v0 = gvar49
	if v0 == gvar46 {
		goto LABEL2
	}
	if v0 == gvar47 {
		goto LABEL3
	}
	goto LABEL4
LABEL2:
	ns.ObjectOff(ns.GetTrigger())
	ns.DestroyEveryChat()
	obj50 = ns.GetCaller()
	ns.SetCallback(obj50, 5, GolemDie)
	ns.ClearOwner(obj20)
	ns.CreatureIdle(obj50)
	ns.ClearOwner(obj50)
	ns.AggressionLevel(obj50, 0)
	gvar49 = gvar47
	goto LABEL5
LABEL3:
	ns.ObjectOff(ns.GetTrigger())
	ns.DestroyEveryChat()
	obj51 = ns.GetCaller()
	ns.SetCallback(obj51, 5, GolemDie)
	ns.ClearOwner(obj21)
	ns.ClearOwner(obj51)
	ns.CreatureIdle(obj51)
	ns.AggressionLevel(obj51, 0)
	ns.LookAtObject(obj50, ns.GetHost())
	ns.LookAtObject(obj51, ns.GetHost())
	flag52 = true
	ns.AggressionLevel(obj50, 0.83)
	ns.AggressionLevel(obj51, 0.83)
	RemoveLetterbox()
	goto LABEL5
LABEL4:
	goto LABEL5
LABEL5:
	return
}
func SummonGolems() {
	ns.CastSpellObjectLocation(ns.SPELL_SUMMON_STONE_GOLEM, obj20, ns.GetWaypointX(wp31), ns.GetWaypointY(wp31))
	ns.CastSpellObjectLocation(ns.SPELL_SUMMON_STONE_GOLEM, obj21, ns.GetWaypointX(wp32), ns.GetWaypointY(wp32))
	ns.AudioEvent(ns.SummonCast, wp33)
	ns.AudioEvent(ns.SummonCast, wp34)
	ns.ObjectGroupOn(gvar26)
	ns.ObjectGroupOff(gvar25)
	ns.FrameTimer(60, PriestsExit)
}
func ImpsEnter() {
	ns.WallGroupOpen(gvar28)
	ns.Move(obj22, wp33)
	ns.Move(obj23, wp34)
}
func GolemSetPiece() {
	if gvar48 == gvar41 {
		goto LABEL1
	}
	return
LABEL1:
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL2
	}
	ns.MusicPushEvent()
	ns.Music(10, 100)
	wp57 = ns.Waypoint("GolemSP_AudioOrigin")
	ns.MoveWaypoint(wp57, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	gvar48 = gvar42
	ns.WallGroupClose(gvar27)
	ns.Frozen(ns.GetHost(), true)
	ns.CreatureIdle(ns.GetHost())
	ns.WideScreen(true)
	ns.SetOwner(ns.GetHost(), obj22)
	ns.SetOwner(ns.GetHost(), obj23)
	ns.FrameTimer(15, ImpsEnter)
LABEL2:
	return
}
func Golem01() {
	ns.Chat(obj50, "Hi. I'm Golem1.")
}
func Golem02() {
	ns.Chat(obj50, "Hi. I'm Golem2, and I'm an alcholic.")
}
func RetrieveWierdling() {
	if !flag58 {
		goto LABEL1
	}
	return
LABEL1:
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL2
	}
	ns.SetQuestStatus(1, "Chapter8:HasWeirdling")
	flag58 = true
	ns.SetHalberd(2)
	ns.ObjectOff(obj24)
	ns.WallGroupOpen(gvar29)
	ns.JournalEdit(ns.GetHost(), "Chapter8Wierdling", 4)
	ns.PrintToAll("Con02a:ObjectiveComplete")
	ns.JournalEntry(ns.GetHost(), "Chapter8MeetCaptain", 2)
	ns.PrintToAll("Con01a:NewJournalEntry")
	ns.MusicPopEvent()
LABEL2:
	return
}
func OnEvent(typ string) {
	switch typ {
	case "PlayerDeath":
		PlayerDeath()
	case "MapInitialize":
		MapInitialize()
	}
}
