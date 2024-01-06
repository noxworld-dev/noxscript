package con06a

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	obj4   ns.ObjectID
	obj5   ns.ObjectID
	obj6   ns.ObjectID
	obj7   ns.ObjectID
	obj8   ns.ObjectID
	obj9   ns.ObjectID
	obj10  ns.ObjectID
	wp11   ns.WaypointID
	gvar12 ns.ObjectGroupID
	gvar13 ns.ObjectGroupID
	gvar14 ns.ObjectGroupID
	gvar15 ns.ObjectGroupID
	flag16 bool
	flag17 bool
	obj18  ns.ObjectID
	obj19  ns.ObjectID
	wp20   ns.WaypointID
	gvar21 ns.WallGroupID
	flag22 bool
	ivar23 int
	obj24  ns.ObjectID
	obj25  ns.ObjectID
	obj26  ns.ObjectID
	obj27  ns.ObjectID
	obj28  ns.ObjectID
	obj29  ns.ObjectID
	obj30  ns.ObjectID
	gvar31 ns.ObjectGroupID
	gvar32 ns.ObjectGroupID
	gvar33 ns.ObjectGroupID
	gvar34 ns.ObjectGroupID
	flag35 bool
	obj36  ns.ObjectID
	obj37  ns.ObjectID
	obj38  ns.ObjectID
	obj39  ns.ObjectID
	obj40  ns.ObjectID
	obj41  ns.ObjectID
	obj42  ns.ObjectID
	obj43  ns.ObjectID
	gvar44 ns.ObjectGroupID
	wp45   ns.WaypointID
	wp46   ns.WaypointID
	obj47  ns.ObjectID
	ivar48 int
	flag49 bool
	obj50  ns.ObjectID
	obj51  ns.ObjectID
	obj52  ns.ObjectID
	obj53  ns.ObjectID
	obj54  ns.ObjectID
	obj55  ns.ObjectID
	obj56  ns.ObjectID
	obj57  ns.ObjectID
	obj58  ns.ObjectID
	obj59  ns.ObjectID
	obj60  ns.ObjectID
	obj61  ns.ObjectID
	wp62   ns.WaypointID
	gvar63 ns.ObjectGroupID
)

