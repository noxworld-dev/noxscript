package war03a

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
	obj20  ns.ObjectID
	obj21  ns.ObjectID
	gvar22 ns.ObjectGroupID
	obj23  ns.ObjectID
	obj24  ns.ObjectID
	obj25  ns.ObjectID
	obj26  ns.ObjectID
	gvar27 ns.WallGroupID
	gvar28 int
	gvar29 int
	gvar30 int
	gvar31 int
	gvar32 int
	gvar33 int
	gvar34 int
	gvar35 int
	gvar36 int
	gvar37 int
	gvar38 int
	gvar39 int
	gvar40 ns.ObjectGroupID
	gvar41 ns.ObjectGroupID
	gvar42 ns.ObjectGroupID
	obj43  ns.ObjectID
	obj44  ns.ObjectID
	wp45   ns.WaypointID
	wp46   ns.WaypointID
	wp47   ns.WaypointID
	flag48 bool
)

func init() {
	gvar28 = 0
	gvar29 = 1
	gvar30 = gvar28
	gvar31 = 0
	gvar32 = 1
	gvar33 = gvar31
	gvar34 = 0
	gvar35 = 1
	gvar36 = gvar34
	gvar37 = 0
	gvar38 = 1
	gvar39 = gvar37
	flag48 = true
}
func PlayerDeath() {
	ns.DeathScreen(3)
}
func BanditAttack() {
	ns.Chat(obj26, "War01A.scr:Bully1")
}
func HorrendousTalkStart() {
	var v0 int
	v0 = gvar30
	if v0 == gvar28 {
		goto LABEL1
	}
	if v0 == gvar29 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.TellStory(ns.SwordsmanHurt, "War03a:HorrendousSpeech")
	gvar30 = gvar29
	goto LABEL3
LABEL2:
	ns.TellStory(ns.SwordsmanHurt, "War03a:HorrendousEnd")
	goto LABEL3
LABEL3:
	return
}
func HorrendousTalkEnd() {
	ns.JournalEntry(ns.GetHost(), "War03aIxQuest", 2)
}
func DunMirGuard1TalkStart() {
	ns.LookAtObject(obj8, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War03a:DunMirGuard1")
}
func DunMirGuard1TalkEnd() {
}
func DunMirGuard2TalkStart() {
	ns.LookAtObject(obj9, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War03a:DunMirGuard2")
}
func DunMirGuard2TalkEnd() {
}
func DunMirGuard3TalkStart() {
	ns.TellStory(ns.SwordsmanHurt, "War03a:DunMirGuard3")
}
func DunMirGuard3TalkEnd() {
}
func MineGuardTalkStart() {
	ns.LookAtObject(obj15, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War03a:MineGuard")
}
func MineGuardTalkEnd() {
}
func GalavaGuard1TalkStart() {
	var v0 int
	ns.LookAtObject(obj11, ns.GetHost())
	v0 = gvar33
	if v0 == gvar31 {
		goto LABEL1
	}
	if v0 == gvar32 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.TellStory(ns.SwordsmanHurt, "War03a:GalavaGuardWarn")
	goto LABEL3
LABEL2:
	ns.TellStory(ns.SwordsmanHurt, "War03a:GalavaGuardEnd")
	goto LABEL3
LABEL3:
	return
}
func GalavaGuard1TalkEnd() {
	var v0 int
	v0 = gvar33
	if v0 == gvar31 {
		goto LABEL1
	}
	if v0 == gvar32 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	goto LABEL3
LABEL2:
	goto LABEL3
LABEL3:
	return
}
func GalavaGuard2TalkStart() {
	var v0 int
	ns.LookAtObject(obj12, ns.GetHost())
	v0 = gvar33
	if v0 == gvar31 {
		goto LABEL1
	}
	if v0 == gvar32 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.TellStory(ns.SwordsmanHurt, "War03a:GalavaGuardEnd")
	goto LABEL3
LABEL2:
	ns.TellStory(ns.SwordsmanHurt, "War03a:GalavaGuardEnd")
	goto LABEL3
LABEL3:
	return
}
func GalavaGuard2TalkEnd() {
	var v0 int
	v0 = gvar33
	if v0 == gvar31 {
		goto LABEL1
	}
	if v0 == gvar32 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	goto LABEL3
LABEL2:
	goto LABEL3
LABEL3:
	return
}
func IxGuard1TalkStart() {
	var v0 int
	ns.LookAtObject(obj18, ns.GetHost())
	v0 = gvar36
	if v0 == gvar34 {
		goto LABEL1
	}
	if v0 == gvar35 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.TellStory(ns.SwordsmanHurt, "War03a:IxGuard1Intro")
	goto LABEL3
LABEL2:
	ns.TellStory(ns.SwordsmanHurt, "War03a:IxGuard1End")
	goto LABEL3
LABEL3:
	return
}
func IxGuard1TalkEnd() {
	var v0 int
	v0 = gvar36
	if v0 == gvar34 {
		goto LABEL1
	}
	if v0 == gvar35 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.UnlockDoor(obj20)
	ns.UnlockDoor(obj21)
	gvar36 = gvar35
	gvar39 = gvar38
	goto LABEL3
LABEL2:
	goto LABEL3
LABEL3:
	return
}
func IxGuard2TalkStart() {
	var v0 int
	ns.LookAtObject(obj19, ns.GetHost())
	v0 = gvar39
	if v0 == gvar37 {
		goto LABEL1
	}
	if v0 == gvar38 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.TellStory(ns.SwordsmanHurt, "War03a:IxGuard2Intro")
	goto LABEL3
LABEL2:
	ns.TellStory(ns.SwordsmanHurt, "War03a:IxGuard2End")
	goto LABEL3
LABEL3:
	return
}
func IxGuard2TalkEnd() {
	var v0 int
	v0 = gvar39
	if v0 == gvar37 {
		goto LABEL1
	}
	if v0 == gvar38 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.UnlockDoor(obj20)
	ns.UnlockDoor(obj21)
	gvar39 = gvar38
	gvar36 = gvar35
	goto LABEL3
LABEL2:
	goto LABEL3
LABEL3:
	return
}
func PlayWanderMusic() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.Music(22, 100)
LABEL1:
	return
}
func MapEntry() {
	PlayWanderMusic()
	ns.UnBlind()
	if ns.GetQuestStatus("War03b:QuestComplete") != 1 {
		goto LABEL1
	}
	ns.CancelDialog(obj18)
	ns.CancelDialog(obj19)
LABEL1:
	return
}
func MapInitialize() {
	obj4 = ns.Object("MissionTrig")
	obj8 = ns.Object("Lance")
	obj9 = ns.Object("Kenneth")
	obj10 = ns.Object("Brigadin")
	obj11 = ns.Object("Rastur")
	obj12 = ns.Object("Kirik")
	obj13 = ns.Object("GalavaGate1")
	obj14 = ns.Object("GalavaGate2")
	ns.LockDoor(obj13)
	ns.LockDoor(obj14)
	obj15 = ns.Object("Millard")
	obj16 = ns.Object("MineDoor1")
	obj17 = ns.Object("MineDoor2")
	ns.LockDoor(obj16)
	ns.LockDoor(obj17)
	gvar22 = ns.ObjectGroup("GuardGroup")
	obj18 = ns.Object("Janero")
	obj19 = ns.Object("Horst")
	obj20 = ns.Object("IxGate01")
	obj21 = ns.Object("IxGate02")
	ns.LockDoor(obj20)
	ns.LockDoor(obj21)
	obj23 = ns.Object("Loproc")
	obj24 = ns.Object("BiffordByzanti")
	obj25 = ns.Object("GarretByzanti")
	ns.GroupSetOwner(ns.GetHost(), gvar22)
	ns.SetOwner(ns.GetHost(), obj15)
	ns.SetOwner(ns.GetHost(), obj10)
	gvar40 = ns.ObjectGroup("WaspGroup1")
	gvar41 = ns.ObjectGroup("WaspGroup2")
	obj43 = ns.Object("Wolf01")
	gvar42 = ns.ObjectGroup("FishGroup")
	ns.GroupWander(gvar40)
	ns.GroupWander(gvar41)
	ns.Wander(obj43)
	ns.GroupWander(gvar42)
	obj44 = ns.Object("Osborn")
	ns.Damage(obj44, 0, 200, 2)
	ns.StoryPic(obj5, "HorrendousPic")
	ns.SetDialog(obj5, ns.NORMAL, HorrendousTalkStart, HorrendousTalkEnd)
	ns.StoryPic(obj8, "Warrior2Pic")
	ns.SetDialog(obj8, ns.NORMAL, DunMirGuard1TalkStart, DunMirGuard1TalkEnd)
	ns.StoryPic(obj9, "Warrior3Pic")
	ns.SetDialog(obj9, ns.NORMAL, DunMirGuard2TalkStart, DunMirGuard2TalkEnd)
	ns.StoryPic(obj10, "Warrior4Pic")
	ns.SetDialog(obj10, ns.NORMAL, DunMirGuard3TalkStart, DunMirGuard3TalkEnd)
	ns.StoryPic(obj11, "WizardGuard1Pic")
	ns.SetDialog(obj11, ns.NORMAL, GalavaGuard1TalkStart, GalavaGuard1TalkEnd)
	ns.StoryPic(obj12, "WizardGuard2Pic")
	ns.SetDialog(obj12, ns.NORMAL, GalavaGuard2TalkStart, GalavaGuard2TalkEnd)
	ns.StoryPic(obj15, "MalePic2")
	ns.SetDialog(obj15, ns.NORMAL, MineGuardTalkStart, MineGuardTalkEnd)
	ns.StoryPic(obj18, "IxGuard1Pic")
	ns.SetDialog(obj18, ns.NORMAL, IxGuard1TalkStart, IxGuard1TalkEnd)
	ns.StoryPic(obj19, "IxGuard2Pic")
	ns.SetDialog(obj19, ns.NORMAL, IxGuard2TalkStart, IxGuard2TalkEnd)
	obj6 = ns.Object("DMDoorL")
	obj7 = ns.Object("DMDoorR")
	ns.LockDoor(obj6)
	ns.LockDoor(obj7)
	obj26 = ns.Object("Bandit01")
	gvar27 = ns.WallGroup("RogueSecretWallGroup")
	wp45 = ns.Waypoint("Secret01WP")
	wp46 = ns.Waypoint("Secret01aWP")
	wp47 = ns.Waypoint("Secret02WP")
}
func SetHorrendous() {
	ns.ObjectOff(obj5)
	ns.MoveObject(obj5, 5569, 5071)
}
func ClearIxGuardDialog() {
	ns.CancelDialog(obj18)
	ns.CancelDialog(obj19)
}
func OpenRogueSecretWallGroup() {
	ns.WallGroupOpen(gvar27)
}
func WolfIdle() {
	ns.CreatureIdle(obj43)
}
func WaspIdle() {
	ns.CreatureGroupIdle(gvar41)
}
func MonsterGoHome() {
	r1 := ns.IsAttackedBy(ns.GetCaller(), ns.GetHost())
	if !r1 {
		goto LABEL1
	}
	ns.GoBackHome(ns.GetCaller())
LABEL1:
	return
}
func PlaySubMusic() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.Music(20, 100)
LABEL1:
	return
}
func PlayActionMusic() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.Music(5, 100)
LABEL1:
	return
}
func RobberRecognize() {
	if !(ns.IsCaller(ns.GetHost()) && flag48) {
		goto LABEL1
	}
	flag48 = false
	ns.Attack(ns.GetTrigger(), ns.GetHost())
	ns.Chat(ns.GetTrigger(), "War01A.scr:Bully1")
LABEL1:
	return
}
func RobberDie() {
	ns.DestroyEveryChat()
}
func Secret01() {
	ns.AudioEvent(ns.SecretFound, wp45)
	ns.AudioEvent(ns.SecretFound, wp46)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 50)
}
func Secret02() {
	ns.AudioEvent(ns.SecretFound, wp47)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 100)
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
