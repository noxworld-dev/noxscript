package war07b

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
	gvar19 ns.ObjectGroupID
	gvar20 ns.ObjectGroupID
	gvar21 ns.ObjectGroupID
	gvar22 ns.ObjectGroupID
	gvar23 ns.WallID
	gvar24 ns.WallGroupID
	gvar25 ns.WallGroupID
	wp26   ns.WaypointID
	flag27 bool
	flag28 bool
	flag29 bool
	flag30 bool
)

func init() {
	flag27 = false
	flag28 = false
	flag29 = false
	flag30 = false
}
func EnableReceptionist() {
	ns.ObjectOn(obj5)
	if !(ns.CurrentHealth(obj5) > 0) {
		goto LABEL1
	}
	ns.Chat(obj5, "War07B.scr:MainReceptionistTalk01")
LABEL1:
	ns.ObjectGroupOff(gvar22)
	ns.CreatureHunt(obj11)
	ns.CreatureHunt(obj12)
	ns.CreatureHunt(obj13)
}
func ClassroomAttacks() {
	ns.ObjectGroupOn(gvar19)
	ns.ObjectOn(obj6)
}
func BirdManAttacks() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	if flag29 {
		goto LABEL1
	}
	ns.Chat(obj7, "War07B.scr:BirdManTalk01")
	flag29 = true
LABEL1:
	return
}
func OpenLabBookcase() {
	ns.ObjectOn(obj16)
	ns.ObjectOn(obj17)
	ns.WallOpen(gvar23)
	ns.WallGroupBreak(gvar24)
	ns.ObjectGroupOff(gvar21)
	ns.GroupDelete(gvar21)
	ns.ObjectOff(obj15)
	flag28 = true
}
func DestroyLab() {
	if flag28 {
		goto LABEL1
	}
	ns.Chat(obj8, "War07B.scr:TempWizardTalk01")
	ns.ObjectGroupOn(gvar20)
	ns.WallGroupOpen(gvar25)
	ns.Delete(obj14)
	flag28 = true
LABEL1:
	return
}
func EnterClassroom() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	if !(ns.CurrentHealth(obj11) > 0) {
		goto LABEL2
	}
	ns.Frozen(obj11, true)
LABEL2:
	if !(ns.CurrentHealth(obj12) > 0) {
		goto LABEL3
	}
	ns.Frozen(obj12, true)
LABEL3:
	if !(ns.CurrentHealth(obj13) > 0) {
		goto LABEL4
	}
	ns.Frozen(obj13, true)
LABEL4:
	if !(ns.CurrentHealth(obj6) > 0) {
		goto LABEL5
	}
	ns.Frozen(obj6, true)
	ns.StartDialog(obj6, ns.GetHost())
	goto LABEL1
LABEL5:
	ClassroomAttacks()
LABEL1:
	return
}
func TeacherStart() {
	ns.TellStory(ns.HumanMaleEatFood, "War07B.scr:TeacherTalk01")
}
func TeacherEnd() {
	ns.CancelDialog(obj6)
	if !(ns.CurrentHealth(obj11) > 0) {
		goto LABEL1
	}
	ns.Frozen(obj11, false)
LABEL1:
	if !(ns.CurrentHealth(obj12) > 0) {
		goto LABEL2
	}
	ns.Frozen(obj12, false)
LABEL2:
	if !(ns.CurrentHealth(obj13) > 0) {
		goto LABEL3
	}
	ns.Frozen(obj13, false)
LABEL3:
	ns.Frozen(obj6, false)
	ClassroomAttacks()
}
func MapInitialize() {
	obj4 = ns.Object("ChamberSpawn")
	obj5 = ns.Object("MainReceptionist")
	obj6 = ns.Object("Teacher")
	obj7 = ns.Object("BirdMan")
	obj8 = ns.Object("TempWizard")
	obj9 = ns.Object("TrapWiz1")
	obj10 = ns.Object("TrapWiz2")
	obj11 = ns.Object("MageHunter1")
	obj12 = ns.Object("MageHunter2")
	obj13 = ns.Object("MageHunter3")
	obj14 = ns.Object("FloorOneMagicAxe")
	obj15 = ns.Object("LabExplodeTrigger")
	obj16 = ns.Object("LabBookcaseMover1")
	obj17 = ns.Object("LabBookcaseMover2")
	obj18 = ns.Object("ToLevelTwoTP")
	gvar19 = ns.ObjectGroup("Students")
	gvar20 = ns.ObjectGroup("LabFlames")
	gvar21 = ns.ObjectGroup("ExplodingBarrels")
	gvar22 = ns.ObjectGroup("ReceptionTriggers")
	gvar23 = ns.Wall(123, 191)
	gvar24 = ns.WallGroup("LabWalls")
	gvar25 = ns.WallGroup("LabSecretWalls2")
	wp26 = ns.Waypoint("PlayerSounds")
	ns.ObjectOff(obj5)
	ns.ObjectOff(obj6)
	ns.ObjectGroupOff(gvar19)
	ns.ObjectGroupOff(gvar20)
	ns.StoryPic(obj6, "WizardGuard2Pic")
	ns.SetDialog(obj6, ns.NORMAL, TeacherStart, TeacherEnd)
}
func PlayerDeath() {
	ns.DeathScreen(7)
}
func MapEntry() {
	ns.Music(16, 100)
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
