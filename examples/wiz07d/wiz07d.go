package wiz07d

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
	gvar16 int
	gvar17 int
	gvar18 int
)

func init() {
	gvar16 = 0
	gvar17 = 1
}
func CreatureSetup() {
	ns.SetOwner(ns.GetHost(), obj14)
	ns.SetOwner(ns.GetHost(), obj15)
	ns.SetOwner(ns.GetHost(), obj11)
	ns.SetOwner(ns.GetHost(), obj12)
	ns.Wander(obj15)
	ns.Music(15, 100)
}
func ReceptionTalkStart() {
	var (
		v0 int
		v1 int
		v2 int
	)
	v2 = gvar18
	if v2 == gvar16 {
		goto LABEL1
	}
	if v2 == gvar17 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.TellStory(ns.HumanMaleEatFood, "Wiz07:Reception01")
	gvar18 = gvar17
	goto LABEL3
LABEL2:
	v0 = ns.Random(1, 2)
	v1 = v0
	if v1 == 1 {
		goto LABEL4
	}
	if v1 == 2 {
		goto LABEL5
	}
	goto LABEL3
LABEL4:
	ns.TellStory(ns.HumanMaleEatFood, "Wiz07:Reception02")
	goto LABEL3
LABEL5:
	ns.TellStory(ns.HumanMaleEatFood, "Wiz07:Reception03")
	goto LABEL3
LABEL3:
	return
}
func ReceptionTalkEnd() {
	ns.SetDialog(obj14, ns.NORMAL, ReceptionTalkStart, ReceptionTalkEnd)
}
func GanemStart() {
	ns.LookAtObject(obj11, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "Wiz07:Ganem01")
}
func GanemEnd() {
}
func MapInitialize() {
	obj14 = ns.Object("Main_Receptionist")
	obj15 = ns.Object("Lewis")
	obj11 = ns.Object("Ganem")
	obj12 = ns.Object("Mystic")
	obj4 = ns.Object("NorthDoor1")
	obj5 = ns.Object("NorthDoor2")
	obj6 = ns.Object("SouthDoorUp1")
	obj7 = ns.Object("SouthDoorUp2")
	obj8 = ns.Object("ClassDoor")
	obj9 = ns.Object("ArchiveGate1")
	obj10 = ns.Object("ArchiveGate2")
	obj13 = ns.Object("LabEntrance")
	ns.LockDoor(obj4)
	ns.LockDoor(obj5)
	ns.LockDoor(obj6)
	ns.LockDoor(obj7)
	ns.LockDoor(obj8)
	ns.LockDoor(obj9)
	ns.LockDoor(obj10)
	ns.LockDoor(obj13)
	gvar18 = gvar16
	ns.SetDialog(obj14, ns.NORMAL, ReceptionTalkStart, ReceptionTalkEnd)
	ns.StoryPic(obj14, "WizardGuard2Pic")
	ns.StoryPic(obj11, "WizardGuard1Pic")
	ns.SetDialog(obj11, ns.NORMAL, GanemStart, GanemEnd)
	ns.FrameTimer(1, CreatureSetup)
}
func PlayerDeath() {
	ns.DeathScreen(10)
}
func OnEvent(typ string) {
	switch typ {
	case "MapInitialize":
		MapInitialize()
	case "PlayerDeath":
		PlayerDeath()
	}
}