func init() {
	flag16 = false
	flag17 = true
	flag22 = false
	ivar23 = 0
	flag35 = true
	flag49 = true
}
func KillTownie() {
	ns.AudioEvent(ns.HumanMaleDie, wp11)
	ns.MoveObject(obj4, ns.GetWaypointX(wp11), ns.GetWaypointY(wp11))
	ns.MoveObject(obj7, 4071, 1828)
	ns.MoveObject(obj6, 3999, 1844)
}
func townieTalk1Start() {
	ns.ObjectGroupOff(gvar14)
	ns.TellStory(ns.UrchinRecognize, "War06a:TownspeopleSpeak3")
}
func townieTalk1End() {
	ns.ObjectGroupOn(gvar14)
	ns.CancelDialog(obj4)
	ns.Delete(obj5)
	ns.FrameTimer(1, KillTownie)
}
func EnableSkeletonGroup() {
	ns.ObjectOff(ns.GetTrigger())
	ns.ObjectGroupOn(gvar13)
}
func captainVanishes() {
	ns.ObjectOff(ns.GetTrigger())
	ns.Delete(obj10)
	ns.Delete(obj8)
	ns.Delete(obj9)
	ns.ObjectGroupOn(gvar12)
	EnableSkeletonGroup()
}
func GuardiansAttack() {
	ns.AggressionLevel(obj24, 0.83)
	ns.Attack(obj24, obj18)
	ns.AggressionLevel(obj25, 0.83)
	ns.Attack(obj25, obj18)
	ns.AggressionLevel(obj26, 0.83)
	ns.Attack(obj26, obj19)
	ns.AggressionLevel(obj18, 0.83)
	ns.Attack(obj18, obj24)
	ns.AggressionLevel(obj19, 0.83)
	ns.Attack(obj19, obj26)
}
func Guardian01DialogStart() {
	ns.TellStory(ns.SwordsmanHurt, "Con06a:CityGuardSpeak1")
}
func Guardian01DialogEnd() {
	ns.CancelDialog(obj18)
	ns.NoWallSound(true)
	ns.WallGroupOpen(gvar21)
	ns.NoWallSound(false)
	ns.BecomePet(obj18)
	ns.BecomePet(obj19)
	ns.CreatureFollow(obj18, ns.GetHost())
	ns.CreatureFollow(obj19, ns.GetHost())
	ns.Frozen(ns.GetHost(), false)
	ns.WideScreen(false)
	ns.ObjectGroupOn(gvar15)
	ns.ObjectGroupOn(gvar13)
}
func Zombie1Rise() {
	var (
		v0 float32
		v1 float32
	)
	v0 = ns.GetObjectX(obj24)
	v1 = ns.GetObjectY(obj24)
	ns.RaiseZombie(obj24)
	ns.MoveWaypoint(wp20, v0, v1)
	ns.AudioEvent(ns.ZombieRecognize, ns.Waypoint("AudioOrigin"))
}
func Zombie2Rise() {
	var (
		v0 float32
		v1 float32
	)
	v0 = ns.GetObjectX(obj25)
	v1 = ns.GetObjectY(obj25)
	ns.RaiseZombie(obj25)
	ns.MoveWaypoint(wp20, v0, v1)
	ns.AudioEvent(ns.ZombieRecognize, ns.Waypoint("AudioOrigin"))
}
func Zombie3Rise() {
	var (
		v0 float32
		v1 float32
	)
	v0 = ns.GetObjectX(obj26)
	v1 = ns.GetObjectY(obj26)
	ns.RaiseZombie(obj26)
	ns.MoveWaypoint(wp20, v0, v1)
	ns.AudioEvent(ns.ZombieRecognize, ns.Waypoint("AudioOrigin"))
}
func SetpieceZombieDie() {
	ivar23 += 1
	if !(ivar23 >= 3) {
		goto LABEL1
	}
	ns.StoryPic(obj18, "WarriorPic")
	ns.SetDialog(obj18, ns.NORMAL, Guardian01DialogStart, Guardian01DialogEnd)
	ns.LookAtObject(obj18, ns.GetHost())
	ns.LookAtObject(obj19, ns.GetHost())
	ns.StartDialog(obj18, ns.GetHost())
	ns.SetCallback(obj24, 5, NullCallback)
	ns.SetCallback(obj25, 5, NullCallback)
	ns.SetCallback(obj26, 5, NullCallback)
LABEL1:
	return
}
func SetZombieDieCallbacks() {
	ns.SetCallback(obj24, 5, SetpieceZombieDie)
	ns.SetCallback(obj25, 5, SetpieceZombieDie)
	ns.SetCallback(obj26, 5, SetpieceZombieDie)
}
func NullCallback() {
}
func InitializeGuardianSetpiece() {
	obj24 = ns.Object("SetpieceZombie01")
	obj25 = ns.Object("SetpieceZombie02")
	obj26 = ns.Object("SetpieceZombie03")
	ns.Damage(obj24, 0, 100, 0)
	ns.Damage(obj25, 0, 100, 0)
	ns.Damage(obj26, 0, 100, 0)
	ns.ZombieStayDown(obj24)
	ns.ZombieStayDown(obj25)
	ns.ZombieStayDown(obj26)
	ns.FrameTimer(30, SetZombieDieCallbacks)
}
func StartGuardianSetpiece() {
	if !(ns.IsCaller(ns.GetHost()) && !flag22) {
		goto LABEL1
	}
	flag22 = true
	ns.CreatureIdle(ns.GetHost())
	ns.Frozen(ns.GetHost(), true)
	ns.WideScreen(true)
	ns.FrameTimer(15, Zombie1Rise)
	ns.FrameTimer(25, Zombie2Rise)
	ns.FrameTimer(20, Zombie3Rise)
	ns.FrameTimer(45, GuardiansAttack)
	ns.ObjectGroupOff(gvar13)
LABEL1:
	return
}
func TurnOffLights() {
	ns.ObjectToggle(obj30)
	ns.ObjectGroupToggle(gvar31)
	ns.ObjectGroupToggle(gvar32)
	ns.ObjectGroupToggle(gvar33)
	ns.ObjectGroupToggle(gvar34)
	ns.ObjectToggle(obj27)
	ns.ObjectToggle(obj28)
	ns.ObjectToggle(obj29)
}
func PlayerDeath() {
	ns.DeathScreen(6)
}
func captainTalk() {
	ns.StartDialog(obj10, ns.GetHost())
}
func captainTalk1Start() {
	ns.Frozen(ns.GetHost(), true)
	ns.TellStory(ns.ArcherRecognize, "Con01a:CaptainProd")
}
func captainTalk1End() {
	ns.Frozen(ns.GetHost(), false)
	ns.SetDialog(obj10, ns.NORMAL, captainTalk1Start, captainTalk1End)
}
func undeadBreakIn() {
	ns.WallBreak(ns.Wall(220, 118))
	ns.WallBreak(ns.Wall(219, 119))
	ns.WallBreak(ns.Wall(220, 120))
}
func KirikEnd1() {
	ns.SetDialog(obj43, ns.NORMAL, KirikStart1, KirikEnd1)
}
func KirikStart1() {
	ns.TellStory(ns.HumanMaleEatFood, "War01A.scr:Guard1Talk01")
}
func SecretRoom() {
	ns.ObjectGroupOff(gvar44)
	ns.MoveWaypoint(wp45, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretFound, wp45)
	ns.GiveXp(ns.GetHost(), 100)
	ns.PrintToAll("GeneralPrint:SecretFound")
}
func InitializeTownspeople() {
	obj50 = ns.Object("Townsman1")
	obj51 = ns.Object("Townsman2")
	obj52 = ns.Object("Townsman3")
	obj53 = ns.Object("Townsman4")
	obj54 = ns.Object("Townsman5")
	obj55 = ns.Object("Townsman6")
	obj56 = ns.Object("Townswoman1")
	obj57 = ns.Object("Townswoman2")
	obj58 = ns.Object("Townswoman3")
	obj59 = ns.Object("Townswoman4")
	obj60 = ns.Object("Townswoman5")
	obj61 = ns.Object("Townswoman6")
	wp62 = ns.Waypoint("TownspeopleStorage")
	gvar63 = ns.ObjectGroup("Townspeople")
	ns.StoryPic(obj18, "WarriorPic")
	ns.StoryPic(obj19, "WarriorPic")
	ns.StoryPic(obj40, "WarriorPic")
	ns.SetDialog(obj40, ns.NORMAL, fireKnightsTalk1Start, fireKnightsTalk1End)
	ns.StoryPic(obj41, "Warrior2Pic")
	ns.SetDialog(obj41, ns.NORMAL, fireKnightsTalk2Start, fireKnightsTalk2End)
	ns.StoryPic(obj42, "Warrior3Pic")
	ns.SetDialog(obj42, ns.NORMAL, fireKnightsTalk3Start, fireKnightsTalk3End)
	ns.FrameTimer(15, OwnTownspeople)
}
func NewJournalEntry() {
	ns.JournalEntry(ns.GetHost(), "War6Fortress", 2)
	ns.PrintToAll("Con01a:NewJournalEntry")
}
func MapInitialize() {
	obj10 = ns.Object("Captain")
	obj8 = ns.Object("Basket")
	obj9 = ns.Object("BasketShadow")
	obj4 = ns.Object("Townie1")
	obj6 = ns.Object("DeadGuyBoots")
	obj7 = ns.Object("DeadGuyChain")
	obj5 = ns.Object("F6Elite2")
	obj36 = ns.Object("GauntletDoorR")
	obj37 = ns.Object("GauntletDoorL")
	obj38 = ns.Object("IxDoorR")
	obj39 = ns.Object("IxDoorL")
	obj18 = ns.Object("F6FireGuard10")
	obj19 = ns.Object("F6FireGuard7")
	obj40 = ns.Object("F6FireGuard9")
	obj41 = ns.Object("F2Fire6")
	obj42 = ns.Object("F2Fire8")
	obj27 = ns.Object("GearLight1")
	obj28 = ns.Object("GearLight2")
	obj29 = ns.Object("GearLight3")
	obj30 = ns.Object("StupidTorch")
	obj43 = ns.Object("Kirik")
	ns.SetOwner(ns.GetHost(), obj43)
	ns.SetOwner(ns.GetHost(), obj5)
	gvar12 = ns.ObjectGroup("ZombieGroupA")
	gvar13 = ns.ObjectGroup("SkeletonGroup01")
	gvar31 = ns.ObjectGroup("StreetLamps")
	gvar32 = ns.ObjectGroup("StreetLamps2")
	gvar33 = ns.ObjectGroup("LightGears")
	gvar34 = ns.ObjectGroup("LightGears2")
	gvar44 = ns.ObjectGroup("SecretRoom2Triggers")
	gvar15 = ns.ObjectGroup("EveryMonsterOnMap")
	gvar14 = ns.ObjectGroup("FleemanMonsters")
	wp11 = ns.Waypoint("TownieWP")
	wp20 = ns.Waypoint("AudioOrigin")
	wp45 = ns.Waypoint("PlayerSounds")
	wp46 = ns.Waypoint("WellWP")
	gvar21 = ns.WallGroup("GrdnStpcWalls")
	ns.StoryPic(obj10, "AirshipCaptainPic")
	ns.SetDialog(obj10, ns.NORMAL, captainTalk1Start, captainTalk1End)
	ns.StoryPic(obj5, "WoundedWarriorPic")
	ns.SetDialog(obj5, ns.NORMAL, townieTalk1Start, townieTalk1End)
	ns.StoryPic(obj43, "Warrior4Pic")
	ns.SetDialog(obj43, ns.NORMAL, KirikStart1, KirikEnd1)
	ns.LockDoor(obj36)
	ns.LockDoor(obj37)
	ns.LockDoor(obj38)
	ns.LockDoor(obj39)
	InitializeGuardianSetpiece()
	InitializeTownspeople()
	ns.StartupScreen(6)
	ns.FrameTimer(15, NewJournalEntry)
	ns.Music(16, 100)
	ns.FrameTimer(1, captainTalk)
	ns.Damage(obj4, 0, 100, 0)
}
func NoEnemys() {
	r1 := ns.IsAttackedBy(ns.GetCaller(), ns.GetHost())
	if !r1 {
		goto LABEL1
	}
	ns.GoBackHome(ns.GetCaller())
LABEL1:
	return
}
func MonstersGoHome() {
	r1 := ns.IsAttackedBy(ns.GetCaller(), ns.GetHost())
	if !r1 {
		goto LABEL1
	}
	ns.AggressionLevel(ns.GetCaller(), 0.16)
	ns.GoBackHome(ns.GetCaller())
LABEL1:
	return
}
func OwnTownspeople() {
	ns.GroupSetOwner(ns.GetHost(), gvar63)
}
func fireKnightsRecognize() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ivar48 = ns.Random(1, 3)
	if ivar48 != 1 {
		goto LABEL2
	}
	ns.StoryPic(ns.GetTrigger(), "WarriorPic")
	ns.SetDialog(ns.GetTrigger(), ns.NORMAL, fireKnightsTalk1Start, fireKnightsTalk1End)
