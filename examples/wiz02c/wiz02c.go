package wiz02c

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	obj4   ns.ObjectID
	obj5   ns.ObjectID
	obj6   ns.ObjectID
	obj7   ns.ObjectID
	obj8   ns.ObjectID
	obj9   ns.ObjectID
	obj10  ns.ObjectID
	obj11  ns.ObjectID
	obj12  ns.ObjectID
	obj13  ns.ObjectID
	obj14  ns.ObjectID
	obj15  ns.ObjectID
	obj16  ns.ObjectID
	obj17  ns.ObjectID
	obj18  ns.ObjectID
	obj19  ns.ObjectID
	gvar20 ns.ObjectGroupID
	gvar21 ns.WallGroupID
	gvar22 ns.WallGroupID
	gvar23 ns.WallGroupID
	gvar24 ns.WallGroupID
	gvar25 ns.WallGroupID
	gvar26 ns.WallGroupID
	gvar27 ns.WallGroupID
	gvar28 ns.WallGroupID
	gvar29 ns.WallGroupID
	gvar30 ns.WallGroupID
	flag31 bool
	flag32 bool
	flag33 bool
	flag34 bool
	flag35 bool
	flag36 bool
	flag37 bool
	flag38 bool
	flag39 bool
	flag40 bool
	wp41   ns.WaypointID
	wp42   ns.WaypointID
	wp43   ns.WaypointID
	obj44  ns.ObjectID
	obj45  ns.ObjectID
	obj46  ns.ObjectID
	obj47  ns.ObjectID
	obj48  ns.ObjectID
	gvar49 ns.ObjectGroupID
	gvar50 ns.ObjectGroupID
	gvar51 ns.ObjectGroupID
	flag52 bool
	flag53 bool
	wp54   ns.WaypointID
	wp55   ns.WaypointID
	wp56   ns.WaypointID
	wp57   ns.WaypointID
	obj58  ns.ObjectID
	obj59  ns.ObjectID
	obj60  ns.ObjectID
	obj61  ns.ObjectID
	obj62  ns.ObjectID
	obj63  ns.ObjectID
	obj64  ns.ObjectID
	obj65  ns.ObjectID
	obj66  ns.ObjectID
	gvar67 ns.WallGroupID
	gvar68 ns.WallGroupID
	gvar69 ns.WallGroupID
	gvar70 ns.WallGroupID
	gvar71 ns.WallGroupID
	gvar72 ns.WallGroupID
	gvar73 ns.WallGroupID
	flag74 bool
	flag75 bool
	flag76 bool
	flag77 bool
	flag78 bool
	flag79 bool
	flag80 bool
	flag81 bool
	flag82 bool
	flag83 bool
	flag84 bool
	fvar85 float32
	fvar86 float32
)

