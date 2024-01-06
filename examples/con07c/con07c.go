package con07c

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
	obj22  ns.ObjectID
	gvar23 ns.ObjectGroupID
	gvar24 ns.ObjectGroupID
	obj25  ns.ObjectID
	obj26  ns.ObjectID
	gvar27 ns.ObjectGroupID
	gvar28 ns.ObjectGroupID
	gvar29 ns.ObjectGroupID
	wp30   ns.WaypointID
)

func CreatureSetup() {
	if !(ns.CurrentHealth(obj4) > 0) {
		goto LABEL1
	}
	ns.CreatureHunt(obj4)
LABEL1:
	if !(ns.CurrentHealth(obj5) > 0) {
		goto LABEL2
	}
	ns.CreatureHunt(obj5)
LABEL2:
	if !(ns.CurrentHealth(obj6) > 0) {
		goto LABEL3
	}
	ns.CreatureHunt(obj6)
LABEL3:
	if !(ns.CurrentHealth(obj7) > 0) {
		goto LABEL4
	}
	ns.CreatureHunt(obj7)
LABEL4:
	if !(ns.CurrentHealth(obj8) > 0) {
		goto LABEL5
	}
	ns.CreatureHunt(obj8)
LABEL5:
	if !(ns.CurrentHealth(obj9) > 0) {
		goto LABEL6
	}
	ns.CreatureHunt(obj9)
LABEL6:
	if !(ns.CurrentHealth(obj10) > 0) {
		goto LABEL7
	}
	ns.CreatureHunt(obj10)
LABEL7:
	if !(ns.CurrentHealth(obj11) > 0) {
		goto LABEL8
	}
	ns.CreatureHunt(obj11)
LABEL8:
	if !(ns.CurrentHealth(obj12) > 0) {
		goto LABEL9
	}
	ns.CreatureHunt(obj12)
LABEL9:
	if !(ns.CurrentHealth(obj13) > 0) {
		goto LABEL10
	}
	ns.CreatureHunt(obj13)
LABEL10:
	if !(ns.CurrentHealth(obj14) > 0) {
		goto LABEL11
	}
	ns.CreatureHunt(obj14)
LABEL11:
	if !(ns.CurrentHealth(obj15) > 0) {
		goto LABEL12
	}
	ns.CreatureHunt(obj15)
LABEL12:
	if !(ns.CurrentHealth(obj16) > 0) {
		goto LABEL13
	}
	ns.CreatureHunt(obj16)
LABEL13:
	if !(ns.CurrentHealth(obj17) > 0) {
		goto LABEL14
	}
	ns.CreatureHunt(obj17)
LABEL14:
	if !(ns.CurrentHealth(obj18) > 0) {
		goto LABEL15
	}
	ns.CreatureHunt(obj18)
LABEL15:
	if !(ns.CurrentHealth(obj19) > 0) {
		goto LABEL16
	}
	ns.CreatureHunt(obj19)
LABEL16:
	if !(ns.CurrentHealth(obj20) > 0) {
		goto LABEL17
	}
	ns.CreatureHunt(obj20)
LABEL17:
	ns.SetOwner(ns.GetHost(), obj10)
	ns.SetOwner(ns.GetHost(), obj11)
	ns.SetOwner(ns.GetHost(), obj12)
	ns.SetOwner(ns.GetHost(), obj13)
	ns.SetOwner(ns.GetHost(), obj14)
	ns.SetOwner(ns.GetHost(), obj15)
	ns.SetOwner(ns.GetHost(), obj16)
	ns.SetOwner(ns.GetHost(), obj17)
	ns.SetOwner(ns.GetHost(), obj18)
	ns.SetOwner(ns.GetHost(), obj19)
	ns.SetOwner(ns.GetHost(), obj20)
}
func FirstSetPiece() {
	r1 := ns.IsAttackedBy(ns.GetHost(), ns.GetCaller())
	if r1 {
		goto LABEL1
	}
	ns.WideScreen(true)
	ns.Frozen(ns.GetHost(), true)
	ns.AggressionLevel(obj10, 0)
	ns.ObjectOn(obj10)
	ns.Move(obj10, wp30)
	ns.LookAtObject(ns.GetHost(), obj10)
LABEL1:
	return
}
func PCMageTalk() {
	ns.Frozen(obj10, true)
	ns.CreatureIdle(obj10)
	ns.LookAtObject(obj10, ns.GetHost())
	ns.SetDialog(obj10, ns.NORMAL, PCMageStart, PCMageEnd)
	ns.StoryPic(obj10, "WizardGuard4Pic")
	ns.StartDialog(obj10, ns.GetHost())
}
func EndSetPiece() {
	ns.Frozen(ns.GetHost(), false)
	ns.Frozen(obj10, false)
	ns.ObjectOn(obj10)
	ns.ObjectGroupOn(gvar29)
	ns.CreatureHunt(obj10)
	ns.WideScreen(false)
	ns.ObjectGroupOff(gvar28)
}
func StartTrapWiz() {
	ns.ObjectGroupOn(gvar27)
}
func StartCenter() {
	ns.ObjectGroupOff(gvar24)
	ns.ObjectGroupOn(gvar23)
	if !(ns.CurrentHealth(obj15) > 0) {
		goto LABEL1
	}
	ns.CreatureHunt(obj15)
LABEL1:
	if !(ns.CurrentHealth(obj16) > 0) {
		goto LABEL2
	}
	ns.CreatureHunt(obj16)
LABEL2:
	if !(ns.CurrentHealth(obj17) > 0) {
		goto LABEL3
	}
	ns.CreatureHunt(obj17)
LABEL3:
	return
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
func PCMageStart() {
	ns.TellStory(ns.HumanMaleEatFood, "Con07:MiscMage03")
}
func PCMageEnd() {
	ns.CancelDialog(obj10)
	EndSetPiece()
}
func MapInitialize() {
	obj4 = ns.Object("OgreHunter1")
	obj5 = ns.Object("OgreHunter2")
	obj6 = ns.Object("OgreHunter3")
	obj7 = ns.Object("OgreHunter4")
	obj8 = ns.Object("OgreHunter5")
	obj9 = ns.Object("OgreHunter6")
	obj10 = ns.Object("Ganem")
	obj11 = ns.Object("PCMage2")
	obj12 = ns.Object("PCMage3")
	obj13 = ns.Object("PCMage4")
	obj14 = ns.Object("PCMage5")
	obj15 = ns.Object("PCMage6")
	obj16 = ns.Object("PCMage7")
	obj17 = ns.Object("PCMage8")
	obj18 = ns.Object("PCMage9")
	obj19 = ns.Object("PCMage10")
	obj20 = ns.Object("PCMage11")
	obj21 = ns.Object("SetOgre1")
	obj22 = ns.Object("SetOgre2")
	gvar23 = ns.ObjectGroup("CenterSetPiece")
	gvar24 = ns.ObjectGroup("CenterTriggers")
	obj25 = ns.Object("ArchiveGate1")
	obj26 = ns.Object("ArchiveGate2")
	gvar27 = ns.ObjectGroup("KeyRoomSetPiece")
	gvar28 = ns.ObjectGroup("FirstSetPieceTriggers")
	gvar29 = ns.ObjectGroup("AllOgres")
	wp30 = ns.Waypoint("PCMage1WP")
	ns.ObjectGroupOff(gvar27)
	ns.ObjectGroupOff(gvar29)
}
func MapEntry() {
	ns.WideScreen(false)
	ns.Music(27, 100)
	ns.FrameTimer(1, CreatureSetup)
}
func PlayerDeath() {
	ns.DeathScreen(7)
}
func OnEvent(typ string) {
	switch typ {
	case "MapInitialize":
		MapInitialize()
	case "MapEntry":
		MapEntry()
	case "PlayerDeath":
		PlayerDeath()
	}
}