LABEL2:
	if ivar48 != 2 {
		goto LABEL3
	}
	ns.StoryPic(ns.GetTrigger(), "Warrior2Pic")
	ns.SetDialog(ns.GetTrigger(), ns.NORMAL, fireKnightsTalk2Start, fireKnightsTalk2End)
LABEL3:
	if ivar48 != 3 {
		goto LABEL1
	}
	ns.StoryPic(ns.GetTrigger(), "Warrior3Pic")
	ns.SetDialog(ns.GetTrigger(), ns.NORMAL, fireKnightsTalk3Start, fireKnightsTalk3End)
LABEL1:
	return
}
func fireKnightsTalk1Start() {
	ns.Frozen(ns.GetHost(), true)
	ns.ObjectGroupOff(gvar15)
	ns.TellStory(ns.SwordsmanRecognize, "Con06a:CityGuardSpeak1")
}
func fireKnightsTalk1End() {
	ns.Frozen(ns.GetHost(), false)
	ns.ObjectGroupOn(gvar15)
	ns.SetCallback(ns.GetTrigger(), 3, NullCallback)
	ns.CancelDialog(ns.GetTrigger())
	ns.BecomePet(ns.GetTrigger())
	ns.CreatureFollow(ns.GetTrigger(), ns.GetHost())
	ns.AggressionLevel(ns.GetTrigger(), 0.83)
}
func fireKnightsTalk2Start() {
	ns.Frozen(ns.GetHost(), true)
	ns.ObjectGroupOff(gvar15)
	ns.TellStory(ns.SwordsmanRecognize, "Con06a:CityGuardSpeak2")
}
func fireKnightsTalk2End() {
	ns.Frozen(ns.GetHost(), false)
	ns.ObjectGroupOn(gvar15)
	ns.SetCallback(ns.GetTrigger(), 3, NullCallback)
	ns.CancelDialog(ns.GetTrigger())
	ns.BecomePet(ns.GetTrigger())
	ns.CreatureFollow(ns.GetTrigger(), ns.GetHost())
	ns.AggressionLevel(ns.GetTrigger(), 0.83)
}
func fireKnightsTalk3Start() {
	ns.Frozen(ns.GetHost(), true)
	ns.ObjectGroupOff(gvar15)
	ns.TellStory(ns.SwordsmanRecognize, "War06a:CityGuardSpeak3")
}
func fireKnightsTalk3End() {
	ns.Frozen(ns.GetHost(), false)
	ns.ObjectGroupOn(gvar15)
	ns.SetCallback(ns.GetTrigger(), 3, NullCallback)
	ns.CancelDialog(ns.GetTrigger())
	ns.CreatureFollow(ns.GetTrigger(), ns.GetHost())
	ns.BecomePet(ns.GetTrigger())
	ns.AggressionLevel(ns.GetTrigger(), 0.83)
}
func fireKnightsStayTalkInit() {
	if !(ns.IsCaller(obj18) || ns.IsCaller(obj19) || ns.IsCaller(obj40) || ns.IsCaller(obj41) || ns.IsCaller(obj42)) {
		goto LABEL1
	}
	if !ns.IsVisibleTo(ns.GetCaller(), ns.GetHost()) {
		goto LABEL2
	}
	ns.CreatureIdle(ns.GetCaller())
	ns.BecomeEnemy(ns.GetCaller())
	ns.CreatureGuard(ns.GetCaller(), ns.GetObjectX(ns.GetCaller()), ns.GetObjectY(ns.GetCaller()), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 150)
	ns.AggressionLevel(ns.GetCaller(), 0.5)
	ns.Wander(ns.GetCaller())
	if !flag49 {
		goto LABEL3
	}
	flag49 = false
	obj47 = ns.GetCaller()
	ns.SetDialog(ns.GetCaller(), ns.NORMAL, fireKnightsStayTalkStart, fireKnightsStayTalkEnd)
	ns.StartDialog(ns.GetCaller(), ns.GetHost())
LABEL3:
	goto LABEL4
LABEL2:
	ns.AggressionLevel(ns.GetCaller(), 0.5)
	ns.Wander(ns.GetCaller())
LABEL4:
	goto LABEL5
LABEL1:
	ns.AggressionLevel(ns.GetCaller(), 0.5)
	ns.Wander(ns.GetCaller())
LABEL5:
	return
}
func fireKnightsStayTalkStart() {
	ns.TellStory(ns.SwordsmanRecognize, "Con06:SaveDunMir06Conjurer")
}
func fireKnightsStayTalkEnd() {
	ns.CancelDialog(ns.GetTrigger())
}
func townsMenRecognize() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ivar48 = ns.Random(1, 3)
	if ivar48 != 1 {
		goto LABEL2
	}
	ns.StoryPic(ns.GetTrigger(), "Townsman1Pic")
	ns.SetDialog(ns.GetTrigger(), ns.NORMAL, townsmenTalk1Start, townsmenTalk1End)