func init() {
	flag31 = false
	flag32 = false
	flag33 = false
	flag34 = false
	flag35 = false
	flag36 = false
	flag37 = false
	flag38 = false
	flag39 = false
	flag40 = false
	flag52 = false
	flag53 = false
	flag74 = false
	flag75 = false
	flag76 = false
	flag77 = false
	flag78 = false
	flag79 = false
	flag80 = false
	flag81 = false
	flag82 = true
	flag83 = false
	flag84 = false
}
func BookcaseSetup() {
	obj4 = ns.Object("Bookcase1AMover")
	obj5 = ns.Object("Bookcase1BMover")
	obj6 = ns.Object("Bookcase2AMover")
	obj7 = ns.Object("Bookcase2BMover")
	obj8 = ns.Object("Bookcase3AMover")
	obj9 = ns.Object("Bookcase3BMover")
	obj10 = ns.Object("Bookcase4AMover")
	obj11 = ns.Object("Bookcase4BMover")
	obj12 = ns.Object("Bookcase5AMover")
	obj13 = ns.Object("Bookcase5BMover")
	obj14 = ns.Object("Bookcase6AMover")
	obj15 = ns.Object("Bookcase7AMover")
	obj16 = ns.Object("Bookcase7BMover")
	obj17 = ns.Object("Orb1Mover")
	obj18 = ns.Object("Orb1BookcaseMover")
	gvar20 = ns.ObjectGroup("CreatureGroup1")
	gvar21 = ns.WallGroup("SecretWalls1")
	gvar22 = ns.WallGroup("SecretWalls2")
	gvar23 = ns.WallGroup("SecretWalls3")
	gvar24 = ns.WallGroup("SecretWalls4")
	gvar25 = ns.WallGroup("SecretWalls5")
	gvar26 = ns.WallGroup("SecretWalls6")
	gvar27 = ns.WallGroup("SecretWalls7")
	gvar28 = ns.WallGroup("SecretWalls10")
	gvar29 = ns.WallGroup("BreakableWalls1")
	gvar30 = ns.WallGroup("Orb1Walls")
	ns.ObjectOn(obj17)
}
func ActivateBookcase1() {
	ns.ObjectOff(ns.GetTrigger())
	ns.NoWallSound(true)
	ns.MoveWaypoint(wp42, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretWallWoodOpen, wp42)
	ns.ObjectOn(obj4)
	ns.ObjectOn(obj5)
	ns.WallGroupOpen(gvar21)
	ns.NoWallSound(false)
	ns.WallGroupBreak(gvar29)
}
func ActivateBookcase2() {
	ns.ObjectOff(ns.GetTrigger())
	ns.NoWallSound(true)
	ns.MoveWaypoint(wp42, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretWallWoodOpen, wp42)
	ns.ObjectOn(obj6)
	ns.ObjectOn(obj7)
	ns.WallGroupOpen(gvar22)
	ns.NoWallSound(false)
}
func ActivateBookcase3() {
	ns.ObjectOff(ns.GetTrigger())
	ns.NoWallSound(true)
	ns.MoveWaypoint(wp42, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretWallWoodOpen, wp42)
	ns.ObjectOn(obj8)
	ns.ObjectOn(obj9)
	ns.WallGroupOpen(gvar23)
	ns.NoWallSound(false)
}
func ActivateBookcase4() {
	ns.ObjectOff(ns.GetTrigger())
	ns.NoWallSound(true)
	ns.MoveWaypoint(wp42, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretWallWoodOpen, wp42)
	ns.ObjectOn(obj10)
	ns.ObjectOn(obj11)
	ns.WallGroupOpen(gvar24)
	ns.NoWallSound(false)
}
func ActivateBookcase5() {
	ns.ObjectOff(ns.GetTrigger())
	ns.NoWallSound(true)
	ns.MoveWaypoint(wp42, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretWallWoodOpen, wp42)
	ns.ObjectOn(obj12)
	ns.ObjectOn(obj13)
	ns.WallGroupOpen(gvar25)
	ns.NoWallSound(false)
	if flag32 {
		goto LABEL1
	}
	flag32 = true
	ns.PauseObject(ns.GetCaller(), 30)
	ns.MoveWaypoint(wp42, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.NecromancerTaunt, wp42)
LABEL1:
	return
}
func ActivateBookcase6() {
	ns.ObjectOff(ns.GetTrigger())
	ns.NoWallSound(true)
	ns.MoveWaypoint(wp42, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretWallWoodOpen, wp42)
	ns.ObjectOn(obj14)
	ns.WallGroupOpen(gvar26)
	ns.NoWallSound(false)
}
func ActivateBookcase7() {
	ns.ObjectOff(ns.GetTrigger())
	ns.NoWallSound(true)
	ns.MoveWaypoint(wp42, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretWallWoodOpen, wp42)
	ns.ObjectOn(obj15)
	ns.ObjectOn(obj16)
	ns.WallGroupOpen(gvar27)
	ns.NoWallSound(false)
}
func OpenOrb1Bookcase() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.NoWallSound(true)
	ns.MoveWaypoint(wp42, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretWallWoodOpen, wp42)
	ns.WallGroupOpen(gvar30)
	ns.ObjectOn(obj18)
	ns.NoWallSound(false)
LABEL1:
	return
}
func Orb1Secret() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.MoveWaypoint(wp43, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretFound, wp43)
	ns.GiveXp(ns.GetHost(), 50)
	ns.PrintToAll("GeneralPrint:SecretFound")
