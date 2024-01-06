package con07d

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	obj4   ns.ObjectID
	wp5    ns.WaypointID
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
	gvar18 ns.ObjectGroupID
	obj19  ns.ObjectID
	obj20  ns.ObjectID
	obj21  ns.ObjectID
	obj22  ns.ObjectID
	obj23  ns.ObjectID
	obj24  ns.ObjectID
	obj25  ns.ObjectID
	obj26  ns.ObjectID
	obj27  ns.ObjectID
	obj28  ns.ObjectID
	obj29  ns.ObjectID
	gvar30 ns.ObjectGroupID
	gvar31 ns.ObjectGroupID
	gvar32 ns.ObjectGroupID
	gvar33 ns.ObjectGroupID
	gvar34 ns.WallID
)

func CreatureSetup() {
	if !(ns.CurrentHealth(obj6) > 0) {
		goto LABEL1
	}
	ns.CreatureHunt(obj6)
LABEL1:
	if !(ns.CurrentHealth(obj7) > 0) {
		goto LABEL2
	}
	ns.CreatureHunt(obj7)
LABEL2:
	ns.SetOwner(ns.GetHost(), obj10)
	ns.SetOwner(ns.GetHost(), obj13)
	ns.SetOwner(ns.GetHost(), obj14)
	ns.SetOwner(ns.GetHost(), obj8)
	ns.SetOwner(ns.GetHost(), obj9)
	ns.ObjectOn(obj24)
	ns.ObjectOn(obj25)
	if !(ns.CurrentHealth(obj8) > 0) {
		goto LABEL3
	}
	ns.CreatureHunt(obj8)
LABEL3:
	if !(ns.CurrentHealth(obj9) > 0) {
		goto LABEL4
	}
	ns.CreatureHunt(obj9)
LABEL4:
	return
}
func StartMageA() {
	ns.ObjectGroupOn(gvar18)
}
func OpenBookcase1() {
	ns.ObjectOn(obj26)
	ns.ObjectOn(obj27)
	ns.WallOpen(gvar34)
}
func CornerSetPiece() {
	ns.ObjectOn(obj10)
	ns.ObjectGroupOn(gvar30)
	ns.ObjectOff(obj29)
	ns.HitFarLocation(obj11, ns.GetObjectX(obj10), ns.GetObjectY(obj10))
	ns.HitFarLocation(obj12, ns.GetObjectX(obj10), ns.GetObjectY(obj10))
}
func CorneredWizardDie() {
	ns.Chat(obj10, "War04a:GraveRobberDie")
}
func StartGearRoom() {
	if !(ns.CurrentHealth(obj13) > 0) {
		goto LABEL1
	}
	ns.Chat(obj14, "Con07C.scr:OpeningHelp")
	ns.ObjectOn(obj14)
LABEL1:
	ns.ObjectOn(obj13)
	ns.ObjectGroupOn(gvar31)
	ns.ObjectGroupOff(gvar33)
	CheckGearDemonHealth()
}
func ActivateGears() {
	if !ns.IsObjectOn(obj19) {
		goto LABEL1
	}
	ns.FrameTimer(1, StopGears)
	goto LABEL2
LABEL1:
	ns.FrameTimer(1, StartGears)
LABEL2:
	return
}
func StartGears() {
	ns.ObjectGroupOn(gvar32)
	ns.ObjectOn(obj20)
	ns.ObjectOn(obj21)
}
func StopGears() {
	ns.ObjectGroupOff(gvar32)
	ns.ObjectOff(obj20)
	ns.ObjectOff(obj21)
	ns.ObjectOff(obj22)
	ns.ObjectOff(obj23)
}
func CheckGearDemonHealth() {
	if !(ns.CurrentHealth(obj15) <= 0 && ns.CurrentHealth(obj16) <= 0 && ns.CurrentHealth(obj17) <= 0) {
		goto LABEL1
	}
	ns.CreatureGuard(obj13, ns.GetObjectX(obj13), ns.GetObjectY(obj13), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	ns.CreatureGuard(obj14, ns.GetObjectX(obj14), ns.GetObjectY(obj14), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	if !(ns.CurrentHealth(obj13) > 0) {
		goto LABEL2
	}
	ns.SetDialog(obj13, ns.NORMAL, GearWizStart, GearWizEnd)
	ns.StoryPic(obj13, "WizardGuard1Pic")
	goto LABEL3
LABEL2:
	if !(ns.CurrentHealth(obj14) > 0) {
		goto LABEL3
	}
	ns.SetDialog(obj14, ns.NORMAL, GearWizStart, GearWizEnd)
	ns.StoryPic(obj14, "WizardGuard2Pic")
LABEL3:
	goto LABEL4
LABEL1:
	ns.FrameTimer(15, CheckGearDemonHealth)
LABEL4:
	return
}
func GearWizStart() {
	ns.TellStory(ns.HumanMaleEatFood, "Con07:MiscMage02")
}
func GearWizEnd() {
	ns.CancelDialog(obj13)
	ns.CancelDialog(obj14)
}
func FollowPlayer() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	if !(ns.CurrentHealth(ns.GetTrigger()) > 0) {
		goto LABEL1
	}
	ns.CreatureFollow(ns.GetTrigger(), ns.GetHost())
LABEL1:
	return
}
func CancelBubble() {
	ns.DestroyChat(ns.GetTrigger())
}
func MapInitialize() {
	obj8 = ns.Object("PCMageA")
	obj9 = ns.Object("PCMageB")
	obj10 = ns.Object("CorneredWizard")
	obj11 = ns.Object("CDemon1")
	obj12 = ns.Object("CDemon2")
	obj13 = ns.Object("Ganem")
	obj14 = ns.Object("Kelvin")
	obj15 = ns.Object("GearDemon1")
	obj16 = ns.Object("GearDemon2")
	obj17 = ns.Object("GearDemon3")
	gvar18 = ns.ObjectGroup("MageAGroup")
	obj19 = ns.Object("TestGear")
	obj20 = ns.Object("Globe1")
	obj21 = ns.Object("Globe2")
	obj22 = ns.Object("Globe3")
	obj23 = ns.Object("Globe4")
	obj24 = ns.Object("Globe1Mover")
	obj25 = ns.Object("Globe2Mover")
	obj26 = ns.Object("Bookcase1Mover")
	obj27 = ns.Object("Bookcase2Mover")
	obj4 = ns.Object("ToLevelOneTP")
	obj28 = ns.Object("ToThreeTP")
	obj29 = ns.Object("CornerTrigger")
	gvar30 = ns.ObjectGroup("CornerDemons")
	gvar31 = ns.ObjectGroup("GearDemons")
	gvar32 = ns.ObjectGroup("Level2Gears")
	gvar33 = ns.ObjectGroup("StartGearRoomTriggers")
	gvar34 = ns.Wall(142, 54)
	wp5 = ns.Waypoint("PlayerSounds")
	ns.ObjectOff(obj10)
	ns.ObjectOff(obj13)
	ns.ObjectOff(obj14)
	ns.ObjectGroupOff(gvar31)
	ns.ObjectGroupOff(gvar30)
	ns.ObjectGroupOff(gvar18)
	ns.SetOwner(ns.GetHost(), obj10)
	ns.SetOwner(ns.GetHost(), obj13)
	ns.SetOwner(ns.GetHost(), obj14)
}
func PlayerDeath() {
	ns.DeathScreen(7)
}
func MapEntry() {
	ns.Music(27, 100)
	ns.FrameTimer(1, CreatureSetup)
}
func OnEvent(typ string) {
	switch typ {
	case "MapInitialize":
		MapInitialize()
	case "PlayerDeath":
		PlayerDeath()
	case "MapEntry":
		MapEntry()
	}
}