LABEL2:
	if ivar48 != 2 {
		goto LABEL3
	}
	ns.StoryPic(ns.GetTrigger(), "Townsman2Pic")
	ns.SetDialog(ns.GetTrigger(), ns.NORMAL, townsmenTalk2Start, townsmenTalk2End)
LABEL3:
	if ivar48 != 3 {
		goto LABEL1
	}
	ns.StoryPic(ns.GetTrigger(), "Townsman3Pic")
	ns.SetDialog(ns.GetTrigger(), ns.NORMAL, townsmenTalk3Start, townsmenTalk3End)
LABEL1:
	return
}
func townsmenTalk1Start() {
	ns.Frozen(ns.GetHost(), true)
	ns.ObjectGroupOff(gvar15)
	ns.TellStory(ns.UrchinRecognize, "War06a:TownspeopleSpeak1")
}
func townsmenTalk1End() {
	ns.Frozen(ns.GetHost(), false)
	ns.ObjectGroupOn(gvar15)
}
func townsmenTalk2Start() {
	ns.Frozen(ns.GetHost(), true)
	ns.ObjectGroupOff(gvar15)
	ns.TellStory(ns.UrchinRecognize, "Con06a:TownspeopleSpeak2")
}
func townsmenTalk2End() {
	ns.Frozen(ns.GetHost(), false)
	ns.ObjectGroupOn(gvar15)
}
func townsmenTalk3Start() {
	ns.Frozen(ns.GetHost(), true)
	ns.ObjectGroupOff(gvar15)
	ns.TellStory(ns.UrchinRecognize, "War06a:TownspeopleSpeak3")
}
func townsmenTalk3End() {
	ns.Frozen(ns.GetHost(), false)
	ns.ObjectGroupOn(gvar15)
}
func townsWomenRecognize() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ivar48 = ns.Random(1, 3)
	if ivar48 != 1 {
		goto LABEL2
	}
	ns.StoryPic(ns.GetTrigger(), "MaidenPic")
	ns.SetDialog(ns.GetTrigger(), ns.NORMAL, townsWomenTalk1Start, townsWomenTalk1End)