LABEL1:
	return
}
func OpenWalls() {
	ns.WallGroupOpen(gvar28)
}
func CloseWalls() {
	ns.WallGroupClose(gvar28)
}
func SetPieceSetup() {
	obj19 = ns.Object("Necromancer")
	obj45 = ns.Object("NecroTemp")
	obj48 = ns.Object("ThePit")
	gvar50 = ns.ObjectGroup("Piece1Triggers")
	gvar51 = ns.ObjectGroup("Piece2Triggers")
	gvar49 = ns.ObjectGroup("NecroRunTriggers")
	wp54 = ns.Waypoint("NecroWP1")
	wp41 = ns.Waypoint("NecroExitWP1")
	wp55 = ns.Waypoint("SummonWP1")
	wp56 = ns.Waypoint("SummonWP2")
	wp57 = ns.Waypoint("NecroPitWP")
	wp42 = ns.Waypoint("PlayerSounds")
	ns.StoryPic(obj19, "NecromancerPic")
	ns.FrameTimer(1, DisableStuff)
}
func InitScript() {
	obj58 = ns.Object("StatueN")
	obj59 = ns.Object("StatueS")
	obj60 = ns.Object("StatueE")
	obj61 = ns.Object("StatueW")
	obj62 = ns.Object("StatueTriggerN")
	obj64 = ns.Object("StatueTriggerE")
	obj63 = ns.Object("StatueTriggerW")
	obj65 = ns.Object("WallLight1")
	obj66 = ns.Object("WallLight2")
	gvar67 = ns.WallGroup("SpikeWalls1")
	gvar68 = ns.WallGroup("SpikeWalls2")
	gvar69 = ns.WallGroup("StatueWalls")
	gvar70 = ns.WallGroup("EastWalls")
	gvar71 = ns.WallGroup("EastExitWalls")
	gvar72 = ns.WallGroup("WestWalls")
	gvar73 = ns.WallGroup("WestExitWalls")
	fvar85 = ns.GetObjectX(obj59)
	fvar86 = ns.GetObjectY(obj59)
}
func LastSetup() {
	ns.ObjectGroupOff(gvar20)
}
func MapEntry() {
	ns.UnBlind()
}
func MapInitialize() {
	obj44 = ns.Object("BookOfOblivion")
	wp43 = ns.Waypoint("SecretSounds")
	BookcaseSetup()
	SetPieceSetup()
	InitScript()
	ns.FrameTimer(1, LastSetup)
}
func PlayerDeath() {
	ns.DeathScreen(2)
}
func DisableStuff() {
	ns.ObjectOff(obj48)
}
func Piece1Start() {
	ns.ObjectGroupOff(gvar50)
	ns.WideScreen(true)
	ns.Frozen(ns.GetHost(), true)
	ns.LookAtObject(ns.GetHost(), ns.ObjectID(wp54)) // FIXME
	ns.MoveObject(obj19, 1214, 4744)
	ns.Move(obj19, wp54)
}
func NecroSeePlayer() {
	if !ns.IsCaller(obj19) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.Frozen(obj19, true)
	ns.CreatureIdle(obj19)
	ns.LookAtObject(obj19, ns.GetHost())
	ns.Chat(obj19, "Wiz02:Necromancer02")
	ns.SecondTimer(2, SummonSpiders)
LABEL1:
	return
}
func SummonSpiders() {
	ns.CastSpellObjectLocation(ns.SPELL_SUMMON_SPIDER, obj19, ns.GetWaypointX(wp55), ns.GetWaypointY(wp55))
	ns.CastSpellLocationLocation(ns.SPELL_SUMMON_SPIDER, ns.GetObjectX(obj19), ns.GetObjectY(obj19), ns.GetWaypointX(wp56), ns.GetWaypointY(wp56))
}
func Spider1Trigger() {
	if ns.IsCaller(obj19) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	obj46 = ns.GetCaller()
	ns.Frozen(obj46, true)
	ns.CreatureIdle(obj46)
	ns.LookAtObject(obj46, ns.GetHost())
	ns.Frozen(obj19, false)
	ns.SetOwner(obj19, obj46)
	ns.Move(obj19, wp41)
LABEL1:
	return
}
func Spider2Trigger() {
	if ns.IsCaller(obj19) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	obj47 = ns.GetCaller()
	ns.Frozen(obj47, true)
	ns.CreatureIdle(obj47)
	ns.LookAtObject(obj47, ns.GetHost())
	ns.Frozen(obj19, false)
	ns.SetOwner(obj19, obj47)
	ns.Move(obj19, wp41)
LABEL1:
	return
}
func SetPiece1Over() {
	ns.ObjectOff(ns.GetTrigger())
	ns.WideScreen(false)
	ns.Frozen(ns.GetHost(), false)
	ns.Frozen(obj46, false)
	ns.Frozen(obj47, false)
	ns.SetOwner(obj46, obj47)
	ns.MoveObject(obj19, 3507, 3496)
	ns.CreatureIdle(obj19)
	ns.ObjectOff(obj19)
	ns.ObjectGroupOn(gvar20)
}
func PlayerFollowCheck() {
	ns.ObjectGroupOff(gvar51)
	ns.MoveObject(obj19, 3874, 3383)
	flag52 = true
	ns.WideScreen(true)
	ns.Frozen(ns.GetHost(), true)
	ns.CreatureIdle(ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj19)
	ns.ObjectOn(obj19)
	ns.FrameTimer(1, NecroGo)
}
func NecroRun() {
	ns.ObjectGroupOff(gvar49)
	ns.ObjectOn(obj19)
	ns.FrameTimer(1, NecroGo)
}
func NecroGo() {
	ns.Wander(obj19)
}
func Piece2Start() {
	if !flag52 {
		goto LABEL1
	}
	if !ns.IsCaller(obj19) {
		goto LABEL2
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.Frozen(obj19, true)
	ns.CreatureIdle(obj19)
	ns.LookAtObject(obj19, ns.GetHost())
	ns.SetDialog(obj19, ns.NORMAL, NecroTauntStart, NecroTauntEnd)
	ns.StartDialog(obj19, ns.GetHost())
LABEL2:
	goto LABEL3
LABEL1:
	ns.FrameTimer(1, Piece2Start)
LABEL3:
	return
}
func NecroWalkOnPit() {
	ns.Frozen(obj19, false)
	ns.Move(obj19, wp57)
}
func NecroScream1() {
	if flag53 {
		goto LABEL1
	}
	if !ns.IsCaller(obj19) {
		goto LABEL1
	}
	flag53 = true
	ns.MoveWaypoint(wp42, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.WallDestroyed, wp42)
	ns.Chat(obj19, "Wiz02:Necromancer05")
LABEL1:
	return
}
func NecroScream2() {
	if !ns.IsCaller(obj19) {
		goto LABEL1
	}
	ns.DestroyEveryChat()
	ns.AudioEvent(ns.TrollRecognize, wp42)
	ns.MoveObject(obj19, 3935, 2442)
	ns.Chat(obj48, "Wiz02:Necromancer06")
	ns.WideScreen(false)
	ns.Frozen(ns.GetHost(), false)
	ns.FrameTimer(1, KillNecroTemp)
LABEL1:
	return
}
func NecroTauntStart() {
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02:Necromancer04")
}
func NecroTauntEnd() {
	ns.CancelDialog(obj19)
	NecroWalkOnPit()
}
func CheckForBook() {
	ns.JournalEdit(ns.GetHost(), "GetNecromancer", 4)
	ns.MoveWaypoint(wp42, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.FlagCapture, wp42)
	ns.GiveXp(ns.GetHost(), 250)
	ns.JournalEntry(ns.GetHost(), "ArchivistBook", 2)
	ns.PrintToAll("Con01a:NewJournalEntry")
}
func KillNecroTemp() {
	ns.Damage(obj45, 0, 500, 0)
}
func OpenSpikeWalls1() {
	if flag74 {
		goto LABEL1
	}
	flag74 = true
	ns.WallGroupOpen(gvar67)
	ns.MoveObject(obj65, 1871, 3998)
	goto LABEL2
LABEL1:
	ns.WallGroupOpen(gvar68)
	ns.MoveObject(obj66, 1818, 3928)
LABEL2:
	return
}
func OpenSpikeWalls2() {
	if flag74 {
		goto LABEL1
	}
	flag74 = true
	ns.WallGroupOpen(gvar67)
	ns.MoveObject(obj65, 1871, 3998)
	goto LABEL2
LABEL1:
	ns.WallGroupOpen(gvar68)
	ns.MoveObject(obj66, 1818, 3928)
LABEL2:
	return
}
func TriggerNHit() {
	flag76 = true
	ns.MoveObject(obj59, 1957, 3768)
	ns.MoveObject(obj61, 1957, 3768)
	ns.MoveObject(obj60, 1957, 3768)
	ns.MoveObject(obj58, fvar85, fvar86)
	ns.WallGroupOpen(gvar69)
	ns.WallGroupClose(gvar72)
	ns.WallGroupClose(gvar70)
}
func TriggerEHit() {
	flag77 = true
	ns.MoveObject(obj59, 1957, 3768)
	ns.MoveObject(obj61, 1957, 3768)
	ns.MoveObject(obj58, 1957, 3768)
	ns.MoveObject(obj60, fvar85, fvar86)
	ns.WallGroupOpen(gvar70)
	ns.WallGroupClose(gvar69)
	ns.WallGroupClose(gvar72)
}
func TriggerWHit() {
	flag78 = true
	ns.MoveObject(obj59, 1957, 3768)
	ns.MoveObject(obj60, 1957, 3768)
	ns.MoveObject(obj58, 1957, 3768)
	ns.MoveObject(obj61, fvar85, fvar86)
	ns.WallGroupOpen(gvar72)
	ns.WallGroupClose(gvar70)
	ns.WallGroupClose(gvar69)
}
func TriggerNRelease() {
	flag76 = false
	if !flag79 {
		goto LABEL1
	}
	flag79 = false
	ns.WallGroupClose(gvar69)
	ns.WallGroupClose(gvar70)
	ns.WallGroupClose(gvar72)
LABEL1:
	return
}
func TriggerWRelease() {
	flag78 = false
	if !flag81 {
		goto LABEL1
	}
	flag81 = false
	ns.WallGroupClose(gvar72)
	ns.WallGroupClose(gvar69)
	ns.WallGroupClose(gvar70)
LABEL1:
	return
}
func TriggerERelease() {
	flag77 = false
	if !flag80 {
		goto LABEL1
	}
	flag79 = false
	ns.WallGroupClose(gvar70)
	ns.WallGroupClose(gvar69)
	ns.WallGroupClose(gvar72)
LABEL1:
	return
}
func OpenEastExitWalls() {
	ns.WallGroupOpen(gvar71)
}
func CloseEastExitWalls() {
	ns.WallGroupClose(gvar71)
}
func OpenWestExitWalls() {
	ns.WallGroupOpen(gvar73)
}
func CloseWestExitWalls() {
	ns.WallGroupClose(gvar73)
}
func OpenPuzzleWalls() {
	ns.WallGroupOpen(gvar69)
	ns.WallGroupOpen(gvar67)
	ns.WallGroupOpen(gvar68)
}
func SecretXP1() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	if flag83 {
		goto LABEL1
	}
	flag83 = true
	ns.MoveWaypoint(wp43, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretFound, wp43)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 100)
LABEL1:
	return
}
func SecretXP2() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	if flag84 {
		goto LABEL1
	}
	flag84 = true
	ns.MoveWaypoint(wp43, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretFound, wp43)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 100)
LABEL1:
	return
}
func HallSecret1() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.MoveWaypoint(wp43, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretFound, wp43)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 25)
LABEL1:
	return
}
func HallSecret2() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.MoveWaypoint(wp43, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretFound, wp43)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 25)
LABEL1:
	return
}
func SuperSecretXP() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.MoveWaypoint(wp43, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretFound, wp43)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 150)
LABEL1:
	return
}
func SuperSecret2XP() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.MoveWaypoint(wp43, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretFound, wp43)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 50)
LABEL1:
	return
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