LABEL2:
	if ivar48 != 2 {
		goto LABEL3
	}
	ns.StoryPic(ns.GetTrigger(), "MaidenPic2")
	ns.SetDialog(ns.GetTrigger(), ns.NORMAL, townsWomenTalk2Start, townsWomenTalk2End)
LABEL3:
	if ivar48 != 3 {
		goto LABEL1
	}
	ns.StoryPic(ns.GetTrigger(), "MaidenPic3")
	ns.SetDialog(ns.GetTrigger(), ns.NORMAL, townsWomenTalk3Start, townsWomenTalk3End)
LABEL1:
	return
}
func townsWomenTalk1Start() {
	ns.Frozen(ns.GetHost(), true)
	ns.ObjectGroupOff(gvar15)
	ns.TellStory(ns.GhostRecognize, "War06a:WomenSpeak1")
}
func townsWomenTalk1End() {
	ns.Frozen(ns.GetHost(), false)
	ns.ObjectGroupOn(gvar15)
}
func townsWomenTalk2Start() {
	ns.Frozen(ns.GetHost(), true)
	ns.ObjectGroupOff(gvar15)
	ns.TellStory(ns.GhostRecognize, "War06a:WomenSpeak1")
}
func townsWomenTalk2End() {
	ns.Frozen(ns.GetHost(), false)
	ns.ObjectGroupOn(gvar15)
}
func townsWomenTalk3Start() {
	ns.Frozen(ns.GetHost(), true)
	ns.ObjectGroupOff(gvar15)
	ns.TellStory(ns.GhostRecognize, "War06a:WomenSpeak1")
}
func townsWomenTalk3End() {
	ns.Frozen(ns.GetHost(), false)
	ns.ObjectGroupOn(gvar15)
}
func leavingDunMir() {
	if !(ns.IsCaller(obj50) || ns.IsCaller(obj51) || ns.IsCaller(obj52) || ns.IsCaller(obj53) || ns.IsCaller(obj54) || ns.IsCaller(obj55) || ns.IsCaller(obj56) || ns.IsCaller(obj57) || ns.IsCaller(obj58) || ns.IsCaller(obj59) || ns.IsCaller(obj60) || ns.IsCaller(obj61)) {
		goto LABEL1
	}
	ns.MoveObject(ns.GetCaller(), ns.GetWaypointX(wp62), ns.GetWaypointY(wp62))
	ns.ObjectOff(ns.GetCaller())
LABEL1:
	return
}
func flee() {
	ns.Wander(ns.GetTrigger())
}
func MakeAWish() {
	var v0 int
	v0 = ns.MaxHealth(ns.GetCaller())
	ns.RestoreHealth(ns.GetCaller(), v0-ns.CurrentHealth(ns.GetCaller()))
	ns.PrintToAll("GeneralPrint:WellSignRefresh")
	ns.AudioEvent(ns.RestoreHealthName, wp46)
}
func OnEvent(typ string) {
	switch typ {
	case "PlayerDeath":
		PlayerDeath()
	case "MapInitialize":
		MapInitialize()
	}
}
